package credentials

type Builder interface {
	UseGcpSecretManager(projectId string) Builder
	SetSecretManagerVersion(version string) Builder
	SetSecretName(name string) Builder
	Build() (CredentialsGetter, error)
}

type CredentialsGetter interface {
	GetCredentials() Credentials
	Implementation() interface{}
}

type Credentials struct {
	CacheUserName string `envconfig:"CACHE_USER_NAME" json:"CACHE_USER_NAME"`
	CacheUserPw   string `envconfig:"CACHE_USER_PW" json:"CACHE_USER_PW"`

	SqlUserName string `envconfig:"SQL_USER_NAME" json:"SQL_USER_NAME"`
	SqlUserPw   string `envconfig:"SQL_USER_PW" json:"SQL_USER_PW"`

	NatsUser string `envconfig:"NATS_USER" json:"NATS_USER"`
	NatsPw   string `envconfig:"NATS_PW" json:"NATS_PW"`

	EncryptKey string `envconfig:"ENCRYPT_KEY" json:"ENCRYPT_KEY"`
	JwtSecret  string `envconfig:"JWT_SECRET" json:"JWT_SECRET"`

	MessagesProviders string `envconfig:"MESSAGES_PROVIDERS" json:"MESSAGES_PROVIDERS"`

	MapaLoginName string `envconfig:"MAPA_LOGIN_NAME" json:"MAPA_LOGIN_NAME"`
	MapaToken     string `envconfig:"MAPA_TOKEN" json:"MAPA_TOKEN"`
}
