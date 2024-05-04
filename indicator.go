package indicatorsdk

type IIndicator interface {
	RSI(prices []float64, length int) float64
	EMA(prices []float64, length int) float64
}

// client implements IIndicator interface
type indicator struct{}

// NewClient creates and returns moov client
func NewIndicator() IIndicator {
	cli := indicator{}
	return &cli
}
