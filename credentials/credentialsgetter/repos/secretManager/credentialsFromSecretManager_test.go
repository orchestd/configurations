package secretManager

import (
	"testing"
)

func TestNewCredentialsFromEnvVariables(t *testing.T) {
	type args struct {
		projectId string
		version   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "testCreds", args: args{
			projectId: "groovy-autumn-242113",
			version:   "",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCredentialsFromSecretManager(tt.args.projectId, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCredentialsFromEnvVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("NewCredentialsFromEnvVariables() got = nil, want Credentials struct")
			}
		})
	}
}
