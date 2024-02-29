package pem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/akimsavvin/savvin_go_lib_common/logger/sl"
)

type KeysParser struct {
	log sl.Log
	Err error
}

func NewKeysParser(log sl.Log) *KeysParser {
	return &KeysParser{
		log: log.With(sl.Pkg("pem"), sl.Mdl("KeysParses")),
	}
}

func (p *KeysParser) ParsePublicKey(pubPEM string) *rsa.PublicKey {
	const op = "ParsePublicKey"
	log := p.log.With(sl.Op(op))

	if p.Err != nil {
		log.Warn("parsing did not start due to existing error")

		return nil
	}

	log.Debug("key parsing started")

	block, _ := pem.Decode([]byte(pubPEM))

	if block == nil {
		log.Error("could not decode public key PEM block")

		p.Err = errors.New("failed to parse PEM block containing the key")

		return nil
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Error("could not parse public key", sl.Err(err))

		p.Err = err

		return nil
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		log.Debug("key parsed")
		return pub
	default:
		break
	}

	log.Error("could not parse key since key type is not RSA")

	p.Err = errors.New("key type is not RSA")

	return nil
}

func (p *KeysParser) ParsePrivateKey(prvPEM string) *rsa.PrivateKey {
	const op = "ParsePrivateKey"
	log := p.log.With(sl.Op(op))

	if p.Err != nil {
		log.Warn("parsing did not start due to existing error")

		return nil
	}

	log.Debug("key parsing started")

	block, _ := pem.Decode([]byte(prvPEM))

	if block == nil {
		log.Error("could not decode public key PEM block")

		p.Err = errors.New("failed to parse PEM block containing the key")

		return nil
	}

	prv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Error("could not parse private key", sl.Err(err))

		p.Err = err

		return nil
	}

	switch prv := prv.(type) {
	case *rsa.PrivateKey:
		log.Debug("key parsed")
		return prv
	default:
		break
	}

	log.Error("could not parse key since key type is not RSA")

	p.Err = errors.New("key type is not RSA")

	return nil
}
