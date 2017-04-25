package dao

import (
	"github.com/songrgg/backeye/model"
	"github.com/songrgg/backeye/model/form"
	"github.com/songrgg/backeye/std"
	modelmapper "gopkg.in/jeevatkm/go-model.v0"

	"time"
)

// NewProject creates a new project
func NewProject(f *form.Project) error {
	nowTime := time.Now()
	proj := model.Project{
		TimeMixin: std.TimeMixin{
			CreatedAt: nowTime,
			UpdatedAt: nowTime,
		},
	}
	errors := modelmapper.Copy(&proj, f)
	if len(errors) > 0 {
		return errors[0]
	}

	// save task
	model.DB().Save(&proj)
	if err := model.DB().Error; err != nil {
		std.LogErrorc("mysql", err, "failed to create project")
		return err
	}
	return nil
}

// GetProjects fetches projects
func GetProjects() ([]model.Project, error) {
	// save task
	p := make([]model.Project, 0)
	model.DB().Find(&p)
	return p, nil
}

// GetProject fetches the specified project
func GetProject(id int64) (model.Project, error) {
	// save task
	var p model.Project
	model.DB().Find(&p, map[string]interface{}{"id": id})
	return p, nil
}

// RemoveProject removes the specified project
func RemoveProject(id int64) error {
	var proj model.Project
	model.DB().Find(&proj, map[string]interface{}{"id": id}).Updates(map[string]interface{}{
		"updatedAt": time.Now(),
		"status":    "deleted",
	})
	return model.DB().Error
}

// UpdateProject updates the specified project
func UpdateProject(id int64, f *form.Project) error {
	var proj model.Project
	errors := modelmapper.Copy(&proj, f)
	if len(errors) > 0 {
		return errors[0]
	}

	model.DB().Find(&proj, map[string]interface{}{"id": id}).Updates(proj)
	return model.DB().Error
}
