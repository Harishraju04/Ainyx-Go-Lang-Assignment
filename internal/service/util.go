package service

import "time"

func CalculateAge(dob time.Time) *int32 {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	agePtr := int32(age)
	return &agePtr
}

func ParseDate(dob string) (time.Time, error) {
	parsedDob, err := time.Parse(
		"2006-01-02",
		dob,
	)
	return parsedDob, err
}
