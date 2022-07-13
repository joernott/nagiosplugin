package nagiosplugin

import (
	"math"
	"testing"
)

func TestPerfdata(t *testing.T) {
	expected := "'badness'=9003.4ms;4000;9000;10;"
	pd, err := NewPerfDatum("badness", "ms", tNewFloatPerfDatumValue(t, 9003.4), tNewSimpleRangeFromFloat(t, 0, 4000), tNewSimpleRangeFromFloat(t, 0, 9000), newFloat64Ptr(10), nil)
	if err != nil {
		t.Errorf("Could not render perfdata: %v", err)
	}
	if pd.String() != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, pd)
	}
}

func TestRenderPerfdata(t *testing.T) {
	expected := " | 'goodness'=3.141592653589793kb;;;3;34.55751918948773 'goodness'=6.283185307179586kb;;;3;34.55751918948773 'goodness'=9.42477796076938kb;;;3;34.55751918948773 'goodness'=12.566370614359172kb;;;3;34.55751918948773 'goodness'=15.707963267948966kb;;;3;34.55751918948773 'goodness'=18.84955592153876kb;;;3;34.55751918948773 'goodness'=21.991148575128552kb;;;3;34.55751918948773 'goodness'=25.132741228718345kb;;;3;34.55751918948773 'goodness'=28.274333882308138kb;;;3;34.55751918948773 'goodness'=31.41592653589793kb;;;3;34.55751918948773"
	var pd []PerfDatum
	for i := 0; i < 10; i++ {
		datum, err := NewPerfDatum("goodness", "kb", tNewFloatPerfDatumValue(t, math.Pi*float64(i+1)), nil, nil, newFloat64Ptr(3.0), newFloat64Ptr(math.Pi*11))
		if err != nil {
			t.Errorf("Could not create perfdata: %v", err)
		}
		pd = append(pd, *datum)
	}
	result := RenderPerfdata(pd)
	if result != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, result)
	}
}

func TestRenderPerfdataWithOmissions(t *testing.T) {
	var pd []PerfDatum
	datum, err := NewPerfDatum(
		"age",                             // label
		"s",                               // UOM
		tNewFloatPerfDatumValue(t, 0.123), // value
		tNewSimpleRangeFromFloat(t, 0, math.NaN()), // warn: NaN -> omit
		tNewSimpleRangeFromFloat(t, 0, 0.5),        // crit
		newFloat64Ptr(0.0),                         // min
		newFloat64Ptr(math.Inf(1)))                 // max: +Inf -> omit
	if err != nil {
		t.Errorf("Could not create perfdata: %v", err)
	}
	pd = append(pd, *datum)

	// 'label'=value[UOM];[warn];[crit];[min];[max]
	expected := " | 'age'=0.123s;;0.5;0;"
	result := RenderPerfdata(pd)
	if result != expected {
		t.Errorf("Perfdata rendering error: expected %s, got %v", expected, result)
	}
}
