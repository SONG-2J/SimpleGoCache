package main

import (
	"SimpleGoCache/simplegocache"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// 模拟键值对数据库
var db = map[string]string{
	"P_A": "100",
	"P_B": "99",
	"P_C": "89",
	"P_D": "80",
	"P_E": "59",
}

func createGroup() *simplegocache.Group {
	return simplegocache.NewGroup("scores", 2<<10, simplegocache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

// 启动缓存服务器
func startCacheServer(addr string, addrs []string, gee *simplegocache.Group) {
	peers := simplegocache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("[SimpleGoCache] is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

// 启用API服务
func startAPIServer(apiAddr string, gee *simplegocache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "SimpleGoCache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}
