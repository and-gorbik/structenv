## structenv
Библиотека для чтения переменных окружения в поля структур.

### Примеры использования

#### Чтение переменной окружения из значения поля

```go
type Creds struct {
    Login       string
    Password    string `env:""`
}

c := Creds{
    Login: "nobody",
    Password: "PASSWORD",
}

os.Setenv("PASSWORD", "12345")
SetFromEnvs(&c)

fmt.Println(c.Password) // 12345
```

#### Чтение переменной окружения из значения тега

```go
type Creds struct {
    Login       string
    Password    string `env:"PASSWORD"`
}

c := Creds{
    Login: "nobody",
}

os.Setenv("PASSWORD", "12345")
SetFromEnvs(&c)

fmt.Println(c.Password) // 12345
```

#### Работа с вложенными структурами

### Дополнительные возможности

#### Типизированные ошибки

#### Переопределение тела ошибок

#### Переопределение имени тега 