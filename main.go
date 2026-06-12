package main

import "fmt"

func userData() (fromCur string, toCur string, curAmount float64) {
	for {
		fmt.Println("Введите исходный тип валюты (USD, EUR, RUB):")
		fmt.Scan(&fromCur)
		if fromCur != "USD" && fromCur != "EUR" && fromCur != "RUB" {
			fmt.Println("Неверный тип валюты. Пожалуйста, введите USD, EUR или RUB.")
		} else {
			break
		}
	}
	for {
		fmt.Println("Введите сумму:")
		fmt.Scan(&curAmount)
		if curAmount <= 0 {
			fmt.Println("Неверная сумма. Пожалуйста, введите положительное число.")
		} else {
			break
		}
	}
	for {
		fmt.Println("Введите целевой тип валюты (USD, EUR, RUB):")
		fmt.Scan(&toCur)
		if toCur != "USD" && toCur != "EUR" && toCur != "RUB" {
			fmt.Println("Неверный тип валюты. Пожалуйста, введите USD, EUR или RUB.")
			continue
		}
		if toCur == fromCur {
			fmt.Println("Целевой тип валюты должен отличаться от исходного. Пожалуйста, выберите другой тип.")
			continue
		}
		return fromCur, toCur, curAmount
	}

}

func convertCurrency(fromCur string, toCur string, curAmount float64) (conversionResult float64) {
	myCurMap := map[string]float64{"USDEUR": 0.85, "USDRUB": 75.0, "EURUSD": 1 / 0.85, "EURRUB": 75.0 / 0.85, "RUBUSD": 1 / 75.0, "RUBEUR": 0.85 / 75.0}
	conversionResult = curAmount * myCurMap[fromCur+toCur]
	// const usdToEur = 0.85
	// const usdToRub = 75.0
	// const eurToRub = usdToEur * usdToRub
	// switch fromCur {
	// case "USD":
	// 	if toCur == "EUR" {
	// 		conversionResult = curAmount * usdToEur
	// 	} else if toCur == "RUB" {
	// 		conversionResult = curAmount * usdToRub
	// 	}
	// case "EUR":
	// 	if toCur == "USD" {
	// 		conversionResult = curAmount / usdToEur
	// 	} else if toCur == "RUB" {
	// 		conversionResult = curAmount / usdToEur * usdToRub
	// 	}
	// case "RUB":
	// 	if toCur == "USD" {
	// 		conversionResult = curAmount / usdToRub
	// 	} else if toCur == "EUR" {
	// 		conversionResult = curAmount / usdToRub * usdToEur
	// 	}
	// }
	return conversionResult
}

func main() {
	fromCur, toCur, curAmount := userData()
	result := convertCurrency(fromCur, toCur, curAmount)
	fmt.Printf("%.2f %s = %.2f %s\n", curAmount, fromCur, result, toCur)
}
