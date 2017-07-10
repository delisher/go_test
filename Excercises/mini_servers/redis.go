package main

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"log"
	"net"
	"strings"
)

type Command struct {
	Fields []string
	Result chan string
}

const (
	ExpirationTime = 0
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	commands := make(chan Command)
	go redisServer(commands, client)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(commands, conn)
	}

	// pong, err := client.Ping().Result()
	// fmt.Println("WAT", pong, "|||", err)
}

// func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd

func redisServer(commands chan Command, client *redis.Client) {
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 arguments"
			continue
		}

		fmt.Println("GOT COMMAND", cmd)

		switch cmd.Fields[0] {
		// GET <KEY>
		case "GET":
			key := cmd.Fields[1]
			val := client.Get(key).String()
			cmd.Result <- val

		// SET <KEY> <VALUE>
		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "EXPECTED VALUE"
				continue
			}
			key := cmd.Fields[1]
			val := cmd.Fields[2]
			client.Set(key, val, ExpirationTime)
			cmd.Result <- ""

		// DEL <KEY>
		case "DEL":
			keys := cmd.Fields[1:]
			client.Del(keys...)
			cmd.Result <- ""
		default:
			cmd.Result <- "INVALID COMMAND " + cmd.Fields[0] + "\n"
		}
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}

}
