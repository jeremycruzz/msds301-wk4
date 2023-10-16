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
- Run `go build -o houseanalyzerfast.exe ./cmd/houseanalyzerfast`
- Run `go build -o houseanalyzerfaster.exe ./cmd/houseanalyzerfaster`
- Run `go build -o houseanalyzerfastest.exe ./cmd/houseanalyzerfastest`

### Running Go executable
- Run `./houseanalyzer.exe {READFILE} {WRITEFILE}`
- Example `./houseanalyzer.exe ./data/housesInput.csv ./data/housesOutGo.txt`

- Run `./houseanalyzerfast.exe`

### Running benchmark
- run `chmod +x benchmark.sh`
- run `./benchmark.sh {RUNS}` where {RUNS} is the amount of runs you want.

### Results

<details>
<summary> Results from initial implementation commit hash:36eec8e2c3210a318916e5c79728b2028d1e27e2</summary>

-------------------------------------------------------------------
| Language     | Test Run 1 (nanoseconds) | Test Run 2 (nanoseconds) | Test Run 3 (nanoseconds) |
|--------------|---------------------------|---------------------------|---------------------------|
| Go           |        3,730,301,800       |        3,748,057,300       |        3,717,731,200       |
| Python       |        1,755,509,700       |        1,729,278,400       |        1,735,856,500       |
| R            |        3,106,460,900       |        3,110,648,100       |        3,132,068,400       |
-------------------------------------------------------------------

These results were suprising at first but I realize that go performed the worst because of my implementation of the app. Both the R code and Python accomplished what my Go program accomplished in 7 lines of code. The main issue with my implementation was the way I designed the Block struct. When reading the data we go over `n` rows. With analyze we then go over these rows of data an additional time when creating the pivot table. We then analyze the metrics and convert the data there back to the block struct before writing it.
</details>

<details>
<summary>Results from fast implementation 1</summary>

| Language     | Test Run 1 (nanoseconds) | Test Run 2 (nanoseconds) | Test Run 3 (nanoseconds) | Test Run 4 (nanoseconds) |
|--------------|---------------------------|---------------------------|---------------------------|---------------------------|
| Go           |        3,735,719,000      |        3,738,952,500      |        3,754,574,850      |        3,743,969,050      |
| Go (fast)    |        3,449,520,300      |        3,429,550,100      |        3,435,929,850      |        3,430,158,550      |
| Python       |        1,751,711,700      |        1,749,594,900      |        1,729,447,600      |        1,730,606,250      |
| R            |        3,115,111,200      |        3,100,016,100      |        3,118,881,450      |        3,084,558,750      |


</details>

<summary>Results from fast implementation 2</summary>

| Language     | Test Run 1 (nanoseconds) | Test Run 2 (nanoseconds) | Test Run 3 (nanoseconds) | Test Run 4 (nanoseconds) |
|--------------|---------------------------|---------------------------|---------------------------|---------------------------|
| Go           |        3,735,719,000      |        3,738,952,500      |        3,754,574,850      |        3,743,969,050      |
| Go (fast)    |        3,449,520,300      |        3,429,550,100      |        3,435,929,850      |        3,430,158,550      |
| Python       |        1,751,711,700      |        1,749,594,900      |        1,729,447,600      |        1,730,606,250      |
| R            |        3,115,111,200      |        3,100,016,100      |        3,118,881,450      |        3,084,558,750      |


</details>