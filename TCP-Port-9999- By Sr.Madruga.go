//go 1.6

package main
import ( 
	"encoding/gob"
	"fmt"
    "net"
)

func Server(){
	In, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for{
		c, err := In.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}


func handleServerConnection(c net.Conn){
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Received", msg)
	}
	
	c.Close()
}


func Cliente(){
	c, err := net.Dial("tcp", "120.0.0.1:9999")
	if err != nil{
		fmt.Println(err)
		return
	}
	
	msg := "Bem Vindo ao Open Serve -> By Sr.Madruga"
	fmt.Println("sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	
	}
	
	c.Close()
}


func main(){
	go Server()
	go Cliente()
	
	var input string
	fmt.Scanln(&input)
	
	fmt.Println("-> Feito por Sr.Madruga")
}