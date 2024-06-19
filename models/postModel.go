package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model // update id uuid or guid
	Title      string
	Body       string
}
