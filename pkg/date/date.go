package date

import (
	"fmt"
	"strconv"
	"time"
)

func ExpireDateString(notafter time.Time) string {
	year := strconv.Itoa(notafter.Year())
	month := strconv.Itoa(int(notafter.Month()))
	day := strconv.Itoa(notafter.Day())

	if len(month) == 1 {
		month = fmt.Sprintf("0%s", month)
	}
	if len(day) == 1 {
		day = fmt.Sprintf("0%s", day)
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

func ExpireDiffDays(notafter time.Time) int {
	days := 1

	now := time.Now()
	diff := notafter.Sub(now)

	hours := int(diff.Hours())

	if hours > 24 {
		days = hours / 24
	}

	return days
}
