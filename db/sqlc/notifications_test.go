package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"Notifications/util"
)

func CreateRandomNotification(t *testing.T, userID uuid.UUID) Notification {
	data := CreateNotificationParams{
		NotificationID: util.CreateUUID(),
		UserID:         userID,
		Title:          util.GenerateTitle(),
		Message:        util.GenerateMessage(),
		Type:           util.GenerateType(),
	}

	notif, err := testStore.CreateNotification(context.Background(), data)
	require.NoError(t, err)
	require.NotEmpty(t, notif)
	require.Equal(t, data.NotificationID, notif.NotificationID)
	require.Equal(t, data.UserID, notif.UserID)
	require.Equal(t, data.Title.String, notif.Title.String)
	require.Equal(t, data.Message.String, notif.Message.String)
	require.Equal(t, data.Type, notif.Type)
	require.False(t, notif.IsRead)
	require.NotZero(t, notif.CreatedAt)

	return notif
}

func TestCreateNotification(t *testing.T) {
	id := util.CreateUUID()
	CreateRandomNotification(t, id)
}

func TestGetNotificationsList(t *testing.T) {
	id := util.CreateUUID()
	for i := 0; i < 10; i++ {
		CreateRandomNotification(t, id)
	}

	datalist, err := testStore.GetNotificationsList(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, datalist)
	require.GreaterOrEqual(t, len(datalist), 10)
}
