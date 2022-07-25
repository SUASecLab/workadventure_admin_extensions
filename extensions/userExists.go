package extensions

func UserExists(adminExtensionsURL, token string) (bool, string) {
	// Check whether user exists
	response, err := Request(adminExtensionsURL + "/userExists/" + token)
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
