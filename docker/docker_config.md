## go basic
- http://golang.site/go/basics
- 도커에서 실행시키도록 한다.
    - centos7 사용
```
docker run -it --name go_lang -v /Users/wshid/Workspace/docker/go_lang/:/root/workspace centos:latest
docker exec -it go_lang /bin/bash
```
- 이후 압축 해제 이후 $PATH에 추가

### wget install
```
yum install wget
yum upgrade
```