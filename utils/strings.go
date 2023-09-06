package utils

func ContainString(list []string, target string) bool {
	for _, el := range list {
		if el == target {
			return true
		}
	}
	return false
}
