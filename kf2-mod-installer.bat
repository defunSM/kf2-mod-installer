@echo off
setlocal

echo DO NOT CLOSE UNTIL FINISHED
echo CHOOSE WHERE YOUR 'KillingFloor2' FOLDER IS (DO NOT CHOOSE KFGAME FOLDER)

set "psCommand="(new-object -COM 'Shell.Application')^
.BrowseForFolder(0,'Please choose the KillingFloor2 Folder.',0,0).self.path""

for /f "usebackq delims=" %%I in (`powershell %psCommand%`) do set "folder=%%I"

setlocal enabledelayedexpansion
echo You chose !folder! to install KF2 mods.

echo downloading....
powershell -Command "(New-Object Net.WebClient).DownloadFile('https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j', 'KFGame.zip')"
::powershell -Command "wget --load-cookies /tmp/cookies.txt "https://docs.google.com/uc?export=download"&"confirm=$(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate 'https://docs.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j' -O- | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')"&"id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j" -O KFGame.zip && rm -rf /tmp/cookies.txt"
::move KFGame.zip !folder!
::powershell -Command "Start-Process cmd -Verb RunAs"
::curl -c /tmp/cookies "https://drive.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j" > /tmp/intermezzo.html
::curl -L -b /tmp/cookies "https://drive.google.com$(cat /tmp/intermezzo.html | grep -Po 'uc-download-link" [^>]* href="\K[^"]*' | sed 's/\&amp;/\&/g')" > KFGame.zip
:: https://anonfiles.com/B1SaX7Hfof/KFGame_zip
::wget --load-cookies /tmp/cookies.txt "https://docs.google.com/uc?export=download&confirm=$(wget --quiet --save-cookies /tmp/cookies.txt --keep-session-cookies --no-check-certificate 'https://docs.google.com/uc?export=download&id=1yQzYTafK3aLS0HMmDR7OJBUHsl7tuz9j' -O- | sed -rn 's/.*confirm=([0-9A-Za-z_]+).*/\1\n/p')&id=" -O KFGame.zip && rm -rf /tmp/cookies.txt


echo extracting....

cd /d %~dp0
Call :UnZipFile "!folder!" "KFGame.zip"
exit /b

:UnZipFile !folder! KFGame.zip
set vbs="%temp%\_.vbs"
if exist %vbs% del /f /q %vbs%
>%vbs%  echo Set fso = CreateObject("Scripting.FileSystemObject")
>>%vbs% echo If NOT fso.FolderExists(%1) Then
>>%vbs% echo fso.CreateFolder(%1)
>>%vbs% echo End If
>>%vbs% echo set objShell = CreateObject("Shell.Application")
>>%vbs% echo set FilesInZip=objShell.NameSpace(%2).items
>>%vbs% echo objShell.NameSpace(%1).CopyHere(FilesInZip)
>>%vbs% echo Set fso = Nothing
>>%vbs% echo Set objShell = Nothing
cscript //nologo %vbs%
if exist %vbs% del /f /q %vbs%


::powershell Expand-Archive -Force KFGame.zip -DestinationPath !folder!
echo KF2 MODS DOWNLOAD COMPLETED 
echo YOU MAY CLOSE THIS NOW...
echo -- Takaeshi


PAUSE
endlocal