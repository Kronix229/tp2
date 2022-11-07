package main

import (
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
	usuarios_registrados := TDADICC.CrearHash[string, int]()
	faux.Escaneararchivo(usuarios, &usuarios_registrados)
	defer usuarios.Close()
	scanner := bufio.NewScanner(os.Stdin)
	login := c.Login{User: "", Conectado: false}
	for scanner.Scan() {
		parametros := strings.Split(scanner.Text(), " ")
		comando := parametros[0]
		switch comando {
		case "Login":
			if !faux.ControlDeParametros(len(parametros[1:]), PARAMETROS_NECESARIOS) {
				fmt.Println(errores.ErrorParametros{})
				continue
			}
			login.User = parametros[1]
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

		default:
			// Esto es algo estetico jaja
			fmt.Printf("Comandos disponibles: \n")
			for i := 0; i < 6; i++ {
				fmt.Println(faux.Devolvercomandos(i))
			}
		}
	}
}
