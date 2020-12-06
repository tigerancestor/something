echo off
REM 声明采用UTF-8编码
chcp 65001
:: chcp 936
cd ..\..\..\
set PROJECTBASE=%cd%
echo PROJECTBASE: %PROJECTBASE%
set CONFIGDATABASE=%PROJECTBASE%\scheme\数值配置
echo CONFIGDATABASE: %CONFIGDATABASE%
set GOOUT=%PROJECTBASE%\programe\server\src\game-stage\data\struct
echo GOOUT: %GOOUT%
set JSONOUT=%PROJECTBASE%\programe\public\data
echo JSONOUT: %JSONOUT%
set TSOUT=%PROJECTBASE%\programe\public\ts
echo TSOUT: %TSOUT%
rd /S /Q %GOOUT%
rd /S /Q %JSONOUT%
rd /S /Q %TSOUT%
mkdir %GOOUT%
mkdir %JSONOUT%
mkdir %TSOUT%
cd /d %PROJECTBASE%\programe\public\tools
echo %cd%
for /R %CONFIGDATABASE% %%f in (*.xlsx,*.xlsm) do (
gen-xlsx-data.exe --go_out=%GOOUT% --json_out=%JSONOUT% --ts_out=%TSOUT% %%f
echo %%f generate data success!
)
::mkdir %JSONOUT%\script
::copy /y %PROJECTBASE%\programe\public\battle\BattleLaya\release\layaweb\web\code.js  %JSONOUT%\script\code.js
::rd %SERVERJSON%
::mkdir %SERVERJSON%
::robocopy %JSONOUT% %SERVERJSON% /e
::del %OUTGOBASE%\msg2id.json
::copy /y %MESSAGEBASE%\msg2id.json %OUTGOBASE%\msg2id.json
pause
%GOPLUGIN%
exit