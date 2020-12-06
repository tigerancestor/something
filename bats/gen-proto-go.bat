echo off
cd ..\..\
set PROGRAMEBASE=%cd%
::echo PROGRAMEBASE: %PROGRAMEBASE%
set MESSAGEBASE=%PROGRAMEBASE%\public\message
svn update %MESSAGEBASE%
::echo MESSAGEBASE:%MESSAGEBASE%
set OUTGOBASE=%PROGRAMEBASE%\server\src\game-stage\msg
::echo OUTGOBASE: %OUTGOBASE%
cd /d %PROGRAMEBASE%\public\tools
echo %cd%
for /R %MESSAGEBASE% %%f in (*.proto) do (
protoc.exe  --plugin=protoc-gen-gogofast.exe --gogofast_out=%OUTGOBASE% --proto_path=%MESSAGEBASE% %%f
echo %%f generate go success!
)
gen-msg-id.exe --msg_path=%MESSAGEBASE% --msg2id_file=%MESSAGEBASE%\msg2id.json --go_reg_file=%OUTGOBASE%\regist.go
echo generate msg-id success!
::del %OUTGOBASE%\msg2id.json
::copy /y %MESSAGEBASE%\msg2id.json %OUTGOBASE%\msg2id.json
pause
%GOPLUGIN%
exit                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         