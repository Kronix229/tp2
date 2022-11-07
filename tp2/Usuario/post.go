package usuario

//Post modela una publicación, con la información pertinente a una publicación y la capacidad de actualizarse y mostrar
//la publicación según el input del usuario
type Post interface {
	//MostrarPost muestra el ID, el texto y cantidad de likes del post
	MostrarPost()

	//LikearPost actuliza la cantidad de likes y los usuarios que le dieron likes a un post
	LikearPost(usuario Usuario)

	//MostrarLikes muestra la cantidad de likes que tiene un post y los usuarios que le dieron like a dicho post
	MostrarLikes() error
}
