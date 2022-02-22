package service

import (
	"github.com/golang-jwt/jwt"
	"time"
	"web-server-in-docker/pkg/models"
	"web-server-in-docker/pkg/store"
)

func GiveToken(u models.LoginUser) string {
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer pgxconn.Close(context.Background())
	//psqlconn := fmt.Sprintf(GetConfig())
	//db, err := sql.Open("postgres", psqlconn)
	//if err != nil {
	//panic(err)
	//}
	//defer db.Close()
	token := CreateToken(u.LoginMail)
	err := store.InsertToken(u.LoginMail, token)
	if err != nil {
		return err.Error()
	}
	//defer insertToken.Close()
	return token
}

func CreateToken(email string) string {
	//pgxconn, err := pgx.Connect(context.Background(), GetConfig())
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//db, err := sqlx.Connect("postgres", store.GetConfig())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//var Fname string
	//var Lname string
	//err = db.Get(&Fname, "SELECT firstname FROM signed_users WHERE email=$1", email)
	//if err == nil {
	//	fmt.Sprintln(err)
	//}
	//err = db.Get(&Lname, "SELECT lastname FROM signed_users WHERE email=$1", email)
	//if err == nil {
	//	fmt.Sprintln(err)
	//}
	Fname, Lname, err := store.GetNames(email)
	if err != nil {
		return err.Error()
	}
	hmacSampleSecret := []byte(store.GetKey())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      Fname + " " + Lname,
		"ExpiresAt": time.Now().Add(12 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}
	return tokenString
}
