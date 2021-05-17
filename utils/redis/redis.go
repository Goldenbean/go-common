package redist

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v7"
)

type MyRedis struct {
	Client *redis.Client
}

type RateLimitParam struct {
	Threshold int64
	Period    time.Duration
}

// NewMyRedis : init
func NewMyRedis(addr string, password string) *MyRedis {

	r := &MyRedis{
		Client: createClient(addr, password),
	}

	return r
}

// RateLimit : Rate Limit
func (r *MyRedis) RateLimit(key string, limit RateLimitParam) (pass bool, err error) {
	client := r.Client
	_key := fmt.Sprintf("rate_limit:%s:%d", key, limit.Period)

	cmd := client.SetNX(_key, 1, limit.Period)

	ttl := client.TTL(_key)

	if ttl.Val() < 0 {
		fmt.Println("getR:", time.Now(), _key, limit.Period, ttl, cmd)
		time.Sleep(time.Duration(10000) * time.Millisecond)

		ttl := client.TTL(_key)
		fmt.Println("getR:", time.Now(), _key, limit.Period, ttl, cmd)

		if ttl.Val() < 0 {
			//panic("")
			client.Expire(_key, time.Second)
			ttl := client.TTL(_key)
			fmt.Println("EgetR:", time.Now(), _key, limit.Period, ttl)
		}
	}

	if cmd.Err() != nil {
		//fmt.Println("case 1:", cmd.Err())
		return true, cmd.Err()
	}

	if cmd.Val() { //first time
		//fmt.Println("case 2:", cmd.Val())
		return true, nil
	}

	incrCmd := client.Incr(_key)

	if incrCmd.Err() != nil {
		//fmt.Println("case 3:", cmd.Err())
		return true, cmd.Err()
	}

	//fmt.Println(limit.Threshold, incrCmd.Val())

	if incrCmd.Val() > limit.Threshold {
		//fmt.Println("case 4:")
		return false, nil
	}

	//fmt.Println("case 5:")
	return true, nil
}

func createClient(addr string, password string) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password, // no password set
		DB:           0,        // use default DB
		Network:      "tcp",
		PoolSize:     10,
		MinIdleConns: 5,
		DialTimeout:  100 * time.Millisecond,
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
		IdleTimeout:  10 * time.Second,
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}

	return client

}

func exampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.59.73.116:6379",
		Password: "B213547b69b13224", // no password set
		DB:       0,                  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
