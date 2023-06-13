package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var configFile = flag.String("f", "etc/ws.yaml", "the config file")

func main() {
	// 创建Redis Cluster客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer func() {
		err := client.Close()
		if err != nil {
			log.Println("关闭Redis连接出错:", err)
			return
		}
	}()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	context.WithDeadline()
	cancel()
	// 设置一个键值对
	err := client.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		log.Println("设置键值对出错:", err)
		return
	}

	// 获取一个键的值
	val, err := client.Get(ctx, "key1").Result()
	if err == redis.Nil {
		fmt.Println("键不存在")
	} else if err != nil {
		log.Println("获取键值出错:", err)
		return
	} else {
		fmt.Println("键值为:", val)
	}

	// 执行事务操作
	err = client.Watch(ctx, func(tx *redis.Tx) error {
		// 开启事务
		pipe := tx.Pipeline()

		// 在事务中执行多个命令
		pipe.Incr(ctx, "counter")
		pipe.Set(ctx, "key2", "value2", 0)

		// 执行事务
		_, err := pipe.Exec(ctx)
		return err
	}, "counter", "key2")
	if err != nil {
		log.Println("执行事务操作出错:", err)
		return
	}

	// 获取计数器的值
	counterVal, err := client.Get(ctx, "counter").Int64()
	if err != nil {
		log.Println("获取计数器值出错:", err)
		return
	}
	fmt.Println("计数器的值为:", counterVal)

}

//func main() {
//	res := countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})
//	println(res)
//	//flag.Parse()
//	//
//	//var c config.Config
//	//conf.MustLoad(*configFile, &c)
//	//c.Timeout = 0
//	//c.Verbose = false
//	//ctx := svc.NewServiceContext(c)
//	//server := rest.MustNewServer(c.RestConf)
//	//defer server.Stop()
//	//
//	//handler.RegisterHandlers(server, ctx)
//	//
//	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
//	//logx.MustSetup(logx.LogConf{
//	//	Encoding: "plain",
//	//})
//	//server.Start()
//}
