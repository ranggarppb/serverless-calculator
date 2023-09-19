package utils

func ContainString(list []string, target string) bool {
	for _, el := range list {
		if el == target {
			return true
		}
	}
	return false
}

func Revert(list []string) []string {
	res := []string{}

	for i, _ := range list {
		res = append(res, list[len(list)-i-1])
	}

	return res
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
