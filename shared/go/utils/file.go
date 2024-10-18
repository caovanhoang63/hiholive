package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type File struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (File) TableName() string { return "images" }

// Scan scan value into Jsonb,
// decode jsonb in db into struct
// implements sql.Scanner interface
func (j *File) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value: %v", value))
	}
	var image File
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}
	*j = image
	return nil
}

// Value return json value;
// encode struct to []byte aka jsonb
// ;implement driver.Valuer interface
func (j *File) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Files []File

func (j *Files) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value: %v", value))
	}
	var images Files
	if err := json.Unmarshal(bytes, &images); err != nil {
		return err
	}
	*j = images
	return nil
}

func (j *Files) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
