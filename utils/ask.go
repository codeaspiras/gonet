package utils

func Ask(template string, args ...interface{}) string {
	return contact(template+" > ", args)
}
