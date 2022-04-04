package envConfiguration

type EnvConfiguration struct {
	DebugMode  bool    `json:"debugMode,omitempty"`
	HeilaEnv   string  `json:"HEILA_ENV"`
	DockerName string  `json:"DOCKER_NAME"`
	ProjectId  *string `json:"PROJECT_ID,omitempty"`
}
