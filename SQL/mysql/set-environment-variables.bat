@ECHO OFF

SET EnvVarNamePrefix=MYSQL_TEST_

SET EnvVarNameBase_NetProtocol=PROTO
SET EnvVarNameBase_Driver=DRIVER
SET EnvVarNameBase_Host=HOST
SET EnvVarNameBase_Port=PORT
SET EnvVarNameBase_Database=DB
SET EnvVarNameBase_User=USER
SET EnvVarNameBase_Password=PWD

SET EnvValue_NetProtocol=tcp
SET Env_Value_Driver=mysql
SET Env_Value_Host=localhost
SET Env_Value_Port=3306
SET Env_Value_Database=test
SET Env_Value_User=test
SET Env_Value_Password=test

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_NetProtocol%
@ECHO ON
SETX %VarName% %EnvValue_NetProtocol%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Driver%
@ECHO ON
SETX %VarName% %Env_Value_Driver%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Host%
@ECHO ON
SETX %VarName% %Env_Value_Host%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Port%
@ECHO ON
SETX %VarName% %Env_Value_Port%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Database%
@ECHO ON
SETX %VarName% %Env_Value_Database%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_User%
@ECHO ON
SETX %VarName% %Env_Value_User%
@ECHO OFF

SET VarName=%EnvVarNamePrefix%%EnvVarNameBase_Password%
@ECHO ON
SETX %VarName% %Env_Value_Password%
@ECHO OFF
