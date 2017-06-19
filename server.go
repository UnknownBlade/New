package main

import "net"
import "fmt"
import "bufio"
import "os"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "time"
import "strings"
import "encoding/json"

//struct for data in mongoDB
//mongoDB got variable id and variable first_name
type student struct {
	ID	bson.ObjectId 	`bson:"_id,omitempty"`
	id  string 			`bson:"id"`
	first_name	string 	`bson:"first_name"`
}

func main() {

	fmt.Println("Launching the server...")

	//duration of server and mongoDB connect(trying to connect by default but not succeded)
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout("localhost:27017", maxWait) //local host for mongoDB(default)
	if err != nil {
			panic(err)
	}
	defer session.Close()

	a := session.DB("StudentList").C("student")		//DB stand for database, C stand for collection
	if err != nil{
    	fmt.Println(err)
    	os.Exit(1)
    }

	//listen on all interfaces
	listen, _ := net.Listen("tcp", ":8081")		//port for client

	// accept connection on port
	conn, _ := listen.Accept()

	//(infinite lopp) ctrl-c to stop
	for{
    // will listen for the message
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//output message received
		fmt.Print("Message Received:", string(message))
		//sample process for string received
		newmessage := strings.TrimSpace(message)
		//send back to client

		if newmessage == "2016728301"{
			
			var display []student //struct
			err = a.Find(bson.M{"id": "2016728301"}).All(&display)
			if err != nil {
				panic(err)
			}
			outpt, err := json.Marshal(display)		//json is used to call function Marshal which contain data from struct
			if err != nil {
				panic (err)
			}
			conn.Write([]byte(string(outpt) + "\n"))

		}else if newmessage == "2016728302"{
			
			var display []student  //struct
			err = a.Find(bson.M{"id": "2016728302"}).All(&display)
			if err != nil {
				panic(err)
			}
			
			outpt, err := json.Marshal(display)		//json is used to call function Marshal which contain data from struct
			if err != nil {
				panic (err)
			}
			conn.Write([]byte(string(outpt) + "\n"))

		}else if newmessage == "2016728303"{
			
			var display []student  //struct
			err = a.Find(bson.M{"id": "2016728302"}).All(&display)
			if err != nil {
				panic(err)
			}
			
			outpt, err := json.Marshal(display)
			if err != nil {
				panic (err)
			}
			conn.Write([]byte(string(outpt) + "\n"))
		}
		if err != nil{
    	fmt.Println(err)
    	os.Exit(1)
    } 
  }
}
