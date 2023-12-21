# MySQL Importer Template

A template for an application importing data into _MySQL_ database.

This application is not automated. The configuration is hard-coded.  
That is why, this is not a tool, but a template for a tool instead.  

## Why ?

The built-in tools for _MySQL_ do not work properly.

**Example #1.**  
`MySQL Workbench` shows random numbers of rows in a table when the table is 
completely static. 

**Example #2.**  
`MySQL Shell` hangs infinitely during script execution, while this tiny tool 
written in _Go_ language executes the script in several seconds. 
