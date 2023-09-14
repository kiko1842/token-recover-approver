package wireset

import (
	"errors"

	"github.com/bnb-chain/airdrop-service/internal/config"
	"github.com/bnb-chain/airdrop-service/pkg/keymanager"
	"github.com/bnb-chain/airdrop-service/pkg/keymanager/aws"
	"github.com/bnb-chain/airdrop-service/pkg/keymanager/local"
)

type SecretType string

const (
	LocalKey         SecretType = "local"
	AWSSecretManager SecretType = "aws"
)

func InitKeyManager(config *config.Config) (keymanager.KeyManager, error) {
	switch SecretType(config.Secret.Type) {
	case LocalKey:
		return local.NewLocalKeyManager(config.Secret.LocalSecretConfig.PrivateKey)
	case AWSSecretManager:
		return aws.NewSecretManager(config.Secret.AWSSecretManagerConfig.SecretName, config.Secret.AWSSecretManagerConfig.Region)
	default:
		return nil, errors.New("invalid secret type")
	}
}
