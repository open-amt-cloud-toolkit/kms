package config

type Configuration struct {
	ID            int              `yaml:"id"`
	Name          string           `yaml:"name"`
	Configuration RemoteManagement `yaml:"configuration"`
}

type RemoteManagement struct {
	GeneralSettings     GeneralSettings     `yaml:"generalSettings"`
	Network             Network             `yaml:"network"`
	Authentication      Authentication      `yaml:"authentication"`
	TLS                 TLS                 `yaml:"tls"`
	Redirection         Redirection         `yaml:"redirection"`
	UserAccounts        UserAccounts        `yaml:"userAccounts"`
	EnterpriseAssistant EnterpriseAssistant `yaml:"enterpriseAssistant"`
	AMTSpecific         AMTSpecific         `yaml:"amtSpecific"`
	BMCSpecific         BMCSpecific         `yaml:"bmcSpecific"`
	DASHSpecific        DASHSpecific        `yaml:"dashSpecific"`
	RedfishSpecific     RedfishSpecific     `yaml:"redfishSpecific"`
}

type GeneralSettings struct {
	SharedFQDN              bool `yaml:"sharedFQDN"`
	NetworkInterfaceEnabled int  `yaml:"networkInterfaceEnabled"`
	PingResponseEnabled     bool `yaml:"pingResponseEnabled"`
}

type Network struct {
	Wired    Wired    `yaml:"wired"`
	Wireless Wireless `yaml:"wireless"`
}

type Wired struct {
	DHCPEnabled    bool   `yaml:"dhcpEnabled"`
	IPSyncEnabled  bool   `yaml:"ipSyncEnabled"`
	SharedStaticIP bool   `yaml:"sharedStaticIP"`
	IPAddress      string `yaml:"ipAddress"`
	SubnetMask     string `yaml:"subnetMask"`
	DefaultGateway string `yaml:"defaultGateway"`
	PrimaryDNS     string `yaml:"primaryDNS"`
	SecondaryDNS   string `yaml:"secondaryDNS"`
	Authentication string `yaml:"authentication"`
}

type Wireless struct {
	Profiles []string `yaml:"profiles"`
}

type Authentication struct {
	Profiles []string `yaml:"profiles"`
}

type TLS struct {
	MutualAuthentication bool     `yaml:"mutualAuthentication"`
	Enabled              bool     `yaml:"enabled"`
	TrustedCN            []string `yaml:"trustedCN"`
}

type Redirection struct {
	Enabled     bool     `yaml:"enabled"`
	Services    Services `yaml:"services"`
	UserConsent string   `yaml:"userConsent"`
}

type Services struct {
	KVM  bool `yaml:"kvm"`
	SOL  bool `yaml:"sol"`
	IDER bool `yaml:"ider"`
}

type UserAccounts struct {
	UserAccounts []string `yaml:"userAccounts"`
}

type EnterpriseAssistant struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AMTSpecific struct {
	ControlMode         string `yaml:"controlMode"`
	AdminPassword       string `yaml:"adminPassword"`
	ProvisioningCert    string `yaml:"provisioningCert"`
	ProvisioningCertPwd string `yaml:"provisioningCertPwd"`
	MEBXPassword        string `yaml:"mebxPassword"`
}

type BMCSpecific struct {
	AdminPassword string `yaml:"adminPassword"`
}

type DASHSpecific struct {
	AdminPassword string `yaml:"adminPassword"`
}

type RedfishSpecific struct {
	AdminPassword string `yaml:"adminPassword"`
}
