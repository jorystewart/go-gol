package main

import (
  "flag"
  "fmt"
  tm "github.com/buger/goterm"
  "math/rand/v2"
  "time"
)

func main() {

  xArg := flag.Int("x", 0, "x size of matrix")
  yArg := flag.Int("y", 0, "y size of matrix")
  flag.Parse()
  if (*xArg < 5) || (*yArg < 5) {
    fmt.Println("x and y must be specified and greater than 5")
    return
  }

  generationCount := 1
  // fmt.Printf("x: %d, y: %d\n", *xArg, *yArg)
  matrix := initializeMatrix(*xArg, *yArg)

  tickTimer := time.NewTimer(time.Second)
  tm.Clear()
  tm.MoveCursor(1, 1)
  drawMatrix(matrix)
  tm.Println("Generation: ", generationCount)
  tm.Flush()
  // newMatrix := matrix

  for {
    tm.MoveCursor(1, 1)

    // newMatrix = tick(matrix)
    matrix = tick(matrix)
    // drawMatrix(newMatrix)
    drawMatrix(matrix)
    generationCount++
    tm.Println("Generation: ", generationCount)
    <-tickTimer.C
    tm.Flush()
    tickTimer.Reset(time.Second)
  }
}

func initializeMatrix(x int, y int) *[][]bool {
  matrix := make([][]bool, y)
  for i := range matrix {
    newRow := make([]bool, x)
    for j := range newRow {
      initialState := rand.IntN(2)
      if initialState == 0 {
        newRow[j] = false
      } else {
        newRow[j] = true
      }
    }
    matrix[i] = newRow
  }
  return &matrix
}

func drawMatrix(matrix *[][]bool) {
  for _, row := range *matrix {
    for _, cell := range row {
      switch cell {
      case true:
        tm.Printf("\u25A0 ")
      default:
        tm.Printf("  ")
      }
    }
    tm.Println()
  }
}

func countNeighbours(matrix *[][]bool, x int, y int) int8 {
  var neighbours int8 = 0
  if y-1 > 0 {
    row := (*matrix)[y-1]
    if x-1 > 0 {
      if row[x-1] == true {
        neighbours++
      }
      if row[x] == true {
        neighbours++
      }
      if x+1 < len(row) {
        if row[x+1] == true {
          neighbours++
        }
      }
    }
  }

  row := (*matrix)[y]
  if x-1 > 0 {
    if row[x-1] == true {
      neighbours++
    }
  }
  if x+1 < len(row) {
    if row[x+1] == true {
      neighbours++
    }
  }

  if y+1 < len(*matrix) {
    row := (*matrix)[y+1]
    if x-1 > 0 {
      if row[x-1] == true {
        neighbours++
      }
    }
    if row[x] == true {
      neighbours++
    }
    if x+1 < len(row) {
      if row[x+1] == true {
        neighbours++
      }
    }
  }

  return neighbours
}

// if <2 neighbours, die
// if 2-3 neighbours, no state change
// if >3 neighbours, die

// if 3 neighbours and dead, live

func tick(matrix *[][]bool) *[][]bool {
  newMatrix := make([][]bool, len(*matrix))
  for i := range newMatrix {
    newMatrix[i] = make([]bool, len((*matrix)[i]))
    copy(newMatrix[i], (*matrix)[i])
  }

  for colNum, row := range *matrix {
    for rowNum, cell := range row {
      if cell == true {
        if countNeighbours(matrix, rowNum, colNum) < 2 {
          newMatrix[colNum][rowNum] = false
        } else if countNeighbours(matrix, rowNum, colNum) > 3 {
          newMatrix[colNum][rowNum] = false
        }
      } else {
        if countNeighbours(matrix, rowNum, colNum) == 3 {
          newMatrix[colNum][rowNum] = true
        }
      }
    }
  }

  return &newMatrix
}
