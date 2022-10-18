package domain

type UserCore struct {
	ID       uint
	Name     string
	Email    string
	Phone    string
	Address  string
	Password string
}

type Repository interface {
	GetMyUser() (UserCore, error)
	Update(updatedUser UserCore) (UserCore, error)
	Delete(deletedUser UserCore) (UserCore, error)
	GetByEmail(email string) (UserCore, error)
	AddUser(newUser UserCore) (UserCore, error)
	GetUser(user UserCore) (UserCore, error)
}

type Service interface {
	MyProfile() (UserCore, error)
	UpdateProfile(updatedUser UserCore) (UserCore, error)
	Deactivate(deletedUser UserCore) (UserCore, error)
	ShowByEmail(email string) (UserCore, error)
	Register(newUser UserCore) (UserCore, error)
	Login(user UserCore) (UserCore, error)
}
