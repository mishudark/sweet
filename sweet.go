/*
  Copyright (c) 2012 Victor Mishel Vera Sanchez, mishudark<mishu.drk@gmail.com>

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package sweet

import(
  "encoding/json"
  "fmt"
  "github.com/gosexy/sugar"
  "github.com/gosexy/to"
  "strings"
)

type List []interface{}

//List.push(item)
func (self *List) push(data interface{}){
  l := len(*self)
  l_total := l
  if l+1 > cap(*self){
    l_total++
  }

  newList := make(List, l_total)
  newList[l] = data

  copy(newList, *self)
  *self = newList
}


//List.pop()
func (self *List) pop(){
  l := len(*self)
  if l < 1{
    return
  }

  newList := make(List, l-1)
  copy(newList, *self)

  *self = newList
}


//json decoder helper
func json_decode(bytes []byte) (interface{}, error) {
  var result []sugar.Tuple

  err := json.Unmarshal(bytes, &result)
  s := to.String(err)

  //FIXME: there is another way to make this?
  if strings.Contains(s, "json: cannot unmarshal object into Go") {
    var result_simple sugar.Tuple
    err = json.Unmarshal(bytes, &result_simple)

    if err != nil {
      //convert siple to array
      array := [1]sugar.Tuple{result_simple}
      return array, err
    }
  }

  if err != nil {
    fmt.Println("error:", err)
    return result, err
  }
  return result, err
}

/*
func main() {
  var jsonBlob = []byte(`
  [
  {"Name": "Platypus", "Order": "Monotremata"},
  {"Name": "Platypus", "Order": "Monotremata"}
  ]
  `)
  json_decode(jsonBlob)
}
*/
