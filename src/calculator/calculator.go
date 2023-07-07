package calculator

import (
	"strconv"
	"strings"
	"time"
)

type AgeCalculator interface {
	CalculateAge(dateStr string) int
}

type ageCalculator struct{}

func NewAgeCalculator() AgeCalculator {
	return &ageCalculator{}
}

func (ac *ageCalculator) CalculateAge(dateStr string) int {
	now := time.Now()

	parts := strings.Split(dateStr, "/")
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	birthDay := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	age := now.Year() - birthDay.Year()

	if now.Month() < birthDay.Month() ||
		(now.Month() == birthDay.Month() && now.Day() < birthDay.Day()) {
		age--
	}
	return age
}
