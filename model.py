
import numpy as np
from io import StringIO
from sklearn.neural_network import MLPClassifier
import pandas as pd
import sklearn
import sklearn.tree
from sklearn import datasets
from sklearn import preprocessing
from sklearn import model_selection
from sklearn import metrics
from matplotlib import pyplot as plt
import numpy as np
import logging
import time

def main():
    logging.basicConfig(encoding='utf-8', level=logging.INFO)
    logging.info("starting model.py")
    start = time.time()
    
    df = readData("test.csv")
    logging.info("df length - {}".format(len(df)))

    size = len(df) // 50
    logging.info("using size - {}".format(size))

    tenthDf = df.iloc[:size,]
    wdf = cleanCopy(tenthDf)
    X, y = splitData(wdf)

    logging.info("converting X and y to numpy arrays")
    X = X.to_numpy()
    y = y.to_numpy()

    logging.warning("starting MLP Classifier - good luck.")
    oneA_relu = model("Relu: 2 hidden layers with 20 nodes per layer", (20,60), "relu", X, y)

    logging.info("MLP Classifier completed, results:")
    logging.info(oneA_relu)

    end = time.time()
    print("executed model.py in :", end-start)

def readData(path):
    logging.info("reading data from {}".format(path))
    return pd.read_csv(path)

def cleanCopy(df):
    wdf = df.copy()
    wdf = wdf.drop('player_id', axis=1)
    return wdf

def splitData(df):
    X = df.drop('fs_total', axis=1)
    y = df['fs_total']
    return X, y

def model(description, arch, activation, X, y):
    kfold = model_selection.KFold(5, shuffle=True, random_state=2)

    prec, rec, f1 = [], [], []

    for train_idx, test_idx in kfold.split(X):
        X_train, X_test = X[train_idx], X[test_idx]
        y_train, y_test = y[train_idx], y[test_idx]
        
        mlp_clf = MLPClassifier(hidden_layer_sizes=arch, max_iter=2000, activation=activation, random_state=2, solver='lbfgs') # number of hidden layers is passed to the classifier
        X_scaler = preprocessing.MinMaxScaler()
        
        X_train = X_scaler.fit_transform(X_train)    
        
        mlp_clf.fit(X_train, y_train)
        
        X_test = X_scaler.transform(X_test)
        y_pred = mlp_clf.predict(X_test)
        
        # print(y_pred)
        # print(y_test)
        
        rec += [metrics.recall_score(y_pred, y_test, average="weighted")]
        prec += [metrics.precision_score(y_pred, y_test, average="weighted")]
        f1 += [metrics.f1_score(y_pred, y_test, average="weighted")]
        
        # print()

    results = [
        description,
        "{:.4f} ±{:.4f}".format(np.mean(rec), np.std(rec)), # recall
        "{:.4f} ±{:.4f}".format(np.mean(prec), np.std(prec)), # precision
        "{:.4f} ±{:.4f}".format(np.mean(f1), np.std(f1)) #f1
    ]
    return results

main()
