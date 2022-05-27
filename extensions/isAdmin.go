package extensions

func UserIsAdmin(adminExtensionsURL, uuid string) (bool, string) {
	// Check if the uuid is valid
	isValid := IsUUIDValid(uuid)
	if !isValid {
		msg := "Invalid user ID"
		return false, msg
	}

	// Check whether user is admin
	response, err := Request(adminExtensionsURL + "/isAdmin/" + uuid)
	if err != nil {
		msg := "Could not connect to admin services:"
		return false, msg
	}

	result, err := JSONToUserIsAdminResponse(response)
	if err != nil {
		msg := "Could not interpret the server's response:"
		return false, msg
	}

	if !result.IsAdmin {
		return false, ""
	}
	return true, ""
}
