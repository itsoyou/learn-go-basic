package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface does not define method, but simply define certain method exists!
// unlike other languages, you don't have to explicitly connect the interface to struct
type saver interface {
	Save() error
}

// embedding existing interface
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1.5)
	printSomething(100)
	printSomething("whatever")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
	// no need to exit when error, this is the end of the program anyway
}

// Accepts any type of input as value
func printSomething(value interface{}){ // Also you can use "any"

	// instead of Swith, to check the type of value
 	intVal, ok := value.(int)
	if ok {
		fmt.Println(intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println(floatVal)
		return
	}



	switch value.(type) {
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)
	}
	// Any other value will accepted, but nothing will happen.
}

func outputData(data outputtable) error{
	data.Display()
	return saveData(data)
}

func saveData(data saver) error{
	err := data.Save()
	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded.")
	return nil
}

func getNoteData() (string, string){
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ", prompt)
	// "Scanln" is for single word input
	// var value string
	// fmt.Scanln(&value)

	// for longer text this is better way
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // stop reading from linebreaker. For byte is single quote.
	if err != nil {
		return ""
	}
	// now we have to remove the linebreak from the text
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows linebreak is \r\n
	return text
}