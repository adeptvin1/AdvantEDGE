#!/usr/bin/env bash

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
ln -s $NVM_DIR ~/nvm
if [ -d $NVM_DIR ]
then
echo "Sucess! NVM dir in $NVM_DIR"
else
echo "Error! Selected path is not a correct directory"; exit 2
fi