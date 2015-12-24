package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello! String")
    fmt.Fprintln(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Struct")
    fmt.Fprintln(w, s.Greeting, s.Punct, s.Who)
}


func main() {
    // your http.Handle calls here

    /* 
    * ListenAndServe, which is blocking. (BTW, if ListenAndServe didn't block, main would return and the process would exit)
    * Register the handlers before that. 
    */
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

    // 2번째 인자 handler를 넘기게 되면 위 handler 등록이 의미 없어짐.
    http.ListenAndServe("localhost:4000", nil)
    
}
