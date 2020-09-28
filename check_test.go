package nagiosplugin

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func tNewFloatPerfDatumValue(t *testing.T, value float64) FloatPerfDatumValue {
	pdata, err := NewFloatPerfDatumValue(value)
	if err != nil {
		t.Fatalf("creating new perfdata from float: %s", err)
	}
	return pdata
}

func tNewSimpleFrangeFromFloat(t *testing.T, start, end float64) *Range {
	r, err := NewSimpleFrangeFromFloat(start, end)
	if err != nil {
		t.Fatalf("parsing range: %s", err)
	}
	return r
}

func newFloat64Ptr(value float64) *float64 {
	v := value
	return &v
}

func TestCheck(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	c := NewCheck()
	expected := "CRITICAL: 200000 terrifying space monkeys in the engineroom | 'space_monkeys'=200000c;10000;100000;0;4294967296"
	nSpaceMonkeys := float64(200000)
	maxSpaceMonkeys := float64(1 << 32)
	c.AddPerfDatum("space_monkeys", "c", tNewFloatPerfDatumValue(t, nSpaceMonkeys), tNewSimpleFrangeFromFloat(t, 0, 10000), tNewSimpleFrangeFromFloat(t, 0, 100000), newFloat64Ptr(0), newFloat64Ptr(maxSpaceMonkeys))
	c.AddResult(CRITICAL, fmt.Sprintf("%v terrifying space monkeys in the engineroom", nSpaceMonkeys))
	// Check a WARNING can't override a CRITICAL
	c.AddResult(WARNING, fmt.Sprintf("%v slightly annoying space monkeys in the engineroom", nSpaceMonkeys))
	result := c.String()
	if expected != result {
		t.Errorf("Expected check output %v, got check output %v", expected, result)
	}
}

func TestDefaultStatusPolicy(t *testing.T) {
	c := NewCheck()
	c.AddResult(WARNING, "Isolated-frame flux emission outside threshold")
	c.AddResult(UNKNOWN, "No response from betaform amplifier")

	expected := "UNKNOWN"
	actual := strings.SplitN(c.String(), ":", 2)[0]
	if actual != expected {
		t.Errorf("Expected %v status, got %v", expected, actual)
	}
}

func TestOUWCStatusPolicy(t *testing.T) {
	c := NewCheckWithOptions(CheckOptions{
		StatusPolicy: NewOUWCStatusPolicy(),
	})
	c.AddResult(WARNING, "Isolated-frame flux emission outside threshold")
	c.AddResult(UNKNOWN, "No response from betaform amplifier")

	expected := "WARNING"
	actual := strings.SplitN(c.String(), ":", 2)[0]
	if actual != expected {
		t.Errorf("Expected %v status, got %v", expected, actual)
	}
}
