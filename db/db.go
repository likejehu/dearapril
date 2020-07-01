package db

import (
	"database/sql"
	"log"

	"github.com/likejehu/dearapril/models"

	"github.com/pkg/errors"
)

// Error404 is 404 err for store, when key is not found
var Error404 = errors.New("key not found")

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection and postgre cloud
type dbStore struct {
	db *sql.DB
}

var pgURL = "postgres://minzlhay:xnHd-hp3uJljOBWnLDwQkb3zt53phs55@balarama.db.elephantsql.com:5432/minzlhay"

var psqlDB *sql.DB

//Store is for storing
var Store = dbStore{
	db: psqlDB,
}

func connectDB() (err error) {
	psqlDB, err = sql.Open("postgres", pgURL)
	return err
}

//Projects

func (store *dbStore) CreateProject(p *models.Project) (err error) {
	_, err = store.db.Exec(`INSERT INTO projects (name, description) VALUES ($1, $2);`, p.Name, p.Description)
	return err
}
func (store *dbStore) GetProject(id int) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.db.QueryRow(`SELECT id, name, description FROM projects WHERE id=$1;`, id).Scan(&project.ID, &project.Name, &project.Description)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) GetAllProjects() (project *models.Project, err error) {

	projects := []*models.Project{}
	rows, err := store.db.Query(`SELECT id, name, description  FROM projects;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		project = new(models.Project)
		err = rows.Scan(&project.ID, &project.Name)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return project, err

}
func (store *dbStore) UpdateProject(id int, project *models.Project) (err error) {

	_, err = store.db.Exec(`UPDATE projects SET name = $2 description =$3 WHERE id = $1;`, id, project.Name, project.Description)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}
func (store *dbStore) DeleteProject(id int) (err error) {
	_, err = store.db.Exec(`DELETE FROM projects WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//Columns

func (store *dbStore) CreateColumn(p *models.Project) (err error) {
	_, err = store.db.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}
func (store *dbStore) GetColumn(key string) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.db.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) GetAllColumns() (project *models.Project, err error) {
	project = new(models.Project)
	key := "1"
	err = store.db.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) UpdateColumn(key string) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.db.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) DeleteColumn(p *models.Project) (err error) {
	_ = store.db.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}

//Tasks

func (store *dbStore) CreateTask(p *models.Project) (err error) {
	_ = store.db.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}
func (store *dbStore) GetTask(key string) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.db.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) GetAllTasks() (project *models.Project, err error) {
	project = new(models.Project)
	key := "1"
	err = store.cloud.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) UpdateTask(key string) (project *models.Project, err error) {
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
func (store *dbStore) DeleteTask(p *models.Project) (err error) {
	_ = store.cloud.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}

//Comments

func (store *dbStore) CreateComment(p *models.Project) (err error) {
	_ = store.cloud.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}
func (store *dbStore) GetComment(key string) (project *models.Project, err error) {
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
func (store *dbStore) GetAllComments() (project *models.Project, err error) {
	project = new(models.Project)
	key := "1"
	err = store.cloud.QueryRow(key).Scan(&project)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}
func (store *dbStore) UpdateComment(key string) (project *models.Project, err error) {
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
func (store *dbStore) DeleteComment(p *models.Project) (err error) {
	_ = store.cloud.Query("INSERT INTO projects VALUES ($1)", p)
	return err
}
