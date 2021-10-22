<h1 align="center">firefly-wallet </h1>

<!--
<p align="center">
  <a href="https://circleci.com/gh/filecoin-project/lotus"><img src="https://circleci.com/gh/filecoin-project/lotus.svg?style=svg"></a>
  <a href="https://codecov.io/gh/filecoin-project/lotus"><img src="https://codecov.io/gh/filecoin-project/lotus/branch/master/graph/badge.svg"></a>
  <a href="https://goreportcard.com/report/github.com/filecoin-project/lotus"><img src="https://goreportcard.com/badge/github.com/filecoin-project/lotus" /></a>  
  <a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.15.5-blue.svg" /></a>
  <br>
</p>
-->

lotus-wallet-tool用于管理客户钱包账户，可以创建钱包，支持转账，miner提现，设置owner，worker key等等。

## 目录

- [构建](#构建)
- [安装](#安装)
- [使用](#使用)

## 构建

### 准备构建环境

- linux

```
   sudo apt install build-essential golang libhwloc-dev  

```

- macos

```
   xcode-select --install 
   brew install go bzr jq pkg-config rustup hwloc 
```

### 开始构建

```
   git clone --recursive https://github.com/fireflyblock/firefly-wallet
   cd firefly-wallet
make 
```

## 安装

### 准备运行环境

- linux

```
   sudo apt install hwloc
```

- macos

```
   brew install hwloc
```

### 安装二进制

- 从构建目录复制

```
   cp firefly-wallet /usr/local/bin
   chmod +x /usr/local/bin/firefly-wallet
```

- 从[github](https://github.com/fireflyblock/firefly-wallet)发布页下载后复制


## 使用

*注意*： 由于涉及资产规模较大，请在专业人士指导下使用。

### 导出lotus全节点API 

```
  export FULLNODE_API_INFO=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUifX0.Y4KESS8FmSOF4Kkk6So2-hjm6Jy3dKyUyYF4Ra-iork:/ip4/115.236.21.26/tcp/1234/
```

### 导入助记词

- 编辑 mnemonic.txt，输入12个助记词
- 使用init命令导入助记词

```
$ firefly-wallet init --key-file  mnemonic.txt 
请输入密码(长度至少6位):******
请再次输入密码：******
{"mnemonic":"Izp0/4JZAWwM2UiW5wP4fS+X2bH7x0UuFfz7OzT2eV6f7fiYvDTKEYWE6/tvk8O0RKPhYOeV5G6NULFdZrAcNsRmC1SI6im7/Ap+NnopHTE=","iv":"bkmnXfbmKbND5V/IzsiSLA==","salt":"0N1N0TML0SG1Berww1/PPekN1uHIt3f7kNlYcnMgTyM="}

f1mb62jel5pe2ni4lmaqgsppskdegtscqwvubpx6i
```

导入后助记词后，会在$HOME/.lotuswallettool建立数据库，存储加密后的助记词。

*注意*： 导入助记词后，务必删除助记词明文文件。
*注意*： 保管好助记词。
*注意*： 如果原有数据不需要，可以在保管好助记词的前提下，通过删除 $HOME/.lotuswallettool目录来清理环境。

### 列出钱包地址及余额

```
$ firefly-wallet list
请输入密码(长度至少6位):******
Address                                    Balance  Nonce  
f1mb62jel5pe2ni4lmaqgsppskdegtscqwvubpx6i  0 FIL    0
```
### 创建新地址

不带BLS选项生成的地址，使用SECP256-K1生成的地址，与以太坊签名算法兼容。  
与外部交易所和钱包之间，最好采用这种形式，通用和兼容性较好。  
```
$ firefly-wallet new-address
请输入密码(长度至少6位):******
f1d5rhpq5icldkw4djdis5oufefutgyqkqunmb2oy

```

带上BLS选项生成的地址，使用BLS签名算法生成的地址，可以用于矿工worker key。  

```
$ firefly-wallet new-address --bls
请输入密码(长度至少6位):******
f3wzaczkndddpf55nlmkyni5q3mqgkgycvj6hnqp2b27zo2yvtbl27ru7dykghi63clxtrv4dsvzvhrmb65coq
```

### 导出私钥（慎重）


```
$ ./firefly-wallet  list
请输入密码(长度至少6位):******
Address                                                                                 Balance  Nonce
f12uc2ntixp4kmx7j5yk7pnrur2w6ggeutm45z6ja                                               0 FIL    0
f174jo52r2rq5b7nqiy3eqphheqxsis2o2avwotbq                                               0 FIL    0
f3qa2axjy5ejuzbvyk3weja7hvnjb7y4mh2s3y4b44nhfzmdhgly6edmcdjar6nfy7vccjvjrswom5pqera52q  0 FIL    0

$ ./firefly-wallet  export-address --address  f3qa2axjy5ejuzbvyk3weja7hvnjb7y4mh2s3y4b44nhfzmdhgly6edmcdjar6nfy7vccjvjrswom5pqera52q
请输入密码(长度至少6位):******
7b2254797065223a22626c73222c22507269766172222b6579223a222b6c37557763586952756f34492f366d2b56514566511111756338746d616c4f6e7676324b66506c62696333337d
```


### 从矿工帐号提现

```
$ firefly-wallet withdraw f0144528 100
&{Version:0 To:f0839084 From:f3sr6lzdwbhq56v5u2emunag4gwg6xdjsm2spgfdrksro47bkf7txk5vooirrttjeyxrvtjmbxeay5ixnd4jaa Nonce:94879 Value:+0 GasLimit:11799568 GasFeeCap:+1694977307 GasPremium:+100514 Method:16 Params:[129 74 0 12 119 228 37 104 99 216 0 0]}
Requested rewards withdrawal in message bafy3bzacecato7b8eyqelweyagargtikrk5hsjx3uep2tvrumkd4ygtxdfava
```

提现后可在区块链浏览器filfox/filscout的矿工帐号下查询交易记录。

### 转帐

```
$ firefly-wallet send --from f3rri57tna736dimimctjh5mig4cxpkuvwvvqk5f2ewr3vssg7kl4b3m2pecvoe5sks356ry2yuklidov6yc5q f3uivirtgact2spxjruxynschbzegnhrf2ynv23tdsygkzkvq5eyigzplhcsxzhkqxpegnjckvd4egtbchicya 194.26
bafy2bzaceblmwzxvmdrnx2lqaoe2qiirpy3tpzd6okt6svu9s6ph6pqlbqwle

```
提现后可在区块链浏览器filfox/filscout的owner帐号下查询交易记录。


### 导入旧owner账号

```
$ firefly-wallet   import  
请输入密码(长度至少6位):******
请输入私钥: 7b2254797065223a22736563703235366b31222c22507269766174654b6579223a2238463966687a5a41434a7530643734576d6a59587537533670554c45723962566f477876384436644954493d227d
成功导入钱包： f1d5rhpq5icldkw4djdis5oufefutgyqkqunmb2oy

```

### 更换owner key

设置过程中这个命令需要被执行两次, 第一次用旧的ownr地址发送, 第二次用新的owner地址发送。这样可以防止用户把owner地址转移到一个不受自己控制的账号上。

```
$ firefly-wallet set-owner --really-do-it f02420 f1xxxxxnew  f1xxxxxold
$ firefly-wallet set-owner --really-do-it f02420 f1xxxxxnew  f1xxxxxnew
```

### 更换worker key

- worker key必须是BLS地址，参见前述地址生成。
- 导出私钥，参见前述导出私钥

将私钥放入一个文本文件，并用密码加密后的worker_key.txt.gpg通过邮件发送给运维，密码通过其他通信方式告知。

```
$ echo 7b2254797065223a22626c73222c22507269766172222b6579223a222b6c37557763586952756f34492f366d2b56514566511111756338746d616c4f6e7676324b66506c62696333337d > /tmp/worker_key.txt
$ gpg  --symmetric /tmp/worker_key.txt
$ ls /tmp/worker_key.txt.gpg
/tmp/worker_key.txt.gpg
```
- 更换worker key

*注意*：
- 更换worker key后，消息上链7.5小时后生效。  
- 需要在生效前，加入到具备签名能力的钱包服务中,否则会导致挖矿出块失败。  

```
$ firefly-wallet  propose-change-worker  --really-do-it  f02420  f3qbjatohy7evsb4hi7qjbyig6egpclztrgyhc25vaqdvtdhxip3aqkzfhtw57k5r2nc6tobves66qdak75msa
请输入密码(长度至少6位):******
Propose Message CID: bafy2bzaceazzo2hhwhg2amfcmjhtlcd5etg5pwlcetyjups4iyeolbjqaewvs
Worker key change to f3qbjatohy7evsb4hi7qjbyig6egpclztrgyhc25vaqdvtdhxip3aqkzfhtw57k5r2nc6tobves66qdak75msa successfully proposed.
Call 'confirm-change-worker' at or after height 1010006 to complete.
```
