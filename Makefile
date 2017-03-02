prepare:
	glide install

run:
	go run main.go

build:
	go build -o ultronpym

buildx:
	for GOOS in $${GOOS_LIST:-darwin linux}; do \
		GOOS=$$GOOS GOARCH=amd64 go build -v -o ultronpym-$$GOOS-amd64 ; \
	done

docker-package:
	docker run -e GOOS_LIST="$${GOOS_LIST:=darwin linux}" --name $${PACKAGE_NAME:=ultronpym} $${BUILD_IMAGE:-ultronpym} bash -c 'make buildx'; \
	CONTAINER_ID=$$(docker ps -aqf "name=$$PACKAGE_NAME"); \
		for GOOS in $${GOOS_LIST:-darwin linux}; do \
			docker cp $$CONTAINER_ID:/go/src/github.com/wheniwork/ultronpym/ultronpym-$$GOOS-amd64 $$PACKAGE_NAME-$$GOOS ; \
			chmod +x $$PACKAGE_NAME-$$GOOS; \
		done ; \
		docker rm $$CONTAINER_ID
