// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"gorm.io/gorm"
)

const TableNameImage = "ims_image"

// Image mapped from table <ims_image>
type Image struct {
	ID              int32          `gorm:"column:id;type:INTEGER" json:"id"`
	Tag             string         `gorm:"column:tag" json:"tag"`
	BuildGit        string         `gorm:"column:build_git" json:"buildGit"`
	BuildDockerfile string         `gorm:"column:build_dockerfile" json:"buildDockerfile"`
	BuildZip        string         `gorm:"column:build_zip" json:"buildZip"`
	Status          int32          `gorm:"column:status" json:"status"`
	StatusStep      string         `gorm:"column:status_step" json:"statusStep"`
	Message         string         `gorm:"column:message" json:"message"`
	CreatedAt       int32          `gorm:"column:created_at" json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
	Md5             string         `gorm:"column:md5" json:"md5"`
}

// TableName Image's table name
func (*Image) TableName() string {
	return TableNameImage
}
