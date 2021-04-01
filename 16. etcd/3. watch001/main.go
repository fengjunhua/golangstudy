package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"reflect"
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

	watcher := client.Watch(context.Background(),"test")
	fmt.Println(watcher)
	for changer := range watcher{
		fmt.Println(changer)
		for ek,ev := range changer.Events{

			fmt.Println(ek)
			fmt.Println("========================")
			fmt.Println(reflect.TypeOf(ev))
			fmt.Println("========================")
			fmt.Println(reflect.ValueOf(ev))
			fmt.Println("========================")
			fmt.Printf("%+v",ev)
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			fmt.Println(ev.Kv.Version,ev.Kv.CreateRevision,ev.Kv.Lease)
		}
	}

}
