package main
import(
"io"
"log"
"net/http"
)
func testHandler(w http.ResponseWriter,r *http.Request) {
   io.WriteString(w,"Hello World!")
}

func main(){
http.HandleFunc("/p1",helloHandler)
err := http.ListenAndServer(":5001",nil)
if err != nil{
log.Fatal("ListenAndServer: ",err.Error())
}
}
