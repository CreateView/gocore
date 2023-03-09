package config

import (
	"fmt"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Config struct {
	Env             string `env:"ENV,default=production"`
	ProjectID       string `env:"PROJECT_ID,default=createviewprototypeproject"`
	SubscriptionID  string `env:"SUBSCRIPTION_ID,default=sensor-pg-data"`
	CredentialsFile string `env:"GOOGLE_APPLICATION_CREDENTIALS"`
	PostgresDNS     string `env:"POSTGRES_DNS"`
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		want    *Config
		wantErr string
	}{
		{
			name: "ok-defaults",
			envVars: map[string]string{
				"GOOGLE_APPLICATION_CREDENTIALS": "file.json",
				"POSTGRES_DNS":                   "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
			},
			want: &Config{
				Env:             "production",
				ProjectID:       "createviewprototypeproject",
				SubscriptionID:  "sensor-pg-data",
				CredentialsFile: "file.json",
				PostgresDNS:     "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
			},
		},
		{
			name: "missing-envs",
			envVars: map[string]string{
				"ENV":                            "staging",
				"PROJECT_ID":                     "",
				"SUBSCRIPTION_ID":                "subscription-id",
				"GOOGLE_APPLICATION_CREDENTIALS": "file.json",
				"POSTGRES_DNS":                   "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
			},
			wantErr: `ProjectID: cannot be blank.`,
		},
		{
			name: "invalid-env-type",
			envVars: map[string]string{
				"ENV":                            "not-a-valid-env-type",
				"PROJECT_ID":                     "createview",
				"SUBSCRIPTION_ID":                "subscription-id",
				"GOOGLE_APPLICATION_CREDENTIALS": "file.json",
				"POSTGRES_DNS":                   "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
			},
			wantErr: `Env: 'not-a-valid-env-type' is not a valid value. Valid values: development, staging, production.`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookuper := envconfig.MapLookuper(tt.envVars)

			cfg := &Config{}
			err := load(cfg, lookuper)
			assert.NoError(t, err)

			err = validate(
				cfg,
				validation.Field(
					&cfg.Env,
					validation.Required,
					validation.In("development", "staging", "production").Error(
						fmt.Sprintf(
							"'%s' is not a valid value. Valid values: development, staging, production", cfg.Env),
					),
				),
				validation.Field(&cfg.ProjectID, validation.Required),
				validation.Field(&cfg.SubscriptionID, validation.Required),
				validation.Field(&cfg.CredentialsFile, validation.Required),
				validation.Field(&cfg.PostgresDNS, validation.Required),
			)

			if len(tt.wantErr) > 0 {
				require.Error(t, err)
				require.Equal(t, tt.wantErr, err.Error())
				return
			}
			assert.Equal(t, tt.want, cfg)
		})
	}
}
