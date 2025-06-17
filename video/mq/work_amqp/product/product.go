package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
	"time"
)

func BodyFrom(args []string) string {
	fmt.Println(args)
	fmt.Println(len(args))
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "test"
	} else {
		s = strings.Join(args[1:], " ")
		fmt.Println(s)
	}
	return s
}
func fail(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	fail(err, "connect error")
	ch, err := conn.Channel()
	fail(err, "channel error")
	err = ch.ExchangeDeclare("logs", amqp.ExchangeFanout, true, false, false, false, nil)
	fail(err, "fanout error")
	//q, err := ch.QueueDeclare("test", true, false, false, false, nil)
	//fail(err, "failed to queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := BodyFrom(os.Args)
	err = ch.PublishWithContext(ctx, "logs", "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	if err != nil {
		fail(err, "send msg error")
	}
}

docker run \
-itd \
-p 2379:2379 \
-p 2380:2380 \
-e ALLOW_NONE_AUTHENTICATION=yes  \
--volume=./member0:/etcd-data \
--name etcd0 quay.io/coreos/etcd:v3.5.21 \
/usr/local/bin/etcd \
--data-dir=/etcd-data --name etcd-node0 \
--initial-advertise-peer-urls http://172.17.0.3:2380 --listen-peer-urls http://0.0.0.0:2380 \
--advertise-client-urls http://172.17.0.3:2379 --listen-client-urls http://0.0.0.0:2379 \
--initial-cluster  etcd-node0=http://172.17.0.3:2380,etcd-node1=http://172.17.0.4:2380,etcd-node2=http://172.17.0.5:2380 \
--initial-cluster-state new --initial-cluster-token my-etcd-token


docker run \
-itd \
-p 2379:2379 \
-p 2380:2380 \
-e ALLOW_NONE_AUTHENTICATION=yes  \
-e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
--volume=./member0:/etcd-data \
--name etcd0 bitnami/etcd \
/opt/bitnami/etcd/bin/etcd \
--data-dir=/etcd-data --name etcd-node0 \
--initial-advertise-peer-urls http://172.17.0.3:2380 --listen-peer-urls http://0.0.0.0:2380 \
--advertise-client-urls http://172.17.0.3:2379 --listen-client-urls http://0.0.0.0:2379 \
--initial-cluster  etcd-node0=http://172.17.0.3:2380,etcd-node1=http://172.17.0.4:2380,etcd-node2=http://172.17.0.5:2380 \
--initial-cluster-state new --initial-cluster-token my-etcd-token


docker network create --subnet=172.80.0.0/16 etcd-net

docker run -d --name etcd1 --net etcd-net --ip 172.80.0.2 \
-p 2379:2379 \
-e ALLOW_NONE_AUTHENTICATION=yes \
-e ETCD_ADVERTISE_CLIENT_URLS="http://172.80.0.2:2379" \ 		##列出此成员客户端 URL 以通告给集群的其余部分
-e ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379" \       		##用于监听客户端流量的 URL 列表
-e ETCD_LISTEN_PEER_URLS="http://0.0.0.0:2380" \         		##用于监听对等流量的 URL 列表
-e ETCD_INITIAL_ADVERTISE_PEER_URLS="http://172.80.0.2:2380" \  ##列出此成员对等 URL，以便在引导时向集群的其余部分进行通告
-e ETCD_NAME=etcd1 \
-e ETCD_INITIAL_CLUSTER="etcd1=http://172.80.0.2:2380,etcd2=http://172.80.0.3:2380,etcd3=http://172.80.0.4:2380" \
-e ETCD_INITIAL_CLUSTER_STATE=new \
bitnami/etcd:latest

docker run -d --name etcd1 --net etcd-net --ip 172.80.0.2 \
-e ALLOW_NONE_AUTHENTICATION=yes \
-p 2379:2379 \
-e ETCD_ADVERTISE_CLIENT_URLS="http://172.80.0.2:2379" \
-e ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379" \
-e ETCD_LISTEN_PEER_URLS="http://0.0.0.0:2380" \
-e ETCD_INITIAL_ADVERTISE_PEER_URLS="http://172.80.0.2:2380" \
-e ETCD_NAME=etcd1 \
-e ETCD_INITIAL_CLUSTER="etcd1=http://172.80.0.2:2380,etcd2=http://172.80.0.3:2380,etcd3=http://172.80.0.4:2380" \
-e ETCD_INITIAL_CLUSTER_STATE=new \
bitnami/etcd:latest


docker run -d --name etcd2 --net etcd-net --ip 172.80.0.3 \
-e ALLOW_NONE_AUTHENTICATION=yes \
-p 2381:2379 \
-e ETCD_ADVERTISE_CLIENT_URLS="http://172.80.0.3:2379" \
-e ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379" \
-e ETCD_LISTEN_PEER_URLS="http://0.0.0.0:2380" \
-e ETCD_INITIAL_ADVERTISE_PEER_URLS="http://172.80.0.3:2380" \
-e ETCD_NAME=etcd2 \
-e ETCD_INITIAL_CLUSTER="etcd1=http://172.80.0.2:2380,etcd2=http://172.80.0.3:2380,etcd3=http://172.80.0.4:2380" \
-e ETCD_INITIAL_CLUSTER_STATE=new \
bitnami/etcd:latest


docker run -d --name etcd3 --net etcd-net --ip 172.80.0.4 \
-e ALLOW_NONE_AUTHENTICATION=yes \
-p 2382:2379 \
-e ETCD_ADVERTISE_CLIENT_URLS="http://172.80.0.4:2379" \
-e ETCD_LISTEN_CLIENT_URLS="http://0.0.0.0:2379" \
-e ETCD_LISTEN_PEER_URLS="http://0.0.0.0:2380" \
-e ETCD_INITIAL_ADVERTISE_PEER_URLS="http://172.80.0.4:2380" \
-e ETCD_NAME=etcd3 \
-e ETCD_INITIAL_CLUSTER="etcd1=http://172.80.0.2:2380,etcd2=http://172.80.0.3:2380,etcd3=http://172.80.0.4:2380" \
-e ETCD_INITIAL_CLUSTER_STATE=new \
bitnami/etcd:latest