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
	client,err := clientv3.New(config)

	if err != nil{
		fmt.Println(err)
	} else{
		fmt.Println("connect etcd success!")
	}
	defer client.Close()

	//使用Grant方法创建一个租约
	leaseGrantResp, err := client.Grant(context.TODO(), 100)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(leaseGrantResp.ID)

	//创建一个key,并将一个租约赋加给这个key
	put, err := client.Put(context.TODO(), "name", "lbwnb", clientv3.WithLease(leaseGrantResp.ID))
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(put)
    time.Sleep(60 * time.Second)
	get, err := client.Get(context.TODO(), "name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(get.Kvs)
	time.Sleep(110 * time.Second)

	get2, err := client.Get(context.TODO(), "name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(get2.Kvs)

}
