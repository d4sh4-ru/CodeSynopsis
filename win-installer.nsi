; installer.nsi - Скрипт для создания установщика CodeSynopsis

; Настройки установщика
Name "CodeSynopsis"
OutFile "build\CodeSynopsis-Installer.exe"
InstallDir "$LOCALAPPDATA\CodeSynopsis"  ; Локальная установка для пользователя
RequestExecutionLevel user  ; Не требуются права администратора

; Включить современный интерфейс
!include "MUI2.nsh"

!define MUI_ICON "build\ico.ico"

; Язык интерфейса
!insertmacro MUI_LANGUAGE "English"

; Настройка текста для английского языка
LangString MUI_TEXT_WELCOME_INFO_TITLE ${LANG_ENGLISH} "Welcome to the CodeSynopsis Setup Wizard"
LangString MUI_TEXT_WELCOME_INFO_TEXT ${LANG_ENGLISH} "This wizard will guide you through the installation of CodeSynopsis."
LangString MUI_TEXT_DIRECTORY_TITLE ${LANG_ENGLISH} "Choose Install Location"
LangString MUI_TEXT_DIRECTORY_SUBTITLE ${LANG_ENGLISH} "Select the folder where CodeSynopsis will be installed."
LangString MUI_TEXT_INSTALLING_TITLE ${LANG_ENGLISH} "Installing"
LangString MUI_TEXT_INSTALLING_SUBTITLE ${LANG_ENGLISH} "Please wait while CodeSynopsis is being installed."
LangString MUI_TEXT_FINISH_INFO_TITLE ${LANG_ENGLISH} "Installation Complete"
LangString MUI_TEXT_FINISH_INFO_TEXT ${LANG_ENGLISH} "CodeSynopsis has been successfully installed on your system!"
LangString MUI_BUTTONTEXT_FINISH ${LANG_ENGLISH} "Finish"

; Отключение неиспользуемых элементов
!undef MUI_HEADER_TEXT
!undef MUI_HEADER_SUBTEXT
!undef MUI_BRANDING_TEXT

; Определение страниц установщика
!insertmacro MUI_PAGE_WELCOME
!insertmacro MUI_PAGE_DIRECTORY
!insertmacro MUI_PAGE_INSTFILES
!insertmacro MUI_PAGE_FINISH

!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES
!insertmacro MUI_UNPAGE_FINISH

; Секция установки
Section "Install"
    ; Создаем основную директорию
    SetOutPath "$INSTDIR"
    File "LICENSE"
    File "README.md"
    File "build\modify_path.exe"
    File "build\ico.ico"

    ; Создаем поддиректорию bin и копируем переименованный исполняемый файл
    CreateDirectory "$INSTDIR\bin"
    SetOutPath "$INSTDIR\bin"
    File "/oname=cs.exe" "build\CodeSynopsis-windows-amd64.exe"

    ; Добавляем путь к bin в пользовательскую переменную PATH
    nsExec::Exec '"$INSTDIR\modify_path.exe" add "$INSTDIR\bin"'

    ; Регистрация в реестре
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "DisplayName" "CodeSynopsis"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "UninstallString" "$\"$INSTDIR\uninstall.exe$\""
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "InstallLocation" "$INSTDIR"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "Publisher" "d4sh4-ru"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "DisplayVersion" "0.1.0"
    WriteRegStr HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "DisplayIcon" "$INSTDIR\ico.ico"
    WriteRegDWORD HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "NoModify" 1
    WriteRegDWORD HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis" "NoRepair" 1

    ; Создание uninstaller
    WriteUninstaller "$INSTDIR\uninstall.exe"
SectionEnd

; Секция удаления
Section "Uninstall"
    ; Удаляем путь из пользовательской переменной PATH
    nsExec::Exec '"$INSTDIR\modify_path.exe" remove "$INSTDIR\bin"'

    ; Удаляем файлы
    Delete "$INSTDIR\bin\cs.exe"
    Delete "$INSTDIR\modify_path.exe"
    Delete "$INSTDIR\LICENSE"
    Delete "$INSTDIR\README.md"
    Delete "$SMPROGRAMS\CodeSynopsis\CodeSynopsis.lnk"
    Delete "$INSTDIR\uninstall.exe"

    ; Удаляем записи из реестра о приложении
    DeleteRegKey HKCU "Software\Microsoft\Windows\CurrentVersion\Uninstall\CodeSynopsis"
    
    ; Удаляем папки
    RMDir "$INSTDIR\bin"
    RMDir "$INSTDIR"
SectionEnd