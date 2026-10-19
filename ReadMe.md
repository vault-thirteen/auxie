# Auxie

A collection of auxiliary types, methods, functions and other small useful 
things.  

The name "**Auxie**" comes from the short name "**AUX**" which in its turn comes 
from a full word "**Auxiliary**". For historical reasons, the word "**AUX**" 
is forbidden for usage as a folder name in operating systems based on _MS-DOS_ 
and _Windows_.  

## List of included libraries

| #  | Folder                 | ReadMe                                             | Description                                                         |
|:---|:-----------------------|:---------------------------------------------------|:--------------------------------------------------------------------|
| 1  | **Array Search**       | [as](as/ReadMe.md)                                 | Various methods for searching in arrays.                            |
| 2  | **Basic Types**        | [BasicTypes](BasicTypes/ReadMe.md)                 | A collection of common sense basic types.                           |
| 3  | **Bit**                | [Bit](bit/ReadMe.md)                               | A simple library to work with bits in _Go_ programming language.    |
| 4  | **BOM**                | [BOM](BOM/ReadMe.md)                               | Functions and methods to work with _Unicode_ byte order mark.       |
| 5  | **Boolean**            | [Boolean](boolean/ReadMe.md)                       | A boolean-to-string parser.                                         |
| 6  | **Cache**              | [Cache](Cache/ReadMe.md)                           | A cache with several features.                                      |
| 7  | **cio**                | [cio](cio/ReadMe.md)                               | Console input & output.                                             |
| 8  | **csv**                | [csv](csv/ReadMe.md)                               | A writer for _CSV_ format.                                          |
| 9  | **dev-inf**            | [Dev Inf](dev-inf/ReadMe.md)                       | Various development infrastructure scripts.                         |
| 10 | **endianness**         | [endianness](endianness/ReadMe.md)                 | Detection of "endianness" of a machine.                             |
| 11 | **Env**                | [Env](env/ReadMe.md)                               | A simple library to work with environment variables.                |
| 12 | **Errors**             | [Errors](errors/ReadMe.md)                         | A simple library to work with built-in errors.                      |
| 13 | **File**               | [File](file/ReadMe.md)                             | A simple library to work with files and folders.                    |
| 14 | **hash**               | [hash](hash/ReadMe.md)                             | Helper functions for hash sums.                                     |
| 15 | **Header**             | [Header](header/ReadMe.md)                         | _IANA_ registered message header names.                             |
| 16 | **HTML-DOM**           | [HTML-DOM](HTML-DOM/ReadMe.md)                     | A library to work with _HTML_ document object model.                |
| 17 | **HTTP Helper**        | [HTTP Helper](http-helper/ReadMe.md)               | A package that helps in processing and testing _HTTP_ requests.     |
| 18 | **IPA**                | [IPA](IPA/ReadMe.md)                               | Methods to work with Internet Protocol Addresses.                   |
| 19 | **MemUC**              | [MemUC](MemUC/ReadMe.md)                           | A library to cope with memory usage in _Go_ programming language.   |
| 20 | **MIME**               | [MIME](MIME/ReadMe.md)                             | MIME types and basic functions to work with them.                   |
| 21 | **MySQL-IT**           | [MySQL-IT](MySQL-IT/ReadMe.md)                     | A template for an application importing data into _MySQL_ database. |
| 22 | **NTS**                | [NTS](NTS/ReadMe.md)                               | Methods for processing null-terminated strings.                     |
| 23 | **Number**             | [Number](number/ReadMe.md)                         | A simple library to work with numbers in _Go_ programming language. |
| 24 | **os**                 | [os](os/ReadMe.md)                                 | Various objects for interaction with operating systems.             |
| 25 | **Random**             | [Random](random/ReadMe.md)                         | A simple library to work with random integer numbers.               |
| 26 | **Range**              | [Range](range/ReadMe.md)                           | Methods to work with ranges of floating point numbers.              |
| 27 | **Reader**             | [Reader](reader/ReadMe.md)                         | Auxiliary methods and functions for reading.                        |
| 28 | **ReaderSeeker**       | [ReaderSeeker](ReaderSeeker/ReadMe.md)             | An interface which combines _io.Reader_ and _io.Seeker_ interfaces. |
| 29 | **RPoFS**              | [RPoFS](rpofs/ReadMe.md)                           | Random Password of Fixed Size.                                      |
| 30 | **RS**                 | [RS](rs/ReadMe.md)                                 | A Seeker extension to the reader package.                           |
| 31 | **SAAC**               | [SAAC](SAAC/ReadMe.md)                             | Simple Arithmetic Average Calculator.                               |
| 32 | **Simple-File-Server** | [Simple-File-Server](Simple-File-Server/ReadMe.md) | A simple file server library.                                       |
| 33 | **SLReader**           | [SLReader](SLReader/ReadMe.md)                     | A reader based on io.Reader and using speed limits.                 |
| 34 | **SMA**                | [SMA](SMA/ReadMe.md)                               | Simple Moving Average.                                              |
| 35 | **SMS**                | [SMS](SMS/ReadMe.md)                               | Simple Merge Sorter.                                                |
| 36 | **SQL**                | [SQL](SQL/ReadMe.md)                               | Various _SQL_ helpers.                                              |
| 37 | **SSC**                | [SSC](SSC/ReadMe.md)                               | Simple Stupid Cache.                                                |
| 38 | **Tester**             | [Tester](tester/ReadMe.md)                         | Various useful functions to perform tests.                          |
| 39 | **Time**               | [Time](time/ReadMe.md)                             | A simple library to work with time in _Go_ programming language.    |
| 40 | **Time Metadata**      | [Time Metadata](time-metadata/ReadMe.md)           | An auxiliary type for storing time of creation and time of update.  |
| 41 | **TSB**                | [TSB](TSB/ReadMe.md)                               | Tri-State Boolean.                                                  |
| 42 | **Unicode**            | [Unicode](unicode/ReadMe.md)                       | _Unicode_ symbols processing package.                               |
| 43 | **VCS**                | [VCS](VCS/ReadMe.md)                               | A collection of helpers for version control systems.                |
| 44 | **Versioneer**         | [Versioneer](Versioneer/ReadMe.md)                 | Product and dependency version helper.                              |
| 45 | **Zipper**             | [Zipper](zipper/ReadMe.md)                         | Functions to work with compressed files.                            |

### Microsoft Windows OS support

_Go_ language 1.20 is the last version supported on old _Windows_ operating 
systems.

As stated in release notes of _Go_ 1.20 located at 
https://go.dev/doc/go1.20#windows, 
> Go 1.20 is the last release that will run on any release of Windows 7, 8, 
> Server 2008 and Server 2012.

As stated in release notes of _Go_ 1.21 located at 
https://go.dev/doc/go1.21#windows,
> Go 1.21 requires at least Windows 10 or Windows Server 2016; support for 
> previous versions has been discontinued.

_Go_ language 1.21 provides some breaking changes, so do not rush into upgrading 
to it at any cost.

Unfortunately, developers of _Go_ language do not support old versions of the language. This is not _Java_.
