package data

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mittacy/ego-layout/internal/model"
	"github.com/mittacy/ego-layout/pkg/cache"
	"github.com/mittacy/ego-layout/pkg/log"
	"github.com/mittacy/ego-layout/pkg/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var User userData

func init() {
	l := log.New("user")

	User = userData{
		db:          mysql.NewClientByName("localhost"),
		cache:       cache.NewClientByName("localhost", 0),
		cacheKeyPre: fmt.Sprintf("%s:user", viper.GetString("APP_NAME")),
		logger:      l,
	}
}

type userData struct {
	db          *gorm.DB
	cache       *redis.Client
	cacheKeyPre string
	logger      *log.Logger
}

func (ctl *userData) GetById(id int64) (*model.User, error) {
	//user := model.User{}
	//if err := ctl.db.Where("id = ?", id).First(&user).Error; err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, apierr.ErrUserNoExist
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
