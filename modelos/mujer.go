package modelos

type Mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

func (h *Mujer) Respirar()    { h.respirando = true }
func (h *Mujer) Comer()       { h.respirando = true }
func (h *Mujer) Pensar()      { h.respirando = true }
func (h *Mujer) Sexo() string { return "Mujer" }
