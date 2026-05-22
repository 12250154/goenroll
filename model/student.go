package model

import "myapp/datastore/postgres"

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES($1, $2, $3, $4);"
const queryGetUser = "SELECT * FROM student WHERE stdid = $1;"
const queryUpdateUser = "UPDATE student set stdid = $1, firstname = $2, lastname = $3, email = $4 WHERE stdid = $5 RETURNING stdid;"
const queryDeleteUser = "DELETE from student where stdid=$1 RETURNING stdid;"

// model is responsible for interacting with database
// stud -> s
func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

func (s *Student) Read() error {
	row := postgres.Db.QueryRow(queryGetUser, s.StdId)
	return row.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)

}

func (s *Student) Update(oldID int64) error {
	row := postgres.Db.QueryRow(queryUpdateUser, s.StdId, s.FirstName, s.LastName, s.Email, oldID)
	return row.Scan(&s.StdId)
}

// method
func (s *Student) Delete() error {
	return postgres.Db.QueryRow(queryDeleteUser, s.StdId).Scan(&s.StdId) //scan is used to return an error

}

// function
// func Delete(StdId int64) error {

// }

func GetAllStudents() ([]Student, error) {
	rows, geterr := postgres.Db.Query("Select * from student;")

	if geterr != nil {
		return nil, geterr
	}
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return students, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}
