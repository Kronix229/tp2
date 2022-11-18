package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	errores "algogram/errores"
	dicc "algogram/hash"
)

// Controla si el usuario ya esta registrado o si no existe, devolviendo el error correspondiente
func DevolverErrorAlLoggearse(usuarios_registrados dicc.Diccionario[string, TDAUSUARIO.Usuario[string]], usuario string, yalogueado bool) error {
	if !usuarios_registrados.Pertenece(usuario) {
		return errores.ErrorUsuarioInexistente{}
	}
	if yalogueado {
		return errores.ErrorUsuarioLoggeado{}
	}
	return nil
}
