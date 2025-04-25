package models

import (
	"gorm.io/gorm"
	"time"
)

type SysProtocols struct {
	Id        int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"` //标识
	AppCode   int            `gorm:"column:app_code" uri:"id" json:"app_code" form:"app_code" comment:"所属应用"`       //所属应用
	Name      string         `gorm:"column:name" json:"name" form:"name" comment:"协议名称"`                            //协议名称
	Type      string         `gorm:"column:type;default:null" json:"type" form:"type" comment:"协议类型，如用户协议/隐私政策"`    //协议类型，如用户协议/隐私政策
	Content   string         `gorm:"column:content" json:"content" form:"content" comment:"协议内容"`                   //协议内容
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`          //创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`          //更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                   //删除时间
}

func (SysProtocols) TableName() string {
	return "loan.sys_protocols"
}
