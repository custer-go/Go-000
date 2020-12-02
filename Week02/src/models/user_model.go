package model

type UserModelImpl struct {
	UserID      int    `json:"id" form:"id"`
	UserName    string `json:"name" form:"name" binding:"required,UserName"`
	UserPwd     string `json:"user_pwd" binding:"required,min=4"`
	UserAddtime string `json:"addtime"`
}

func (UserModelImpl) TableName() string {
	return "users"
}
