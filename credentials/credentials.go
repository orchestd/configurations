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
	MySqlConnection string `envconfig:"MYSQL_CON" required:"true" json:"MYSQL_CON"`
	MySqlDbName     string `envconfig:"MYSQL_DB_NAME" required:"true" json:"MYSQL_DB_NAME"`
	MongoConnection string `envconfig:"MONGO_CON" required:"true" json:"MONGO_CON"`
	MongoDbName     string `envconfig:"MONGO_DB_NAME" required:"true" json:"MONGO_DB_NAME"`
}
