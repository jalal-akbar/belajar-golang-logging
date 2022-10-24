# belajar-golang-logging
# Ekosistem Loggin
Aplikasi->Log file->Log Aggregator (fluentd,filebit) ->Log Database(Elasticsearch)->Log Management (Kibana, Graylog)

# Logging Library
```go
Logrus : https://github.com/sirupsen/logrus 
Zap : https://github.com/uber-go/zap 
Zerolog : https://github.com/rs/zerolog
```
# Logger
# Level
    - Urutan Level
        1.Trace
        2.Debug
        3.Info
        4.Warn
        5.Error
        5.Fatal
        6.Panic
#  Output
    Secara default, output tujuan log yg kita kirim via logrus adalah ke console
    Output Logger adalah io.Writer
# Formatter
    Text Formatter, Default
    JSON Formatter
# Field
# Entry
```go
https://github.com/sirupsen/logrus/blob/master/entry.go
```