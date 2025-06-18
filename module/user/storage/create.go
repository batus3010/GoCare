package userStorage

import (
	"GoCare/common"
	userModel "GoCare/module/user/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *userModel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrorDB(err)
	}

	return nil
}
