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
	CacheUserName string `envconfig:"CACHE_USER_NAME" json:"CACHE_USER_NAME"`
	CacheUserPw  string `envconfig:"CACHE_USER_PW" json:"CACHE_USER_PW"`

	SqlUserName string `envconfig:"SQL_USER_NAME" json:"SQL_USER_NAME"`
	SqlUserPw  string `envconfig:"SQL_USER_PW" json:"SQL_USER_PW"`
}
