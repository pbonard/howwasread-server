package common

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
)

func CreateTlSConfig(certFile, keyFile, caCertFile string) (tlsConfig *tls.Config, err error) {
	tlsConfig = &tls.Config{InsecureSkipVerify: false}
	if certFile != "" {
		clientCert, err1 := tls.LoadX509KeyPair(certFile, keyFile)
		if err1 != nil {
			return nil, fmt.Errorf("failed to load client cert/key: %w", err1)
		}
		tlsConfig.Certificates = []tls.Certificate{clientCert}
	}
	if caCertFile != "" {
		caCert, err2 := os.ReadFile(caCertFile)
		if err2 != nil {
			return nil, fmt.Errorf("failed to read CA cert: %w", err2)
		}
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("failed to append CA cert to pool")
		}
		tlsConfig.RootCAs = caCertPool
	}
	return tlsConfig, nil
}

func AddJitter(base int64) int64 {
	if base <= 0 {
		return base
	}
	maxJitter := float64(base) * 0.10
	multiplier := (rand.Float64() * 2.0) - 1.0
	jitterAmount := maxJitter * multiplier
	return base + int64(jitterAmount)
}

func Marshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		slog.Error("fail to marshal", "err", err)
		return nil
	}
	return b
}
