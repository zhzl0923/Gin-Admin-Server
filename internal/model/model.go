package model

type Model struct {
	ID        uint64    `gorm:"primary_key" json:"id"` //自增ID
	CreatedAt LocalTime `json:"created_at"`            //创建时间
	UpdatedAt LocalTime `json:"updated_at"`            //修改时间
}
