package models

import (
	"gorm.io/gorm"
	"time"
)

type ChannelStatistics struct {
	Id                  int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                       //标识
	Date                string         `gorm:"column:date" json:"date" form:"date" comment:"每日日期"`                                                  //每日日期
	Type                *int           `gorm:"column:type" json:"type" form:"type" comment:"类型:1=收益统计;2=注册统计;"`                                     //类型:1=收益统计;2=注册统计;
	ChannelName         string         `gorm:"column:channel_name" json:"channel_name" form:"channel_name" comment:"渠道名称"`                          //渠道名称
	ChannelKey          string         `gorm:"column:channel_key" json:"channel_key" form:"channel_key" comment:"合作渠道唯一标识"`                         //合作渠道唯一标识
	TotalCount          *int           `gorm:"column:total_count" json:"total_count" form:"total_count" comment:"总数"`                               //总数
	TotalAmount         float64        `gorm:"column:total_amount" json:"total_amount" form:"total_amount" comment:"总金额"`                           //总金额
	DeductionPercentage float64        `gorm:"column:deduction_percentage" json:"deduction_percentage" form:"deduction_percentage" comment:"扣量百分比"` //扣量百分比
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建日期"`                                //创建日期
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                                //修改时间
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                         //删除时间
}

func (ChannelStatistics) TableName() string {
	return "loan.channel_statistics"
}
