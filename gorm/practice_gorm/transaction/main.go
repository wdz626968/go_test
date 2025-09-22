package main

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := initDB()
	if err != nil {
	}
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	accountA := Account{
		Model:   gorm.Model{},
		Name:    "dzhwu",
		Balance: 100,
	}
	db.Save(&accountA)
	accountB := Account{
		Model:   gorm.Model{},
		Name:    "mia",
		Balance: 0,
	}
	db.Save(&accountB)
	transfer(db, accountA.Name, accountB.Name, 100)
}

func transfer(db *gorm.DB, fromName string, toName string, amount uint) {
	fromAccount := Account{}
	toAccount := Account{}
	db.Model(&Account{}).Select("*").Where("name = ?", fromName).First(&fromAccount)
	db.Model(&Account{}).Select("*").Where("name = ?", toName).First(&toAccount)

	db.Transaction(func(tx *gorm.DB) error {
		if fromAccount.Balance < amount {
			return errors.New("not enough balance")
		}
		fromAccount.Balance -= amount

		t := tx.Model(&fromAccount).Where("name", fromName).Update("balance", fromAccount.Balance)
		if t.Error != nil {
			return t.Error
		}
		toAccount.Balance += amount
		update := tx.Model(&toAccount).Where("name", toName).Update("balance", toAccount.Balance)
		if update.Error != nil {
			return update.Error
		}

		transaction := Transaction{
			Model:         gorm.Model{},
			FromAccountId: fromAccount.ID,
			ToAccountId:   toAccount.ID,
			Amount:        amount,
		}
		create := tx.Debug().Create(&transaction)
		if create.Error != nil {
			return create.Error
		}
		return nil
	})

}

type Account struct {
	gorm.Model
	Balance uint `gorm:"type:uint"`
	Name    string
}

type Transaction struct {
	gorm.Model
	FromAccountId uint
	ToAccountId   uint
	Amount        uint
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
