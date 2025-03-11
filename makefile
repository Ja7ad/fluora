gen_instructions:
	@for file in $(shell find instructions -type f -name "*.yml"); do \
		sub_dir=$$(basename $$(dirname $$file)); \
		name=$$(basename $$file .yml); \
		output_dir="internal/instructions/$$sub_dir"; \
		echo "Generating $$name.gen.go in $$output_dir..."; \
		mkdir -p $$output_dir; \
		go run cmd/instruction/main.go -instruction $$file -output $$output_dir -package $$sub_dir; \
	done

fmt:
	go tool gofumpt -l -w .

swagger:
	go tool swag init -g ./api/rest/server.go -o ./api/rest/docs
