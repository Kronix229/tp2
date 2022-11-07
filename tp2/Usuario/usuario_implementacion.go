package usuario

import (
	"algogram/errores"
	TDAHeap "algogram/heap"
)

type usuarioImplementacion[T comparable] struct {
	nombre   string //creo que no es necesario
	feed     TDAHeap.ColaPrioridad[T]
	afinidad int
}
type postPrioridad struct {
	prioridad int
	post      Post
}

func CrearUsuario(nombre string, afinidad int) Usuario[string, int] {
	usuario := new(usuarioImplementacion[string])
	usuario.nombre = nombre
	usuario.afinidad = afinidad
	usuario.feed = TDAHeap.CrearHeap[postPrioridad](cmp_prioridad)
	return usuario
}

func (usuario usuarioImplementacion[string]) PublicarPost(arr []string, dicc TDAHeap.Diccionario[K, V], post Post, afinidad int) {
	for _, nombre := range arr {
		user := dicc.Obtener(nombre) // dicc debe ser un diccionario del TDA Usuario
		if user.nombre == usuario.nombre {
			continue
		}
		post_prio := new(postPrioridad)
		post_prio.prioridad = val_abs(user.afinidad - afinidad)
		post_prio.post = post
		user.feed.Encolar(post_prio)
	}
}
func (usuario *usuarioImplementacion[string]) ScrollFeed() error {
	if usuario.feed.EstaVacia() {
		return errores.ErrorFinFeed{}
	}
	post_prio := usuario.feed.Desencolar()
	post_prio.post.MostrarPost()
	return nil
}
func val_abs(afinidad int) int {
	if afinidad < 0 {
		return -afinidad
	}
	return afinidad
}
func cmp_prioridad(prio1, prio2 postPrioridad) int {
	if prio1.prioridad > prio2.prioridad {
		return -1
	}
	if prio1.prioridad < prio2.prioridad {
		return 1
	}
	return 0
}
