package converter

// Service interface for working with coin convertion
type Service interface {
	Convert(amount, convertFrom, convertTo string) (float64, error)
}
