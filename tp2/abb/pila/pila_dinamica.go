package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaEnlazada[T any] struct {
	datos    []T
	cantidad int
}

var _FACTOR_REDIMENSION int = 2

func (p *pilaEnlazada[T]) redimensionar() {
	if p.cantidad == len(p.datos) {
		lista_temp := p.datos
		p.datos = make([]T, p.cantidad*_FACTOR_REDIMENSION)
		_ = copy(p.datos, lista_temp)
	} else if (p.cantidad * _FACTOR_REDIMENSION * 2) <= len(p.datos) {
		lista_temp := p.datos
		p.datos = make([]T, len(p.datos)/_FACTOR_REDIMENSION)
		_ = copy(p.datos, lista_temp)

	}
}

func (p pilaEnlazada[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaEnlazada[T]) Apilar(elem T) {
	p.datos[p.cantidad] = elem
	p.cantidad++
	p.redimensionar()

}

func (p *pilaEnlazada[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		elem := p.datos[p.cantidad-1]
		p.redimensionar()
		p.cantidad--
		return elem
	}
}

func (p pilaEnlazada[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		elem := p.datos[p.cantidad-1]
		return elem
	}
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaEnlazada[T])
	pila.datos = make([]T, 2)
	pila.cantidad = 0
	return pila
}
