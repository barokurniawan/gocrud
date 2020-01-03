package model

import (
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/barokurniawan/gocrud/entity"

	"github.com/barokurniawan/gocrud/sys"
)

type Guestbook struct {
	db *sys.Database
}

func (gb *Guestbook) SetDB(db *sys.Database) {
	gb.db = db
}

func (gb *Guestbook) Delete(ID int64) (bool, error) {
	conn := gb.db.Connection
	stmt, err := conn.Prepare("delete from guestbook where id = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (gb *Guestbook) CreateNew(Name string, Message string, created_at time.Time) (bool, error) {
	conn := gb.db.Connection
	stmt, err := conn.Prepare("INSERT INTO guestbook (Name, Message, created_at) VALUES (?, ?, ?)")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(Name, Message, created_at)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (gb *Guestbook) AdvanceShowList() ([]entity.Guestbook, error) {
	var ListGuestbook []entity.Guestbook
	var guestbook entity.Guestbook

	conn := gb.db.Connection
	rows, err := conn.Query("select * from guestbook order by created_at desc")
	if err != nil {
		return ListGuestbook, err
	}

	for rows.Next() {
		var ID int64
		var Name string
		var Message string
		var CreateAT mysql.NullTime

		rows.Scan(
			&ID,
			&Name,
			&Message,
			&CreateAT,
		)

		guestbook.ID = ID
		guestbook.Name = Name
		guestbook.Message = Message
		if CreateAT.Valid {
			guestbook.CreatedAT = CreateAT.Time
		}

		ListGuestbook = append(ListGuestbook, guestbook)
	}

	return ListGuestbook, nil
}
