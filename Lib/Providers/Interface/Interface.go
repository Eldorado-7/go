package Interface

type ProviderInterface interface {
	Connect() bool
	Query(params map[string]string)
}
