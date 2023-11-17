package tailf //读取日志文件的内容
import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/hpcloud/tail"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct { //读取日志配置的结构体，由于本包要使用该结构体，因此从main包中迁移过来
	LogPath string `json:"logPath"`
	Topic   string `json:"topic"`
}

type TailObj struct { //TailObj结构体
	tail     *tail.Tail  //tail包的Tail结构体，可以用来调用tail包查看日志的方法
	conf     CollectConf //配置文件结构体
	status   int
	exitChan chan int
}

type TailObjMgr struct { //TailObjMgr结构体
	tailObjs []*TailObj
	MsgChan  chan *TextMsg

	Lock sync.Mutex
}

type TextMsg struct { //读取日志的字符串结构体
	Msg   string
	Topic string
}

var (
	tailObjMgr *TailObjMgr //声明
)

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.MsgChan
	return
}

func UpdateConfig(confs []CollectConf) (err error) {
	tailObjMgr.Lock.Lock()
	defer tailObjMgr.Lock.Unlock()

	for _, oneConf := range confs {
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}

		if isRunning {
			continue
		}

		createNewTask(oneConf)
	}

	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}

		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}

	tailObjMgr.tailObjs = tailObjs
	return
}

func createNewTask(conf CollectConf) {
	fmt.Println("开启了一个新的协程读取文件", conf.Topic)
	obj := &TailObj{
		conf: conf,
	}
	tails, err := tail.TailFile(conf.LogPath, tail.Config{ //
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, err)
		return
	}

	obj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFromTail(obj) //每一个CollectConf开启一个协程，进行日志文件的读取

	return
}

func InitTail(conf []CollectConf, chanSize int) (err error) {

	tailObjMgr = &TailObjMgr{ //声明一个tial对象管理结构体
		MsgChan: make(chan *TextMsg, chanSize), //make 给管道分配地址及空间
	}

	if len(conf) == 0 { //判断传入的CollectConf切片是否为空
		logs.Error("invalid config for log collect, conf:%v", conf)
		return
	}

	for _, v := range conf { //将切片中的CollectConf取出
		createNewTask(v)
	}

	return
}

func readFromTail(tailObj *TailObj) { //读取日志文件方法
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines: //读取日志文件的一行 ，循环读取
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			textMsg := &TextMsg{ //将该行数据写入到消息结构体中
				Msg:   line.Text,
				Topic: tailObj.conf.Topic,
			}

			tailObjMgr.MsgChan <- textMsg //将消息结构体加入管道，后面发送给main
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf:%v", tailObj.conf)
			fmt.Println("tail obj will exited, conf:%v", tailObj.conf)
			return
		}

	}
}
