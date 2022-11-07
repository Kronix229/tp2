package comandos

import "fmt"

// Un pequeño tda que maneja si ya hay alguien loggeado y saluda al usuario al registrarse
type Login struct {
	User      string
	Conectado bool
}

func (c Login) Saludar() string {
	return fmt.Sprintf("Hola %s", c.User)
}

func (c *Login) EstadoDelUsuario() {
	if c.Conectado {
		c.Conectado = false
	} else {
		c.Conectado = true
	}
}

// Despide al usuario
type Logout struct{}

func (c Logout) Despedir() string {
	return "Adios"
}
