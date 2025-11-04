package errno

import (
	"github.com/xh-polaris/psych-profile/pkg/errorx/code"
)

const (
	ErrUnAuth        = 1000
	ErrUnImplement   = 888
	ErrInvalidParams = 999
)

func init() {
	code.Register(
		ErrUnAuth,
		"身份认证失败",
		code.WithAffectStability(false),
	)
	code.Register(
		ErrUnImplement,
		"功能暂未实现",
		code.WithAffectStability(true),
	)
	code.Register(
		ErrInvalidParams,
		"参数{field}={type}错误",
		code.WithAffectStability(false),
	)
}

//// 定义常量错误
//// TODO: 定义错误常量
//var (
//	// User模块相关错误码
//	// User模块相关错误码
//	ErrUserSignUp           = errorx.New(10001, "用户注册失败，请重试")
//	ErrUserPhoneExist       = errorx.New(10002, "该手机号已被注册")
//	ErrUserVerify           = errorx.New(10003, "验证码错误")
//	ErrUserPasswordMismatch = errorx.New(10004, "密码不匹配")
//	ErrUserNotExist         = errorx.New(10005, "用户账号不存在")
//	ErrUserGetInfo          = errorx.New(10006, "获取用户信息失败")
//	ErrUserSignIn           = errorx.New(10007, "用户登录失败")
//
//	// Unit模块相关错误码
//	ErrUnitSignUp           = errorx.New(11001, "单位注册失败，请重试")
//	ErrUnitPhoneExist       = errorx.New(11002, "该手机号已被注册为单位账号")
//	ErrUnitVerify           = errorx.New(11003, "单位账号验证失败")
//	ErrUnitPasswordMismatch = errorx.New(11004, "密码不匹配")
//	ErrUnitNotExist         = errorx.New(11005, "单位账号不存在")
//	ErrUnitLinkUser         = errorx.New(11006, "关联用户失败")
//	ErrUnitGetInfo          = errorx.New(11007, "获取单位信息失败")
//	ErrUnitCreateUser       = errorx.New(11008, "创建用户失败")
//)
//
//// ErrInvalidParams 调用时错误
//var (
//	ErrInvalidParams = NewErrno(codes.InvalidArgument, errors.New("参数错误"))
//	ErrCall          = NewErrno(codes.Unknown, errors.New("调用接口失败，请重试"))
//)
//
//// 数据库相关错误
//var (
//	ErrNotFound        = NewErrno(codes.NotFound, errors.New("not found"))
//	ErrInvalidObjectId = NewErrno(codes.InvalidArgument, errors.New("无效的id "))
//)
