package base

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var HttpClient *HttpCli

const (
	post    = "POST"
	deletes = "DELETE"
	get     = "GET"
	put     = "put"
)

type HttpCli struct {
	Client  *http.Client
	BaseURL *url.URL
	Ctx     context.Context
}

type HttpResponse struct {
	Body []byte
	Code int
}

func InitHttpClient() error {
	client, err := New("")
	if err != nil {
		return err
	}
	HttpClient = client
	return nil
}

func New(baseURL string, timeout ...int) (*HttpCli, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	to := 10
	if len(timeout) > 0 && timeout[0] > 0 {
		to = timeout[0]
	}
	hc := &HttpCli{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives: true,
			},
			Timeout: time.Second * time.Duration(to),
		},
		BaseURL: base,
		Ctx:     context.TODO(),
	}
	return hc, nil
}

func (h *HttpCli) Post(url string, body []byte) (*HttpResponse, error) {
	return h.PostJSONWithContext(url, body, h.Ctx)
}

func (h *HttpCli) Get(url string, hd ...map[interface{}]interface{}) (*HttpResponse, error) {
	if hd != nil && len(hd) > 0 {
		return h.GetWithContext(url, h.Ctx, hd[0])
	}
	return h.GetWithContext(url, h.Ctx)
}

func (h *HttpCli) GetWithContext(url string, ctx context.Context, hd ...map[interface{}]interface{}) (*HttpResponse, error) {
	ul, err := h.parse(url)
	if err != nil {
		return nil, err
	}
	requ, err := http.NewRequestWithContext(ctx, get, ul.String(), nil)
	if err != nil {
		return nil, err
	}
	if hd != nil && len(hd) > 0 {
		for k, v := range hd[0] {
			requ.Header.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
		}
	}
	return h.do(requ)
}

func (h *HttpCli) PostJSONWithContext(url string, body []byte, ctx context.Context) (*HttpResponse, error) {
	ul, err := h.parse(url)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if body != nil && len(body) > 0 {
		buf = bytes.NewBuffer(body)
	}
	requ, err := http.NewRequestWithContext(ctx, post, ul.String(), buf)
	if err != nil {
		return nil, err
	}
	requ.Header.Set("Content-Type", "application/json")
	return h.do(requ)
}

func (h *HttpCli) PostWithHeader(url string, header map[interface{}]interface{}) (*HttpResponse, error) {
	ul, err := h.parse(url)
	if err != nil {
		return nil, err
	}
	requ, err := http.NewRequestWithContext(h.Ctx, post, ul.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		requ.Header.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}
	return h.do(requ)
}

func (h *HttpCli) DeleteWithHeader(url string, header map[interface{}]interface{}) (*HttpResponse, error) {
	ul, err := h.parse(url)
	if err != nil {
		return nil, err
	}
	requ, err := http.NewRequestWithContext(h.Ctx, deletes, ul.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		requ.Header.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}
	return h.do(requ)
}

func (h *HttpCli) do(requ *http.Request) (*HttpResponse, error) {
	fmt.Println("url : ", requ.URL.String())
	retres, err := h.Client.Do(requ)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(retres.Body)
	if err != nil {
		return nil, err
	}
	defer retres.Body.Close()
	return &HttpResponse{
		Body: all,
		Code: retres.StatusCode,
	}, nil
}

func (h *HttpCli) parse(url string) (*url.URL, error) {
	return h.BaseURL.Parse(url)
}
