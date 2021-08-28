<h1 align='center'>nanoid</h1>
<p align="center">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/lint.yml/badge.svg" alt="Lint Status">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/test.yml/badge.svg" alt="Test and Benchmark Status">
  <a href="https://pkg.go.dev/github.com/mohitsinghs/nanoid"><img src="https://pkg.go.dev/badge/github.com/mohitsinghs/nanoid.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/badge/github.com/mohitsinghs/nanoid"><img src="https://goreportcard.com/badge/github.com/mohitsinghs/nanoid" alt="Go Report Card"></a>
  <a href="https://github.com/mohitsinghs/nanoid/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mohitsinghs/nanoid" alt="license MIT"></a>
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

### Benchmark

So, I quickly compared this with other implementations.

```sh
BenchmarkLong/matoous/nanoid-8            590540              2072 ns/op             144 B/op          3 allocs/op
BenchmarkLong/aidarkhanov/nanoid-8        329972              3424 ns/op              72 B/op          2 allocs/op
BenchmarkLong/jkomyno/nanoid-8            641336              1859 ns/op              72 B/op          3 allocs/op
BenchmarkLong/bsm/nanoid-8               1824750               661.6 ns/op            24 B/op          1 allocs/op
BenchmarkLong/mohitsinghs/nanoid-8       1951998               617.4 ns/op            24 B/op          1 allocs/op
BenchmarkShort/matoous/nanoid-8           544934              1896 ns/op              56 B/op          3 allocs/op
BenchmarkShort/aidarkhanov/nanoid-8       563295              1801 ns/op              24 B/op          2 allocs/op
BenchmarkShort/jkomyno/nanoid-8           671174              1802 ns/op              24 B/op          3 allocs/op
BenchmarkShort/bsm/nanoid-8              4687411               252.0 ns/op             8 B/op          1 allocs/op
BenchmarkShort/mohitsinghs/nanoid-8      5015697               244.2 ns/op             8 B/op          1 allocs/op
```

### Notes

- Custom alphabets are not supported yet

### Credits

- [ai/nanoid](https://github.com/ai/nanoid)
- [matoous/go-nanoid](https://github.com/matoous/go-nanoid)

### Alternatives

- [matoous/go-nanoid](https://github.com/matoous/go-nanoid)
- [bsm/nanoid](https://github.com/bsm/nanoid)
- [aidarkhanov/nanoid](https://github.com/aidarkhanov/nanoid)
- [jkomyno/nanoid](https://github.com/jkomyno/nanoid)
