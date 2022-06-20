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
docker run -p 8138:8138
--add-host=http://documentcloud.github.com:204.232.175.78 \
--add-host=http://github.com:207.97.227.239 \
--add-host=http://gist.github.com:204.232.175.94 \
--add-host=http://help.github.com:107.21.116.220 \
--add-host=http://nodeload.github.com:207.97.227.252 \
--add-host=http://raw.github.com:199.27.76.130 \
--add-host=http://status.github.com:107.22.3.110 \
--add-host=http://www.github.com:207.97.227.243 \
--name ${app_name} \
-e TZ="Asia/Shanghai" \
-d ${group_name}/${app_name}