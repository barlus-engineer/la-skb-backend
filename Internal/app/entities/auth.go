package entities

type AuthReturnData struct {
	Status	int
	Message	string
}

type AuthFormData struct {
	Username string
	Password string
}