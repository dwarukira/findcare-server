package entity

import (
	"github.com/jinzhu/gorm"
)

type Providers []Provider

type Provider struct {
	gorm.Model
	Name        string    `gorm:"type:VARCHAR(255);unique" json:"Name" yaml:"Name"`
	Description string    `gorm:"type:text;" json:"Description" yaml:"Description"`
	ProviderLat float32   `gorm:"type:FLOAT;index;" json:"Lat" yaml:"Lat,omitempty"`
	ProviderLng float32   `gorm:"type:FLOAT;index;" json:"Lng" yaml:"Lng,omitempty"`
	Website     string    `gorm:"type:VARCHAR(255);" json:"Website" yaml:"Website"`
	Phone       string    `gorm:"type:VARCHAR(255);" json:"Phone" yaml:"Phone"`
	Photo       string    `gorm:"type:VARCHAR(255);" json:"Photo" yaml:"Photo"`
	Services    []Service `gorm:"many2many:provider_services;"`
}

func (p *Provider) Create() error {
	return Db().Create(p).Error
}

func CreateDefaultProviders() {
	provider := Provider{
		Name:        "Mater kenya",
		Description: "Mater Misericordiae Hospital was opened in 1962 by the Sisters of Mercy, a Catholic Order of Nuns originating from Ireland, three years after registering themselves as the Registered Trustees of an entity under the Perpetual Succession Act (the succeeding legislation after independence).",
	}

	if err := provider.Create(); err == nil {
		return
	}

}
