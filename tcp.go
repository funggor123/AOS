package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	Number    int
	Timestamp string
	Date      int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

// bcServer handles incoming concurrent Blocks
var bcServer chan []Block
var mutex = &sync.Mutex{}

func main() {
	err := godotenv.Load("tcp.env")
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	// create genesis block
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	httpPort := os.Getenv("PORT")

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	io.WriteString(conn, "enter your new blockdate:")

	scanner := bufio.NewScanner(conn)

	// take in Date from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			Date, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], Date)
			if err != nil {
				log.Println(err)
				continue
			}
			if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\nenter your new blockdate:")
		}
	}()

	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			mutex.Lock()
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}

}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Number+1 != newBlock.Number {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
	mutex.Unlock()
}

func calculateHash(block Block) string {
	record := string(block.Number) + block.Timestamp + string(block.Date) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Date int) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Number = oldBlock.Number + 1
	newBlock.Timestamp = t.String()
	newBlock.Date = Date
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
