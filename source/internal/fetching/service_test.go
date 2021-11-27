package fetching

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"

	"github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal"
	mockRepo "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal/storage/mocks"
)

const VALID_SYMBOL = "BTMT"
const INVALID_SYMBOL = "ABCD"

func getCryptoMapFixtures() (map[string]cryptocli.Crypto) {
	return  map[string]cryptocli.Crypto{
		VALID_SYMBOL: {
			Symbol: VALID_SYMBOL,
			LastPrice: 1.01,
			HighPrice: 1.001,
			LowPrice:  1.12,
			Volume:    0.1212,
			LastId:    100,
			FirstId:   101,
		},
	}
} 


func TestGetCryptoByIdShouldReturnACryptoEntity(t *testing.T) {
	repo := new(mockRepo.MockRespository)
	testService := NewService(repo)

	repo.On("GetAll").Return(getCryptoMapFixtures(), nil)

	result, _ := testService.GetCryptoById(VALID_SYMBOL)

	repo.AssertExpectations(t)

	assert.Equal(t, VALID_SYMBOL, result.Symbol)
}

func TestGetCryptoByIdReturnAnError(t *testing.T) {
	repo := new(mockRepo.MockRespository)
	testService := NewService(repo)

	expectedErrorMsg := fmt.Sprintf("Crypto %s not found", INVALID_SYMBOL)

	repo.On("GetAll").Return(getCryptoMapFixtures(), nil)

	result, resultErr := testService.GetCryptoById(INVALID_SYMBOL)

	repo.AssertExpectations(t)

	assert.Equal(t, cryptocli.Crypto{}, result)
	assert.EqualErrorf(t, resultErr, expectedErrorMsg, "Error should be: %v, got: %v", expectedErrorMsg, resultErr)
}


