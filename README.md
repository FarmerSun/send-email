## main.go

### Field need to set
```bash
sender_name   (发件者名字)       
sender_email  (发件者邮箱）
auth_email    (认证邮箱，一般同上)
auth_password (认证邮箱密码，QQ邮箱需自己开启，非登陆密码)
auth_host     (如 smtp.qq.com)
host          (如 smtp.qq.com:25)
subject       (邮件主题)
content       (邮件内容)
```

### Run Command
```go run main.go file_name```
***
**file_name** 为保存邮件信息的文件， 这里用sealyun/fetch拉下来的邮件格式
