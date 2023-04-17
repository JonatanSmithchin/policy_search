package etcd_client

import (
	"context"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type HttpClient struct {
	Client        *clientv3.Client
	HttpEndpoints []string
}

var mutex sync.Mutex

func NewHttpClient(EtcdIp string, EtcdPort int, EtcdPrefix string) *HttpClient {
	keyName := EtcdPrefix + "/http"

	client := &HttpClient{}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdIp + ":" + strconv.Itoa(EtcdPort)},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	res, err := cli.Get(context.Background(), keyName)
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range res.Kvs {
		endPoints := kv.Value
		err := json.Unmarshal(endPoints, &client.HttpEndpoints)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
	if len(client.HttpEndpoints) <= 0 {
		log.Fatal("no http service")
	}

	client.Client = cli

	rch := cli.Watch(context.Background(), keyName)
	go func() {
		for watchResponse := range rch {
			for _, ev := range watchResponse.Events {
				mutex.Lock()
				err := json.Unmarshal(ev.Kv.Value, &client.HttpEndpoints)
				mutex.Unlock()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}()

	return client
}

func (c *HttpClient) GetHttpServIp() string {
	rand.Seed(time.Now().Unix())
	n := len(c.HttpEndpoints)
	return c.HttpEndpoints[rand.Intn(n)]
}
