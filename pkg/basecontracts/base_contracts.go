package basecontracts

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"update_at" gorm:"updated_at;autoUpdateTime"`
}

type PhoneNumber struct {
	CountryCode string `gorm:"column:country_code;not null" json:"country_code"`
	Number      string `gorm:"column:number;not null;unique" json:"number"`
}
