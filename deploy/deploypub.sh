#!/opt/homebrew/bin/bash


# 交叉编译
rm -r deploy/output
for name in $(ls app); do
cd app/$name && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO_ENV=online go build -o $name
cd ../..
done 


# 编译完成后，复制到output目录
for name in $(ls app); do
    mkdir -p deploy/output/$name
    mv app/$name/$name deploy/output/$name
# 复制配置文件
    cp -r app/$name/conf deploy/output/$name
done

# 复制docker-compose.yml
cp docker-compose.yml deploy/output

# 打包
tar -czf deploy/output.tar.gz -C deploy/output .
