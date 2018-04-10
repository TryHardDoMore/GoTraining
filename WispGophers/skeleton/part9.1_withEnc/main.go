// Skeleton to part 9 of the Whispering Gophers code lab.
//
// This program extends part 8.
//
// It connects to the peer specified by -peer.
// It accepts connections from peers and receives messages from them.
// When it sees a peer with an address it hasn't seen before, it makes a
// connection to that peer.
// It adds an ID field containing a random string to each outgoing message.
// When it recevies a message with an ID it hasn't seen before, it broadcasts
// that message to all connected peers.
//
package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/campoy/whispering-gophers/util"
)

var (
	peerAddr = flag.String("peer", "", "peer host:port")
	self     string
)

var aeskey = []byte("AbsolutlyRandomKeyForAesEncryption")[:32]
var iv = []byte("532b6195636c6127")[:aes.BlockSize]

type Message struct {
	ID   string
	Addr string
	Body []byte
	Size int
}

func main() {
	flag.Parse()

	l, err := util.Listen()
	if err != nil {
		log.Fatal(err)
	}
	self = l.Addr().String()
	log.Println("Listening on", self)

	go dial(*peerAddr)
	go readInput()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(c)
	}
}

var peers = &Peers{m: make(map[string]chan<- Message)}

type Peers struct {
	m  map[string]chan<- Message
	mu sync.RWMutex
}

// Add creates and returns a new channel for the given peer address.
// If an address already exists in the registry, it returns nil.
func (p *Peers) Add(addr string) <-chan Message {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.m[addr]; ok {
		return nil
	}
	ch := make(chan Message)
	p.m[addr] = ch
	return ch
}

// Remove deletes the specified peer from the registry.
func (p *Peers) Remove(addr string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.m, addr)
}

// List returns a slice of all active peer channels.
func (p *Peers) List() []chan<- Message {
	p.mu.RLock()
	defer p.mu.RUnlock()
	l := make([]chan<- Message, 0, len(p.m))
	for _, ch := range p.m {
		l = append(l, ch)
	}
	return l
}

func broadcast(m Message) {
	for _, ch := range peers.List() {
		select {
		case ch <- m:
		default:
			// Okay to drop messages sometimes.
		}
	}
}

func serve(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)
	for {
		var m Message
		err := d.Decode(&m)
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: If this message has seen before, ignore it.
		if Seen(m.ID) {
			continue
		}
		decripted := make([]byte, m.Size)
		err = decryptAES(decripted, []byte(m.Body), aeskey, iv)
		if err != nil {
			panic(err)
		}
		if checkSpam(string(decripted)) {
			fmt.Println("SPAM!!!")
		} else {
			fmt.Printf("%s\n", string(decripted))
			broadcast(m)
			go dial(m.Addr)
		}
	}
}

func readInput() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := []byte(s.Text())
		encrypted := make([]byte, len(text))
		err := encryptAES(encrypted, text, aeskey, iv)
		if err != nil {
			panic(err)
		}
		m := Message{
			// TODO: use util.RandomID to populate the ID field.
			ID:   util.RandomID(),
			Addr: self,
			Body: encrypted,
			Size: len(text),
		}
		// TODO: Mark the message ID as seen.
		Seen(m.ID)
		broadcast(m)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func dial(addr string) {
	if addr == self {
		return // Don't try to dial self.
	}

	ch := peers.Add(addr)
	if ch == nil {
		return // Peer already connected.
	}
	defer peers.Remove(addr)

	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(addr, err)
		return
	}
	defer c.Close()

	e := json.NewEncoder(c)
	for m := range ch {
		err := e.Encode(m)
		if err != nil {
			log.Println(addr, err)
			return
		}
	}
}

// TODO: Create a new map of seen message IDs and a mutex to protect it.
var seenIDs = struct {
	m    map[string]bool
	lock sync.Mutex
}{m: make(map[string]bool)}

// Seen returns true if the specified id has been seen before.
// If not, it returns false and marks the given id as "seen".
func Seen(id string) bool {
	// TODO: Get a write lock on the seen message IDs map and unlock it at before returning.
	seenIDs.lock.Lock()
	// TODO: Check if the id has been seen before and return that later.
	ok := seenIDs.m[id]
	// TODO: Mark the ID as seen in the map.
	seenIDs.m[id] = true
	seenIDs.lock.Unlock()
	return ok
}

func encryptAES(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func decryptAES(dst, src, key, iv []byte) error {
	aesBlockEncryptor, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncryptor, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func checkSpam(message string) bool {
	fields := strings.Fields(message)
	if len(fields) < 1 {
		log.Println("Empty Message")
		return true
	} else if len(fields) == 1 {
		for sub := range getUniqSub(message) {
			if strings.Count(message, sub) > 2 {
				return true
			}
		}
	} else {
		for _, word := range strings.Fields(message) {
			if strings.Count(message, word) > 2 {
				return true
			}
		}
	}
	return false
}

func getUniqSub(message string) map[string]bool {
	substr := make(map[string]bool)
	for f := 0; f < len(message); f++ {
		for u := len(message); u > f; u-- {
			sub := message[f:u]
			_, ok := substr[sub]
			if !ok {
				substr[sub] = true
			}
		}
	}
	return substr
}
