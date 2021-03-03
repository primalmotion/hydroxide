package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func TLS(certPath string, keyPath string, clientCAPath string) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	var didSomething bool

	if certPath != "" && keyPath != "" {
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			return nil, fmt.Errorf("error: unable load key pair: %s", err)
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
		didSomething = true
	}

	if clientCAPath != "" {
		data, err := ioutil.ReadFile(clientCAPath)
		if err != nil {
			return nil, fmt.Errorf("error: unable read CA file: %s", err)
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(data)

		tlsConfig.ClientCAs = pool
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		didSomething = true
	}

	if !didSomething {
		return nil, nil
	}

	return tlsConfig, nil
}
