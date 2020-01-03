package contract

import (
	"github.com/barokurniawan/gocrud/sys"
)

type Database interface {
	CreateConnection() *sys.Database
	GetCurrentConnection() *sys.Database
}
