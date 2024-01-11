package modelos

type Mujer struct {
	Hombre
}

// EstaVivo implements interfaces.SerVivo.
func (*Mujer) EstaVivo() bool {
	panic("unimplemented")
}

func (m *Mujer) Respirar()     { m.respirando = true }
func (m *Mujer) Comer()        { m.comiendo = true }
func (m *Mujer) Pensar()       { m.pensando = true }
func (m *Mujer) Sexo() string  { return "Mujer" }
func (m *Mujer) SerVivo() bool { return true }
