package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysBannelDao struct {
	db     *gorm.DB
	models *models.SysBannel
}

func NewSysBannelDao(db *gorm.DB) *SysBannelDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysBannelDao{db: db}
}

// Create
func (dao *SysBannelDao) Create(sysBannel models.SysBannel) error {
	return dao.db.Create(&sysBannel).Error
}

// DeleteById
func (dao *SysBannelDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysBannelDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysBannelDao) Update(sysBannel models.SysBannel) error {
	return dao.db.Updates(&sysBannel).Error
}

// GetList
func (dao *SysBannelDao) GetList() (list []models.SysBannel) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysBannelDao) GetPage(page page.PageParam) (list []models.SysBannel, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("title", "LIKE", page.FilterMap["title"]), asql.Where("action", "=", page.FilterMap["action"]), asql.Where("status", "=", page.FilterMap["status"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysBannelDao) GetById(id int) (row models.SysBannel) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
