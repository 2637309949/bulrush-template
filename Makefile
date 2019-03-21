all: target

target: dependence doc bin migration clean

dependence:
	@echo "\033[32mInstalling dependence\033[0m"
	@npm install apidoc --global
	@go get github.com/kardianos/govendor
	@go get github.com/json-iterator/go
	@govendor sync
bin:
	@echo "\033[32mBuilding bin file\033[0m"
	@go build -tags=jsoniter -o web .

doc:
	@echo "\033[32mBuilding doc file\033[0m"
	@apidoc -i routes/ -o ./assets/public/apidoc

migration:
	@echo "\033[32mMigrating files\033[0m"
	@rm -rf ./build
	@mkdir ./build
	@cp -rf Dockerfile ./build
	@cp -rf assets ./build
	@cp -rf conf ./build
	@cp -rf logs ./build
	@cp web ./build
clean:
	@echo "\033[32mClean files\033[0m"
	@rm -rf web

.PHONY: all target dependence doc bin migration clean 