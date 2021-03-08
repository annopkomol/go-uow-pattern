package tx

import "gorm.io/gorm"

type ARepoImp struct {
	*gorm.DB
}

func (a *ARepoImp) Update(msg string) error {
	return a.DB.First(&A{}, 1).Update("msg", msg).Error
}

func (a *ARepoImp) SetTx(tx Tx) ARepo {
	return &ARepoImp{
		DB: tx.(*gorm.DB),
	}
}

func (a *ARepoImp) Find(id uint) (model A) {
	a.DB.First(&model, id)
	return
}

type BRepoImp struct {
	*gorm.DB
}

func (b *BRepoImp) SetTx(tx Tx) BRepo {
	return &BRepoImp{DB: tx.(*gorm.DB)}
}

func (b *BRepoImp) Find(id uint) (model B) {
	b.DB.First(&model, id)
	return
}

type UOWImp struct {
	*gorm.DB
}

func (u *UOWImp) Process(callback func(Tx) error) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		return callback(tx)
	})
}
