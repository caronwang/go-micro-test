package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getNode(srv []*registry.Service,f func ([]*registry.Service ) selector.Next) *registry.Node{
	next:=f(srv)
	node,err:=next()
	if err !=nil{
		log.Fatal(err)
	}
	return node
}


func random(srv []*registry.Service){
	for{
		next:=selector.Random(srv)
		node,err:=next()
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Println(node)
		time.Sleep(time.Second)
	}
}

func roundRobin(srv []*registry.Service){
	for{
		next:=selector.RoundRobin(srv)
		node,err:=next()
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Println(node)
		time.Sleep(time.Second)
	}
}


func callAPI(addr string, path string, method string) (string,error) {
	req,err := http.NewRequest(method,"http://"+addr+path,nil)
	client := http.DefaultClient
	res,err := client.Do(req)
	if err!=nil{
		return "",err
	}
	defer res.Body.Close()
	buf,_ := ioutil.ReadAll(res.Body)
	return string(buf),nil

}


func callAPI2()


func main(){
	csReg :=consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	getService,err :=csReg.GetService("prodService")
	if err!=nil{
		log.Fatal(err)
	}

	//random(getService)

	//oundRobin(getService)

	node := getNode(getService,selector.Random)
	callres,err := callAPI(node.Address ,"/v1/prod","GET")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(callres)
}