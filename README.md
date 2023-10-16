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
- Run `./houseanalyzerfaster.exe`
- Run `./houseanalyzerfasestt.exe`

### Running benchmark
- run `chmod +x benchmark.sh`
- run `./benchmark.sh {RUNS}` where {RUNS} is the amount of runs you want.

### Results

I ended up not getting the results I thought I would so I implmented a few changes.

Fast - Tries to minimize the data transformations and reads each csv row only once
Faster - Additionally doesn't use the tab writer
Fastest - Same as faster but reads all the rows at once then iterates over them

Note: These version of the the app do not have tests.

Then I removed all print statements in main for the last runs

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

I implemented a version with less transformations to the data and was quite disapointed when only a little bit of time was saved.

</details>

<details>
<summary>Results from fast implementation 2</summary>

| Language               | Test Run 1 (nanoseconds) | Test Run 2 (nanoseconds) | Test Run 3 (nanoseconds) | Test Run 4 (nanoseconds) |
|----------------------|---------------------------|---------------------------|---------------------------|---------------------------|
| Go                   |       3,728,930,100       |       3,736,011,000       |       3,730,175,700       |       3,742,231,500       |
| Go (fast)            |       3,424,380,500       |       3,420,223,500       |       3,419,356,100       |       3,437,085,350       |
| Go (faster)          |       3,408,618,800       |       3,401,338,700       |       3,405,574,600       |       3,412,590,000       |
| Go (fastest)         |       3,505,374,200       |       3,517,643,700       |       3,499,446,050       |       3,516,287,150       |
| Python               |       1,755,800,600       |       1,710,396,000       |       1,738,588,250       |       1,756,981,200       |
| R                    |       3,080,280,200       |       3,086,912,400       |       3,069,791,350       |       3,095,916,600       |

I implemented two more version removing the tab writer and was still disapointed when the times changed even less.


</details>

<details>
<summary>Results from fast implementation without printing</summary>

| Language               | Test Run 1 (nanoseconds) | Test Run 2 (nanoseconds) | Test Run 3 (nanoseconds) | Test Run 4 (nanoseconds) |
|----------------------|---------------------------|---------------------------|---------------------------|---------------------------|
| Go                   |       3,793,349,200       |       3,734,792,700       |       3,739,142,000       |       3,746,888,200       |
| Go (fast)            |       3,497,688,700       |       3,449,193,200       |       3,450,075,600       |       3,456,628,600       |
| Go (faster)          |       3,470,810,000       |       3,419,502,900       |       3,434,581,700       |       3,437,296,200       |
| Go (fastest)         |       3,554,384,100       |       3,499,173,900       |       3,518,411,200       |       3,512,986,700       |
| Python               |       1,726,555,000       |       1,719,113,300       |       1,717,965,200       |       1,705,428,000       |
| R                    |       3,087,121,700       |       3,074,617,600       |       3,096,046,300       |       3,110,616,800       |

Since I was using bash to time all of these I thought that removing the print statements would decrease times but it really did nothing.

</details>


Overall, I'm disapointed with the results since I thought Go would run faster. After running my implmentations vs python and R I would say that it would be worth it to use python as a reader/writer for summary data. I think I may be using the stats package sub-optimally so further tests should be done in order to determine which language we should use for reading and writing summaries for data.