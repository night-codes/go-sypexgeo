# go-sypexgeo
SypexGeo API bindings for Go (provides access data from [SypexGeo](https://sypexgeo.net/) IP database files). Accepts only SypexGeo 2.2 or above format.
   
For more information about Sypex Geo databases and their features please visit http://sypexgeo.net.

Direct DB links that might became broken over time:

  * Sypex Geo City DB (free) - http://sypexgeo.net/files/SxGeoCity_utf8.zip
  * Sypex Geo City Max DB - https://sypexgeo.net/ru/buy/


## Documentation
[Docs on godoc.org](https://godoc.org/gopkg.in/night-codes/go-sypexgeo.v1)   


## How To Install   
```bash
go get gopkg.in/night-codes/go-sypexgeo.v1
```

   

## Getting Started
```go
package main

import "gopkg.in/night-codes/go-sypexgeo.v1"

func main() {
    geo := sypexgeo.New("/path/to/db/SxGeoCity.dat")
    info, err := geo.GetCityFull("93.73.35.74")
}
```


## License
   
MIT License   
   
Copyright (C) 2015 Oleksiy Chechel (alex.mirrr@gmail.com)   
   
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:   
   
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.   
   
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
