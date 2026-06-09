package main
import "fmt"

func main() {
	proposal := CommercialProposal{
		ClientName:     "Empresa Júnior Cliente",
		EstimatedHours: 150.0, // Exemplo: 150 horas estimadas
		HourlyRate:     80.0,  // Exemplo: R$ 80/hora
	}

	fmt.Printf("--- Gerando Proposta para: %s ---\n", proposal.ClientName)

	// 2. O cliente prefere fechar um pacote com Preço Fixo
	fixedStrategy := &FixedPriceStrategy{}
	proposal.SetPricingStrategy(fixedStrategy)
	
	totalFixed := proposal.CalculateTotal()
	fmt.Printf("Total do Orçamento: R$ %.2f\n\n", totalFixed)

	// 3. Mas depois decidem mudar o modelo de contrato para "Por Hora"
	// Mudamos a estratégia em tempo de execução, sem alterar os dados da proposta!
	hourlyStrategy := &HourlyStrategy{}
	proposal.SetPricingStrategy(hourlyStrategy)
	
	totalHourly := proposal.CalculateTotal()
	fmt.Printf("Total do Orçamento: R$ %.2f\n", totalHourly)
}