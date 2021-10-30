package domain

type Journey [][]string

func (j *Journey) Initialize(travelTime *int) {
	*j = make([][]string, *travelTime)
	for i := range *j {
		(*j)[i] = make([]string, 0, 10)
	}
}
