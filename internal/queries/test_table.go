package test_table_queries

type Queries struct {
	InsertOne string
	SelectAll string
}

var InsertOne = "INSERT INTO test_table (name) VALUES ($1) RETURNING id"
var SelectAll = "SELECT id, name FROM test_table"

func TestQueries() Queries {
	return Queries{
		InsertOne: InsertOne,
		SelectAll: SelectAll,
	}
}
