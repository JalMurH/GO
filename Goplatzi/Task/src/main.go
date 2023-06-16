package main

import "fmt"

type task struct {
	name        string
	description string
	done        bool
}

type tasksList struct {
	tasks []*task
}

func main() {
	t := &task{
		name:        "Completar",
		description: "/",
		done:        false,
	}
	t2 := &task{
		name:        "curso",
		description: "un nuevo curso",
		done:        false,
	}
	fmt.Printf("%+v \n", t)

	t.taskDone()
	t.udateTaskDescription("Nueva descripcion de la tarea")
	t.updateTaskName("new task")

	fmt.Printf("%+v \n", t)

	list := &tasksList{
		tasks: []*task{t, t2},
	}
	fmt.Printf("%+v \n", list)
}

func (t *task) taskDone() {
	t.done = true
}

func (t *task) updateTaskName(name string) {
	t.name = name
}

func (t *task) udateTaskDescription(description string) {
	t.description = description
}

func (t *tasksList) addTask(tl *task) {
	t.tasks = append(t.tasks, tl)
}
