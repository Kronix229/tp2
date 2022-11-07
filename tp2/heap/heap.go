package cola_prioridad

const (
	CAPACIDAD_INICIAL  = 10
	FACTOR_REDIMENSION = 2
)

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h, arreglo := inicializar_heap[T]()
	h.datos = arreglo
	h.cmp = funcion_cmp
	return h
}
func CrearHeapArr[T comparable](Arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(Arreglo) != 0 {
		h, arr_iniciado := inicializar_heap[T]()
		h.datos = Arreglo
		h.cantidad = len(h.datos)
		h.redimensionar(len(arr_iniciado))
		h.cmp = funcion_cmp
		heapify(&h.datos, h.cmp)
		return h
	}
	return CrearHeap(funcion_cmp)
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}
func (h *heap[T]) Encolar(elem T) {
	h.datos[h.cantidad] = elem
	upheap(h.datos, h.cmp, h.cantidad, pos_padre(h.cantidad))
	h.cantidad++
	if h.cantidad == cap(h.datos) {
		h.redimensionar(cap(h.datos) * FACTOR_REDIMENSION)
	}
}
func (h *heap[T]) VerMax() T {
	if h.cantidad == 0 {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.cantidad == 0 {
		panic("La cola esta vacia")
	}
	if h.cantidad*4 <= cap(h.datos) && cap(h.datos) > CAPACIDAD_INICIAL {
		h.redimensionar(cap(h.datos) / FACTOR_REDIMENSION)
	}
	res := h.datos[0]
	h.cantidad--
	swap(&h.datos, 0, h.cantidad)
	downheap(&h.datos, h.cantidad-1, 0, h.cmp)
	return res
}
func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

// Le da forma de heap al arreglo
func heapify[T comparable](arreglo *[]T, cmp func(T, T) int) []T {
	downheap(arreglo, len(*arreglo)-1, len(*arreglo)-1, cmp)
	return *arreglo
}

// Encuentra la posicion de los hijos
func pos_hijos(padre int) (int, int) {
	pos_hijo_dere := 2*padre + 1
	pos_hijo_izqu := 2*padre + 2
	return pos_hijo_dere, pos_hijo_izqu
}

// Encuentra la posicion del padre del hijo que lo llama
func pos_padre(hijo int) int {
	pos_padre := (hijo - 1) / 2
	return pos_padre
}

// Cambia de posicion entre dos elementos
func swap[T comparable](arreglo *[]T, a int, b int) {
	(*arreglo)[a], (*arreglo)[b] = (*arreglo)[b], (*arreglo)[a]
}

func upheap[T comparable](h []T, cmp func(T, T) int, posicion, padre int) {
	if posicion == 0 || cmp(h[posicion], h[padre]) < 0 {
		return
	}
	swap(&h, posicion, padre)
	posicion = padre
	upheap(h, cmp, posicion, pos_padre(posicion))
}

// Aplico dowheap del ultimo al primero
func downheap[T comparable](arreglo *[]T, longitud int, posicion int, cmp func(T, T) int) {
	hijo_dere, hijo_izqu := pos_hijos(posicion)
	downheap_recu(arreglo, longitud, posicion, hijo_dere, hijo_izqu, cmp)
}

// func downheap_recu[T comparable](arreglo *[]T, longitud int, posicion int, hijo_dere int, hijo_izqu int, cmp func(T, T) int) {
// 	if posicion == -1 {
// 		return
// 	}
// 	if hijo_dere <= longitud && hijo_izqu <= longitud {
// 		if cmp((*arreglo)[posicion], (*arreglo)[hijo_dere]) < 0 && cmp((*arreglo)[posicion], (*arreglo)[hijo_izqu]) < 0 {
// 			posicion_max := maximo(*arreglo, hijo_dere, hijo_izqu, cmp)
// 			swap(arreglo, posicion, posicion_max)
// 			posicion = posicion_max
// 			hijo_dere, hijo_izqu = pos_hijos(posicion)
// 			downheap_recu(arreglo, longitud, posicion, hijo_dere, hijo_izqu, cmp)
// 		}
// 	}
// 	if hijo_dere <= longitud && cmp((*arreglo)[posicion], (*arreglo)[hijo_dere]) < 0 {
// 		swap(arreglo, posicion, hijo_dere)
// 		posicion = hijo_dere
// 	} else if hijo_izqu <= longitud && cmp((*arreglo)[posicion], (*arreglo)[hijo_izqu]) < 0 {
// 		swap(arreglo, posicion, hijo_izqu)
// 		posicion = hijo_dere
// 	} else {
// 		posicion--
// 	}

// 	hijo_dere, hijo_izqu = pos_hijos(posicion)
// 	downheap_recu(arreglo, longitud, posicion, hijo_dere, hijo_izqu, cmp)
// }

func downheap_recu[T comparable](arreglo *[]T, longitud int, posicion int, hijo_dere int, hijo_izqu int, cmp func(T, T) int) {
	if posicion == -1 {
		return
	}
	if hijo_dere <= longitud && hijo_izqu <= longitud {
		pos_valor_max := maximo(*arreglo, hijo_dere, hijo_izqu, cmp)
		if cmp((*arreglo)[posicion], (*arreglo)[pos_valor_max]) < 0 {
			swap(arreglo, posicion, pos_valor_max)
			posicion = pos_valor_max
		} else {
			posicion--
		}
	} else {
		if hijo_dere <= longitud && cmp((*arreglo)[posicion], (*arreglo)[hijo_dere]) < 0 {
			swap(arreglo, posicion, hijo_dere)
			posicion = hijo_dere
		} else if hijo_izqu <= longitud && cmp((*arreglo)[posicion], (*arreglo)[hijo_izqu]) < 0 {
			swap(arreglo, posicion, hijo_izqu)
			posicion = hijo_izqu
		} else {
			posicion--
		}
	}
	hijo_dere, hijo_izqu = pos_hijos(posicion)
	downheap_recu(arreglo, longitud, posicion, hijo_dere, hijo_izqu, cmp)
}

func (h *heap[T]) redimensionar(capacidad int) {
	nuevo := make([]T, capacidad)
	copy(nuevo, h.datos)
	h.datos = nuevo
}

func HeapSort[T comparable](elementos []T, cmp func(T, T) int) {
	elementos = heapify(&elementos, cmp)
	largo := len(elementos) - 1
	for largo > 0 {
		swap(&elementos, 0, largo)
		largo--
		downheap(&elementos, largo, 0, cmp)
	}
}

// Devuelve la posicion que posea el valor mas alto dentro del arreglo
func maximo[T comparable](arreglo []T, pos1 int, pos2 int, cmp func(T, T) int) int {
	if cmp(arreglo[pos1], arreglo[pos2]) > 0 {
		return pos1
	}
	return pos2
}

func inicializar_heap[T comparable]() (*heap[T], []T) {
	h := new(heap[T])
	arreglo := make([]T, CAPACIDAD_INICIAL)
	return h, arreglo
}
