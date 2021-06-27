:: This file builds the application backend from the sources.
@echo off
setlocal enabledelayedexpansion
echo Building MaX Pondus backend

:: Make Go to install necessary packages globally.
set GO111MODULE=off

:: Resolve the absolute path of the project root from the script path.
set rootpath=%~dp0%
set rootpath=%rootpath:~0,-9%
cd %rootpath%

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

:: -----------------------------
:: Source code format validation
:: -----------------------------

:: Install the goimports if not yet installed.
if not exist %GOPATH%\bin\goimports.exe (
  echo Installing goimports...
  go get golang.org/x/tools/cmd/goimports || exit /B 1
  echo Installing goimports completed.
)

:: Use goimports to check the format correctness.
echo Validating code formatting...
set formatok=1
for /f "skip=1 delims=" %%i in ('goimports -e -d ./') do (echo %%i && set formatok=0)
if %formatok% == 0 (
  echo Validating code formatting failed.
  exit /B 1
)
echo Validating code formatting passed.

:: ---------
:: Run tests
:: ---------

echo Running tests...
go test -v -failfast -race ./internal/... || (
  echo Running tests failed.
  exit /B 1
)
echo Running tests passed.

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
echo Compilation completed:
echo    Executable      %executablepath%
echo    Start time      %compilationstart%
echo    End time        %compilationend%
echo Build completed.