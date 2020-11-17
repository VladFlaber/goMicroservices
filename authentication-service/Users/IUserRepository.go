package Users

type IuserRepository interface {
	GetById(id int) (*User,error)
	GetAll () (*[]User,error)
	GetByEmail(email string) (*User,error)
	Create (user *User)error
	Edit  (user *User)error
	Remove (user *User) (int,error)
	Block (user *User) (int,error)
}