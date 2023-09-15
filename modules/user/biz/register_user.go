package biz

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/user/model"
	"strings"
)

type ResisterUserStorage interface {
	RegisterUser(ctx context.Context, data *model.UserCreation) error
}

type registerUserBiz struct {
	store ResisterUserStorage
}

func NewRegisterUserBiz(store ResisterUserStorage) *registerUserBiz {
	return &registerUserBiz{store: store}
}

func (biz *registerUserBiz) RegisterNewUser(ctx context.Context, data *model.UserCreation) error {
	email := strings.TrimSpace(data.Email)

	if email == "" {
		return model.ErrEmailIsBlank
	}

	name := strings.TrimSpace(data.Name)

	if name == "" {
		return model.ErrNameIsBlank
	}

	username := strings.TrimSpace(data.Username)

	if username == "" {
		return model.ErrUsernameIsBlank
	}

	password := strings.TrimSpace(data.Password)

	if password == "" {
		return model.ErrPasswordIsBlank
	}

	if err := biz.store.RegisterUser(ctx, data); err != nil {
		return common.CannotCreateEntity(model.EntityName, err)
	}

	return nil
}
