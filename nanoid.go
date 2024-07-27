package nanoid

import (
	"bufio"
	"crypto/rand"
	"errors"
	"io"
	"sync"
	"unsafe"
)

const (
	DefaultAlphabet = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DefaultSize     = 21
)

var (
	entropyPool = sync.Pool{
		New: func() interface{} {
			return bufio.NewReader(rand.Reader)
		},
	}
	ErrInvalidLength   = errors.New("invalid ID length")
	ErrUnexpectedParam = errors.New("unexpected parameter")
)

// Generate new id for given size.
func New(s ...int) (string, error) {
	var size int

	switch {
	case len(s) == 0:
		size = DefaultSize
	case len(s) == 1:
		size = s[0]
		if size < 1 {
			return "", ErrInvalidLength
		}
	default:
		return "", ErrUnexpectedParam
	}

	entropy := entropyPool.Get().(*bufio.Reader)
	defer entropyPool.Put(entropy)

	bytes := make([]byte, size)
	if _, err := io.ReadAtLeast(entropy, bytes, size); err != nil {
		return "", err
	}

	for idx, b := range bytes {
		bytes[idx] = DefaultAlphabet[b&63]
	}

	// zero copy conversion from []byte to string
	return unsafe.String(unsafe.SliceData(bytes), len(bytes)), nil
}

// Same as New but panics on error.
func Must(size ...int) string {
	id, err := New(size...)
	if err != nil {
		panic(err)
	}

	return id
}
