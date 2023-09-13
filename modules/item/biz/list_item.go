package biz

import (
	"context"
	"goobike-backend/common"
	"goobike-backend/modules/item/model"
)

// interface to storage
type ListItemStorage interface {
	ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, morekeys ...string) ([]model.TodoItem, error)
}

// business chir dungf interface ko biet storage dang dungf phuong phap gif de luu tru
type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
