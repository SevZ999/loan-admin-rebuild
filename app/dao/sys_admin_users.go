package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/app/entity/models"
)

type SysAdminUsersDao struct {
	db     *gorm.DB
	models *models.SysAdminUsers
}

func NewSysAdminUsersDao(db *gorm.DB) *SysAdminUsersDao {
	if db == nil {
		db = ant.Db()
	}
	return &SysAdminUsersDao{db: db}
}

// Create
func (dao *SysAdminUsersDao) Create(sysAdminUsers models.SysAdminUsers) error {
	return dao.db.Create(&sysAdminUsers).Error
}

// DeleteById
func (dao *SysAdminUsersDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *SysAdminUsersDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *SysAdminUsersDao) Update(sysAdminUsers models.SysAdminUsers) error {
	return dao.db.Updates(&sysAdminUsers).Error
}

// Update
func (dao *SysAdminUsersDao) UpdateAndRole(sysAdminUsers models.SysAdminUsers) error {
	return dao.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&sysAdminUsers).Error
}

// GetList
func (dao *SysAdminUsersDao) GetList() (list []models.SysAdminUsers) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *SysAdminUsersDao) GetPage(page page.PageParam) (list []models.SysAdminUsers, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("username", "LIKE", page.FilterMap["username"]),
		asql.Where("nick_name", "LIKE", page.FilterMap["nick_name"]),
		asql.Where("phone", "LIKE", page.FilterMap["phone"]),
		asql.Where("email", "LIKE", page.FilterMap["email"]),
		asql.Where("status", "=", page.FilterMap["status"]),
		asql.Where("parent_id", "=", page.FilterMap["parent_id"]),
		asql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Omit("password").Preload(clause.Associations).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *SysAdminUsersDao) GetById(id int) (row models.SysAdminUsers) {
	dao.db.Model(&dao.models).Omit("password").Preload(clause.Associations).Where("id=?", id).Limit(1).Find(&row)
	return row
}

// GetByUserNameAndStatus
func (dao *SysAdminUsersDao) GetByUserNameAndStatus(username string) (row models.SysAdminUsers) {
	dao.db.Model(&dao.models).Where("username=? AND status=2", username).Preload(clause.Associations).Limit(1).Find(&row)
	return row
}

// GetByUserNameAndStatus
func (dao *SysAdminUsersDao) GetByUserName(username string) (row models.SysAdminUsers) {
	dao.db.Model(&dao.models).Where("username=?", username).Limit(1).Find(&row)
	return row
}
