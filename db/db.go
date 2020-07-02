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

func (store *dbStore) CreateProject(project *models.Project) (err error) {
	_, err = store.db.Exec(`INSERT INTO projects (name, description) VALUES ($1, $2);`, project.Name, project.Description)
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
func (store *dbStore) GetAllProjects() (projects []*models.Project, err error) {

	projects = []*models.Project{}
	rows, err := store.db.Query(`SELECT id, name, description  FROM projects;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		project := new(models.Project)
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

	return projects, err

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

func (store *dbStore) CreateColumn(column *models.Column) (err error) {
	_, err = store.db.Exec(`INSERT INTO columns (name, position) VALUES ($1, $2);`, column.Name, column.Position)
	return err

}
func (store *dbStore) GetColumn(id int) (column *models.Column, err error) {
	column = new(models.Column)

	err = store.db.QueryRow(`SELECT id, name, position FROM columns WHERE id=$1;`, id).Scan(&column)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return column, err

}
func (store *dbStore) GetAllColumns() (columns []*models.Column, err error) {

	columns = []*models.Column{}
	rows, err := store.db.Query(`SELECT id, name, position  FROM columns;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		column := new(models.Column)
		err = rows.Scan(&column.ID, &column.Name, &column.Position)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, column)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return columns, err

}
func (store *dbStore) UpdateColumn(id int, column *models.Column) (err error) {

	_, err = store.db.Exec(`UPDATE columns SET name = $2 position =$3 WHERE id = $1;`, id, column.Name, column.Position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}
func (store *dbStore) DeleteColumn(id int) (err error) {
	_, err = store.db.Exec(`DELETE FROM columns WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//Tasks

func (store *dbStore) CreateTask(task *models.Task) (err error) {
	_, err = store.db.Exec(`INSERT INTO tasks (name, description, position) VALUES ($1, $2, $3);`, task.Name, task.Description, task.Position)
	return err

}
func (store *dbStore) GetTask(id int) (task *models.Task, err error) {
	task = new(models.Task)

	err = store.db.QueryRow(`SELECT id, name, description, position FROM tasks WHERE id=$1;`, id).Scan(&task)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return task, err

}
func (store *dbStore) GetAllTasks() (tasks []*models.Task, err error) {

	tasks = []*models.Task{}
	rows, err := store.db.Query(`SELECT id,  name, description, position  FROM tasks;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		task := new(models.Task)
		err = rows.Scan(&task.ID, &task.Name, task.Description, &task.Position)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tasks, err

}
func (store *dbStore) UpdateTask(id int, task *models.Task) (err error) {

	_, err = store.db.Exec(`UPDATE tasks SET name = $2 description = $3 position =$4 WHERE id = $1;`, id, task.Name, task.Description, task.Position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}
func (store *dbStore) DeleteTask(id int) (err error) {
	_, err = store.db.Exec(`DELETE FROM tasks WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//Comments

func (store *dbStore) CreateComment(comment *models.Comment) (err error) {
	_, err = store.db.Exec(`INSERT INTO comments (name, text, date) VALUES ($1, $2, $3);`, comment.Name, comment.Text, comment.Date)
	return err

}
func (store *dbStore) GetComment(id int) (comment *models.Comment, err error) {
	comment = new(models.Comment)

	err = store.db.QueryRow(`SELECT id, name, text, date FROM comments WHERE id=$1;`, id).Scan(&comment)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return comment, err

}
func (store *dbStore) GetAllComments() (comments []*models.Comment, err error) {

	comments = []*models.Comment{}
	rows, err := store.db.Query(`SELECT id,  name, description, position  FROM comments;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		comment := new(models.Comment)
		err = rows.Scan(&comment.ID, &comment.Name, &comment.Text, &comment.Date)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return comments, err

}
func (store *dbStore) UpdateComment(id int, comment *models.Comment) (err error) {

	_, err = store.db.Exec(`UPDATE comments SET name = $2 description = $3 position =$4 WHERE id = $1;`, id, comment.Name, comment.Text, comment.Date)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}
func (store *dbStore) DeleteComment(id int) (err error) {
	_, err = store.db.Exec(`DELETE FROM comments WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
