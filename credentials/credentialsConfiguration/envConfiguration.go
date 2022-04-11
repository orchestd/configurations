package credentialsConfiguration

type CredentialsConfiguration struct {
	SecretManager        bool    `json:"ENABLE_SECRET_MANAGER,omitempty"`
	SecretManagerVersion *string `json:"SECRET_MANAGER_VERSION,omitempty"`
}
