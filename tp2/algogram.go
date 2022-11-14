package main

import (
	TDAUSUARIO "algogram/Usuario"
	c "algogram/comandos"
	errores "algogram/errores"
	faux "algogram/funciones_aux"
	TDADICC "algogram/hash"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// Lo puse para compara, probablemente habra que cambiarle el nombre
	PARAMETROS_NECESARIOS = 1
)

func main() {
	usuarios, err := faux.Leerarchivo(os.Args[1:])
	// muestro el error si los hay
	if err != nil {
		fmt.Print(err)
		return
	}
	// Cambie el hash, en vez de guardad nombre y posicion, guardo el nombre y como dato el tda usuario
	// habra que cambiar algo porque estoy guardando dos veces el nombre(como clave y en el struct)
	usuarios_registrados := TDADICC.CrearHash[string, TDAUSUARIO.Usuario[string]]()
	faux.Escaneararchivo(usuarios, &usuarios_registrados)
	defer usuarios.Close()
	scanner := bufio.NewScanner(os.Stdin)
	login := c.Login{User: "", Conectado: false}
	for scanner.Scan() {
		comando, parametro, _ := strings.Cut(scanner.Text(), " ")
		comando, parametro = string(comando), string(parametro)
		switch comando {
		case "Login":
			parametros := strings.Split(parametro, " ")
			if !faux.ControlDeParametros(len(parametros), PARAMETROS_NECESARIOS) {
				fmt.Println(errores.ErrorParametros{})
				continue
			}
			login.User = parametro
			err = faux.DevolverErrorAlLoggearse(usuarios_registrados, login.User, login.Conectado)
			if err != nil {
				fmt.Println(err)
				continue
			}
			login.EstadoDelUsuario()
			fmt.Println(login.Saludar())
		case "Logout":
			if !login.Conectado {
				fmt.Println(errores.ErrorUsuarioNoLoggeado{})
				continue
			}
			fmt.Println(c.Logout{}.Despedir())
			login.EstadoDelUsuario()
		case "Publicar":
			if !login.Conectado {
				fmt.Println(errores.ErrorUsuarioNoLoggeado{})
				continue
			}
			// obtengo el usuario que esta actualmente loggeado
			usuario_loggeado := usuarios_registrados.Obtener(login.User)
			// Creo su post pasando por parametros la cantidad de post que hizo(me sirve para tener conteo con el id), el texto  su nombre
			nuevo_Post := TDAUSUARIO.CrearPost(usuario_loggeado.DevolverCantidadPost(), parametro, usuario_loggeado.DevolverNombre())
			// Al publicar le paso el usuario actual, todos los usuarios registrados, el post y la afinidad del usuario actual
			usuario_loggeado.PublicarPost(usuario_loggeado, usuarios_registrados, nuevo_Post, usuario_loggeado.DevolverAfinidad())
			fmt.Println(c.Publicar{}.ConfirmarPublicacion())
		case "Ver_siguiente_feed":
			if !login.Conectado {
				fmt.Println(errores.ErrorUsuarioNoLoggeado{})
				continue
			}
			usuario_loggeado := usuarios_registrados.Obtener(login.User)
			fmt.Println(usuario_loggeado.ScrollFeed())

		default:
			// Esto es algo estetico jaja
			fmt.Printf("Comandos disponibles: \n")
			for i := 0; i < 6; i++ {
				fmt.Println(faux.Devolvercomandos(i))
			}
		}
	}
}
