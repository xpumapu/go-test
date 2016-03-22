package main

import "net"
import "fmt"
import "bufio"
import "strings" 
import "os"

const (  
    CONN_HOST = ""
    CONN_PORT = "3000"
    CONN_TYPE = "tcp"
)


func main() {

  fmt.Println("Start server...")

  // listen on all interfaces
  ln, err := net.Listen(CONN_TYPE, ":" + CONN_PORT)
  
  if err != nil {
    fmt.Println("Listen error:", err.Error())
    os.Exit(-1)
  }

  defer ln.Close()
  
  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Accept error:", err.Error())
      os.Exit(-1)
    }
  
    fmt.Println("client connected:", conn.RemoteAddr().String())
    
    go handle_rx(conn)
  }
  
}  
 
 
func handle_rx(conn net.Conn) { 
  // will listen for message to process ending in newline (\n)
  message, _ := bufio.NewReader(conn).ReadString('\n')
  // output message received
  fmt.Print("Message Received (", len(message), ") :", string(message))
  // sample process for string received
  newmessage := strings.ToUpper(message)
  // send new string back to client
  conn.Write([]byte(newmessage + "\n"))
  conn.Close()
}