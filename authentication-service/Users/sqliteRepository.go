package Users

import (
	"database/sql"
	"time"
)
type UserRepository struct{
	DB *sql.DB
}


func (u UserRepository) GetById(id string) (*User,error) {

	res ,err:= u.DB.Query("select * from users U where U.id=" +id)
	user:=User{}
	if err != nil {
		return nil, err
	}
	res.Scan(&user.Id,&user.Email,&user.Password,
		&user.IsBlocked,&user.CreatedAt,&user.UpdatedAt)
	return &user,nil
}

func (u UserRepository) GetAll() (*[]User,error) {
	rows,err:= u.DB.Query("SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}
		users:=[]User{}


	for rows.Next() {
		user:=User{}

		rows.Scan(&user.Id,&user.Email,&user.Password,
			&user.IsBlocked,&user.CreatedAt,&user.UpdatedAt)
		users=append(users,user)
	}
	return &users,nil
}

func (u UserRepository) GetByEmail(email string) (*User,error) {
	res ,err:= u.DB.Query("select * from users U where U.id=$1", email)
	user:=User{}
	if err != nil {
		return nil, err
	}
	res.Scan(&user.Id,&user.Email,&user.Password,
		&user.IsBlocked,&user.CreatedAt,&user.UpdatedAt)
	return &user,nil
}

func (u UserRepository) Create(user *User) (int64,error) {
	res,err:=u.DB.Exec("INSERT INTO USERS (email,password,isblocked,createdAt,updatedAt)" +
		"VALUES $1,$2,$3,$4,$5",user.Email,user.Password,user.IsBlocked,user.CreatedAt,user.UpdatedAt)
	if err != nil {
		return -1 ,err
	}
	id,_:=res.LastInsertId()
	return id,nil
}

func (u UserRepository) Edit(user *User) error {
	user,err:=u.GetById(user.Id)
	if err != nil {
		return err
	}
	_, err = u.DB.Exec("update Users set email=$1 , password=$2,updatedAt=$3",
		user.Email, user.Password, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Remove(user *User) (int,error) {
	panic("implement me")

}

func (u UserRepository) Block(user *User) (int,error) {
	panic("implement me")
}

