package main

import (
	"net/http"
	"text/template"
	"os"
	"fmt"
)

const port = ":8080"

func main() {
	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		requestFile := req.URL.Path[1:]
		fmt.Println(requestFile)
		lookup := templates.Lookup(requestFile + ".html")
		if lookup != nil {
			lookup.Execute(w,nil)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("404 - not found"))
		}

	})
	http.ListenAndServe(port,nil)
}

func populateTemplates() *template.Template {
	result :=  template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)

	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" +pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}


//
//func main() {
//	//http.ListenAndServe(":8080",http.FileServer(http.Dir("public")))
//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
//		w.Header().Add("Content-Type", "text/html")
//		templates := template.New("template")
//		templates.New("test").Parse(doc)
//		templates.New("header").Parse(header)
//		templates.New("footer").Parse(footer)
//		context := Context{
//			[3]string{"Lemon", "Orange", "Apple"},
//			"the title",
//		}
//		templates.Lookup("test").Execute(w,context)
//	})
//
//	http.ListenAndServe(":8080",nil)
//}
//
//const doc = `
//	{{template "header" .Title }}
//			<h1>List of fruits</h1>
//			<ul>
//				{{range .Fruit}}
//					<li>{{.}}</li>
//				{{end}}
//			</ul>
//    {{template "footer"}}
//`
//
//const header = `
//	<!DOCTYPE html>
//	<html>
//		<head>
//			<title>{{.}}</title>
//		</head>
//		<body> `
//
//const footer = `
//	</body></html>
//`
//
//type Context struct {
//	Fruit [3]string
//	Title string
//}

//
//func main (){
//	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//	//	w.Write([]byte("Hello World"))
//	//})
//
//	http.Handle("/", new(MyHandler))
//
//	fmt.Printf("Listening on port "+ port)
//	http.ListenAndServe(port ,nil)
//}


//type MyHandler struct {
//	http.Handler
//}
//
//func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
//	path := "public/" + req.URL.Path
//	//data, err := ioutil.ReadFile(string(path))
//	file, err := os.Open(path)
//
//	if err == nil {
//		bufferReader := bufio.NewReader(file)
//		var contentType string
//
//		if strings.HasSuffix(path,".css"){
//			contentType = "text/css"
//		} else if strings.HasSuffix(path,".html"){
//			contentType = "text/html"
//		} else if strings.HasSuffix(path,".js"){
//			contentType = "application/javascript"
//		}else if strings.HasSuffix(path,".png"){
//			contentType = "image/png"
//		}else if strings.HasSuffix(path,".woff"){
//			contentType = "font/woff"
//		} else if strings.HasSuffix(path,".woff2"){
//			contentType = "font/woff2"
//		} else if strings.HasSuffix(path,".appcache"){
//			contentType = "text/cache-manifest"
//		} else {
//			contentType = "text/plain"
//		}
//		w.Header().Add("Content-Type", contentType )
//		bufferReader.WriteTo(w)
//
//	} else {
//		w.WriteHeader(404)
//		w.Write([]byte("404 - " + http.StatusText(404)))
//	}
//}
