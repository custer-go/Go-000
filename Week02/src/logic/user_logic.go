package logic

import (
	"Week02/src/dao"
	"Week02/src/models"
	//"github.com/sirupsen/logrus"
	"log"
)

type userLogic struct{}

var DefaultUser = userLogic{}

func (u userLogic) FindOne(uid interface{}) (*model.UserModelImpl, error) {
	user, err := dao.DefaultUser.FindOneByUid(uid)
	if err != nil {
		log.Printf("%+v\n", err) // %+v 打印出详细信息，包含堆栈
	}
	return user, err
}
