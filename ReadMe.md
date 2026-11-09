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
| 3  | **bencode**            | [bencode](bencode/ReadMe.md)                       | A library for encoding and decoding the _Bencode_ format.           |
| 4  | **Bit**                | [Bit](bit/ReadMe.md)                               | A simple library to work with bits in _Go_ programming language.    |
| 5  | **BOM**                | [BOM](BOM/ReadMe.md)                               | Functions and methods to work with _Unicode_ byte order mark.       |
| 6  | **Boolean**            | [Boolean](boolean/ReadMe.md)                       | A boolean-to-string parser.                                         |
| 7  | **Cache**              | [Cache](Cache/ReadMe.md)                           | A cache with several features.                                      |
| 8  | **cio**                | [cio](cio/ReadMe.md)                               | Console input & output.                                             |
| 9  | **counter**            | [counter](counter/ReadMe.md)                       | A counter with a practically unlimited value.                       |
| 10 | **csv**                | [csv](csv/ReadMe.md)                               | A writer for _CSV_ format.                                          |     
| 11 | **dev-inf**            | [Dev Inf](dev-inf/ReadMe.md)                       | Various development infrastructure scripts.                         |
| 12 | **endianness**         | [endianness](endianness/ReadMe.md)                 | Detection of "endianness" of a machine.                             |
| 13 | **Env**                | [Env](env/ReadMe.md)                               | A simple library to work with environment variables.                |
| 14 | **Errors**             | [Errors](errors/ReadMe.md)                         | A simple library to work with built-in errors.                      |
| 15 | **File**               | [File](file/ReadMe.md)                             | A simple library to work with files and folders.                    |
| 16 | **hash**               | [hash](hash/ReadMe.md)                             | Helper functions for hash sums.                                     |
| 17 | **Header**             | [Header](header/ReadMe.md)                         | _IANA_ registered message header names.                             |
| 18 | **HTML-DOM**           | [HTML-DOM](HTML-DOM/ReadMe.md)                     | A library to work with _HTML_ document object model.                |
| 19 | **HTTP Helper**        | [HTTP Helper](http-helper/ReadMe.md)               | A package that helps in processing and testing _HTTP_ requests.     |
| 20 | **IPA**                | [IPA](IPA/ReadMe.md)                               | Methods to work with Internet Protocol Addresses.                   |
| 21 | **MemUC**              | [MemUC](MemUC/ReadMe.md)                           | A library to cope with memory usage in _Go_ programming language.   |
| 22 | **MIME**               | [MIME](MIME/ReadMe.md)                             | MIME types and basic functions to work with them.                   |
| 23 | **MySQL-IT**           | [MySQL-IT](MySQL-IT/ReadMe.md)                     | A template for an application importing data into _MySQL_ database. |
| 24 | **NTS**                | [NTS](NTS/ReadMe.md)                               | Methods for processing null-terminated strings.                     |
| 25 | **Number**             | [Number](number/ReadMe.md)                         | A simple library to work with numbers in _Go_ programming language. |
| 26 | **os**                 | [os](os/ReadMe.md)                                 | Various objects for interaction with operating systems.             |
| 27 | **queue**              | [queue](queue/ReadMe.md)                           | A simple queue.                                                     |
| 28 | **Random**             | [Random](random/ReadMe.md)                         | A simple library to work with random integer numbers.               |
| 29 | **Range**              | [Range](range/ReadMe.md)                           | Methods to work with ranges of floating point numbers.              |
| 30 | **Reader**             | [Reader](reader/ReadMe.md)                         | Auxiliary methods and functions for reading.                        |
| 31 | **ReaderSeeker**       | [ReaderSeeker](ReaderSeeker/ReadMe.md)             | An interface which combines _io.Reader_ and _io.Seeker_ interfaces. |
| 32 | **RPoFS**              | [RPoFS](rpofs/ReadMe.md)                           | Random Password of Fixed Size.                                      |
| 33 | **RS**                 | [RS](rs/ReadMe.md)                                 | A Seeker extension to the reader package.                           |
| 34 | **SAAC**               | [SAAC](SAAC/ReadMe.md)                             | Simple Arithmetic Average Calculator.                               |
| 35 | **Simple-File-Server** | [Simple-File-Server](Simple-File-Server/ReadMe.md) | A simple file server library.                                       |
| 36 | **SLReader**           | [SLReader](SLReader/ReadMe.md)                     | A reader based on io.Reader and using speed limits.                 |
| 37 | **SMA**                | [SMA](SMA/ReadMe.md)                               | Simple Moving Average.                                              |
| 38 | **SMS**                | [SMS](SMS/ReadMe.md)                               | Simple Merge Sorter.                                                |
| 39 | **SQL**                | [SQL](SQL/ReadMe.md)                               | Various _SQL_ helpers.                                              |
| 40 | **SSC**                | [SSC](SSC/ReadMe.md)                               | Simple Stupid Cache.                                                |
| 41 | **Start-Controller**   | [Start-Controller](Start-Controller/ReadMe.md)     | A start-stop controlling mechanism.                                 |
| 42 | **task-queue**         | [task-queue](task-queue/ReadMe.md)                 | Task queue.                                                         |
| 43 | **Tester**             | [Tester](tester/ReadMe.md)                         | Various useful functions to perform tests.                          |
| 44 | **Time**               | [Time](time/ReadMe.md)                             | A simple library to work with time in _Go_ programming language.    |
| 45 | **Time Metadata**      | [Time Metadata](time-metadata/ReadMe.md)           | An auxiliary type for storing time of creation and time of update.  |
| 46 | **TSB**                | [TSB](TSB/ReadMe.md)                               | Tri-State Boolean.                                                  |
| 47 | **Unicode**            | [Unicode](unicode/ReadMe.md)                       | _Unicode_ symbols processing package.                               |
| 48 | **VCS**                | [VCS](VCS/ReadMe.md)                               | A collection of helpers for version control systems.                |
| 49 | **Versioneer**         | [Versioneer](Versioneer/ReadMe.md)                 | Product and dependency version helper.                              |
| 50 | **Zipper**             | [Zipper](zipper/ReadMe.md)                         | Functions to work with compressed files.                            |

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
