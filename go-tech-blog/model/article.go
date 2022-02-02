package model

// Article ...
type Article struct {
	ID      int    `db:"id"`
	Article string `db:"article"`
}
