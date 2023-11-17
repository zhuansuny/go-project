package service

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	ProductStatusNormal       = 0
	ProductStatusSaldOut      = 1
	ProductStatusForceSaldOut = 2
)

//redis相关
type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

//etcd相关
type EtcdConf struct {
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

//用户访问限制相关
type AccessLimitConf struct {
	IPSecAccessLimit   int
	UserSecAccessLimit int
	IPMinAccessLimit   int
	UserMinAccessLimit int
}

//秒杀总配置
type SecSkillConf struct {
	RedisBlackConf       RedisConf //黑名单redis配置
	RedisProxy2LayerConf RedisConf //
	RedisLayer2ProxyConf RedisConf

	EtcdConf EtcdConf //etcd配置

	LogPath  string //日志路径
	LogLevel string //日志级别

	SecProductInfoMap map[int]*SecProductInfoConf //秒杀商品信息
	RWSecProductLock  sync.RWMutex                //读写锁
	CookieSecretKey   string                      //cookie密钥

	AccessLimitConf AccessLimitConf //用户访问限制结构体

	ipBlackMap map[string]bool //ip黑名单
	idBlackMap map[int]bool    //id黑名单

	blackRedisPool       *redis.Pool //黑名单redis数据库池
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool

	secLimitMgr *SecLimitMgr //

	ReferWhiteList []string

	RWBlackLock                  sync.RWMutex
	WriteProxy2LayerGoroutineNum int //写的协程个数
	ReadProxy2LayerGoroutineNum  int //读的协程个数

	SecReqChan     chan *SecRequest
	SecReqChanSize int

	UserConnMap     map[string]chan *SecResult
	UserConnMapLock sync.Mutex
}

type SecProductInfoConf struct {
	ProductId int   //商品id
	StartTime int64 //秒杀开始时间
	EndTime   int64 //秒杀结束时间
	Status    int   //抢购状态
	Total     int   //秒杀总量
	Left      int   //数量剩余
}

type SecRequest struct {
	ProductId     int
	Source        string      //来源
	AuthCode      string      //生成的校验码
	SecTime       string      //时间
	Nance         string      //随机数
	UserId        int         //用户id
	UserAuthSign  string      //用户cookie码签名（md5加密） 包含密钥和userId
	AccessTime    time.Time   //接受时间
	ClientAddr    string      //用户地址
	ClientRefence string      //refence 从哪个页面过来
	CloseNotify   <-chan bool `json:"-"`

	ResultChan chan *SecResult `json:"-"` //抢购结果
}

type SecResult struct {
	ProductId int
	UserId    int
	Code      int
	Token     string
}
