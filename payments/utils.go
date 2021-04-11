package payments

func truncateFloat(k float64) float64 {
	return float64(int(k*100)) / 100
}
