package cases

type Case struct {
	ID           string
	Province     string
	Gender       string
	Neighborhood string
	Age          int
	Stage        string
	Dead         string
}

type CaseSearch struct {
	Province     *string
	Age          *int
	Stage        *string
	Dead         *string
	Neighborhood *string
	Gender       *string
}
