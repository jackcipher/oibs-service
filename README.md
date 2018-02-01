# oibs-service
![Travis](https://img.shields.io/travis/rust-lang/rust.svg)![done](http://progressed.io/bar/18?title=done)

## Optimised Identity-based Signature Service

OIBS是一种在线的Restful API服务，用户对消息进行基于身份的签名。

本服务提供，消息签名、消息认证两大主要功能。





## 依赖类库

- PBC Libraray
- PBC GoWrap



## 核心API

OIBS包含两大核心API：签名与验证

签名是针对数字签名发布者，而验证是针对所有人



## 底层函数

私钥提取以及其他函数，已经完成，但没有做权限处理。

extract 提取私钥函数

sign 签名函数

verify 验证签名函数