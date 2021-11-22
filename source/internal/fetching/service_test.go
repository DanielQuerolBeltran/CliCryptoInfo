package fetching

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal"
	mockRepo "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/storage/mocks"
)


func TestGetCryptoById(t *testing.T) {
	repo := new(mockRepo.MockRespository)
	testService := NewService(repo)

	symbol := "BTMT"

	expectedCrypt := cryptocli.Crypto{
		Symbol: symbol,
		LastPrice: 1.01,
		HighPrice: 1.001,
		LowPrice:  1.12,
		Volume:    0.1212,
		LastId:    100,
		FirstId:   101,
	}

	expectedCryptoMap := map[string]cryptocli.Crypto{
		symbol: expectedCrypt,
	}


	repo.On("GetAll").Return(expectedCryptoMap, nil)

	result, _ := testService.GetCryptoById(symbol)

	repo.AssertExpectations(t)

	assert.Equal(t, symbol, result.Symbol)
}


