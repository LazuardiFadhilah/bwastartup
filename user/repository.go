package user

import "gorm.io/gorm"

type Repository interface {
	// Function Save User
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(ID int) (User, error)
	Update(user User) (User, error)
	Delete(ID int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindById(id int) (User, error) {

	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error

	if err != nil {

		return user, err

	}

	return user, nil

}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
