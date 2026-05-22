package model

import "myapp/datastore/postgres"

// course

type Course struct {
	CID        int    `json:"courseid"`
	CourseName string `json:"coursename"`
}

const queryInsertCourse = "INSERT INTO course(cid, coursename) VALUES($1, $2);"
const queryGetCourse = "SELECT * FROM course WHERE cid = $1;"
const queryUpdateCourse = "UPDATE course SET cid = $1, coursename = $2 WHERE cid = $3 RETURNING cid;"
const queryDeleteCourse = "DELETE from course where cid=$1 RETURNING cid;"

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertCourse, c.CID, c.CourseName)
	return err
}

func (c *Course) Read() error {
	row := postgres.Db.QueryRow(queryGetCourse, c.CID)
	return row.Scan(&c.CID, &c.CourseName)

}

func (c *Course) Update(oldCID int64) error {
	row := postgres.Db.QueryRow(
		queryUpdateCourse,
		c.CID,        // new cid
		c.CourseName, // new name
		oldCID,       // old cid in URL
	)
	return row.Scan(&c.CID)
}

func (c *Course) Delete() error {
	return postgres.Db.QueryRow(queryDeleteCourse, c.CID).Scan(&c.CID) //scan is used to return an error

}

//get all courses

func GetAllCourses() ([]Course, error) {
	rows, geterr := postgres.Db.Query("Select * from course;")

	if geterr != nil {
		return nil, geterr
	}
	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.CID, &c.CourseName)
		if dbErr != nil {
			return courses, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}
