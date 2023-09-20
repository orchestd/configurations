package credentials

import (
	"encoding/json"
	"fmt"
)

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

	MapsProviderUserName string `envconfig:"MAPS_PROVIDER_USER_NAME" json:"MAPS_PROVIDER_USER_NAME"`
	MapsProviderToken    string `envconfig:"MAPS_PROVIDER_TOKEN" json:"MAPS_PROVIDER_TOKEN"`

	CreditCardServiceUserName string `envconfig:"CREDIT_CARD_SERVICE_USER_NAME" json:"CREDIT_CARD_SERVICE_USER_NAME"`
	CreditCardServiceUserPw   string `envconfig:"CREDIT_CARD_SERVICE_USER_PW" json:"CREDIT_CARD_SERVICE_USER_PW"`

	PaymentProviders string `envconfig:"PAYMENT_PROVIDERS" json:"PAYMENT_PROVIDERS"`

	EnablePprofPw bool `envconfig:"ENABLE_PPROF_PW" json:"ENABLE_PPROF_PW"`
}

func (cr Credentials) GetPaymentProvider(name string, provider interface{}) error {
	paymentProviders := map[string]json.RawMessage{}
	err := json.Unmarshal([]byte(cr.PaymentProviders), &paymentProviders)
	if err != nil {
		return fmt.Errorf("can't unmarshal paymentProviders. error: %v", err)
	}

	providerJSON, ok := paymentProviders[name]
	if !ok {
		return fmt.Errorf("payment provider %s not found in credentials PaymentProviders", name)
	}
	err = json.Unmarshal(providerJSON, provider)
	if err != nil {
		return fmt.Errorf("can't unmarshal payment provider %s. error: %v", name, err)
	}
	return nil
}
