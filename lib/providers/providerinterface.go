package providers

type ProviderInterface interface {
	Find()
	List()
	Insert()
	Update()
	Delete()
}
