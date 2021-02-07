package lb

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestLb(t *testing.T){
	var serverList string
	var port int
	flag.StringVar(&serverList,"backends","","Load balanced backends, use commas to separate")
	flag.IntVar(&port,"port",3030,"Port to serve")
	flag.Parse()

	if len(serverList)==0{
		t.Fatal("the server list is none")
		return
	}

	tokens:=strings.Split(serverList,",")
	for _,tok := range tokens{
		serverUrl,err:=url.Parse(tok)
		if err!=nil{
			t.Fatalf("failed to parse the url;err%s\n",err.Error())
			return
		}

		proxy:=httputil.NewSingleHostReverseProxy(serverUrl)
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			log.Printf("[error]%s:%s\n",serverUrl.Host,err.Error())
			retries:=GetRetryFromContext(request)
			if retries<3{
				select {
				case <-time.After(10 * time.Millisecond):
					ctx:=context.WithValue(request.Context(),Retry,retries+1)
					proxy.ServeHTTP(writer,request.WithContext(ctx))
				}
				return
			}

			serverPool.MarkBackendStatus(serverUrl,false)

			attempts:=GetAttemptsFromContext(request)
			log.Printf("[lb]%s(%s) attempt retry %d\n",request.RemoteAddr,request.URL.Path,attempts)
			ctx:=context.WithValue(request.Context(),Attempts,attempts+1)
			lb(writer,request.WithContext(ctx))
		}

		serverPool.AddBackend(&Backend{
			URL:          serverUrl,
			Alive:        true,
			ReverseProxy: proxy,
		})
		log.Printf("Configured server:%s\n",serverUrl)

	}

	server:=http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: http.HandlerFunc(lb),
	}

	go healthCheck()

	log.Printf("Load Balance started at [%d]\n",port)
	if err:=server.ListenAndServe();err!=nil{
		t.Fatalf(err.Error())
	}
}