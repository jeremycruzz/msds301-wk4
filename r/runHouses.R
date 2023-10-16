N = 100 
sink("./data/benchmark/housesOutputR.txt")
for (i in 1:N) {
    houses = read.csv(file = "./data/housesInput.csv", header = TRUE)
    print(summary(houses)) 
}
sink()
