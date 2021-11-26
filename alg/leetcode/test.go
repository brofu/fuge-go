package leetcode

import (
	"testing"
)

type TestCase interface {
	Skipped() bool
	GetName() string
	DebugInfo(*testing.T)
	Execute(*testing.T)
	Verify(*testing.T) bool
}

func RunTestCases(t *testing.T, cases []TestCase) {

	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()

	for _, tc := range cases {
		if tc.Skipped() {
			continue
		}
		t.Log(tc.GetName(), "...")
		tc.Execute(t)
		tc.DebugInfo(t)
		if tc.Verify(t) {
			t.Log(tc.GetName(), "PASS")
		} else {
			t.Log(tc.GetName(), "FAIL")
		}
	}
}

type BaseCase struct {
	skipped bool
	name    string
	debug   bool
}

func (c *BaseCase) Skipped() bool {
	return c.skipped
}

func (c *BaseCase) GetName() string {
	return c.name
}

var baseCase *BaseCase = new(BaseCase)
