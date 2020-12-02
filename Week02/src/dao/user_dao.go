package dao

import (
	"Week02/src/dbs"
	"Week02/src/models"
	"fmt"
	"github.com/pkg/errors"
)

type UserDAO struct{}

var DefaultUser = UserDAO{}

func (u UserDAO) FindOneByUid(uid interface{}) (*model.UserModelImpl, error) {
	user := new(model.UserModelImpl)
	db := dbs.Orm.Find(user, uid)
	if db.Error != nil {
		return user, errors.Wrap(db.Error, fmt.Sprintf("根据uid: %v,查找用户信息发生错误", uid))
	}
	if db.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("根据uid: %v,未找到用户信息", uid))
	}
	return user, nil
}
