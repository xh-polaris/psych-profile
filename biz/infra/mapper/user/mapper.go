package user

import (
	"context"

	"github.com/xh-polaris/psych-profile/biz/infra/config"
	"github.com/xh-polaris/psych-profile/biz/infra/cst"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ IMongoMapper = (*mongoMapper)(nil)

const (
	prefixUserCacheKey = "cache:user"
	collectionName     = "user"
)

type IMongoMapper interface {
	FindOneByPhone(ctx context.Context, phone string) (*User, error)
	FindOneByStudentID(ctx context.Context, studentId string) (*User, error)
	FindOneByAccount(ctx context.Context, account string) (*User, error)
	FindOne(ctx context.Context, id primitive.ObjectID) (*User, error)
	Insert(ctx context.Context, user *User) error
	UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByStudentID(ctx context.Context, studentId string) (bool, error)
}

type mongoMapper struct {
	mapper.IMongoMapper[User]
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) IMongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, collectionName, config.Cache)
	return &mongoMapper{
		IMongoMapper: mapper.NewMongoMapper[User](conn),
		conn:         conn,
	}
}

// FindOneByPhone 根据电话号码查询用户
func (m *mongoMapper) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	return m.FindOneByField(ctx, cst.Code, phone)
}

// FindOneByStudentID 根据学号查询用户
func (m *mongoMapper) FindOneByStudentID(ctx context.Context, studentId string) (*User, error) {
	return m.FindOneByField(ctx, cst.Code, studentId)
}

// FindOneByAccount 根据账号查询用户
func (m *mongoMapper) FindOneByAccount(ctx context.Context, account string) (*User, error) {
	return m.FindOneByField(ctx, cst.Code, account)
}

// ExistsByPhone 根据电话号码查询用户是否存在
func (m *mongoMapper) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	return m.ExistsByField(ctx, cst.Code, phone)
}

// ExistsByStudentID 根据学号查询用户是否存在
func (m *mongoMapper) ExistsByStudentID(ctx context.Context, studentId string) (bool, error) {
	return m.ExistsByField(ctx, cst.Code, studentId)
}
