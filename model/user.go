package model

import (
	"github.com/jinzhu/gorm"
)

// UserModel 用户表
type UserModel struct {
	Model

	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	FaceToken string `gorm:"column:face_token" json:"face_token"`
	FaceUrl   string `gorm:"column:face_url" json:"face_url"`
}

// TableName 返回asc_door 表名称
func (UserModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "user")
}

// AddUser insert a new UserModel into database and returns
// last inserted Id on success.
func AddUser(m *UserModel) (err error) {
	err = DB.Create(m).Error
	return err
}

// GetUserById retrieves UserModel by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *UserModel, err error) {
	v = new(UserModel)
	if err = DB.First(&v, 10).Error;
		err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserByUsername retrieves UserModel by username. Returns error if
// Id doesn't exist
func GetUserByUsername(username string) (v *UserModel, err error) {
	v = new(UserModel)
	if err = DB.Where("username = ?", username).Find(&v).Error; err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all UserModel matches certain condition. Returns empty list if
// no records exist
func GetAllUser() (v []*UserModel, err error) {
	if err = DB.Order("id desc").Select("id,username,face_url").Find(&v).Error; err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUser updates UserModel by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *UserModel) (err error) {
	// ascertain id exists in the database
	if err = DB.First(&m, m.Id).Error; err == nil {
		err = DB.Save(m).Error
	}
	return err
}

// DeleteUser deletes UserModel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	v := UserModel{}
	// ascertain id exists in the database
	if err = DB.First(&v, id).Error; err == nil {
		err = DB.Where("id = ?", id).Delete(v).Error
	}
	return err
}

func IsUserExist(username string) (exist bool, err error) {
	var count int
	if err = DB.Model(&UserModel{}).Where("username = ?", username).Count(&count).Error; err == nil {
		return count > 0, nil
	}
	return false, err
}
