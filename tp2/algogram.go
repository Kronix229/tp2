package main

import (
	TDAUSUARIO "algogram/Usuario"
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
	scanner, login, posteos, id_global := faux.Inicializarvariables()
	for scanner.Scan() {
		comando, parametro := faux.InicializarComandosyParametro(scanner)
		switch comando {
		case "login":
			faux.Login(usuarios_registrados, (&login), parametro)
		case "logout":
			faux.Logout(&login)
		case "publicar":
			faux.Publicar(login, &id_global, parametro, posteos, usuarios_registrados)
		case "ver_siguiente_feed":
			faux.Mostrar_Siguiente_Feed(login, usuarios_registrados)
		case "likear_post":
			faux.Likear_Post(parametro, login, posteos)
		case "mostrar_likes":
			faux.Mostrar_Likes(parametro, posteos)
		}
	}
}
