install: go-build docker-build
go-build:
	mkdir -p bin
	cp `which kubectl` .
	cp /root/.kube/config .
	docker run  --rm -v `pwd`:/go registry.cn-huhehaote.aliyuncs.com/shengzhihao/go-build:v1  go build -o /go/bin/app /go/src/app
docker-build:
	docker build -t nodecheck:release-1.0.0 .
