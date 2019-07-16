package giphy

import (
	"errors"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

const(
	API_KEY = "dGiRPc2btcRacSHPVKR4V7NzR47Shliw"
	URL = "https://api.giphy.com/v1/gifs/random"
)

func GetRandomGiphy() (gifUrl string, err error) {

	resp, err := resty.R().SetQueryParams(
		map[string]string{
			"api_key":API_KEY,
			"rating":"G",
		}).
		Get(URL)

	if err != nil {
		return "", errors.New("Error on Resty:" + err.Error())
	}


	gifUrl = gjson.Get(string(resp.Body()),"data.image_url").String()

	if gifUrl == "" {
		return "", errors.New("no giphy found")
	}

	return gifUrl, nil

}