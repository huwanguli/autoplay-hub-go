package models

// ApiScript 用于返回脚本列表的结构体
type ApiScript struct {
	Id          int64  `json:"id,string" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Owner       string `json:"owner"`
	CreateAt    string `json:"create_at" db:"create_at"`
	UpdateAt    string `json:"update_at" db:"update_at"`
}

type Script struct {
	ID          int64  `json:"id,string" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	OwnerID     int64  `json:"owner_id,string" db:"owner_id"`
	Content     string `json:"content" db:"content"`
	CreateAt    string `json:"create_at" db:"created_at"`
	UpdateAt    string `json:"update_at" db:"updated_at"`
}

// ApiScriptDetail 用于返回脚本详情的结构体
type ApiScriptDetail struct {
	ID          int64  `json:"id,string" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Content     string `json:"content" db:"content"`
	CreateAt    string `json:"create_at" db:"created_at"`
	UpdateAt    string `json:"update_at" db:"updated_at"`
}
