package main

import "net"
import "fmt"
import "bufio"
import "strings"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  for { 


    // read input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Input student a ID: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")

    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    result := strings.Split(message, ",")
    for x := range result{
          fmt.Println(result[x])
  }
}
}
