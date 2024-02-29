package jwt

import (
	"context"
	"crypto/rsa"
	"errors"

	"github.com/akimsavvin/savvin_go_lib_common/logger/sl"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/json"
)

type JWT = *jose.JSONWebSignature

type Claim = any

type Claims = map[string]Claim

type Manager struct {
	log    sl.Log
	pubKey *rsa.PublicKey
}

func NewManager(log sl.Log, pubKey *rsa.PublicKey) *Manager {
	log.Debug("creating jwt manager", sl.Pkg("jwt"), sl.Op("NewManager"))

	return &Manager{
		log:    log.With(sl.Pkg("jwt"), sl.Mdl("Manager")),
		pubKey: pubKey,
	}
}

var ErrInvalidToken = errors.New("invalid token")

func (m *Manager) ParseToken(ctx context.Context, tokenStr string) (JWT, error) {
	const op = "ParseToken"
	log := m.log.With(sl.Op(op), sl.CorId(ctx), sl.ReqId(ctx))

	token, err := jose.ParseSigned(tokenStr)
	if err != nil {
		log.Error("could not parse token", sl.Err(err))
		return nil, err
	}

	return token, nil
}

func (m *Manager) GetTokenClaims(ctx context.Context, token JWT) (Claims, error) {
	const op = "GetTokenClaims"
	log := m.log.With(sl.Op(op), sl.CorId(ctx), sl.ReqId(ctx))

	payload, err := token.Verify(m.pubKey)
	if err != nil {
		log.Error("could not verify token", sl.Err(err))
		return nil, ErrInvalidToken
	}

	var claims map[string]interface{}

	if err := json.Unmarshal(payload, &claims); err != nil {
		log.Error("could not unmarshal token payload", sl.Err(err))
		return nil, err
	}

	return claims, nil
}
