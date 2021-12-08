#!/usr/bin/env python
# coding: utf-8
​
# In[1]:
​
​
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import sklearn
from sklearn import neighbors, linear_model, datasets
from sklearn.model_selection import train_test_split
​
​
# In[2]:
​
​
dt = pd.DataFrame
df = pd.read_csv('test.csv')
​
​
# In[3]:
​
​
df
​
​
# In[4]:
​
​
target = df[("fs_total")]
​
​
# In[5]:
​
​
#casual_X_train, casual_X_test, casual_y_train, casual_y_test 
#= train_test_split(df_features, casual_target, test_size = .33, random_state=42)
​
​
# In[6]:
​
​
casual_X_train, casual_X_test, casual_y_train, casual_y_test = train_test_split(df, target, test_size = .33, 
                                                                                random_state=42)
​
​
# In[7]:
​
​
len(df)
​
​
# In[8]:
​
​
len(casual_X_train)
​
​
# In[9]:
​
​
len(casual_X_test)
​
​
# In[10]:
​
​
len(casual_y_test)
​
​
# In[11]:
​
​
len(casual_y_train)
​
​
# In[12]:
​
​
#KNN
knn_regression = neighbors.KNeighborsRegressor(n_neighbors=4, weights='uniform')
knn_regression.fit(casual_X_train, casual_y_train)
​
print("k-nearest neighbor")
print("Mean Squared Error: %.2f"
      % np.mean((knn_regression.predict(casual_X_test) - casual_y_test) ** 2))
print('R^2: %.2f' % knn_regression.score(casual_X_test, casual_y_test))
​
​
# In[ ]:
​