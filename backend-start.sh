docker run --rm -it -d -p 8080:8080 -p 443:443 \
    --link postgresql:postgresql \
    -v /root/gin-log:/eventbackend/log \
    -e TZ=Asia/Shanghai \
    eventbackend:test