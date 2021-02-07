package lb

import (
	"log"
	"net/url"
	"sync/atomic"
)

type HealthStatus int

const(
	Health HealthStatus = iota
	UnHealth
)

type ServerPool struct{
	Backends []*Backend
	Index uint64
}

func (s *ServerPool)AddBackend(backend *Backend){
	s.Backends = append(s.Backends,backend)
}

func (s *ServerPool)NextIndex()int{
	return int(atomic.AddUint64(&s.Index,uint64(1))%uint64(len(s.Backends)))
}

func (s *ServerPool)MarkBackendStatus(url *url.URL,alive bool){
	for _,b:=range s.Backends{
		if b.URL.String() == url.String(){
			b.SetAlive(alive)
			return
		}
	}
}

func (s *ServerPool)GetNextPeer()*Backend{
	backendLen:=len(s.Backends)
	next:=s.NextIndex()
	l:=backendLen+ next
	for i:=next;i<l;i++{
		idx:=i%backendLen
		if s.Backends[idx].IsAlive(){
			if i!=next{
				atomic.StoreUint64(&s.Index,uint64(idx))
			}
			return s.Backends[idx]
		}
	}
	return nil
}

func (s *ServerPool) HealthCheck(){
	for _,v:=range s.Backends{
		alive:=isBackendAlive(v.URL)
		v.SetAlive(alive)
		status:=Health
		if !alive{
			status=UnHealth
		}
		log.Printf("[HealthCheck]%s:%d\n",v.URL,status)
	}
}


