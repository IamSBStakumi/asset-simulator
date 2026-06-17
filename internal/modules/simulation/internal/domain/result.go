package domain

type SimulationResult struct {
	Rows []SimulationResultRow
}

type SimulationResultRow struct {
	Year int
	TotalAmount int64
	Principal int64
	Profit	int64
}