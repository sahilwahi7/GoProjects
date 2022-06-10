package main

import (
	"fmt"
	"log"
	"net/http"
)


func hellohandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method  is not supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(w,"hello")

}

func formhandler(w http.ResponseWriter, r *http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParserForm() err %v",err)
	}
	fmt.Fprintf(w,"Post request Successful")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"name=%s\n",name)
	fmt.Fprintf(w,"address=%s",address)
}
func main(){
	fileserver :=http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",hellohandler)

	fmt.Printf("Starting server on port 8080\n")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)

	}
}