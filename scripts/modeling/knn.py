'''
Creates a Multilayer Perceptron Classifier (MLPC) model. The
model is trained using the csv data provided. Evaulation is
conducted using 5-fold cross validation. The model is serialized
and saved at the end of execution regardless of performance.
'''
import logging
import time

import pandas as pd
import numpy as np
from sklearn import neighbors, model_selection, preprocessing, metrics

PATH = "../../dist/20210712_214459_formatted-data.csv"
LIMITER = 1000
K_VALUE = 5

def main():
    ''' main execution function '''
    logging.basicConfig(
        encoding='utf-8',
        level=logging.INFO,
        format="%(levelname)s ::: %(message)s")

    logging.info("starting model.py")
    start = time.time()

    data_frame = read_csv(PATH)
    logging.info(
        "read in %s records to data frame with %s features",
        len(data_frame),
        len(data_frame.columns))

    size = len(data_frame) // LIMITER
    logging.info("will use %s records for MLP Classifier", size)

    subset = data_frame.iloc[:size,]
    working_subset = clean_copy(subset)
    predictors, target = split_data(working_subset)

    logging.info("converting predictors and target to numpy arrays")
    predictors = predictors.to_numpy()
    target = target.to_numpy()

    logging.warning("starting KNN Regression - good luck.")
    results = model(
        "KNN, where K = " +str(K_VALUE),
        (K_VALUE),predictors, target)

    logging.info("KNN Regression completed, results:")
    logging.info(results)

    end = time.time()
    print("executed model.py in :", end-start)

def read_csv(path):
    ''' read_csv reads csv from provided path and return dataframe '''
    return pd.read_csv(path)

def clean_copy(data_frame):
    ''' clean_copy copies the given data frame and drops player_id '''
    wdf = data_frame.copy()
    wdf = wdf.drop('player_id', axis=1)
    return wdf

def split_data(data_frame):
    ''' splits dataframe into predictor and target dataframes '''
    predictors = data_frame.drop('fs_total', axis=1)
    target = data_frame['fs_total']
    return predictors, target

def model(description, K_VALUE, predictors, target):
    ''' uses 5-fold cross validation to create and save MLPC model '''
    kfold = model_selection.KFold(5, shuffle=True, random_state=2)

    prec, rec, f_score = [], [], []

    for train_idx, test_idx in kfold.split(predictors):
        predictors_train, predictors_test = predictors[train_idx], predictors[test_idx]
        target_train, target_test = target[train_idx], target[test_idx]

        knn_regression = neighbors.KNeighborsRegressor(n_neighbors=K_VALUE, weights='uniform')

        # with ignore_warnings(category=UndefinedMetricWarning):

        knn_regression.fit(predictors_train, target_train)

        rec += [metrics.recall_score(knn_regression.predict(predictors_test), target_test, average="weighted")]
        prec += [metrics.precision_score(knn_regression.predict(predictors_test), target_test, average="weighted")]
        f_score += [metrics.f1_score(knn_regression.predict(predictors_test), target_test, average="weighted")]

    results = [
        description,
        "{:.4f} ±{:.4f}".format(np.mean(rec), np.std(rec)), # recall
        "{:.4f} ±{:.4f}".format(np.mean(prec), np.std(prec)), # precision
        "{:.4f} ±{:.4f}".format(np.mean(f_score), np.std(f_score)) #f_score (f1)
    ]
    return results

main()
