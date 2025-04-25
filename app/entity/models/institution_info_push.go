package models

import (
	"gorm.io/gorm"
	"time"
)

type InstitutionInfoPush struct {
	Id                   int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                                 //标识
	InstitutionInfoGeoId int            `gorm:"column:institution_info_geo_id" json:"institution_info_geo_id" form:"institution_info_geo_id" comment:"结构城市标识"` //结构城市标识
	InstitutionInfoId    int            `gorm:"column:institution_info_id" json:"institution_info_id" form:"institution_info_id" comment:"机构标识"`               //结构城市标识
	WeekDay              string         `gorm:"column:week_day;default:null" json:"week_day" form:"week_day" comment:"星期几"`                                    //星期几
	StartTime            *string        `gorm:"column:start_time" json:"start_time" form:"start_time" comment:"开始时间"`                                          //开始时间
	EndTime              *string        `gorm:"column:end_time" json:"end_time" form:"end_time" comment:"结束时间"`                                                //结束时间
	CreatedAt            time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                          //创建时间
	UpdatedAt            time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                                          //修改时间
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                   //删除时间
}

func (InstitutionInfoPush) TableName() string {
	return "loan.institution_info_push"
}
