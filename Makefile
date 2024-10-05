# Define the current date
DATE			:= $(shell date +%FT%T%z)
DIST_DIR		?= ./dist
DOC_DIR			?= ./doc
# Define the git tag
GIT_TAG			:= $(shell \
					git describe --tags --always --match=v* 2> /dev/null \
					|| echo "no-git-tag")
GOFLAGS			?= -trimpath -mod=readonly -modcacherw
INSTALL_PATH	?= /usr/local/bin
MAN_PATH		?= /usr/local/share/man/man1
MESSAGE_PREFIX= $(shell \
					if [ "$$(tput colors 2> /dev/null \
						|| echo 0)" -ge 8 ]; then \
							printf "\033[34;1m▶\033[0m" \
					else printf "▶"; \
					fi)
MODULE			= $(shell go list -m | sed 's|.*/||')
PKGS			= $(shell go list)
# PLATFORMS defines which platforms to build release binaries for.
PLATFORMS		:=	darwin/amd64 darwin/arm64 \
					linux/amd64 linux/arm64   \
					windows/amd64 windows/arm64
# Define dynamic variable reassignments.
# Define the semantic version of the application.
VERSION			:= v1.0.0
BUILD_DATE		:= ${DATE}
LD_VAR_PATH		:= github.com/pierow2k/polyhymnia/cmd
LDFLAGS			:= -s -w                                            \
				-X '${LD_VAR_PATH}.BuildDate=${BUILD_DATE}'         \
				-X '${LD_VAR_PATH}.Version=${VERSION}'

.PHONY: audit
audit: #@HELP Runs quality control checks
	$(info ${MESSAGE_PREFIX} Running quality control checks for ${MODULE})
	@echo "${MESSAGE_PREFIX} Running go mod verify:"
	go mod verify
	@echo "${MESSAGE_PREFIX} Running go vet:"
	go vet ./...
	@echo "${MESSAGE_PREFIX} Running staticcheck:"
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all ./...
	@echo "${MESSAGE_PREFIX} Running govulncheck:"
	go run golang.org/x/vuln/cmd/govulncheck@latest -show verbose ./...
	@echo "${MESSAGE_PREFIX} Removing test cache for ${MODULE}:"
	go clean -testcache
	@echo "${MESSAGE_PREFIX} Running go test:"
	go test -race -vet=off ./...


.PHONY: all
all: ${PLATFORMS} # @HELP Builds binaries for all platforms

${PLATFORMS}: man
	@${MAKE} GOOS=$(firstword $(subst /, ,$@)) GOARCH=$(lastword $(subst /, ,$@)) binary
	@cd ${DIST_DIR} && gsha256sum *.gz > sha256sums.txt


.PHONY: binary
binary: # @HELP Builds binary release in ${DIST_DIR} (Used Internally)
	@$(info ${MESSAGE_PREFIX} "Building ${MODULE}_${VERSION}_${GOOS}-${GOARCH}")
	@GOFLAGS="${GOFLAGS}" \
		go build \
			-ldflags="${LDFLAGS}" \
			-o ${DIST_DIR}/${MODULE}$(if $(findstring windows,${GOOS}),.exe,) \
			${PKGS}
	@cp ${DOC_DIR}/${MODULE}.1 ${DIST_DIR}/
	@tar \
		-czf ${DIST_DIR}/${MODULE}_${VERSION}_${GOOS}-${GOARCH}.tar.gz \
    	-C dist \
		${MODULE}$(if $(findstring windows,${GOOS}),.exe,) \
		${MODULE}.1
	@rm \
		${DIST_DIR}/${MODULE}$(if $(findstring windows,${GOOS}),.exe,) \
		${DIST_DIR}/${MODULE}.1


.PHONY: build
build: #@HELP Builds binary for the current architecture in ./bin
	$(info ${MESSAGE_PREFIX} Building binary as bin/${MODULE})
	@GOFLAGS="${GOFLAGS}" go build -ldflags="${LDFLAGS}" -o bin/${MODULE} ${PKGS}
	@chmod 755 bin/${MODULE}


.PHONY: clean
clean: # @HELP Remove generated files and caches
	$(info ${MESSAGE_PREFIX} Removing generated files and caches)
	@echo "${MESSAGE_PREFIX} Removing test cache for ${MODULE}:"
	@go clean -testcache
	@echo "${MESSAGE_PREFIX} Removing ./bin and ${DIST_DIR}/ directories"
	@rm -Rf \
		./bin/ \
		${DIST_DIR}/
	@echo "${MESSAGE_PREFIX} Recreating ./bin and ${DIST_DIR}/ directories"
	@mkdir -p ./bin ${DIST_DIR}/
	@echo "${MESSAGE_PREFIX} Creating .gitkeep in ./bin and ${DIST_DIR}/"
	@touch \
		./bin/.gitkeep \
		${DIST_DIR}/.gitkeep
	@echo "${MESSAGE_PREFIX} Removing generated manpages:"
	@rm -f \
		${DOC_DIR}/${MODULE}.1          \
		${DOC_DIR}/${MODULE}.1-build.md \
		${DOC_DIR}/${MODULE}.1.pdf
	@echo "${MESSAGE_PREFIX} Removing coverage.out:"
	@rm -f coverage.out
	@echo "${MESSAGE_PREFIX} Running go mod tidy:"
	@go mod tidy


.PHONY: coverage
coverage: # @HELP Produces HTML unit test coverage report in ${DOC_DIR}
	$(info ${MESSAGE_PREFIX} Producing HTML unit test coverage report)
	@echo "${MESSAGE_PREFIX} Removing test cache for ${MODULE}:"
	go clean -testcache
	@echo "${MESSAGE_PREFIX} Running go test:"
	go test ./... -race -cover -covermode=atomic -coverprofile=coverage.out
	@echo "${MESSAGE_PREFIX} Generating coverage report in ${DOC_DIR}:"
	go tool cover -html=coverage.out -o ${DOC_DIR}/coverage.html


.PHONY: help
help: # @HELP Prints this message
	$(info ${MESSAGE_PREFIX} HELP:)
	@echo "VARIABLES:"
	@echo "  BUILD_DATE = ${BUILD_DATE}"
	@echo "  GOARCH     = ${GOARCH}"
	@echo "  GOFLAGS    = ${GOFLAGS}"
	@echo "  GOOS       = ${GOOS}"
	@echo "  LDFLAGS    = ${LDFLAGS}"
	@echo "  MODULE     = '${MODULE}'"
	@echo "  PKGS       = ${PKGS}"
	@echo "  VERSION    = ${VERSION}"
	@echo
	@echo "BUILD PLATFORMS:"
	@echo "  ${PLATFORMS}"
	@echo
	@echo "Targets for make:"
	@grep -E '^.*:.*# *@HELP' ${MAKEFILE_LIST} |   \
		awk 'BEGIN {FS = ":.*# *@HELP"};           \
			 { gsub(/:.*/, "", $$1);               \
			 printf "  %-12s %s\n", $$1, $$2 } ' | \
		sort


.PHONY: install
install: build man # @HELP Install binary for the current architecture
	$(info ${MESSAGE_PREFIX} Installing to ${INSTALL_PATH}/${MODULE})
	@install -m 755 bin/${MODULE} ${INSTALL_PATH}/${MODULE}
	$(info ${MESSAGE_PREFIX} Ensuring existence of ${MAN_PATH})
	@mkdir -p "${MAN_PATH}"
	$(info ${MESSAGE_PREFIX} Installing manpage to ${MAN_PATH})
	@install -m 644 ${DOC_DIR}/${MODULE}.1 ${MAN_PATH}/${MODULE}.1

.PHONY: lint
lint: # @HELP Formats and lints the Go source code
	$(info ${MESSAGE_PREFIX} Running gocheckfile.sh on all files)
	@find . -type f -iname "*.go" -exec gocheckfile.sh {} \;


.PHONY: man
man: # @HELP Uses Pandoc to build a manpage from Markdown
	$(info ${MESSAGE_PREFIX} Creating manpage from Markdown file)
	@echo "${MESSAGE_PREFIX} Substituting variables in manpage"
	@sed \
		-e 's/{{date}}/${BUILD_DATE}/g'               \
		-e 's/{{version}}/${VERSION}/g'               \
		${DOC_DIR}/${MODULE}.1.md                     \
			> ${DOC_DIR}/${MODULE}.1-build.md
	@echo "${MESSAGE_PREFIX} Generating manpage from template"
	@pandoc \
		--standalone --to man           \
		--output ${DOC_DIR}/${MODULE}.1 \
		${DOC_DIR}/${MODULE}.1-build.md
	@echo "${MESSAGE_PREFIX} Generating pdf from manpage"
	@pandoc \
		--standalone -V geometry:margin=1in          \
		--to pdf --output ${DOC_DIR}/${MODULE}.1.pdf \
		${DOC_DIR}/${MODULE}.1
	@echo "${MESSAGE_PREFIX} Removing temp file"
	@rm -f ${DOC_DIR}/${MODULE}.1-build.md


.PHONY: run
run: # @HELP Runs the application
	$(info ${MESSAGE_PREFIX} Running ${PKGS})
	@go run ${PKGS}


.PHONY: show-man
show-man: # @HELP Shows the manpage (as built) using nroff
	$(info ${MESSAGE_PREFIX} Displaying ${DOC_DIR}/${MODULE}.1 manpage)
	@nroff -man ${DOC_DIR}/${MODULE}.1


.PHONY: show-version
show-version: # @HELP Shows the application version
	$(info ${MESSAGE_PREFIX} ${MODULE} Version)
	@echo "${VERSION}"


.PHONY: test
test: # @HELP Runs unit tests
	$(info ${MESSAGE_PREFIX} Removing test cache for ${MODULE})
	@go clean -testcache
	$(info ${MESSAGE_PREFIX} Running unit tests for ${MODULE})
	@go test -race ./...


.PHONY: tree
tree: # @HELP Produces tree output for the project directory
	$(info ${MESSAGE_PREFIX} ${MODULE} project tree)
	@tree -I _WIP -I '*.sublime-project' -I '*.sublime-workspace' .
