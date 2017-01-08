package main

import(
	"fmt"
	"github.com/husio/irc"
	"strings"
	"os"
	"time"
	"bufio"
	"flag"
)

var(
	address = "irc.freenode.net:6667"
	nick = "gobot"
	name = "gobot"
)

func main(){
        //making connection and sending essential details
	conn,err := irc.Connect(address)
	if(err!=nil){
		fmt.Println("Connection error occured")
		return
	}
	conn.Send("USER %s %s * :github.com/husio/irc",name,address)
        conn.Send("NICK %s",nick)

	time.Sleep(time.Millisecond*50)
	
	// joining to a freenode
	flag.Parse()
	for _,name := range flag.Args(){
		if !strings.HasPrefix(name,"#"){
			name = "#"+name
		}
		conn.Send("JOIN %s",name)
	}
	
	//user reply= read  from terminal and send to server
	reader := bufio.NewReader(os.Stdin)
	for{
		line,err := reader.ReadString('\n')
		if(err!=nil){
			fmt.Println("Unable to read from Terminal")
		}
		line = strings.TrimSpace(line)
		if len(line)==0{continue}
		conn.Send(line)
	}
	
	//gobot reply= read from server and echoing it
	for{
		message,err := conn.ReadMessage()
		if(err!=nil){
			fmt.Println("Unable to read From server")
			return
		}
	
		//echoing the recieved message
		if message.Command == "PRIVMSG"{
			if strings.HasPrefix(message.Trailing, nick){
				text:=message.Trailing[len(nick):]
				conn.Send("PRIVMSG %s: %s \"%s\"", message.Params[0], message.Nick(), text)
			}
		}	
	}	
}
