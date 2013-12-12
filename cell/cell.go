package cell

const (
    Live int = iota
    Dead
)

type Cell struct {
    state int
}

func (cell *Cell) mutate(neighbours int) {
    if neighbours == 2 {
        cell.state = Live
    } else {
        cell.state = Dead
    }
}
