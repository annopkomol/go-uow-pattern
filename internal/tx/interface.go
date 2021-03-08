package tx

import "gorm.io/gorm"

type ARepo interface {
	SetTx(tx Tx) ARepo
	Find(id uint) A
	Update(msg string) error
}

type BRepo interface {
	SetTx(tx Tx) BRepo
	Find(id uint) B
}

type UOW interface {
	Process(func(tx Tx) error) error
}

type Tx interface {
}

type A struct {
	gorm.Model
	Msg string
}

type B struct {
	gorm.Model
}
