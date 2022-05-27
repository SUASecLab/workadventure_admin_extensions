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

func UnmarshalResponse(content string, format struct{}) (struct{}, error) {
	bytes := []byte(content)

	err := json.Unmarshal(bytes, &format)
	if err != nil {
		return struct{}{}, err
	}

	return format, nil
}

func UnmarshalledJSONRequestResponse(url string, format struct{}) (struct{}, error) {
	response, err := Request(url)
	if err != nil {
		return struct{}{}, err
	}

	result, err := UnmarshalResponse(response, format)
	if err != nil {
		return struct{}{}, err
	}

	return result, nil
}
