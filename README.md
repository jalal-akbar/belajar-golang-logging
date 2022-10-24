# belajar-golang-logging
- Log file adalah file yang berisikan informasi kejadian dari sebuah sistem
- Biasanya dalam log file, terdapat informasi waktu kejadian dan pesan kejadian
- Logging adalah aksi menambah informasi log ke log file
- Logging sudah menjadi standar industri untuk menampilkan informasi yang terjadi di aplikasi yang kita buat
- Logging bukan hanya untuk menampilkan informasi, kadang digunakan untuk proses debugging ketika terjadi masalah di aplikasi kita

# Ekosistem Logging
- Aplikasi->Log file->Log Aggregator (fluentd,filebit) ->Log Database(Elasticsearch)->Log Management (Kibana, Graylog)

# Logging Library
- Menggunakan Go-Lang sebenarnya kita bisa package log untuk melakukan logging
- Hanya saja karena fiturnya terbatas, oleh karena itu kebanyakan programmer tidak menggunakanny

- Logrus :
```go
https://github.com/sirupsen/logrus
```
- Zap : 
```go
https://github.com/uber-go/zap
```
- Zerolog : 
```go
https://github.com/rs/zerolog
```  

# Logger
- Logger adalah struct utama pada Logrus untuk melakukan logging
- Untuk membuat Logger, kita bisa menggunakan function New() pada package logrus
- Dan hasil dari function New() adalah sebuah pointer Logger

# Level
- Dalam Logging, hal yang paling penting adalah Level
- Level adah penentuan prioritas atau jenis dari sebuah kejadian
- Level itu dimulai dari level terendah sampai level tertinggi
- Logrus mendukung banyak sekali level

- (1).Trace
- (2).Debug
- (3).Info
- (4).Warn
- (5).Error
- (6).Fatal
- (7).Panic

#  Output
- Secara default, output tujuan log yang kita kirim via Logrus adalah ke Console
- Kadang kita butuh mengubah output tujuan log, misal ke File atau Database
- Output Logger adalah io.Writer, jadi kita bisa menggunakan io.Writer dari Go-Lang atau bisa membuat sendiri mengikuti kontrak io.Writer


# Formatter
- Saat Logger mengirim data ke Output, log yang kita kirim akan diformat menggunakan object Formatter
- Logrus secara default memiliki dua formatter :
- TextFormatter, yang secara default digunakan
- JSONFormatter, yang bisa digunakan untuk memformat pesan log menjadi data JSON
- Untuk mengubah formatter, kita bisa gunakan function logger.SetFormatter()

# Field
1. Field
- Saat kita mengirim informasi log, kadang kita ingin menyisipkan sesuatu pada log tersebut
- Misal saja, kita ingin menyisipkan informasi siapa yang login di log nya
- Cara manual kita bisa menambahkan informasi di message nya, tapi Logrus menyediakan cara yang lebih baik, yaitu menggunakan fitur Field
- Dengan fitur Field, kita bisa menambahkan Field tambahan di informasi Log yang kita kirim
- Sebelum melakukan logging, kita bisa gunakan function logger.WithField() untuk menambahkan Field yang kita inginkan
2. Fields
- Atau, kita juga bisa langsung memasukkan beberapa Field dengan menggunakan Fields
- Fields adalah alias untuk map[string]interface{}
- Caranya kita bisa menggunakan function logger.WithFields()


# Entry
- Entry adalah sebuah Struct representasi dari log yang kita kirim menggunakan Logrus Logger
- Setiap log yang kita kirim, maka akan dibuatkan object Entry
- Contohnya ketika kita membuat Formatter sendiri, maka parameter yang kita dapat untuk melakukan formatting bukanlah string message, melainkan - object Entry
```go
https://github.com/sirupsen/logrus/blob/master/entry.go
``` 
- Untuk membuat entry, kita bisa menggunakan function logrus.NewEntry()

# Hook
- Hook adalah sebuah Struct yang bisa kita tambahkan ke Logger sebagai callback yang akan dieksekusi ketika terdapat kejadian log untuk level tertentu
- Contohnya misal, ketika ada log error, kita ingin mengirim notifikasi via chat ke programmer, dan lain-lain
- Kita bisa menambah Hook ke Logger dengan menggunakan function logger.AddHook()
- Dan kita juga bisa menambahkan lebih dari satu Hook ke Logger


# Singleton
- Logrus sendiri memiliki singleton object untuk Logger, sehingga kita tidak perlu membuat object Logger sendiri sebenarnya
- Namun artinya, jika kita ubah data Logger singleton tersebut, maka secara otomatis yang menggunakan Logger tersebut akan berubah
- Secara default, Logger singleton yang ada di logrus menggunakan TextFormatter dan Info Level
- Cara menggunakan Logger singleton ini, kita bisa langsung menggunakan package logrus nya saja
