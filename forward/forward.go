package forward

import (
	"crypto/tls"
	"net/http"
	"io"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/sirupsen/logrus"
)


type Handler struct {
	TlsConfig *tls.Config

}

func (h Handler) ServeHTTP(res http.ResponseWriter,req *http.Request) {
	//TODO forward http request and sent it back
	var host  string = req.Host
	var path string = req.URL.Path
	var method string = req.Method

	logrus.Info("%s\t%s\t%s", method, path, host)

	roundTripper := http3.RoundTripper{
		TLSClientConfig: h.TlsConfig,
		QuicConfig: &quic.Config{},
	}	
	client := http.Client{Transport: &roundTripper}

	new_res, err := client.Do(req)

	if err != nil {
		(res).WriteHeader(305)
		return 
	}

	defer new_res.Body.Close()

	bodyBytes, readErr := io.ReadAll(new_res.Body)

	if readErr != nil {
		(res).WriteHeader(305)
		return 
	}

	_, err = res.Write(bodyBytes)

	if err != nil {
		(res).WriteHeader(305)
		return 
	}



}


