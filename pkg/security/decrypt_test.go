package security

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"io/fs"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/config"
)

var (
	validKey           = "ThisismyveryStrongkey32byteslong"
	wrongKey           = "ThisismyveryStrongkey32bytelong!"
	shortKey           = "shortKey"
	missingKey         = ""
	validMessageText   = "Hello, World!"
	invalidMessageText = "Invalid_Base64_String!@#"
	expectedConfigFile = config.Configuration{
		ID:   1234,
		Name: "Test",
		Configuration: config.RemoteManagement{
			GeneralSettings: config.GeneralSettings{
				SharedFQDN:              true,
				NetworkInterfaceEnabled: 1,
				PingResponseEnabled:     true,
			},
			Network: config.Network{
				Wired: config.Wired{
					DHCPEnabled:    true,
					IPSyncEnabled:  true,
					SharedStaticIP: true,
					IPAddress:      "",
					SubnetMask:     "",
					DefaultGateway: "",
					PrimaryDNS:     "",
					SecondaryDNS:   "",
					Authentication: "",
				},
				Wireless: config.Wireless{
					Profiles: []string{},
				},
			},
			Authentication: config.Authentication{
				Profiles: []string{},
			},
			TLS: config.TLS{
				MutualAuthentication: false,
				Enabled:              false,
				TrustedCN:            []string{},
			},
			Redirection: config.Redirection{
				Enabled: true,
				Services: config.Services{
					KVM:  true,
					SOL:  true,
					IDER: true,
				},
				UserConsent: "none",
			},
			UserAccounts: config.UserAccounts{
				UserAccounts: []string{},
			},
			EnterpriseAssistant: config.EnterpriseAssistant{},
			AMTSpecific: config.AMTSpecific{
				ControlMode:         "ccmactivate",
				AdminPassword:       "adminPassword",
				ProvisioningCert:    "",
				ProvisioningCertPwd: "",
				MEBXPassword:        "",
			},
			BMCSpecific:     config.BMCSpecific{},
			DASHSpecific:    config.DASHSpecific{},
			RedfishSpecific: config.RedfishSpecific{},
		},
	}
)

type expectedError struct {
	Base64Error         bool
	NewCipherError      bool
	AuthenticationError bool
	InvalidKeySizeError bool
	FileReadError       bool
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		message       string
		key           string
		expectedError expectedError
		errorMsg      error
		expected      []byte
	}{
		{
			name:          "successful decryption",
			message:       validMessageText,
			key:           validKey,
			expectedError: expectedError{},
			errorMsg:      nil,
			expected:      []byte("Hello World"),
		},
		{
			name:          "fail to decode base64",
			message:       invalidMessageText,
			key:           validKey,
			expectedError: expectedError{Base64Error: true},
			errorMsg:      base64.CorruptInputError(7),
			expected:      nil,
		},
		{
			name:          "fail to create new cipher",
			message:       validMessageText,
			key:           missingKey,
			expectedError: expectedError{NewCipherError: true},
			errorMsg:      aes.KeySizeError(0),
			expected:      nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var err error

			var decryptedString []byte

			cryptor := Crypto{}

			if tc.expectedError.Base64Error {
				_, err = cryptor.Decrypt(tc.message, []byte(tc.key))
				assert.Equal(t, tc.errorMsg, err)
				assert.Equal(t, tc.expected, decryptedString)
			}

			if tc.expectedError.NewCipherError {
				encryptedString, _ := cryptor.Encrypt([]byte(tc.message), tc.key)
				decryptedString, err = cryptor.Decrypt(encryptedString, []byte(tc.key))
				assert.Equal(t, tc.errorMsg, err)
				assert.Equal(t, tc.expected, decryptedString)
			}

			if !tc.expectedError.Base64Error && !tc.expectedError.NewCipherError {
				encryptedString, _ := cryptor.Encrypt([]byte(tc.message), tc.key)
				decryptedString, err = cryptor.Decrypt(encryptedString, []byte(tc.key))
				assert.Equal(t, tc.message, string(decryptedString))
				assert.NoError(t, err)
			}
		})
	}
}

func TestReadAndDecryptFile(t *testing.T) {
	t.Parallel()

	byteArrayConfigFile, _ := yaml.Marshal(expectedConfigFile)
	tests := []struct {
		name          string
		filePath      string
		key           string
		expectedError expectedError
		errorMsg      error
		expected      []byte
	}{
		{
			name:          "successful decryption",
			filePath:      "testing/encryptedConfig.yaml",
			key:           validKey,
			expectedError: expectedError{},
			errorMsg:      nil,
			expected:      byteArrayConfigFile,
		},
		{
			name:          "incorrect key size",
			filePath:      "testing/encryptedConfig.yaml",
			key:           shortKey,
			expectedError: expectedError{InvalidKeySizeError: true},
			errorMsg:      aes.KeySizeError(8),
			expected:      []byte("Hello World"),
		},
		{
			name:          "incorrect key",
			filePath:      "testing/encryptedConfig.yaml",
			key:           wrongKey,
			expectedError: expectedError{AuthenticationError: true},
			errorMsg:      errors.New("cipher: message authentication failed"),
			expected:      []byte("Hello World"),
		},
		{
			name:          "unable to read file",
			filePath:      "testing/doesnotexist.yaml",
			key:           validKey,
			expectedError: expectedError{FileReadError: true},
			errorMsg:      &fs.PathError{Op: "open", Path: "testing/doesnotexist.yaml", Err: syscall.ENOENT},
			expected:      []byte(""),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cryptor := Crypto{}
			_, err := cryptor.ReadAndDecryptFile(test.filePath, []byte(test.key))

			if !test.expectedError.InvalidKeySizeError && !test.expectedError.AuthenticationError && !test.expectedError.NewCipherError && !test.expectedError.Base64Error && !test.expectedError.FileReadError {
				decryptedFile, err := cryptor.ReadAndDecryptFile(test.filePath, []byte(test.key))
				assert.Equal(t, expectedConfigFile, decryptedFile)
				assert.NoError(t, err)
			} else {
				assert.Equal(t, test.errorMsg, err)
			}
		})
	}
}
