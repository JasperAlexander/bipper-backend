package models

import (
	"time"
)

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Name     string `json:"name" bson:"name,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	// Password    *string   `json:"password" bson:"password,omitempty"`
	CreatedAt   time.Time `json:"createdat" bson:"createdat,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	ProfileImg  string    `json:"profileimg" bson:"profileimg,omitempty"`
	BannerImg   string    `json:"bannerimg" bson:"bannerimg,omitempty"`
	Dob         int       `json:"dob" bson:"dob,omitempty"`
	// PhoneNr       *string   `json:"phonenr" bson:"phonenr,omitempty"`
	Location     string   `json:"location" bson:"location,omitempty"`
	Url          string   `json:"url" bson:"url,omitempty"`
	PinnedBeepID string   `json:"pinnedbeepid" bson:"pinnedbeepid,omitempty"`
	Protected    bool     `json:"protected" bson:"protected,omitempty"`
	Config       string   `json:"config" bson:"config,omitempty"`
	Beeps        []string `json:"beeps" bson:"beeps,omitempty"`
}

func (User) IsNode() {}
