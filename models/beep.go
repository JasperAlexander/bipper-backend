package models

import (
	"time"
)

type Beep struct {
	ID                string    `json:"id" bson:"_id,omitempty"`
	UserID            string    `json:"userid" bson:"userid,omitempty"`
	Text              string    `json:"text" bson:"text,omitempty"`
	ConversationID    string    `json:"conversationid" bson:"conversationid,omitempty"`
	ReferenceID       string    `json:"referenceid" bson:"referenceid,omitempty"`
	CreatedAt         time.Time `json:"createdat" bson:"createdat,omitempty"`
	Sensitive         bool      `json:"sensitive" bson:"sensitive,omitempty"`
	Config            string    `json:"config" bson:"config,omitempty"`
	Location          string    `json:"location" bson:"location,omitempty"`
	ImpressionCount   int32     `json:"impressioncount" bson:"impressioncount,omitempty"`
	RebeepCount       int32     `json:"rebeepcount" bson:"rebeepcount,omitempty"`
	QuoteCount        int32     `json:"quotecount" bson:"quotecount,omitempty"`
	LikeCount         int32     `json:"likecount" bson:"likecount,omitempty"`
	ReplyCount        int32     `json:"replycount" bson:"replycount,omitempty"`
	UrlClickCount     int32     `json:"urlclickcount" bson:"urlclickcount,omitempty"`
	ProfileClickCount int32     `json:"profileclickcount" bson:"profileclickcount,omitempty"`
	DetailsClickCount int32     `json:"detailsclickcount" bson:"detailsclickcount,omitempty"`
	VideoViewCount    int32     `json:"videoviewcount" bson:"videoviewcount,omitempty"`
}

func (Beep) IsNode() {}
