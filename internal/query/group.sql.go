// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: group.sql

package query

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createGroupAdmin = `-- name: CreateGroupAdmin :exec
INSERT INTO group_admin (
    group_name, position, roll_number
)
VALUES (
    (SELECT name from groups WHERE name = $1 or alias = $1),
    $2,
    $3
)
`

type CreateGroupAdminParams struct {
	Name       string `json:"name"`
	Position   string `json:"position"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) CreateGroupAdmin(ctx context.Context, arg CreateGroupAdminParams) error {
	_, err := q.db.ExecContext(ctx, createGroupAdmin, arg.Name, arg.Position, arg.RollNumber)
	return err
}

const createGroupFaculty = `-- name: CreateGroupFaculty :exec
INSERT INTO group_faculty (
    group_name, emp_id
)
VALUES (
    (SELECT g.name from groups g WHERE g.name = $1 or g.alias = $1),
    $2
)
`

type CreateGroupFacultyParams struct {
	Name  string `json:"name"`
	EmpID int32  `json:"emp_id"`
}

func (q *Queries) CreateGroupFaculty(ctx context.Context, arg CreateGroupFacultyParams) error {
	_, err := q.db.ExecContext(ctx, createGroupFaculty, arg.Name, arg.EmpID)
	return err
}

const createGroupMember = `-- name: CreateGroupMember :exec
INSERT INTO group_member (
    group_name, roll_number
)
VALUES (
    (SELECT name from groups WHERE name = $1 or alias = $1),
    $2
)
`

type CreateGroupMemberParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) CreateGroupMember(ctx context.Context, arg CreateGroupMemberParams) error {
	_, err := q.db.ExecContext(ctx, createGroupMember, arg.Name, arg.RollNumber)
	return err
}

const createGroupSocial = `-- name: CreateGroupSocial :exec
INSERT INTO group_social (
    name, platform_type, link
)
VALUES (
    (SELECT g.name from groups g WHERE g.name = $1 or g.alias = $1),
    $2,
    $3
)
`

type CreateGroupSocialParams struct {
	Name         string `json:"name"`
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
}

func (q *Queries) CreateGroupSocial(ctx context.Context, arg CreateGroupSocialParams) error {
	_, err := q.db.ExecContext(ctx, createGroupSocial, arg.Name, arg.PlatformType, arg.Link)
	return err
}

const deleteGroupAdmin = `-- name: DeleteGroupAdmin :exec
DELETE FROM group_admin
WHERE
    group_name = (SELECT name FROM groups WHERE name = $1 OR alias = $1)
    AND roll_number = $2
`

type DeleteGroupAdminParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) DeleteGroupAdmin(ctx context.Context, arg DeleteGroupAdminParams) error {
	_, err := q.db.ExecContext(ctx, deleteGroupAdmin, arg.Name, arg.RollNumber)
	return err
}

const deleteGroupFaculty = `-- name: DeleteGroupFaculty :exec
DELETE FROM group_faculty gf
WHERE
    gf.group_name = (SELECT g.name FROM groups g WHERE g.name = $1 OR g.alias = $1)
    AND gf.emp_id = $2
`

type DeleteGroupFacultyParams struct {
	Name  string `json:"name"`
	EmpID int32  `json:"emp_id"`
}

func (q *Queries) DeleteGroupFaculty(ctx context.Context, arg DeleteGroupFacultyParams) error {
	_, err := q.db.ExecContext(ctx, deleteGroupFaculty, arg.Name, arg.EmpID)
	return err
}

const deleteGroupMember = `-- name: DeleteGroupMember :exec
DELETE FROM group_member
WHERE
    group_name = (SELECT name FROM groups WHERE name = $1 OR alias = $1)
    AND roll_number = $2
`

type DeleteGroupMemberParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) DeleteGroupMember(ctx context.Context, arg DeleteGroupMemberParams) error {
	_, err := q.db.ExecContext(ctx, deleteGroupMember, arg.Name, arg.RollNumber)
	return err
}

const deleteGroupSocial = `-- name: DeleteGroupSocial :exec
DELETE FROM group_social
WHERE
    group_name = (SELECT name FROM groups WHERE name = $1 OR alias = $1)
    AND platform_type = $2
`

type DeleteGroupSocialParams struct {
	Name         string `json:"name"`
	PlatformType string `json:"platform_type"`
}

func (q *Queries) DeleteGroupSocial(ctx context.Context, arg DeleteGroupSocialParams) error {
	_, err := q.db.ExecContext(ctx, deleteGroupSocial, arg.Name, arg.PlatformType)
	return err
}

const getGroup = `-- name: GetGroup :one
SELECT
    g.name, g.alias, g.branch, g.kind, g.description,
    CAST(ARRAY(SELECT f.name FROM faculty AS f JOIN group_faculty AS gf ON f.emp_id = gf.emp_id WHERE g.name = gf.group_name) AS text[]) AS faculty_names,
    CAST(ARRAY(SELECT f.mobile FROM faculty AS f JOIN group_faculty AS gf ON f.emp_id = gf.emp_id WHERE g.name = gf.group_name) AS text[]) AS faculty_mobiles,
    CAST(ARRAY(SELECT gs.platform_type FROM group_social AS gs WHERE g.name = gs.group_name) AS text[]) AS social_types,
    CAST(ARRAY(SELECT gs.link FROM group_social AS gs WHERE g.name = gs.group_name) AS text[]) AS social_links,
    CAST(ARRAY(SELECT ga.position FROM group_admin AS ga WHERE g.name = ga.group_name) AS text[]) AS admin_positions,
    CAST(ARRAY(SELECT ga.roll_number FROM group_admin AS ga WHERE g.name = ga.group_name) AS bigint[]) AS admin_rolls,
    CAST(ARRAY(SELECT gm.roll_number FROM group_member AS gm WHERE g.name = gm.group_name) AS bigint[]) AS members
FROM
    groups AS g
WHERE
    g.name = $1
    OR g.alias = $1
`

type GetGroupRow struct {
	Name           string         `json:"name"`
	Alias          sql.NullString `json:"alias"`
	Branch         []string       `json:"branch"`
	Kind           string         `json:"kind"`
	Description    string         `json:"description"`
	FacultyNames   []string       `json:"faculty_names"`
	FacultyMobiles []string       `json:"faculty_mobiles"`
	SocialTypes    []string       `json:"social_types"`
	SocialLinks    []string       `json:"social_links"`
	AdminPositions []string       `json:"admin_positions"`
	AdminRolls     []int64        `json:"admin_rolls"`
	Members        []int64        `json:"members"`
}

func (q *Queries) GetGroup(ctx context.Context, name string) (GetGroupRow, error) {
	row := q.db.QueryRowContext(ctx, getGroup, name)
	var i GetGroupRow
	err := row.Scan(
		&i.Name,
		&i.Alias,
		pq.Array(&i.Branch),
		&i.Kind,
		&i.Description,
		pq.Array(&i.FacultyNames),
		pq.Array(&i.FacultyMobiles),
		pq.Array(&i.SocialTypes),
		pq.Array(&i.SocialLinks),
		pq.Array(&i.AdminPositions),
		pq.Array(&i.AdminRolls),
		pq.Array(&i.Members),
	)
	return i, err
}

const getGroupAdmins = `-- name: GetGroupAdmins :many
SELECT
    s.roll_number, s.section, s.name, s.gender, s.mobile, s.birth_date, s.email, s.batch, s.hostel_id, s.room_id, s.discord_id, s.is_verified, admin.position
FROM
    student s
    JOIN group_admin admin ON s.roll_number = admin.roll_number
WHERE
    admin.group_name = $1
    OR $1 = (SELECT alias FROM groups WHERE name = admin.group_name)
`

type GetGroupAdminsRow struct {
	RollNumber string         `json:"roll_number"`
	Section    string         `json:"section"`
	Name       string         `json:"name"`
	Gender     sql.NullString `json:"gender"`
	Mobile     sql.NullString `json:"mobile"`
	BirthDate  sql.NullTime   `json:"birth_date"`
	Email      string         `json:"email"`
	Batch      int16          `json:"batch"`
	HostelID   sql.NullString `json:"hostel_id"`
	RoomID     sql.NullString `json:"room_id"`
	DiscordID  sql.NullInt64  `json:"discord_id"`
	IsVerified bool           `json:"is_verified"`
	Position   string         `json:"position"`
}

func (q *Queries) GetGroupAdmins(ctx context.Context, groupName string) ([]GetGroupAdminsRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroupAdmins, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupAdminsRow
	for rows.Next() {
		var i GetGroupAdminsRow
		if err := rows.Scan(
			&i.RollNumber,
			&i.Section,
			&i.Name,
			&i.Gender,
			&i.Mobile,
			&i.BirthDate,
			&i.Email,
			&i.Batch,
			&i.HostelID,
			&i.RoomID,
			&i.DiscordID,
			&i.IsVerified,
			&i.Position,
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

const getGroupFaculty = `-- name: GetGroupFaculty :many
SELECT
    f.name, f.mobile
FROM
    faculty AS f
    JOIN group_faculty AS gf ON f.emp_id = gf.emp_id
WHERE
    gf.group_name = $1
    OR $1 = (SELECT alias FROM groups WHERE name = gf.group_name)
`

type GetGroupFacultyRow struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func (q *Queries) GetGroupFaculty(ctx context.Context, groupName string) ([]GetGroupFacultyRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroupFaculty, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupFacultyRow
	for rows.Next() {
		var i GetGroupFacultyRow
		if err := rows.Scan(&i.Name, &i.Mobile); err != nil {
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

const getGroupMembers = `-- name: GetGroupMembers :many
SELECT
    s.roll_number, s.section, s.name, s.gender, s.mobile, s.birth_date, s.email, s.batch, s.hostel_id, s.room_id, s.discord_id, s.is_verified
FROM
    student s
    JOIN group_member member ON s.roll_number = member.roll_number
WHERE
    member.group_name = $1
    OR $1 = (SELECT alias FROM groups WHERE name = member.group_name)
`

func (q *Queries) GetGroupMembers(ctx context.Context, groupName string) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getGroupMembers, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.RollNumber,
			&i.Section,
			&i.Name,
			&i.Gender,
			&i.Mobile,
			&i.BirthDate,
			&i.Email,
			&i.Batch,
			&i.HostelID,
			&i.RoomID,
			&i.DiscordID,
			&i.IsVerified,
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

const getGroupSocials = `-- name: GetGroupSocials :many
SELECT
    platform_type,
    link
FROM
    group_social
WHERE
    group_name = $1
    OR $1 = (SELECT alias FROM groups WHERE name = group_name)
`

type GetGroupSocialsRow struct {
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
}

func (q *Queries) GetGroupSocials(ctx context.Context, groupName string) ([]GetGroupSocialsRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroupSocials, groupName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupSocialsRow
	for rows.Next() {
		var i GetGroupSocialsRow
		if err := rows.Scan(&i.PlatformType, &i.Link); err != nil {
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

const getGroups = `-- name: GetGroups :many
SELECT
    g.name, g.alias, g.branch, g.kind, g.description,
    CAST(ARRAY(SELECT f.name FROM faculty AS f JOIN group_faculty AS gf ON f.emp_id = gf.emp_id WHERE g.name = gf.group_name) AS text[]) AS faculty_names,
    CAST(ARRAY(SELECT f.mobile FROM faculty AS f JOIN group_faculty AS gf ON f.emp_id = gf.emp_id WHERE g.name = gf.group_name) AS text[]) AS faculty_mobiles,
    CAST(ARRAY(SELECT gs.platform_type FROM group_social AS gs WHERE g.name = gs.group_name) AS text[]) AS social_types,
    CAST(ARRAY(SELECT gs.link FROM group_social AS gs WHERE g.name = gs.group_name) AS text[]) AS social_links,
    CAST(ARRAY(SELECT ga.position FROM group_admin AS ga WHERE g.name = ga.group_name) AS text[]) AS admin_positions,
    CAST(ARRAY(SELECT ga.roll_number FROM group_admin AS ga WHERE g.name = ga.group_name) AS bigint[]) AS admin_rolls,
    CAST(ARRAY(SELECT gm.roll_number FROM group_member AS gm WHERE g.name = gm.group_name) AS bigint[]) AS members
FROM
    groups AS g
`

type GetGroupsRow struct {
	Name           string         `json:"name"`
	Alias          sql.NullString `json:"alias"`
	Branch         []string       `json:"branch"`
	Kind           string         `json:"kind"`
	Description    string         `json:"description"`
	FacultyNames   []string       `json:"faculty_names"`
	FacultyMobiles []string       `json:"faculty_mobiles"`
	SocialTypes    []string       `json:"social_types"`
	SocialLinks    []string       `json:"social_links"`
	AdminPositions []string       `json:"admin_positions"`
	AdminRolls     []int64        `json:"admin_rolls"`
	Members        []int64        `json:"members"`
}

func (q *Queries) GetGroups(ctx context.Context) ([]GetGroupsRow, error) {
	rows, err := q.db.QueryContext(ctx, getGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupsRow
	for rows.Next() {
		var i GetGroupsRow
		if err := rows.Scan(
			&i.Name,
			&i.Alias,
			pq.Array(&i.Branch),
			&i.Kind,
			&i.Description,
			pq.Array(&i.FacultyNames),
			pq.Array(&i.FacultyMobiles),
			pq.Array(&i.SocialTypes),
			pq.Array(&i.SocialLinks),
			pq.Array(&i.AdminPositions),
			pq.Array(&i.AdminRolls),
			pq.Array(&i.Members),
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

const updateGroupSocials = `-- name: UpdateGroupSocials :exec
UPDATE
    group_social
SET
    link = $2
WHERE
    platform_type = $1
    AND group_name = $3
    OR $3 = (SELECT alias FROM groups WHERE name = group_name)
`

type UpdateGroupSocialsParams struct {
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
	GroupName    string `json:"group_name"`
}

func (q *Queries) UpdateGroupSocials(ctx context.Context, arg UpdateGroupSocialsParams) error {
	_, err := q.db.ExecContext(ctx, updateGroupSocials, arg.PlatformType, arg.Link, arg.GroupName)
	return err
}
