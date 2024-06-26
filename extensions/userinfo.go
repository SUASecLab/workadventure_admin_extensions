package extensions

import (
	"encoding/json"
	"log"
)

type UserInfo struct {
	Exists  bool   `json:"exists"`
	IsAdmin bool   `json:"isAdmin"`
	UUID    string `json:"uuid"`
}

func GetUserInfo(url string) (UserInfo, error) {
	response, err := Request(url)
	if err != nil {
		log.Println("Could not send userinfo request:", err)
		return UserInfo{
			Exists:  false,
			IsAdmin: false,
		}, err
	}

	bytes := []byte(response)
	var result UserInfo

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		errorMsg := "Could not interpret userinfo result"
		log.Println(errorMsg, err)
		return UserInfo{
			Exists:  false,
			IsAdmin: false,
		}, err
	}

	return result, nil
}
