package diccionario

import TDAPila "diccionario/pila"

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
	padre, hijo := abb.raiz.buscar(clave, abb)
	if padre == nil {
		if hijo == nil {
			abb.raiz = crearNodo(clave, valor)
			abb.cantidad++
			return
		}
		abb.raiz.dato = valor
		return
	}
	if abb.cmp(padre.clave, clave) > 0 {
		actualizar_hijo(abb, &padre.izq, clave, valor)
	} else {
		actualizar_hijo(abb, &padre.der, clave, valor)
	}
}
func actualizar_hijo[K comparable, V any](abb *abb[K, V], nodo **nodoAbb[K, V], clave K, valor V) {
	if *nodo != nil {
		(*nodo).dato = valor
		return
	}
	*nodo = crearNodo(clave, valor)
	abb.cantidad++
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, hijo := abb.raiz.buscar(clave, abb)
	return hijo != nil
}
func (abb *abb[K, V]) Obtener(clave K) V {
	_, hijo := abb.raiz.buscar(clave, abb)
	if hijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return hijo.dato
}
func (abb *abb[K, V]) Borrar(clave K) V {
	padre, hijo := abb.raiz.buscar(clave, abb)
	if hijo == nil {
		panic("La clave no pertenece al diccionario")
	}
	res := hijo.dato
	if hijo.der == nil || hijo.izq == nil {
		padre.borrar_no_dos_hijos(abb, hijo)
	} else {
		padre.borrar_dos_hijos(abb, hijo)
	}
	return res
}
func (padre *nodoAbb[K, V]) borrar_no_dos_hijos(abb *abb[K, V], hijo *nodoAbb[K, V]) {
	var reemplazo *nodoAbb[K, V]
	if hijo.izq != nil {
		reemplazo = hijo.izq
	}
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
}
func (padre *nodoAbb[K, V]) borrar_dos_hijos(abb *abb[K, V], hijo *nodoAbb[K, V]) {
	reemplazo := hijo.izq // caso 2 hijos
	for reemplazo.der != nil {
		reemplazo = reemplazo.der
	}
	nueva_clave := reemplazo.clave
	nuevo_dato := abb.Borrar(reemplazo.clave)
	hijo.clave = nueva_clave
	hijo.dato = nuevo_dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}
func (d *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	nodo_actual := d.raiz
	seguir := true
	cont := &seguir
	nodo_actual.iterar(visitar, cont)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(clave K, valor V) bool, seguir *bool) {
	if nodo == nil || !(*seguir) {
		return
	}
	nodo.izq.iterar(visitar, seguir)
	if *seguir && !visitar(nodo.clave, nodo.dato) {
		*seguir = false
	}
	nodo.der.iterar(visitar, seguir)
}

func (d *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador, pila := d.inicializar_iterador_pila()
	iterador.dicc = d
	iterador.pila = pila
	if d.raiz == nil {
		return iterador
	}
	iterador.actual = d.raiz
	iterador.desde = d.valor_minimo()
	iterador.hasta = d.valor_maximo()
	iterador.actualizar_actual()
	return iterador
}

func (nodo *nodoAbb[K, V]) buscar(clave K, abb *abb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if abb.raiz == nil {
		return nil, nil
	}
	if abb.raiz.clave == clave {
		return nil, abb.raiz
	}
	return nodo.buscar_recu(clave, abb)
}
func (nodo *nodoAbb[K, V]) buscar_recu(clave K, abb *abb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
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
	}
	return !i.pila.EstaVacia() // caso para el iterador sin rango
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
	i.actualizar_actual()
	return nodo_anterior.clave
}
func (d *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador, pila := d.inicializar_iterador_pila()
	if d.raiz == nil {
		iterador.pila = pila
		iterador.dicc = d
		return iterador
	}
	if desde != nil || hasta != nil {
		iterador.actual = d.raiz
		iterador.desde, iterador.hasta = d.asignar_autovalores_desde_hasta(desde, hasta)
		iterador.dicc = d
		iterador.pila = pila
		iterador.actualizar_actual()
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
			seguir := true
			cont := &seguir
			d.iter_rango_recur(d.raiz, desde, hasta, visitar, cont)
		}
	}
}

func (d *abb[K, V]) iter_rango_recur(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, valor V) bool, seguir *bool) {
	if nodo == nil || !(*seguir) {
		return
	}
	if d.cmp(nodo.clave, *desde) > 0 {
		d.iter_rango_recur(nodo.izq, desde, hasta, visitar, seguir)
	}
	if *seguir {
		if d.cmp(nodo.clave, *desde) >= 0 && d.cmp(nodo.clave, *hasta) <= 0 {
			if !visitar(nodo.clave, nodo.dato) {
				*seguir = false
			}
		}
	}
	if d.cmp(nodo.clave, *hasta) < 0 {
		d.iter_rango_recur(nodo.der, desde, hasta, visitar, seguir)
	}
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = valor
	return nodo
}

// si desde es nil le asigno la clave de la raiz y si
// hasta es nil le asigno el valor mas grande del arbol
func (d *abb[K, V]) asignar_autovalores_desde_hasta(desde *K, hasta *K) (*K, *K) {
	if desde == nil {
		desde = d.valor_minimo()
	} else if hasta == nil {
		hasta = d.valor_maximo()
	}
	return desde, hasta
}

// crea un iterador y una pila para el iterador externo con y sin rango
func (d *abb[K, V]) inicializar_iterador_pila() (*iteradorExternoAbb[K, V], TDAPila.Pila[nodoAbb[K, V]]) {
	iterador := new(iteradorExternoAbb[K, V])
	p := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	return iterador, p
}

func (i *iteradorExternoAbb[K, V]) actualizar_actual() {
	for i.actual != nil {
		if i.dicc.cmp(i.actual.clave, *i.desde) >= 0 && i.dicc.cmp(i.actual.clave, *i.hasta) <= 0 {
			i.pila.Apilar(*i.actual)
			if i.actual.clave == *i.desde {
				break
			}
		}
		if i.dicc.cmp(i.actual.clave, *i.desde) > 0 {
			i.actual = i.actual.izq
		} else {
			i.actual = i.actual.der
		}
	}
}

func (d abb[K, V]) valor_minimo() *K {
	nodo_actual := d.raiz
	for nodo_actual.izq != nil {
		nodo_actual = nodo_actual.izq
	}
	return &nodo_actual.clave
}

func (d abb[K, V]) valor_maximo() *K {
	nodo_actual := d.raiz
	for nodo_actual.der != nil {
		nodo_actual = nodo_actual.der
	}
	return &nodo_actual.clave
}
