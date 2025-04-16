package service

import (
	"PeoplePilot/backend/model"
	"time"
)

// CostService verwaltet Berechnungen und Abfragen zu Personalkosten
type CostService struct{}

// NewCostService erstellt einen neuen CostService
func NewCostService() *CostService {
	return &CostService{}
}

// CalculateMonthlyLaborCosts berechnet die monatlichen Personalkosten basierend auf aktiven Mitarbeitern
func (s *CostService) CalculateMonthlyLaborCosts(employees []*model.Employee) float64 {
	var total float64
	for _, emp := range employees {
		if emp.Status == model.EmployeeStatusActive || emp.Status == model.EmployeeStatusRemote || emp.Status == model.EmployeeStatusOnLeave {
			total += emp.Salary
		}
	}
	return total
}

// CalculateAvgCostPerEmployee berechnet die durchschnittlichen Kosten pro Mitarbeiter
func (s *CostService) CalculateAvgCostPerEmployee(totalCost float64, employeeCount int) float64 {
	if employeeCount <= 0 {
		return 0
	}
	return totalCost / float64(employeeCount)
}

// GenerateMonthlyLaborCostsTrend erzeugt Trenddaten für monatliche Personalkosten
func (s *CostService) GenerateMonthlyLaborCostsTrend(currentCosts float64) []float64 {
	// Generiert historische Monatsdaten basierend auf dem aktuellen Wert
	// mit leichter Variation, um einen realistischen Trend zu simulieren
	monthlyTrend := []float64{
		currentCosts * 0.95, // Jan
		currentCosts * 0.96, // Feb
		currentCosts * 0.98, // März
		currentCosts * 0.99, // April
		currentCosts * 0.99, // Mai
		currentCosts * 1.00, // Juni
		currentCosts * 1.01, // Juli
		currentCosts * 1.02, // August
		currentCosts * 1.02, // Sept
		currentCosts * 1.03, // Okt
		currentCosts * 1.04, // Nov
		currentCosts * 1.05, // Dez
	}

	return monthlyTrend
}

// CountEmployeesByDepartment zählt Mitarbeiter pro Abteilung
func (s *CostService) CountEmployeesByDepartment(employees []*model.Employee) ([]string, []int) {
	// Zähle Mitarbeiter pro Abteilung
	departmentCount := make(map[string]int)
	for _, emp := range employees {
		departmentCount[string(emp.Department)]++
	}

	// Konvertiere in Listen für das Chart
	var labels []string
	var data []int

	for dept, count := range departmentCount {
		labels = append(labels, dept)
		data = append(data, count)
	}

	// Verwende Standardwerte, wenn keine Daten vorhanden sind
	if len(labels) == 0 {
		labels = []string{"IT", "Vertrieb", "HR", "Marketing", "Finanzen", "Produktion"}
		data = []int{12, 8, 3, 5, 6, 10}
	}

	return labels, data
}

// GenerateExpectedReviews berechnet und generiert anstehende Mitarbeitergespräche
func (s *CostService) GenerateExpectedReviews(employees []*model.Employee) []map[string]string {
	// In einer echten Anwendung würden wir hier Daten aus einer Datenbank abfragen
	// Für dieses Beispiel erstellen wir simulierte Daten

	reviews := []map[string]string{
		{
			"EmployeeName": "Max Mustermann",
			"ReviewType":   "Leistungsbeurteilung",
			"Date":         time.Now().AddDate(0, 0, 5).Format("02.01.2006"),
		},
		{
			"EmployeeName": "Erika Musterfrau",
			"ReviewType":   "Beförderungsgespräch",
			"Date":         time.Now().AddDate(0, 0, 9).Format("02.01.2006"),
		},
		{
			"EmployeeName": "John Doe",
			"ReviewType":   "Einarbeitung",
			"Date":         time.Now().AddDate(0, 0, 12).Format("02.01.2006"),
		},
	}

	// Wenn wir echte Mitarbeiterdaten haben, könnten wir Namen verwenden
	if len(employees) >= 3 {
		reviews[0]["EmployeeName"] = employees[0].FirstName + " " + employees[0].LastName
		reviews[1]["EmployeeName"] = employees[1].FirstName + " " + employees[1].LastName
		reviews[2]["EmployeeName"] = employees[2].FirstName + " " + employees[2].LastName
	}

	return reviews
}
