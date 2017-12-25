目录结构
--common
    --common_log:封装的日志记录库
    --myconobj:封装的玩家网络对象结构
    --mydb:数据库连接擦造作结构
    --net:protobuf消息文件 以及通用消息结构
--logdir
    --日志文件
--msg
    --c2php:客户端发向php的请求
    --c2smsg:客户端发向服务器的请求
    --s2cmsg:服务器给客户端的响应
--main
    --程序总入口，包括收发消息

命令行参数说明：
-d : 是否以守护进程方式运行,默认以非守护进程方式运行（true/false）
-begin : 登录消息发往PHP的openid起始,默认0,必须传
-end:登录消息发往php的openid末尾  end-begin:代表登录的玩家个数,必须传
-maxnum：程序中协程的数量 maxnum必须等于2*(end-begin)，必须传
-logdir:日志文件目录 可不传(默认："../../logdir/)
-num:每天保留的日志文件的数量，可不传（默认：10）
-cap:每个日志文件最大的保留的内容，可不传（1024*1024*50）
-days:最多保留几天的日志，可不传（默认3）

运行示例：

go run main.go -d=true -begin=10000 -end=10010 -maxnum=20