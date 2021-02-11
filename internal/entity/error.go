package entity

import (
	"time"

	"github.com/dwarukira/findcare/internal/event"
	"github.com/sirupsen/logrus"
)

// Error represents an error message log.
type Error struct {
	ID           uint      `gorm:"primary_key" json:"ID" yaml:"ID"`
	ErrorTime    time.Time `sql:"index" json:"Time" yaml:"Time"`
	ErrorLevel   string    `gorm:"type:VARBINARY(32)" json:"Level" yaml:"Level"`
	ErrorMessage string    `gorm:"type:VARBINARY(2048)" json:"Message" yaml:"Message"`
}

type Errors []Error

// SaveErrorMessages subscribes to error logs and stored them in the errors table.
func SaveErrorMessages() {
	s := event.Subscribe("log.*")

	defer func() {
		event.Unsubscribe(s)
	}()

	for msg := range s.Receiver {
		level, ok := msg.Fields["level"]

		if !ok {
			continue
		}

		logLevel, err := logrus.ParseLevel(level.(string))

		if err != nil || logLevel >= logrus.InfoLevel {
			continue
		}

		newError := Error{ErrorLevel: logLevel.String()}

		if val, ok := msg.Fields["message"]; ok {
			newError.ErrorMessage = val.(string)
		}

		if val, ok := msg.Fields["time"]; ok {
			newError.ErrorTime = val.(time.Time)
		}

		Db().Create(&newError)
	}
}
