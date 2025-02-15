-- name: CreateNotification :one
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
) RETURNING *;

-- name: GetNotificationsList :many
SELECT 
    * 
FROM 
    notifications
WHERE 
    USER_ID = $1
ORDER BY NOTIFICATION_ID DESC;