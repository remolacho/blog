package post

import (
	"blog/pkg/engineDB"
)

type PostService struct {
	Data *engineDB.Data
}

func (service *PostService) All() ([]Post, error) {
	var posts []Post
	records := service.Data.DB.Find(&posts)
	return posts, records.Error
}

func (service *PostService) Find(id uint) (Post, error) {
	var post Post
	record := service.Data.DB.First(&post, id)
	return post, record.Error
}

func (service *PostService) FindByUser(userID uint) ([]Post, error) {
	var posts []Post
	record := service.Data.DB.Where("user_id = ?", userID).Find(&posts)
	return posts, record.Error
}

func (service *PostService) Create(post *Post) error {
	result := service.Data.DB.Create(&post) // puntero de la data a crear
	return result.Error
}

func (service *PostService) Update(id uint, p Post) error {
	var post Post
	record := service.Data.DB.First(&post, id)

	if record.Error != nil {
		return record.Error
	}

	record = service.Data.DB.Model(&post).Updates(p)
	return record.Error
}

func (service *PostService) Delete(id uint) error {
	var post Post
	record := service.Data.DB.Delete(&post, id)
	return record.Error
}
