<h1 align='center'>nanoid</h1>
<p align="center">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/lint.yml/badge.svg" alt="Lint Status">
  <img src="https://github.com/mohitsinghs/nanoid/actions/workflows/test.yml/badge.svg" alt="Test and Benchmark Status">
  <a href="https://pkg.go.dev/github.com/mohitsinghs/nanoid"><img src="https://pkg.go.dev/badge/github.com/mohitsinghs/nanoid.svg" alt="Go Reference"></a>
</p>
<p align="center">
  <b>Yet another nanoID in Go</b><br/>
  <sub>Derived from <a href="https://github.com/ai/nanoid">nanoid</a> and <a href="https://github.com/matoous/go-nanoid">go-nanoid</a> but faster</sub>
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

## Notes

- Custom alphabets are not supported yet

## Credits

- [ai](https://github.com/ai) : [nanoid](https://github.com/ai/nanoid)
- [matoous](https://github.com/matoous) : [go-nanoid](https://github.com/matoous/go-nanoid)
- Did I forget anyone ?
