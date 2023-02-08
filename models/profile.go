package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Profile struct {
	Model

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`

	State int `json:"state"`
}

// GetTags gets a list of tags based on paging and constraints
func GetProfiles(pageNum int, pageSize int, maps interface{}) ([]Profile, error) {
	var (
		profiles []Profile
		err      error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&profiles).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&profiles).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return profiles, nil
}

func GetProfileTotal(maps interface{}) (count int) {
	db.Model(&Profile{}).Where(maps).Count(&count)

	return
}

func ExistProfileByName(name string) bool {
	var profile Profile
	db.Select("id").Where("name = ?", name).First(&profile)
	if profile.ID > 0 {
		return true
	}

	return false
}

func AddProfile(title string, desc string, content string, cover_image_url string, creater_id string, create_by string) bool {
	db.Create(&Profile{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: cover_image_url,
		CreatedBy:     creater_id,
	})

	return true
}

func ModifyProfile(title string, desc string, content string, cover_image_url string, creater_id string, create_by string) bool {
	db.Save(&Profile{
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: cover_image_url,
		CreatedBy:     creater_id,
	})

	return true
}

func (profile *Profile) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (profile *Profile) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistProfileByID(id int) bool {
	var profile Profile
	db.Select("id").Where("id = ?", id).First(&profile)
	if profile.ID > 0 {
		return true
	}

	return false
}

func DeleteProfile(id int) bool {
	db.Where("id = ?", id).Delete(&Profile{})

	return true
}

func EditProfile(id int, data interface{}) bool {
	db.Model(&Profile{}).Where("id = ?", id).Updates(data)

	return true
}

func CleanAllProfile() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Profile{})

	return true
}
