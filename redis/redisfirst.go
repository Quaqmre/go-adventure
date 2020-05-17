package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis", // no password set
		DB:       0,       // use default DB
	})
	// err := client.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }

	// i, _ := client.RPush("mylist", "deneme123").Result()

	// client.RPush("mylist2", "{client:deneme}")
	// t, _ := client.LRange("mylist2", 0, 1).Result()
	// fmt.Println(t)
	// client.RPop("mylist2").Result() // 1 elemt çıkarıyor sağdan

	// client.LPush("mylist", "soldan eklendi").Result()
	// client.LTrim("mylist", 0, 2).Result() // son eklenen 3 item dışındakileri siler

	// l, _ := client.BRPop(time.Second*50, "mylist3").Result() // block the code if list null and wait when some item added mylist3

	type okul struct {
		Name string
	}
	// client.HSet("test", "asd", "qwe", "fff", "ggg")

	result := client.HExists("test", "asd")
	fmt.Println(result.Val())
	fmt.Println(client.HLen("test"))

	fmt.Println(client.HGetAll("test").Val())

	// fmt.Println(l)

}
