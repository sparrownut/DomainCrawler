package utils

func IsInList(element string, list []string) bool {
	for _, el := range list {
		if el == element {
			return true
		}
	}
	return false
}
