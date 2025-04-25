package dao

import (
	"github.com/small-ek/antgo/db/adb/asql"
	"github.com/small-ek/antgo/frame/ant"
	"github.com/small-ek/antgo/utils/page"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/app/entity/models"
)

type FeedbackDao struct {
	db     *gorm.DB
	models *models.Feedback
}

func NewFeedbackDao(db *gorm.DB) *FeedbackDao {
	if db == nil {
		db = ant.Db()
	}
	return &FeedbackDao{db: db}
}

// Create
func (dao *FeedbackDao) Create(feedback models.Feedback) error {
	return dao.db.Create(&feedback).Error
}

// DeleteById
func (dao *FeedbackDao) DeleteById(id int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// DeleteByIds
func (dao *FeedbackDao) DeleteByIds(id []int) error {
	return dao.db.Delete(&dao.models, id).Error
}

// Update
func (dao *FeedbackDao) Update(feedback models.Feedback) error {
	return dao.db.Updates(&feedback).Error
}

// GetList
func (dao *FeedbackDao) GetList() (list []models.Feedback) {
	dao.db.Model(&dao.models).Find(&list)
	return list
}

// GetPage
func (dao *FeedbackDao) GetPage(page page.PageParam) (list []models.Feedback, total int64, err error) {
	err = dao.db.Model(&dao.models).Scopes(
		asql.Where("user_id", "IN", page.FilterMap["user_id"]),
		asql.Where("description", "LIKE", page.FilterMap["description"]),
		asql.Where("created_at", "BETWEEN", page.FilterMap["created_at"]),
		asql.Filters(page.Filter),
		asql.Order(page.Order, page.Desc),
		asql.Paginate(page.PageSize, page.CurrentPage),
	).Preload(clause.Associations).Find(&list).Offset(-1).Limit(1).Count(&total).Error
	return list, total, err
}

// GetById
func (dao *FeedbackDao) GetById(id int) (row models.Feedback) {
	dao.db.Model(&dao.models).Where("id=?", id).Limit(1).Find(&row)
	return row
}
