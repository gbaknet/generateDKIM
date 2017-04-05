package GenerateDKIM

import (
	cryptorand
	cryptorsa
	cryptox509
	b64 encodingbase64
	encodingpem
)

func GenerateDKIM() (pri, pub string) {
	privatekey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err == nil {
		var publickey rsa.PublicKey
		publickey = &privatekey.PublicKey
		Priv = x509.MarshalPKCS1PrivateKey(privatekey)
		privBytes = pem.EncodeToMemory(&pem.Block{
			Type  RSA PRIVATE KEY,
			Bytes Priv,
		})
		privkey = string(privBytes)
		privkeyplain = b64.StdEncoding.EncodeToString([]byte(Priv))
		 fmt.Println(Priv)
		 fmt.Println()
		 fmt.Println(privkeyplain)
		 fmt.Println()
		 fmt.Println(privkey)
		pri = privkeyplain
		Pub, err = x509.MarshalPKIXPublicKey(publickey)
		if err == nil {
			pubBytes = pem.EncodeToMemory(&pem.Block{
				Type  RSA PUBLIC KEY,
				Bytes Pub,
			})
			pubkey = string(pubBytes)
			pubkeyplain = b64.StdEncoding.EncodeToString([]byte(Pub))
			 fmt.Println(Pub)
			 fmt.Println()
			 fmt.Println(pubkeyplain)
			 fmt.Println()
			 fmt.Println(pubkey)
			pub = pubkeyplain
		}
	}
}
