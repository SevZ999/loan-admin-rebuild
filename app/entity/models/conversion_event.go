package models

import (
	"time"
)

type ConversionEvent struct {
	Id             int       `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:""`
	AppCode        *int      `gorm:"column:app_code" json:"app_code" form:"app_code" comment:""`
	UserId         *int      `gorm:"column:user_id" json:"user_id" form:"user_id" comment:"用户主键id"` //用户主键id
	AppId          string    `gorm:"column:app_id" json:"app_id" form:"app_id" comment:""`
	Channel        string    `gorm:"column:channel" json:"channel" form:"channel" comment:"广告渠道"`               //广告渠道
	ChannelId      *int      `gorm:"column:channel_id" json:"channel_id" form:"channel_id" comment:"广告渠道的主键id"` //广告渠道的主键id
	Imei           string    `gorm:"column:imei" json:"imei" form:"imei" comment:""`
	ImeiMd5        string    `gorm:"column:imei_md5" json:"imei_md5" form:"imei_md5" comment:""`
	Oaid           string    `gorm:"column:oaid" json:"oaid" form:"oaid" comment:""`
	OaidMd5        string    `gorm:"column:oaid_md5" json:"oaid_md5" form:"oaid_md5" comment:""`
	EventType      *int      `gorm:"column:event_type" json:"event_type" form:"event_type" comment:"事件类型1=激活;2=注册;3=认证;4=进件;5=完件"`     //事件类型1=激活;2=注册;3=认证;4=进件;5=完件
	EventTime      time.Time `gorm:"column:event_time" json:"event_time" form:"event_time" comment:"事件发生是俺"`                           //事件发生是俺
	CallbackStatus string    `gorm:"column:callback_status;default:null" json:"callback_status" form:"callback_status" comment:"回调状态"` //回调状态
	CallbackUrl    string    `gorm:"column:callback_url" json:"callback_url" form:"callback_url" comment:"回调连接，需拼接event_type"`         //回调连接，需拼接event_type
	RetryCount     *int      `gorm:"column:retry_count" json:"retry_count" form:"retry_count" comment:"重试次数"`                          //重试次数
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                             //创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                             //修改时间
}

func (ConversionEvent) TableName() string {
	return "loan.conversion_event"
}
