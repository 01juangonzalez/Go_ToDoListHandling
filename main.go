package main

import (
	"fmt"
	"slices"
)

type TaskList struct {
	Tasks []Task
}

type Task struct {
	taskTitle  string
	taskDesc   string
	taskStatus bool
}

/*
CreateTask

	  		Fields created:
				taskTitle: string
				taskDesc: string
				taskStatus: Open by default, not requested
			For the moment, only one word per item
*/
func (l *TaskList) CreateTask(t Task) {

	l.Tasks = append(l.Tasks, t)

	fmt.Printf("Task created! \t Title: %s \t Description: %s \t Status: Open \n", t.taskTitle, t.taskDesc)

}

/*
MarkDeletionTask

	Task selected to be deleted.
	List is debugged before marking task to delete
*/
func (l *TaskList) MarkDeletionTask(index int) {

	l.DebugList()
	l.Tasks[index-1].taskStatus = false
	fmt.Printf("Item %d -> %s marked successfully!", index, l.Tasks[index-1].taskTitle)
}

/*
MarkDeletionTaskDebugList

	Item already selected to be deleted is removed from list
*/
func (l *TaskList) DebugList() {

	for i, t := range l.Tasks {

		fmt.Println(t)
		if !(l.Tasks[i].taskStatus) {
			fmt.Println("Item deleted: ", i)
			l.Tasks = slices.Delete(l.Tasks, i, i+1)
			break

		}

		fmt.Println("List debugging completed!")
	}
}

/*
ListTasks

	List is displayed
*/
func (l *TaskList) ListTasks() {
	for i, t := range l.Tasks {
		fmt.Printf("%d %s %s %t \n", i+1, t.taskDesc, t.taskTitle, t.taskStatus)
	}
}

func main() {
	/*
		var newTask Task
		var itemsInArray []Task //TaskList
		//var addedItem []Task*/
	var taskSelection int

	var List TaskList
	var task Task
	var item int

	fmt.Println("Entering To Do List Management...")
	for {
		fmt.Println("Enter selection -> ")
		fmt.Scan(&taskSelection)
		switch taskSelection {
		case 1:
			var newTitle, newDesc string

			fmt.Printf("Insert Task Title: ")
			fmt.Scan(&newTitle)

			fmt.Printf("Insert Task Description: ")
			fmt.Scan(&newDesc)

			task.taskTitle = newTitle
			task.taskDesc = newDesc
			task.taskStatus = true

			List.CreateTask(task)

		case 2:
			fmt.Println("Mark task for deletion")
			fmt.Printf("Digit item number to mark: ")
			fmt.Scan(&item)
			List.MarkDeletionTask(item)

		case 3:
			fmt.Println("Debug mode list")
			List.DebugList()

		case 4:
			fmt.Println("To Do List status:")
			List.ListTasks()

		case 5:
			fmt.Println("Exiting system")

		default:
			//No task to be done
		}
		if taskSelection == 5 {
			break
		}
	}

	fmt.Println("Leaving To Do List session!")

}
