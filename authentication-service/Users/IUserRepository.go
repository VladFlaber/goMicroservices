package Users

type IuserRepository interface {
	GetById(id int) *User
	GetAll () *[]User
	GetByEmail(email string) *User
	Create (user *User)
	Edit  (user *User)
	Remove (user *User) int
	Block (user *User) int
}