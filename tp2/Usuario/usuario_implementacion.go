package usuario

import (
	"algogram/errores"
	TDAHash "algogram/hash"
	TDAHeap "algogram/heap"
)

type usuarioImplementacion[T comparable] struct {
	nombre   string //creo que no es necesario
	feed     TDAHeap.ColaPrioridad[*postPrioridad]
	afinidad int
}
type postPrioridad struct {
	prioridad int
	post      Post
}

func CrearUsuario(nombre string, afinidad int) Usuario[string] {
	usuario := new(usuarioImplementacion[string])
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap(cmp_prioridad)
	return usuario
}

// Cambie el for por el iterador del diccionario asi evitada tener que crear un array de los usuarios
// le actualiza a todos su feed menos el que esta loggeado(asi fue como lo entendi)
func (usuario *usuarioImplementacion[T]) PublicarPost(usuario_loggeado Usuario[string], dicc TDAHash.Diccionario[string, Usuario[string]], post Post, afinidad int) {
	dicc.Iterar(func(clave string, dato Usuario[string]) bool {
		if clave == usuario_loggeado.DevolverNombre() {
			return true
		}
		post_prio := new(postPrioridad)
		post_prio.prioridad = val_abs(dato.DevolverAfinidad() - afinidad)
		post_prio.post = post
		dato.DevolverFeed().Encolar(post_prio)
		return true
	})
}
func (usuario *usuarioImplementacion[string]) ScrollFeed() error {
	if usuario.feed.EstaVacia() {
		return errores.ErrorFindFeed{}
	}
	post_prio := usuario.feed.Desencolar()
	post_prio.post.MostrarPost()
	return nil
}

// le agrege por conveniencia
func (usuario usuarioImplementacion[T]) DevolverNombre() string {
	return usuario.nombre
}

// le agrege por conveniencia
func (usuario usuarioImplementacion[T]) DevolverAfinidad() int {
	return usuario.afinidad
}

func (usuario usuarioImplementacion[T]) DevolverFeed() TDAHeap.ColaPrioridad[*postPrioridad] {
	return usuario.feed
}

func val_abs(afinidad int) int {
	if afinidad < 0 {
		return -afinidad
	}
	return afinidad
}
func cmp_prioridad(prio1, prio2 *postPrioridad) int {
	if prio1.prioridad > prio2.prioridad {
		return -1
	}
	if prio1.prioridad < prio2.prioridad {
		return 1
	}
	if prio1.post.DevolverId() > prio2.post.DevolverId() {
		return -1
	}
	return 1
}
