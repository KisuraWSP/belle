build:
	@go build

run:
	./simp_

version:
	./simp_ version

help:
	./simp_ help

exec_name:
	echo "executable file_name: ./simp_"

prog:
	@go run .

prog_ver:
	@go run . version


