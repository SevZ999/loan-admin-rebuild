package dao

import (
	"github.com/small-ek/antgo/frame/ant"
	"gorm.io/gorm"
	"server/app/entity/models"
)

type SysAdminUsersSysRoleDao struct {
	db     *gorm.DB
	models *models.SysAdminUsersSysRole
}

func NewSysAdminUsersSysRoleDao(db *gorm.DB) *SysAdminUsersSysRoleDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysAdminUsersSysRoleDao{db: db}
}

// DeleteByAdminUserId
func (dao *SysAdminUsersSysRoleDao) DeleteByAdminUserId(sys_admin_users_id int) error {
	return dao.db.Where("sys_admin_users_id=?", sys_admin_users_id).Delete(&dao.models).Error
}

// DeleteByAdminUserId
func (dao *SysAdminUsersSysRoleDao) DeleteByAdminUserIds(sys_admin_users_id []int) error {
	return dao.db.Where("sys_admin_users_id IN ?", sys_admin_users_id).Delete(&dao.models).Error
}
