package cell

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestLiveCell(t *testing.T) {
    Convey("Subject: Any live cell with fewer than two live neighbours dies", t, func() {
        var cell *Cell = new(Cell)

        Convey("Given: Initial live cell and 0 neighbour", func() {
            cell.state = Live
            var neighbour = 0

            Convey("When: cell mutate by given neighbour", func() {
                cell.mutate(neighbour)

                Convey("Then: cell is dead", func() {
                    So(cell.state, ShouldEqual, Dead)
                })
            })
        })
    })
}

