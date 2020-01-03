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

func (i *InitService) Init() *InitService {
	DB := NewDatabase()
	i.DB = DB

	GB := NewGuestbook(DB)
	i.GB = GB

	return i
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
