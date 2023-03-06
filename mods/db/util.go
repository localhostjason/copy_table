package db

func Like(s string) string {
	return "%" + s + "%"
}
