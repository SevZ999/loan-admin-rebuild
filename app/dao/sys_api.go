package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysApiDao struct {
	db     *gorm.DB
	models *models.SysApi
}

func NewSysApiDao(db *gorm.DB) *SysApiDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysApiDao{db: db}
}

// Create
func (dao *SysApiDao) Create(sysApi models.SysApi) error {
	return dao.db.Create(&sysApi).Error
}

// DeleteById
func (dao *SysApiDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysApiDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysApiDao) Update(sysApi models.SysApi) error {
	return dao.db.Updates(&sysApi).Error
}

// GetList
func (dao *SysApiDao) GetList() (list []models.SysApi) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysApiDao) GetPage(page page.PageParam) (list []models.SysApi, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("title", "LIKE", page.FilterMap["title"]),
		asql.Where("path", "LIKE", page.FilterMap["path"]),
		asql.Where("group", "LIKE", page.FilterMap["group"]),
		asql.Where("method", "=", page.FilterMap["method"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysApiDao) GetById(id int) (row models.SysApi) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetById
func (dao *SysApiDao) GetByList(ids []int) (list []models.SysApi) {
	dao.db.Model(&dao.models).Where("id IN ?", ids).Find(&list)
	return list
}
