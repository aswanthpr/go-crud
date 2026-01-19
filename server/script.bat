@echo off
setlocal enabledelayedexpansion

REM ============================================================================
REM Goose migration helper for Windows
REM Usage:
REM   script.bat createmigrate <migration_name>
REM   script.bat migrationup
REM   script.bat migratedown
REM   script.bat migratedelete
REM
REM NOTE: This script assumes you run it from the `server` directory.
REM       Migrations will be stored in `database\migration`.
REM       Requires `goose` in your PATH.
REM       DB URL: sqlserver://localhost?database=crud_app&trusted_connection=true
REM ============================================================================

set "MIGRATION_DIR=database\migrations"
set "DB_DRIVER=sqlserver"
REM Keep '&' unescaped; safe inside set "VAR=..."
REM (Escaping with '^' breaks when launched from PowerShell)
set "DB_URL=sqlserver://localhost?database=crud_app&trusted_connection=true"
set "DB_NAME=crud_app"
set "DB_SERVER=localhost"

if "%~1"=="" (
    echo Usage:
    echo   %~nx0 createmigrate ^<migration_name^>
    echo   %~nx0 migrateup
    echo   %~nx0 migratedown
    echo   %~nx0 migratedelete
    exit /b 1
)

if /i "%~1"=="createmigrate" goto :CreateMigrate
if /i "%~1"=="migrateup" goto :MigrateUp
if /i "%~1"=="migrateup" goto :MigrateUp
if /i "%~1"=="migratedown" goto :MigrationDown
if /i "%~1"=="migratedelete" goto :MigrationDelete

echo Unknown command: %~1
echo Valid commands: createmigrate, migrationup ^(or migrateup^), migratedown, migratedelete
exit /b 1

:EnsureDir
if not exist "%MIGRATION_DIR%" (
    mkdir "%MIGRATION_DIR%"
)
goto :eof

:EnsureDatabase
REM Best-effort: create database if missing (requires sqlcmd)
where sqlcmd >nul 2>nul
if errorlevel 1 (
    REM sqlcmd not found; can't auto-create DB
    goto :eof
)
sqlcmd -S "%DB_SERVER%" -E -b -Q "IF DB_ID(N'%DB_NAME%') IS NULL BEGIN CREATE DATABASE [%DB_NAME%] END"
goto :eof

:CreateMigrate
if "%~2"=="" (
    echo Missing migration name.
    echo Example: %~nx0 createmigrate create_users_table
    exit /b 1
)

call :EnsureDir
echo Creating new goose migration "%~2" in "%MIGRATION_DIR%"...
goose -dir "%MIGRATION_DIR%" create "%~2" sql
exit /b %ERRORLEVEL%

:MigrateUp
call :EnsureDir
call :EnsureDatabase
echo Applying all up migrations...
goose -dir "%MIGRATION_DIR%" %DB_DRIVER% "%DB_URL%" up
exit /b %ERRORLEVEL%

:MigrationDown
call :EnsureDir
call :EnsureDatabase
echo Rolling back all migrations (down to version 0)...
goose -dir "%MIGRATION_DIR%" %DB_DRIVER% "%DB_URL%" down-to 0
exit /b %ERRORLEVEL%

:MigrationDelete
call :EnsureDir
call :EnsureDatabase

REM First roll back the last migration (one step down)
echo Rolling back last migration...
goose -dir "%MIGRATION_DIR%" %DB_DRIVER% "%DB_URL%" down
if errorlevel 1 (
    echo Failed to roll back last migration. Aborting delete.
    exit /b 1
)

REM Find the latest migration file (by name, descending)
set "LAST_MIGRATION_FILE="
for /f "delims=" %%F in ('dir /b /a-d /o-n "%MIGRATION_DIR%\*.sql" 2^>nul') do (
    if not defined LAST_MIGRATION_FILE (
        set "LAST_MIGRATION_FILE=%%F"
    )
)

if not defined LAST_MIGRATION_FILE (
    echo No migration files found to delete.
    exit /b 0
)

echo Deleting last migration file: "%MIGRATION_DIR%\!LAST_MIGRATION_FILE!"...
del "%MIGRATION_DIR%\!LAST_MIGRATION_FILE!"
exit /b %ERRORLEVEL%

