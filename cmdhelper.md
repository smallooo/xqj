### 根据端口查询进程号
sudo lsof -iTCP -sTCP:LISTEN -n -P
### 查询进程号关闭进程
kill -9 PID


### 重新起动进程
#!/bin/bash
#查找进程端口号
ss -nltp | grep id
#判断端口是否存在
Statu=`$?`
while true; do
#20秒运行
sleep 20
if [ ${Statu} == 0 ];then
#进程运行中

    else
    #启动进程

done