package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
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
	return "Error: no habia usuario"
}

type ErrorPostInexistente struct{}

func (e ErrorPostInexistente) Error() string {
	return "Error: Usuario no loggeado o Post inexistente."
}

type ErrorFinFeed struct{}

func (e ErrorFinFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver."
}

type ErrorSinLikes struct{}

func (e ErrorSinLikes) Error() string {
	return "Error: Post inexistente o sin likes."
}
