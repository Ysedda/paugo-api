package models

type TestInput struct {
	Name string `json:"name"`
}

type TestRow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
