ifndef VERBOSE
.SILENT:
endif

NUMBERS = 1 2 3 4 5

build:
	go install -v ./...
	go build

container:
	docker build -t cmattoon/dir2cm:latest .

test: build
	mkdir -p configfiles
	$(foreach i,$(NUMBERS),echo "This is file $(i)" >> "configfiles/file_$(i).txt";)
	./dir2cm -dir configfiles -name my-configs

clean:
	rm -rf configfiles
