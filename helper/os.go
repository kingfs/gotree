package helper

import (
	"errors"
	"os"
	"strings"

	"github.com/huandu/go-tls/g"
)

var (
	_testing   bool = false
	ErrBreaker      = errors.New("Fuse")
)

func init() {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "/unit") || strings.Contains(wd, "\\unit") {
		_testing = true
	}
}

type VoidValue struct {
	Void byte
}

func Testing() bool {
	return _testing
}

func Exit(errorMsg ...string) {
	if len(errorMsg) == 0 {
		os.Exit(0)
	}
	Log().Error(errorMsg)
	os.Exit(-1)
}

func RuntimePointer() interface{} {
	return g.G()
}
