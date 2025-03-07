package encryption_test

import (
	"testing"

	"github.com/Duong2903/client-encryption-go/encryption"
	"github.com/Duong2903/client-encryption-go/jwe"
	"github.com/Duong2903/client-encryption-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestEncryptPayload_ShouldEncryptRootArrays(t *testing.T) {
	payload := "[" +
		"   {}," +
		"   {}" +
		"]"

	decryptionKeyPath := "../testdata/keys/pkcs8/test_key_pkcs8-2048.der"
	certificatePath := "../testdata/certificates/test_certificate-2048.der"

	decryptionKey, err := utils.LoadUnencryptedDecryptionKey(decryptionKeyPath)
	assert.Nil(t, err)
	certificate, err := utils.LoadEncryptionCertificate(certificatePath)
	assert.Nil(t, err)

	cb := jwe.NewJWEConfigBuilder()
	jweConfig := cb.WithDecryptionKey(decryptionKey).
		WithEncryptionPath("$", "$").
		WithDecryptionPath("encryptedData", "$").
		WithEncryptedValueFieldName("encryptedData").
		WithCertificate(certificate).
		Build()

	encryptedPayload := encryption.EncryptPayload(payload, *jweConfig)
	decryptedPayload := encryption.DecryptPayload(encryptedPayload, *jweConfig)

	assert.Equal(t, "[{},{}]", decryptedPayload)
}

func TestDecryptPayload_ShouldEncryptRootArrays(t *testing.T) {
	const encryptedPayload = "{" +
		"    \"encryptedData\": \"eyJraWQiOiI3NjFiMDAzYzFlYWRlM2E1NDkwZTUwMDBkMzc4ODdiYWE1ZTZlYzBlMjI2YzA3NzA2ZTU5OTQ1MWZjMDMyYTc5IiwiY3R5IjoiYXBwbGljYXRpb24vanNvbiIsImVuYyI6IkEyNTZHQ00iLCJhbGciOiJSU0EtT0FFUC0yNTYifQ.IcTIce59pgtjODJn4PhR7oK3F-gxcd7dishTrT7T9y5VC0U5ZS_JdMoRe59_UTkJMY8Nykb2rv3Oh_jSDYRmGB_CWMIciXYMLHQptLTF5xI1ZauDPnooDMWoOCBD_d3I0wTJNcM7I658rK0ZWSByVK9YqhEo8UaIf4e6egRHQdZ2_IGKgICwmglv_uXQrYewOWFTKR1uMpya1N50MDnWax2NtnW3SljP3mARUBLBnRmOyubQCg-Mgn8fsOWWXm-KL9RrQq9AF_HJceoJl1rRgzPW7g6SLK6EjiGW_ArTmrLaOHg9bYOY_LrbyokK_M1pMo9qup70DHvjHkMZqIL3aQ.vtma3jBIo2STkquxTUX9PQ.9ZoQG0sFvQ.ms4bW3OFd03neRlex-zZ8w\"" +
		"}"

	decryptionKeyPath := "../testdata/keys/pkcs8/test_key_pkcs8-2048.der"
	certificatePath := "../testdata/certificates/test_certificate-2048.der"

	decryptionKey, err := utils.LoadUnencryptedDecryptionKey(decryptionKeyPath)
	assert.Nil(t, err)
	certificate, err := utils.LoadEncryptionCertificate(certificatePath)
	assert.Nil(t, err)

	cb := jwe.NewJWEConfigBuilder()
	jweConfig := cb.WithDecryptionKey(decryptionKey).
		WithCertificate(certificate).
		WithEncryptedValueFieldName("encryptedData").
		WithDecryptionPath("encryptedData", "$").
		Build()

	decryptedPayload := encryption.DecryptPayload(encryptedPayload, *jweConfig)
	assert.Equal(t, "[{},{}]", decryptedPayload)
}
