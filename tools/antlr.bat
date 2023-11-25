@ECHO OFF
SET cur_dir=%~dp0
SET CLASSPATH=.;%cur_dir%\antlr-4.13.1-complete.jar;%CLASSPATH%
SET TEST_CURRENT_DIR=%CLASSPATH:.;=%
if "%TEST_CURRENT_DIR%" == "%CLASSPATH%" ( SET CLASSPATH=.;%CLASSPATH% )
java org.antlr.v4.Tool %*