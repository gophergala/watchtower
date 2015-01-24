package users

import (
	"net/http"
)

type User interface {
	ID() uint32
	Send(sender uint32, content string)
}

// TODO: Use \n as keep-alive and send messages from incoming channel
/*
   fmt.Fprintf(res, "sending first line of data")
   if f, ok := res.(http.Flusher); ok {
      f.Flush()
   } else {
      log.Println("Damn, no flush");
   }
   sleep(10) //not real code
   fmt.Fprintf(res, "sending second line of data")
*/
type httpStreamUser struct {
	id uint32
	w  http.ResponseWriter
}
