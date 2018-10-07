# crawler
网络爬虫


# 存储/部署
Elastcisearch存储 采用Docker部署

# crawler_distribute 分布式服务
# crawler 本地服务

# 启动 Elasticsearch 服务器
docker run -d -p 9200:9200 imageid

# 启动并发单机服务
go run -v localserver.go

# 启动分布式服务
# 端口：9000 worker服务器
go run -v crawler_distribute/worker/workerRpcServer.go
# 端口：1234 saver服务器
go run -v crawler_distribute/persisr/saverrpcserver.go
# engine 服务器
go run -v crawler_distribute/server.go

# 启动后台服务
go run -v frontend/starter.go 端口：9999

# docker
https://download.daocloud.io/Docker_Mirror/Docker_for_Windows_Mac

# docker 加速器
https://www.daocloud.io/mirror

# elasticserch client
https://github.com/olivere/elastic

# goland IDE
https://www.jetbrains.com/go/?fromMenu
