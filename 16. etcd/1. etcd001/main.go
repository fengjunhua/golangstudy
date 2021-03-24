package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints: []string{"182.92.236.162:2379"},
		DialTimeout: 5 * time.Second,
	}
    client,_ := clientv3.New(config)
    defer client.Close()
    /*
    kv := clientv3.NewKV(client)
    ctx := context.Background()
    kv.Put(ctx,"/services/user","user1")
    */
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	put, err := client.Put(ctx, "test", "test")
	cancel()
	if err != nil{
		panic(err)
	}
    fmt.Println(put.Header.String())
	/*
	ctx, cancel := context.WithTimeout(context.Background(), 5)
	get, err := client.Get(ctx, "/services/test")
	cancel()
	if err != nil{
		panic(err)
	}

	fmt.Println(get)

	 */
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, "test")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("\n")
	fmt.Println(resp.Kvs)
	fmt.Println("\n")
	/*
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

	 */
}