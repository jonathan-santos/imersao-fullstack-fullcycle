package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/jonathan-santos/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Jonathan Santos"
	email := "jonathan.santosaugusto@gmail.com"

	user, err := model.NewUser(name, email)
	
	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)
}
