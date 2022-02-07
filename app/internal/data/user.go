package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/app/internal/model"
	"github.com/mittacy/ego/library/eMongo"
	"github.com/mittacy/ego/library/eMysql"
	"github.com/mittacy/ego/library/eRedis"
	"github.com/spf13/viper"
)

type User struct {
	eMysql.EGorm
	eRedis.ERedis
	eMongo.EMongo
	cacheKeyPre string
}

func NewUser() User {
	return User{
		EGorm: eMysql.EGorm{
			MysqlConfName: "localhost",
		},
		ERedis: eRedis.ERedis{
			RedisConfName: "localhost",
			RedisDB:       0,
		},
		EMongo: eMongo.EMongo{
			MongoConfName: "localhost",
			CollationName: "user",
		},
		cacheKeyPre: fmt.Sprintf("%s:user", viper.GetString("APP_NAME")),
	}
}

func (ctl *User) GetById(c *gin.Context, id int64) (model.User, error) {
	//if err := ctl.RDB().Set(context.Background(), "name", "xiyangyang", time.Second * 10).Err(); err != nil {
	//	return model.User{}, err
	//}
	//
	//user := model.User{}
	//if err := ctl.GDB().Where("id = ?", id).First(&user).Error; err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return model.User{}, apierr.UserNoExist
	//	}
	//
	//	return model.User{}, errors.WithStack(err)
	//}
	//
	//return user, nil
	return model.User{
		Id:        id,
		Name:      "测试",
		Introduce: "测试",
		Password:  "密码",
		Deleted:   0,
		CreatedAt: 0,
		UpdatedAt: 0,
	}, nil
}
