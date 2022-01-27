package users

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
