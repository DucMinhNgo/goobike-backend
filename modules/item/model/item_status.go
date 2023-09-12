package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// Enum
type ItemStatus int

const (
	ItamStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemstatuses = [3]string{"doing", "done", "deleted"}

func (item *ItemStatus) String() string {
	return allItemstatuses[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemstatuses {
		if allItemstatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status")
}

// Doc du lieu mang len Itemstatus
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*item = v

	return nil
}

// Itemstatus veef duwx lieeuj
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// Ho tro Itemstatus thanh JsonValue
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// JsonValue sang Items status
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
