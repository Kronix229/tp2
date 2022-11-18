package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	c "algogram/comandos"
	errores "algogram/errores"
	TDADICC "algogram/hash"
	"fmt"
	"strconv"
)

func Login(usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]], login *c.Login, parametro string) {
	(*login).User = parametro
	err := DevolverErrorAlLoggearse(usuarios_registrados, (*login).User, (*login).Conectado)
	if err != nil {
		fmt.Println(err)
		return
	}

	login.EstadoDelUsuario()
	fmt.Println(login.Saludar())
}
func Logout(login *c.Login) {
	if !login.Conectado {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{})
		return
	}

	fmt.Println(c.Logout{}.Despedir())
	(*login).EstadoDelUsuario()
}
func Publicar(login c.Login, id_global *int, parametro string, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post], usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	if !login.Conectado {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{})
		return
	}
	// Creo su post pasando por parametros la cantidad de post que hizo(me sirve para tener conteo con el id), el texto  su nombre
	nuevo_Post := TDAUSUARIO.CrearPost(*id_global, parametro, login.User)
	posteos.Guardar(*id_global, nuevo_Post)
	(*id_global)++
	// obtengo el TDAusuario que esta actualmente loggeado
	usuario_loggeado := usuarios_registrados.Obtener(login.User)

	ActualizaFeedDelResto(login.User, usuario_loggeado.DevolverAfinidad(), nuevo_Post, usuarios_registrados)
	fmt.Println(c.Publicar{}.ConfirmarPublicacion())
}
func Mostrar_Siguiente_Feed(login c.Login, usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	if !login.Conectado {
		fmt.Println(errores.ErrorFindFeed{})
		return
	}
	usuario_loggeado := usuarios_registrados.Obtener(login.User)
	err := usuario_loggeado.ScrollFeed()
	if err != nil {
		fmt.Println(err)
	}
}
func Likear_Post(parametro string, login c.Login, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post]) {
	id, err := strconv.Atoi(parametro)
	if err != nil {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	if !posteos.Pertenece(id) || !login.Conectado {
		fmt.Println(errores.ErrorPostInexistente{})
		return
	}
	post := posteos.Obtener(id)
	post.LikearPost(login.User)
	fmt.Println(c.Likear{}.ConfirmarLike())
}
func Mostrar_Likes(parametro string, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post]) {
	id, err := strconv.Atoi(parametro)
	if err != nil {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	if !posteos.Pertenece(id) {
		fmt.Println(errores.ErrorSinLikes{})
		return
	}
	post := posteos.Obtener(id)
	err = post.MostrarLikes()
	if err != nil {
		fmt.Println(err)
	}
}
