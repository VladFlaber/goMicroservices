package authentication_service

import (
	"database/sql"
	users "github.com/VladFlaber/goMicroservices/authentication-service/Users"
	)
func main()  {
	sql.Open("","")
	repo:=users.UserRepository{}

}