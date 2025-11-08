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

func (m *mongoMapper) FindOneByPhone(ctx context.Context, phone string) (*Unit, error) {
	return m.FindOneByField(ctx, cst.Phone, phone)
}
