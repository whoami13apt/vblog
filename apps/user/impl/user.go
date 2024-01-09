package impl

import (
	"context"

	"gitlab.com/go-course-project/go13/vblog/apps/user"
)

// 实现 user.Service
// 怎么判断这个服务有没有实现这个接口喃？
// &UserServiceImpl{} 是会分配内存, 怎么才能不分配内存
// nil 如何生命 *UserServiceImpl 的 nil
// (*UserServiceImpl)(nil) ---> int8 1  int32(1)  (int32)(1)
// nil 就是一个*UserServiceImpl的空指针
var _ user.Service = (*UserServiceImpl)(nil)

// 用户创建
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (
	*user.User, error) {
	return nil, nil
}

// 查询用户列表, 对象列表 [{}]
func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	return nil, nil
}

// 查询用户详情, 通过Id查询,
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (
	*user.User, error) {
	return nil, nil
}
