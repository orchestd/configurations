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
	SqlConnection   string `envconfig:"SQL_CON" json:"SQL_CON"`
	CacheConnection string `envconfig:"CACHE_CON" json:"CACHE_CON"`
}
