package extensions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthDecision struct {
	Allowed bool `json:"allowed"`
}

type UserInfo struct {
	Exists  bool `json:"exists"`
	IsAdmin bool `json:"isAdmin"`
}

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

func GetAuthDecision(answer string) (bool, error) {
	bytes := []byte(answer)
	var result AuthDecision

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return false, err
	}

	return result.Allowed, nil
}

func AuthRequestAndDecision(url string) (bool, error) {
	authResult, err := Request(url)
	if err != nil {
		log.Println("Could not send auth request:", err)
		return false, err
	}

	allowed, err := GetAuthDecision(authResult)
	if err != nil {
		log.Println("Could not interpret auth response:", err)
		return false, err
	}

	return allowed, nil
}

func GetUserInfo(url string) (bool, bool, error) {
	response, err := Request(url)
	if err != nil {
		log.Println("Could not send userinfo request:", err)
		return false, false, err
	}

	bytes := []byte(response)
	var result UserInfo

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Println("Could not unmarshal userinfo result:", err)
		return false, false, err
	}

	return result.Exists, result.IsAdmin, nil
}
