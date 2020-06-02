package incident

import (
	"github.com/jinzhu/gorm"
)

type Opts func(*Incident)

type Incident struct {
	gorm.Model
	Email  string `json:"email,omitempty"`
	Report string `json:"report,omitempty"`
}

func WithEmail(email string) Opts {
	return func(i *Incident) {
		i.Email = email
	}
}

func WithReport(report string) Opts {
	return func(i *Incident) {
		i.Report = report
	}
}

func New(optionFunc ...Opts) *Incident {
	inc := new(Incident)
	for _, fn := range optionFunc {
		fn(inc)
	}
	return inc
}
