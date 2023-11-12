package model_dto

import (
	"time"
)

type UserAndEmail struct {
	Subject   string    `gorm:"column:subject;not null" json:"subject"`
	Body      string    `gorm:"column:body;not null" json:"body"`
	SentAt    time.Time `gorm:"column:sent_at;default:pg_systimestamp()" json:"sent_at"`
	CcList    string    `gorm:"column:cc_list" json:"cc_list"`
	BccList   string    `gorm:"column:bcc_list" json:"bcc_list"`
	IsDraft   bool      `gorm:"column:is_draft" json:"is_draft"`
	IsDeleted bool      `gorm:"column:is_deleted" json:"is_deleted"`
	SenderID  int32     `gorm:"column:sender_id" json:"sender_id"`
	Img       string    `gorm:"column:img" json:"img"`
	
	EmailID     int32     `gorm:"column:email_id" json:"email_id"`
	RecipientID int32     `gorm:"column:recipient_id" json:"recipient_id"`
	IsRead      bool      `gorm:"column:is_read" json:"is_read"`
	ReadAt      time.Time `gorm:"column:read_at" json:"read_at"`
	Folder      int32     `gorm:"column:folder;not null;default:1;comment:1：收件箱，2：发件箱，3：草稿，4：垃圾箱" json:"folder"` // 1：收件箱，2：发件箱，3：草稿，4：垃圾箱
}
