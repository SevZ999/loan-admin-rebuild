package models

import (
	"gorm.io/gorm"
	"time"
)

type SysConfig struct {
	Id          int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`        //标识
	ConfigKey   string         `gorm:"column:config_key" json:"config_key" form:"config_key" comment:"唯一键名"`                 //唯一键名
	ConfigValue string         `gorm:"column:config_value" json:"config_value" form:"config_value" comment:"配置的值"`           //配置的值
	ConfigType  string         `gorm:"column:config_type;default:null" json:"config_type" form:"config_type" comment:"配置类型"` //配置类型
	Description string         `gorm:"column:description" json:"description" form:"description" comment:"配置描述"`              //配置描述
	Status      *int           `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用"`                     //状态:1=禁用;2=启用
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                 //创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                 //修改时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                          //删除时间
}

func (SysConfig) TableName() string {
	return "loan.sys_config"
}
