package biz

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/user/model"
)

type GetUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

// business chir dungf interface ko biet storage dang dungf phuong phap gif de luu tru
type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*model.User, error) {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}

func (biz *getUserBiz) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
