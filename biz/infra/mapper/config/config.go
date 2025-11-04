package config

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string    `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string    `json:"appId,omitempty" bson:"appId,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

type TTS struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string    `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string    `json:"appId,omitempty" bson:"appId,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

type Report struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string    `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string    `json:"appId,omitempty" bson:"appId,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
}

type Config struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       int32              `json:"type,omitempty" bson:"type,omitempty"` // Chain | End2End
	Chat       *Chat              `json:"chat,omitempty" bson:"chat,omitempty"`
	TTS        *TTS               `json:"tts,omitempty" bson:"tts,omitempty"`
	Report     *Report            `json:"report,omitempty" bson:"report,omitempty"`
	Status     int32              `json:"status,omitempty" bson:"status,omitempty"`
	CreateTime time.Time          `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime time.Time          `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
	DeleteTime time.Time          `json:"deleteTime,omitempty" bson:"deleteTime,omitempty"`
}
