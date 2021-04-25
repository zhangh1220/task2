package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

const dsn = "{user}:{password}@tcp(127.0.0.1:3306)/{database}?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8"


//创建db
func NewDB()( *sql.DB, error) {
	d,err :=sql.Open("mysql",dsn)
	if err != nil {
		err = errors.WithStack(err)
		return nil,err
	}
	d.SetMaxOpenConns(20)
	d.SetMaxIdleConns(10)
	d.SetConnMaxLifetime(time.Hour*4)
	return d, nil
}