<h1 align='center'>nanoid</h1>
<p align="center">
  <img src="https://github.com/noquark/nanoid/actions/workflows/test.yml/badge.svg" alt="CI Status">
  <a href="https://pkg.go.dev/github.com/noquark/nanoid"><img src="https://pkg.go.dev/badge/github.com/noquark/nanoid.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/badge/github.com/noquark/nanoid"><img src="https://goreportcard.com/badge/github.com/noquark/nanoid" alt="Go Report Card"></a>
  <a href="https://github.com/noquark/nanoid/blob/master/LICENSE"><img src="https://img.shields.io/github/license/noquark/nanoid" alt="license MIT"></a>
</p>
<p align="center">
  <b>A NanoID generator in Go</b><br/>
  <sub>Based on <a href="https://github.com/ai/nanoid">nanoid</a> and <a href="https://github.com/matoous/go-nanoid">go-nanoid</a></sub>
</p>
<br />

## Install

```bash
$ go get github.com/noquark/nanoid
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

### Credits

- [ai/nanoid](https://github.com/ai/nanoid) for original nanoid
- [matoous/go-nanoid](https://github.com/matoous/go-nanoid) for tests

### Alternatives

- [matoous/go-nanoid](https://github.com/matoous/go-nanoid)
- [bsm/nanoid](https://github.com/bsm/nanoid)
- [aidarkhanov/nanoid](https://github.com/aidarkhanov/nanoid)
- [jkomyno/nanoid](https://github.com/jkomyno/nanoid)
