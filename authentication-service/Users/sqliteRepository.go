package Users

type UserRepository struct{

}

func (u UserRepository) GetById(id int) *User {
	panic("implement me")
}

func (u UserRepository) GetAll() *[]User {
	panic("implement me")
}

func (u UserRepository) GetByEmail(email string) *User {
	panic("implement me")
}

func (u UserRepository) Create(user *User) {
	panic("implement me")
}

func (u UserRepository) Edit(user *User) {
	panic("implement me")
}

func (u UserRepository) Remove(user *User) int {
	panic("implement me")
}

func (u UserRepository) Block(user *User) int {
	panic("implement me")
}


