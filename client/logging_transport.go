package client

import (
	"io"
	"log"
	"net/http"
)

type loggingTransport struct {
	base http.RoundTripper
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("[TANGLE-CLIENT] [INFO] Disparando requisição HTTP %s %s", req.Method, req.URL.String())

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		log.Printf(
			"[TANGLE-CLIENT] [ERROR] Falha na requisição HTTP: method=%s url=%s error=%v",
			req.Method,
			req.URL.String(),
			err,
		)
		return nil, err
	}

	log.Printf(
		"[TANGLE-CLIENT] [INFO] Resposta HTTP recebida: method=%s url=%s status=%d",
		req.Method,
		req.URL.String(),
		resp.StatusCode,
	)

	resp.Body = &loggingReadCloser{
		ReadCloser: resp.Body,
		method:     req.Method,
		url:        req.URL.String(),
		statusCode: resp.StatusCode,
	}

	return resp, nil
}

type loggingReadCloser struct {
	io.ReadCloser
	method     string
	url        string
	statusCode int
	bytesRead  int64
	logged     bool
}

func (l *loggingReadCloser) Read(p []byte) (int, error) {
	n, err := l.ReadCloser.Read(p)
	l.bytesRead += int64(n)

	if err == io.EOF && !l.logged {
		l.logged = true
		log.Printf(
			"[TANGLE-CLIENT] [INFO] Payload lido: method=%s url=%s status=%d payload_bytes=%d",
			l.method,
			l.url,
			l.statusCode,
			l.bytesRead,
		)
	}

	return n, err
}

func (l *loggingReadCloser) Close() error {
	err := l.ReadCloser.Close()

	if err != nil {
		log.Printf(
			"[TANGLE-CLIENT] [ERROR] Falha ao fechar response body: method=%s url=%s error=%v",
			l.method,
			l.url,
			err,
		)
	}

	if !l.logged {
		l.logged = true
		log.Printf(
			"[TANGLE-CLIENT] [INFO] Payload lido (body fechado antes do EOF): method=%s url=%s status=%d payload_bytes=%d",
			l.method,
			l.url,
			l.statusCode,
			l.bytesRead,
		)
	}

	return err
}
