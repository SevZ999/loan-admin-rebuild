package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/conv"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/app/entity/models"
)

type UserInfoDao struct {
	db     *gorm.DB
	models *models.UserInfo
}

func NewUserInfoDao(db *gorm.DB) *UserInfoDao {
	if db == nil {
		db = ant.Db()
	}
	return &UserInfoDao{db: db}
}

// Create
func (dao *UserInfoDao) Create(userInfo models.UserInfo) error {
	return dao.db.Create(&userInfo).Error
}

// DeleteById
func (dao *UserInfoDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *UserInfoDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *UserInfoDao) Update(userInfo models.UserInfo) error {
	return dao.db.Updates(&userInfo).Error
}

// GetList
func (dao *UserInfoDao) GetList() (list []models.UserInfo) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *UserInfoDao) GetPage(page page.PageParam) (list []models.UserInfo, total int64, err error) {
	// 构建查询条件
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("app_code", "=", page.FilterMap["app_code"]),
		asql.Where("id_card", "like", conv.String(page.FilterMap["id_card"])+"%"),
		asql.Where("channel_code", "=", page.FilterMap["channel_code"]),
		asql.Where("name", "LIKE", conv.String(page.FilterMap["name"])+"%"),
		asql.Where("institution_info_id", "=", page.FilterMap["institution_info_id"]),
		asql.Where("status", "=", page.FilterMap["status"]),
		asql.Where("credit_status", "=", page.FilterMap["credit_status"]),
		asql.Where("age", "=", page.FilterMap["age"]),
		asql.Where("city_name", "=", page.FilterMap["city_name"]),
		asql.Where("wechat_id", "=", page.FilterMap["wechat_id"]),
		asql.Where("occupation", "=", page.FilterMap["occupation"]),
		asql.Where("push_time", "BETWEEN", page.FilterMap["push_time"]),
		asql.Where("gender", "=", page.FilterMap["gender"]),
		asql.Where("policy_duration", "=", page.FilterMap["policy_duration"]),
		asql.Where("funding_gap", "=", page.FilterMap["funding_gap"]),
		asql.Where("sesame_score", "=", page.FilterMap["sesame_score"]),
		asql.Where("type", "=", page.FilterMap["type"]),
		asql.Where("phone", "LIKE", page.FilterMap["phone"]),
		asql.Where("city_id", "=", page.FilterMap["city_id"]),
		asql.Where("has_car", "=", page.FilterMap["has_car"]),
		asql.Where("has_house", "=", page.FilterMap["has_house"]),
		asql.Where("has_social_security", "=", page.FilterMap["has_social_security"]),
		asql.Where("has_housing_fund", "=", page.FilterMap["has_housing_fund"]),
		asql.Where("has_company", "=", page.FilterMap["has_company"]),
		asql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Preload(clause.Associations).Find(&list).Offset(-1).Limit(1).Count(&total).Error

	return list, total, err
}

// GetById
func (dao *UserInfoDao) GetById(id int) (row models.UserInfo) {
	dao.db.Model(&dao.models).Preload(clause.Associations).Where("id=?", id).Limit(1).Preload(clause.Associations).Find(&row)
	return row
}
