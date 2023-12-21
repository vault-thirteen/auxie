SETLOCAL

:: Distribution name may not contain a space symbol !
SET Distribution=Debian

SET WslIsoFolder=J:\WSL\ISO
SET ImagePath=%WslIsoFolder%\Debian\wsl-debian.tar
SET WslVmFolder=J:\WSL\VM
SET MachinePath=%WslVmFolder%\Debian\wsl-debian

wsl --install

wsl -l -v
wsl --update

wsl --list
wsl --list --online
wsl --install -d %Distribution%
wsl --setdefault %Distribution%

:: Move the VM to another place.
wsl --export %Distribution% "%ImagePath%"
wsl --unregister %Distribution%
wsl --import %Distribution% "%MachinePath%" "%ImagePath%"
wsl --setdefault %Distribution%
ENDLOCAL

:: Docker Desktop for Windows is bugged.
:: https://github.com/docker/for-win/issues/13345
::
:: WSL in Windows 10 is still bugged.
:: https://github.com/microsoft/WSL/issues/5895
::
:: What the hell is going on ?
