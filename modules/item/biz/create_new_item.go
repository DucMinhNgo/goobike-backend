package biz

import (
	"context"
	"goobike-backend/modules/item/model"
	"strings"
)

// interface to storage
type CreateItemItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

// business chir dungf interface ko biet storage dang dungf phuong phap gif de luu tru
type createItemBiz struct {
	store CreateItemItemStorage
}

func NewCreateItemBiz(store CreateItemItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
