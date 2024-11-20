#!/bin/bash

##########################################
# VSCode Extensions Installer ############
##########################################

# Text colors ############################
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' #### no color

echo -e "${BLUE}Installing VSCode Extensions...${NC}\n"

# Read extensions #########################
# from .vscode/extensions.json ############
EXTENSIONS=$(jq -r '.recommendations[]' .vscode/extensions.json)

# Install each extension #################
for ext in $EXTENSIONS
do
    if [ ! -z "$ext" ]; then
        echo -e "Installing ${GREEN}$ext${NC}"
        code --install-extension "$ext" --force
    fi
done

echo -e "\n${BLUE}Installation Complete!${NC}"






##########################################
# 💞 @octokas ############################
##########################################
