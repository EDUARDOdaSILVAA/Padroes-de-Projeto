package main

import "fmt"

type CommercialProposal struct {
	ClientName     string
	EstimatedHours float64
	HourlyRate     float64
	Strategy       PricingStrategy 
}


func (p *CommercialProposal) SetPricingStrategy(strategy PricingStrategy) {
	p.Strategy = strategy
}


func (p *CommercialProposal) CalculateTotal() float64 {
	if p.Strategy == nil {
		fmt.Println("Erro: Nenhuma estratégia de precificação definida!")
		return 0.0
	}
	return p.Strategy.CalculatePrice(p.EstimatedHours, p.HourlyRate)
}