package main

import (
	"fmt"
	"math"
	"time"
)

type EUR float64
type Percent float64

func (me Percent) Value() float64 {
	return float64(me) / 100
}

func (me EUR) String() string {
	return fmt.Sprintf("%10.2f EUR", float64(me))
}

func main() {

	investment := EUR(1.0)

	//totalExpenseRatio := Percent(0.37)
	//orderFee := 1.5
	returnRate := Percent(1)

	purchaseDate := time.Date(2019, time.November, 1, 9, 0, 0, 0, time.Local)
	now := time.Now()

	purchaseYear := purchaseDate.Year()
	purchaseMonth := purchaseDate.Month()

	nowYear := now.Year()
	nowMonth := now.Month()

	fullYears := nowYear - purchaseYear - 1

	purchaseYearReturn := monthlyCompoundInterestStarting(investment, returnRate, purchaseMonth)
	fullYearsReturn := compoundInterest(purchaseYearReturn, returnRate, fullYears)
	totalReturn := monthlyCompoundInterestUntil(fullYearsReturn, returnRate, nowMonth)

	fmt.Printf("Investment:\t\t %s\n", investment.String())
	fmt.Printf("Purchase Year:\t %s\n", purchaseYearReturn.String())
	fmt.Printf("Full Years:\t\t %s\n", fullYearsReturn.String())
	fmt.Printf("Total:\t\t\t %s\n", totalReturn.String())
}

func compoundInterest(initialInvestment EUR, returnRate Percent, years int) EUR {
	return EUR(float64(initialInvestment) * math.Pow(1+returnRate.Value(), float64(years)))
}

func monthlyCompoundInterestStarting(initialInvestment EUR, returnRate Percent, investmentStartMonth time.Month) EUR {
	return monthlyCompoundInterest(initialInvestment, returnRate, 1-yearShare(investmentStartMonth-1))
}

func monthlyCompoundInterestUntil(initialInvestment EUR, returnRate Percent, investmentEndMonth time.Month) EUR {
	return monthlyCompoundInterest(initialInvestment, returnRate, float64(investmentEndMonth)/12)
}

func monthlyCompoundInterest(initialInvestment EUR, returnRate Percent, share float64) EUR {
	return EUR(float64(initialInvestment) * (1 + returnRate.Value()*share))
}

func yearShare(m time.Month) float64 {
	if m == time.January {
		return 1
	}
	return float64(m) / 12
}
