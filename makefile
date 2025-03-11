ROOT = "github.com/imag-er/wendingcup"

# 服务列表
SERVICES := submit judge

# 定义生成客户端和服务器端的通用规则
define gen-rpc-client
	@echo "Generating $(1) client-side..."
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/$(1).proto --service $(1) -module $(ROOT)/rpc_gen -I ../idl/
endef

define gen-rpc-server
	@echo "Generating $(1) server-side..."
	@cd app/$(1) && cwgo server --type RPC --service $(1) -module $(ROOT)/app/$(1) --pass "-use $(ROOT)/rpc_gen/kitex_gen" -I ../../idl/ --idl ../../idl/$(1).proto
endef

# 生成所有服务的Makefile内容
.PHONY: gen-all
gen-all: $(addprefix gen-,$(SERVICES))
# 为每个服务生成Makefile内容
$(foreach service,$(SERVICES),$(eval gen-$(service): ; $(call gen-rpc-client,$(service)) $(call gen-rpc-server,$(service))))

# 清理目标
.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@rm -rf rpc_gen/kitex_gen app/*/kitex_gen
