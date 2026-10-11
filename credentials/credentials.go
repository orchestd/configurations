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

	SqlReadOnlyUserName string `envconfig:"SQL_READ_ONLY_USER_NAME" json:"SQL_READ_ONLY_USER_NAME"`
	SqlReadOnlyUserPw   string `envconfig:"SQL_READ_ONLY_USER_PW" json:"SQL_READ_ONLY_USER_PW"`

	NatsJwtSignatureSeed string `envconfig:"NATS_JWT_SIGNATURE_SEED" json:"NATS_JWT_SIGNATURE_SEED"`

	EncryptKey string `envconfig:"ENCRYPT_KEY" json:"ENCRYPT_KEY"`
	JwtSecret  string `envconfig:"JWT_SECRET" json:"JWT_SECRET"`

	MessagesProviders string `envconfig:"MESSAGES_PROVIDERS" json:"MESSAGES_PROVIDERS"`

	MapsProviderUserName string `envconfig:"MAPS_PROVIDER_USER_NAME" json:"MAPS_PROVIDER_USER_NAME"`
	MapsProviderToken    string `envconfig:"MAPS_PROVIDER_TOKEN" json:"MAPS_PROVIDER_TOKEN"`

	CreditCardServiceUserName string `envconfig:"CREDIT_CARD_SERVICE_USER_NAME" json:"CREDIT_CARD_SERVICE_USER_NAME"`
	CreditCardServiceUserPw   string `envconfig:"CREDIT_CARD_SERVICE_USER_PW" json:"CREDIT_CARD_SERVICE_USER_PW"`

	PaymentProviders string `envconfig:"PAYMENT_PROVIDERS" json:"PAYMENT_PROVIDERS"`

	NatsClients string `envconfig:"NATS_CLIENTS" json:"NATS_CLIENTS"`

	EnablePprofPw string `envconfig:"ENABLE_PPROF_PW" json:"ENABLE_PPROF_PW"`

	GoogleApiKey     string `envconfig:"GOOGLE_API_KEY" json:"GOOGLE_API_KEY"`
	RecaptchaSiteKey string `envconfig:"RECAPTCHA_SITE_KEY" json:"RECAPTCHA_SITE_KEY"`

	OtherProviders string `envconfig:"OTHER_PROVIDERS" json:"OTHER_PROVIDERS"`
}

type NatsCredentials struct {
	// Type support "be" or "fe", determine which URL will be used for connection
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
	Seed     string `json:"seed"`
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

// example for NATS_CLIENTS structure:
// { "service-id-1": { "be": NatsCredentials, "fe": NatsCredentials }, "service-id-2": { ... } }
func (cr Credentials) GetServiceNatsCredentials(serviceId, key string) (NatsCredentials, error) {
	var res map[string]NatsCredentials
	err := getJSON("NATS_CLIENTS", cr.NatsClients, serviceId, &res)
	if err != nil {
		return NatsCredentials{}, err
	}

	creds, ok := res[key]
	if !ok {
		return creds, fmt.Errorf("did not found credentials for NATS_CLIENTS[%s][%s]", serviceId, key)
	}
	return creds, err
}

func getJSON(name, stringifyObj, key string, target interface{}) error {
	m := map[string]json.RawMessage{}
	err := json.Unmarshal([]byte(stringifyObj), &m)
	if err != nil {
		return fmt.Errorf("can't unmarshal %s. error: %v", name, err)
	}

	obj, ok := m[key]
	if !ok {
		return fmt.Errorf("key %s not found in credentials %s", key, name)
	}
	err = json.Unmarshal(obj, target)
	if err != nil {
		return fmt.Errorf("can't unmarshal obj %s for key %s to type %T. error: %v", name, key, target, err)
	}
	return nil
}
