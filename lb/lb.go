package lb

import (
	"log"
	"net/http"
	"time"
)

type Status int

const (
	Attempts Status = iota
	Retry
)

var(
	HealthCheckTime = 2 * time.Minute
)

func GetAttemptsFromContext(r *http.Request)int{
	if attempts,ok:=r.Context().Value(Attempts).(int);ok{
		return attempts
	}
	return int(Retry)
}

func GetRetryFromContext(r *http.Request)int{
	if retry,ok:=r.Context().Value(Retry).(int);ok{
		return retry
	}
	return int(Attempts)
}

func lb(w http.ResponseWriter,r *http.Request){
	attempts:=GetAttemptsFromContext(r)
	if attempts > 3{
		log.Printf("[lb]%s(%s) is terminating because of the max attempts\n",r.RemoteAddr,r.URL.Path)
		http.Error(w,"Service is not available",http.StatusServiceUnavailable)
		return
	}
	peer:=serverPool.GetNextPeer()
	if peer!=nil{
		peer.ReverseProxy.ServeHTTP(w,r)
		return
	}
	http.Error(w,"Service is not available",http.StatusServiceUnavailable)
	return
}

var serverPool ServerPool

func healthCheck(){
	t:=time.NewTicker(HealthCheckTime)
	for{
		select {
		case <-t.C:
			log.Println("Starting health check")
			serverPool.HealthCheck()
			log.Println("health check done")
		}
	}
}