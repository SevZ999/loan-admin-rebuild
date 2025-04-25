package models

import (
	"gorm.io/gorm"
	"time"
)

type InstitutionInfoGeo struct {
	Id                  int                   `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                                           //标识
	InstitutionInfoId   int                   `gorm:"column:institution_info_id" json:"institution_info_id" form:"institution_info_id" comment:"机构标识"`                         //机构标识
	GeoId               *int                  `gorm:"column:geo_id" json:"geo_id" form:"geo_id" comment:"城市标识"`                                                                //城市标识
	GeoName             string                `gorm:"column:geo_name" json:"geo_name" form:"geo_name" comment:"城市名称"`                                                          //城市名称
	MaxLoanAmount       *int                  `gorm:"column:max_loan_amount" json:"max_loan_amount" form:"max_loan_amount" comment:"最高借款金额(万)"`                                //最高借款金额(万)
	AnnualRate          *float64              `gorm:"column:annual_rate" json:"annual_rate" form:"annual_rate" comment:"年利率"`                                                  //年利率
	LoanDuration        *int                  `gorm:"column:loan_duration" json:"loan_duration" form:"loan_duration" comment:"借款期限(月)"`                                        //借款期限(月)
	DailyLimit          *int                  `gorm:"column:daily_limit" json:"daily_limit" form:"daily_limit" comment:"每日量"`                                                  //每日量
	Price               float64               `gorm:"column:price" json:"price" form:"price" comment:"价格"`                                                                     //价格
	MinSesameScoreRange *int                  `gorm:"column:min_sesame_score_range" json:"min_sesame_score_range" form:"min_sesame_score_range" comment:"最小芝麻分区间"`             //最小芝麻分区间
	MaxSesameScoreRange *int                  `gorm:"column:max_sesame_score_range" json:"max_sesame_score_range" form:"max_sesame_score_range" comment:"最大芝麻分区间"`             //最大芝麻分区间
	Type                *int                  `gorm:"column:type" json:"type" form:"type" comment:"类型:1=债务;2=助贷;3=芝麻"`                                                         //类型:1=债务;2=助贷;3=芝麻
	HasCar              *int                  `gorm:"column:has_car" json:"has_car" form:"has_car" comment:"是否有车"`                                                             //是否有车
	HasHouse            *int                  `gorm:"column:has_house" json:"has_house" form:"has_house" comment:"是否有房"`                                                       //是否有房
	HasHousingFund      *int                  `gorm:"column:has_housing_fund" json:"has_housing_fund" form:"has_housing_fund" comment:"是否公积金"`                                 //是否公积金
	HasSocialSecurity   *int                  `gorm:"column:has_social_security" json:"has_social_security" form:"has_social_security" comment:"是否有社保"`                        //是否有社保
	HasCompany          *int                  `gorm:"column:has_company" json:"has_company" form:"has_company" comment:"是否有企业/个体"`                                             //是否有企业/个体
	Status              *int                  `gorm:"column:status" json:"status" form:"status" comment:"状态:1=禁用;2=启用"`                                                        //状态:1=禁用;2=启用
	CreatedAt           time.Time             `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                                    //创建时间
	UpdatedAt           time.Time             `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                                    //更新时间
	DeletedAt           gorm.DeletedAt        `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                             //删除时间
	InstitutionInfoPush []InstitutionInfoPush `gorm:"foreignKey:InstitutionInfoGeoId;references:Id;" json:"institution_info_push" form:"institution_info_push" comment:"推送时间"` //推送时间
}

func (InstitutionInfoGeo) TableName() string {
	return "loan.institution_info_geo"
}
