docker run  --rm -p 8080:8080 -p 443:443 \
    --link postgresql:postgresql \
    -v /root/gin-log:/eventbackend/log \
    -e TZ=Asia/Shanghai \
    --entrypoint "./backend" \
    eventbackend:test