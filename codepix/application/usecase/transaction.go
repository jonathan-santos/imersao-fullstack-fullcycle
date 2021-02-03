package usecase

import (
	"errors"
	"github.com/jonathan-santos/imersao-fullstack-fullcycle/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository      model.PixKeyRepositoryInterface
}

func (t* TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) error {
	account, err := t.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := t.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	t.TransactionRepository.Save(transaction)
	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to processs this transaction")
}

func (t* TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found",transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err := t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t* TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found",transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err := t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t* TransactionUseCase) Error(transactionId string) (*model.Transaction, error) {
	transaction, err := t.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found",transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionError
	err := t.transactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
