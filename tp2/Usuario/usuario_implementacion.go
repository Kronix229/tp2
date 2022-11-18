package usuario

import (
	"algogram/errores"
	TDAHeap "algogram/heap"
)

type usuarioImplementacion[T comparable] struct {
	feed     TDAHeap.ColaPrioridad[*postPrioridad]
	afinidad int
}
type postPrioridad struct {
	prioridad int
	post      Post
}

func CrearUsuario(afinidad int) Usuario[string] {
	usuario := new(usuarioImplementacion[string])
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap(cmp_prioridad)
	return usuario
}

func (usuario *usuarioImplementacion[T]) PublicarPost(dato Usuario[string], post Post, afinidad int) {
	post_prio := new(postPrioridad)
	post_prio.prioridad = val_abs(dato.DevolverAfinidad() - afinidad)
	post_prio.post = post
	dato.DevolverFeed().Encolar(post_prio)
}
func (usuario *usuarioImplementacion[string]) ScrollFeed() error {
	if usuario.feed.EstaVacia() {
		return errores.ErrorFindFeed{}
	}
	post_prio := usuario.feed.Desencolar()
	post_prio.post.MostrarPost()
	return nil
}

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
