package models

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ValidationErr struct {
	FieldValue string
	ErrMassage string
}

type LoginUser struct {
	LoginMail     string `json:"email"`
	LoginPassword string `json:"password"`
}

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Key      string `yaml:"key"`
}

type TokenResponse struct {
	ResponseMessage string
	Token           string
}
