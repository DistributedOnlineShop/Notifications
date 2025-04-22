package util

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/exp/rand"
)

func GenerateTitle() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookTitle(),
		Valid:  true,
	}
}
func GenerateMessage() pgtype.Text {
	return pgtype.Text{
		String: gofakeit.BookGenre(),
		Valid:  true,
	}
}

func GenerateType() string {
	notifications := []string{
		"Order-Related",
		"Payment-Related",
		"Promotions",
		"Account & Security",
		"Customer Support",
		"Social",
		"Product & Subscription",
	}

	return notifications[rand.Intn(len(notifications))]
}
