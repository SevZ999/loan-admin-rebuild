package models

import (
	"gorm.io/gorm"
	"time"
)

type InstitutionInfoGeoStatistics struct {
	Id                  int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                         //标识
	Date                string         `gorm:"column:date" json:"date" form:"date" comment:"每日日期"`                                                    //每日日期
	InstitutionInfoId   *int           `gorm:"column:institution_info_id" json:"institution_info_id" form:"institution_info_id" comment:"机构标识"`       //机构标识
	InstitutionInfoName string         `gorm:"column:institution_info_name" json:"institution_info_name" form:"institution_info_name" comment:"机构名称"` //机构名称
	Price               *float64       `gorm:"column:price" json:"price" form:"price" comment:"价格"`                                                   //价格
	GeoId               *int           `gorm:"column:geo_id" json:"geo_id" form:"geo_id" comment:"城市标识"`                                              //城市标识
	GeoName             string         `gorm:"column:geo_name" json:"geo_name" form:"geo_name" comment:"城市名称"`                                        //城市名称
	Count               *int           `gorm:"column:count" json:"count" form:"count" comment:"总数"`                                                   //总数
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                  //创建时间
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                  //更新时间
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                           //删除时间
}

func (InstitutionInfoGeoStatistics) TableName() string {
	return "loan.institution_info_geo_statistics"
}
