package extensions

type UserIsAdminResponse struct {
	IsAdmin bool   `json:"isAdmin"`
	Error   string `json:"error"`
}

type UserExistsResponse struct {
	Exists bool   `json:"exists"`
	Error  string `json:"error"`
}
