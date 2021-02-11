package entity

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	Name  string `gorm:"unique";"index"`
	Icons string
}

func (s *Service) Create() error {
	return Db().Create(s).Error
}

func CreateDefaultServices() {
	serviceInpatient := Service{
		Name: "Inpatient",
	}

	if err := serviceInpatient.Create(); err != nil {
		log.Debug(err)
	}

	provider := &Provider{}

	if err := Db().First(provider); err != nil {
		log.Debug(err)
		return
	}

	provider.Services = append(provider.Services, serviceInpatient)

	if err := Db().Save(provider); err != nil {
		log.Debug(err)
		return
	}

	serviceDental := Service{
		Name: "Dental",
	}

	if err := serviceDental.Create(); err != nil {
		log.Debug(err)
	}

	serviceMaternity := Service{
		Name: "Maternity",
	}

	if err := serviceMaternity.Create(); err != nil {
		log.Debug(err)
	}

	servicePhysiotherapy := Service{
		Name: "Physiotherapy",
	}

	if err := servicePhysiotherapy.Create(); err != nil {
		log.Debug(err)
	}
}
