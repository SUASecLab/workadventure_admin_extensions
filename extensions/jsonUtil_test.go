package extensions

import "testing"

var testDataUserIsAdmin = []struct {
	name     string
	json     string
	response UserIsAdminResponse
}{
	{"UserIsAdmin 1", "", UserIsAdminResponse{}},
	{"UserIsAdmin 2", "sdfsdf", UserIsAdminResponse{}},
	{"UserIsAdmin 3", "{\"isAdmin\":false, \"error\":\"\"}", UserIsAdminResponse{
		IsAdmin: false,
		Error:   "",
	}},
	{"UserIsAdmin 4", "{\"isAdmin\":false, \"error\":\"Invalid UUID\"}", UserIsAdminResponse{
		IsAdmin: false,
		Error:   "Invalid UUID",
	}},
	{"UserIsAdmin 5", "{\"isAdmin\":true, \"error\":\"\"}", UserIsAdminResponse{
		IsAdmin: true,
		Error:   "",
	}},
}

func TestJSONToUserIsAdminResponse(t *testing.T) {
	for _, data := range testDataUserIsAdmin {
		t.Run(data.name, func(t *testing.T) {
			data := data
			t.Parallel()
			result, _ := JSONToUserIsAdminResponse(data.json)
			if result != data.response {
				t.Errorf("Error while converting JSON: expected %v, received %v", data.response, result)
			}
		})
	}
}

var testDataUserExists = []struct {
	name     string
	json     string
	response UserExistsResponse
}{
	{"UserExists 1", "", UserExistsResponse{}},
	{"UserExists 2", "sdfsdf", UserExistsResponse{}},
	{"UserExists 3", "{\"exists\":false, \"error\":\"\"}", UserExistsResponse{
		Exists: false,
		Error:  "",
	}},
	{"UserExists 4", "{\"exists\":false, \"error\":\"Invalid UUID\"}", UserExistsResponse{
		Exists: false,
		Error:  "Invalid UUID",
	}},
	{"UserExists 5", "{\"exists\":true, \"error\":\"\"}", UserExistsResponse{
		Exists: true,
		Error:  "",
	}},
}

func TestJSONToUserExistsResponse(t *testing.T) {
	for _, data := range testDataUserExists {
		t.Run(data.name, func(t *testing.T) {
			data := data
			t.Parallel()
			result, _ := JSONToUserExistsResponse(data.json)
			if result != data.response {
				t.Errorf("Error while converting JSON: expected %v, received %v", data.response, result)
			}
		})
	}
}
