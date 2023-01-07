build/download-schema:
	@go build -o bin/download-schema ./cmd/$(notdir $@)/...

download-schema: build/download-schema
	@./bin/download-schema

clean:
	rm bin/*
	rm -r proxmox/access/*
	rm -r proxmox/cluster/*
	rm -r proxmox/nodes/*
	rm -r proxmox/pools/*
	rm -r proxmox/storage/*
	rm -r proxmox/version/*

test:
	@go test $(TEST)
