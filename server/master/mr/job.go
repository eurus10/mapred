package mr

const (
	MAPPER  = "Map"
	REDUCER = "Reduce"
)

type Job struct {
	ID       int
	Name     string
	Type     string
	DepJobID int
	Worker   string
}
