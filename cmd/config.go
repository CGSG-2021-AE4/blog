package main

import (
	"flag"
	"time"
)

type Config struct {
	Domain string `json:"domain"`

	CertFilename    string `json:"certFilename"`
	PrivKeyFilename string `json:"privKeyFilename"`

	UserStoreFilename   string `json:"userStoreFilename"`
	UserSvcSecret       string `json:"usersSecret"`
	UserTokenExpTimeout int64  `json:"userTokenExpTimeout"`

	ArticleStoreFilename string `json:"articleStoreFilename"`
	ArticleContentDir    string `json:"ArticleContentDir"`
}

func NewConfigFromFlags() Config {
	var config Config

	flag.StringVar(&config.Domain, "domain", "localhost:8080", "Domain address")

	flag.StringVar(&config.CertFilename, "cert-file", "out/cert.pem", "Certification filename")
	flag.StringVar(&config.PrivKeyFilename, "key-file", "out/key.pem", "Private key filename")

	flag.StringVar(&config.UserStoreFilename, "users-file", "out/user_store.json", "User store json database filename")
	flag.StringVar(&config.UserSvcSecret, "secret", "cgsgforever", "User service hmac secret for tokens signing")
	flag.Int64Var(&config.UserTokenExpTimeout, "token-exp-timeout", int64(3*time.Hour), "User token expiration timeout")

	flag.StringVar(&config.ArticleStoreFilename, "articles-file", "out/article_store.json", "Article store json database filename")
	flag.StringVar(&config.ArticleContentDir, "articles-content-dir", "out/db/articles/", "Article content files' dir")

	flag.Parse()

	return config
}
