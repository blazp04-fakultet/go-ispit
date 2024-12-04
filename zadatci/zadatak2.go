package zadatci

import "fmt"

type Employee struct {
	Name     string
	Position string
	Tasks    []string
}

func (e *Employee) AddTask(task string) {
	e.Tasks = append(e.Tasks, task)
}

func (e *Employee) PrintAllTasks() {
	fmt.Printf("\n-----------[SVI ZADATCI - %s]-----------\n", e.Name)

	for _, task := range e.Tasks {
		fmt.Println(task)
	}
}

func (e *Employee) CompleateTask() {
	fmt.Printf("\n-----------[ZADATAK ZAVRšEN - %s]-----------\n", e.Name)

	for _, task := range e.Tasks {
		fmt.Printf("Employee %s compleated task: %s \n", e.Name, task)
		e.Tasks = e.Tasks[1:]
	}

}
