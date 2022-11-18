package utils

// Constants for all supported currencies
const (
	USD = "USD"
	RUB = "RUB"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, RUB:
		return true
	}
	return false
}
