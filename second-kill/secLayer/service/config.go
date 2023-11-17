package service

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	clientV3 "go.etcd.io/etcd/client/v3"
)

var (
	secLayerContext = &SecLayerContext{}
)

// 秒杀商品结构体
type SecProductInfoConf struct {
	ProductId         int     //商品id
	StartTime         int64   //开始时间
	EndTime           int64   //结束时间
	Status            int     //秒杀状态
	Total             int     //一共有多少个
	Left              int     //剩余数量
	OnePersonBuyLimit int     //限制购买数量
	BuyRate           float64 //秒杀概率
	//每秒最多能卖多少个
	SoldMaxLimit int
	//限速控制
	secLimit *SecLimit
}

type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
	RedisQueueName   string
}

type EtcdConf struct {
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

type SecLayerConf struct {
	Proxy2LayerRedis RedisConf
	Layer2ProxyRedis RedisConf
	EtcdConfig       EtcdConf
	LogPath          string
	LogLevel         string

	WriteGoroutineNum      int
	ReadGoroutineNum       int
	HandleUserGoroutineNum int
	Read2handleChanSize    int
	Handle2WriteChanSize   int
	MaxRequestWaitTimeout  int

	SendToWriteChanTimeout  int
	SendToHandleChanTimeout int

	SecProductInfoMap map[int]*SecProductInfoConf
	TokenPasswd       string
}

type SecLayerContext struct {
	proxy2LayerRedisPool *redis.Pool
	layer2ProxyRedisPool *redis.Pool
	etcdClient           *clientV3.Client
	RWSecProductLock     sync.RWMutex

	secLayerConf     *SecLayerConf
	waitGroup        sync.WaitGroup
	Read2HandleChan  chan *SecRequest
	Handle2WriteChan chan *SecResponse

	HistoryMap     map[int]*UserBuyHistory
	HistoryMapLock sync.Mutex

	//商品的计数
	productCountMgr *ProductCountMgr
}

type SecRequest struct {
	ProductId     int
	Source        string
	AuthCode      string
	SecTime       string
	Nance         string
	UserId        int
	UserAuthSign  string
	AccessTime    time.Time
	ClientAddr    string
	ClientRefence string
	//CloseNotify   <-chan bool

	//ResultChan chan *SecResult
}

type SecResponse struct {
	ProductId int
	UserId    int
	Token     string
	TokenTime int64
	Code      int
}
