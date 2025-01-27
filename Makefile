.PHONY: gen
gen:
	@scripts/gen.sh ${svc}

.PHONY: gen-server
gen-server:
	# @cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/bz-2021/gomall/app/${svc} --pass "-use github.com/bz-2021/gomall/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/${svc}.proto
