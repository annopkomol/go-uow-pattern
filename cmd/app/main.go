package main

import (
	"errors"
	"fmt"
	"github.com/annopkomol/go-uow-pattern/internal/tx"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	dsn := "database connection string"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&tx.A{}, &tx.B{})

	uow := tx.UOWImp{
		DB: db,
	}
	repoA := tx.ARepoImp{
		DB: db,
	}
	repoB := tx.BRepoImp{
		DB: db,
	}

	err = uow.Process(func(tx tx.Tx) error {
		a := repoA.SetTx(tx)
		b := repoB.SetTx(tx)
		a.Find(1)
		err := a.Update("hi hi")
		if err != nil {
			return err
		}
		b.Find(1)
		return errors.New("some error -> must rollback tx")

	})
	if err != nil {
		fmt.Println(err)
	}

}
