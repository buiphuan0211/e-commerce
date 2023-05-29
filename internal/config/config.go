package config

import (
	"ecommerce/internal/constant"
	"os"
)

type ENV struct {
	Env string

	ZookeeperPrefixEcommerceCommon string
	ZookeeperPrefixEcommerceAdmin  string

	Admin struct {
		Server string
		Port   string
	}

	MongoDB struct {
		URI    string
		DBName string

		ReplicaSet          string
		CAPem               string
		CertPem             string
		CertKeyFilePassword string
		ReadPrefMode        string
	}

	Auth struct {
		SecretKey string
	}
}

var env ENV

func GetENV() *ENV {
	return &env
}

func IsEnvDevelop() bool {
	return env.Env == constant.EnvDevelop
}
func IsEnvProduction() bool {
	return env.Env == constant.EnvProduction
}

func Init() {
	env = ENV{
		Env:                            os.Getenv("ENV"),
		ZookeeperPrefixEcommerceCommon: os.Getenv("ZOOKEEPER_PREFIX_ECOMMERCE_COMMON"),
		ZookeeperPrefixEcommerceAdmin:  os.Getenv("ZOOKEEPER_PREFIX_ECOMMERCE_ADMIN"),
	}
}
