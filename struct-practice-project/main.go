package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct-practice-project/note"
	"struct-practice-project/todo"
)

// 'saver' interface defines a contract for any type that can save its data.
// The type implementing this must define a 'Save' method returning an error.
type saver interface {
	Save() error
}

// type displayer interface{
// 	Display()
// }

// 'outputtable' is a composite interface that embeds 'saver' and adds 'Display'.
// Any type implementing 'outputtable' must satisfy both 'saver' and 'Display'.
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("test")

	title, content := getNoteData()

	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	printSomething(todo)

	err = outputData(todo)

	if err != nil {
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	// The 'value' parameter is of type 'interface{}',
	// which can hold any value of any type in Go.

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }

	/* Alternative way to check value type */

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Integer", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("Integer", stringVal)
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save data json.")
		return err
	}

	fmt.Println("Successfully saved data json.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
