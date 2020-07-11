package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/likejehu/dearapril/models"
)

//DBStore is the stuct for database operations
type DBStore struct {
	DB *sql.DB
}

var pgURL = "postgres://minzlhay:xnHd-hp3uJljOBWnLDwQkb3zt53phs55@balarama.db.elephantsql.com:5432/minzlhay"

// PostgreStore is instance of redis storage
var PostgreStore = NewPostgreConn(pgURL)

// NewPostgreConn returns a new DBStore Instance
func NewPostgreConn(address string) *DBStore {
	var postgreDB = &DBStore{}
	// Initialize the postgre connection to a redis instance running on your local machine
	conn, err := sql.Open("postgres", address)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	postgreDB.DB = conn
	log.Println("connetion established")
	return postgreDB
}

//Projects

//CreateProject is for inserting project to database
func (store *DBStore) CreateProject(project *models.Project) (id int, err error) {

	res, err := store.DB.Exec(`INSERT INTO projects (name, description) VALUES ($1, $2 );`, project.Name, project.Description)
	if err != nil {
		log.Fatal(err)
	}
	lastid, err := res.LastInsertId()
	id = int(lastid)
	return id, err
}

//GetProject is for getting project from database
func (store *DBStore) GetProject(id int) (project *models.Project, err error) {
	project = new(models.Project)

	err = store.DB.QueryRow(`SELECT id, name, description FROM projects WHERE id=$1;`, id).Scan(&project.ID, &project.Name, &project.Description)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return project, err

}

//GetAllProjects is for getting projects from database
func (store *DBStore) GetAllProjects() (projects []*models.Project, err error) {

	projects = []*models.Project{}
	rows, err := store.DB.Query(`SELECT id, name, description  FROM projects ORDER BY name;`)
	log.Println(rows)
	if err == sql.ErrNoRows {
		log.Println(err)
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		project := new(models.Project)
		err = rows.Scan(&project.ID, &project.Name, &project.Description)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
		projects = append(projects, project)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	return projects, err

}

//UpdateProject is for  updating project in database
func (store *DBStore) UpdateProject(id int, project *models.Project) (err error) {

	_, err = store.DB.Exec(`UPDATE projects SET name = $2 description =$3 WHERE id = $1;`, id, project.Name, project.Description)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//DeleteProject is for  deleting project in database
func (store *DBStore) DeleteProject(id int) (err error) {
	_, err = store.DB.Exec(`DELETE FROM projects WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//Columns

//CreateColumn is for inserting column into database
func (store *DBStore) CreateColumn(column *models.Column) (id int, err error) {
	res, err := store.DB.Exec(`INSERT INTO columns (name, position) VALUES ($1, $2);`, column.Name, column.Position)
	if err != nil {
		log.Fatal(err)
	}
	lastid, err := res.LastInsertId()
	id = int(lastid)
	return id, err

}

//GetColumn is for getting column from database
func (store *DBStore) GetColumn(id int) (column *models.Column, err error) {
	column = new(models.Column)

	err = store.DB.QueryRow(`SELECT id, name, position FROM columns WHERE id=$1;`, id).Scan(&column)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return column, err

}

//GetAllColumns is for getting all columns from database
func (store *DBStore) GetAllColumns() (columns []*models.Column, err error) {

	columns = []*models.Column{}
	rows, err := store.DB.Query(`SELECT id, name, position  FROM columns ORDER BY position;`)
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

//UpdateColumn is for updating column in database
func (store *DBStore) UpdateColumn(id int, column *models.Column) (err error) {

	_, err = store.DB.Exec(`UPDATE columns SET name = $2 position =$3 WHERE id = $1;`, id, column.Name, column.Position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//DeleteColumn is for deleting column in database
func (store *DBStore) DeleteColumn(id int) (err error) {
	_, err = store.DB.Exec(`DELETE FROM columns WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateColumnPosition is for updating column position in database
func (store *DBStore) UpdateColumnPosition(id int, position int) (err error) {

	_, err = store.DB.Exec(`UPDATE columns SET position =$2 WHERE id = $1;`, id, position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//CountColumns is for counting columns in database
func (store *DBStore) CountColumns(id int, position int) (err error) {

	_, err = store.DB.Exec(`UPDATE columns SET position =$2 WHERE id = $1;`, id, position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//Tasks

//CreateTask is for inserting task to database
func (store *DBStore) CreateTask(task *models.Task) (id int, err error) {
	res, err := store.DB.Exec(`INSERT INTO tasks (name, description, position) VALUES ($1, $2, $3);`, task.Name, task.Description, task.Position)
	if err != nil {
		log.Fatal(err)
	}
	lastid, err := res.LastInsertId()
	id = int(lastid)
	return id, err

}

//GetTask is for getting task from database
func (store *DBStore) GetTask(id int) (task *models.Task, err error) {
	task = new(models.Task)

	err = store.DB.QueryRow(`SELECT id, name, description, position FROM tasks WHERE id=$1;`, id).Scan(&task)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return task, err

}

//GetAllTasks is for all getting tasks from database
func (store *DBStore) GetAllTasks() (tasks []*models.Task, err error) {

	tasks = []*models.Task{}
	rows, err := store.DB.Query(`SELECT id,  name, description, position  FROM tasks ORDER BY position;`)
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

//UpdateTask is for updating task in  database
func (store *DBStore) UpdateTask(id int, task *models.Task) (err error) {

	_, err = store.DB.Exec(`UPDATE tasks SET name = $2 description = $3 position =$4 WHERE id = $1;`, id, task.Name, task.Description, task.Position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//DeleteTask is for deleting task in  database
func (store *DBStore) DeleteTask(id int) (err error) {
	_, err = store.DB.Exec(`DELETE FROM tasks WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateTaskPosition is for updating  task position in  database
func (store *DBStore) UpdateTaskPosition(id int, position int) (err error) {

	_, err = store.DB.Exec(`UPDATE tasks SET position = $2 WHERE id = $1;`, id, position)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//Comments

//CreateComment is for inserting comment in database
func (store *DBStore) CreateComment(comment *models.Comment) (id int, err error) {
	res, err := store.DB.Exec(`INSERT INTO comments (text, date) VALUES ($1, $2);`, comment.Text, comment.Date)
	if err != nil {
		log.Fatal(err)
	}
	lastid, err := res.LastInsertId()
	id = int(lastid)
	return id, err

}

//GetComment is for getting comment from database
func (store *DBStore) GetComment(id int) (comment *models.Comment, err error) {
	comment = new(models.Comment)

	err = store.DB.QueryRow(`SELECT id, text, date FROM comments WHERE id=$1;`, id).Scan(&comment)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return comment, err

}

//GetAllComments is for getting  all comments from database
func (store *DBStore) GetAllComments() (comments []*models.Comment, err error) {

	comments = []*models.Comment{}
	rows, err := store.DB.Query(`SELECT id,  text, date  FROM comments ORDER BY date DESC;`)
	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		comment := new(models.Comment)
		err = rows.Scan(&comment.ID, &comment.Text, &comment.Date)
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

//UpdateComment is for  updating   comment in database
func (store *DBStore) UpdateComment(id int, comment *models.Comment) (err error) {

	_, err = store.DB.Exec(`UPDATE comments SET text = $2 date = $3  WHERE id = $1;`, id, comment.Text, comment.Date)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err

}

//DeleteComment is for  deleting  comment in database
func (store *DBStore) DeleteComment(id int) (err error) {
	_, err = store.DB.Exec(`DELETE FROM comments WHERE id = $1;`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//Connections

//CreateProjectColumns inserting ids of projects and columns to database
func (store *DBStore) CreateProjectColumns(projid int, colid int) (err error) {
	_, err = store.DB.Exec(`INSERT INTO projects_columns (project_id, column_id) VALUES ($1, $2);`, projid, colid)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//CreateColumnTasks inserting ids of columns and tasks to database
func (store *DBStore) CreateColumnTasks(taskid int, colid int) (err error) {
	_, err = store.DB.Exec(`INSERT INTO columns_tasks (column_id, task_id) VALUES ($1, $2);`, colid, taskid)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//CreateTaskComments inserting ids of tasks and comments to database
func (store *DBStore) CreateTaskComments(taskid int, comid int) (err error) {
	_, err = store.DB.Exec(`INSERT INTO tasks_comments (task_id, comment_id) VALUES ($1, $2);`, taskid, comid)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//UpdateColumnTasks is for updating columns_tasks table
func (store *DBStore) UpdateColumnTasks(colid int, taskid int) (err error) {
	_, err = store.DB.Exec(`UPDATE columns_tasks SET column_id = $1 task_id = $2;`, colid, taskid)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//MoveTasks is for moving tasks from one column to another
func (store *DBStore) MoveTasks(colid int, nextid int) (err error) {
	_, err = store.DB.Exec(`UPDATE columns_tasks SET column_id = $2 WHERE column_id = $1;`, colid, nextid)

	if err == sql.ErrNoRows {
		log.Fatal("No Results Found")
	}
	if err != nil {
		log.Fatal(err)
	}
	return err
}
