<p align="center">
<img src="https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small@2x.png?raw=true" alt="zanphp logo" srcset="https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small.png?raw=true 1x, https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small@2x.png?raw=true 2x, https://github.com/youzan/zanphp.io/blob/master/src/img/zan-logo-small.png?raw=true" width="210" height="210">
</p>

<p align="center">所采用的签名方案，在随机预言机模型下的可证明安全</p>

<p align="center">基于 PBC Library 的在线数字签名服务，面向 Web 服务提供开箱即用的 Restful API 接口。</p>



<p align="center">

[![License](https://img.shields.io/badge/license-mit-blue.svg)](LICENSE)
![done](http://progressed.io/bar/18?title=done)
</p>





## Optimised Identity-based Signature Service

Identity-based Signature是指基于身份的数字签名。OIBS是我提出的一种高效的、在随机寓言机模型下可证明安全的基于身份的数字签名方案。前期我使用了C语言在PBC密码学配对库中得到了比较好的仿真结果。为了更好地将该签名方案应用至Web系统。我继续在Golang中完成仿真实验，并不断地优化，提取核心的签名和认证功能。

OIBSS就是这种开箱即用的在线Restful API服务。本服务对普通用户是透明的，使得开发者只需专注业务本身。

本服务提供，消息签名、消息认证两大主要功能。





## 依赖类库

- [PBC Library](https://github.com/blynn/pbc)
- [PBC GoWrap](https://github.com/Nik-U/pbc)



## 核心API

OIBS 包含两大核心API：签名与验证

签名是针对数字签名发布者，而验证是针对所有人



## 底层函数

私钥提取以及其他函数，已经完成，但没有做权限处理。

extract 提取私钥函数

sign 签名函数

verify 验证签名函数



## License

OIBS 服务基于 [MIT license](https://opensource.org/licenses/MIT) 进行开源