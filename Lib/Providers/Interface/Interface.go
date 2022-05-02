package Interface

type ProviderInterface interface {
	Find()
	List()
	Insert() bool
	Update() bool
	Delete() bool
	Connect() bool
}
