// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package query

import (
	"context"
)

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