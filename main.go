package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// Counter Created Counter struct.
// added sync.Mutex to handle race condition.
// mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
type Counter struct {
	sync.Mutex
	counter int
}

func (cs *Counter) get(w http.ResponseWriter, req *http.Request) {
	log.Printf("get received: %v", cs.counter)
	fmt.Fprintf(w, "got: %d\n", cs.counter)
}

func (cs *Counter) set(w http.ResponseWriter, req *http.Request) {
	cs.Lock()
	defer cs.Unlock()
	log.Printf("set %v", req)
	val := req.URL.Query().Get("val")
	intval, err := strconv.Atoi(val)

	if err != nil {
		panic("unhandled error")
	}

	cs.counter = intval
	log.Printf("set to: %v", cs.counter)

}

func (cs *Counter) inc(_ http.ResponseWriter, _ *http.Request) {
	cs.Lock()
	defer cs.Unlock()
	cs.counter++
	fmt.Println("incremented to", cs.counter)
	log.Printf("incremented to: %v", cs.counter)
}

func main() {
	count := Counter{counter: 0}
	http.HandleFunc("/get", count.get)
	http.HandleFunc("/set", count.set)
	http.HandleFunc("/increment", count.inc)

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
		log.Printf("Going  %d\n", os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}
