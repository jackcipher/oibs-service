<p align="center">
<img src="https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small@2x.png?raw=true" alt="zanphp logo" srcset="https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small.png?raw=true 1x, https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small@2x.png?raw=true 2x, https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small.png?raw=true" width="210" height="210">
</p>

<p align="center">基于 PBC Library的在线数字签名服务，面向Web服务提供开箱即用的Restful API接口。</p>
<p align="center">所采用的签名方案，在随机预言机模型下的可证明安全</p>



[![License](https://img.shields.io/badge/license-mit-blue.svg)](LICENSE)
![done](http://progressed.io/bar/18?title=done)






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