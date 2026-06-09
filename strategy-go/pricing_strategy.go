package main

import "fmt"

type PricingStrategy interface {
	CalculatePrice(estimatedHours float64, hourlyRate float64) float64
}
type FixedPriceStrategy struct{}

func (f *FixedPriceStrategy) CalculatePrice(estimatedHours float64, hourlyRate float64) float64 {
	basePrice := estimatedHours * hourlyRate
	riskMargin := basePrice * 0.20
	fmt.Println("Aplicando Estratégia de Preço Fixo (com margem de risco)...")
	return basePrice + riskMargin
}

type HourlyStrategy struct{}

func (h *HourlyStrategy) CalculatePrice(estimatedHours float64, hourlyRate float64) float64 {
	fmt.Println("Aplicando Estratégia de Preço por Hora Trabalhada...")
	return estimatedHours * hourlyRate
}