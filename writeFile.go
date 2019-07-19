// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
    //"bufio"
    "fmt"
    //"io/ioutil"
    "os"
    "time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    readLastLine("/tmp/dat2");
    // To start, here's how to dump a string (or just
    // bytes) into a file.
   // d1 := []byte("hello\ngo\n")
    //err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    //check(err)

    // For more granular writes, open a file for writing.
    f, err := os.OpenFile("/tmp/dat2", os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
    check(err)

    // It's idiomatic to defer a `Close` immediately
    // after opening a file.
    defer f.Close()

    // You can `Write` byte slices as you'd expect.
    //d2 := []byte{115, 111, 109, 101, 10}
    //n2, err := f.Write(d2)
    //check(err)
    //fmt.Printf("wrote %d bytes\n", n2)

    // A `WriteString` is also available.
    t := time.Now()
    n3, err := f.WriteString(t.Format(time.RFC3339)+"\n")
    fmt.Printf("wrote %d bytes for time\n", n3)

    // Issue a `Sync` to flush writes to stable storage.
    f.Sync()

    // `bufio` provides buffered writers in addition
    // to the buffered readers we saw earlier.
   // w := bufio.NewWriter(f)
    
    
    //n4, err := w.WriteString(t.Format(time.RFC3339)+"\n")
    //fmt.Printf("wrote %d bytes\n", n4)

    // Use `Flush` to ensure all buffered operations have
    // been applied to the underlying writer.
    //w.Flush()


}
func readLastLine(fname string) {
    file, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fi, err := file.Stat()
    if err != nil {
        fmt.Println(err)
    }

    buf := make([]byte, 26)
    //fmt.Printf("wrote %d bytes for 32\n", len(buf))
    n, err := file.ReadAt(buf, fi.Size()-int64(len(buf)))
    if err != nil {
        fmt.Println(err)
    }
    buf = buf[:n]
    fmt.Printf("%s", buf)

}

