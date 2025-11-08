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
	"github.com/xh-polaris/psych-profile/biz/infra/util/encrypt"
	"github.com/xh-polaris/psych-profile/biz/infra/util/enum"
	"github.com/xh-polaris/psych-profile/biz/infra/util/reg"
	"github.com/xh-polaris/psych-profile/pkg/errorx"
	"github.com/xh-polaris/psych-profile/pkg/logs"
	"github.com/xh-polaris/psych-profile/types/errno"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ IUnitService = (*UnitService)(nil)

type IUnitService interface {
	UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (*profile.UnitSignUpResp, error)
	UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (*profile.UnitSignInResp, error)
	UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (*profile.UnitGetInfoResp, error)
	UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (*basic.Response, error)
	UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (*basic.Response, error)
	UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (*basic.Response, error)
	UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (*basic.Response, error)
}

type UnitService struct {
	UnitMapper unit.IMongoMapper
	UserMapper user.IMongoMapper
}

var UnitServiceSet = wire.NewSet(
	wire.Struct(new(UnitService), "*"),
	wire.Bind(new(IConfigService), new(*UnitService)),
)

func (u *UnitService) UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (*profile.UnitSignUpResp, error) {
	// 参数校验
	if req.Unit == nil {
		return nil, errorx.New(errno.ErrMissingEntity, errorx.KV("entity", "单位用户"))
	}
	if req.Unit.Name == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位名称"))
	}
	if req.Unit.Phone == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "电话号码"))
	}
	if req.Unit.Password == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
	}

	if !reg.CheckMobile(req.Unit.Phone) {
		return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "电话号码"))
	}

	// 检查手机号是否已注册
	if unitDAO, err := u.UnitMapper.FindOneByPhone(ctx, req.Unit.Phone); err != nil {
		logs.Errorf("find unit by phone error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	} else if unitDAO != nil {
		return nil, errorx.New(errno.ErrPhoneAlreadyExist)
	}
	if userDAO, err := u.UserMapper.FindOneByPhone(ctx, req.Unit.Phone); err != nil {
		logs.Errorf("find user by phone error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	} else if userDAO != nil {
		return nil, errorx.New(errno.ErrPhoneAlreadyExist)
	}

	// 密码加密
	hashedPwd, err := encrypt.BcryptEncrypt(req.Unit.Password)
	if err != nil {
		logs.Errorf("bcrypt encrypt error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 构造 Unit
	unitDAO := &unit.Unit{
		ID:         primitive.NewObjectID(),
		Phone:      req.Unit.Phone,
		Password:   hashedPwd,
		Name:       req.Unit.Name,
		Address:    req.Unit.Address,
		Contact:    req.Unit.Contact,
		Level:      int(req.Unit.Level),
		Status:     enum.Active,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	// 插入数据库
	if err := u.UnitMapper.Insert(ctx, unitDAO); err != nil {
		logs.Errorf("insert unit error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 获得单位状态
	statusStr, ok := enum.GetStatus(unitDAO.Status)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}

	// 构造返回结果
	return &profile.UnitSignUpResp{
		Unit: &profile.Unit{
			Id:         unitDAO.ID.Hex(),
			Phone:      unitDAO.Phone,
			Name:       unitDAO.Name,
			Address:    unitDAO.Address,
			Contact:    unitDAO.Contact,
			Level:      int32(unitDAO.Level),
			Status:     statusStr,
			CreateTime: unitDAO.CreateTime,
			UpdateTime: unitDAO.UpdateTime,
			DeleteTime: unitDAO.DeleteTime,
		},
	}, nil
}

func (u *UnitService) UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (*profile.UnitSignInResp, error) {
	// 参数校验
	if req.Phone == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "电话号码"))
	}
	if req.AuthType == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证方式"))
	}
	if req.VerifyCode == "" && req.AuthType == cst.AuthTypePhonePassword {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
	}
	if req.VerifyCode == "" && req.AuthType == cst.AuthTypePhoneCode {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证码"))
	}

	if !reg.CheckMobile(req.Phone) {
		return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "电话号码"))
	}

	// 验证方式
	unitDAO := &unit.Unit{}
	switch req.AuthType {
	// 密码登录
	case cst.AuthTypePhonePassword:
		// 获得用户
		unitDAO, err := u.UnitMapper.FindOneByPhone(ctx, req.Phone)
		if err != nil {
			logs.Errorf("find unit by phone error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		} else if unitDAO == nil {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}

		// 获得密码
		if !encrypt.BcryptCheck(unitDAO.Password, req.VerifyCode) {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}
	// 验证码登录
	case cst.AuthTypePhoneCode:
		return nil, errorx.New(errno.ErrUnImplement) // TODO: 验证码登录
	}

	// 构造返回结果
	return &profile.UnitSignInResp{UnitId: unitDAO.ID.Hex()}, nil
}

func (u *UnitService) UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (*profile.UnitGetInfoResp, error) {
	// 参数校验
	if req.UnitId == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位ID"))
	}

	unitId, err := primitive.ObjectIDFromHex(req.UnitId)
	logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
	if err != nil {
		return nil, err
	}

	// 查询单位
	unitDAO, err := u.UnitMapper.FindOne(ctx, unitId)
	if err != nil {
		logs.Errorf("find unit error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 获得单位状态
	statusStr, ok := enum.GetStatus(unitDAO.Status)
	if !ok {
		return nil, errorx.New(errno.ErrInternalError)
	}

	// 构造返回结果
	return &profile.UnitGetInfoResp{
		Unit: &profile.Unit{
			Id:         unitDAO.ID.Hex(),
			Phone:      unitDAO.Phone,
			Name:       unitDAO.Name,
			Address:    unitDAO.Address,
			Contact:    unitDAO.Contact,
			Level:      int32(unitDAO.Level),
			Status:     statusStr,
			CreateTime: unitDAO.CreateTime,
			UpdateTime: unitDAO.UpdateTime,
			DeleteTime: unitDAO.DeleteTime,
		},
	}, nil
}

func (u *UnitService) UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (*basic.Response, error) {
	// 参数校验
	if req.Unit.Id == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位ID"))
	}

	// 不允许修改手机号、密码、验证方式、level、状态
	// 密码、验证方式需要通过其他接口修改
	unitId, err := primitive.ObjectIDFromHex(req.Unit.Id)
	if err != nil {
		logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 构建更新字段
	update := make(bson.M)
	if req.Unit.Name != "" {
		update[cst.Name] = req.Unit.Name
	}
	if req.Unit.Address != "" {
		update[cst.Address] = req.Unit.Address
	}
	if req.Unit.Contact != "" {
		update[cst.Contact] = req.Unit.Contact
	}
	update[cst.UpdateTime] = time.Now().Unix()

	// 一次更新所有字段
	if len(update) > 0 {
		if err = u.UnitMapper.UpdateField(ctx, unitId, update); err != nil {
			logs.Errorf("update unit error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
	}

	// 构造返回结果
	return &basic.Response{}, nil
}

func (u *UnitService) UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (*basic.Response, error) {
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

	unitId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 验证方式
	unitDAO := &unit.Unit{}
	switch req.AuthType {
	// 验证码
	case cst.AuthTypePhoneCode:
		return nil, errorx.New(errno.ErrUnImplement) // TODO: 验证码登录
	// 密码
	case cst.AuthTypePhonePassword:
		// 获取密码
		unitDAO, err := u.UnitMapper.FindOne(ctx, unitId)
		if err != nil {
			logs.Errorf("find unit by phone error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
		if !encrypt.BcryptCheck(unitDAO.Password, req.AuthValue) {
			return nil, errorx.New(errno.ErrWrongAccountOrPassword)
		}
	}

	// 加密密码
	newPwd, err := encrypt.BcryptEncrypt(req.NewPassword)
	if err != nil {
		logs.Errorf("bcrypt encrypt error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 更新密码
	if err = u.UnitMapper.UpdateField(ctx, unitDAO.ID, bson.M{
		cst.Password:   newPwd,
		cst.UpdateTime: time.Now().Unix(),
	}); err != nil {
		logs.Errorf("update unit error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	// 构造返回结果
	return &basic.Response{}, nil
}

func (u *UnitService) UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (*basic.Response, error) {
	// 参数校验
	if req.UnitId == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位ID"))
	}
	if req.UserId == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "用户ID"))
	}

	// 转换ID
	unitId, err := primitive.ObjectIDFromHex(req.UnitId)
	if err != nil {
		logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}
	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		logs.Errorf("parse user id error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	if err := u.UnitMapper.UpdateField(ctx, userId, bson.M{cst.UnitID: unitId}); err != nil {
		logs.Errorf("update user error: %s", errorx.ErrorWithoutStack(err))
		return nil, err
	}

	return &basic.Response{}, nil
}

func (u *UnitService) UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (*basic.Response, error) {
	// 参数校验
	if req.UnitId == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "单位ID"))
	}
	if req.CodeType == "" {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "验证方式"))
	}
	if len(req.Users) == 0 {
		return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "用户列表"))
	}

	// 验证方式标记
	isAuthTypePhone := req.CodeType == cst.AuthTypePhoneCode

	// 插入用户
	for _, userReq := range req.Users {
		// 参数校验
		if userReq.Code == "" && isAuthTypePhone {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "电话"))
		}
		if userReq.Code == "" && !isAuthTypePhone {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "学号"))
		}
		if userReq.Name == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "姓名"))
		}
		if userReq.Password == "" {
			return nil, errorx.New(errno.ErrMissingParams, errorx.KV("field", "密码"))
		}

		if isAuthTypePhone {
			// 如果说验证方式是手机，则需要检测手机号的格式
			if !reg.CheckMobile(userReq.Code) {
				return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "手机号"))
			}

			// 检查手机号是否已注册
			if unitDAO, err := u.UnitMapper.FindOneByPhone(ctx, userReq.Code); err != nil {
				logs.Errorf("find unit by phone error: %s", errorx.ErrorWithoutStack(err))
				return nil, err
			} else if unitDAO != nil {
				return nil, errorx.New(errno.ErrPhoneAlreadyExist)
			}
			if userDAO, err := u.UserMapper.FindOneByPhone(ctx, userReq.Code); err != nil {
				logs.Errorf("find user by phone error: %s", errorx.ErrorWithoutStack(err))
				return nil, err
			} else if userDAO != nil {
				return nil, errorx.New(errno.ErrPhoneAlreadyExist)
			}
		} else {
			// 检查学号是否已注册
			if userDAO, err := u.UserMapper.FindOneByStudentID(ctx, userReq.Code); err != nil {
				logs.Errorf("find unit by student id error: %s", errorx.ErrorWithoutStack(err))
				return nil, err
			} else if userDAO != nil {
				return nil, errorx.New(errno.ErrStudentIDAlreadyExist)
			}
		}

		// 加密密码
		hashedPwd, err := encrypt.BcryptEncrypt(userReq.Password)
		if err != nil {
			logs.Errorf("bcrypt encrypt error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}

		// 提取枚举值
		gender, ok := enum.ParseGender(userReq.Gender)
		if !ok {
			return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "性别"))
		}
		codeType, ok := enum.ParseCodeType(req.CodeType)
		if !ok {
			return nil, errorx.New(errno.ErrInvalidParams, errorx.KV("field", "验证方式"))
		}

		// 转换ID
		var unitId primitive.ObjectID
		if req.UnitId != "" {
			unitId, err = primitive.ObjectIDFromHex(req.UnitId)
			if err != nil {
				logs.Errorf("parse unit id error: %s", errorx.ErrorWithoutStack(err))
				return nil, err
			}
		}

		// 构造用户
		userDAO := &user.User{
			ID:         primitive.NewObjectID(),
			CodeType:   codeType,
			Code:       userReq.Code,
			Password:   hashedPwd,
			Name:       userReq.Name,
			Birth:      userReq.Birth,
			Gender:     gender,
			Status:     enum.Active,
			Class:      userReq.Class,
			Grade:      userReq.Grade,
			EnrollYear: userReq.EnrollYear,
			UnitID:     unitId,
			UpdateTime: time.Now().Unix(),
			CreateTime: time.Now().Unix(),
		}

		// 插入用户
		if err = u.UserMapper.Insert(ctx, userDAO); err != nil {
			logs.Errorf("insert user error: %s", errorx.ErrorWithoutStack(err))
			return nil, err
		}
	}

	return &basic.Response{}, nil
}
