package utils

func IndexOf(values []string, target string) int {
	for i, v := range values {
		if v == target {
			return i
		}
	}

	return -1
}
