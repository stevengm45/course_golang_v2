package main

import (
	"fmt"
)

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Nombre", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Nombre", tarea.descripcion)
		}
	}
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNOmbre(nombre string) {
	t.nombre = nombre
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	t1 := &task{
		nombre:      "completar mi curso de GO",
		descripcion: "completar mi curso de go de platzi en esta semana",
	}

	t2 := &task{
		nombre:      "completar mi curso de Python",
		descripcion: "completar mi curso de python de platzi en esta semana",
	}

	t3 := &task{
		nombre:      "completar mi curso de Docker",
		descripcion: "completar mi curso de Docker de platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	// fmt.Println(lista.tasks[1])
	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("Index", index, "nombre", tarea.nombre)
	}

	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()
	// fmt.Println(len(lista.tasks))
	// lista.eliminarDeLista(1)
	// fmt.Println(len(lista.tasks))
	// fmt.Println(t)
	// fmt.Printf("%+v\n", t)
	// t.marcarCompleta()
	// t.actualizarNOmbre("FInzalizar mi curso de go")
	// t.actualizarDescripcion("Completa mi curso de go antes")
	// fmt.Printf("%+v\n", t)

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Steven"] = lista

	t4 := &task{
		nombre:      "completar mi curso de Linux",
		descripcion: "completar mi curso de Linux de platzi en esta semana",
	}

	t5 := &task{
		nombre:      "completar mi curso de GCP",
		descripcion: "completar mi curso de GCP de platzi en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Felipe"] = lista2

	fmt.Println("Tareas Steven")
	mapaTareas["Steven"].imprimirLista()

	fmt.Println("Tareas Felipe")
	mapaTareas["Felipe"].imprimirLista()
}
