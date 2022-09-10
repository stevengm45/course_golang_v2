package main

import "net/http"

type Server struct {
	//puerto del servidor para escuchar las conexiones
	port string
	//se agrega el atributo router que es un apuntador al struct Router de router.go
	router *Router
}

//Sirve para instanciar el servidor y que sea capaz de escuchar las conexiones
//recive el puerto que tiene que estar escuchando y devuelve el servidor como tal
func NewServer(port string) *Server {
	return &Server{

		port: port,
		//router instanciado
		router: NewRouter(),
		//con esto el servidor ya es capaz de instanciar el router y de tenerlo como propiedad
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//Funcion tipo receiver, del struct Server, devuelve un error en caso de que haya problemas al conectar
func (s *Server) Listen() error {

	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}
	//si la ejecucion salio bien, retorna un valor nil
	return nil
}
