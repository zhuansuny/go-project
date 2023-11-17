# log-transfer

功能:

① 程序刚开始运行时，从Etcd中,读取配置信息（主要是日志路径以及kafka的topic）

②实时监听etcd的key(/logAgent/config/127.0.0.1)的变化（主要是topic的变化），将变化结果返回给程序，再开启新的协程

③从kafka中读取指定topic的日志内容，并根据etcd更改的配置，不断更新读取的topic

④将从kafka中读取的日志内容上传到elasticSearch上
