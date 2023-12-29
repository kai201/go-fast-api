package accessor

import (
	"context"
	"errors"

	"github.com/fast/internal/models"
	"github.com/fast/pkg/mysql/query"
	"gorm.io/gorm"
)

type UserAccessor interface {
	Create(ctx context.Context, table *models.SysUser) error
	GetListByCondition(ctx context.Context, condition *query.Conditions) (*models.SysUser, error)
	GetList(ctx context.Context, params *query.Params) ([]*models.SysUser, int64, error)
}

type userAccessor struct {
	db *gorm.DB
}

// GetListByCondition implements UserAccessor.
func (d *userAccessor) GetListByCondition(ctx context.Context, condition *query.Conditions) (*models.SysUser, error) {
	queryStr, args, err := condition.ConvertToGorm()
	if err != nil {
		return nil, err
	}

	table := &models.SysUser{}
	err = d.db.WithContext(ctx).Where(queryStr, args...).First(table).Error
	if err != nil {
		return nil, err
	}

	return table, nil
}

func NewUserAccessor() UserAccessor {
	return &userAccessor{
		db: Get(),
	}
}

// Create implements UserAccessor.
func (d *userAccessor) Create(ctx context.Context, table *models.SysUser) error {
	err := d.db.WithContext(ctx).Create(table).Error
	return err
}

// GetList implements UserAccessor.
func (d *userAccessor) GetList(ctx context.Context, params *query.Params) ([]*models.SysUser, int64, error) {

	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&models.SysUser{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*models.SysUser{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}
