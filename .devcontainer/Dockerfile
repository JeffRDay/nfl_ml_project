# [Choice] Python version (use -bullseye variants on local arm64/Apple Silicon): 3, 3.10, 3.9, 3.8, 3.7, 3.6, 3-bullseye, 3.10-bullseye, 3.9-bullseye, 3.8-bullseye, 3.7-bullseye, 3.6-bullseye, 3-buster, 3.10-buster, 3.9-buster, 3.8-buster, 3.7-buster, 3.6-buster
ARG VARIANT=3-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/python:0-${VARIANT}

ENV PYTHONUNBUFFERED 1

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

ARG MONGO_TOOLS_VERSION=5.0
RUN . /etc/os-release \
    && if [ "${VERSION_CODENAME}" = "buster" ] && [ "$(dpkg --print-architecture)" = "amd64" ]; then \
        curl -sSL "https://www.mongodb.org/static/pgp/server-${MONGO_TOOLS_VERSION}.asc" | gpg --dearmor > /usr/share/keyrings/mongodb-archive-keyring.gpg \
        && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/mongodb-archive-keyring.gpg] http://repo.mongodb.org/apt/debian $(lsb_release -cs)/mongodb-org/${MONGO_TOOLS_VERSION} main" | tee /etc/apt/sources.list.d/mongodb-org-${MONGO_TOOLS_VERSION}.list \
        && apt-get update && export DEBIAN_FRONTEND=noninteractive \
        && apt-get install -y mongodb-database-tools mongodb-mongosh \
        && apt-get clean -y && rm -rf /var/lib/apt/lists/*; \
    fi

# [Optional] If your requirements rarely change, uncomment this section to add them to the image.
# COPY requirements.txt /tmp/pip-tmp/
# RUN pip3 --disable-pip-version-check --no-cache-dir install -r /tmp/pip-tmp/requirements.txt \
#    && rm -rf /tmp/pip-tmp

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>



