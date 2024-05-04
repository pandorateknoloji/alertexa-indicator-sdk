package indicatorsdk

// Invoke implements IIndicator.
func (i *indicator) EMA(prices []float64, length int) float64 {
	var value float64

	multiplier := 2.0 / float64(length+1)

	sma := sma(prices[:length], length)
	value = sma

	for i := length; i < len(prices); i++ {
		value = prices[i]*multiplier + value*(1-multiplier)
	}

	return value
}

func sma(prices []float64, period int) float64 {
	var sum float64
	for i := 0; i < period; i++ {
		sum += prices[i]
	}
	return sum / float64(period)
}
