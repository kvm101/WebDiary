package pgdb

type Tasks struct {
	Task_ID     int64  `json:"task_id"`
	Date        string `json:"date"`
	User_id     int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Due_date    string `json:"due_date"`
}

func (t Tasks) Get() []byte {
	return nil
}

func (t Tasks) Add() string {
	return ""
}

func (t Tasks) Delete() string {
	return ""
}

func (t Tasks) Update() string {
	return ""
}
