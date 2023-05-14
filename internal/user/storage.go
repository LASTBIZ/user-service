package user

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Storage struct {
	db gorm.DB
}

func NewUserStorage(db gorm.DB) Storage {
	return Storage{
		db: db,
	}
}

func (s Storage) CreateUser(userCreate User) (*User, error) {
	var findUser User
	err := s.db.Create(&userCreate).Where(&User{Email: userCreate.Email}).Find(&findUser).Error
	if err != nil {
		return nil, err
	}
	return &findUser, nil
}

func (s Storage) DeleteUser(userId uint32) error {
	err := s.db.Delete(&User{ID: userId}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) UpdateUser(userUpdate User) (*User, error) {
	err := s.db.Model(&User{}).Where(&User{ID: userUpdate.ID}).
		Updates(map[string]interface{}{
			"blocked":      userUpdate.Blocked,
			"role":         userUpdate.Role,
			"firstname":    userUpdate.Firstname,
			"lastname":     userUpdate.Lastname,
			"is_verify":    userUpdate.IsVerify,
			"phone":        userUpdate.Phone,
			"messengers":   userUpdate.Messengers,
			"organization": userUpdate.Organization,
			"updated_at":   time.Now(),
		}).
		First(&userUpdate).Error
	return &userUpdate, err

}

func (s Storage) GetUser(userId uint32) (*User, error) {
	var getUser User
	err := s.db.
		Model(&User{}).
		Where(&User{ID: userId}).
		First(&getUser).
		Error
	return &getUser, err
}

func (s Storage) GetUserByEmail(email string) (*User, error) {
	var getUser User
	err := s.db.
		Model(&User{}).
		Where(&User{Email: strings.ToLower(email)}).
		First(&getUser).Error
	return &getUser, err
}
