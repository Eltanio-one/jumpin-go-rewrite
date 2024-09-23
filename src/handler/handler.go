package handler

func formatDate(date string) string {
	var d string
	for _, char := range date {
		charS := string(char)
		if charS != "T" {
			d += charS
			continue
		}
		break
	}
	return d
}
