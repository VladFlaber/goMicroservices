package main

import (
	"database/sql"
	"fmt"
	users "github.com/VladFlaber/goMicroservices/authentication-service/Users"
	"log"
	"strconv"
	_"github.com/mattn/go-sqlite3"
)
func main()  {

	db,err:=sql.Open("sqlite3","Users.sqlite")


	if err != nil {
		log.Println("failed to connect to DATABASE",err)
	}
	_,err=db.Exec(createTable)
	if err != nil {
		fmt.Println(err)
	}
	repo:=users.UserRepository{DB:db }
	u:=users.NewUser("1","qwewqe@gmail.com","asdasd")
	id,err:=repo.Create(u)
	if err != nil {
		fmt.Println(err)
	}

	user,_:=repo.GetById(strconv.Itoa(int(id)))
	fmt.Println(user)

}

const createTable ="CREATE TABLE IF NOT EXISTS (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,email TEXT UNIQUE,psw TEXT,isBlocked INTEGER NOT NULL DEFAULT(0) CHECK (isBlocked in (0,1),createdAt time,updateAT time)"
