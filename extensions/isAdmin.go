package extensions

func UserIsAdmin(adminExtensionsURL, token string) (bool, string) {
	// Check whether user is admin
	response, err := Request(adminExtensionsURL + "/isAdmin/" + token)
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
