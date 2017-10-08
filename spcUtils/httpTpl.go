package spcUtils

func SetSelected(str1 string, str2 string) string {
	if str1 == str2 {
		return "selected"
	}
	return ""
}
