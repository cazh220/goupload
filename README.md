# 上传服务员

1. 上传文件

URL： https://zyrs.cnty.cn:80/file/upload

Method: POST

参数： 

file  File
prj   string  项目  hw 环卫  fl 分类 hs 回收 yf 医废 cc 餐厨
tp    int   类型1 video 2 picture 3 others


2. 查看文件列表

URL： https://zyrs.cnty.cn:80/file/list

Method: GET

参数：

tp  int     类型
prj string  项目