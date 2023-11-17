package main

// Cookie是由浏览器维持的，存储在客户端的一小段文本信息，伴随着用户请求和页面在Web服务器和浏览器之间传递。
// 用户每次访问站点时，Web应用程序都可以读取cookie包含的信息。浏览器设置里面有cookie隐私数据选项，打开它，
// 可以看到很多已访问网站的cookies
import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) { //创建与读取cookie
	cookie, _ := r.Cookie("username")
	if cookie.Value == "" {
		expiration := time.Now()
		fmt.Println(expiration)
		expiration = expiration.AddDate(1, 0, 0)
		fmt.Println(expiration)
		cookie := http.Cookie{Name: "username", Value: "zsyy", Expires: expiration}
		http.SetCookie(w, &cookie)
	} else {
		fmt.Fprint(w, cookie.Value)
	}

}

func main() {
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
