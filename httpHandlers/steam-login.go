package httpHandlers
// Steam login hnadler
// based on https://github.com/solovev/steam_go

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"fmt"
	"os"
	"encoding/json"
	"github.com/deflexor/gonewsticker/structs"
	"github.com/deflexor/gonewsticker/session"
)

type OpenId struct {
	root      string
	returnUrl string
	data      url.Values
}

var (
	steam_login = "https://steamcommunity.com/openid/login"

	openId_mode       = "checkid_setup"
	openId_ns         = "http://specs.openid.net/auth/2.0"
	openId_identifier = "http://specs.openid.net/auth/2.0/identifier_select"

	validation_regexp        = regexp.MustCompile("^(http|https)://steamcommunity.com/openid/id/[0-9]{15,25}$")
	digits_extraction_regexp = regexp.MustCompile("\\D+")
)

func SteamLogin(w http.ResponseWriter, r *http.Request) {
	opId := NewOpenId(r)
	switch opId.Mode() {
	case "":
		http.Redirect(w, r, opId.AuthUrl(), 301)
		break
	case "cancel":
		w.Write([]byte("Authorization cancelled"))
		break
	default:
		apiKey, ok := os.LookupEnv("STEAM_KEY")
		if !ok {
			http.Error(w, "STEAM_KEY env variable is not set, cannot access Steam api!", http.StatusInternalServerError)
			return
		}
		user, err := opId.ValidateAndGetUser(apiKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// log.Printf("got steam user: %v", user)
		guid := session.Login(*user)
		http.SetCookie(w, &http.Cookie{Name: "SESSIONV2", Value: guid, Path: "/", HttpOnly: true})
		w.Write([]byte("<html><body><script>window.close();</script></body></html>"))
	}
}

func NewOpenId(r *http.Request) *OpenId {
	id := new(OpenId)

	proto := "http://"
	if r.TLS != nil {
		proto = "https://"
	}
	id.root = proto + r.Host

	uri := r.RequestURI
	if i := strings.Index(uri, "openid"); i != -1 {
		uri = uri[0 : i-1]
	}
	id.returnUrl = id.root + uri

	switch r.Method {
	case "POST":
		id.data = r.Form
	case "GET":
		id.data = r.URL.Query()
	}

	return id
}

func (id OpenId) AuthUrl() string {
	data := map[string]string{
		"openid.claimed_id": openId_identifier,
		"openid.identity":   openId_identifier,
		"openid.mode":       openId_mode,
		"openid.ns":         openId_ns,
		"openid.realm":      id.root,
		"openid.return_to":  id.returnUrl,
	}

	i := 0
	url := steam_login + "?"
	for key, value := range data {
		url += key + "=" + value
		if i != len(data)-1 {
			url += "&"
		}
		i++
	}
	return url
}

func (id *OpenId) ValidateAndGetId() (string, error) {
	if id.Mode() != "id_res" {
		return "", errors.New("Mode must equal to \"id_res\".")
	}

	if id.data.Get("openid.return_to") != id.returnUrl {
		return "", errors.New("The \"return_to url\" must match the url of current request.")
	}

	params := make(url.Values)
	params.Set("openid.assoc_handle", id.data.Get("openid.assoc_handle"))
	params.Set("openid.signed", id.data.Get("openid.signed"))
	params.Set("openid.sig", id.data.Get("openid.sig"))
	params.Set("openid.ns", id.data.Get("openid.ns"))

	split := strings.Split(id.data.Get("openid.signed"), ",")
	for _, item := range split {
		params.Set("openid."+item, id.data.Get("openid."+item))
	}
	params.Set("openid.mode", "check_authentication")

	resp, err := http.PostForm(steam_login, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := strings.Split(string(content), "\n")
	if response[0] != "ns:"+openId_ns {
		return "", errors.New("Wrong ns in the response.")
	}
	if strings.HasSuffix(response[1], "false") {
		return "", errors.New("Unable validate openId.")
	}

	openIdUrl := id.data.Get("openid.claimed_id")
	if !validation_regexp.MatchString(openIdUrl) {
		return "", errors.New("Invalid steam id pattern.")
	}

	return digits_extraction_regexp.ReplaceAllString(openIdUrl, ""), nil
}

func (id OpenId) ValidateAndGetUser(apiKey string) (*structs.PlayerSummary, error) {
	steamId, err := id.ValidateAndGetId()
	if err != nil {
		return nil, err
	}
	return GetPlayerSummaries(steamId, apiKey)
}

func (id OpenId) Mode() string {
	return id.data.Get("openid.mode")
}

func GetPlayerSummaries(steamId, apiKey string) (*structs.PlayerSummary, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", apiKey, steamId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type Result struct {
		Response struct {
			Players []structs.PlayerSummary `json:"players"`
		} `json:"response"`
	}
	var data Result
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data.Response.Players[0], err
}