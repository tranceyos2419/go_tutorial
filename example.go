//* Example

// URL: https://qiita.com/nyamage/items/46f17318473aa141ffd5

package main

// type JsonTestDate struct {
//   Time string `json:"time"`
//   Msec int64  `json:"milliseconds_since_epoch"`
//   Date string `json:"date"`
// }

// func main(){
//   useDecoder := false
//   flag.BoolVar(&useDecoder, "decoder", false, "use json decoder")
//   flag.Parse()

//   res, err := http.Get("http://date.jsontest.com/")
//   if err != nil {
//     log.Fatal(err)
//   }
//   defer res.Body.Close()
//   if res.StatusCode != 200 {
//     fmt.Println("StatusCode=%d", res.StatusCode)
//     return
//   }
//   //Dump response status line
//   fmt.Printf("======Status line======\n")
//   fmt.Printf("%s %s\n", res.Proto, res.Status)

//   //Dump response header
//   fmt.Printf("======Header======\n")
//   for k, v := range res.Header {
//     fmt.Printf("%s:%s\n", k, v)
//   }

//   var d JsonTestDate
//   if !useDecoder {
//     //Dump response body (use json.Unmarshal)
//     fmt.Printf("======Body (use json.Unmarshal)======\n")
//     body, err := ioutil.ReadAll(res.Body)
//     if err != nil {
//       log.Fatal(err)
//     }

//     err = json.Unmarshal(body, &d)
//     if err != nil {
//       log.Fatal(err)
//     }
//   } else {
//     //Dump response body (use Decoder)
//     fmt.Printf("======Body(use json.Decoder)======\n")
//     decoder := json.NewDecoder(res.Body)
//     err = decoder.Decode(&d)
//     if err != nil {
//       log.Fatal(err)
//     }
//   }
//   fmt.Printf("%v\n", d)
// }
