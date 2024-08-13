package customDb

// GetDsnString from the passed map returns a string of settings for the database if there are keys, otherwise an empty string.
func GetDsnString(envData map[string]string) string {
	var dsnStr string
	dsnStr = "host=db "
	if val, ok := envData["DB_USER"]; ok {
		dsnStr += "user=" + val + " "
	}
	if val, ok := envData["DB_PASSWORD"]; ok {
		dsnStr += "password=" + val + " "
	}
	if val, ok := envData["DB_NAME"]; ok {
		dsnStr += "dbname=" + val + " "
	}
	dsnStr += "port=5432 sslmode=disable TimeZone=Europe/Samara"
	return dsnStr
}
