package funcionesaux

import (
	errores "algogram/errores"
	dicc "algogram/hash"
)

// Controla si el usuario ya esta registrado o si no existe devolviendo el error correspondiente
func DevolverErrorAlLoggearse(usuarios_registrados dicc.Diccionario[string, int], usuario string, yalogueado bool) error {
	if !usuarios_registrados.Pertenece(usuario) {
		return errores.ErrorUsuarioInexistente{}
	}
	if yalogueado {
		return errores.ErrorUsuarioLoggeado{}
	}
	return nil
}

// Muestra los comandos disponibles con sus respectivos parametros
func Devolvercomandos(indice int) string {
	comandos := []string{"Login ", "Logout ", "Publicar ", "Ver_siguiente_feed ", "Likear_post ", "Mostrar_likes "}
	parametros := []string{"<Nombre de usuario>", "", "<Texto>", "", "<Id del post>", "<Id del post>"}
	comando_parametro := "-" + comandos[indice] + parametros[indice]
	return comando_parametro
}
