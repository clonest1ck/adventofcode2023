package main

import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

type card struct {
  index int
  winners []int
  numbers []int
  wins int
  copies int
}

func handle(e error) {
  if e != nil {
    panic(e)
  }
}

func strToIntSlice(s []string) []int {
  values := []int {}

  for i := 0; i < len(s); i++ {
    val, err := strconv.Atoi(s[i])
    handle(err)
    values = append(values, val)
  }

  return values
}

func parseLine(line string) card {
  split1 := strings.Split(line, ":")
  card_index_str := strings.Fields(split1[0])[1]
  split2 := strings.Split(split1[1], "|")
  winners_str := strings.Fields(split2[0])
  numbers_str := strings.Fields(split2[1])

  card_index, err := strconv.Atoi(card_index_str)
  handle(err)

  winners := strToIntSlice(winners_str)
  numbers := strToIntSlice(numbers_str)

  c := card {
    card_index,
    winners,
    numbers,
    0,
    1,
  }

  return c
}

func calculateWins(c card) int {
  wins := 0

  for i := 0; i < len(c.winners); i++ {
    for j := 0; j < len(c.numbers); j++ {
      if c.winners[i] == c.numbers[j] {
        wins++
        break
      }
    }
  }

  return wins
}

func solve() {
  file, err := os.Open("input.txt")
  handle(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)
  cards := []card {}

  for scanner.Scan() {
    line := scanner.Text()
    c := parseLine(line)
    c.wins = calculateWins(c)
    cards = append(cards, c)
  }

  score_part1 := 0
  score_part2 := 0
  for i := 0; i < len(cards); i++ {
    wins := cards[i].wins
    score_part2 += cards[i].copies
    if (wins > 0) {
      score_part1 += int(math.Pow(2, float64(wins - 1)))
    }

    for j := 1; j <= wins; j++ {
      cards[i + j].copies += cards[i].copies
    }
  }

  fmt.Println("Part 1: ", score_part1)
  fmt.Println("Part 2: ", score_part2)
}

func main() {
  solve()
}

