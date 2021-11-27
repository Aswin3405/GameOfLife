package main

import (
	"fmt"
	"time"
)

type Board struct {
	Cells           [][]bool
	Length, Breadth int
}

type Coordinate struct {
	X, Y int
}

func NewBoard(length, breadth int) *Board {
	b := &Board{
		Length:  length,
		Breadth: breadth,
	}
	b.Cells = make([][]bool, length)
	for i := 0; i < length; i++ {
		b.Cells[i] = make([]bool, breadth)
	}
	return b
}

func (b *Board) Print() {
	fmt.Printf("%5v", " ")
	for i := 0; i < b.Breadth; i++ {
		fmt.Printf("%5v", i)
	}
	fmt.Println()
	for i := 0; i < b.Length; i++ {
		//fmt.Printf("%d  " ,i )
		fmt.Printf("%5v", i)
		for j := 0; j < b.Breadth; j++ {
			//	fmt.Printf("%d  ",j )
			if b.Cells[i][j] {
				fmt.Printf("%5v", "X")
			} else {
				fmt.Printf("%5v", ".")
			}
		}
		fmt.Println()
	}

	fmt.Println()

}

func (b *Board) GetNeighbours(x, y int) []Coordinate {
	var coordinates []Coordinate
	coordinates = append(coordinates, Coordinate{
		X: (b.Length + x - 1) % b.Length,
		Y: (y + 1)%b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: x % b.Length,
		Y: (y + 1) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (x + 1) % b.Length,
		Y: (y + 1) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (b.Length + x - 1) % b.Length,
		Y: (y) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (x + 1) % b.Length,
		Y: (y) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (b.Length + x - 1) % b.Length,
		Y: (b.Breadth + y - 1) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (x) % b.Length,
		Y: (b.Breadth + y - 1) % b.Breadth,
	})
	coordinates = append(coordinates, Coordinate{
		X: (x + 1) % b.Length,
		Y: (b.Breadth + y - 1) % b.Breadth,
	})
	//fmt.Printf("\n%d , %d, %v", x, y, coordinates)
	return coordinates
}

func (b *Board) GetNextStatus(x, y int) bool {
	liveNeighbourCount := b.GetLiveNeighbourCount(x, y)
	if b.IsAlive(x, y) {
		//Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
		if liveNeighbourCount < 2 {
			return false
		}
		//Any live cell with two or three live neighbors lives on to the next generation.
		if liveNeighbourCount == 2 || liveNeighbourCount == 3 {
			return b.IsAlive(x, y)
		}
		//Any live cell with more than three live neighbors dies, as if by overcrowding.
		return false
	}
	//Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction
	if liveNeighbourCount == 3 {
		return true
	}
	return false
}

func (b *Board) GetLiveNeighbourCount(x, y int) int {
	liveNeighbourCount := 0
	neighbourCoordinates := b.GetNeighbours(x, y)
	for _, coordinate := range neighbourCoordinates {
		if b.IsAlive(coordinate.X, coordinate.Y) {
			liveNeighbourCount++
		}
	}
	return liveNeighbourCount
}

func (b *Board) GetDeadNeighbourCount(x, y int) int {
	return 8 - b.GetLiveNeighbourCount(x, y)
}

func (b *Board) IsAlive(x, y int) bool {
	return b.Cells[x][y]
}

func (b *Board) Step() {
	newCells := make([][]bool, b.Length)
	for i := 0; i < b.Length; i++ {
		newCells[i] = make([]bool, b.Breadth)
	}
	for i := 0; i < b.Length; i++ {
		for j := 0; j < b.Breadth; j++ {
			newCells[i][j] = b.GetNextStatus(i, j)
		}
	}
	b.Cells = newCells
}

func main() {
	b := NewBoard(25, 25)
	//fmt.Print("\x0c", b)
	b.Cells[11][12] = true
	b.Cells[12][13] = true
	b.Cells[13][11] = true
	b.Cells[13][12] = true
	b.Cells[13][13] = true

	for i:=0; i<100; i++ {
		fmt.Printf("Generation: %d\n", i)
		b.Print()
		b.Step()
		time.Sleep(time.Second * 1)
	}
}
