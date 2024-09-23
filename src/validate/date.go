package validate

import "time"

func Date(date string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil

}
