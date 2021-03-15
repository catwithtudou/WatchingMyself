package etcd

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"os"
	"strings"
	"testing"
)

// 可以设置一些参数，比如节点ID
var (
	nodeID    = flag.Int("id", 0, "node ID")
	addr      = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
	electName = flag.String("name", "my-test-elect", "election name")
)

func TestLeaderChoose(t *testing.T) {
	flag.Parse()

	// 将etcd的地址解析成slice of string
	endpoints := strings.Split(*addr, ",")

	// 生成一个etcd的clien
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// 创建session,如果程序宕机导致session断掉，etcd能检测到
	session, err := concurrency.NewSession(cli)
	defer session.Close()

	// 生成一个选举对象。下面主要使用它进行选举和查询等操作
	// 另一个方法ResumeElection可以使用既有的leader初始化Election
	e1 := concurrency.NewElection(session, *electName)

	// 从命令行读取命令
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		switch action {
		case "elect": // 选举命令
			go elect(e1, *electName)
		case "proclaim": // 只更新leader的value
			proclaim(e1, *electName)
		case "resign": // 辞去leader,重新选举
			resign(e1, *electName)
		case "watch": // 监控leader的变动
			go watch(e1, *electName)
		case "query": // 查询当前的leader
			query(e1, *electName)
		case "rev":
			rev(e1, *electName)
		default:
			fmt.Println("unknown action")
		}
	}
}

var count int

func elect(e1 *concurrency.Election, electName string) {
	log.Println("acampaigning for ID:", *nodeID)
	if err := e1.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

func proclaim(e1 *concurrency.Election, electName string) {
	log.Println("proclaiming for ID:", *nodeID)
	if err := e1.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}

func resign(e1 *concurrency.Election, electName string) {
	log.Println("resigning for ID:", *nodeID)
	if err := e1.Resign(context.TODO()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}

func watch(e1 *concurrency.Election, electName string) {
	ch := e1.Observe(context.TODO())

	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}

func query(e1 *concurrency.Election, electName string) {
	resp, err := e1.Leader(context.Background())
	if err != nil {
		log.Printf("failed to get the current leader: %v", err)
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

func rev(e1 *concurrency.Election, electName string) {
	rev := e1.Rev()
	log.Println("current rev:", rev)
}
