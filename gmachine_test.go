package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	g := gmachine.New()

	// Add ADD instruction to our current machine state
	g.Memory[0] = gmachine.NoOpcode
	g.Memory[1] = gmachine.NoOpcode

	g.Run()

	if g.P != 2 {
		t.Errorf("Wanted 2 but got %+v\n", g.P)
	}
}
