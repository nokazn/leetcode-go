# !/usr/bin/env make -f

.DEFAULT_GOAL := help

.PHONY: test
test: # Run test codes
	go test ./problems/**

.PHONY: help
help: # Show all commands
	@echo
	@echo "📗 Displays help information for make commands."
	@echo
	@echo "Commands:"
# コマンド一覧 -> ":" で改行 -> ":" を含む行 (前半) の \s を ", " に置換、"#" を含む行 (後半) からコメントを抽出 -> ":" で分けた個所を再連結 -> column で整形
	@grep -E '^[a-zA-Z]\S+(\s\S+)*:.*' ./Makefile \
		| sed -E -e "s/:/:\n/" \
		| sed -E -e "/:/ s/\s/, /g" -e "s/^.*[#|;]+\s*//" \
		| sed -E -e "N" -e "s/:\n/:/g;" -e "s/^/  /" \
		| column -s ":" -t
