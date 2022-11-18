package usuario

import (
	TDAabb "algogram/abb"
	"algogram/errores"
	"fmt"
	"strings"
)

const (
	CERO = 0
)

type postImplementacion struct {
	ID            int
	Texto         string
	Creador       string
	Likes         int
	Usuarios_like TDAabb.DiccionarioOrdenado[string, int]
}

func CrearPost(id int, texto string, creador string) Post {
	post := new(postImplementacion)
	post.ID = id
	post.Texto = texto
	post.Creador = creador
	post.Usuarios_like = TDAabb.CrearABB[string, int](strings.Compare)
	return post
}
func (post postImplementacion) MostrarPost() {
	fmt.Printf("Post ID %d\n", post.ID)
	fmt.Printf("%s dijo: %s\n", post.Creador, post.Texto)
	fmt.Printf("Likes: %d\n", post.Likes)
}
func (post *postImplementacion) LikearPost(nombre_del_usuario string) { //podria ser el nombre en ves de TDA usuario
	if post.Usuarios_like.Pertenece(nombre_del_usuario) {
		return
	}
	post.Usuarios_like.Guardar(nombre_del_usuario, 0)
	post.Likes++
}

func (post postImplementacion) MostrarLikes() error {
	if post.Likes == CERO {
		return errores.ErrorSinLikes{}
	}
	fmt.Printf("El post tiene %d likes:\n", post.Likes)
	iter := post.Usuarios_like.Iterador()
	for iter.HaySiguiente() {
		fmt.Printf("\t%s\n", iter.Siguiente())
	}
	return nil
}

func (post postImplementacion) DevolverId() int {
	return post.ID
}
