package goapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGoapi(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "goapi Test Suite")
}
