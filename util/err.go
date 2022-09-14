package util

func HandleErr(err error) {
	if err != nil {
		LogError(err)
	}
}
