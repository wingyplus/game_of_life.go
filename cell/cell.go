package cell

const (
    Live int = iota
    Dead
)

type Cell struct {
    state int
}

func (cell *Cell) Mutate(neighbours int) {
    if neighbours == 2 {
        cell.state = Live
    } else {
        cell.state = Dead
    }
}
