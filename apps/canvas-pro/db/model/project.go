package model

import "github.com/cza14h/nino-work/pkg/db"

type ProjectGroup struct {
	db.BaseModel
	Name      string
	Workspace string
}

func (p ProjectGroup) TableName() string {
	return "project_prefixs"
}

type ProjectModel struct {
	db.BaseModel
	Name       string
	Code       string
	Version    string
	RootConfig string `gorm:"type:blob"`
}

func (p ProjectModel) TableName() string {
	return "projects"
}