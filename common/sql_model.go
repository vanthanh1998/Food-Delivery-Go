package common

import "time"

type SqlModel struct {
	Id        int        `json:"_" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"column:id;"` // convert id to string
	Status    int        `json:"status" gorm:"column:status; default:1;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (m *SqlModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
