package models

type Mysql struct {
	userName     string
	password     string
	host         string
	port         int
	dbName       string
	maxOpenConns int
	maxIdleConns int
}
