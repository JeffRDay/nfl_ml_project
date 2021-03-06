# nfl_ml_project

NOTE: this is an academic project and cannot accept contributions until 12/14/2021

### Abstract

Fantasy football is a multi-billion-dollar industry where individuals from across the world join fantasy football leagues to draft National Football League (NFL) football players onto their own fantasy team and compete against other individual’s fantasy teams. The individuals, also referred to as “owners” or “team owners”, assess NFL player and team statistics to project which NFL players on their active roster and bench will accumulate the most points in the fantasy football scoring system. The objective of this project is to explore machine learning approaches to predicting the fantasy scores achieved by NFL players on a per game basis. The project utilized several different supervised learning techniques to identify the optimal solution and provide analysis for use in future efforts.

## Table of Contents

1. Directory Structure
2. Machine Learning Pipeline
3. Development Environment Setup
4. Executing Scripts/Code

## Directory Structure

```bash
├───.devcontainer # dev environment
├───code # all code developed by the project team
│   ├───data_understanding # random scripts to understand contents of raw data
│   ├───modeling # jupyter NB and requirements.txt for running models
│   ├───post_processing # random scripts to review post modeling results
│   └───pre_processing # golang script to take raw data and produce formatted csv
├───data # all data associated with the project
│   ├───clean_formatted # pre-processed CSV used for modeling
│   ├───raw # the raw csv data requiring pre-processing before modeling
│   └───results # all performance metrics and testing results
└───imgs # images associated with the project
```

## Machine Learning Pipeline

![ML Pipeline](imgs/CS5644_ML_Pipeline-Detailed.drawio.png)

## Development Environment Setup

1. Install VS Code ([link](https://code.visualstudio.com/Download))
2. Follow Instructions for Remote Containers ([link](https://code.visualstudio.com/docs/remote/containers))
3. Clone this repository ([clone link](https://github.com/JeffRDay/nfl_ml_project.git))
4. Open `nfl_ml_project` directory in VS Code
5. If not prompted, press F1 (if windows) or otherwise open command palette.
6. Type `remote-containers`
7. Select `Remote-Containers: Open Folder in Container`
8. Click `open`
9. VS Code should launch the development environment
10. Contact the project team if any issues persist

## Executing scripts

This project uses Go to pre-process the raw data located in `/data/raw`. To execute the go script:

### OPTION 1: Download the latest release

1. Download latest release for your system architecture from [releases](https://github.com/JeffRDay/nfl_ml_project/releases)

NOTE: Due to resource limitations, we are only able to test on 64-bit linux and windows operating systems.

2. copy the executable into `/code/pre_processing`. This **MUST** be done. The script is dependent on the file structure of this repo at this time.

3. execute the binary. results will be output to `/data/clean_formatted` directory.

if linux:
```bash
./<binary name>
```

if windows: double click on the applicable `*.exe`.

### OPTION 2: Use the executables in the source code project

1. cd to `/code/pre_processing` and execute the executable for your system architecture.

results will be output to `/data/clean_formatted` directory.

### OPTION 3: Follow the development environment setup instructions

After completing the development environment setup instructions, cd into `/code/pre-processing`

```bash
cd code/pre-processing
```

Copy/paste the command below into your terminal of choice to execute the go script.
```bash
go run ./...
```

*note - do NOT use `go run go-main.go`. This is a multi-file script requiring the execution of the command above rather than running a specific file.*

results will be output to `/data/clean_formatted` directory.