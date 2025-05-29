# About
This repository contains simplified version of [upstream CloudMuX](https://github.com/yunionio/cloudmux) for research purposes.

Field mapping between different Cloud providers is [here](https://github.com/yunionio/cloudmux/blob/0718ac214f296c5c94d0d0ab7e3d0c6d51b358ba/pkg/cloudprovider/cloudprovider.go#L106-L143). Translated version is below.

Account information. The fields for each platform are different.
The following are the fields required for account creation on each platform:
| Cloud Platform | Field | Translation | Required | Default Value | Updateable | How to Get |
| ------ |------ | ------ | --------- | -------- |-------- |-------- |
|Aliyun |access_key_id |Secret Key ID | Yes | | Yes | |
|Aliyun |access_key_secret |Secret Key | Yes | | Yes | |
|Azure |directory_id |Directory ID | Yes | | No | |
|Azure |environment |Region | Yes | | No | |
|Azure |client_id |Client ID | Yes | | Yes | |
|Azure |client_secret |Client secret | Yes | | Yes | |
|Aws |access_key_id |Secret key ID | Yes | | Yes | |
|Aws |access_key_secret |Secret key | Yes | | Yes | |
|Aws |environment |Region | Yes | | No | |
|Google |project_id |Project ID | Yes | | No | |
|Google |client_email |Client email | Yes | | No | |
|Google |private_key_id |Secret key ID | Yes | | Yes | |
|Google |private_key |Secret keyKey | yes | | yes | |
|VMware |username |User name | Yes | | Yes | |
|VMware |password |Password | Yes | | Yes | |
|VMware |host |Host IP or domain name | Yes | | No | |
|VMware |port |Host port | No | |443 | No | |
