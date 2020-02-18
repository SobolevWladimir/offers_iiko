package status

func FindFromStatusCode(code string) (Statuses, error) {
	sts, err := FindAll()
	if err != nil {
		return Statuses{}, err
	}
	return sts.GetStatusesFromStatusCode(code)
}
