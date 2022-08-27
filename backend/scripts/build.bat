:: This file builds the application backend from the sources.
@echo off
setlocal enabledelayedexpansion
echo Building MaX Pondus backend

:: Resolve the absolute path of the project root from the script path.
set rootpath=%~dp0%
set rootpath=%rootpath:~0,-9%
cd %rootpath%

:: Resolve and store the version of the Go environment.
for /f "delims=" %%i in ('go version') do set goversion=%%i

:: Resolve the folder where to put all build results.
set buildpath=%rootpath%\bin\server

:: Resolve the folder which contains the main of the application.
set cmdpath=%rootpath%\cmd\server

:: Show the information related to compilation environment.
echo Detected environment information
echo    Project root    %rootpath%
echo    Build path      %buildpath%
echo    Cmd path        %cmdpath%
echo    Go version      %goversion%
echo    Go path         %GOPATH%

:: ------------------------------
:: Source code format and linting
:: ------------------------------

:: Install the golangci-lint if not yet installed.
if not exist %GOPATH%\bin\golangci-lint.exe (
  echo Installing golangci-lint...
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0 || exit /B 1
  echo Installing golangci-lint completed.
)

:: Use goimports to check the format correctness.
echo Linting and validating code...
set lintresult=%rootpath%\lint.out
golangci-lint run || exit /B 1
echo Validating code formatting passed.

:: ----------------------------
:: Run tests and check coverage
:: ----------------------------

set coveragethreshold=95.0%
echo Running tests and checking test coverage (threshold: %coveragethreshold%)...
go test -v -failfast -race -coverprofile coverage.out ./internal/... || (
  echo Running tests failed.
  exit /B 1
)
set coverage=0.0%
for /f "tokens=3" %%i in ('go tool cover -func ./coverage.out') do set coverage=%%i
call :percentage_string_gte %coverage% %coveragethreshold% coveragepassed
if %coveragepassed% equ 0 (
  echo Checking test coverage failed. Coverage %coverage% is less than %coveragethreshold%.
  exit /B 1
)
echo Running tests and checking coverage (threshold: %coveragethreshold%) passed.

:: ----------------------------------
:: Compile the application executable
:: ----------------------------------

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
exit /B 0

:: Compare two percentage strings by checking whether the lhs is greater or equal than rhs. Both
:: strings must be in a numeric string format with one decimal number and with a percentage suffix.
:percentage_string_gte
set lhs=%~1
set rhs=%~2
set lhs_decimal=%lhs:~-1%
set rhs_decimal=%rhs:~-1%
set lhs=%lhs:~0,-2%
set rhs=%rhs:~0,-2%
set lhs=%lhs%%lhs_decimal%
set rhs=%rhs%%rhs_decimal%
set /a lhs=%lhs%
set /a rhs=%rhs%
if %lhs% geq %rhs% ( set %~3=1 ) else ( set %~3=0 )
exit /B 0
