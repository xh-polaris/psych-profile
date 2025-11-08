// generic_mapper.go
package mapper

import (
	"context"

	"github.com/xh-polaris/psych-profile/biz/infra/cst"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMongoMapper[T any] interface {
	FindOneByField(ctx context.Context, field string, value any) (*T, error)
	FindOne(ctx context.Context, id primitive.ObjectID) (*T, error)
	Insert(ctx context.Context, data *T) error
	UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error
}

type mongoMapper[T any] struct {
	conn *monc.Model
}

func NewMongoMapper[T any](conn *monc.Model) IMongoMapper[T] {
	return &mongoMapper[T]{conn: conn}
}

func (m *mongoMapper[T]) FindOneByField(ctx context.Context, field string, value any) (*T, error) {
	result := new(T)
	if err := m.conn.FindOneNoCache(ctx, result, bson.M{field: value}); err != nil {
		return nil, err
	}
	return result, nil
}

func (m *mongoMapper[T]) FindOne(ctx context.Context, id primitive.ObjectID) (*T, error) {
	return m.FindOneByField(ctx, cst.ID, id)
}

func (m *mongoMapper[T]) Insert(ctx context.Context, data *T) error {
	_, err := m.conn.InsertOneNoCache(ctx, data)
	return err
}

func (m *mongoMapper[T]) UpdateField(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := m.conn.UpdateOneNoCache(ctx, id, update)
	return err
}
