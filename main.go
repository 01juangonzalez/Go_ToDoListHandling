package main

import "fmt"

type TaskList struct {
	Tasks []Task
}

type Task struct {
	taskTitle  string
	taskDesc   string
	taskStatus bool
}

func (l *TaskList) CreateTask(t Task) {
	/*	InsertTask
		  		Fields:
					taskTitle: string
					taskDesc: string
					taskStatus: Open by default, not requested
				For the moment, only one word per item
	*/
	l.Tasks = append(l.Tasks, t)
	//{taskTitle, taskDesc, true}
	fmt.Printf("Task created! \t Title: %s \t Description: %s \t Status: Open \n", t.taskTitle, t.taskDesc)

}

func (l *TaskList) MarkDeletionTask(index int) {
	l.Tasks[index-1].taskStatus = false
	fmt.Printf("Item %d -> %s marked successfully!", index, l.Tasks[index-1].taskTitle)
}

func (l *TaskList) DebugList() {

	//temporalList := make([]Task, len(l))
	//temporalList := make([]TaskList, 5, 10)
	//var destLower, destUpper, srcLower int

	for i, t := range l.Tasks {

		fmt.Println(t)
		if !(l.Tasks[i].taskStatus) {
			//temporalList = append(temporalList, l.Tasks[i])
			//l.Tasks = append(l.Tasks[0:1], l.Tasks[2:]...)
			//l.Tasks = append(l.Tasks[destLower:destUpper],l.Tasks[srcLower:]...)
		}
		/*	destLower++
			destUpper++
			srcLower++
		*/
	}
	//fmt.Printf("%d, %d, %d", destLower, destUpper, srcLower)

	//copy(l.Tasks, temporalList)

	//l.Tasks = append(l.Tasks[0:1], l.Tasks[2:]...)

	fmt.Println("List debugging completed!")
}

//}
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
