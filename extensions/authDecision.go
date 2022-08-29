package extensions

import (
	"encoding/json"
	"log"
)

type AuthDecision struct {
	Allowed bool `json:"allowed"`
}

func GetAuthDecision(url string) (AuthDecision, error) {
	authResult, err := Request(url)
	if err != nil {
		log.Println("Could not send auth request:", err)
		return AuthDecision{Allowed: false}, err
	}

	bytes := []byte(authResult)
	var result AuthDecision

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return AuthDecision{Allowed: false}, err
	}

	return result, nil
}
