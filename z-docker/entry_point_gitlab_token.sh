#!/usr/bin/env bash

# most change
GIT_LAB_COMMIT_CODE=${WEB_GIT_LAB_COMMIT_CODE}
GIT_LAB_BRANCH=${WEB_GIT_LAB_BRANCH}
GIT_LAB_TOKEN=${WEB_GIT_LAB_TOKEN}

GIT_LAB_REPO=${WEB_GIT_LAB_REPO}
GIT_LAB_GROUP=${WEB_GIT_LAB_GROUP}
GIT_LAB_HOST=${WEB_GIT_LAB_HOST}
GIT_LAB_API=http://gitlab-ci-token
INFO_ENV_PREFIX="WEB_"

run_path=$(pwd)
shell_run_name=$(basename $0)
shell_run_path=$(cd `dirname $0`; pwd)

pV(){
    echo -e "\033[;36m$1\033[0m"
}
pI(){
    echo -e "\033[;32m$1\033[0m"
}
pD(){
    echo -e "\033[;34m$1\033[0m"
}
pW(){
    echo -e "\033[;33m$1\033[0m"
}
pE(){
    echo -e "\033[;31m$1\033[0m"
}

checkFuncBack(){
  if [[ $? -ne 0 ]]; then
    echo -e "\033[;31mRun [ $1 ] error exit code 1\033[0m"
    exit 1
  fi
}

checkBinary(){
  binary_checker=`which $1`
  checkFuncBack "which $1"
  if [[ ! -n "${binary_checker}" ]]; then
    echo -e "\033[;31mCheck binary [ $1 ] error exit\033[0m"
    exit 1
#  else
#    echo -e "\033[;32mCli [ $1 ] event check success\033[0m\n-> \033[;34m$1 at Path: ${evn_checker}\033[0m"
  fi
}

checkEnv(){
	if [[ $(env | grep "$1" | wc -l) -ne 1 ]]; then
		echo -e "\033[;31mCheck environment [ $1 ] error exit\033[0m"
		exit 1
	fi
}

check_root(){
  if [[ ${EUID} != 0 ]]; then
    echo "no not root user"
  fi
}

checkBinary git
#checkBinary golang

if test ${GIT_LAB_TOKEN}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_TOKEN not set!"
    exit 1
fi

if test ${GIT_LAB_REPO}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_REPO not set!"
    exit 1
fi

if test ${GIT_LAB_GROUP}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_GROUP not set!"
    exit 1
fi

if test ${GIT_LAB_HOST}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_HOST not set!"
    exit 1
fi

if test ${GIT_LAB_BRANCH}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_BRANCH not set!"
    exit 1
fi

if test ${GIT_LAB_COMMIT_CODE}; then
    pW "env ${INFO_ENV_PREFIX}GIT_LAB_COMMIT_CODE not set!"
    exit 1
fi

echo -e "Now config of environment
ENV_PREFIX      -> ${INFO_ENV_PREFIX}
GIT_LAB_API     -> ${GIT_LAB_API}
GIT_LAB_HOST    -> ${GIT_LAB_HOST}
GIT_LAB_GROUP   -> ${GIT_LAB_GROUP}
GIT_LAB_REPO    -> ${GIT_LAB_REPO}
GIT_LAB_BRANCH  -> ${GIT_LAB_BRANCH}
GIT_LAB_COMMIT_CODE -> ${GIT_LAB_COMMIT_CODE}
"

if [[ ! -d ${GIT_LAB_REPO} ]];then
	git clone ${GIT_LAB_API}:${GIT_LAB_TOKEN}@${GIT_LAB_HOST}/${GIT_LAB_GROUP}/${GIT_LAB_REPO}.git
	checkFuncBack "git clone ${GIT_LAB_API}:${GIT_LAB_TOKEN}@${GIT_LAB_HOST}/${GIT_LAB_GROUP}/${GIT_LAB_REPO}.git"
fi

# to repo
cd ${GIT_LAB_REPO}
echo "run at $PWD"

now_path_commit_code=$(git reset --head HEAD)
if [[ "${GIT_LAB_COMMIT_CODE}" -ne "${now_path_commit_code}" ]];then
	git reset --head HEAD
	#git rev-parse refs/remotes/origin/master^{commit}
	#git rev-parse refs/remotes/origin/origin/master^{commit}
	git checkout origin/${GIT_LAB_BRANCH}
	git checkout -f ${GIT_LAB_COMMIT_CODE}
#git rev-list ${GIT_LAB_COMMIT_CODE}
fi

pW "must change run code here"
