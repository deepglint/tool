
#Hopper V1.0.0
```
实现face平台抓拍数据跳过网闸的功能
1.将kafka某个topic的信息写入FTP
2.从FTP读取信息写入到kafka某个topic
3.适配vsd1.1.0，matrix 0.8.0
4.修复没有删除ftp数据问题
```

##使用说明
```
1. kafkainputconfig里面配置外网接入系统的kafka信息
2. ftpoutputconfig里面配置外网ftp信息，用来写入数据（maxthread由于ftp限制暂时只能设置成1）
3. ftpinputconfig里面配置内网ftp信息，用来抓取数据（maxthread由于ftp限制暂时只能设置成1）
4. kafkaoutputconfig里面配置内网推送系统的kafka信息
5. speedlimit参数用来控制内网写入速度，外网写入速度没有限制，有就写
6. hopperout/in参数用来控制开启功能，外网部署该程序hopperout为true，公安内网部署改程序hopperin为true，同一网段内可以都配成true用来转发
7. config.json文件中新加了“Devices”字段，是一个map，用来存放内网接入的对应外网服务器的虚拟设备信息，key为外网服务设备的sensorid，value为虚拟设备的sensorid；
8. 外网服务器isd（v2.1.1 or later）配置信息中，"IsImageFormatURL”必须设置为false才会传图片，而不是url；
9. 内网服务器bingo（v2.6.x or later）配置信息中，”full_image_body", “cutboard_image_body"都必须设置为true，才会处理图片，而不是url；
10. 使用时现在内网服务器建立比对库，然后新建虚拟设备，然后将虚拟设备布控比对库，并将虚拟设备的sensorid写到hopper的配置文件中

#Hopper V0.2.0
```
实现face平台抓拍数据跳过网闸的功能
1.将kafka某个topic的信息写入FTP
2.从FTP读取信息写入到kafka某个topic
```

##使用说明
```
1. kafkainputconfig里面配置外网接入系统的kafka信息
2. ftpoutputconfig里面配置外网ftp信息，用来写入数据（maxthread由于ftp限制暂时只能设置成1）
3. ftpinputconfig里面配置内网ftp信息，用来抓取数据（maxthread由于ftp限制暂时只能设置成1）
4. kafkaoutputconfig里面配置内网推送系统的kafka信息
5. speedlimit参数用来控制内网写入速度，外网写入速度没有限制，有就写
6. hopperout/in参数用来控制开启功能，外网部署该程序hopperout为true，公安内网部署改程序hopperin为true，同一网段内可以都配成true用来转发
7. config.json文件中新加了“Devices”字段，是一个map，用来存放内网接入的对应外网服务器的虚拟设备信息，key为外网服务设备的sensorid，value为虚拟设备的sensorid；
8. 外网服务器isd（v2.1.1 or later）配置信息中，"IsImageFormatURL”必须设置为false才会传图片，而不是url；
9. 内网服务器bingo（v2.6.x or later）配置信息中，”full_image_body", “cutboard_image_body"都必须设置为true，才会处理图片，而不是url；
10. 使用时现在内网服务器建立比对库，然后新建虚拟设备，然后将虚拟设备布控比对库，并将虚拟设备的sensorid写到hopper的配置文件中
```


#Hopper V0.1.1
```
实现face平台抓拍数据跳过网闸的功能
1.将kafka某个topic的信息写入FTP
2.从FTP读取信息写入到kafka某个topic
```

##使用说明
```
1. kafkainputconfig里面配置外网接入系统的kafka信息
2. ftpoutputconfig里面配置外网ftp信息，用来写入数据（maxthread由于ftp限制暂时只能设置成1）
3. ftpinputconfig里面配置内网ftp信息，用来抓取数据（maxthread由于ftp限制暂时只能设置成1）
4. kafkaoutputconfig里面配置内网推送系统的kafka信息
5. speedlimit参数用来控制内网写入速度，外网写入速度没有限制，有就写
6. hopperout/in参数用来控制开启功能，外网部署该程序hopperout为true，公安内网部署改程序hopperin为true，同一网段内可以都配成true用来转发
7. config.json文件中新加了“Devices”字段，是一个map，用来存放内网接入的对应外网服务器的虚拟设备信息，key为外网服务器在ftp上的文件夹名称，value为虚拟设备的sensorid；
8. 外网服务器isd（v2.1.1 or later）配置信息中，"IsImageFormatURL”必须设置为false才会传图片，而不是url；
9. 内网服务器bingo（v2.6.x or later）配置信息中，”full_image_body", “cutboard_image_body"都必须设置为true，才会处理图片，而不是url；
10. 使用时现在内网服务器建立比对库，然后新建虚拟设备，然后将虚拟设备布控比对库，并将虚拟设备的sensorid写到hopper的配置文件中
```
