package redis
import (
	"gopkg.in/redis.v4"
	"sync"
	"time"
	"encoding/json"

	"fmt"
)

type RedisS struct {
	sync.RWMutex
	Conn *redis.Client
}

func main()  {
	r,_:=CreateClient(0,"192.168.43.11:6379","Root1q2w")
	var RR RedisS
	RR.Conn = r
	RR.StringSet("key",1)
}

// 创建 redis 客户端
func CreateClient(dbint int, Addr, passwd string) (r *redis.Client, err error) {

	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: passwd,
		DB:       dbint,
		PoolSize: 15,
	})

	_, err = client.Ping().Result()
	//fmt.Println(pong)
	if err != nil {
		//err = errors.New(fmt.Sprintf("redis conn %s ,passwd :%s fail,err:%s \n", Addr, passwd, err))
		return
	}

	return client, nil
}

func (c *RedisS) StringExists(key string) bool {
	client := c.Conn

	bool1 := client.Exists(key)
	return bool1.Val()
}

func (c *RedisS) StringSet(key string, value int) (err error) {
	c.Lock()
	defer c.Unlock()
	client := c.Conn
	err = client.Set(key, value, 0).Err()
	if err != nil {
		return
	}
	return
}

func (c *RedisS) StringGet(key string) (val string, err error) {
	c.Lock()
	defer c.Unlock()
	client := c.Conn
	val, err = client.Get(key).Result()
	if err != nil {
		return
	}
	return
	//fmt.Println("name", val)
}

func (c *RedisS) SetJson(key string, value interface{}, ex time.Duration)(err error) {
	client := c.Conn

	b, err := json.Marshal(value)
	fmt.Println(string(b))
	if err != nil {
		//logger.Printf("key %s convert json faile data:%+v,err:%s\n", key, value, err)
		return
	}
	//fmt.Println(b,e)
	//fmt.Println("=========================================")
	//log.Printf("存入redis json %s\n", string(b))

	c.Lock()
	defer c.Unlock()

	err = client.Set(key, b, ex).Err()
	if err != nil {
		//logger.Printf("key %s 存入redis json 失败 data:%+v, ERR: %s\n", key, value, err)
		return
	}
	return nil
}

//
func (c *RedisS) GetJson(key string ,h interface{}) (err error){
	client := c.Conn

	c.Lock()
	defer c.Unlock()

	value, err := client.Get(key).Result()
	if err != nil {
		//log.Printf("key %s output redis faile!\n", key)

		return
	}
	json.Unmarshal([]byte(value), h)
	//fmt.Println("==============================")
	//fmt.Println(value)
	//fmt.Println(h)
	//fmt.Println("==============================")
	return

}

func (c *RedisS) Expire(key string, time time.Duration) bool {
	client := c.Conn

	c.Lock()
	defer c.Unlock()

	b := client.Expire(key, time)
	return b.Val()

}

func (c *RedisS) SetAdd(key, val string) (bool) {

	client := c.Conn

	c.Lock()
	defer c.Unlock()
	i := client.SAdd(key, val)
	//fmt.Println("=================",i.Val())
	if i.Val() == 1 { // 1代表成功添加，里面没有此UUID,
		//log.Printf("uuid:%s, 任务ID:%s 加入redis 任务组合 成功！", val, key)
		return true
	} else {
		//log.Printf("uuid:%s, 任务ID:%s 加入redis 任务组合 失败！", val, key)

		return false
	}

}

func (c *RedisS) SetInMember(key, val string) (bool) {
	client := c.Conn

	c.Lock()
	defer c.Unlock()

	b := client.SIsMember(key, val)
	return b.Val()

}

func (c *RedisS) SetRemMember(key, appname string, ex time.Duration) bool {
	client := c.Conn

	c.Lock()
	defer c.Unlock()
	//err := client.Set("name", "xys", time.Duration(10)*time.Second).Err()

	err := client.SRem(key, appname).Err()
	if err != nil {
		//log.Printf("uuid %s 存入redis app_Set 失败 data:%+v, ERR: %s\n", key, appname, err)
		return false
	}
	return true
}
