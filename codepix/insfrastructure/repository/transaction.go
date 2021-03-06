package repository

import "gorm.io/'gorm"

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDb) Register(transaction *Transaction) error {
	err := r.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDb) Save(transaction *Transaction) error {
	err := r.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDb) Find(id string) (*Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
