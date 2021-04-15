package CustodyAPI

import (
	"time"
)

type Valid interface {
	OK() error
}

type CustodyCase struct {
	ID          int
	Description string
	StartDate   time.Time
}

func (c *CustodyCase) OK() error {
	if c.Description == "" {
		//return error.New("Description required")
	}
}
func (c *CustodyCase) CustodyCaseAdd(d string) error {
	s := c.Description
	println(s)
}
