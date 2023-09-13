package biz

import (
	"context"
	"goobike-backend/modules/item/model"
)

// interface to storage
type GetItemItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

// business chir dungf interface ko biet storage dang dungf phuong phap gif de luu tru
type getItemBiz struct {
	store GetItemItemStorage
}

func NewGetItemBiz(store GetItemItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}