package utils

var ENV string
var PORT string

func InitENV() {
	if PORT == "" {
		PORT = "8080"
	}

	if ENV == "" {
		ENV = "dev"
	}
}

type EnvironmentVariables struct {
	ENV  string
	PORT string
}

func GetENVVars() EnvironmentVariables {
	return EnvironmentVariables{
		ENV:  ENV,
		PORT: PORT,
	}
}
