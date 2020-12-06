echo off
REM 声明采用UTF-8编码
chcp 65001
cd ..\
svn commit -m"提交协议和配置文件" .
pause
exit