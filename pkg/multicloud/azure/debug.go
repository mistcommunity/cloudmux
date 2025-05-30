
package azure

import (
	"net/http"
	"net/http/httputil"

	"github.com/Azure/go-autorest/autorest"

	"yunion.io/x/log"
)

func LogRequest() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				log.Errorln(err)
			}
			auth := r.Header.Get("Authorization")
			if len(auth) > 0 {
				log.Debugf("Authorization: %s", auth)
			}
			return r, err
		})
	}
}

func LogResponse() autorest.RespondDecorator {
	return func(p autorest.Responder) autorest.Responder {
		return autorest.ResponderFunc(func(r *http.Response) error {
			err := p.Respond(r)
			if err != nil {
				log.Errorln(err)
			}
			dump, _ := httputil.DumpResponse(r, true)
			log.Errorf("%s", string(dump))
			return err
		})
	}
}
