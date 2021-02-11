package entity

import "time"

type Providers []Provider

type Provider struct {
	ID        string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false;" json:"ProviderID" yaml:"ProviderID"`
	Name      string    `gorm:"type:VARCHAR(255);" json:"Name" yaml:"Name"`
	CreatedAt time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt time.Time `json:"UpdatedAt" yaml:"-"`
}
