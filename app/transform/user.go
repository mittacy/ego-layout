package transform

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mittacy/ego-layout/app/model"
	"github.com/mittacy/ego-layout/app/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/logger"
	"github.com/mittacy/ego-layout/pkg/response"
)

type User struct {
	logger *logger.CustomLogger
}

func NewUser(customLogger *logger.CustomLogger) User {
	return User{logger: customLogger}
}

// UserPack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func (ctl *User) UserPack(data *model.User) (*userValidator.GetReply, error) {
	reply := userValidator.GetReply{}

	if err := copier.Copy(&reply, data); err != nil {
		return nil, err
	}

	return &reply, nil
}

// UsersPack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func (ctl *User) UsersPack(data []model.User) (reply []userValidator.ListReply, err error) {
	err = copier.Copy(&reply, &data)
	return
}

// GetReply 详情响应包装
// @param data 数据库数据
func (ctl *User) GetReply(c *gin.Context, data *model.User) {
	reply, err := ctl.UserPack(data)
	if err != nil {
		ctl.logger.CopierErrLog(err)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"user": reply,
	}

	response.Success(c, res)
}

// ListReply 列表响应包装
// @param data 数据库列表数据
// @param totalSize 记录总数
func (ctl *User) ListReply(c *gin.Context, data []model.User, totalSize int64) {
	ctl.logger.Info("this is transform")
	list, err := ctl.UsersPack(data)
	if err != nil {
		ctl.logger.CopierErrLog(err)
		response.Unknown(c)
		return
	}

	res := map[string]interface{}{
		"list":       list,
		"total_size": totalSize,
	}

	response.Success(c, res)
}
