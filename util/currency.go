package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
)

// IsSupportedCurrency returns true if the currency is supported
//this function is used to replace "one of" with "currency"
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR:
		return true
	}
	return false
}
