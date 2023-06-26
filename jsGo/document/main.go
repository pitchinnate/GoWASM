package document

import (
	"fmt"
	"github.com/google/uuid"
	"svelteGo/jsGo/vars"
	"syscall/js"
)

type JsDoc struct {
	Document js.Value
	App      js.Value
	Elements map[string]Element
}

type Element struct {
	Id        string
	UniqueId  string
	Document  JsDoc
	Element   js.Value
	Listeners map[string][]func(any)
}

func Get() JsDoc {
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", "app")
	elements := make(map[string]Element)
	jsDoc := JsDoc{
		Document: doc,
		App:      element,
		Elements: elements,
	}
	return jsDoc
}

func (jsDoc *JsDoc) AddScript(script string) {
	jsElement := jsDoc.Document.Call("createElement", "script")
	jsElement.Set("innerHTML", script)
	jsDoc.App.Call("appendChild", jsElement)
}

func (jsDoc *JsDoc) GetElementById(id string) Element {
	foundElement, ok := jsDoc.Elements[id]
	if ok {
		return foundElement
	}
	element := jsDoc.Document.Call("getElementById", id)
	listeners := make(map[string][]func(any))
	uniqueId := uuid.New().String()
	newElement := Element{id, uniqueId, *jsDoc, element, listeners}
	jsDoc.Elements[id] = newElement
	return newElement
}

func (element *Element) Set(prop string, val any) {
	element.Element.Set(prop, val)
}

func (element *Element) Bind(prop string, myVar *vars.Variable) {
	myVar.OnUpdate(func(newVal any) {
		element.Element.Set(prop, newVal)
	})
}

func (element *Element) On(event string, myFunc func(any)) {
	_, ok := element.Listeners[event]
	if !ok {
		element.Listeners[event] = []func(any){}
		element.Document.AddScript(fmt.Sprintf(`
(() => {
const element = document.getElementById('%s');
element.addEventListener('%s', (event) => { GoEventHandler('%s', '%s', event); });
})();
`, element.Id, event, element.Id, event))
	}
	element.Listeners[event] = append(element.Listeners[event], myFunc)
}

func (element *Element) Action(event string, jsEvent js.Value) {
	fns, ok := element.Listeners[event]
	if !ok {
		return
	}
	for _, fn := range fns {
		go fn(jsEvent)
	}
}
