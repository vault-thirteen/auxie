@ECHO OFF

SET EnvVarNamePrefix=MYSQL_TEST_

SET EnvVarNameBase_NetProtocol=PROTO
SET EnvVarNameBase_Driver=DRIVER
SET EnvVarNameBase_Host=HOST
SET EnvVarNameBase_Port=PORT
SET EnvVarNameBase_Database=DB
SET EnvVarNameBase_User=USER
SET EnvVarNameBase_Password=PWD

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_NetProtocol%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Driver%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Host%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Port%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Database%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_User%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Password%
@ECHO ON
REG delete HKCU\Environment /F /V %VarName%
@ECHO OFF
