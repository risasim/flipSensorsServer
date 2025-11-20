package db

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var db *sql.DB

type DataEntry struct {
	Id          int
	TimeStamp   timestamppb.Timestamp
	Gas         int
	Temperature int
}
