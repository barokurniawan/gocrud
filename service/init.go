package service

import (
	"github.com/barokurniawan/gocrud/contract"
	"github.com/barokurniawan/gocrud/model"
	"github.com/barokurniawan/gocrud/sys"
)

type InitService struct {
	DB *sys.Database
	GB *model.Guestbook
}

func NewInitService(DB *sys.Database, GB *model.Guestbook) *InitService {
	return &InitService{
		DB: DB,
		GB: GB,
	}
}

func NewDatabase() *sys.Database {
	db := sys.Database{}
	return db.CreateConnection()
}

func NewGuestbook(db contract.Database) *model.Guestbook {
	model := model.Guestbook{}
	model.SetDB(db.GetCurrentConnection())

	return &model
}
