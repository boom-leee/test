// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertUser = `-- name: InsertUser :one
INSERT INTO users (
    id, auth_id, email, name, role, image_url, created_at, updated_at, deleted_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         )
    RETURNING id, auth_id, email, name, role, image_url, created_at, updated_at, deleted_at
`

type InsertUserParams struct {
	ID        pgtype.UUID
	AuthID    string
	Email     pgtype.Text
	Name      pgtype.Text
	Role      pgtype.Text
	ImageUrl  pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRow(ctx, insertUser,
		arg.ID,
		arg.AuthID,
		arg.Email,
		arg.Name,
		arg.Role,
		arg.ImageUrl,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AuthID,
		&i.Email,
		&i.Name,
		&i.Role,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
