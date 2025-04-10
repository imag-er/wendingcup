ROOT = "github.com/imag-er/wendingcup"

# 服务列表
SERVICES := submit board user
CONTAINERS := api $(SERVICES) frontend judger

# 定义生成客户端和服务器端的通用规则
ROOTDIR:= $(abspath .)

define gen-rpc
	@echo "Generating $(1) client-side..."
	@cd rpc_gen && cwgo client --type RPC --idl $(ROOTDIR)/idl/$(1).proto -I $(ROOTDIR)/idl --service $(1) -module $(ROOT)/rpc_gen   # 修改为父级目录

	@echo "Generating $(1) server-side..."
	@cd app/ && mkdir -p $(1) && cd $(1) && cwgo server --type RPC --service $(1) -module $(ROOT)/app/$(1) --pass "-use $(ROOT)/rpc_gen/kitex_gen" -I $(ROOTDIR)/idl --idl $(ROOTDIR)/idl/$(1).proto
endef

define gen-docker
	docker build -f deploy/go.Dockerfile -t wendingcup-$(1) --build-arg SVC=$(1)  .
endef

define gen-docker-frontend
	docker build -f deploy/svelte.Dockerfile --format=docker -t wendingcup-frontend frontend 
endef

define gen-docker-api
	docker build -f deploy/go.Dockerfile -t wendingcup-api --build-arg SVC=api .	
endef

define gen-docker-judger
	docker build -f deploy/python.Dockerfile -t wendingcup-judger app/submit/python
endef



# 生成所有服务的Makefile内容
.PHONY: gen-all
gen-all: $(addprefix gen-,$(SERVICES))
# 为每个服务生成Makefile内容
$(foreach service,$(SERVICES),$(eval gen-$(service): ; $(call gen-rpc,$(service)) ))

.PHONY: gen-docker
gen-docker: $(addprefix gen-docker-,$(CONTAINERS))
$(foreach service,$(SERVICES),$(eval gen-docker-$(service): ; $(call gen-docker,$(service)) ))
$(eval gen-docker-frontend: ; $(call gen-docker-frontend))
$(eval gen-docker-judger: ; $(call gen-docker-judger))
$(eval gen-docker-api: ; $(call gen-docker-api))

# 导出所有container
.PHONY: save_images
save_images:
	docker save -o deploy/backend_image.tar $(CONTAINERS)
	docker save -o deploy/frontend_image.tar frontend:latest



.PHONY: clean
clean:
	rm -rf rpc_gen/*
	rm submit_files/*.zip
