package diccionario

import TDAPila "algogram/abb/pila"

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type iteradorExternoAbb[K comparable, V any] struct {
	dicc   *abb[K, V]
	actual *nodoAbb[K, V]
	pila   TDAPila.Pila[nodoAbb[K, V]]
	desde  *K
	hasta  *K
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.cmp = funcion_cmp
	return abb
}
func (abb *abb[K, V]) Guardar(clave K, valor V) {
	if abb.raiz == nil {
		abb.raiz = abb.crearNodo(clave, valor)
		abb.cantidad++
	} else if abb.raiz.clave == clave {
		abb.raiz.dato = valor
	} else {
		abb.raiz.guardar(clave, valor, abb)
	}
}
func (nodo *nodoAbb[K, V]) guardar(clave K, valor V, abb *abb[K, V]) {
	padre, _ := abb.raiz.buscar(clave, abb)
	if abb.cmp(padre.clave, clave) > 0 {
		if padre.izq != nil {
			padre.izq.dato = valor
			return
		}
		padre.izq = abb.crearNodo(clave, valor)
	} else {
		if padre.der != nil {
			padre.der.dato = valor
			return
		}
		padre.der = abb.crearNodo(clave, valor)
	}
	abb.cantidad++

}
func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.raiz == nil {
		return false
	}
	if abb.raiz.clave == clave {
		return true
	}
	_, hijo := abb.raiz.buscar(clave, abb)
	return hijo != nil
}
func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	if abb.raiz.clave == clave {
		return abb.raiz.dato
	}
	_, hijo := abb.raiz.buscar(clave, abb)
	if hijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return hijo.dato
}
func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	return abb.raiz.borrar(clave, abb)
}
func (nodo *nodoAbb[K, V]) borrar(clave K, abb *abb[K, V]) V {
	var padre *nodoAbb[K, V]
	var hijo *nodoAbb[K, V]
	if abb.raiz.clave == clave {
		padre = nil
		hijo = abb.raiz
	} else {
		padre, hijo = nodo.buscar(clave, abb)
	}
	if hijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	res := hijo.dato
	if hijo.der == nil && hijo.izq == nil {
		if padre != nil {
			if padre.der == hijo {
				padre.der = nil
			} else {
				padre.izq = nil
			}
		} else {
			abb.raiz = nil
		}
		abb.cantidad--
		return res
	}
	if hijo.der != nil && hijo.izq == nil || hijo.der == nil && hijo.izq != nil {
		reemplazo := hijo.izq
		if hijo.der != nil {
			reemplazo = hijo.der
		}
		if padre != nil {
			if padre.der == hijo {
				padre.der = reemplazo
			} else {
				padre.izq = reemplazo
			}
		} else {
			abb.raiz = reemplazo
		}
		abb.cantidad--
		return res
	}
	reemplazo := hijo.izq // caso 2 hijos
	for reemplazo.der != nil {
		reemplazo = reemplazo.der
	}
	nueva_clave := reemplazo.clave
	nuevo_dato := abb.Borrar(reemplazo.clave)
	hijo.clave = nueva_clave
	hijo.dato = nuevo_dato
	return res

}
func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}
func (d *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	nodo_actual := d.raiz
	nodo_actual.iterar(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(clave K, valor V) bool) {
	if nodo == nil {
		return
	}
	nodo.izq.iterar(visitar)
	visitar(nodo.clave, nodo.dato)
	nodo.der.iterar(visitar)
}
func (d *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador, p := d.inicializar_iterador_pila()
	iterador.actual = d.raiz
	for iterador.actual != nil {
		p.Apilar(*iterador.actual)
		iterador.actual = iterador.actual.izq
	}
	iterador.pila = p
	return iterador
}
func (nodo *nodoAbb[K, V]) buscar(clave K, abb *abb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	cmp := abb.cmp(nodo.clave, clave)
	if cmp > 0 && nodo.izq == nil {
		return nodo, nil
	}
	if cmp > 0 {
		if nodo.izq.clave == clave {
			return nodo, nodo.izq
		}
		return nodo.izq.buscar(clave, abb)
	}
	if cmp <= 0 && nodo.der == nil {
		return nodo, nil
	}
	if nodo.der.clave == clave {
		return nodo, nodo.der
	}
	return nodo.der.buscar(clave, abb)
}
func (i *iteradorExternoAbb[K, V]) HaySiguiente() bool {
	if i.desde != nil && i.hasta != nil { // caso para el iterador con rango
		if i.pila.EstaVacia() || i.dicc.cmp(i.pila.VerTope().clave, *i.hasta) > 0 {
			return false
		}
	} else if i.pila.EstaVacia() { // caso para el iterador sin rango
		return false
	}
	return true
}

func (i *iteradorExternoAbb[K, V]) VerActual() (clave K, valor V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.pila.VerTope().clave, i.pila.VerTope().dato
}

func (i *iteradorExternoAbb[K, V]) Siguiente() K {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo_anterior := i.pila.Desapilar()
	i.actual = nodo_anterior.der
	for i.actual != nil {
		if i.desde != nil && i.hasta != nil {
			if i.dicc.cmp(i.actual.clave, *i.desde) > 0 {
				i.pila.Apilar(*i.actual)
				i.actual = i.actual.izq
			}
		} else {
			i.pila.Apilar(*i.actual)
			i.actual = i.actual.izq
		}
	}
	return nodo_anterior.clave
}
func (d *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador, p := d.inicializar_iterador_pila()
	if d.raiz == nil {
		iterador.pila = p
		iterador.dicc = d
		return iterador
	}
	iterador.actual = d.raiz
	if desde != nil || hasta != nil {
		desde, hasta = d.asignar_autovalores_desde_hasta(desde, hasta)
		for iterador.actual != nil {
			if d.cmp(iterador.actual.clave, *desde) >= 0 && d.cmp(iterador.actual.clave, *hasta) <= 0 {
				p.Apilar(*iterador.actual)
				if iterador.actual.clave == *desde {
					break
				}
			}
			if d.cmp(iterador.actual.clave, *desde) > 0 {
				iterador.actual = iterador.actual.izq
			} else {
				iterador.actual = iterador.actual.der
			}
		}
		iterador.desde = desde
		iterador.hasta = hasta
		iterador.pila = p
		iterador.dicc = d
		return iterador
	}
	return d.Iterador()
}

func (d *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if d.raiz != nil {
		if desde == nil && hasta == nil {
			d.Iterar(visitar)
		} else {
			desde, hasta = d.asignar_autovalores_desde_hasta(desde, hasta)
			d.iter_rango_recur(d.raiz, desde, hasta, visitar)
		}
	}
}

func (d *abb[K, V]) iter_rango_recur(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, valor V) bool) {
	if nodo == nil {
		return
	}
	if d.cmp(nodo.clave, *desde) > 0 {
		d.iter_rango_recur(nodo.izq, desde, hasta, visitar)
	}
	if d.cmp(nodo.clave, *desde) >= 0 && d.cmp(nodo.clave, *hasta) <= 0 {
		visitar(nodo.clave, nodo.dato)
	}
	if d.cmp(nodo.clave, *hasta) < 0 {
		d.iter_rango_recur(nodo.der, desde, hasta, visitar)
	}
}

func (abb *abb[K, V]) crearNodo(clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = valor
	return nodo
}

// si desde es nil le asigno la clave de la raiz y si
// hasta es nil le asigno el valor mas grande del arbol
func (d *abb[K, V]) asignar_autovalores_desde_hasta(desde *K, hasta *K) (*K, *K) {
	nodo_actual := *d.raiz
	if desde == nil {
		desde = &nodo_actual.clave
	} else if hasta == nil {
		for nodo_actual.der != nil {
			nodo_actual = *nodo_actual.der
		}
		hasta = &nodo_actual.clave
	}
	return desde, hasta
}

// crea un iterador y una pila para el iterador externo con y sin rango
func (d *abb[K, V]) inicializar_iterador_pila() (*iteradorExternoAbb[K, V], TDAPila.Pila[nodoAbb[K, V]]) {
	iterador := new(iteradorExternoAbb[K, V])
	p := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	return iterador, p
}
