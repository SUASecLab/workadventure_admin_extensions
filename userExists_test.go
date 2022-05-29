package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SUASecLab/workadventure_admin_extensions/db"
	"github.com/gorilla/mux"
)

func Test_userExistsHandler(t *testing.T) {
	type dataDescription struct {
		uuid           string
		expectedResult string
	}
	tests := []struct {
		name        string
		description dataDescription
	}{
		{"userExists positive", dataDescription{db.MockUUID, "{\"exists\":true,\"error\":\"\"}"}},
		{"userExists negative", dataDescription{"2a579658-ff5a-479f-bea4-2c50543955ea", "{\"exists\":false,\"error\":\"\"}"}},
		{"userExists empty uuid", dataDescription{"", "404 page not found"}},
		{"userExists no uuid", dataDescription{"abc", "{\"exists\":false,\"error\":\"The provided UUID is invalid\"}"}},
	}
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			path := fmt.Sprintf("/userExists/%s", data.description.uuid)
			expectedResult := data.description.expectedResult

			//result, err := extensions.Request(url)
			request, err := http.NewRequest(http.MethodGet, path, nil)
			if err != nil {
				t.Errorf("Could not send request: %v", err)
			}

			recorder := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/userExists/{uuid}", userExistsHandler)
			router.ServeHTTP(recorder, request)

			result := recorder.Body.String()
			if !strings.Contains(result, expectedResult) {
				t.Errorf("Error when testing userExists: expected %v, received %v", expectedResult, result)
			}
		})
	}
}
