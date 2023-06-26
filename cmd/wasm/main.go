package main

import (
	"fmt"
	"svelteGo/jsGo/document"
	"svelteGo/jsGo/vars"
	"syscall/js"
	"time"
)

func GoEventHandler(doc *document.JsDoc) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		uniqueId := args[0]
		action := args[1]
		event := args[2]
		element := doc.GetElementById(uniqueId.String())
		element.Action(action.String(), event)
		return nil
	})
}

func main() {
	fmt.Println("Go Web Assembly")

	counter := 0
	doc := document.Get()
	input := doc.GetElementById("test")
	button := doc.GetElementById("button")
	messageDiv := doc.GetElementById("message")

	test := vars.Create(counter)
	input.Bind("value", &test)

	button.On("click", func(event any) {
		counter += 1
		test.Update(counter)
	})

	button.On("click", func(event any) {
		messageDiv.Set("innerHTML", "clicked")
		time.Sleep(time.Second * 2)
		messageDiv.Set("innerHTML", "")
	})

	js.Global().Set("GoEventHandler", GoEventHandler(&doc))
	<-make(chan bool)
}
