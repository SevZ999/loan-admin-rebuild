package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type Feedback struct {
	Id          int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	UserId      *int           `gorm:"column:user_id" json:"user_id" form:"user_id" comment:"用户标识"`                   //用户标识
	Description string         `gorm:"column:description" json:"description" form:"description" comment:"问题描述"`       //问题描述
	ImageUrl    asql.Json      `gorm:"column:image_url" json:"image_url" form:"image_url" comment:"上传图片路径"`           //上传图片路径
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"反馈时间"`          //反馈时间
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`          //更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
	Users       Users          `gorm:"foreignKey:Id;references:UserId" json:"users" form:"users" comment:"用户"`        //用户
}

func (Feedback) TableName() string {
	return "loan.feedback"
}
