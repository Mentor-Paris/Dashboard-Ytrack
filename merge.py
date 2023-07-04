import pandas as pd
import os

cwd = os.getcwd()

usernational = pd.read_json(cwd+"/assets/json/usersnational.json")
usercampus = pd.read_json(cwd+"/assets/json/usersnationalxp.json")

usercampus = usercampus.rename(columns={'login': 'login'})

merged = usernational.merge(usercampus, on='login', how="outer")

merged["xps_aggregate"] = merged["xps_aggregate"].fillna(0)

for i in range (0, len(merged)):
    if merged.xps_aggregate[i] :
        merged.xps_aggregate[i] = merged.xps_aggregate[i]['aggregate']['sum']['amount']
    else :
        merged.xps_aggregate[i] = 0

merged.to_csv(cwd+"/assets/files/merge_data_user.csv", index=False)