package tcp

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

func TestHandler(t *testing.T) {
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:48879")
	if nil != err {
		t.Error(err)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if nil != err {
		t.Error(err)
	}

	go func() {
		for {
			// Listen for connections and let Handler handle them
			conn, err := listener.Accept()
			if err != nil {
				t.Error(err)
			}
			go Handle(conn)
		}
	}()

	c, err := net.Dial("tcp", "127.0.0.1:48879")
	if err != nil {
		t.Error(err)
	}

	c.Write([]byte{registerMessageType})
	buffer := bufferPool.Get().([]byte)
	bytes, err := c.Read(buffer)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Response to registration: Ack: [%c], ID: [%d]", buffer[0], binary.LittleEndian.Uint32(buffer[1:bytes]))

}
