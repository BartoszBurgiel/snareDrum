package interpreter

// Pointer represents a pointer on a turing machine
type Pointer struct {
	Pos int
}

// Incr the pointer value by one -> move to the left
func (p *Pointer) Incr() {
	p.Pos++
}

// Decr the pointer value by one -> move to the right
func (p *Pointer) Decr() {
	p.Pos--
}
