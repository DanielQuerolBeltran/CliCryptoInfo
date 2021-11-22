package mocks_test

import (
	cryptocli "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal"
	"github.com/stretchr/testify/mock"
)

type MockRespository struct {
	mock.Mock
}

func (m *MockRespository) GetAll() (map[string]cryptocli.Crypto, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(map[string]cryptocli.Crypto), args.Error(1)
}
