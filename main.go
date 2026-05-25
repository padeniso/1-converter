package main

import "fmt"

func userData() (curType string, curAmount float64) {
	fmt.Println("Введите исходный тип валюты (USD, EUR, RUB):")
	fmt.Scan(&curType)
	fmt.Println("Введите сумму:")
	fmt.Scan(&curAmount)
	return curType, curAmount
}

func convertCurrency(fromCur string, toCur string, curAmount float64) (conversionResult, float64) {

}

func main() {
	const usdToEur = 0.85
	const usdToRub = 75.0
	const eurToRub = usdToEur * usdToRub
}
