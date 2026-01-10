package tokenmanager

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type TokenManager interface {
	Generate(claims map[string]interface{}) (string, error)
	Verify(token string) (map[string]interface{}, error)
}

type JWEManager struct {
	signPrivateKey *rsa.PrivateKey
	signPublicKey  *rsa.PublicKey
	encryptionKey  []byte
}

// NewJWEManager takes both RSA private & public keys + AES-256 encryption key.
func NewJWEManager(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, encryptionKey []byte) (TokenManager, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("RSA private and public keys are required")
	}
	if len(encryptionKey) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes (got %d)", len(encryptionKey))
	}

	return &JWEManager{
		signPrivateKey: privateKey,
		signPublicKey:  publicKey,
		encryptionKey:  encryptionKey,
	}, nil
}

// Generate creates a signed JWS and then encrypts it as a JWE.
func (m *JWEManager) Generate(claims map[string]interface{}) (string, error) {
	if len(claims) == 0 {
		return "", errors.New("claims cannot be empty")
	}

	builder := jwt.NewBuilder()
	for k, v := range claims {
		builder = builder.Claim(k, v)
	}

	token, err := builder.Build()
	if err != nil {
		return "", fmt.Errorf("failed to build JWT: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.EdDSA, m.signPrivateKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	encrypted, err := jwe.Encrypt(
		signed,
		jwe.WithContentEncryption(jwa.A256GCM),
		jwe.WithKey(jwa.DIRECT, m.encryptionKey),
	)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt JWE: %w", err)
	}

	return string(encrypted), nil
}

// Verify decrypts the JWE, verifies the JWS signature, and returns the claims.
func (m *JWEManager) Verify(token string) (map[string]interface{}, error) {
	if token == "" {
		return nil, errors.New("token cannot be empty")
	}

	decrypted, err := jwe.Decrypt(
		[]byte(token),
		jwe.WithKey(jwa.DIRECT, m.encryptionKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt JWE: %w", err)
	}

	parsed, err := jwt.Parse(
		decrypted,
		jwt.WithKey(jwa.EdDSA, m.signPrivateKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to verify JWS signature: %w", err)
	}

	claims := make(map[string]interface{})
	for k, v := range parsed.PrivateClaims() {
		claims[k] = v
	}
	if sub := parsed.Subject(); sub != "" {
		claims["sub"] = sub
	}
	if aud := parsed.Audience(); len(aud) > 0 {
		claims["aud"] = aud
	}
	if iss := parsed.Issuer(); iss != "" {
		claims["iss"] = iss
	}

	return claims, nil
}

// ParseRSAKeys parses PEM strings into *rsa.PrivateKey and *rsa.PublicKey
func ParseRSAKeys(privateKeyPEM string, publicKeyPEM string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, nil, fmt.Errorf("failed to parse RSA private key PEM")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse RSA private key: %w", err)
	}

	// Parse public key
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	if blockPub == nil {
		return nil, nil, fmt.Errorf("failed to parse RSA public key PEM")
	}
	pubKeyInterface, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse RSA public key: %w", err)
	}
	pubKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("parsed key is not RSA public key")
	}

	return privKey, pubKey, nil
}

// DecodeAESKey decodes the base64-encoded AES-256 key
func DecodeAESKey(aeskeybase64 string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(aeskeybase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode AES key: %w", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("AES key must be 32 bytes")
	}
	return key, nil
}

func ParseEdDSAKeyFromPEM(privateKeyPEM string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return nil, nil, fmt.Errorf("failed to parse EdDSA private key PEM")
	}
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse EdDSA private key: %w", err)
	}
	edPriv, ok := privKey.(ed25519.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("parsed key is not Ed25519 private key")
	}
	edPub := edPriv.Public().(ed25519.PublicKey)
	return edPriv, edPub, nil
}
