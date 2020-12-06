echo off
cd ..\..\
set REDIS="base.it-crazy.net:6379"
set REDISPWD="123456"
set MASTER=""
set PROGRAMEBASE=%cd%
cd /d %PROGRAMEBASE%\public\tools
set DATAPATH=%PROGRAMEBASE%\public\data
set BRANCH="gw"
set VERSION="0.30"
set PREFIX="hssg/data/%BRANCH%/%VERSION%"

data-to-redis.exe --redis_addr=%REDIS% --redis_pwd=%REDISPWD% --data_path=%DATAPATH% --prefix=%PREFIX% --master=%MASTER%
pause
exit