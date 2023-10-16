# msds301-wk4

### Setup (Windows 11 + Git Bash)
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk4.git`
- Install go dependencies `go mod tidy`
- Install py dependencies `py -m pip install pandas`
- Install ginkgo `go install github.com/onsi/ginkgo/v2/ginkgo`
- Install mockgen `go install github.com/golang/mock/mockgen@v1.6.0`
- Create mocks `go generate ./...`

### Running Tests
- Run tests with `ginkgo -r`

### Building executable
- Run `go build -o houseanalyzer.exe ./cmd/houseanalyzer`

### Running Go executable
- Run `./houseanalyzer {READFILE} {WRITEFILE}`
- Example `./houseanalyzer ./data/housesInput.csv ./data/housesOutGo.txt`

### Running benchmark
- run `chmod +x benchmark.sh`
- run `./benchmark.sh {RUNS}` where {RUNS} is the amount of runs you want.