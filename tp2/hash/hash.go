package diccionario

import (
	"fmt"
)

type estados int

const (
	VACIO estados = iota
	OCUPADO
	BORRADO
)

const (
	COTA_MAYOR_DE_FACTOR_DE_CARGA = 0.7
	COTA_MENOR_DE_FACTOR_DE_CARGA = 0.1
	FACTOR_DE_REDIMENSION         = 2
	TAMAÑO_INICIAL                = 13
)

type celdas[K comparable, V any] struct {
	clave  K
	valor  V
	estado estados
}

type dicc_implementado[K comparable, V any] struct {
	tamaño       uint
	arreglo      []celdas[K, V]
	cant_borrado int
	cant_ocupado int
}

type iteradorExterno[K comparable, V any] struct {
	dicc   *dicc_implementado[K, V]
	actual int
}

// crear el diccionario
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	Hash := new(dicc_implementado[K, V])
	Hash.tamaño = TAMAÑO_INICIAL
	Hash.arreglo = crear_tabla[K, V](Hash.tamaño)
	return Hash
}

func (d *dicc_implementado[K, V]) Guardar(clave K, valor V) {
	indice := generador_indices(clave, d.tamaño)
	factor_carga := d.calculo_de_factor_de_carga()
	if d.arreglo[indice].estado == OCUPADO {
		d.posicionar_indice(clave, &indice)
	}
	d.arreglo[indice].clave = clave
	d.arreglo[indice].valor = valor
	if d.arreglo[indice].estado != OCUPADO {
		d.arreglo[indice].estado = OCUPADO
		d.cant_ocupado++
	}
	if float32(factor_carga) > COTA_MAYOR_DE_FACTOR_DE_CARGA {
		nuevo_tamaño := d.nuevo_tamaño(factor_carga)
		if nuevo_tamaño > TAMAÑO_INICIAL {
			d.redimensionar(nuevo_tamaño, d.tamaño, d.arreglo)
		}
	}
}

func (d dicc_implementado[K, V]) Pertenece(clave K) bool {
	indice := generador_indices(clave, d.tamaño)
	d.posicionar_indice(clave, &indice)
	return d.arreglo[indice].estado == OCUPADO

}

func (d dicc_implementado[K, V]) Obtener(clave K) V {
	if d.Pertenece(clave) {
		indice := generador_indices(clave, d.tamaño)
		d.posicionar_indice(clave, &indice)
		return d.arreglo[indice].valor
	}
	panic("La clave no pertenece al diccionario")
}

func (d *dicc_implementado[K, V]) Borrar(clave K) V {
	indice := generador_indices(clave, d.tamaño)
	factor_carga := d.calculo_de_factor_de_carga()
	d.posicionar_indice(clave, &indice)
	if d.Pertenece(clave) {
		d.arreglo[indice].estado = BORRADO
		d.cant_borrado++
		valor := d.arreglo[indice].valor
		if factor_carga < COTA_MENOR_DE_FACTOR_DE_CARGA {
			nuevo_tamaño := d.nuevo_tamaño(factor_carga)
			if nuevo_tamaño > TAMAÑO_INICIAL {
				d.redimensionar(nuevo_tamaño, d.tamaño, d.arreglo)
			}
		}
		return valor
	}

	panic("La clave no pertenece al diccionario")
}

func (d dicc_implementado[K, V]) Cantidad() int {
	return d.cant_ocupado - d.cant_borrado
}

func (d dicc_implementado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	var actual uint
	for actual < d.tamaño {
		if d.arreglo[actual].estado != OCUPADO {
			if actual == d.tamaño-1 {
				break
			}
			actual++
		} else if visitar(d.arreglo[actual].clave, d.arreglo[actual].valor) {
			actual++
		} else {
			break
		}
	}
}

func (d *dicc_implementado[K, V]) redimensionar(nuevo_tamaño uint, tamaño_anterior uint, arreglo_temp []celdas[K, V]) {
	d.tamaño = nuevo_tamaño
	d.arreglo = crear_tabla[K, V](d.tamaño)
	d.cant_ocupado = 0
	d.cant_borrado = 0
	for i := uint(0); i < tamaño_anterior; i++ {
		if arreglo_temp[i].estado == OCUPADO {
			d.Guardar(arreglo_temp[i].clave, arreglo_temp[i].valor)
		}
	}
}

func (d *dicc_implementado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iteradorExterno[K, V])
	iter.dicc = d
	iter.saltarse_vacios()
	return iter
}
func (iter *iteradorExterno[K, V]) HaySiguiente() bool {
	return iter.actual != int(iter.dicc.tamaño)
}
func (iter *iteradorExterno[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.dicc.arreglo[iter.actual].clave, iter.dicc.arreglo[iter.actual].valor
}
func (iter *iteradorExterno[K, V]) Siguiente() K {
	if iter.HaySiguiente() {
		valor := iter.dicc.arreglo[iter.actual].clave
		iter.actual++
		if iter.actual != int(iter.dicc.tamaño) {
			iter.saltarse_vacios()
		}
		return valor
	} else {
		panic("El iterador termino de iterar")
	}
}

// funciones auxiliares

func es_primo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Posiciona al indice en un lugar vacio, si la clave ya se encuentra en
// el arreglo lo posiona en el lugar que ocupa
func (d dicc_implementado[K, V]) posicionar_indice(clave K, indice *int) {
	for d.arreglo[*indice].estado != VACIO {
		if d.arreglo[*indice].clave == clave {
			break
		}
		if *indice == int(d.tamaño)-1 {
			*indice = -1
		}
		(*indice)++
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// genera indices para cada clave pasada por parametro
func generador_indices[K comparable](clave K, tamaño uint) int {
	clave_byte := convertirABytes(clave)
	indice := funcion_hash(clave_byte, tamaño)
	return indice
}

// jenkins hash
func funcion_hash(clave []byte, tamaño uint) int {
	hash := uint32(0)
	i := 0
	for i != len(clave) {
		hash += uint32(clave[i])
		hash += hash << 10
		hash ^= hash >> 6
		i++
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return int(hash) % int(tamaño)
}

func crear_tabla[K comparable, V any](tamaño uint) []celdas[K, V] {
	return make([]celdas[K, V], tamaño)
}

func (d dicc_implementado[K, V]) calculo_de_factor_de_carga() float32 {
	factor_carga := float32((d.cant_ocupado)) / float32(d.tamaño)
	return factor_carga
}

// Determina un nuevo tamaño para el arreglo corroborando siempre que sea
// un numero primo
func (d dicc_implementado[K, V]) nuevo_tamaño(factor_carga float32) uint {
	nuevo_tamaño := 0
	if float32(factor_carga) > COTA_MAYOR_DE_FACTOR_DE_CARGA {
		for nuevo_tamaño = int(d.tamaño) * FACTOR_DE_REDIMENSION; !es_primo(nuevo_tamaño); {
			nuevo_tamaño--
		}
	} else if float32(factor_carga) < COTA_MENOR_DE_FACTOR_DE_CARGA {
		for nuevo_tamaño = int(d.tamaño) / FACTOR_DE_REDIMENSION; !es_primo(nuevo_tamaño); {
			nuevo_tamaño--
		}
	}
	return uint(nuevo_tamaño)
}

// Recorre el iterador hasta encontrase con una posicion ocupada
func (iter *iteradorExterno[K, V]) saltarse_vacios() {
	for iter.dicc.arreglo[iter.actual].estado != OCUPADO {
		iter.actual++
		if uint(iter.actual) == iter.dicc.tamaño {
			break
		}
	}
}
