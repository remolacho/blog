package user

// Repository handle the CRUD operations with Users.
type Repository interface {
	All() ([]User, error)
	Find(id uint) (User, error)
	FindByUsername(username string) (User, error)
	Create(user *User) error
	Update(id uint, user User) error
	Delete(id uint) error
}
