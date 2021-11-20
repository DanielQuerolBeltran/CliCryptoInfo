package cryptocli

import "fmt"

type Crypto struct {
	Symbol string
	LastPrice float64
	HighPrice float64
	LowPrice float64
	Volume float64
	LastId int
	FirstId int
}

type CryptoRespoInteraface interface {
	GetAll() (map[string]Crypto, error)
}

func (c Crypto) Print() {
	fmt.Printf("Symbol: %s \nPrice: %f \nLast Price: %f \nLow Price %f \n \n",
	c.Symbol, c.LastPrice, c.HighPrice, c.LowPrice);
}

