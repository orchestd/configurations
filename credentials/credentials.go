package credentials

type Builder interface {
	UseGcpSecretManager(projectId string) Builder
	SetSecretManagerVersion(version string) Builder
	Build() (CredentialsGetter, error)
}

type CredentialsGetter interface {
	GetCredentials() Credentials
	Implementation() interface{}
}
type Credentials struct {
	DbUsername string `envconfig:"DB_USERNAME" required:"true" json:"DB_USERNAME"`
	DbPassword string `envconfig:"DB_PASSWORD" required:"true" json:"DB_PASSWORD"`
	DbHost     string `envconfig:"DB_HOST" required:"true" json:"DB_HOST"`
	DbName     string `envconfig:"DB_NAME" required:"true" json:"DB_NAME"`
}
