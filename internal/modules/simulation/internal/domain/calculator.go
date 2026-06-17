package domain

import "math"

type Calculator struct {
	input SimulationInput
}

func NewCalculator(input SimulationInput) *Calculator{
	return &Calculator{
		input: input,
	}
}

func (c *Calculator) Calculate() []SimulationResultRow{
	rowsByYear := c.calculateRowsByYear()

	return c.buildRows(rowsByYear)
}

func (c *Calculator) calculateRowsByYear() map[int]SimulationResultRow {
	amount := float64(c.input.CurrentAmount())
	monthlyRate := c.input.MonthlyRate()
	maxYear := maxInt(c.input.ProjectionYears)

	rowsByYear := make(map[int]SimulationResultRow, maxYear)

	for month := 1; month <= maxYear*12; month++ {
		amount = calculateNextMonthAmount(
			amount,
			monthlyRate,
			c.input.MonthlyContribution,
		)

		if isEndOfYear(month) {
			year := monthToYear(month)
			rowsByYear[year] = c.buildResultRow(year, month, amount)
		}
	}

	return rowsByYear
}


// buildRows は、指定された出力対象年数の順序に合わせて結果を並べます。
func (c *Calculator) buildRows(rowsByYear map[int]SimulationResultRow) []SimulationResultRow {
	rows := make([]SimulationResultRow, 0, len(c.input.ProjectionYears))

	for _, year := range c.input.ProjectionYears {
		rows = append(rows, rowsByYear[year])
	}

	return rows
}

// buildResultRow は、指定年数時点のシミュレーション結果を生成します。
func (c *Calculator) buildResultRow(
	year int,
	elapsedMonths int,
	amount float64,
) SimulationResultRow {
	totalAmount := int64(math.Round(amount))
	principal := c.input.PrincipalAmount + c.input.MonthlyContribution*int64(elapsedMonths)

	return SimulationResultRow{
		Year:        year,
		TotalAmount: totalAmount,
		Principal:   principal,
		Profit:      totalAmount - principal,
	}
}

// calculateNextMonthAmount は、1ヶ月後の資産額を計算します。
// MVPでは、月次複利で運用した後、月末に積立額を追加します。
func calculateNextMonthAmount(
	currentAmount float64,
	monthlyRate float64,
	monthlyContribution int64,
) float64 {
	return currentAmount*(1+monthlyRate) + float64(monthlyContribution)
}

func isEndOfYear(month int) bool {
	return month%12 == 0
}

func monthToYear(month int) int {
	return month / 12
}

func maxInt(values []int) int {
	maxValue := values[0]

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}