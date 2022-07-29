package gohttpCounter

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type FileSystem interface {
// 	Open(name string) (http.File, error)
// }
// type Page struct {
// 	Title string
// 	Body  []byte
// }

// func ServerStart() {
// 	http.HandleFunc("/root", getRoot)
// 	http.HandleFunc("/hello", getHello)
// 	// http.HandleFunc("/index", returnIndex)
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

// func FileServer(root FileSystem) Handler{

// }

// func returnIndex(w http.ResponseWriter, r *http.Request) (string, error) {
// 	// title := r.URL.Path
// 	resp, _ := http.Get("index.html")
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return string(body), nil

// }
