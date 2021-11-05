package googleplay

import (
   "bytes"
   "fmt"
   "os"
   "testing"
)

func TestAuth(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   // read token
   tok := new(Token)
   {
      r, err := os.Open(cache + "/googleplay/token.json")
      if err != nil {
         t.Fatal(err)
      }
      defer r.Close()
      if err := tok.Decode(r); err != nil {
         t.Fatal(err)
      }
   }
   o, err := tok.Auth()
   if err != nil {
      t.Fatal(err)
   }
   // read checkin
   check := new(Checkin)
   {
      r, err := os.Open(cache + "/googleplay/checkin.json")
      if err != nil {
         t.Fatal(err)
      }
      defer r.Close()
      if err := check.Decode(r); err != nil {
         t.Fatal(err)
      }
   }
   // details
   det, err := o.Details(check.String(), "com.google.android.youtube")
   if err != nil {
      t.Fatal(err)
   }
   vers := []string{"16.", "16.4", "16.43.", "16.43.3", "16.43.34"}
   for _, ver := range vers {
      if bytes.Contains(det, []byte(ver)) {
         fmt.Println("pass", ver)
      } else {
         fmt.Println("fail", ver)
      }
   }
}
