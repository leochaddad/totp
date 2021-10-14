package main
import (
	"fmt"
	"io"
	"crypto/sha1"
	"crypto/hmac"

    "encoding/hex"
	"time"

)

func main(){

	fmt.Println(h("JBSWY3DPEHPK3PXP", "122"))
	counter()
}

func hotp(key string ) string{
	h := sha1.New()
	io.WriteString(h, key)
	sum:=h.Sum(nil)[:]


	sEnc := hex.EncodeToString(sum)

	return sEnc

}

func h(secretToken string, payloadBody string) string{
	mac := hmac.New(sha1.New, []byte(secretToken))
	mac.Write([]byte(payloadBody))
	expectedMAC := mac.Sum(nil)
	return "sha1=" + expectedMAC
}

func counter() {
	t0 := time.Now().Unix()
	var tx int64 = 30000
  
    fmt.Printf("%v\n", t0/tx)
}