package data_test

import (
	"go-testtest/data"
	"go-testtest/packageclient"
	"testing"
)

func Test(t *testing.T) {
	foo := data.NewFoo("foo")
	_ = foo

	packageclient.NewRD()
}
