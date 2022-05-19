package googleplay

import (
   "os"
   "testing"
   "time"
)

func TestCheckin(t *testing.T) {
   cache, err := os.UserCacheDir()
   if err != nil {
      t.Fatal(err)
   }
   for _, platform := range Platforms {
      device, err := Phone.Checkin(platform)
      if err != nil {
         t.Fatal(err)
      }
      platform += ".json"
      if err := device.Create(cache, "googleplay", platform); err != nil {
         t.Fatal(err)
      }
      time.Sleep(Sleep)
   }
}
