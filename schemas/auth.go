package schemas

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}