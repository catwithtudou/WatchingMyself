package lb

import (
	"log"
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

var(
	BackendAliveTimeOut = 2 * time.Second
)

type Backend struct{
	URL *url.URL
	Alive bool
	Mux sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (b *Backend)SetAlive(alive bool){
	b.Mux.Lock()
	defer b.Mux.Unlock()
	b.Alive = alive
}

func (b *Backend)IsAlive()bool{
	b.Mux.RLock()
	defer b.Mux.RUnlock()
	alive:=b.Alive
	return alive
}


func isBackendAlive(u *url.URL)bool{
	conn,err:=net.DialTimeout("tcp",u.Host,BackendAliveTimeOut)
	if err!=nil{
		log.Printf("failed to dial url;err:%s\n",err.Error())
		return false
	}
	_ = conn.Close()
	return true
}