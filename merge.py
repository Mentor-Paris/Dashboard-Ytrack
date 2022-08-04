import pandas as pd
import os

cwd = os.getcwd()

usernational = pd.read_json(cwd+"/assets/json/usersnational.json")
usercampus = pd.read_json(cwd+"/assets/json/usersnationalxp.json")

usercampus = usercampus.rename(columns={'firstName': 'login'})

merged = usernational.merge(usercampus, on='login', how="outer")
merged.to_csv(cwd+"/assets/files/merge_data_user2.csv", index=False)