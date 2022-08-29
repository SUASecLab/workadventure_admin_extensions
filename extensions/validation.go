package extensions

import (
	"encoding/json"
	"log"
)

type Validation struct {
	Error string `json:"error"`
	Valid bool   `json:"valid"`
}

func GetValidationResult(url string) (Validation, error) {
	response, err := Request(url)
	if err != nil {
		errorMsg := "Could not send validation request"
		log.Println(errorMsg, err)
		return Validation{
			Error: errorMsg,
			Valid: false,
		}, err
	}

	bytes := []byte(response)
	var result Validation

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		errorMsg := "Could not interpret validation result"
		log.Println(errorMsg, err)
		return Validation{
			Error: errorMsg,
			Valid: false,
		}, err
	}

	return result, nil
}
