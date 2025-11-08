package unit

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
	prefixUnitCacheKey = "cache:unit"
	collectionName     = "unit"
)

type IMongoMapper interface {
	FindOneByPhone(ctx context.Context, phone string) (*Unit, error)
	FindOne(ctx context.Context, id primitive.ObjectID) (*Unit, error)
	Insert(ctx context.Context, unit *Unit) error
	UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
}

type mongoMapper struct {
	mapper.IMongoMapper[Unit]
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) IMongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, collectionName, config.Cache)
	return &mongoMapper{
		IMongoMapper: mapper.NewMongoMapper[Unit](conn),
		conn:         conn,
	}
}

// FindOneByPhone 根据手机号查询单位
func (m *mongoMapper) FindOneByPhone(ctx context.Context, phone string) (*Unit, error) {
	return m.FindOneByField(ctx, cst.Phone, phone)
}

// ExistsByPhone 根据手机号查询单位是否存在
func (m *mongoMapper) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	return m.ExistsByField(ctx, cst.Phone, phone)
}
