package usuario

import (
	"algogram/errores"
	TDAHash "algogram/hash"
	TDAHeap "algogram/heap"
)

type usuarioImplementacion[T comparable] struct {
	nombre           string //creo que no es necesario
	feed             TDAHeap.ColaPrioridad[*postPrioridad]
	afinidad         int
	cantidad_de_post int
}
type postPrioridad struct {
	prioridad int
	post      Post
}

func CrearUsuario(nombre string, afinidad int) Usuario[string] {
	usuario := new(usuarioImplementacion[string])
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap[*postPrioridad](cmp_prioridad)
	return usuario
}

func (usuario *usuarioImplementacion[T]) PublicarPost(usuario_loggeado Usuario[string], dicc TDAHash.Diccionario[string, Usuario[string]], post Post, afinidad int) {
	dicc.Iterar(func(clave string, dato Usuario[string]) bool {
		if clave == usuario_loggeado.DevolverNombre() {
			return true
		}
		post_prio := new(postPrioridad)
		post_prio.prioridad = val_abs(dato.DevolverAfinidad() - afinidad)
		post_prio.post = post
		dato.ActualizarFeed(post_prio)
		return true
	})
	usuario.cantidad_de_post++
}
func (usuario *usuarioImplementacion[string]) ScrollFeed() error {
	if usuario.feed.EstaVacia() {
		return errores.ErrorFindFeed{}
	}
	post_prio := usuario.feed.Desencolar()
	post_prio.post.MostrarPost()
	return nil
}

func (usuario usuarioImplementacion[T]) DevolverNombre() string {
	return usuario.nombre
}

func (usuario usuarioImplementacion[T]) DevolverAfinidad() int {
	return usuario.afinidad
}

func (usuario usuarioImplementacion[T]) DevolverCantidadPost() int {
	return usuario.cantidad_de_post
}

func (usuario *usuarioImplementacion[T]) ActualizarFeed(post *postPrioridad) {
	usuario.feed.Encolar(post)
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
	return 0
}
