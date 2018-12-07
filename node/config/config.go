package config

import (
	"github.com/caarlos0/env"
	"log"
)

const (
	PlasmaRangeSpace = 2 ^ 24 - 1
)

type OperatorConfig struct {
	OperatorPort          int    `env:"operator_port" envDefault:"3001"`
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x240d6ad83930067d82e0803696996f743acd78d8fa6a5f6e4f148fd9def37c55"`
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0x28088c91385a83563c6392ee09f1d0d386f0cd7f"`
	GethAccount           string `env:"geth_account" envDefault:"ws://127.0.0.1:8546"`
}

type VerifierConfig struct {
	VerifierPort          int    `env:"verifier_port" envDefault:"2000"`                                                                          // port where verifier server runs
	MainAccountPrivateKey string `env:"main_account_private_key" envDefault:"0x19980f942c4a65850f8b5f730695cc3b3ff666fc028ff2addc9756676b235476"` // private key of account who deploy plasma contract and who push blocks to it (operator)
	MainAccountPublicKey  string `env:"main_account_public_key" envDefault:"0xb2571928F73a6Ecd86c17b60863e6F9cF1Cf2084"`                          // public key of account who deploy plasma contract and who push blocks to it (operator)
	PlasmaContractAddress string `env:"plasma_contract_address" envDefault:"0xa86a2c6B81C22d25D8EBAf8cE52895F46A263348"`                          // address of plasma smart contract
	GethHost              string `env:"geth_host" envDefault:"ws://127.0.0.1:8546"`
	OperatorHost          string `env:"operator_host" envDefault:"http://127.0.0.1:3001"`
}

var (
	operatorConfigInstance *OperatorConfig
	verifierConfigInstance *VerifierConfig
)

// GetOperator gets operator config instance.
func GetOperator() *OperatorConfig {
	if operatorConfigInstance == nil {
		operatorConfigInstance = &OperatorConfig{}
		err := env.Parse(operatorConfigInstance)
		if err != nil {
			log.Fatalf("error initializing config: %s", err)
		}
	}
	return operatorConfigInstance
}

// GetVerifier gets verifier config instance.
func GetVerifier() *VerifierConfig {
	if verifierConfigInstance == nil {
		verifierConfigInstance = &VerifierConfig{}
		err := env.Parse(verifierConfigInstance)
		if err != nil {
			log.Fatalf("error initializing config: %s", err)
		}
	}
	return verifierConfigInstance
}
