package binance


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	cli "github.com/DanielQuerolBeltran/CliCryptoInfo/source/internal"
)

const (
	endpoint = "/ticker/24hr"
	url      = "https://api2.binance.com/api/v3"
)

type jsonElement struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
	HighPrice string `json:"highPrice"`
	LowPrice  string `json:"lowPrice"`
	Volume    string `json:"volume"`
	LastId    int    `json:"lastId"`
	FirstId   int    `json:"firstId"`
}

type repository struct {
	url string
}

func NewRepository() cli.CryptoRespoInteraface {
	return &repository{url: url}
}

func (c *repository) GetAll() (map[string]cli.Crypto, error) {
	var jsonElements []jsonElement
	contents, err := getData()
	cryptoMap := make(map[string]cli.Crypto)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &jsonElements)
	if err != nil {
		return nil, err
	}

	for _, element := range jsonElements {
		var lastPrice, _ = strconv.ParseFloat(element.LastPrice, 32)
		var higtPrice, _ = strconv.ParseFloat(element.HighPrice, 32)
		var lowPrice, _ = strconv.ParseFloat(element.LowPrice, 32)
		var volume, _ = strconv.ParseFloat(element.Volume, 32)

		cryptoMap[element.Symbol] = cli.Crypto{
			Symbol:    element.Symbol,
			LastPrice: lastPrice,
			HighPrice: higtPrice,
			LowPrice:  lowPrice,
			Volume:    volume,
			LastId:    element.LastId,
			FirstId:   element.FirstId,
		}

	}

	return cryptoMap, err
}

func getData() ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("%v%v", url, endpoint))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}