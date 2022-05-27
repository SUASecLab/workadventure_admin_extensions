package extensions

func UserExists(adminExtensionsURL, uuid string) (bool, string) {
	// Check if the uuid is valid
	isValid := IsUUIDValid(uuid)
	if !isValid {
		msg := "Invalid user ID"
		return false, msg
	}

	// Check whether user exists
	response, err := Request(adminExtensionsURL + "/userExists/" + uuid)
	if err != nil {
		msg := "Could not connect to admin services:"
		return false, msg
	}

	result, err := JSONToUserExistsResponse(response)
	if err != nil {
		msg := "Could not interpret the server's response:"
		return false, msg
	}

	if !result.Exists {
		msg := "User does not exist"
		return false, msg
	}
	return true, ""
}
