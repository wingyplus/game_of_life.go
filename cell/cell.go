package cell

const (
    Live int = iota
    Dead
)

type Cell struct {
    state int
}

func (cell *Cell) mutate(neighbours int) {
    cell.state = Dead
}
