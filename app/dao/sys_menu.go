package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysMenuDao struct {
	db     *gorm.DB
	models *models.SysMenu
}

func NewSysMenuDao(db *gorm.DB) *SysMenuDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysMenuDao{db: db}
}

// Create
func (dao *SysMenuDao) Create(sysMenu models.SysMenu) error {
	return dao.db.Create(&sysMenu).Error
}

// DeleteById
func (dao *SysMenuDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysMenuDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysMenuDao) Update(sysMenu models.SysMenu) error {
	return dao.db.Updates(&sysMenu).Error
}

// GetList
func (dao *SysMenuDao) GetList() (list []models.SysMenu) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysMenuDao) GetPage(page page.PageParam) (list []models.SysMenu, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("title", "LIKE", page.FilterMap["title"]),
		asql.Where("path", "LIKE", page.FilterMap["path"]),
		asql.Where("component", "LIKE", page.FilterMap["component"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Order("sort desc,id asc").Find(&list).Offset(0).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysMenuDao) GetById(id int) (row models.SysMenu) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByIds
func (dao *SysMenuDao) GetByIds(ids []int) (list []models.SysMenu) {
	dao.db.Model(&dao.models).Where("id in ?", ids).Order("sort desc,id asc").Find(&list)
	return list
}
