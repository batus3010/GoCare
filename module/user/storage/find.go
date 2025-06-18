package userStorage

import (
	"GoCare/common"
	userModel "GoCare/module/user/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}) (*userModel.User, error) {
	db := s.db.Table(userModel.User{}.TableName())
	var user userModel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorEntityNotFound(userModel.EntityUser, err)
		}
		return nil, common.ErrorDB(err)
	}
	return &user, nil
}
