package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	errores "algogram/errores"
	dicc "algogram/hash"
	"bufio"
	"os"
)

const (
	USUARIOS_TXT = 1
)

// Lee los usarios del archivo pasado por consola
func Leerarchivo(archi []string) (*os.File, error) {
	if len(archi) != USUARIOS_TXT {
		return nil, errores.ErrorParametros{}
	}
	archivo, err := os.Open(archi[0])
	return archivo, err
}

// Escanea los usuarios devolviendo un diccionario y los errores correspondientes si los hay,
// donde la clave y el valor del diccionario son el nombre de los usuarios y su TDA correspondiente
func Escaneararchivo(usuarios *os.File, usuarios_registrados *dicc.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	scanner := bufio.NewScanner(usuarios)
	posicion := 1
	for scanner.Scan() {
		usuario := TDAUSUARIO.CrearUsuario(posicion)
		(*usuarios_registrados).Guardar(scanner.Text(), usuario)
		posicion++
	}
}
