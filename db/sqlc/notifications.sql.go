// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: notifications.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createNotification = `-- name: CreateNotification :one
INSERT INTO notifications (
    NOTIFICATION_ID,
    USER_ID,
    TITLE,
    MESSAGE,
    TYPE
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING notification_id, user_id, title, message, type, is_read, created_at, updated_at
`

type CreateNotificationParams struct {
	NotificationID uuid.UUID   `json:"notification_id"`
	UserID         uuid.UUID   `json:"user_id"`
	Title          pgtype.Text `json:"title"`
	Message        pgtype.Text `json:"message"`
	Type           string      `json:"type"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error) {
	row := q.db.QueryRow(ctx, createNotification,
		arg.NotificationID,
		arg.UserID,
		arg.Title,
		arg.Message,
		arg.Type,
	)
	var i Notification
	err := row.Scan(
		&i.NotificationID,
		&i.UserID,
		&i.Title,
		&i.Message,
		&i.Type,
		&i.IsRead,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNotificationsList = `-- name: GetNotificationsList :many
SELECT 
    notification_id, user_id, title, message, type, is_read, created_at, updated_at 
FROM 
    notifications
WHERE 
    USER_ID = $1
ORDER BY NOTIFICATION_ID DESC
`

func (q *Queries) GetNotificationsList(ctx context.Context, userID uuid.UUID) ([]Notification, error) {
	rows, err := q.db.Query(ctx, getNotificationsList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Notification{}
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.NotificationID,
			&i.UserID,
			&i.Title,
			&i.Message,
			&i.Type,
			&i.IsRead,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
