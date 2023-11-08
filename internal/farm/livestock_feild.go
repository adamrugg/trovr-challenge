package farm

type LiveStockField[livestock comparable] interface {
	Visit()
	FindLivestock(id int) (livestock, error)
	Name() string
}
