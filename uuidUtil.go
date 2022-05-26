package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

func isUUIDValid(requestedUUID string) (bool, []byte) {
	_, err := uuid.Parse(requestedUUID)

	response := UserExistsResponse{}

	if err != nil {
		response.Exists = false
		response.Error = "The provided UUID is invalid"
		responseToString, err := json.Marshal(response)

		if err != nil {
			log.Println(err)
			return false, nil
		}
		return false, responseToString
	}

	return true, nil
}
