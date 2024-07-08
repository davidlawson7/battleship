package game

type Direction int

const (
    HORIZONTAL Direction = 0
    VERTICAL Direction = 1
)

type Battleship struct {
    Col int
    Row int
    Size int
    Dir Direction
    Hits []bool
}

func (b *Battleship) isDestroyed() bool {
    for _, hit := range b.Hits {
        if !hit {
          return false  
        }
    }

    return true
}

type Grid struct {
    Size int
    Grid []bool
}

func NewGrid(size int) Grid {
    return Grid{
        Size: size,
        Grid: make([]bool, size * size),
    }
}

