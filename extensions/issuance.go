package extensions

import (
	"encoding/json"
	"log"
)

type Issuance struct {
	Error string `json:"error"`
	Token string `json:"token"`
}

func IssueToken(url string) (Issuance, error) {
	issuanceResult, err := Request(url)
	if err != nil {
		errorMsg := "Could not send auth request"
		log.Println(errorMsg, err)
		return Issuance{
			Error: errorMsg,
			Token: "",
		}, err
	}

	bytes := []byte(issuanceResult)
	var result Issuance

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		errorMsg := "Could not interpret auth response"
		return Issuance{
			Error: errorMsg,
			Token: "",
		}, err
	}

	return result, nil
}
