package common

type Provider interface {
	GetName() string
	Request() (GasPrice, error)
}
