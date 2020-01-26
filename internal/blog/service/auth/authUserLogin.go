package auth

type UserLoginParam struct {
	Username string
	Password string
}

func VerifyUser(user UserLoginParam) bool {
	if user.Username == "noxymon" && user.Password == "admin12345678" {
		return true
	}
	return false
}
