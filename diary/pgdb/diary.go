package pgdb

type Diary struct {
	Diary_ID int64  `json:"diary_id"`
	Date     string `json:"date"`
	User_id  int64  `json:"user_id"`
	Content  string `json:"content"`
}

func (d Diary) Get() []byte {
	return nil
}

func (d Diary) Add() string {
	return ""
}

func (d Diary) Delete() string {
	return ""
}

func (d Diary) Update() string {
	return ""
}
