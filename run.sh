#!/bin/bash
# 定义应用组名
group_name='golang'
#定义应用名称
app_name='go-private-proxy'
echo ${app_name}
echo '----copy app----'
docker stop ${app_name}
echo '----stop container----'
docker rm ${app_name}
echo '----rm container----'
# 打包编译docker镜像
docker build -t ${group_name}/${app_name} .
#构建docker应用
docker run -p 8138:8138 --name ${app_name} \
-e TZ="Asia/Shanghai" \
-d ${group_name}/${app_name}