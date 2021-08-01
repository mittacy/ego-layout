package userTransform

import (
	"github.com/mittacy/ego-layout/app/model"
	"github.com/mittacy/ego-layout/app/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// UserPack 数据库数据转化为响应数据
// @param data 数据库数据
// @return reply 响应体数据
// @return err
func UserPack(data *model.User) (*userValidator.GetReply, error) {
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
func UsersPack(data []model.User) (reply []userValidator.ListReply, err error) {
	err = copier.Copy(&reply, &data)
	return
}

// UserToReply 详情响应包装
// @param data 数据库数据
func UserToReply(c *gin.Context, data *model.User) {
	reply, err := UserPack(data)
	if err != nil {
		response.CopierErrAndLog(c, err)
		return
	}

	res := map[string]interface{}{
		"user": reply,
	}

	response.Success(c, res)
}

// UsersToReply 列表响应包装
// @param data 数据库列表数据
// @param totalSize 记录总数
func UsersToReply(c *gin.Context, data []model.User, totalSize int64) {
	list, err := UsersPack(data)
	if err != nil {
		response.CopierErrAndLog(c, err)
		return
	}

	res := map[string]interface{}{
		"list":       list,
		"total_size": totalSize,
	}

	response.Success(c, res)
}
