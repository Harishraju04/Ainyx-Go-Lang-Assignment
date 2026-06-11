package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: *s,
		Valid:  true,
	}
}

func ToDate(time *time.Time) pgtype.Date {
	if time == nil {
		return pgtype.Date{}
	}

	return pgtype.Date{
		Time:  *time,
		Valid: true,
	}
}
