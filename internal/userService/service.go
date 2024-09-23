package userService

type UserService struct {
	repoUsers UserRepository
}

func NewService(repoUsers UserRepository) *UserService {
	return &UserService{repoUsers: repoUsers}
}

func (u *UserService) CreateUser(user User) (User, error) {
	return u.repoUsers.CreateUser(user)
}

func (u *UserService) GetAllUsers() ([]User, error) {
	return u.repoUsers.GetAllUsers()
}

func (u *UserService) DeleteUserByID(id int) (User, error) {
	return u.repoUsers.DeleteUserByID(id)
}

func (u *UserService) UpdateUserByID(id int, user User) (User, error) {
	return u.repoUsers.UpdateUserByID(id, user)
}
