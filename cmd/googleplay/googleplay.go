package main

import (
   "fmt"
   "github.com/89z/googleplay"
   "os"
   "time"
)

func checkin() error {
   tok := new(googleplay.Token)
   cache, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   r, err := os.Open(cache + "/googleplay/token.json")
   if err != nil {
      return err
   }
   defer r.Close()
   if err := tok.Decode(r); err != nil {
      return err
   }
   o, err := tok.Auth()
   if err != nil {
      return err
   }
   c, err := googleplay.NewCheckinRequest().Post()
   if err != nil {
      return err
   }
   if err := o.Upload(c.String(), googleplay.NewDevice()); err != nil {
      return err
   }
   w, err := os.Create(cache + "/googleplay/checkin.json")
   if err != nil {
      return err
   }
   defer w.Close()
   if err := c.Encode(w); err != nil {
      return err
   }
   time.Sleep(googleplay.Sleep)
   return nil
}

func details() (*googleplay.AppDetails, error) {
   cache, err := os.UserCacheDir()
   if err != nil {
      return nil, err
   }
   tok := new(googleplay.Token)
   {
      r, err := os.Open(cache + "/googleplay/token.json")
      if err != nil {
         return nil, err
      }
      defer r.Close()
      if err := tok.Decode(r); err != nil {
         return nil, err
      }
   }
   o, err := tok.Auth()
   if err != nil {
      return nil, err
   }
   check := new(googleplay.Checkin)
   {
      r, err := os.Open(cache + "/googleplay/checkin.json")
      if err != nil {
         return nil, err
      }
      defer r.Close()
      if err := check.Decode(r); err != nil {
         return nil, err
      }
   }
   return o.Details(check.String(), "com.google.android.youtube")
}

func token(email, password string) error {
   tok, err := googleplay.NewToken(email, password)
   if err != nil {
      return err
   }
   c, err := os.UserCacheDir()
   if err != nil {
      return err
   }
   c += "/googleplay"
   os.Mkdir(c, os.ModeDir)
   f, err := os.Create(c + "/token.json")
   if err != nil {
      return err
   }
   defer f.Close()
   return tok.Encode(f)
}

func main() {
   fmt.Println("hello world")
}
