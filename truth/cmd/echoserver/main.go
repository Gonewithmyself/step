package main

import "net"

func main() {
	ln, er := net.Listen("tcp", "127.0.0.1:8888")
	if er != nil {
		panic(er)
	}

	for {
		conn, er := ln.Accept()
		if er != nil {
			break
		}

		go func(conn net.Conn) {
			defer conn.Close()
			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					return
				}

				conn.Write(buf[:n])
			}
		}(conn)
	}
}
