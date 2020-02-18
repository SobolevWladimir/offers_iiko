package offerentity

func containsInArray(target string, arr []string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
