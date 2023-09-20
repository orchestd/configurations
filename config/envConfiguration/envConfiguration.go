package envConfiguration

type EnvConfiguration struct {
	DebugMode   bool    `json:"debugMode,omitempty"`
	HeilaEnv    string  `json:"HEILA_ENV"`
	ServiceName string  `json:"SERVICE_NAME"`
	ProjectId   *string `json:"PROJECT_ID,omitempty"`
}
