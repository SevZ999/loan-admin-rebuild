package models

import (
	"gorm.io/gorm"
	"time"
)

type UserInfo struct {
	Id                  int             `gorm:"column:id;primaryKey;autoIncrement;" uri:"id" json:"id" form:"id" comment:"主键ID"`                         //主键ID
	UserId              *int            `gorm:"column:user_id" json:"user_id" form:"user_id" comment:"用户标识"`                                             //用户标识
	Name                string          `gorm:"column:name" json:"name" form:"name" comment:"姓名"`                                                        //姓名
	InstitutionInfoId   *int            `gorm:"column:institution_info_id" json:"institution_info_id" form:"institution_info_id" comment:"推送机构标识"`       //推送机构标识
	Type                *int            `gorm:"column:type" json:"type" form:"type" comment:"类型:1=信用贷;2=车贷;3=企业贷;4=房产贷"`                                 //类型:1=信用贷;2=车贷;3=企业贷;4=房产贷
	CompanyName         string          `gorm:"column:company_name" json:"company_name" form:"company_name" comment:"公司名称/企业名称"`                         //公司名称/企业名称
	BusinessType        string          `gorm:"column:business_type;default:null" json:"business_type" form:"business_type" comment:"企业性质"`              //企业性质
	RegisteredCapital   *int            `gorm:"column:registered_capital" json:"registered_capital" form:"registered_capital" comment:"企业注册资金(万)"`       //企业注册资金(万)
	OperationYears      string          `gorm:"column:operation_years;default:null" json:"operation_years" form:"operation_years" comment:"企业经营年限"`      //企业经营年限
	WechatId            string          `gorm:"column:wechat_id" json:"wechat_id" form:"wechat_id" comment:"微信号"`                                        //微信号
	IdCard              string          `gorm:"column:id_card" json:"id_card" form:"id_card" comment:"身份证"`                                              //身份证
	Phone               string          `gorm:"column:phone" json:"phone" form:"phone" comment:"手机号"`                                                    //手机号
	Gender              string          `gorm:"column:gender;default:null" json:"gender" form:"gender" comment:"性别"`                                     //性别
	Age                 *int            `gorm:"column:age" json:"age" form:"age" comment:"年龄"`                                                           //年龄
	LicensePlate        string          `gorm:"column:license_plate" json:"license_plate" form:"license_plate" comment:"车牌号"`                            //车牌号
	BuyCarType          *int            `gorm:"column:buy_car_type" json:"buy_car_type" form:"buy_car_type" comment:"购车方式:1=全款;2=按揭"`                    //购车方式:1=全款;2=按揭
	CityId              *int            `gorm:"column:city_id" json:"city_id" form:"city_id" comment:"常驻城市标识"`                                           //常驻城市标识
	ChannelCode         string          `gorm:"column:channel_code" json:"channel_code" form:"channel_code" comment:"渠道标识"`                              //H5上游渠道
	AppCode             *int            `gorm:"column:app_code" json:"app_code" form:"app_code" comment:"APP标识"`                                         //APP标识
	FundingGap          string          `gorm:"column:funding_gap;default:null" json:"funding_gap" form:"funding_gap" comment:"资金缺口/借款额度"`               //资金缺口/借款额度
	CreditStatus        string          `gorm:"column:credit_status;default:null" json:"credit_status" form:"credit_status" comment:"信用情况"`              //信用情况
	HasCar              *int            `gorm:"column:has_car" json:"has_car" form:"has_car" comment:"是否有车"`                                             //是否有车
	HasHouse            *int            `gorm:"column:has_house" json:"has_house" form:"has_house" comment:"是否有房产"`                                      //是否有房产
	HouseType           string          `gorm:"column:house_type;default:null" json:"house_type" form:"house_type" comment:"房产类型"`                       //房产类型
	HouseAddress        string          `gorm:"column:house_address" json:"house_address" form:"house_address" comment:"房产位置"`                           //房产位置
	HasSocialSecurity   *int            `gorm:"column:has_social_security" json:"has_social_security" form:"has_social_security" comment:"是否有社保"`        //是否有社保
	HasHousingFund      *int            `gorm:"column:has_housing_fund" json:"has_housing_fund" form:"has_housing_fund" comment:"是否有公积金"`                //是否有公积金
	HasCompany          *int            `gorm:"column:has_company" json:"has_company" form:"has_company" comment:"是否有营业执照"`                              //是否有营业执照
	PolicyDuration      *int            `gorm:"column:policy_duration" json:"policy_duration" form:"policy_duration" comment:"保单年限"`                     //保单年限
	Occupation          string          `gorm:"column:occupation" json:"occupation" form:"occupation" comment:"职业"`                                      //职业
	SesameScore         string          `gorm:"column:sesame_score;default:null" json:"sesame_score" form:"sesame_score" comment:"支付宝芝麻分"`               //支付宝芝麻分
	IsEmployee          *int            `gorm:"column:is_employee" json:"is_employee" form:"is_employee" comment:"是否为上班族"`                               //是否为上班族
	AnnualIncome        float64         `gorm:"column:annual_income" json:"annual_income" form:"annual_income" comment:"年收入（万）"`                         //年收入（万）
	SalaryPaymentStatus string          `gorm:"column:salary_payment_status" json:"salary_payment_status" form:"salary_payment_status" comment:"工资发放情况"` //工资发放情况
	DebtAmount          float64         `gorm:"column:debt_amount" json:"debt_amount" form:"debt_amount" comment:"负债金额（万）"`                              //负债金额（万）
	PushTime            time.Time       `gorm:"column:push_time" json:"push_time" form:"push_time" comment:"推送时间"`                                       //推送时间
	Status              *int            `gorm:"column:status" json:"status" form:"status" comment:"状态:1=正常生成报告"`                                         //状态:1=正常生成报告
	CreatedAt           time.Time       `gorm:"column:created_at" json:"created_at" form:"created_at" comment:"创建时间"`                                    //创建时间
	UpdatedAt           time.Time       `gorm:"column:updated_at" json:"updated_at" form:"updated_at" comment:"更新时间"`                                    //更新时间
	DeletedAt           gorm.DeletedAt  `gorm:"column:deleted_at" json:"-" form:"deleted_at" comment:"删除时间"`                                             //删除时间
	Users               Users           `gorm:"foreignKey:Id;references:UserId" json:"users" form:"users" comment:"用户"`                                  //用户
	InstitutionInfo     InstitutionInfo `gorm:"foreignKey:Id;references:InstitutionInfoId" json:"institution_info" form:"institution_info" comment:"机构"` //机构
	SysGeo              SysGeo          `gorm:"foreignKey:Id;references:CityId" json:"sys_geo" form:"sys_geo" comment:"城市"`                              //城市
	Channel             Channel         `gorm:"foreignKey:ChannelCode;references:ChannelCode" json:"channel" form:"channel" comment:"渠道"`                //渠道

}

func (UserInfo) TableName() string {
	return "loan.user_info"
}
