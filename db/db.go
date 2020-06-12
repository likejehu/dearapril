package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/likejehu/dearapril/models"

	"github.com/pkg/errors"

	"github.com/eaigner/jet"
)

// Error404 is 404 err for store, when key is not found
var Error404 = errors.New("key not found")

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "minedemo"
)

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection and postgre cloud
type dbStore struct {
	db    *sql.DB
	cloud *jet.Db
}

var pgURL = "postgres://minzlhay:xnHd-hp3uJljOBWnLDwQkb3zt53phs55@balarama.db.elephantsql.com:5432/minzlhay"
var cloudDB *jet.Db

func connectionString() string {
	psqlstr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlstr
}

var connString = connectionString()
var psqlDB *sql.DB

//Store is for storing
var Store = dbStore{
	cloud: cloudDB,
}

func connectDB() (err error) {
	psqlDB, err = sql.Open("postgres", connString)
	return err
}
func connectCloud() (err error) {
	cloudDB, err = jet.Open("postgres", pgURL)
	return err
}

func (store *dbStore) Set(p *models.Project) (err error) {
	_ = store.cloud.Query("INSERT INTO links VALUES ($1)", p)
	return err
}
func (store *dbStore) Get(key string) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.cloud.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
