package model

type Menu struct {
	ID       int `db:"id"`
	ParentID int `db:"parent_id"`
}
