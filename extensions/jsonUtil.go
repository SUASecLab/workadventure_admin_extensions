package extensions

import "encoding/json"

type UserIsAdminResponse struct {
	IsAdmin bool   `json:"isAdmin"`
	Error   string `json:"error"`
}

type UserExistsResponse struct {
	Exists bool   `json:"exists"`
	Error  string `json:"error"`
}

func JSONToUserIsAdminResponse(input string) (UserIsAdminResponse, error) {
	bytes := []byte(input)
	var result UserIsAdminResponse

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func JSONToUserExistsResponse(input string) (UserExistsResponse, error) {
	bytes := []byte(input)
	var result UserExistsResponse

	err := json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
