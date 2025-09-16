# microhod/warppipe

<img src="./img/logo.png" width="64"/>

Package `microhod/warppipe` implements a queue to share any go types across processes using named pipes.

> **NOTE:**
> windows is not currently supported

## Usage

Writing

```go
import "github.com/micrhohod/warppipe"

type Foo struct {
    Bar string
}

func main() {
    w, err := warppipe.NewWriter[Custom]()
    if err != nil {
        panic(err)
    }
    defer close(w)

    v := Foo{
        Bar: "baz",
    }
    if err := w.Write(v); err != nil {
        panic(err)
    }
}
```

Reading

```go
import "github.com/micrhohod/warppipe"

type Foo struct {
    Bar string
}

func main() {
    r, err := warppipe.NewReader[Custom]()
    if err != nil {
        panic(err)
    }
    defer close(r)

    var v Foo
    if err := r.Read(&v); err != nil {
        panic(err)
    }
    fmt.Println(v)
}
```
