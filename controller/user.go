package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"github.com/nEdAy/wx_attendance_api_server/model"
	"github.com/nEdAy/wx_attendance_api_server/util"
	"github.com/nEdAy/wx_attendance_api_server/internal/face"
	"time"
	"strconv"
	"github.com/google/logger"
)

// Binding from Register JSON
type RegisterModel struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password"`
	PrefixCosUrl string `json:"prefixCosUrl" binding:"required"`
	FileName     string `json:"fileName" binding:"required"`
}

// Register 添加用户
func Register(c *gin.Context) {
	registerModel := new(RegisterModel)
	if err := c.ShouldBindWith(&registerModel, binding.JSON); err != nil {
		renderJSONWithError(c, err.Error())
		return
	}

	isUserExist, err := model.IsUserExist(registerModel.Username)
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	if isUserExist {
		renderJSONWithError(c, "用户<"+registerModel.Username+">已注册")
		return
	}

	faceToken := util.UserPwdEncrypt(registerModel.Username)
	faceCount, err := face.GetFaceCount(registerModel.PrefixCosUrl, registerModel.FileName, faceToken)

	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	if faceCount == -1 {
		renderJSONWithError(c, "已存在该用户名的人脸信息")
		return
	}
	if faceCount == 0 {
		renderJSONWithError(c, "未检测到人脸信息")
		return
	}
	if faceCount > 1 {
		renderJSONWithError(c, "请保证人脸照片中只包含一个人脸")
		return
	}

	userModel := new(model.UserModel)
	userModel.Username = registerModel.Username
	userModel.Password = util.UserPwdEncrypt(registerModel.Password)
	userModel.FaceUrl = registerModel.PrefixCosUrl + registerModel.FileName
	userModel.FaceToken = faceToken
	userModel.CreateTime = time.Now().Unix()

	if err := model.AddUser(userModel); err == nil {
		c.JSON(http.StatusCreated, gin.H{"id": userModel.Id})
	} else {
		renderJSONWithError(c, err.Error())
	}
}

// UserList 用户列表
func UserList(c *gin.Context) {
	list := make([]*model.UserModel, 0)
	list, err := model.GetAllUser()
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// Binding from Login JSON
type LoginModel struct {
	Username     string `json:"username" binding:"required"`
	PrefixCosUrl string `json:"prefixCosUrl" binding:"required"`
	FileName     string `json:"fileName" binding:"required"`
}

// Login 登录
func Login(c *gin.Context) {
	loginUserModel := new(LoginModel)
	if err := c.ShouldBindWith(&loginUserModel, binding.JSON); err != nil {
		renderJSONWithError(c, err.Error())
		return
	}

	user, err := model.GetUserByUsername(loginUserModel.Username)
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}

	isMatchFace, err := face.IsMatchFace(loginUserModel.PrefixCosUrl, loginUserModel.FileName, user.FaceToken)
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	if !isMatchFace {
		renderJSONWithError(c, "拍摄照片中未检测到该用户人脸")
		return
	}

	user.Password = ""
	user.FaceToken = ""
	c.JSON(http.StatusOK, user)
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		renderJSONWithError(c, "输入删除用户id非法")
		return
	}
	err = model.DeleteUser(intId)
	if err != nil {
		renderJSONWithError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func renderJSONWithError(c *gin.Context, error string) {
	logger.Error(error)
	c.JSON(http.StatusBadRequest, gin.H{"error": error})
	return
}
