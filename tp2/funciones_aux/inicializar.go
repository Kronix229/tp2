package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	TDADICC "algogram/hash"
	"bufio"
	"os"
	"strings"
)

// Crea y devuelve un escaner, un struc que servira para acceder al nombre del usuario y informara si alguien esta o no conectado,
// un diccionario que guarda como clave la id y como dato el posteo y finalmente crea una id global que sera el conteo de los posteos
func Inicializarvariables() (*bufio.Scanner, TDADICC.Diccionario[int, TDAUSUARIO.Post], int) {
	scanner := bufio.NewScanner(os.Stdin)

	posteos := TDADICC.CrearHash[int, TDAUSUARIO.Post]()
	id_global := 0
	return scanner, posteos, id_global
}

// Separa el parametro del comando de la entrada del usuario
func InicializarComandosyParametro(scanner *bufio.Scanner) (string, string) {
	comando, parametro, _ := strings.Cut(scanner.Text(), " ")
	comando, parametro = string(comando), string(parametro)
	return comando, parametro
}
