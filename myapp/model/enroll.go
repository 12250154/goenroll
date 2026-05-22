package model

import (
	"myapp/datastore/postgres"
)

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll (std_id, course_id, date_enrolled) VALUES ($1, $2, $3);"

func (e *Enroll) EnrollStud() error {
	//fmt.Println("e in the model", e)
	_, err := postgres.Db.Exec(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	return err
}

const queryGetEnrollments = `SELECT std_id, course_id, date_enrolled FROM enroll ORDER BY date_enrolled DESC;`

func GetAllEnrollments() ([]Enroll, error) {

	rows, err := postgres.Db.Query(queryGetEnrollments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enrollments := []Enroll{}

	for rows.Next() {
		var e Enroll
		err := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if err != nil {
			return enrollments, err
		}
		enrollments = append(enrollments, e)
	}

	return enrollments, nil
}
