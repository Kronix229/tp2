package funcionesaux

// Recibe los parametros que el usuario ingreso y los parametros que el comando
// necesite,devolviendo un booleano
func ControlDeParametros(parametros, parametros_necesarios int) bool {
	return parametros >= parametros_necesarios
}
