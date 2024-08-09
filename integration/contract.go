package integration

type InfoTestIntegration interface {
	GetInfo() (interface{}, error)
}
