package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type str_dict struct {
  digit_str string
  digit string
}

func handle(e error) {
  if e != nil {
    panic(e)
  }
}

func isInt(s byte) bool {
  return 48 <= s && s <= 57
}

func findDigit(s string, start int, end int, step int) int {
  for i := start; i != end; i += step {
    if isInt(s[i]) {
      return i
    }
  }
  return -1
}

func firstDigit(s string) int {
  return findDigit(s, 0, len(s), 1);
}

func lastDigit(s string) int {
  return findDigit(s, len(s) - 1, -1, -1);
}

func getDigitStr() []str_dict {
  needles := []str_dict {
    str_dict {
      digit_str: "zero",
      digit: "0",
    },
    str_dict {
      digit_str: "one",
      digit: "1",
    },
    str_dict {
      digit_str: "two",
      digit: "2",
    },
    str_dict {
      digit_str: "three",
      digit: "3",
    },
    str_dict {
      digit_str: "four",
      digit: "4",
    },
    str_dict {
      digit_str: "five",
      digit: "5",
    },
    str_dict {
      digit_str: "six",
      digit: "6",
    },
    str_dict {
      digit_str: "seven",
      digit: "7",
    },
    str_dict {
      digit_str: "eight",
      digit: "8",
    },
    str_dict {
      digit_str: "nine",
      digit: "9",
    },
  }

  return needles
}

func firstDigitStr(s string) (string, int) {
  needles := getDigitStr()
  first_at := len(s)
  first_digit := ""
  for i := 0; i < len(needles); i++ {
    substring_at := strings.Index(s, needles[i].digit_str)
    if substring_at > -1 && substring_at < first_at {
      first_at = substring_at
      first_digit = needles[i].digit
    }
  }

  return first_digit, first_at
}

func lastDigitStr(s string) (string, int) {
  needles := getDigitStr()
  last_at := 0
  last_digit := ""
  for i := 0; i < len(needles); i++ {
    substring_at := strings.LastIndex(s, needles[i].digit_str)
    if substring_at > -1 && substring_at > last_at {
      last_at = substring_at
      last_digit = needles[i].digit
    }
  }

  return last_digit, last_at
}

func do_part(part2 bool) {
  file, err := os.Open("input.txt")
  handle(err)
  defer file.Close()

  scanner := bufio.NewScanner(file)

  sum := 0
  for scanner.Scan() {
    line := scanner.Text()
    value_str := make([]byte, 2)

    first_digit := firstDigit(line)
    last_digit := lastDigit(line)

    first_digit_str, first_digit_str_at := firstDigitStr(line)
    last_digit_str, last_digit_str_at := lastDigitStr(line)

    value_str[0] = line[first_digit]
    if part2 && first_digit_str_at != -1 && first_digit_str_at < first_digit {
      value_str[0] = first_digit_str[0]
    }

    value_str[1] = line[last_digit]
    if part2 && last_digit_str_at != -1 && last_digit_str_at > last_digit {
      value_str[1] = last_digit_str[0]
    }

    value, err := strconv.Atoi(string(value_str))
    handle(err)
    sum += value
  }

  fmt.Printf("Part 1: %d\n", sum)
}

func main() {
  do_part(false)
  do_part(true)
}

