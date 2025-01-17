# Utility for renewal or creation of apple push notification certificate
## Description
Automate steps for certificate creation or renewal for apple push notification service.

### Rough description of tasks:
1. Create a Certificate Signing Request to upload via apple portal and get it signed
2. Download the certificate and convert it so it becomes usable by [mattermost-push-proxy](https://github.com/skymailbr/proxy-push-notification)

### Result:
2 scripts, 2 manual steps:
1. scripted creation of CSR via create_csr/main.go
2. manual upload of generated `certs/csr/*.csr` file
3. manual download of signed `aps.cer` file from apple portal
4. scripted extraction + conversion from `aps.cer` to `certs/converted/*_priv.pem` to be then usable for mattermost-push-proxy

## Usage
### Prerequisites:
```bash
$ openssl version
 OpenSSL 1.1.1f  31 Mar 2020
```

### Steps
1. `$ cd cmd/renew_apple_cert/create_csr`
2. `$ cp config.sample.json config.json`
3. Fill in input information in `config/config.json`
4. `$ go run .`
5. Follow https://developers.mattermost.com/contribute/mobile/push-notifications/ios/ to upload the Certificate Signing Request *.csr generated by the script
6. Download the `aps.cer` from the apple portal and put it in `certs/downloaded/aps.cer`\
7. `$ cd ../convert_cert`
8. `$ go run .`
9. Use `certs/converted/*_priv.pem` in the push proxy configuration as described [here](https://developers.mattermost.com/contribute/mobile/push-notifications/service/#set-up-mattermost-push-notification-service-to-send-ios-push-notifications)
