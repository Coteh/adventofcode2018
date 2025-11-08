# adventofcode2018

[![Test AoC 2018](https://github.com/Coteh/adventofcode2018/actions/workflows/run_aoc_test.yml/badge.svg)](https://github.com/Coteh/adventofcode2018/actions/workflows/run_aoc_test.yml)

This is the first year I did Advent of Code. I used Go for this challenge, as I was learning Go at the time.

## Instructions

To compile and run:

```sh
go run 03/03.go

# pass in input file directly

go run 03/03.go < 03/input
```

Alternatively, the programs can be compiled then run directly:

```sh
go build -o 03/03 03/03.go
```

Then once it's compiled, can simply run using:

```sh
./03/03

# pass in input file directly

./03/03 < 03/input
```

NOTE: Days 1 and 2 are each split into 2 sub-modules, ending in either `-1` or `-2` (for parts 1 and 2 respectively)

## Progress

| Day  | Part 1 | Part 2 |
|------|--------|--------|
|  1   |   ✅   |   ✅   |
|  2   |   ✅   |   ✅   |
|  3   |   ✅   |   ✅   |
|  4   |   ✅   |   ✅   |
|  5   |   ✅   |   ✅   |
|  6   |   ✅   |   ✅   |
|  7   |   ✅   |   ✅   |
|  8   |   ✅   |   ✅   |
|  9   |   ✅   |   ✅   |
|  10  |   ✅*  |   ✅   |
|  11  |   ✅   |   ✅** |
|  12  |   ✅   |        |
|  13  |   ✅   |        |
|  14  |   ✅   |   ✅   |
|  15  |        |        |
|  16  |   ✅   |        |
|  17  |        |        |
|  18  |        |        |
|  19  |        |        |
|  20  |        |        |
|  21  |        |        |
|  22  |        |        |
|  23  |        |        |
|  24  |        |        |
|  25  |        |        |

\* Output has to be verified manually at the moment. See TODO comment in `./10/test.sh` for more info.

\** Day 11 part 2 is incredibly slow at the moment. Will need to see if it can be optimized later.
