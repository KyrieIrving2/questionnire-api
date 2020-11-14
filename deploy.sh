! /bin/sh

ps -ef | grep questionnaire-api| grep -v grep | awk '{print $2}' | xargs kill -9
cd ~/go/src/questionnaire/qustionnire-api/
git pull
./questionnaire-api &
