package fetching

import (
	cli "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal"
	"github.com/pkg/errors"
)

type Service interface {
	GetAllCryptos() ([]cli.Crypto, error)
	GetCryptoById(symbol string) (cli.Crypto, error)
}

type service struct {
	bR cli.CryptoRepoInterface
}

func NewService(r cli.CryptoRepoInterface) Service {
	return &service{r}
}

func (s *service) GetAllCryptos() ([]cli.Crypto, error) {
	cryptoMap, err := s.bR.GetAll()
	var cryptos []cli.Crypto

	for _, value := range cryptoMap {
		cryptos = append(cryptos, value)
	}

	return cryptos, err
}

func (s *service) GetCryptoById(symbol string) (cli.Crypto, error) {
	cryptoMap, _ := s.bR.GetAll()

	if value, ok := cryptoMap[symbol]; ok {
		return value, nil
	}

	return cli.Crypto{}, errors.Errorf("Crypto %s not found", symbol)
}