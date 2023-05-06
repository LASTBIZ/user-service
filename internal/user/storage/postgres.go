package storage

import (
	"gorm.io/gorm"
	"lastbiz/user-service/internal/user"
	"strings"
	"time"
)

type storage struct {
	db gorm.DB
}

func NewUserStorage(db gorm.DB) user.Storage {
	return storage{
		db: db,
	}
}

func (s storage) CreateUser(userCreate user.User) (user.User, error) {
	var findUser user.User
	err := s.db.Create(&userCreate).Where(&user.User{Email: userCreate.Email}).Find(&findUser).Error
	if err != nil {
		return findUser, err
	}
	return findUser, nil
}

func (s storage) DeleteUser(userId uint32) error {
	err := s.db.Delete(&user.User{ID: userId}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s storage) UpdateUser(userUpdate user.User) (user.User, error) {
	err := s.db.Model(&user.User{}).Where(&user.User{ID: userUpdate.ID}).
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
	return userUpdate, err

}

func (s storage) GetUser(userId uint32) (user.User, error) {
	var getUser user.User
	err := s.db.
		Model(&user.User{}).
		Where(&user.User{ID: userId}).
		First(&getUser).
		Error
	return getUser, err
}

func (s storage) GetUserByEmail(email string) (user.User, error) {
	var getUser user.User
	err := s.db.
		Model(&user.User{}).
		Where(&user.User{Email: strings.ToLower(email)}).
		First(&getUser).Error
	return getUser, err
}
