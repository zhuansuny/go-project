# logsys
日志收集系统是用来收集各个服务器上指定项目的日志，将日志信息显示到web界面进行分析，并且可以在web界面修改要收集的日志路径，可以实时更改。整个系统主要包括三大部分，各部分可独立运行，降低耦合性：
1.使用beego的MVC框架实现的web日志配置管理平台
①从Mysql数据库中获取数据，展示在网页上；并可进行数据库的添加、删除、修改操作
②web服务启动时读取数据库中的配置信息发送给etcd配置中心
③.每进行一次删除、添加、修改配置，会将数据重新发送给etcd
2.日志收集客户端：
①程序启动读取受etcd配置信息，读取指定路径的日志
②监听etcd的key,可以通过更改etcd来实时更改读取的日志路径
③将读取的日志信息发送到kafka中
3.日志转发：
①从kafka中消费指定topic的日志内容
②监听etcd更改的配置，不断更新读取的topic信息
③将从kafka中读取的日志内容上传到elasticSearch上最后用Kibana显示到web界面