ROOT = "github.com/imag-er/wendingcup"

# 服务列表
SERVICES := submit board user

# 定义生成客户端和服务器端的通用规则
ROOTDIR:= $(abspath .)
define gen-rpc
	@echo "Generating $(1) client-side..."
	@cd rpc_gen && cwgo client --type RPC --idl $(ROOTDIR)/idl/$(1).proto -I $(ROOTDIR)/idl --service $(1) -module $(ROOT)/rpc_gen   # 修改为父级目录

	@echo "Generating $(1) server-side..."
	@cd app/ && mkdir -p $(1) && cd $(1) && cwgo server --type RPC --service $(1) -module $(ROOT)/app/$(1) --pass "-use $(ROOT)/rpc_gen/kitex_gen" -I $(ROOTDIR)/idl --idl $(ROOTDIR)/idl/$(1).proto

endef

# 生成所有服务的Makefile内容
.PHONY: gen-all
gen-all: $(addprefix gen-,$(SERVICES))
# 为每个服务生成Makefile内容
$(foreach service,$(SERVICES),$(eval gen-$(service): ; $(call gen-rpc,$(service)) ))
# 确保以下命令以制表符开头

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@rm -rf rpc_gen/* app/*

