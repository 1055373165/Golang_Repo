代码运行

> go run worker_original.go -max_wokers 5 

测试（直接在 shell 中执行的 curl 脚本）

> for i in {1..15}; do curl localhost:8080/work -d name=job$i -d delay=$(expr $i % 9 + 1)s; done