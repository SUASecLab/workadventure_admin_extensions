package extensions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func UnmarshalResponse(content string, format interface{}) (interface{}, error) {
	bytes := []byte(content)

	err := json.Unmarshal(bytes, &format)
	if err != nil {
		return nil, err
	}

	return format, nil
}

func UnmarshalledJSONRequestResponse(url string, format interface{}) (interface{}, error) {
	response, err := Request(url)
	if err != nil {
		return nil, err
	}

	result, err := UnmarshalResponse(response, &format)
	if err != nil {
		return nil, err
	}

	return result, nil
}
