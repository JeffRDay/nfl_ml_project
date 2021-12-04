import pandas as pd

df1 = pd.read_csv("../prototype-data/Game_Logs_Runningback.csv")
df2 = pd.read_csv("../prototype-data/Game_Logs_Wide_Receiver_and_Tight_End.csv")

print(df1.describe)
print(df2.describe)

merged = pd.merge(df1, df2)
print(merged.head)
print(merged.describe)
merged.to_csv("working.csv")