package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Faltan parametros."
}

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado."
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente."
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado."
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente."
}

type ErrorFindFeed struct{}

func (e ErrorFindFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver."
}

type ErrorSinLikes struct{}

func (e ErrorSinLikes) Error() string {
	return "Error: Post inexistente o sin likes."
}
