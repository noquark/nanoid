<h1 align='center'>nanoid</h1>
<p align="center">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/lint.yml/badge.svg" alt="Lint Status">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/test.yml/badge.svg" alt="Test and Benchmark Status">
  <a href="https://pkg.go.dev/github.com/mohitsinghs/nanoid"><img src="https://pkg.go.dev/badge/github.com/mohitsinghs/nanoid.svg" alt="Go Reference"></a>
</p>
<p align="center">
  <b>A NanoID generator in Go</b><br/>
  <sub>Based on <a href="https://github.com/ai/nanoid">nanoid</a> and <a href="https://github.com/matoous/go-nanoid">go-nanoid</a></sub>
</p>
<br />

## Install

```bash
$ go get github.com/mohitsinghs/nanoid
```

## Usage

For simple nanoID

```go
id, err := nanoid.New()
```

For nanoID with custom length

```go
id, err := nanoid.New(7)
```

## Benchmark

So, I quickly compared this with 2 other nanoid implementations and 3 popular uuid implementations.

```sh
BenchmarkLong/mohitsinghs/nanoid-8               2018715               603.7 ns/op            24 B/op          1 allocs/op
BenchmarkLong/bsm/nanoid-8                       1832822               657.5 ns/op            24 B/op          1 allocs/op
BenchmarkLong/matoous/go-nanoid-8                 517705              2040 ns/op             144 B/op          3 allocs/op
BenchmarkLong/satori/go.uuid-8                    596088              1750 ns/op              16 B/op          1 allocs/op
BenchmarkLong/google/uuid-8                       585169              1755 ns/op              16 B/op          1 allocs/op
BenchmarkLong/gofrs/uuid-8                        688498              1741 ns/op              16 B/op          1 allocs/op
BenchmarkShort/mohitsinghs/nanoid-8              5009698               240.6 ns/op             8 B/op          1 allocs/op
BenchmarkShort/bsm/nanoid-8                      4735372               253.7 ns/op             8 B/op          1 allocs/op
BenchmarkShort/matoous/go-nanoid-8                562224              1878 ns/op              56 B/op          3 allocs/op
```

## Notes

- Custom alphabets are not supported yet

## Credits and Alternatives

- [ai/nanoid](https://github.com/ai/nanoid)
- [matoous/go-nanoid](https://github.com/matoous/go-nanoid)
- [bsm/nanoid](https://github.com/bsm/nanoid)
