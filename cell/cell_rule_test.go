package cell

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestLiveCell(t *testing.T) {
    var cell *Cell = &Cell { state: Live }

    Convey("Subject: Any live cell with fewer than two live neighbours dies", t, func() {
        Convey("Given: Initial 0 neighbour", func() {
            var neighbour = 0

            Convey("When: cell mutate by given neighbour", func() {
                cell.Mutate(neighbour)

                Convey("Then: cell is dead", func() {
                    So(cell.state, ShouldEqual, Dead)
                })
            })
        })
    })

    Convey("Subject: Any live cell with two or three live neighbours lives", t, func() {
        Convey("Given: Initial 2 neighbours", func() {
            var neighbours = 2

            Convey("When: cell mutate by given neighbours", func() {
                cell.Mutate(neighbours)

                Convey("Then: cell is live", func() {
                    So(cell.state, ShouldEqual, Live)
                })
            })
        })
        Convey("Given: Initial 3 neighbours", func() {
            var neighbours = 3

            Convey("When: cell mutate by given neighbours", func() {
                cell.Mutate(neighbours)

                Convey("Then: cell is live", func() {
                    So(cell.state, ShouldEqual, Live)
                })
            })
        })
    })
    Convey("Subject: Any live cell with more than three live neighbours dies", t, func() {
        Convey("Given: Initial 4 neighbours", func() {
            var neighbours = 4
            Convey("When: cell mutate by given neighbours", func() {
                cell.Mutate(neighbours)

                Convey("Then: cell is dead", func() {
                    So(cell.state, ShouldEqual, Dead)
                })
            })
        })
    })
}

func TestDeadCell(t *testing.T) {
    Convey("Subject: Any dead cell with exactly three live neighbours a live cell", t, func() {
        var cell *Cell = &Cell { state: Dead }

        Convey("Given: Initial 3 neighbours", func() {
            var neighbours = 3
            Convey("When: cell mutate by given neighbours", func() {
                cell.Mutate(neighbours)
                Convey("Then: cell is live", func() {
                    So(cell.state, ShouldEqual, Live)
                })
            })
        })
    })
}
