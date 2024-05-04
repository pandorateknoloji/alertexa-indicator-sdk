package indicatorsdk

// Invoke implements IIndicator.
func (i *indicator) RSI(prices []float64, length int) float64 {
	gain, loss := 0.0, 0.0

	for i := 1; i <= length; i++ {
		change := prices[i] - prices[i-1]
		if change >= 0 {
			gain += change
		} else {
			loss -= change
		}
	}

	avgGain := gain / float64(length)
	avgLoss := loss / float64(length)

	rs := avgGain / avgLoss
	value := 100 - (100 / (1 + rs))

	for i := length + 1; i < len(prices); i++ {
		change := prices[i] - prices[i-1]
		if change >= 0 {
			avgGain = (avgGain*(float64(length)-1) + change) / float64(length)
			avgLoss = avgLoss * (float64(length) - 1) / float64(length)
		} else {
			avgLoss = (avgLoss*(float64(length)-1) - change) / float64(length)
			avgGain = avgGain * (float64(length) - 1) / float64(length)
		}

		rs = avgGain / avgLoss
		value = 100 - (100 / (1 + rs))
	}

	return value
}
