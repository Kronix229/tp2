package usuario

import (
	TDAhash "algogram/hash"
	TDAHeap "algogram/heap"
)

// Usuario modela un usuario, con la capacidad de realizar las acciones permitidas a los usuarios en AlgoGram
type Usuario[K comparable] interface {
	//PublicarPost recibe una lista con todos los nombres de los usuarios, el hash donde estan guardados, el post a publicar y
	//la posicion del usuario que publica el post en archivo de texto y encola en una cola con prioridad el post con la
	//informacion de que tan relevante es el post para ese usuario de manera que en su feed vea posts en orden de relevancia.
	PublicarPost(usuario_loggeado Usuario[K], dicc TDAhash.Diccionario[K, Usuario[K]], post Post, afinidad int)

	//ScrollFeed muestra el proximo post en el feed del usuario. Si no le quedan posts en el feed devuelve el
	//correspondiente error
	ScrollFeed() error

	DevolverNombre() string

	DevolverAfinidad() int

	DevolverFeed() TDAHeap.ColaPrioridad[*postPrioridad]
}
