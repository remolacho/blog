package user

import (
	"blog/pkg/engineDB"
)

type UserService struct {
	Data *engineDB.Data
}

func (service *UserService) All() ([]User, error) {
	var users []User
	records := service.Data.DB.Find(&users)
	return users, records.Error
}

func (service *UserService) Find(id uint) (User, error) {
	var user User
	record := service.Data.DB.First(&user, id)
	return user, record.Error
}

func (service *UserService) FindByUsername(username string) (User, error) {
	var user User
	record := service.Data.DB.Where("username = ?", username).First(&user)
	return user, record.Error
}

func (service *UserService) Create(user *User) error {

	if user.Picture == "" {
		user.Picture = "https://placekitten.com/g/300/300"
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	result := service.Data.DB.Create(&user) // puntero de la data a crear
	return result.Error
}

func (service *UserService) Update(id uint, u User) error {
	var user User
	record := service.Data.DB.First(&user, id)

	if record.Error != nil {
		return record.Error
	}

	record = service.Data.DB.Model(&user).Updates(u)
	return record.Error
}

func (service *UserService) Delete(id uint) error {
	var user User
	record := service.Data.DB.Delete(&user, id)
	return record.Error
}
