package comandos

import "fmt"

// Un pequeño tda que maneja si ya hay alguien loggeado y saluda al usuario al registrarse
type Login struct {
	User      string
	Conectado bool
}

// Saluda al usuario
func (c Login) Saludar() string {
	return fmt.Sprintf("Hola %s", c.User)
}

// Cambia a true o false si el usuario esta o no conectado
func (c *Login) EstadoDelUsuario() {
	if c.Conectado {
		c.Conectado = false
	} else {
		c.Conectado = true
	}
}

type Logout struct{}

// Despide al usuario
func (c Logout) Despedir() string {
	return "Adios"
}

type Publicar struct{}

// Informa al usuario que el post se publico
func (c Publicar) ConfirmarPublicacion() string {
	return "Post publicado"
}

type Likear struct{}

// Informa al usuario que el post se likeo
func (c Likear) ConfirmarLike() string {
	return "Post likeado"
}
