package model

import (
	"github.com/jinzhu/gorm"
	"github.com/nEdAy/wx_attendance_api_server/internal/db"
)

// UserModel 用户表
type UserModel struct {
	Id         int    `gorm:"column:id;primary_key" json:"id"`
	Username   string `gorm:"column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	FaceToken  string `gorm:"column:face_token" json:"face_token"`
	FaceUrl    string `gorm:"column:face_url" json:"face_url"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

// TableName 返回asc_door 表名称
func (UserModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "user")
}

// AddUser insert a new UserModel into database and returns
// last inserted Id on success.
func AddUser(m *UserModel) (err error) {
	err = db.DB.Create(m).Error
	return err
}

// GetUserById retrieves UserModel by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *UserModel, err error) {
	v = new(UserModel)
	if err = db.DB.First(&v, 10).Error;
		err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserByUsername retrieves UserModel by username. Returns error if
// Id doesn't exist
func GetUserByUsername(username string) (v *UserModel, err error) {
	v = new(UserModel)
	if err = db.DB.Where("username = ?", username).Find(&v).Error; err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all UserModel matches certain condition. Returns empty list if
// no records exist
func GetAllUser() (v []*UserModel, err error) {
	if err = db.DB.Order("id desc").Select("id,username,face_url").Find(&v).Error; err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateUser updates UserModel by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *UserModel) (err error) {
	// ascertain id exists in the database
	if err = db.DB.First(&m, m.Id).Error; err == nil {
		err = db.DB.Save(m).Error
	}
	return err
}

// DeleteUser deletes UserModel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	v := UserModel{Id: id}
	// ascertain id exists in the database
	if err = db.DB.First(&v, id).Error; err == nil {
		err = db.DB.Delete(v).Error
	}
	return err
}

func IsUserExist(username string) (exist bool, err error) {
	var count int
	if err = db.DB.Model(&UserModel{}).Where("username = ?", username).Count(&count).Error; err == nil {
		return count > 0, nil
	}
	return false, err
}
