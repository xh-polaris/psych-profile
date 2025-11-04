package unit

import (
	"github.com/xh-polaris/psych-profile/biz/infra/config"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

var _ IMongoMapper = (*mongoMapper)(nil)

const (
	prefixUnitCacheKey = "cache:unit"
	collectionName     = "unit"
)

type IMongoMapper interface {
}

type mongoMapper struct {
	conn *monc.Model
}

func NewMongoMapper(config *config.Config) IMongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, collectionName, config.Cache)
	return &mongoMapper{
		conn: conn,
	}
}
