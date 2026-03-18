; ============================================
; SysTrace Server - Installer
; ============================================
!include "MUI2.nsh"
!include "x64.nsh"
!include "nsDialogs.nsh"
!include "LogicLib.nsh"

Name "SysTrace Server"
OutFile "SysTraceServerInstaller.exe"
InstallDir "$PROGRAMFILES64\SysTrace Server"
RequestExecutionLevel admin

!define MUI_ABORTWARNING

Var DB_Host
Var DB_Port
Var DB_User
Var DB_Password
Var DB_Name
Var ServerPort
Var ContainerExists
Var DockerExe

Var Dialog
Var Label_DBHost
Var Label_DBPort
Var Label_DBUser
Var Label_DBPassword
Var Label_DBName
Var Label_ServerPort
Var Input_DBHost
Var Input_DBPort
Var Input_DBUser
Var Input_DBPassword
Var Input_DBName
Var Input_ServerPort

!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
Page custom ConfigPageShow ConfigPageLeave
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES

!insertmacro MUI_LANGUAGE "English"

; ============================================
Function ConfigPageShow
    nsDialogs::Create 1018
    Pop $Dialog

    ${If} $Dialog == error
        Abort
    ${EndIf}

    ${NSD_CreateLabel} 0 0 100% 12u "Server Port:"
    Pop $Label_ServerPort
    ${NSD_CreateText} 0 13u 100% 12u "8080"
    Pop $Input_ServerPort

    ${NSD_CreateLabel} 0 30u 100% 12u "Database Host:"
    Pop $Label_DBHost
    ${NSD_CreateText} 0 43u 100% 12u "localhost"
    Pop $Input_DBHost

    ${NSD_CreateLabel} 0 60u 100% 12u "Database Port:"
    Pop $Label_DBPort
    ${NSD_CreateText} 0 73u 100% 12u "5432"
    Pop $Input_DBPort

    ${NSD_CreateLabel} 0 90u 100% 12u "Database User:"
    Pop $Label_DBUser
    ${NSD_CreateText} 0 103u 100% 12u "systrace"
    Pop $Input_DBUser

    ${NSD_CreateLabel} 0 120u 100% 12u "Database Password:"
    Pop $Label_DBPassword
    ${NSD_CreatePassword} 0 133u 100% 12u "systrace_secure_password"
    Pop $Input_DBPassword

    ${NSD_CreateLabel} 0 150u 100% 12u "Database Name:"
    Pop $Label_DBName
    ${NSD_CreateText} 0 163u 100% 12u "systrace_db"
    Pop $Input_DBName

    nsDialogs::Show
FunctionEnd

Function ConfigPageLeave
    ${NSD_GetText} $Input_ServerPort $ServerPort
    ${NSD_GetText} $Input_DBHost $DB_Host
    ${NSD_GetText} $Input_DBPort $DB_Port
    ${NSD_GetText} $Input_DBUser $DB_User
    ${NSD_GetText} $Input_DBPassword $DB_Password
    ${NSD_GetText} $Input_DBName $DB_Name

    ${If} $DB_Password == ""
        StrCpy $DB_Password "systrace_secure_password"
    ${EndIf}
FunctionEnd

; ============================================
Section "Main Program" SecMain

    ; Hard pre-check: stop installation immediately if Docker is missing or not running.
    StrCpy $DockerExe "docker"
    nsExec::ExecToLog '$DockerExe --version'
    Pop $0

    ; Fallback for systems where PATH is not available in installer context.
    ${If} $0 != 0
        IfFileExists "$PROGRAMFILES\Docker\Docker\resources\bin\docker.exe" 0 +3
            StrCpy $DockerExe '"$PROGRAMFILES\Docker\Docker\resources\bin\docker.exe"'
            nsExec::ExecToLog '$DockerExe --version'
            Pop $0
    ${EndIf}

    ${If} $0 != 0
        MessageBox MB_OK|MB_ICONEXCLAMATION "Docker is not installed or not running.$\nPlease install/start Docker Desktop and run the installer again.$\nhttps://www.docker.com/products/docker-desktop"
        Abort
    ${EndIf}

    nsExec::ExecToLog '$DockerExe info'
    Pop $1
    ${If} $1 != 0
        MessageBox MB_OK|MB_ICONEXCLAMATION "Docker is not installed or not running.$\nPlease install/start Docker Desktop and run the installer again.$\nhttps://www.docker.com/products/docker-desktop"
        Abort
    ${EndIf}

    SetOutPath "$INSTDIR"
    File "..\..\SysTrace_Server.exe"

    File /r "..\..\templates"

    SetOutPath "$INSTDIR\db\init"
    File "..\..\db\create_tables.sql"

    SetOutPath "$INSTDIR\db"
    FileOpen $0 "$INSTDIR\db\docker-compose.yml" w
    FileWrite $0 "version: '3.8'$\r$\n"
    FileWrite $0 "services:$\r$\n"
    FileWrite $0 "  postgres:$\r$\n"
    FileWrite $0 "    image: postgres:15$\r$\n"
    FileWrite $0 "    container_name: systrace_postgres$\r$\n"
    FileWrite $0 "    environment:$\r$\n"
    FileWrite $0 "      POSTGRES_PASSWORD: $DB_Password$\r$\n"
    FileWrite $0 "      POSTGRES_USER: $DB_User$\r$\n"
    FileWrite $0 "      POSTGRES_DB: $DB_Name$\r$\n"
    FileWrite $0 "    ports:$\r$\n"
    FileWrite $0 '      - "$DB_Port:5432"$\r$\n'
    FileWrite $0 "    volumes:$\r$\n"
    FileWrite $0 "      - postgres_data:/var/lib/postgresql/data$\r$\n"
    FileWrite $0 "      - ./init:/docker-entrypoint-initdb.d$\r$\n"
    FileWrite $0 "    restart: unless-stopped$\r$\n"
    FileWrite $0 "volumes:$\r$\n"
    FileWrite $0 "  postgres_data:$\r$\n"
    FileClose $0

    SetOutPath "$INSTDIR"
    FileOpen $0 "$INSTDIR\.env" w
    FileWrite $0 "# PostgreSQL Database Configuration$\r$\n"
    FileWrite $0 "DB_HOST=$DB_Host$\r$\n"
    FileWrite $0 "DB_PORT=$DB_Port$\r$\n"
    FileWrite $0 "DB_USER=$DB_User$\r$\n"
    FileWrite $0 "DB_PASSWORD=$DB_Password$\r$\n"
    FileWrite $0 "DB_NAME=$DB_Name$\r$\n"
    FileWrite $0 "$\r$\n"
    FileWrite $0 "# Server Configuration$\r$\n"
    FileWrite $0 "SERVER_PORT=$ServerPort$\r$\n"
    FileClose $0

    DetailPrint "Docker detected - preparing database..."
    nsExec::ExecToLog '$DockerExe inspect systrace_postgres'
    Pop $ContainerExists
    ${If} $ContainerExists != 0
        DetailPrint "New container detected - initializing volume..."
        nsExec::ExecToLog '$DockerExe compose version'
        Pop $1
        ${If} $1 == 0
            nsExec::ExecToLog '$DockerExe compose -f "$INSTDIR\db\docker-compose.yml" down -v'
            Pop $1
            nsExec::ExecToLog '$DockerExe compose -f "$INSTDIR\db\docker-compose.yml" up -d'
            Pop $2
        ${Else}
            nsExec::ExecToLog 'docker-compose --version'
            Pop $1
            ${If} $1 == 0
                nsExec::ExecToLog 'docker-compose -f "$INSTDIR\db\docker-compose.yml" down -v'
                Pop $1
                nsExec::ExecToLog 'docker-compose -f "$INSTDIR\db\docker-compose.yml" up -d'
                Pop $2
            ${Else}
                StrCpy $2 1
            ${EndIf}
        ${EndIf}
    ${Else}
        DetailPrint "Existing container detected - starting services..."
        nsExec::ExecToLog '$DockerExe compose version'
        Pop $1
        ${If} $1 == 0
            nsExec::ExecToLog '$DockerExe compose -f "$INSTDIR\db\docker-compose.yml" up -d'
            Pop $2
        ${Else}
            nsExec::ExecToLog 'docker-compose --version'
            Pop $1
            ${If} $1 == 0
                nsExec::ExecToLog 'docker-compose -f "$INSTDIR\db\docker-compose.yml" up -d'
                Pop $2
            ${Else}
                StrCpy $2 1
            ${EndIf}
        ${EndIf}
    ${EndIf}

    ${If} $2 != 0
        MessageBox MB_OK|MB_ICONEXCLAMATION "Database could not be started automatically.$\nPlease start Docker Desktop and run the compose command manually."
    ${Else}
        DetailPrint "Database started successfully."
    ${EndIf}

    WriteUninstaller "$INSTDIR\Uninstall.exe"

    CreateDirectory "$SMPROGRAMS\SysTrace Server"
    CreateShortcut "$SMPROGRAMS\SysTrace Server\SysTrace Server.lnk" "$INSTDIR\SysTrace_Server.exe"
    CreateShortcut "$SMPROGRAMS\SysTrace Server\Uninstall.lnk" "$INSTDIR\Uninstall.exe"

    CreateShortCut "$DESKTOP\SysTrace Server.lnk" "$INSTDIR\SysTrace_Server.exe"

    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server" "DisplayName" "SysTrace Server"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server" "UninstallString" "$INSTDIR\Uninstall.exe"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server" "InstallLocation" "$INSTDIR"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server" "DisplayVersion" "1.0.0"
    WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server" "Publisher" "Elias"

SectionEnd

; ============================================
Section "Uninstall"

    StrCpy $DockerExe "docker"
    nsExec::ExecToLog '$DockerExe --version'
    Pop $0
    ${If} $0 != 0
        IfFileExists "$PROGRAMFILES\Docker\Docker\resources\bin\docker.exe" 0 +3
            StrCpy $DockerExe '"$PROGRAMFILES\Docker\Docker\resources\bin\docker.exe"'
            nsExec::ExecToLog '$DockerExe --version'
            Pop $0
    ${EndIf}

    ${If} $0 == 0
        nsExec::ExecToLog '$DockerExe compose version'
        Pop $0
        ${If} $0 == 0
            nsExec::ExecToLog '$DockerExe compose -f "$INSTDIR\db\docker-compose.yml" down'
            Pop $0
        ${Else}
            nsExec::ExecToLog 'docker-compose --version'
            Pop $0
            ${If} $0 == 0
                nsExec::ExecToLog 'docker-compose -f "$INSTDIR\db\docker-compose.yml" down'
                Pop $0
            ${EndIf}
        ${EndIf}
    ${EndIf}

    Delete "$INSTDIR\SysTrace_Server.exe"
    Delete "$INSTDIR\.env"
    RMDir /r "$INSTDIR\templates"
    RMDir /r "$INSTDIR\db"
    Delete "$INSTDIR\Uninstall.exe"
    RMDir "$INSTDIR"

    Delete "$DESKTOP\SysTrace Server.lnk"
    Delete "$SMPROGRAMS\SysTrace Server\SysTrace Server.lnk"
    Delete "$SMPROGRAMS\SysTrace Server\Uninstall.lnk"
    RMDir "$SMPROGRAMS\SysTrace Server"

    DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\SysTrace_Server"

SectionEnd
