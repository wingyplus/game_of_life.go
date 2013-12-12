package cell

const (
    Live int = iota
    Dead
)

type Cell struct {
    state int
}

func (cell *Cell) Mutate(neighbours int) {
    switch(neighbours) {
    case 2, 3:
        cell.state = Live
    default:
        cell.state = Dead
    }
}
