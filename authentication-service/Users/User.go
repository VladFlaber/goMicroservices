package Users

import (
	"log"
	"regexp"
	"time"
)

type User struct {
	Id string
	Email string
	Password string
	IsBlocked bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, email string, password string) *User {
	return &User{Id: id, Email: email, Password: password,
		CreatedAt: time.Now(),UpdatedAt: time.Now()}
}
func (User) validatePassword(password string) bool  {
	if len(password)==0{
		log.Println("password is empty")
		return false
	}
	return true
}
func (User)validateEmail(email string ) bool {

	regex:="^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	return  regexp.MustCompile(regex).MatchString(email)
}
func (User) ValidateUser (user User)bool{
	if !user.validateEmail(user.Email){
		log.Println("invalid email : ", user.Email)
		return false
	}

	if !user.validatePassword(user.Password){
		log.Println("invalid password : ",user.Password)
		return false
	}
	return true
}
