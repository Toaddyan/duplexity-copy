package oauth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/duplexityio/duplexity/pkg/messages"
	"golang.org/x/oauth2"
)
var (
	clientID     = messages.ClientID
	clientSecret = messages.ClientSecret
)

// NOTE: userID is our clientID from outside the application. 


type User struct {
	UserID string,
	OAuth2Token *oauth2.Token,
	IDTokenClaims *json.RawMessage,
}

type AuthConfig struct {
	config       oauth2.Config,
	verifier     *oidc.IDTokenVerifier,
	provider     *oidc.Provider, 
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (cfg *AuthConfig)loginUser(w http.ResponseWriter, req *http.Request) {
	state, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	nonce, err := randString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	setCallbackCookie(w, r, "state", state)
	setCallbackCookie(w, r, "nonce", nonce)

	http.Redirect(w, r, config.AuthCodeURL(state, oidc.Nonce(nonce)), http.StatusFound)
}

func authenticate(userID string, ctx context.Context) {
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)
	config := oauth2.Config{
		ClientID:     messages.ClientID
		ClientSecret: messages.ClientSecret
		Endpoint:     provider.Endpoint(),
		RedirectURL:  messages.CallbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	
	http.HandleFunc("/", )
}