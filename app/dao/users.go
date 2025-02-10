package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"loan-ext/app/entity/models"
)

type UsersDao struct {
	db     *gorm.DB
	models *models.Users
}

func NewUsersDao(db *gorm.DB) *UsersDao {
	if db == nil {
		db = ant.Db()
	}
	return &UsersDao{db: db}
}

// Create
func (dao *UsersDao) Create(users *models.Users) error {
	return dao.db.Create(&users).Error
}

// DeleteById
func (dao *UsersDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *UsersDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *UsersDao) Update(users models.Users) error {
	return dao.db.Updates(&users).Error
}

// GetList
func (dao *UsersDao) GetList() (list []models.Users) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *UsersDao) GetPage(page page.PageParam) (list []models.Users, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("phone", "=", page.FilterMap["phone"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *UsersDao) GetById(id int) (row models.Users) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByPhoneAndStatus
func (dao *UsersDao) GetByPhoneAndStatus(phone string) (row *models.Users) {
	dao.db.Model(&dao.models).Where("phone=?", phone).Limit(1).Find(&row)
	return row
}

// UpdateAddVerifiedCount
func (dao *UsersDao) UpdateAddVerifiedCount(id int) error {
	return dao.db.Model(&dao.models).Where("id=?", id).Update("verified_count", gorm.Expr("verified_count + ?", 1)).Error
}

// UpdateAddMatchCount
func (dao *UsersDao) UpdateAddMatchCount(id int) error {
	return dao.db.Model(&dao.models).Where("id=?", id).Update("match_count", gorm.Expr("match_count + ?", 1)).Error
}
