resource "google_compute_ssl_certificate" "gcp_compute_ssl_certificates_cert" {
  name_prefix = "gcp-compute-ssl-cert-${var.test_suffix}"
  description = "a description"
  private_key = data.template_file.gcp_compute_ssl_certificates_private_key.rendered
  certificate = data.template_file.gcp_compute_ssl_certificates_certificate.rendered

  lifecycle {
    create_before_destroy = true
  }
}


data "template_file" "gcp_compute_ssl_certificates_private_key" {
  template = <<EOF
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZqZI7lNSN2xUP
azPoIgRWZJy8OYkcbSDFA9xb+Hs5jUMgs9Jf1zZ8OhUTAQL12Qofpt7XM9f2R4ac
dZcy1YHYVcsFBww+q/yESBlhMeso8hPRgjqA/J8nZ0VW/IVbMJv3dU5/SPK1ITuP
BLCbsWFO5yuXxDpQOPYrA1N6UayvE/rtZgMqeWSAOy1XWgrQPRgYbm9NkXGX92VN
7BycOdWRbu4ezGr7ZSFgNJ5VB0/n14RHS3DMsf7BMkwpg733U4J6TQ7C4Ms3FSVh
qm+kPiwQrMfQNDSxeLo1Hw8AdxzjVlouWnSm1vaFTkIBdopdLctcEnnjBktTMGlv
G8qmlXsfAgMBAAECggEAO+5p2kfvgqOpF9a/sxHyucr4MQdyjkYp+LVIbnZrj3wq
2I1KxqLeWLQxa0sjAohhNjffMcgPlbs6AEiMei25k9SDkv3OzE7Ut6OWgWGaS2rk
NBK0gyGLvPC9cecT3Pj0aN1+4KM4WNEusgFrk2Ly1SPnp+Ea4U3d0hgXWx2z+3x1
pJcbND8E3yq9JQIwuZNyx60n/gI22M4FaNbtD7pGapJgVz3ApV/4DEHB65pqqVIt
kZdRmF7XNDjuQtsuW6i6E7UUMDra1NHuZGvdfmXR3W8btWeoHSnhPZDIYLIDW3Po
9nE3fTCOQQlkvUgZXFCJQEJDs2UiqzV9W9jYpRRr0QKBgQD7EUNQ/PPBudeG9C8s
r3vo8iPtnqauFvwftUp24bA7vYJVffoQYExuYBQ5p3CmqpsruSAeChrznZiVw8BE
UlP3aEnsjSpWsVIh04w4dCmWHAY+yrSblVDu3gBgoUy2WO8K6OxRCO0oGm2uC4+b
BKDy38eLuWOOqrQ1/Pb5TG6xmQKBgQDd8ExocbqQoKEwC4f3B23GcxW47JV9OO3K
mVmmp9zq6+cbb342wQ0wEVZ8H6jQ/oSzLwGOCWsVtRHL6FxOWUvmbUmhDJlQkE5P
hS09lRJMR74j1ZIE8jDswguLRFU6yDYh6UGfP6fnDS3B0Rb586QpID3hC3LXcBcu
5r3qDWZ1dwKBgB9aMIXUkLwIcRmxNJLn9xlH46Swwy/KPwHWqc3esRtEtxnl+WxC
GklORjhM6IxnkakMHS6jJGp3q65IG6JshX/HzjN0DW12B0OiH0iNeQP9y+nbdmJX
axvpLTLj8ahzwqYiICCedL8lTb0GRJCfK1opB8ozBHO0bXywckb/fHNBAoGBALI+
PX9cZ3OkLhBCEo6I/tb0sqt0BpMtV3zxMBkyk7BwiYl1P66F2SuToRvK6XAAGV83
D06drc0fQQ28rfWWreiAOTQIxFD5tIsU8EKXKLzumXx6F+20/SoIpfDRjonJJgCS
L0vQee6MnQUeAg/4Zw1Igant4eu4cEYQttH0tSb9AoGBALVlnXXIeLU9+UUnAXxn
Hz/cOPZ0Xk1i3FtfHhZ++MAybUqYQc9Br6M4+5lxnhMB6QgObt0LKisQd8yOcFgL
j/9FQyUq3OlKVkdxjPIqYkjlzAPty5GFbam3EWlBj3l215dI2+5f83CSsNsL9VWU
hjBOlkunJBDtNlp60Zz0Yw8f
-----END PRIVATE KEY-----
EOF
}


data "template_file" "gcp_compute_ssl_certificates_certificate" {
  template = <<EOF
-----BEGIN CERTIFICATE-----
MIIEDTCCAvWgAwIBAgIUNJXGox7wFz/Y/qx0duC/mF+qE8UwDQYJKoZIhvcNAQEL
BQAwgZUxCzAJBgNVBAYTAlVMMQ0wCwYDVQQIDARNYWluMRAwDgYDVQQHDAdDYXBp
dGFsMRQwEgYDVQQKDAtDb3Jwb3JhdGlvbjEeMBwGA1UECwwVY29ycnVwdGlvbiBk
ZXBhcnRtZW50MREwDwYDVQQDDAhvd25lciBvZjEcMBoGCSqGSIb3DQEJARYNdGVz
dEB0ZXN0LmNvbTAeFw0yMTA5MDIwNzQ4NTRaFw0zMTA4MzEwNzQ4NTRaMIGVMQsw
CQYDVQQGEwJVTDENMAsGA1UECAwETWFpbjEQMA4GA1UEBwwHQ2FwaXRhbDEUMBIG
A1UECgwLQ29ycG9yYXRpb24xHjAcBgNVBAsMFWNvcnJ1cHRpb24gZGVwYXJ0bWVu
dDERMA8GA1UEAwwIb3duZXIgb2YxHDAaBgkqhkiG9w0BCQEWDXRlc3RAdGVzdC5j
b20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDZqZI7lNSN2xUPazPo
IgRWZJy8OYkcbSDFA9xb+Hs5jUMgs9Jf1zZ8OhUTAQL12Qofpt7XM9f2R4acdZcy
1YHYVcsFBww+q/yESBlhMeso8hPRgjqA/J8nZ0VW/IVbMJv3dU5/SPK1ITuPBLCb
sWFO5yuXxDpQOPYrA1N6UayvE/rtZgMqeWSAOy1XWgrQPRgYbm9NkXGX92VN7Byc
OdWRbu4ezGr7ZSFgNJ5VB0/n14RHS3DMsf7BMkwpg733U4J6TQ7C4Ms3FSVhqm+k
PiwQrMfQNDSxeLo1Hw8AdxzjVlouWnSm1vaFTkIBdopdLctcEnnjBktTMGlvG8qm
lXsfAgMBAAGjUzBRMB0GA1UdDgQWBBT3bfnTbUQs6brwdCJiqHBuwpMDDDAfBgNV
HSMEGDAWgBT3bfnTbUQs6brwdCJiqHBuwpMDDDAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQAeglfLaLE7E85hI05biJFZoAtd3J+8iibBLfAbyV32
Gf5salTO8GOZuy5lwr8+eI210clV628cXm3MGzicMlLuBUw3UrTFZqJIO0upQizT
u0yi9V8AjMi25Ny99GAYHlgo0rflMCacp3A9TxeUEO5XYRyU8mpvVNlLlbEyecIx
RugJMdAi8WbvfAgUef09hFn3Uz8IJyZfQ394NDwpSNJkHP55sQe5cYKLyMybNC9S
T6P7eyc3YBOCo8ohDYDHf6DFemj3DKeRw9N5w6D5H0C36LTDtL+UZDqXcUxGIiRX
L9utkN8nRSnMlX1UFFnn3BVxhrFWijGgwbtnMNLFCxKz
-----END CERTIFICATE-----
EOF
}