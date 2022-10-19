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
	GetMyUser(userID uint) (UserCore, error)
	Update(updatedUser UserCore, userID uint) (UserCore, error)
	Delete(userID uint) (UserCore, error)
	GetByEmail(email string) (UserCore, error)
	AddUser(newUser UserCore) (UserCore, error)
	GetUser(existUser UserCore) (UserCore, error)
}

type Service interface {
	MyProfile(userID uint) (UserCore, error)
	UpdateProfile(updatedUser UserCore, userID uint) (UserCore, error)
	Deactivate(userID uint) (UserCore, error)
	ShowByEmail(email string) (UserCore, error)
	Register(newUser UserCore) (UserCore, error)
	Login(existUser UserCore) (UserCore, error)
}
