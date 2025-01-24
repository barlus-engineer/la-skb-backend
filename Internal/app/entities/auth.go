package entities

type Auth struct {
	Username string
	Password string
}

type AuthErr struct {
	UserNotFound string
}

type AuthForm struct {
	Username string
	Password string
}
