package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type User struct {
	id int64
	name  string
  dob string
}

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is the Homepage")
}

func addUser(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
         http.ServeFile(w, r, "form.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }

        tt,_ := strconv.ParseInt(r.FormValue("id"), 0, 64)
        _,e := strconv.Atoi("wat")
        fmt.Println(e)

        var p = User{
          id : int64(tt),
          name: r.FormValue("name"),
          dob: r.FormValue("dob"),
        }
        //name := r.FormValue("name")
        //dob := r.FormValue("dob")
        fmt.Fprintf(w, "ID = %s\n", p.id)
        fmt.Fprintf(w, "Name = %s\n", p.name)
        fmt.Fprintf(w, "DOB = %s\n", p.dob)

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {

    //http.HandleFunc("adduser", addUser)
    http.HandleFunc("/",addUser)
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
