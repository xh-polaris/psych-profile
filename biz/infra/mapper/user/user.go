package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CodeType   int32              `json:"codeType,omitempty" bson:"codeType,omitempty"` // Phone | Code
	Code       string             `json:"code,omitempty" bson:"code,omitempty"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Birth      string             `json:"birth,omitempty" bson:"birth,omitempty"`
	Gender     int32              `json:"gender,omitempty" bson:"gender,omitempty"`
	Status     int32              `json:"status,omitempty" bson:"status,omitempty"`
	CreateTime time.Time          `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime time.Time          `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
	DeleteTime time.Time          `json:"deleteTime,omitempty" bson:"deleteTime,omitempty"`
}
