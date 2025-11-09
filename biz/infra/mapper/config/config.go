package config

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string `json:"appId,omitempty" bson:"appId,omitempty"`
}

type TTS struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string `json:"appId,omitempty" bson:"appId,omitempty"`
}

type Report struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Provider    string `json:"provider,omitempty" bson:"provider,omitempty"`
	AppID       string `json:"appId,omitempty" bson:"appId,omitempty"`
}

type Config struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UnitID     primitive.ObjectID `json:"unitId,omitempty" bson:"unitId,omitempty"`
	Type       int32              `json:"type,omitempty" bson:"type,omitempty"` // Chain | End2End
	Chat       *Chat              `json:"chat,omitempty" bson:"chat,omitempty"`
	TTS        *TTS               `json:"tts,omitempty" bson:"tts,omitempty"`
	Report     *Report            `json:"report,omitempty" bson:"report,omitempty"`
	Status     int32              `json:"status,omitempty" bson:"status,omitempty"`
	CreateTime int64              `json:"createTime,omitempty" bson:"createTime,omitempty"`
	UpdateTime int64              `json:"updateTime,omitempty" bson:"updateTime,omitempty"`
	DeleteTime int64              `json:"deleteTime,omitempty" bson:"deleteTime,omitempty"`
}
