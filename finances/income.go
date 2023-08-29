package finances


func TimeToBreakEven(month, flatRate, percentRate, monthlyRevenue float32) {
	var total float32 = 0.00

	for month = 1.00; month < 24.00; month++ {
		total += (monthlyRevenue*percentRate)
		if total > flatRate {
			return month
		} else {
			fmt.Printf("Month: %f\r\nTotal: %f\r\n\n", month, total)
		}
	}
}


