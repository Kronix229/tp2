package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	c "algogram/comandos"
	TDADICC "algogram/hash"
	"bufio"
	"os"
	"strings"
)

// Crea y devuelve un escaner, un struc que servira para acceder al nombre del usuario y informara si alguien esta o no conectado,
// un diccionario que guarda como clave la id y como dato el posteo y finalmente crea la id
func Inicializarvariables() (*bufio.Scanner, c.Login, TDADICC.Diccionario[int, TDAUSUARIO.Post], int) {
	scanner := bufio.NewScanner(os.Stdin)
	login := c.Login{User: "", Conectado: false}
	posteos := TDADICC.CrearHash[int, TDAUSUARIO.Post]()
	id := 0
	return scanner, login, posteos, id
}

// Separa el parametro del comando de la entrada del usuario
func InicializarComandosyParametro(scanner *bufio.Scanner) (string, string) {
	comando, parametro, _ := strings.Cut(scanner.Text(), " ")
	comando, parametro = string(comando), string(parametro)
	return comando, parametro
}
