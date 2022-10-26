## Concurrency: Filter and Fill data

### Warning

Sebelum mengerjakan ini, pastikan kalian sudah belajar materi berikut:

- Goroutine
- Channel dan buffered channel

### Description

Terdapat data yang berisi informasi tentang domain dari sebuah website. Data tersebut akan di simpan ke dalam sebuah struct yang berisi beberapa _field_ :

- `RankWebsite`, akan berisi informasi tentang peringkat dari sebuah website bertipe `int`
- `Domain`, akan berisi informasi tentang domain dari sebuah website bertipe `string`, seperti `google.com` atau `yahoo.co.id`
- `Valid` akan berisi informasi apakah domain tersebut valid atau tidak bertipe `bool`
- `RefIPs` akan berisi informasi tentang total IP yang merujuk ke domain tersebut bertipe `int`
- `TLD` (Top Level Domain), akan berisi informasi tentang TLD dari sebuah website bertipe `string`, seperti `.com` atau `.gov`
- `IDN_TLD` (Indonesia Top Level Domain), akan berisi informasi tentang TLD khusus di wilayah Indonesia dari sebuah website bertipe `string`, seperti `.co.id` atau `.go.id`

```go
type RowData struct {
    RankWebsite int
    Domain      string
    Valid       bool
    RefIPs      int
    TLD         string
    IDN_TLD     string
}
```

Namun, untuk saat ini data yang tersimpan hanya berisi di field `RankWebsite` dan `Domain`, `Valid` dan `RefIPs`. Kalian diminta untuk mengisi dan mengecek data tersebut dengan menggunakan goroutine dan channel.

Untuk mengisi `TLD`, kita dapat mengambil data dari _field_ `Domain`. Contoh, jika domain tersebut berisi `google.com` maka `TLD` nya adalah `.com`. Contoh lain, jika domain tersebut berisi `bukanruang.org` maka `TLD` nya adalah `.org`.

Untuk mendapatkan data `IDN_TLD` hampir sama seperti `TLD`, hanya saja kita harus memeriksa apakah domain tersebut ada kemungkinan memiliki TLD khusus di Indonesia atau tidak. Dalam kasus ini, terdapat 3 jenis TLD yang pasti memiliki TLD khusus di Indonesia, yaitu `.com`, `.org`, dan `.gov`.

- Jika domain tersebut memiliki TLD `.com`, maka `IDN_TLD` nya adalah `.co.id`
- Jika domain tersebut memiliki TLD `.org`, maka `IDN_TLD` nya adalah `.org.id`
- Jika domain tersebut memiliki TLD `.gov`, maka `IDN_TLD` nya adalah `.go.id`

Untuk jenis `TLD` yang lain, maka `IDN_TLD` nya adalah sama dengan `TLD`.

### Task

Dapatkan data `TLD` dari `Domain` yang ada di data tersebut. dengan menggunakan fungsi `ProcessGetTLD` yang sudah disediakan. Fungsi ini akan dijalankan secara _concurrent_ dengan menggunakan goroutine. Terdapat parameter `website` bertipe `RowData` yang berisi satuan data yang akan diolah. Disediakan juga 2 channel yang digunakan untuk:

- channel pertama akan menerima data dari _struct_ `RowData` yang sudah berisi data `TLD` dan `IDN_TLD`
- Channel ke dua akan menerima data `error` jika terjadi error

Kembalikan error jika di dalam data:

- `Domain` kosong, dengan berisi pesan `domain is empty`
- `Valid` berisi `false`, dengan berisi pesan `domain is not valid`
- `RefIPs` berisi `-1`, dengan berisi pesan `domain has no refIPs`

```go
func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
    // kerjakan di sini
}
```

Kita juga akan melakukan filter terhadap data yang ada. Terdapat fungsi `FilterAndFillData` yang akan menerima 2 parameter

- `TLD` bertipe `string` yang berisi data TLD yang akan di filter di data
- `data` bertipe `[]RowData` yang berisi data yang akan dilengkapi terlebih dahulu di fungsi `ProcessGetTLD` dan kemudian di filter

Didalam fungsi ini terdapat 2 channel yang digunakan melakukan komunikasi antar goroutine di fungsi `ProcessGetTLD`.

```go
    ch := make(chan RowData, len(data))
    errCh := make(chan error)
```

Lengkapi fungsi `FilterAndFillData` agar dapat menjalankan goroutine `ProcessGetTLD` secara _concurrent_ dan mengembalikan data yang sudah di filter berdasarkan `TLD` yang diinginkan.

Kembalikan 2 data, yaitu:

- `[]RowData` yang berisi data yang sudah di filter
- `error` jika terjadi error di salah satu goroutine dengan keterangan yang sudah dijelaskan di fungsi `ProcessGetTLD`

```go
// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
    ch := make(chan RowData, len(data))
    errCh := make(chan error)

    for _, website := range data {
        go func(website RowData) {
            FuncProcessGetTLD(website, ch, errCh)
        }(website)
    }

    // lengkapi di sini
}
```

### Constraints

Berikut batasan test data yang akan digunakan untuk memastikan fungsi yang dibuat sudah sesuai :

- List `data` akan selalu berisi minimal 50 data dan maksimal 100 data (tidak akan lebih dari 100)
- Setiap proses berjalan akan terjadi delay selama 100 milidetik
- Selama menjalankan fungsi ini, tidak boleh ada proses yang berjalan secara _blocking_ dan total eksekusi tidak boleh lebih dari 200 milidetik

### Test Case Examples

#### Test Case 1

Ketika fungsi `FilterAndFillData` berjalan :

**Input**:

```txt
TLD = .com
data = [{
    RankWebsite: 1,
    Domain: "google.com",
    Valid: true,
    RefIPs: 1,
    TLD: "",
    IDN_TLD: "",
}, {
    RankWebsite: 2,
    Domain: "bukanruang.org",
    Valid: true,
    RefIPs: 1,
    TLD: "",
    IDN_TLD: "",
}, {
    RankWebsite: 3,
    Domain: "bukanjudi.xyz",
    Valid: true,
    RefIPs: 1,
    TLD: "",
    IDN_TLD: "",
},{
    ...dst
}]
```

**Expected Output / Behavior**:

Data akan terisi lengkap dan melakukan filter untuk data yang memiliki TLD `.com`, dan tidak terjadi error :

```txt
[{
    RankWebsite: 1,
    Domain: "google.com",
    Valid: true,
    RefIPs: 1,
    TLD: ".com",
    IDN_TLD: ".co.id",
},{
    ...dst
}]

error = nil
```

**Explanation**:

Setiap data akan diisi dengan data `TLD` dan `IDN_TLD` yang sesuai dengan data `Domain` yang ada di dalamnya. Kemudian data yang memiliki TLD `.com` akan di filter dan di kembalikan. Error tidak terjadi karena semua data sudah sesuai

#### Test Case 2

Ketika fungsi `FilterAndFillData` berjalan :

**Input**:

```txt
TLD = .com
data = [{
    RankWebsite: 1,
    Domain: "google.com",
    Valid: false,
    RefIPs: 1,
    TLD: "",
    IDN_TLD: "",
},{
    ...dst
}]
```

**Expected Output / Behavior**:

```txt
[]

error = domain is not valid
```

**Explanation**:

Terdapat data yang memiliki `Valid` bernilai `false`, sehingga akan mengembalikan error dengan keterangan `domain is not valid`
