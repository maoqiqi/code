package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
	"math/rand"
	"net/http"
	"html/template"
)

var client *redis.Client

func main() {
	client = newRedis()

	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/code", codeHandle)
	http.HandleFunc("/verify", verifyHandle)

	fmt.Println("Server is at localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	checkErr(err)
}

// 创建 redis pool
func newRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]interface{})
	m["codeKey"] = getCodeKey()

	t, err := template.ParseFiles("index.html")
	checkErr(err)
	t.Execute(w, &m)
}

func codeHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// 得到传递过来地参数
	codeKey := r.Form.Get("codeKey")

	d := []byte(getCodeValue())
	codeValue := ""
	for v := range d {
		d[v] %= 10
		codeValue += strconv.FormatInt(int64(d[v]), 32)
	}
	// 保存到redis
	setCode(client, codeKey, codeValue)
	fmt.Println("生成验证码:", codeKey, "--", codeValue)

	w.Header().Set("Content-Type", "image/png")
	NewImage(d, 100, 40).WriteTo(w)
	fmt.Fprintf(w, codeKey)
}

func verifyHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	codeKey := r.PostForm.Get("codeKey")
	inputCode := r.PostForm.Get("inputCode")
	codeValue := getCode(client, codeKey)

	fmt.Println("验证验证码:", codeKey, "--", codeValue, "--", inputCode)

	if inputCode == codeValue {
		fmt.Println("正确")
		w.Write([]byte("正确"))
	} else {
		fmt.Println("错误")
		w.Write([]byte("错误"))
	}
}

func getCodeKey() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func getCodeValue() string {
	// 以当前纳秒数作为随机数种子
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%04v", rnd.Int31n(10000))
	return code
}

// 保存验证码到redis
func setCode(client *redis.Client, codeKey string, codeValue string) {
	err := client.Set(codeKey, codeValue, 5*60*time.Second).Err()
	checkErr(err)
}

// 得到验证码
func getCode(client *redis.Client, codeKey string) string {
	codeValue, err := client.Get(codeKey).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return codeValue
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
