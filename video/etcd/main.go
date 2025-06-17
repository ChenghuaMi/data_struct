package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2381", "localhost:2382"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
		// handle error!
	}
	cli.Put(context.Background(), "test", "test1")
	res, _ := cli.Get(context.Background(), "test")
	for _, v := range res.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}
	members, _ := cli.MemberList(context.Background())
	for _, v := range members.Members {
		fmt.Println(v.ID, v.Name, "client url:", v.ClientURLs, "peer url:", v.PeerURLs)
	}
	ress, _ := cli.Get(context.Background(), "mi", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	for _, v := range ress.Kvs {
		fmt.Printf("key:%s,val:%s \n", v.Key, v.Value)
	}
	//de, _ := cli.Delete(context.Background(), "mi", clientv3.WithPrefix())
	//fmt.Println(de.Deleted)
	//fmt.Println(len(de.PrevKvs))
	//for _, v := range de.PrevKvs {
	//	fmt.Println("xxxx")
	//	fmt.Printf("key:%s,val:%s \n", v.Key, v.Value)
	//}
	lease, _ := cli.Grant(context.Background(), 5)
	cli.Put(context.Background(), "name", "666", clientv3.WithLease(lease.ID))
	//_, err = cli.Revoke(context.Background(), lease.ID)
	//if err != nil {
	//	panic(err)
	//}
	go func() {
		wc := cli.Watch(context.Background(), "mi", clientv3.WithPrefix())
		for s := range wc {
			for _, v := range s.Events {
				fmt.Printf("%s,key:%s,val:%s \n", v.Type, v.Kv.Key, v.Kv.Value)
			}

		}
	}()
	//go func() {
	//	chanRes, _ := cli.KeepAlive(context.Background(), lease.ID)
	//	for {
	//		ka := <-chanRes
	//		fmt.Println(ka.ID, "ttl:", ka.TTL)
	//	}
	//
	//}()
	time.Sleep(time.Second * 60)
	defer cli.Close()
}
