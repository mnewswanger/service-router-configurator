package main

import (
	"log"

	"github.com/coreos/etcd/client"
	"github.com/fatih/color"
	"golang.org/x/net/context"
)

func publish(config configData, haproxyConfig string) {
	cfg := client.Config{Endpoints: []string{"http://etcd.kubernetes.home.mikenewswanger.com:2379"}}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)
	_, err = kapi.Set(context.Background(), config.HaproxyEtcdKey, haproxyConfig, nil)
	if err != nil {
		if err == context.Canceled {
			color.Red("Etcd request was cancelled")
			panic(err)
		} else if err == context.DeadlineExceeded {
			color.Red("Etcd deadline exceeded")
			panic(err)
		} else if cerr, ok := err.(*client.ClusterError); ok {
			color.Red("Etcd client error")
			panic(cerr.Errors)
		} else {
			color.Red("Failed to connect to etcd endpoint")
			panic(err)
		}
	}
}
