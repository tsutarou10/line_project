package entity

type RegisterEntity struct {
	ID   int64
	URL  string
	Memo string
}

type DeleteEntity struct {
	IDs []int64
}
