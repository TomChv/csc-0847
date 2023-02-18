package db

import "fmt"

type ProviderFunc = func() (*Client, error)

var providers = map[Provider]ProviderFunc{
	Local:   local,
	GCloud:  gcloud,
	Unknown: nil,
}

type Provider string

const (
	Local   Provider = "local"
	GCloud           = "gcloud"
	Unknown          = "unknown"
)

func (p Provider) String() string {
	switch p {
	case Local:
		return "local"
	case GCloud:
		return "gcloud"
	default:
		return "unknown"
	}
}

// stringToProvider converts a simple string into a
// Provider
func stringToProvider(provider string) (Provider, error) {
	switch provider {
	case "local":
		return Local, nil
	case "gcloud":
		return GCloud, nil
	default:
		return Unknown, fmt.Errorf("%s is not a valid provider", provider)
	}
}
