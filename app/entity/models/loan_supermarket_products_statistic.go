package models

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"gorm.io/gorm"
	"time"
)

type LoanSupermarketProductsStatistic struct {
	Id                          int            `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"标识"`                                                    //标识
	Date                        string         `gorm:"column:date" json:"date" form:"date" comment:"每日日期"`                                                                               //每日日期
	LoanSupermarketProductsId   *int           `gorm:"column:loan_supermarket_products_id" json:"loan_supermarket_products_id" form:"loan_supermarket_products_id" comment:"贷超标识"`       //贷超标识
	LoanSupermarketProductsName string         `gorm:"column:loan_supermarket_products_name" json:"loan_supermarket_products_name" form:"loan_supermarket_products_name" comment:"贷超名称"` //贷超名称
	ClickNum                    *int           `gorm:"column:click_num" json:"click_num" form:"click_num" comment:"点击次数"`                                                                //点击次数
	Price                       float64        `gorm:"column:price" json:"price" form:"price" comment:"单价"`                                                                              //单价
	UserIds                     asql.Json      `gorm:"column:user_ids" json:"user_ids" form:"user_ids" comment:"点击标识"`                                                                   //点击标识
	Ips                         asql.Json      `gorm:"column:ips" json:"ips" form:"ips" comment:"ip"`                                                                                    //ip
	CreatedAt                   time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                                             //创建时间
	UpdatedAt                   time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"修改时间"`                                                             //修改时间
	DeletedAt                   gorm.DeletedAt `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                                                      //删除时间
}

func (LoanSupermarketProductsStatistic) TableName() string {
	return "loan.loan_supermarket_products_statistic"
}
