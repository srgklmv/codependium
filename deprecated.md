## Термины

### Общее

- **IaaS (Infrastructure as a Service)** - разновидность облачных сервисов, в которой
  клиенту предоставляются вычислительные мощности для настройки платформы (ОС, виртуализация и пр.),
  деплоя и работы приложения, хранения данных.
- **PaaS (Platform as a Service)** - разновидность облачных сервисов, в которой
  клиенту предоставляется всё для работы его приложения и хранения данных.
- **SaaS (Software as a Service)** - разновидность облачных сервисов, в которой
  предоставляется сервис целиком для реализации потребностей клиента.


## Crib

### Golang

#### Tour of Go - Language essentials

Programs are made up of packages. Main is main.

Caps is an export tools.

Outside functions every statement must begin with a Go keyword.

Basic types:

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
// represents a Unicode code point

float32 float64

complex64 complex128
```

"The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type."

`int`, `float64` & `complex128` is default in short declaration with untyped values.

You can store high-precision numeric values in constants. Unless variables, which are limited by types.

