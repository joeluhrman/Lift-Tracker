package model

import "database/sql"

type User struct {
	Model

	Username sql.NullString `gorm:"unique;not null" json:"username"`
}

type UserPassword struct {
	Model

	Password sql.NullString `gorm:"not null" json:"password"`
}

// Creates a User and corresponding UserPassword in the DB.
func CreateUser(user *User, userPassword *UserPassword) error {
	err := DBConn.Create(user).Error
	if err != nil {
		return err
	}

	userPassword.ID = user.ID

	err = DBConn.Create(userPassword).Error
	if err != nil {
		DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)
	}

	return err
}
