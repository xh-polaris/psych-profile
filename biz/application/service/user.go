package service

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/infra/cst"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/unit"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/user"
	"github.com/xh-polaris/psych-profile/biz/infra/util/convert"
	"github.com/xh-polaris/psych-profile/biz/infra/util/encrypt"
	"github.com/xh-polaris/psych-profile/biz/infra/util/enum"
	"github.com/xh-polaris/psych-profile/biz/infra/util/reg"
	"github.com/xh-polaris/psych-profile/pkg/errorx"
	"github.com/xh-polaris/psych-profile/pkg/logs"
	"github.com/xh-polaris/psych-profile/types/errno"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	UserSignUp(ctx context.Context, req *profile.UserSignUpReq) (*profile.UserSignUpResp, error)
	UserSignIn(ctx context.Context, req *profile.UserSignInReq) (*profile.UserSignInResp, error)
	UserGetInfo(ctx context.Context, req *profile.UserGetInfoReq) (*profile.UserGetInfoResp, error)
	UserUpdateInfo(ctx context.Context, req *profile.UserUpdateInfoReq) (*basic.Response, error)
	UserUpdatePassword(ctx context.Context, req *profile.UserUpdatePasswordReq) (*basic.Response, error)
}

type UserService struct {
	UserMapper user.IMongoMapper
	UnitMapper unit.IMongoMapper
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)

func (u *UserService) UserSignUp(ctx context.Context, req *profile.UserSignUpReq) (*profile.UserSignUpResp, error) {
	// 默认用户通过注册接口，使用手机号注册
	// 参数校验
	if req.User == nil {
		return nil, errorx.New(errno.ErrMissingEntity, errorx.KV("entity", "用户"))
	}
	if req.User.Code == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "电话号码"))
	}
	if req.User.Password == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
	}
	if req.User.Name == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "姓名"))
	}

	// 手机号格式校验
	if !reg.CheckMobile(req.User.Code) {
		return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "电话号码"))
	}

	// 检查手机号是否已注册
	if exists, err := u.UserMapper.ExistsByPhone(ctx, req.User.Code); err != nil {
		logs.Errorf("check phone exists error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	} else if exists {
		return nil, errorx.New(errno.ErrPhoneAlreadyExist)
	}

	// 密码加密
	hashedPwd, err := encrypt.BcryptEncrypt(req.User.Password)
	if err != nil {
		logs.Errorf("bcrypt encrypt error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 转换枚举值
	gender, ok := enum.ParseGender(req.User.Gender)
	if !ok {
		return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "性别"))
	}

	// 转换ID
	var unitId primitive.ObjectID
	if req.User.UnitId != "" {
		unitId, err = primitive.ObjectIDFromHex(req.User.UnitId)
		if err != nil {
			logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
	}

	// 构造用户
	userDAO := &user.User{
		ID:         primitive.NewObjectID(),
		CodeType:   enum.CodeTypePhone,
		Code:       req.User.Code,
		Password:   hashedPwd,
		Name:       req.User.Name,
		Birth:      req.User.Birth,
		Gender:     gender,
		Status:     enum.Active,
		Class:      req.User.Class,
		Grade:      req.User.Grade,
		EnrollYear: req.User.EnrollYear,
		UnitID:     unitId,
		UpdateTime: time.Now().Unix(),
		CreateTime: time.Now().Unix(),
	}

	// 插入用户
	if err = u.UserMapper.Insert(ctx, userDAO); err != nil {
		logs.Errorf("insert user error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 获得枚举值
	genderStr, ok := enum.GetGender(userDAO.Gender)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}
	statusStr, ok := enum.GetStatus(userDAO.Status)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}
	codeTypeStr, ok := enum.GetCodeType(userDAO.CodeType)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}

	// 构造响应
	return &profile.UserSignUpResp{
		User: &profile.User{
			Id:         userDAO.ID.Hex(),
			CodeType:   codeTypeStr,
			Code:       userDAO.Code,
			Name:       userDAO.Name,
			Gender:     genderStr,
			Birth:      userDAO.Birth,
			Class:      userDAO.Class,
			Grade:      userDAO.Grade,
			EnrollYear: userDAO.EnrollYear,
			Status:     statusStr,
			CreateTime: userDAO.CreateTime,
			UpdateTime: userDAO.UpdateTime,
		},
	}, nil
}

func (u *UserService) UserSignIn(ctx context.Context, req *profile.UserSignInReq) (*profile.UserSignInResp, error) {
	// 参数校验
	if req.AuthType == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "登录方式"))
	}

	var strong bool // 强密码
	var err error
	userDAO := &user.User{}

	switch req.AuthType {
	// 手机号密码登录
	case cst.AuthTypePhonePassword:
		if req.AuthId == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "手机号"))
		}
		if req.AuthValue == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
		}
		if !reg.CheckMobile(req.AuthId) {
			return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "电话号码"))
		}

		// 获得用户
		if userDAO, err = u.UserMapper.FindOneByPhone(ctx, req.AuthId); err != nil {
			logs.Errorf("find user by phone error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		} else if userDAO == nil {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		// 密码验证
		if !encrypt.BcryptCheck(req.AuthValue, userDAO.Password) {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		strong = true

	// 手机号验证码登录
	case cst.AuthTypePhoneCode:
		if req.AuthId == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "手机号"))
		}
		if req.AuthValue == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证码"))
		}
		if !reg.CheckMobile(req.AuthId) {
			return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "电话号码"))
		}

		strong = true

		return nil, errorx.New(errno.ErrUnImplement) // TODO: 验证码

	// 学号密码登录
	case cst.AuthTypeStudentIDPassword:
		if req.AuthId == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "学号"))
		}
		if req.AuthValue == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
		}

		// 获得用户
		if userDAO, err := u.UserMapper.FindOneByStudentID(ctx, req.AuthId); err != nil {
			logs.Errorf("find user by studentId error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		} else if userDAO == nil {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		// 密码验证
		if !encrypt.BcryptCheck(req.AuthValue, userDAO.Password) {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		strong = true

	// 弱验证登录
	case cst.AuthTypeWeakAuth:
		if req.AuthId == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "账号"))
		}
		if req.AuthValue == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
		}

		// 获得用户
		if userDAO, err := u.UserMapper.FindOneByAccount(ctx, req.AuthId); err != nil {
			logs.Errorf("find user by account error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		} else if userDAO == nil {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		// 密码验证
		if !encrypt.BcryptCheck(req.AuthValue, userDAO.Password) {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		strong = false
	}

	return &profile.UserSignInResp{
		UnitId: userDAO.UnitID.Hex(),
		UserId: userDAO.ID.Hex(),
		Strong: strong,
	}, nil
}

func (u *UserService) UserGetInfo(ctx context.Context, req *profile.UserGetInfoReq) (*profile.UserGetInfoResp, error) {
	// 参数校验
	if req.UserId == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "用户ID"))
	}

	// 转换用户ID
	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		logs.Errorf("parse user id error: %s", errorx.ErrorWithoutStack(err))
		return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "用户ID"))
	}

	// 获得用户
	userDAO, err := u.UserMapper.FindOne(ctx, userId)
	if err != nil {
		logs.Errorf("find user error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 获得枚举值
	genderStr, ok := enum.GetGender(userDAO.Gender)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}
	statusStr, ok := enum.GetStatus(userDAO.Status)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}
	codeTypeStr, ok := enum.GetCodeType(userDAO.CodeType)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}

	optionsAny, ok := convert.OptionsToAny(userDAO.Options)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}

	return &profile.UserGetInfoResp{
		User: &profile.User{
			Id:         userDAO.ID.Hex(),
			CodeType:   codeTypeStr,
			Code:       userDAO.Code,
			UnitId:     userDAO.UnitID.Hex(),
			Name:       userDAO.Name,
			Gender:     genderStr,
			Birth:      userDAO.Birth,
			Status:     statusStr,
			EnrollYear: userDAO.EnrollYear,
			Class:      userDAO.Class,
			Grade:      userDAO.Grade,
			Options:    optionsAny,
			CreateTime: userDAO.CreateTime,
			UpdateTime: userDAO.UpdateTime,
			DeleteTime: userDAO.DeleteTime,
		},
	}, nil
}

func (u *UserService) UserUpdateInfo(ctx context.Context, req *profile.UserUpdateInfoReq) (*basic.Response, error) {
	// 参数校验
	if req.User.Id == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "用户ID"))
	}

	// 不允许修改手机号、密码、验证方式、单位ID、状态
	// 密码、验证方式需要通过其他接口修改
	userId, err := primitive.ObjectIDFromHex(req.User.Id)
	if err != nil {
		logs.Errorf("parse user id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 构建更新字段
	update := make(bson.M)
	if req.User.Name != "" {
		update[cst.Name] = req.User.Name
	}
	if req.User.Gender != "" {
		gender, ok := enum.ParseGender(req.User.Gender)
		if !ok {
			return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "性别"))
		}
		update[cst.Gender] = gender
	}
	if req.User.Birth != 0 {
		update[cst.Birth] = req.User.Birth
	}
	if req.User.EnrollYear != 0 {
		update[cst.EnrollYear] = req.User.EnrollYear
	}
	if req.User.Class != 0 {
		update[cst.Class] = req.User.Class
	}
	if req.User.Grade != 0 {
		update[cst.Grade] = req.User.Grade
	}
	// TODO: Options

	update[cst.UpdateTime] = time.Now().Unix()

	// 一次更新所有字段
	if len(update) > 0 {
		if err = u.UserMapper.UpdateField(ctx, userId, update); err != nil {
			logs.Errorf("update user error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
	}

	// 构造返回结果
	return &basic.Response{}, nil
}

func (u *UserService) UserUpdatePassword(ctx context.Context, req *profile.UserUpdatePasswordReq) (*basic.Response, error) {
	// 参数校验
	if req.Id == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位ID"))
	}
	if req.AuthType == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证方式"))
	}
	if req.AuthValue == "" && req.AuthType == cst.AuthTypePhonePassword {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "旧密码"))
	}
	if req.AuthValue == "" && req.AuthType == cst.AuthTypePhoneCode {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证码"))
	}
	if req.NewPassword == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "新密码"))
	}

	userId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logs.Errorf("parse user id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 验证方式
	userDAO := &user.User{}
	switch req.AuthType {
	// 验证码
	case cst.AuthTypeCode:
		return nil, errorx.New(errno.ErrUnImplement) // TODO: 验证码登录
	// 密码
	case cst.AuthTypeOldPassword:
		// 获取密码
		userDAO, err := u.UserMapper.FindOne(ctx, userId)
		if err != nil {
			logs.Errorf("find user by phone error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
		if !encrypt.BcryptCheck(req.AuthValue, userDAO.Password) {
			return nil, errorx.New(errno.ErrWrongPassword)
		}
	}

	// 加密密码
	newPwd, err := encrypt.BcryptEncrypt(req.NewPassword)
	if err != nil {
		logs.Errorf("bcrypt encrypt error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 更新密码
	if err = u.UserMapper.UpdateField(ctx, userDAO.ID, bson.M{
		cst.Password:   newPwd,
		cst.UpdateTime: time.Now().Unix(),
	}); err != nil {
		logs.Errorf("update user error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 构造返回结果
	return &basic.Response{}, nil
}
