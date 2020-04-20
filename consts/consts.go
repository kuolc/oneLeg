package consts

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env not found")
		return
	}
}

type Environment = string

const (
	EnvironmentLocal       = "local"
	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

func Env() Environment {
	return Environment(os.Getenv("ENV"))
}

func IsLocal() bool {
	return Env() == EnvironmentLocal
}

func IsDevelopment() bool {
	return Env() == EnvironmentDevelopment
}

func IsProduction() bool {
	return Env() == EnvironmentProduction
}

func TLSCertificatePath() string {
	return os.Getenv("TLS_CERTIFICATE_PATH")
}

func TLSPrivateKeyPath() string {
	return os.Getenv("TLS_PRIVATE_KEY_PATH")
}

func ChannelSecret(botName string) string {
	return os.Getenv(botName + "_SECRET")
}

func ChannelAccessToken(botName string) string {
	return os.Getenv(botName + "_ACCESS_TOKEN")
}

func GroupID(botName string) string {
	return os.Getenv(botName + "_GROUP_ID")
}

func ProblemTemplatePath() string {
	return os.Getenv("PROBLEM_TEMPLATE_PATH")
}

func EditorialTemplatePath() string {
	return os.Getenv("EDITORIAL_TEMPLATE_PATH")
}

func GoogleCredentialPath() string {
	return os.Getenv("GOOGLE_CREDENTIAL_PATH")
}

func ProblemSheetID() string {
	return os.Getenv("PROBLEM_SHEET_ID")
}

func PushProblemAt() int {
	return 9
}

func PushEditorialAt() int {
	return 19
}
