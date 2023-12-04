package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type cubes struct {
  red int
  green int
  blue int
}

type game struct {
  index int
  rounds []cubes
}

func handle(e error) {
  if e != nil {
    panic(e)
  }
}

func possibleGame(outcomes game, available cubes) bool {
  for i := 0; i < len(outcomes.rounds); i++ {
    if outcomes.rounds[i].red > available.red ||
        outcomes.rounds[i].green > available.green ||
        outcomes.rounds[i].blue > available.blue {
      return false
    }
  }

  return true
}

func reduce(outcomes game) cubes {
  reduced := cubes {0, 0, 0}

  for i := 0; i < len(outcomes.rounds); i++ {
    if outcomes.rounds[i].red > reduced.red {
      reduced.red = outcomes.rounds[i].red
    }

    if outcomes.rounds[i].green > reduced.green {
      reduced.green = outcomes.rounds[i].green
    }

    if outcomes.rounds[i].blue > reduced.blue {
      reduced.blue = outcomes.rounds[i].blue
    }
  }

  return reduced
}

func parseLine(line string) game {
  split1 := strings.Split(line, ":")
  game_index_str := strings.Fields(split1[0])[1]
  game_index, err := strconv.Atoi(game_index_str)
  handle(err)

  rounds := []cubes {}
  split2 := strings.Split(split1[1], ";")
  for i := 0; i < len(split2); i++ {
    split3 := strings.Split(split2[i], ",")

    round := cubes {}
    for j := 0; j < len(split3); j++ {
      count_color := strings.Fields(split3[j])
      count, err := strconv.Atoi(count_color[0])
      handle(err)
      color := count_color[1]

      if color == "blue" {
        round.blue = count
      } else if color == "red" {
        round.red = count
      } else if color == "green" {
        round.green = count
      } else {
        panic("unknown color " + color)
      }
    }

    rounds = append(rounds, round)
  }

  g := game {
    game_index,
    rounds,
  }

  return g
}

func solve() {
  file, err := os.Open("input.txt")
  handle(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)

  available_part1 := cubes {
    12,
    13,
    14,
  }

  score_part1 := 0
  score_part2 := 0
  for scanner.Scan() {
    line := scanner.Text()
    g := parseLine(line)

    if possibleGame(g, available_part1) {
      score_part1 += g.index
    }

    reduced := reduce(g)
    score_part2 += reduced.red * reduced.green * reduced.blue
  }


  fmt.Println("Part 1: ", score_part1)
  fmt.Println("Part 2: ", score_part2)
}

func main() {
  solve()
}

