package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db        *sql.DB
	dbParams  string
	closeChan chan struct{}
	stmts     map[string]*sql.Stmt
	gStmtInfo map[string]string
}

type IPrepare interface {
	Prepare(db *sql.DB) (map[string]*sql.Stmt, error)
}

func New(dbParams string, gStmtInfo map[string]string) (*Database, error) {
	db, e := connectDB(dbParams)
	if e != nil {
		return nil, e
	}

	myDB := &Database{
		db:        db,
		dbParams:  dbParams,
		closeChan: make(chan struct{}),
		gStmtInfo: gStmtInfo,
	}

	myDB.Prepare()

	return myDB, nil
}

func connectDB(dbParams string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbParams)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func (mdb *Database) Prepare() {
	mdb.stmts = make(map[string]*sql.Stmt)
	for name, sql := range mdb.gStmtInfo {
		println(name, sql)
		stmt, e := mdb.db.Prepare(sql)
		if e != nil {
			log.Fatal("prepare error ", e)
			mdb.stmts = nil
		}
		mdb.stmts[name] = stmt
	}
}

func (mdb *Database) Query(stmtName string, args ...interface{}) (*sql.Rows, error) {
	if stmt, ok := mdb.stmts[stmtName]; ok {
		return stmt.Query(args...)
	}
	return nil, fmt.Errorf("now stmt %s", stmtName)
}

func (mdb *Database) QueryRow(stmtName string, args ...interface{}) *sql.Row {
	if stmt, ok := mdb.stmts[stmtName]; ok {
		return stmt.QueryRow(args...)
	}
	return nil
}

func (mdb *Database) Exec(stmtName string, args ...interface{}) (sql.Result, error) {
	if stmt, ok := mdb.stmts[stmtName]; ok {
		return stmt.Exec(args...)
	}
	return nil, fmt.Errorf("not stmt %s", stmtName)
}

func (mdb *Database) Close() {
	select {
	case mdb.closeChan <- struct{}{}:
	}
	mdb.db.Close()
}
