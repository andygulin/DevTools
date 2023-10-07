# DevTools 一些实用的开发工具

## 使用方法

```shell
go build
```

### time

```shell
# 输出当前时间-字符串
./DevTools time now
# 输出当前时间-时间戳
./DevTools time timestamp
# 格式化时间戳
./DevTools time format 1694170305
```

### uuid

```shell
# 获取UUID
./DevTools uuid
```

### hash

```shell
# 输出md5
./DevTools md5 Hello
# 输出sha1
./DevTools sha1 Hello
# 输出sha256
./DevTools sha256 Hello
# 输出sha512
./DevTools sha512 Hello
```

### base64

```shell
# 输入字符串，输出base64
./DevTools base64 text Hello
# 输入图片，输出base64
./DevTools base64 image tools/encode/base64/image.jpeg
```

### url

```shell
# URL Encode
./DevTools url encode 你好，世界。
# URL Decode
./DevTools url decode %E4%BD%A0%E5%A5%BD%EF%BC%8C%E4%B8%96%E7%95%8C%E3%80%82
```

### Jwt

```shell
# 创建Jwt，输入userId
./DevTools jwt create 30227
# 解析Jwt，返回userId
./DevTools jwt parse eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMwMjI3LCJleHAiOjE2OTY2Njc5NzF9.bWg87-7TTE7fmKBUrVWO6oivVeLqORI6ZCRhOvAG6ag
```

### number

```shell
# 输入2进制，输出8、10、16进制
./DevTools number 2 101011
# 输入8进制，输出2、10、16进制
./DevTools number 8 53
# 输入10进制，输出2、8、16进制
./DevTools number 10 43
# 输入16进制，输出2、8、10进制
./DevTools number 16 2b
```