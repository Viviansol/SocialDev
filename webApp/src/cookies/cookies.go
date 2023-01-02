package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"webApp/src/config"
)

var s *securecookie.SecureCookie

func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encryptedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encryptedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
