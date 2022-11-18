package funcionesaux

import (
	TDAUSUARIO "algogram/Usuario"
	TDADICC "algogram/hash"
)

// Recibe el nombre del usuario actual con su afinidad, el post que creo y un diccionario con todos los usuarios
// para iterar con cada uno y publicar en sus feed
func ActualizaFeedDelResto(nombre_del_usuario_actual string, afinidad_del_usuario_actual int, nuevo_post TDAUSUARIO.Post, usuarios_registrados TDADICC.Diccionario[string, TDAUSUARIO.Usuario[string]]) {
	usuarios_registrados.Iterar(func(clave string, usuario TDAUSUARIO.Usuario[string]) bool {
		if clave == nombre_del_usuario_actual {
			return true
		}
		usuario.PublicarPost(usuario, nuevo_post, afinidad_del_usuario_actual)
		return true
	})
}
