package comandos

import (
	TDAUSUARIO "algogram/Usuario"
	errores "algogram/errores"
	faux "algogram/funciones_aux"
	TDADICC "algogram/hash"
	"fmt"
	"strconv"
)

// Un pequeño struct que maneja si ya hay alguien loggeado y saluda al usuario al registrarse
type Login struct {
	User      string
	Conectado bool
}

// Saluda al usuario
func (c Login) Saludar() string {
	return fmt.Sprintf("Hola %s", c.User)
}

// Cambia a true o false si el usuario esta o no conectado
func (c *Login) EstadoDelUsuario() {
	if c.Conectado {
		c.Conectado = false
	} else {
		c.Conectado = true
	}
}

// Loggea al usuario y corroborando que el nombre del usuario sea valido y tambien que no haya nadie conectado
func (c *Login) LoggearUsuario(parametro string, usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	c.User = parametro
	err := faux.DevolverErrorAlLoggearse(usuarios_registrados, c.User, c.Conectado)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.EstadoDelUsuario()
	fmt.Println(c.Saludar())
}

type Logout struct{}

// Despide al usuario
func (c Logout) Despedir() string {
	return "Adios"
}

// Desconecta al usuario corroborando primero si esta conectado
func (c Logout) Desloggear(l *Login) {
	if !l.Conectado {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{})
		return
	}

	fmt.Println(c.Despedir())
	(*l).EstadoDelUsuario()
}

type Publicar struct{}

// Publica el post al resto de los usuarios
func (c Publicar) RealizarPublicacion(l Login, id_global *int, parametro string, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post], usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	if !l.Conectado {
		fmt.Println(errores.ErrorUsuarioNoLoggeado{})
		return
	}
	// Creo su post pasando por parametros la cantidad de post que hizo(me sirve para tener conteo con el id), el texto, y el nombre de quien lo creo
	nuevo_Post := TDAUSUARIO.CrearPost(*id_global, parametro, l.User)
	posteos.Guardar(*id_global, nuevo_Post)
	(*id_global)++
	// obtengo el TDAusuario que esta actualmente loggeado
	usuario_loggeado := usuarios_registrados.Obtener(l.User)

	faux.ActualizaFeedDelResto(l.User, usuario_loggeado.DevolverAfinidad(), nuevo_Post, usuarios_registrados)
	fmt.Println(c.ConfirmarPublicacion())
}

// Muestra al usuario los posteos que se crearon
func (c Publicar) Feed(l Login, usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	if !l.Conectado {
		fmt.Println(errores.ErrorFindFeed{})
		return
	}
	usuario_loggeado := usuarios_registrados.Obtener(l.User)
	err := usuario_loggeado.ScrollFeed()
	if err != nil {
		fmt.Println(err)
	}
}

// Informa al usuario que el post se publico
func (c Publicar) ConfirmarPublicacion() string {
	return "Post publicado"
}

type Likes struct{}

// Le suma un like al post
func (c Likes) LikearPost(parametro string, l Login, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post]) {
	id, err := strconv.Atoi(parametro)
	if err != nil {
		fmt.Println(errores.ErrorParametros{})
		return
	}
	if !posteos.Pertenece(id) || !l.Conectado {
		fmt.Println(errores.ErrorPostInexistente{})
		return
	}
	post := posteos.Obtener(id)
	post.LikearPost(l.User)
	fmt.Println(c.ConfirmarLike())
}

// Muestra los likes que tiene el post hasta el momento
func (c Likes) MostrarLikes(parametro string, posteos TDADICC.Diccionario[int, TDAUSUARIO.Post]) {
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

// Informa al usuario que el post se likeo
func (c Likes) ConfirmarLike() string {
	return "Post likeado"
}
