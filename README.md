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

```sh
name                         time/op
NanoID/mohitsinghs/nanoid-8   611ns ± 0%
NanoID/bsm/nanoid-8           646ns ± 0%
NanoID/jkomyno/nanoid-8      1.85µs ± 0%
NanoID/aidarkhanov/nanoid-8  3.39µs ± 0%
NanoID/matoous/go-nanoid-8   2.01µs ± 0%

name                         alloc/op
NanoID/mohitsinghs/nanoid-8   24.0B ± 0%
NanoID/bsm/nanoid-8           24.0B ± 0%
NanoID/jkomyno/nanoid-8       72.0B ± 0%
NanoID/aidarkhanov/nanoid-8   72.0B ± 0%
NanoID/matoous/go-nanoid-8     144B ± 0%

name                         allocs/op
NanoID/mohitsinghs/nanoid-8    1.00 ± 0%
NanoID/bsm/nanoid-8            1.00 ± 0%
NanoID/jkomyno/nanoid-8        3.00 ± 0%
NanoID/aidarkhanov/nanoid-8    2.00 ± 0%
NanoID/matoous/go-nanoid-8     3.00 ± 0%

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
