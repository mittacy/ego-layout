package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego/library/mysql"
	"github.com/mittacy/ego/library/redis"
	"github.com/spf13/viper"
)

type User struct {
	mysql.Gorm
	redis.GoRedis
	cacheKeyPre string
}

func NewUser() User {
	return User{
		Gorm: mysql.Gorm{
			MysqlConfName: "localhost",
		},
		GoRedis: redis.GoRedis{
			RedisConfName: "localhost",
			RedisDB:       0,
		},
		cacheKeyPre: fmt.Sprintf("%s:user", viper.GetString("APP_NAME")),
	}
}

func (ctl *User) GetById(c *gin.Context, id int64) (*model.User, error) {
	//if err := ctl.Redis().Set(context.Background(), "name", "xiyangyang", time.Second * 10).Err(); err != nil {
	//	return nil, err
	//}
	//
	//user := model.User{}
	//if err := ctl.DB().Where("id = ?", id).First(&user).Error; err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, apierr.UserNoExist
	//	}
	//
	//	return nil, errors.WithStack(err)
	//}
	//
	//return &user, nil
	return &model.User{
		Id:        id,
		Name:      "测试",
		Introduce: "测试",
		Password:  "密码",
		Deleted:   0,
		CreatedAt: 0,
		UpdatedAt: 0,
	}, nil
}
