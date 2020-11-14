! /bin/sh

ps -ef | grep questionnaire-api| grep -v grep | awk '{print $2}' | xargs kill -9
cd ~/go/src/questionnaire/
git pull https://github.com/KyrieIrving2/questionnire-api.git
./questionnaire-api &