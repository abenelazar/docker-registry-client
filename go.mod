module github.com/abenelazar/docker-registry-client

go 1.12

require (
	github.com/docker/distribution v0.0.0-20171011171712-7484e51bf6af
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/go-logr/logr v1.0.0-rc1
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.4.1 // indirect
	github.com/golangci/golangci-lint v1.17.2-0.20190909185456-6163a8a79084
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/tools v0.1.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// From/For golangci-lint. Can be removed once v1.17.2 (or newer) is released
replace (
	// https://github.com/ultraware/funlen/pull/1
	github.com/ultraware/funlen => github.com/golangci/funlen v0.0.0-20190909161642-5e59b9546114
	// https://github.com/golang/tools/pull/139
	golang.org/x/tools => github.com/golangci/tools v0.0.0-20190909104219-979bdb7f8cc8
)
