package retry

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNoRetry ...
func TestNoRetry(t *testing.T) {
	a := func() error {
		log.Println("hello, world")
		return nil
	}

	err := Retry(3, 1*time.Millisecond, a)
	assert.Nil(t, err)
}

// TestRetryEveryTime...
func TestRetryEveryTime(t *testing.T) {
	cnt := 0
	err := fmt.Errorf("test error every time")
	a := func() error {
		log.Printf("TestRetryEveryTime ...\n")
		cnt++
		return err
	}
	errFn := Retry(3, 1*time.Millisecond, a)
	assert.Equal(t, err, errFn)
	assert.Equal(t, 4, cnt)
}

// TestRetryOnce...
func TestRetryOnce(t *testing.T) {
	cnt := 0
	err := fmt.Errorf("testing error every time")
	a := func() error {
		log.Printf("TestRetryOnce ...\n")
		if cnt == 0 {
			cnt++
			return err
		} else {
			cnt++
			return nil
		}
	}

	errFn := Retry(3, 1*time.Millisecond, a)
	assert.Nil(t, errFn)
	assert.Equal(t, 2, cnt)
}

// TestRetryStop test stop retry after first call fn.
func TestRetryStop(t *testing.T) {
	cnt := 0
	err := fmt.Errorf("test error every time")
	a := func() error {
		cnt++
		return NoRetryError(err)
	}
	errFn := Retry(3, 1*time.Millisecond, a)
	assert.Equal(t, errFn, err)
	assert.Equal(t, 2, cnt)
}
