# Development Infrastructure

Various development infrastructure scripts.

### Docker-Compose notes
Relative paths in 'Docker Compose' configuration are relative to the 
configuration file, not the command path.

### Legacy products' notes

- `MinIO` storage is now not maintained.  
- Modern versions of `Kafka` do not require `ZooKeeper` and other shit.
- `PostgreSQL` is garbage, use `MySQL` instead.
