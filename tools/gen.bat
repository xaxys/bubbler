@ECHO OFF
SET cur_dir=%~dp0
%cur_dir%antlr.bat -Dlanguage=Go -package parser -o %cur_dir%..\parser -no-listener -visitor %cur_dir%..\bubbler.g4