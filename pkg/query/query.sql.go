// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package query

import (
	"context"

	"github.com/lib/pq"
)

const getAllCourses = `-- name: GetAllCourses :many
SELECT code, title, branch, semester, credits, prereq, type, objectives, content, books, outcomes FROM course
`

func (q *Queries) GetAllCourses(ctx context.Context) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getAllCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.Code,
			&i.Title,
			&i.Branch,
			&i.Semester,
			pq.Array(&i.Credits),
			pq.Array(&i.Prereq),
			&i.Type,
			&i.Objectives,
			&i.Content,
			&i.Books,
			&i.Outcomes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBranchCourses = `-- name: GetBranchCourses :many
SELECT code, title, branch, semester, credits, prereq, type, objectives, content, books, outcomes FROM course
WHERE branch = $1
`

func (q *Queries) GetBranchCourses(ctx context.Context, branch string) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getBranchCourses, branch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.Code,
			&i.Title,
			&i.Branch,
			&i.Semester,
			pq.Array(&i.Credits),
			pq.Array(&i.Prereq),
			&i.Type,
			&i.Objectives,
			&i.Content,
			&i.Books,
			&i.Outcomes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourse = `-- name: GetCourse :one
SELECT code, title, branch, semester, credits, prereq, type, objectives, content, books, outcomes FROM course
WHERE code = $1 LIMIT 1
`

func (q *Queries) GetCourse(ctx context.Context, code string) (Course, error) {
	row := q.db.QueryRowContext(ctx, getCourse, code)
	var i Course
	err := row.Scan(
		&i.Code,
		&i.Title,
		&i.Branch,
		&i.Semester,
		pq.Array(&i.Credits),
		pq.Array(&i.Prereq),
		&i.Type,
		&i.Objectives,
		&i.Content,
		&i.Books,
		&i.Outcomes,
	)
	return i, err
}

const getCourses = `-- name: GetCourses :many
SELECT code, title, branch, semester, credits, prereq, type, objectives, content, books, outcomes FROM course
WHERE branch = $1 AND semester = $2
`

type GetCoursesParams struct {
	Branch   string
	Semester int16
}

func (q *Queries) GetCourses(ctx context.Context, arg GetCoursesParams) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getCourses, arg.Branch, arg.Semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.Code,
			&i.Title,
			&i.Branch,
			&i.Semester,
			pq.Array(&i.Credits),
			pq.Array(&i.Prereq),
			&i.Type,
			&i.Objectives,
			&i.Content,
			&i.Books,
			&i.Outcomes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSemesterCourses = `-- name: GetSemesterCourses :many
SELECT code, title, branch, semester, credits, prereq, type, objectives, content, books, outcomes FROM course
WHERE semester = $1
`

func (q *Queries) GetSemesterCourses(ctx context.Context, semester int16) ([]Course, error) {
	rows, err := q.db.QueryContext(ctx, getSemesterCourses, semester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Course
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.Code,
			&i.Title,
			&i.Branch,
			&i.Semester,
			pq.Array(&i.Credits),
			pq.Array(&i.Prereq),
			&i.Type,
			&i.Objectives,
			&i.Content,
			&i.Books,
			&i.Outcomes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudent = `-- name: GetStudent :one
SELECT roll_number, section, sub_section, name, gender, mobile, birthday, email, batch, hostel_number, room_number, discord_uid, verified FROM student
WHERE roll_number = $1
`

func (q *Queries) GetStudent(ctx context.Context, rollNumber int32) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, rollNumber)
	var i Student
	err := row.Scan(
		&i.RollNumber,
		&i.Section,
		&i.SubSection,
		&i.Name,
		&i.Gender,
		&i.Mobile,
		&i.Birthday,
		&i.Email,
		&i.Batch,
		&i.HostelNumber,
		&i.RoomNumber,
		&i.DiscordUid,
		&i.Verified,
	)
	return i, err
}