# belajar-golang-logging
Log file adalah file yang berisikan informasi kejadian dari sebuah sistem
Biasanya dalam log file, terdapat informasi waktu kejadian dan pesan kejadian
Logging adalah aksi menambah informasi log ke log file
Logging sudah menjadi standar industri untuk menampilkan informasi yang terjadi di aplikasi yang kita buat
Logging bukan hanya untuk menampilkan informasi, kadang digunakan untuk proses debugging ketika terjadi masalah di aplikasi kita

# Ekosistem Logging
Aplikasi->Log file->Log Aggregator (fluentd,filebit) ->Log Database(Elasticsearch)->Log Management (Kibana, Graylog)

# Logging Library
Menggunakan Go-Lang sebenarnya kita bisa package log untuk melakukan logging
Hanya saja karena fiturnya terbatas, oleh karena itu kebanyakan programmer tidak menggunakanny

```go
Logrus : https://github.com/sirupsen/logrus 
Zap : https://github.com/uber-go/zap 
Zerolog : https://github.com/rs/zerolog
```
# Logger
Logger adalah struct utama pada Logrus untuk melakukan logging
Untuk membuat Logger, kita bisa menggunakan function New() pada package logrus
Dan hasil dari function New() adalah sebuah pointer Logger

# Level
Dalam Logging, hal yang paling penting adalah Level
Level adah penentuan prioritas atau jenis dari sebuah kejadian
Level itu dimulai dari level terendah sampai level tertinggi
Logrus mendukung banyak sekali level

1.Trace
2.Debug
3.Info
4.Warn
5.Error
6.Fatal
7.Panic

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
# Singleton
Logrus sendiri memiliki singleton object untuk Logger, sehingga kita tidak perlu membuat object Logger sendiri sebenarnya
Namun artinya, jika kita ubah data Logger singleton tersebut, maka secara otomatis yang menggunakan Logger tersebut akan berubah
Secara default, Logger singleton yang ada di logrus menggunakan TextFormatter dan Info Level
Cara menggunakan Logger singleton ini, kita bisa langsung menggunakan package logrus nya saja
