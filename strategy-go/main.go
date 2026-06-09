package main
import "fmt"

func main() {
	proposal := CommercialProposal{
		ClientName:     "Empresa Júnior Cliente",
		EstimatedHours: 150.0, 
		HourlyRate:     80.0,  
	}

	fmt.Printf("--- Gerando Proposta para: %s ---\n", proposal.ClientName)

	
	fixedStrategy := &FixedPriceStrategy{}
	proposal.SetPricingStrategy(fixedStrategy)
	
	totalFixed := proposal.CalculateTotal()
	fmt.Printf("Total do Orçamento: R$ %.2f\n\n", totalFixed)


	hourlyStrategy := &HourlyStrategy{}
	proposal.SetPricingStrategy(hourlyStrategy)
	
	totalHourly := proposal.CalculateTotal()
	fmt.Printf("Total do Orçamento: R$ %.2f\n", totalHourly)
}