package gateway

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pay interface {
	Pay(taxid string, value float64) string
}

type Pay1 struct {
	ID byte
}
type Pay2 struct {
	ID byte
}

func getUUID() string {
	return uuid.NewString()
}

func (p *Pay1) Pay(taxid string, value float64) string {
	time.Sleep(time.Second * 4)
	return fmt.Sprintln(getUUID(), 1)
}

func (p *Pay2) Pay(taxid string, value float64) string {
	time.Sleep(time.Second * 3)
	return fmt.Sprintln(getUUID(), 2)
}
