package models 

import (
"gorm.io/gorm"
"time"
)

type SysHelp struct { 
    Id int `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`//标识
    Type *int `gorm:"column:type" json:"type" form:"type" comment:"类型:1=常见问题;2=借款相关;3=防骗指南;4=账户相关;5=其他问题"`//类型:1=常见问题;2=借款相关;3=防骗指南;4=账户相关;5=其他问题
    Title string `gorm:"column:title" json:"title" form:"title" comment:"标题"`//标题
    Describe string `gorm:"column:describe" json:"describe" form:"describe" comment:"描述"`//描述
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`//创建时间
    UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`//修改时间
    DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`//删除时间
}

func (SysHelp) TableName() string {
	return "loan.sys_help"
}