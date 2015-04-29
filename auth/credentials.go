package auth

type Credentials interface {
	GetAccessKeyId() string
	GetSecretAccessKey() string
}

type CredentialsProvider interface {
	SetCredentials(creds Credentials)
	GetCredentials() (creds Credentials)
}

//default Credentials
func DefaultCredentials(accessKeyId, secretAccessKey string) Credentials {
	if accessKeyId == "" || secretAccessKey == "" {
		panic("accessKeyId and secretAccessKey can't be nil")
	}
	return &credentials{accessKeyId: accessKeyId, secretAccessKey: secretAccessKey}
}

type credentials struct {
	accessKeyId     string
	secretAccessKey string
}

func (p *credentials) GetAccessKeyId() string {
	return p.accessKeyId
}

func (p *credentials) GetSecretAccessKey() string {
	return p.secretAccessKey
}

//default CredentialsProvider
func DefaultCredentialsProvider(accessKeyId, secretAccessKey string) CredentialsProvider {
	return &credentialsProvider{creds: DefaultCredentials(accessKeyId, secretAccessKey)}
}

type credentialsProvider struct {
	creds Credentials
}

func (p *credentialsProvider) SetCredentials(creds Credentials) {
	p.creds = creds
}

func (p *credentialsProvider) GetCredentials() (creds Credentials) {
	return p.creds
}
