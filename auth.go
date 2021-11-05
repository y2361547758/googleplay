package googleplay

import (
   "bytes"
   "github.com/segmentio/encoding/proto"
   "io"
   "net/http"
   "net/url"
)

type Auth struct {
   url.Values
}

// deviceID is Google Service Framework.
func (a Auth) Details(deviceID, app string) ([]byte, error) {
   req, err := http.NewRequest("GET", origin + "/fdfe/details", nil)
   if err != nil {
      return nil, err
   }
   val := req.URL.Query()
   val.Set("doc", app)
   req.URL.RawQuery = val.Encode()
   req.Header.Set("Authorization", "Bearer " + a.Get("Auth"))
   req.Header.Set("X-DFE-Device-ID", deviceID)
   res, err := roundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}

// This seems to return `StatusOK`, even with invalid requests, and the response
// body only contains a token, that doesnt seem to indicate success or failure.
// Only way I know to check, it to try the `deviceID` with a `details` request
// or similar. Also, after the POST, you need to wait at least 16 seconds
// before the `deviceID` can be used.
func (a Auth) Upload(deviceID string, dev Device) error {
   buf, err := proto.Marshal(dev)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", origin + "/fdfe/uploadDeviceConfig", bytes.NewReader(buf),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.Get("Auth")},
      "User-Agent": {"Android-Finsky (sdk=99,versionCode=99999999)"},
      "X-DFE-Device-ID": {deviceID},
   }
   res, err := roundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}
