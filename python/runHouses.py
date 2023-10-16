import pandas as pd
N = 100 
with open('./data/benchmark/housesOutputPy.txt', 'wt') as outfile:
    for i in range(N):
        houses = pd.read_csv("./data/housesInput.csv")
        outfile.write(houses.describe().to_string(header=True, index=True))
        outfile.write("\n")




