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

func Test_isAdminHandler(t *testing.T) {
	type dataDescription struct {
		uuid           string
		expectedResult string
	}
	tests := []struct {
		name        string
		description dataDescription
	}{
		{"isAdmin positive", dataDescription{db.MockUUID, "{\"isAdmin\":true,\"error\":\"\"}"}},
		{"isAdmin negative", dataDescription{"2a579658-ff5a-479f-bea4-2c50543955ea", "{\"isAdmin\":false,\"error\":\"User does not exist\"}"}},
		{"isAdmin empty uuid", dataDescription{"", "404 page not found"}},
		{"isAdmin no uuid", dataDescription{"abc", "{\"isAdmin\":false,\"error\":\"The provided UUID is invalid\"}"}},
	}
	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			path := fmt.Sprintf("/isAdmin/%s", data.description.uuid)
			expectedResult := data.description.expectedResult

			//result, err := extensions.Request(url)
			request, err := http.NewRequest(http.MethodGet, path, nil)
			if err != nil {
				t.Errorf("Could not send request: %v", err)
			}

			recorder := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/isAdmin/{uuid}", isAdminHandler)
			router.ServeHTTP(recorder, request)

			result := recorder.Body.String()
			if !strings.Contains(result, expectedResult) {
				t.Errorf("Error when testing isAdmin: expected %v, received %v", expectedResult, result)
			}
		})
	}
}
