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
	SqlConnectionString string `envconfig:"SQL_CON_STR" json:"SQL_CON_STR"`
	CacheConnection     string `envconfig:"CACHE_CON" json:"CACHE_CON"`
}
