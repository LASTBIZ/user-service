package user

import (
	"errors"
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
			"updated_at":   time.Now(),
			"organization": userUpdate.Organization,
		}).
		First(&userUpdate).Error
	return &userUpdate, err
}

func (s Storage) AddMessenger(messenger Messenger) (*Messenger, error) {
	var getMessenger Messenger
	err := s.db.Model(&Messenger{}).
		Create(&messenger).
		First(&getMessenger).
		Error

	return &getMessenger, err
}

func (s Storage) RemoveMessenger(id uint32, userID uint32) error {

	err := s.db.Delete(&Messenger{ID: id, UserId: userID}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) GetUser(userId uint32) (*User, error) {
	var getUser User
	result := s.db.Where(&User{ID: userId}).Preload("Messengers").Where(&User{ID: userId}).First(&getUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("USER_NOT_FOUND")
	}
	if result.Error != nil {
		return nil, errors.New("USER_NOT_FOUND")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}

	return &getUser, nil
}

func (s Storage) GetUserByEmail(email string) (*User, error) {
	var getUser *User
	if res := s.db.Where(&User{Email: strings.ToLower(email)}).Preload("Messengers").Where(&User{ID: userId}).First(&getUser); res.RowsAffected == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}
	return getUser, nil
}
