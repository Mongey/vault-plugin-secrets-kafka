// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/identity/mfa/types.proto

package mfa

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Config represents the configuration information used *along with* the MFA
// secret tied to caller's identity, to verify the MFA credentials supplied.
// Configuration information differs by type. Handler of each type should know
// what to expect from the Config field.
type Config struct {
	Type           string `sentinel:"" protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name           string `sentinel:"" protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ID             string `sentinel:"" protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	MountAccessor  string `sentinel:"" protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor,proto3" json:"mount_accessor,omitempty"`
	UsernameFormat string `sentinel:"" protobuf:"bytes,5,opt,name=username_format,json=usernameFormat,proto3" json:"username_format,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*Config_TOTPConfig
	//	*Config_OktaConfig
	//	*Config_DuoConfig
	//	*Config_PingIDConfig
	Config               isConfig_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Config) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Config) GetUsernameFormat() string {
	if m != nil {
		return m.UsernameFormat
	}
	return ""
}

type isConfig_Config interface {
	isConfig_Config()
}

type Config_TOTPConfig struct {
	TOTPConfig *TOTPConfig `sentinel:"" protobuf:"bytes,6,opt,name=totp_config,json=totpConfig,proto3,oneof"`
}

type Config_OktaConfig struct {
	OktaConfig *OktaConfig `sentinel:"" protobuf:"bytes,7,opt,name=okta_config,json=oktaConfig,proto3,oneof"`
}

type Config_DuoConfig struct {
	DuoConfig *DuoConfig `sentinel:"" protobuf:"bytes,8,opt,name=duo_config,json=duoConfig,proto3,oneof"`
}

type Config_PingIDConfig struct {
	PingIDConfig *PingIDConfig `sentinel:"" protobuf:"bytes,9,opt,name=pingid_config,json=pingidConfig,proto3,oneof"`
}

func (*Config_TOTPConfig) isConfig_Config() {}

func (*Config_OktaConfig) isConfig_Config() {}

func (*Config_DuoConfig) isConfig_Config() {}

func (*Config_PingIDConfig) isConfig_Config() {}

func (m *Config) GetConfig() isConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Config) GetTOTPConfig() *TOTPConfig {
	if x, ok := m.GetConfig().(*Config_TOTPConfig); ok {
		return x.TOTPConfig
	}
	return nil
}

func (m *Config) GetOktaConfig() *OktaConfig {
	if x, ok := m.GetConfig().(*Config_OktaConfig); ok {
		return x.OktaConfig
	}
	return nil
}

func (m *Config) GetDuoConfig() *DuoConfig {
	if x, ok := m.GetConfig().(*Config_DuoConfig); ok {
		return x.DuoConfig
	}
	return nil
}

func (m *Config) GetPingIDConfig() *PingIDConfig {
	if x, ok := m.GetConfig().(*Config_PingIDConfig); ok {
		return x.PingIDConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Config) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Config_TOTPConfig)(nil),
		(*Config_OktaConfig)(nil),
		(*Config_DuoConfig)(nil),
		(*Config_PingIDConfig)(nil),
	}
}

// TOTPConfig represents the configuration information required to generate
// a TOTP key. The generated key will be stored in the entity along with these
// options. Validation of credentials supplied over the API will be validated
// by the information stored in the entity and not from the values in the
// configuration.
type TOTPConfig struct {
	Issuer               string   `sentinel:"" protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Period               uint32   `sentinel:"" protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	Algorithm            int32    `sentinel:"" protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Digits               int32    `sentinel:"" protobuf:"varint,4,opt,name=digits,proto3" json:"digits,omitempty"`
	Skew                 uint32   `sentinel:"" protobuf:"varint,5,opt,name=skew,proto3" json:"skew,omitempty"`
	KeySize              uint32   `sentinel:"" protobuf:"varint,6,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	QRSize               int32    `sentinel:"" protobuf:"varint,7,opt,name=qr_size,json=qrSize,proto3" json:"qr_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TOTPConfig) Reset()         { *m = TOTPConfig{} }
func (m *TOTPConfig) String() string { return proto.CompactTextString(m) }
func (*TOTPConfig) ProtoMessage()    {}
func (*TOTPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{1}
}

func (m *TOTPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TOTPConfig.Unmarshal(m, b)
}
func (m *TOTPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TOTPConfig.Marshal(b, m, deterministic)
}
func (m *TOTPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TOTPConfig.Merge(m, src)
}
func (m *TOTPConfig) XXX_Size() int {
	return xxx_messageInfo_TOTPConfig.Size(m)
}
func (m *TOTPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TOTPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TOTPConfig proto.InternalMessageInfo

func (m *TOTPConfig) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *TOTPConfig) GetPeriod() uint32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *TOTPConfig) GetAlgorithm() int32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *TOTPConfig) GetDigits() int32 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *TOTPConfig) GetSkew() uint32 {
	if m != nil {
		return m.Skew
	}
	return 0
}

func (m *TOTPConfig) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *TOTPConfig) GetQRSize() int32 {
	if m != nil {
		return m.QRSize
	}
	return 0
}

// DuoConfig represents the configuration information required to perform
// Duo authentication.
type DuoConfig struct {
	IntegrationKey       string   `sentinel:"" protobuf:"bytes,1,opt,name=integration_key,json=integrationKey,proto3" json:"integration_key,omitempty"`
	SecretKey            string   `sentinel:"" protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	APIHostname          string   `sentinel:"" protobuf:"bytes,3,opt,name=api_hostname,json=apiHostname,proto3" json:"api_hostname,omitempty"`
	PushInfo             string   `sentinel:"" protobuf:"bytes,4,opt,name=push_info,json=pushInfo,proto3" json:"push_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DuoConfig) Reset()         { *m = DuoConfig{} }
func (m *DuoConfig) String() string { return proto.CompactTextString(m) }
func (*DuoConfig) ProtoMessage()    {}
func (*DuoConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{2}
}

func (m *DuoConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DuoConfig.Unmarshal(m, b)
}
func (m *DuoConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DuoConfig.Marshal(b, m, deterministic)
}
func (m *DuoConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DuoConfig.Merge(m, src)
}
func (m *DuoConfig) XXX_Size() int {
	return xxx_messageInfo_DuoConfig.Size(m)
}
func (m *DuoConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DuoConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DuoConfig proto.InternalMessageInfo

func (m *DuoConfig) GetIntegrationKey() string {
	if m != nil {
		return m.IntegrationKey
	}
	return ""
}

func (m *DuoConfig) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *DuoConfig) GetAPIHostname() string {
	if m != nil {
		return m.APIHostname
	}
	return ""
}

func (m *DuoConfig) GetPushInfo() string {
	if m != nil {
		return m.PushInfo
	}
	return ""
}

// OktaConfig contains Okta configuration parameters required to perform Okta
// authentication.
type OktaConfig struct {
	OrgName              string   `sentinel:"" protobuf:"bytes,1,opt,name=org_name,json=orgName,proto3" json:"org_name,omitempty"`
	APIToken             string   `sentinel:"" protobuf:"bytes,2,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	Production           bool     `sentinel:"" protobuf:"varint,3,opt,name=production,proto3" json:"production,omitempty"`
	BaseURL              string   `sentinel:"" protobuf:"bytes,4,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	PrimaryEmail         bool     `sentinel:"" protobuf:"varint,5,opt,name=primary_email,json=primaryEmail,proto3" json:"primary_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OktaConfig) Reset()         { *m = OktaConfig{} }
func (m *OktaConfig) String() string { return proto.CompactTextString(m) }
func (*OktaConfig) ProtoMessage()    {}
func (*OktaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{3}
}

func (m *OktaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OktaConfig.Unmarshal(m, b)
}
func (m *OktaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OktaConfig.Marshal(b, m, deterministic)
}
func (m *OktaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OktaConfig.Merge(m, src)
}
func (m *OktaConfig) XXX_Size() int {
	return xxx_messageInfo_OktaConfig.Size(m)
}
func (m *OktaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OktaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OktaConfig proto.InternalMessageInfo

func (m *OktaConfig) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *OktaConfig) GetAPIToken() string {
	if m != nil {
		return m.APIToken
	}
	return ""
}

func (m *OktaConfig) GetProduction() bool {
	if m != nil {
		return m.Production
	}
	return false
}

func (m *OktaConfig) GetBaseURL() string {
	if m != nil {
		return m.BaseURL
	}
	return ""
}

func (m *OktaConfig) GetPrimaryEmail() bool {
	if m != nil {
		return m.PrimaryEmail
	}
	return false
}

// PingIDConfig contains PingID configuration information
type PingIDConfig struct {
	UseBase64Key         string   `sentinel:"" protobuf:"bytes,1,opt,name=use_base64_key,json=useBase64Key,proto3" json:"use_base64_key,omitempty"`
	UseSignature         bool     `sentinel:"" protobuf:"varint,2,opt,name=use_signature,json=useSignature,proto3" json:"use_signature,omitempty"`
	Token                string   `sentinel:"" protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	IDPURL               string   `sentinel:"" protobuf:"bytes,4,opt,name=idp_url,json=idpUrl,proto3" json:"idp_url,omitempty"`
	OrgAlias             string   `sentinel:"" protobuf:"bytes,5,opt,name=org_alias,json=orgAlias,proto3" json:"org_alias,omitempty"`
	AdminURL             string   `sentinel:"" protobuf:"bytes,6,opt,name=admin_url,json=adminUrl,proto3" json:"admin_url,omitempty"`
	AuthenticatorURL     string   `sentinel:"" protobuf:"bytes,7,opt,name=authenticator_url,json=authenticatorUrl,proto3" json:"authenticator_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingIDConfig) Reset()         { *m = PingIDConfig{} }
func (m *PingIDConfig) String() string { return proto.CompactTextString(m) }
func (*PingIDConfig) ProtoMessage()    {}
func (*PingIDConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{4}
}

func (m *PingIDConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingIDConfig.Unmarshal(m, b)
}
func (m *PingIDConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingIDConfig.Marshal(b, m, deterministic)
}
func (m *PingIDConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingIDConfig.Merge(m, src)
}
func (m *PingIDConfig) XXX_Size() int {
	return xxx_messageInfo_PingIDConfig.Size(m)
}
func (m *PingIDConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PingIDConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PingIDConfig proto.InternalMessageInfo

func (m *PingIDConfig) GetUseBase64Key() string {
	if m != nil {
		return m.UseBase64Key
	}
	return ""
}

func (m *PingIDConfig) GetUseSignature() bool {
	if m != nil {
		return m.UseSignature
	}
	return false
}

func (m *PingIDConfig) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PingIDConfig) GetIDPURL() string {
	if m != nil {
		return m.IDPURL
	}
	return ""
}

func (m *PingIDConfig) GetOrgAlias() string {
	if m != nil {
		return m.OrgAlias
	}
	return ""
}

func (m *PingIDConfig) GetAdminURL() string {
	if m != nil {
		return m.AdminURL
	}
	return ""
}

func (m *PingIDConfig) GetAuthenticatorURL() string {
	if m != nil {
		return m.AuthenticatorURL
	}
	return ""
}

// Secret represents all the types of secrets which the entity can hold.
// Each MFA type should add a secret type to the oneof block in this message.
type Secret struct {
	MethodName string `sentinel:"" protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Secret_TOTPSecret
	Value                isSecret_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{5}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

type isSecret_Value interface {
	isSecret_Value()
}

type Secret_TOTPSecret struct {
	TOTPSecret *TOTPSecret `sentinel:"" protobuf:"bytes,2,opt,name=totp_secret,json=totpSecret,proto3,oneof"`
}

func (*Secret_TOTPSecret) isSecret_Value() {}

func (m *Secret) GetValue() isSecret_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Secret) GetTOTPSecret() *TOTPSecret {
	if x, ok := m.GetValue().(*Secret_TOTPSecret); ok {
		return x.TOTPSecret
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Secret) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Secret_TOTPSecret)(nil),
	}
}

// TOTPSecret represents the secret that gets stored in the entity about a
// particular MFA method. This information is used to validate the MFA
// credential supplied over the API during request time.
type TOTPSecret struct {
	Issuer    string `sentinel:"" protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Period    uint32 `sentinel:"" protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	Algorithm int32  `sentinel:"" protobuf:"varint,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Digits    int32  `sentinel:"" protobuf:"varint,4,opt,name=digits,proto3" json:"digits,omitempty"`
	Skew      uint32 `sentinel:"" protobuf:"varint,5,opt,name=skew,proto3" json:"skew,omitempty"`
	KeySize   uint32 `sentinel:"" protobuf:"varint,6,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	// reserving 7 here just to keep parity with the config message above
	AccountName          string   `sentinel:"" protobuf:"bytes,8,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Key                  string   `sentinel:"" protobuf:"bytes,9,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TOTPSecret) Reset()         { *m = TOTPSecret{} }
func (m *TOTPSecret) String() string { return proto.CompactTextString(m) }
func (*TOTPSecret) ProtoMessage()    {}
func (*TOTPSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb73493aac0ba29, []int{6}
}

func (m *TOTPSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TOTPSecret.Unmarshal(m, b)
}
func (m *TOTPSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TOTPSecret.Marshal(b, m, deterministic)
}
func (m *TOTPSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TOTPSecret.Merge(m, src)
}
func (m *TOTPSecret) XXX_Size() int {
	return xxx_messageInfo_TOTPSecret.Size(m)
}
func (m *TOTPSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TOTPSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TOTPSecret proto.InternalMessageInfo

func (m *TOTPSecret) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *TOTPSecret) GetPeriod() uint32 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *TOTPSecret) GetAlgorithm() int32 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *TOTPSecret) GetDigits() int32 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *TOTPSecret) GetSkew() uint32 {
	if m != nil {
		return m.Skew
	}
	return 0
}

func (m *TOTPSecret) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *TOTPSecret) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *TOTPSecret) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "mfa.Config")
	proto.RegisterType((*TOTPConfig)(nil), "mfa.TOTPConfig")
	proto.RegisterType((*DuoConfig)(nil), "mfa.DuoConfig")
	proto.RegisterType((*OktaConfig)(nil), "mfa.OktaConfig")
	proto.RegisterType((*PingIDConfig)(nil), "mfa.PingIDConfig")
	proto.RegisterType((*Secret)(nil), "mfa.Secret")
	proto.RegisterType((*TOTPSecret)(nil), "mfa.TOTPSecret")
}

func init() { proto.RegisterFile("helper/identity/mfa/types.proto", fileDescriptor_2eb73493aac0ba29) }

var fileDescriptor_2eb73493aac0ba29 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xdd, 0xae, 0xdc, 0x34,
	0x10, 0x66, 0xf7, 0x74, 0xf3, 0x33, 0xfb, 0xd3, 0xd6, 0x42, 0xb0, 0xa8, 0x40, 0xcb, 0x16, 0x44,
	0x25, 0xa4, 0x5d, 0x74, 0x40, 0x88, 0xdb, 0x2e, 0x05, 0xb5, 0x42, 0xa2, 0x55, 0xce, 0xe9, 0x0d,
	0x37, 0x91, 0x4f, 0xe2, 0x24, 0xd6, 0x26, 0x71, 0x6a, 0x3b, 0x45, 0x7b, 0x5e, 0x83, 0x57, 0xe0,
	0x29, 0x78, 0x10, 0xde, 0x02, 0xf1, 0x0a, 0x68, 0xc6, 0xce, 0x6e, 0x90, 0x78, 0x80, 0xde, 0x79,
	0xbe, 0x99, 0x6f, 0x3c, 0xfe, 0x66, 0xc6, 0xf0, 0xb0, 0x12, 0x75, 0x27, 0xf4, 0x4e, 0xe6, 0xa2,
	0xb5, 0xd2, 0x1e, 0x77, 0x4d, 0xc1, 0x77, 0xf6, 0xd8, 0x09, 0xb3, 0xed, 0xb4, 0xb2, 0x8a, 0x5d,
	0x34, 0x05, 0xdf, 0xfc, 0x3d, 0x85, 0xe0, 0x07, 0xd5, 0x16, 0xb2, 0x64, 0x0c, 0xee, 0xa0, 0x7b,
	0x3d, 0x79, 0x34, 0x79, 0x12, 0x27, 0x74, 0x46, 0xac, 0xe5, 0x8d, 0x58, 0x4f, 0x1d, 0x86, 0x67,
	0xb6, 0x82, 0xa9, 0xcc, 0xd7, 0x17, 0x84, 0x4c, 0x65, 0xce, 0xbe, 0x80, 0x55, 0xa3, 0xfa, 0xd6,
	0xa6, 0x3c, 0xcb, 0x84, 0x31, 0x4a, 0xaf, 0xef, 0x90, 0x6f, 0x49, 0xe8, 0x53, 0x0f, 0xb2, 0x2f,
	0xe1, 0x6e, 0x6f, 0x84, 0xc6, 0x14, 0x69, 0xa1, 0x74, 0xc3, 0xed, 0x7a, 0x46, 0x71, 0xab, 0x01,
	0xfe, 0x89, 0x50, 0x76, 0x09, 0x73, 0xab, 0x6c, 0x97, 0x66, 0x54, 0xd6, 0x3a, 0x78, 0x34, 0x79,
	0x32, 0xbf, 0xbc, 0xbb, 0x6d, 0x0a, 0xbe, 0xbd, 0x7e, 0x79, 0xfd, 0xca, 0x55, 0xfb, 0xfc, 0xbd,
	0x04, 0x30, 0xca, 0xd7, 0x7e, 0x09, 0x73, 0x75, 0xb0, 0x7c, 0xe0, 0x84, 0x23, 0xce, 0xcb, 0x83,
	0xe5, 0x67, 0x8e, 0x3a, 0x59, 0x6c, 0x07, 0x90, 0xf7, 0x6a, 0xa0, 0x44, 0x44, 0x59, 0x11, 0xe5,
	0x59, 0xaf, 0x4e, 0x8c, 0x38, 0x1f, 0x0c, 0xf6, 0x3d, 0x2c, 0x3b, 0xd9, 0x96, 0x32, 0x1f, 0x38,
	0x31, 0x71, 0xee, 0x13, 0xe7, 0x95, 0x6c, 0xcb, 0x17, 0xcf, 0x4e, 0xb4, 0x85, 0x8b, 0x74, 0xf6,
	0x3e, 0x82, 0xc0, 0x51, 0x36, 0x7f, 0x4e, 0x00, 0xce, 0xaf, 0x60, 0x1f, 0x40, 0x20, 0x8d, 0xe9,
	0x85, 0xf6, 0xaa, 0x7b, 0x0b, 0xf1, 0x4e, 0x68, 0xa9, 0x72, 0x52, 0x7e, 0x99, 0x78, 0x8b, 0x7d,
	0x0c, 0x31, 0xaf, 0x4b, 0xa5, 0xa5, 0xad, 0x1a, 0x6a, 0xc1, 0x2c, 0x39, 0x03, 0xc8, 0xca, 0x65,
	0x29, 0xad, 0xa1, 0x0e, 0xcc, 0x12, 0x6f, 0x61, 0x17, 0xcd, 0x41, 0xfc, 0x46, 0x7a, 0x2f, 0x13,
	0x3a, 0xb3, 0x8f, 0x20, 0x3a, 0x88, 0x63, 0x6a, 0xe4, 0xad, 0x20, 0x89, 0x97, 0x49, 0x78, 0x10,
	0xc7, 0x2b, 0x79, 0x2b, 0xd8, 0x87, 0x10, 0xbe, 0xd1, 0xce, 0x13, 0xba, 0x3c, 0x6f, 0x34, 0x3a,
	0x36, 0xbf, 0x4f, 0x20, 0x3e, 0x69, 0x83, 0x0d, 0x95, 0xad, 0x15, 0xa5, 0xe6, 0x56, 0xaa, 0x36,
	0x3d, 0x88, 0xa3, 0x7f, 0xc4, 0x6a, 0x04, 0xff, 0x2c, 0x8e, 0xec, 0x13, 0x00, 0x23, 0x32, 0x2d,
	0x2c, 0xc5, 0xb8, 0x51, 0x8a, 0x1d, 0x82, 0xee, 0xcf, 0x60, 0xc1, 0x3b, 0x99, 0x56, 0xca, 0x58,
	0x9a, 0x35, 0x37, 0x59, 0x73, 0xde, 0xc9, 0xe7, 0x1e, 0x62, 0x0f, 0x20, 0xee, 0x7a, 0x53, 0xa5,
	0xb2, 0x2d, 0x94, 0x9f, 0xae, 0x08, 0x81, 0x17, 0x6d, 0xa1, 0x36, 0x7f, 0x4c, 0x00, 0xce, 0x4d,
	0xc6, 0x87, 0x29, 0x5d, 0xa6, 0x94, 0xca, 0xd5, 0x13, 0x2a, 0x5d, 0xfe, 0xe2, 0xd3, 0xe0, 0x4d,
	0x56, 0x1d, 0x44, 0xeb, 0xeb, 0x88, 0x78, 0x27, 0xaf, 0xd1, 0x66, 0x9f, 0x02, 0x74, 0x5a, 0xe5,
	0x7d, 0x86, 0x65, 0x53, 0x11, 0x51, 0x32, 0x42, 0x30, 0xef, 0x0d, 0x37, 0x22, 0xed, 0x75, 0xed,
	0x4b, 0x08, 0xd1, 0x7e, 0xad, 0x6b, 0xf6, 0x18, 0x96, 0x9d, 0x96, 0x0d, 0xd7, 0xc7, 0x54, 0x34,
	0x5c, 0xd6, 0x24, 0x74, 0x94, 0x2c, 0x3c, 0xf8, 0x23, 0x62, 0x9b, 0x7f, 0x26, 0xb0, 0x18, 0x0f,
	0x09, 0xfb, 0x1c, 0x70, 0xf2, 0x53, 0x4c, 0xf2, 0xdd, 0xb7, 0x23, 0xf9, 0x16, 0xbd, 0x11, 0x7b,
	0x02, 0x51, 0x9d, 0xc7, 0xb0, 0xc4, 0x28, 0x23, 0xcb, 0x96, 0xdb, 0x5e, 0xbb, 0x55, 0x8c, 0x28,
	0xe8, 0x6a, 0xc0, 0xd8, 0xfb, 0x30, 0x73, 0x8f, 0x72, 0xda, 0x39, 0x03, 0xfb, 0x28, 0xf3, 0x6e,
	0x54, 0x70, 0x20, 0xf3, 0x0e, 0xeb, 0x7d, 0x00, 0x31, 0x4a, 0xc4, 0x6b, 0xc9, 0x8d, 0x5f, 0x42,
	0xd4, 0xec, 0x29, 0xda, 0x24, 0x52, 0xde, 0xc8, 0x96, 0x78, 0x81, 0x17, 0x09, 0x01, 0x64, 0x7e,
	0x05, 0xf7, 0x79, 0x6f, 0x2b, 0xfc, 0x51, 0x32, 0x6e, 0x95, 0xa6, 0xa0, 0x90, 0x82, 0xee, 0xfd,
	0xc7, 0xf1, 0x5a, 0xd7, 0x9b, 0x02, 0x82, 0x2b, 0xea, 0x32, 0x7b, 0x08, 0xf3, 0x46, 0xd8, 0x4a,
	0xe5, 0xe3, 0xb6, 0x80, 0x83, 0xa8, 0x33, 0xc3, 0xce, 0xbb, 0xa9, 0xa0, 0x37, 0x8e, 0x77, 0xde,
	0xa5, 0x19, 0x76, 0xde, 0x59, 0xfb, 0x10, 0x66, 0x6f, 0x79, 0xdd, 0x8b, 0xcd, 0x5f, 0x7e, 0xa7,
	0xfc, 0x65, 0xef, 0xe4, 0x4e, 0xe1, 0x90, 0x67, 0x19, 0x7d, 0x93, 0x24, 0x41, 0xe4, 0x87, 0xdc,
	0x61, 0xa4, 0xc1, 0x3d, 0xb8, 0xc0, 0x21, 0x88, 0xc9, 0x83, 0xc7, 0xfd, 0xd7, 0xbf, 0x6e, 0x4b,
	0x69, 0xab, 0xfe, 0x66, 0x9b, 0xa9, 0x66, 0x57, 0x71, 0x53, 0xc9, 0x4c, 0xe9, 0x6e, 0xf7, 0x96,
	0xf7, 0xb5, 0xdd, 0xfd, 0xcf, 0xff, 0x7e, 0x13, 0xd0, 0xd7, 0xfe, 0xcd, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0x93, 0x21, 0xaa, 0xfd, 0x05, 0x00, 0x00,
}
