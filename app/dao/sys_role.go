package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysRoleDao struct {
	db     *gorm.DB
	models *models.SysRole
}

func NewSysRoleDao(db *gorm.DB) *SysRoleDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysRoleDao{db: db}
}

// Create
func (dao *SysRoleDao) Create(sysRole models.SysRole) error {
	return dao.db.Create(&sysRole).Error
}

// DeleteById
func (dao *SysRoleDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysRoleDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysRoleDao) Update(sysRole models.SysRole) error {
	return dao.db.Updates(&sysRole).Error
}

// GetList
func (dao *SysRoleDao) GetList() (list []models.SysRole) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysRoleDao) GetPage(page page.PageParam) (list []models.SysRole, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("name", "LIKE", page.FilterMap["name"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysRoleDao) GetById(id int) (row models.SysRole) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByIds
func (dao *SysRoleDao) GetByIds(ids []int) (row []models.SysRole) {
	dao.db.Model(&dao.models).Where("id IN ?", ids).Find(&row)
	return row
}
