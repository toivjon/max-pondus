:: This file builds the application backend from the sources.
@echo off
echo Building MaX Pondus backend

:: Resolve the absolute path of the project root from the script path.
set rootpath=%~dp0%
set rootpath=%rootpath:~0,-9%

:: Resolve and store the version of the Go environment.
for /f "delims=" %%i in ('go version') do set goversion=%%i

:: Resolve the folder where to put all build results.
set buildpath=%rootpath%\bin\backend

:: Resolve the folder which contains the main of the application.
set cmdpath=%rootpath%\cmd\backend

:: Show the information related to compilation environment.
echo Detected environment information
echo    Project root    %rootpath%
echo    Build path      %buildpath%
echo    Cmd path        %cmdpath%
echo    Go version      %goversion%

:: Remove the old build directory to ensure that we get a clean build.
echo Removing the old build directory if it already exists
if exist "%buildpath%" rd /s /q "%buildpath%" || exit /B 1

:: Gather and compile the backend sources files into an executable.
echo Compiling the source files into an executable.
set compilationstart=%time%
mkdir %buildpath%
cd %buildpath%
go build -race %cmdpath% || exit /B 1
set compilationend=%time%

:: Resolve the path of the created executable.
set executablepath=%buildpath%\backend.exe

:: Show information related to compilation.
echo Compilation statistics
echo    Executable      %executablepath%
echo    Start time      %compilationstart%
echo    End time        %compilationend%
echo Build completed.