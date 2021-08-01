package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/copier"
	"github.com/mittacy/ego-layout/app/model"
	"github.com/mittacy/ego-layout/app/transform/userTransform"
	"github.com/mittacy/ego-layout/app/validator/userValidator"
	"github.com/mittacy/ego-layout/pkg/response"
)

type User struct {
	userService IUserService
}

func NewUser(userService IUserService) User {
	return User{userService: userService}
}

type IUserService interface {
	Create(user model.User) (int64, error)
	Delete(id int64) error
	UpdateInfo(user model.User) error
	UpdatePassword(id int64, name string) error
	Get(id int64) (*model.User, error)
	List(page, pageSize int) ([]model.User, int64, error)
}

/**
 * @apiVersion 0.0.1
 * @apiGroup User
 * @api {get} /user 创建用户
 * @apiName User.Create
 *
 * @apiParam {string{1..20}} name 昵称
 * @apiParam {string{1..100}} info 介绍
 * @apiParam {number{1..20}} password 密码
 *
 * @apiSuccess {number} id 用户id
 *
 * @apiSuccessExample {json} Success-Response:
 *     {
 *         "code": 0,
 *         "data": {
 *             "id": 520
 *         },
 *         "msg": "success"
 *     }
 *
 */
func (ctl *User) Create(c *gin.Context) {
	req := userValidator.CreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	user := model.User{}
	if err := copier.Copy(&user, &req); err != nil {
		response.CopierErrAndLog(c, err)
		return
	}

	id, err := ctl.userService.Create(user)
	if err != nil {
		response.CheckErrAndLog(c, "create user", err)
		return
	}

	res := map[string]int64{
		"id": id,
	}
	response.Success(c, res)
}

/**
 * @apiVersion 0.0.1
 * @apiGroup User
 * @api {get} /user 删除用户
 * @apiName User.Delete
 *
 * @apiParam {number{1..}} id 用户id
 *
 * @apiErrorExample {json} 用户不存在
 *     {
 *       "code": 1,
 *       "msg": "对象不存在",
 *       "data": {}
 *     }
 *
 */
func (ctl *User) Delete(c *gin.Context) {
	req := userValidator.DeleteReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	if err := ctl.userService.Delete(req.Id); err != nil {
		response.CheckErrAndLog(c, "delete user", err)
		return
	}

	response.Success(c, nil)
}

/**
 * @apiVersion 0.1.0
 * @apiGroup User
 * @api {patch} /user 编辑用户
 * @apiName User.Update
 *
 * @apiParam {number=1(用户基本信息),2(用户密码)} update_type 更新类型
 * @apiParam {number{1..}} id 用户id
 * @apiParam {string{1..20}} name 昵称, update_type=1时必须
 * @apiParam {string{1..100}} info 介绍, update_type=1时必须
 * @apiParam {number{1..20}} password 密码, update_type=2时必须
 *
 * @apiErrorExample {json} 用户不存在
 *     {
 *       "code": 1,
 *       "msg": "对象不存在",
 *       "data": {}
 *     }
 *
 */
func (ctl *User) Update(c *gin.Context) {
	req := userValidator.UpdateReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	switch req.UpdateType {
	case 1:
		ctl.updateInfo(c)
	case 2:
		ctl.updatePassword(c)
	default:
		response.FailMsg(c, "update_type err")
	}

	return
}

/**
 * @apiVersion 0.0.1
 * @apiGroup User
 * @api {get} /user?id=1 查询用户详情
 * @apiName User.Get
 *
 * @apiParam {number{1..}} id 用户id
 *
 * @apiSuccess {number} id 用户id
 * @apiSuccess {string} name 用户名字
 * @apiSuccess {string} info 用户介绍
 * @apiSuccess {string} created_at 创建时间
 * @apiSuccess {string} updated_at 更新时间
 *
 * @apiSuccessExample {json} Success-Response:
 *     {
 *         "code": 0,
 *         "data": {
 *             "user": {
 *                 "id": 1,
 *                 "name": "name",
 *                 "info": "this is info",
 *                 "created_at": 1627455919,
 *                 "updated_at": 1627659959
 *             }
 *         },
 *         "msg": "success"
 *     }
 *
 * @apiErrorExample {json} 用户不存在
 *     {
 *       "code": 1,
 *       "msg": "对象不存在",
 *       "data": {}
 *     }
 *
 */
func (ctl *User) Get(c *gin.Context) {
	req := userValidator.GetReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	user, err := ctl.userService.Get(req.Id)
	if err != nil {
		response.CheckErrAndLog(c, "get user", err)
		return
	}

	userTransform.UserToReply(c, user)
}

/**
 * @apiVersion 0.0.1
 * @apiGroup User
 * @api {get} /users 用户分页列表
 * @apiName User.List
 *
 * @apiParam {number{1..}} page 页码
 * @apiParam {number{1..50}} page_size 数据分页大小
 *
 * @apiSuccess {number} id 用户id
 * @apiSuccess {string} name 用户名字
 * @apiSuccess {string} created_at 创建时间
 * @apiSuccess {string} updated_at 更新时间
 *
 * @apiSuccessExample {json} Success-Response:
 *     {
 *         "code": 0,
 *         "data": {
 *             "list": [
 *                 {
 *                     "id": 1,
 *                     "name": "Xiao Ming",
 *                     "created_at": 1627455919,
 *                     "updated_at": 1627659959
 *                 },
 *                 {
 *                     "id": 2,
 *                     "name": "Xiao Hong",
 *                     "created_at": 1627455012,
 *                     "updated_at": 1627651203
 *                 },
 *             ],
 *             "total_size": 2
 *         },
 *         "msg": "success"
 *     }
 *
 */
func (ctl *User) List(c *gin.Context) {
	req := userValidator.ListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	user, totalSize, err := ctl.userService.List(req.Page, req.PageSize)
	if err != nil {
		response.CheckErrAndLog(c, "user list", err)
		return
	}

	userTransform.UsersToReply(c, user, totalSize)
}

func (ctl *User) updateInfo(c *gin.Context) {
	req := userValidator.UpdateInfoReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	user := model.User{}
	if err := copier.Copy(&user, &req); err != nil {
		response.CopierErrAndLog(c, err)
		return
	}

	if err := ctl.userService.UpdateInfo(user); err != nil {
		response.CheckErrAndLog(c, "update user info", err)
		return
	}

	response.Success(c, nil)
	return
}

func (ctl *User) updatePassword(c *gin.Context) {
	req := userValidator.UpdatePasswordReq{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		response.ValidateErrAndLog(c, err)
		return
	}

	if err := ctl.userService.UpdatePassword(req.Id, req.Password); err != nil {
		response.CheckErrAndLog(c, "update user name", err)
		return
	}

	response.Success(c, nil)
	return
}
