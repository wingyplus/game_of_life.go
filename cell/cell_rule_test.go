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

    Convey("Subject: Any live cell with two or three live neighbours lives", t, func() {
        var cell *Cell = new(Cell)

        Convey("Given: Initial live cell and 2 neighbours", func() {
            cell.state = Live
            var neighbours = 2

            Convey("When: cell mutate by given neighbours", func() {
                cell.mutate(neighbours)

                Convey("Then: cell is live", func() {
                    So(cell.state, ShouldEqual, Live)
                })
            })
        })
    })
}

