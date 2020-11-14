#! /bin/sh

kill -9 $(pgrep questionnaire-api)
cd ~/go/src/questionnaire/
git pull https://github.com/KyrieIrving2/questionnire-api.git
./questionnaire-api &