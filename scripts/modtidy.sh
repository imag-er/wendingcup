# 定义服务名列表
service_names=("submit" "judge" "gateway")

# 遍历服务名列表
for service in "${service_names[@]}"; do
    cd "app/$service" && go mod tidy
    # 返回上一级目录，以便下一次循环可以正确进入下一个服务目录
    cd ../..
done