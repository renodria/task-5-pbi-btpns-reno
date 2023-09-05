package app

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Update struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Password struct {
	Password string `json:"password"`
}
