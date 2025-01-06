package entities

type Auth struct {
	Username string
	Password string
}

type AuthReturnData struct {
	Status	int
	Message	string
}

type AuthFormData struct {
	Username string
	Password string
}