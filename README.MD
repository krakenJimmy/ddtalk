## 钉钉机器人

#### go发送http请求: [https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7](https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7)

### go build
[https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

### go build with busybox

[https://www.sohamkamani.com/blog/2016/11/22/docker-server-busybox-golang/](https://www.sohamkamani.com/blog/2016/11/22/docker-server-busybox-golang/)


### for linux
env GOOS=linux GOARCH=amd64 go build -o talk *.go


** for busybox **
env GOOS=linux GOARCH=386  CGO_ENABLED=0 go build -o talk *.go