# Humanornot bot library for go
***Made by Lord_Jakub***
Hi, this is a library I made to make bots for humanornot.ai easier and I decided to post it because why not. 

Get this library by using `go get github.com/Lord-Jakub/humanornot-bot-lib`

Here is an example of a simple bot:

```go
  package  main

    import (
	    "fmt"
	    "os"
	    "time"
	    "github.com/Lord-Jakub/humanornot-bot-lib"
    )
    

    func  main() {
	//this is lsit of messages that bot will use
	    messages  := []string{
		    "Message",
		    "Other message"
	    }
	//repeat endlessly
	    for  true {
		    fmt.Println("Searching for partner")
		//create chat
		    chat, err  :=  humanornotbot.CreateChat()
		    if  err  !=  nil {
			    fmt.Println(err)
			    return
		    }
		    fmt.Println("Chat created")
		/check who will start
		    if  !chat.IsMyTurn {
			    fmt.Println("Waiting for partner to send message")
			//wait for other side to start
			    chat, err  =  humanornotbot.WaitMessage(chat.ChatID)
			    if  err  !=  nil {
				    fmt.Println(err)
				    return
			    }
			    if  len(chat.Messages) !=  0 {
				    fmt.Println("Other side: ", chat.Messages[0].Text)
			    }
		    }
		    i := 0
		//while chat is active send messages
		    for  chat.IsActive {
			    message := messages[i]
			    fmt.Println("Me: ", message)
			//send message
			    chat, err  =  humanornotbot.SendMessage(chat.ChatID, message)
			    if  err  !=  nil {
				    fmt.Println(err)
				    return
			    }
			    
			    fmt.Println("Other side: ", chat.Messages[len(chat.Messages)-1].Text)
			    i++
			    }
		    
		    //If you have some way to gues if it human or bot, use next commented line
		    //chat, err  =  humanornotbot.GuessChat(chat.ChatID, "human")
		}
	}
```
Note: if you don't want to use my userID, change humanornotbot.UserID variable.
