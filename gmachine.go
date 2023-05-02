// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	HaltOpcode = 0
	NoOpcode   = 1
)

type machine struct {
	P      uint64
	Memory map[uint64]uint64
}

func New() *machine {
	return &machine{
		P:      0,
		Memory: make(map[uint64]uint64),
	}
}

// TODO - want to be able to use string OPCODES ("HALT")
func (m *machine) Run() {
	for {
		current := m.Memory[m.P]
		if current == HaltOpcode {
			return
		}
		m.P++
	}
}
