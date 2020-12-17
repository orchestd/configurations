package credentials

type Builder interface {
	DevMode() Builder
	Build() (CredentialsGetter, error)
}

type CredentialsGetter interface {
	GetCredentials() Credentials
	Implementation() interface{}
}
type Credentials struct {
	DbUsername string `envconfig:"DB_USERNAME" required:"true"`
	DbPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DbHost     string `envconfig:"DB_HOST" required:"true"`
	DbName     string `envconfig:"DB_NAME" required:"true"`
}
