
image=$(docker images | grep 'none' | awk '{print $3}')
if [ -n "$image" ]; then
  docker rmi $image
fi


cd ../cmd

# 更新 swagger
swag init --parseDependency --parseInternal

rm -rf cmd
# go clean -cache
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
echo "build"

mv cmd ../docker

cd ../docker

cp ../config .

dockerName=oss:$(TZ=CST-8 date '+%Y%m%d-%H%M')

docker build -t $dockerName .

echo $dockerName

rm -rf cmd

rm -rf config