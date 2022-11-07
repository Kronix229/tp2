package usuario

import TDAhash "algogram/hash"

// Usuario modela un usuario, con la capacidad de realizar las acciones permitidas a los usuarios en AlgoGram
type Usuario[K comparable, V any] interface {
	//PublicarPost recibe una lista con todos los nombres de los usuarios, el hash donde estan guardados, el post a publicar y
	//la posicion del usuario que publica el post en archivo de texto y encola en una cola con prioridad el post con la
	//informacion de que tan relevante es el post para ese usuario de manera que en su feed vea posts en orden de relevancia.
	PublicarPost(arr []string, dicc TDAhash.Diccionario[K, V], post Post, afinidad int)

	//ScrollFeed muestra el proximo post en el feed del usuario. Si no le quedan posts en el feed devuelve el
	//correspondiente error
	ScrollFeed() error
}
