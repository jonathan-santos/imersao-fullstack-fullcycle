package usecase

import (
	"errors"
	"github.com/jonathan-santos/imersao-fullstack-fullcycle/codepix/domain/model"
)

type PixKeyUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixKeyUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	err := p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey,.ID == "" {
		return nil, errors.New("unable to create new pixkey at the moment")
	}

	return pixKey, nil
}


func (p *PixKeyUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil	
}