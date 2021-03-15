package etcd

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	recipe "github.com/coreos/etcd/contrib/recipes"
	"log"
	"os"
	"strings"
	"testing"
)


var (
	// addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	queueName = flag.String("name", "my-test-queue", "queue name")
)


func TestQueue(t *testing.T){
	flag.Parse()

	endpoints:=strings.Split(*addr,",")

	cli,err:=clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err!=nil{
		log.Fatal(err)
	}
	defer cli.Close()

	q:=recipe.NewQueue(cli,*queueName)

	consoleScanner:= bufio.NewScanner(os.Stdin)
	for consoleScanner.Scan(){
		action:=consoleScanner.Text()
		items:=strings.Split(action," ")
		switch items[0] {
		case "push":
			if len(items)!=2{
				fmt.Println("the push value must exist")
				continue
			}
			q.Enqueue(items[1])
		case "pop":
			v,err:=q.Dequeue()
			if err!=nil{
				log.Fatal(err)
			}
			fmt.Println(v)
		case "quit","exit":
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
