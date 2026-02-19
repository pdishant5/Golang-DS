package user

type User struct {
	Username string
	Password string
}

func New(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
