package authentication_service

import (
	"database/sql"
	"fmt"
	users "github.com/VladFlaber/goMicroservices/authentication-service/Users"
	"log"
)
func main()  {

	db,err:=sql.Open("sqlite3","./users.db")


	if err != nil {
		log.Println("failed to connect to DATABASE")
	}
	db.Exec(createTable)
	repo:=users.UserRepository{DB:db }
	u:=users.NewUser("1","qwewqe@gmail.com","asdasd",1)
	user,err:=repo.Create(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

}
const createTable ="CREATE TABLE IF NOT EXISTS (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,email TEXT UNIQUE,psw TEXT,isBlocked INTEGER NOT NULL DEFAULT(0) CHECK (isBlocked in (0,1),createdAt time,updateAT time)"
