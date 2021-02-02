package post

// Repository handle the CRUD operations with Posts.
type Repository interface {
	All() ([]Post, error)
	Find(id uint) (Post, error)
	FindByUser(userID uint) ([]Post, error)
	Create(post *Post) error
	Update(id uint, post Post) error
	Delete(id uint) error
}
