package main

import (
	TDAUSUARIO "algogram/Usuario"
	comandos "algogram/comandos"
	faux "algogram/funciones_aux"
	TDADICC "algogram/hash"
	"fmt"
	"os"
)

const (
	PARAMETROS_NECESARIOS = 1
)

func main() {
	usuarios, err := faux.Leerarchivo(os.Args[1:])
	// muestro el error si los hay
	if err != nil {
		fmt.Println(err)
		return
	}
	usuarios_registrados := TDADICC.CrearHash[string, TDAUSUARIO.Usuario[string]]()
	faux.Escaneararchivo(usuarios, &usuarios_registrados)
	defer usuarios.Close()
	login := comandos.Login{User: "", Conectado: false}
	scanner, posteos, id_global := faux.Inicializarvariables()
	for scanner.Scan() {
		comando, parametro := faux.InicializarComandosyParametro(scanner)
		switch comando {
		case "login":
			login.LoggearUsuario(parametro, usuarios_registrados)
		case "logout":
			comandos.Logout{}.Desloggear(&login)
		case "publicar":
			comandos.Publicar{}.RealizarPublicacion(login, &id_global, parametro, posteos, usuarios_registrados)
		case "ver_siguiente_feed":
			comandos.Publicar{}.Feed(login, usuarios_registrados)
		case "likear_post":
			comandos.Likes{}.LikearPost(parametro, login, posteos)
		case "mostrar_likes":
			comandos.Likes{}.MostrarLikes(parametro, posteos)
		}
	}
}
