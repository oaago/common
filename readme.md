由于go对私有gitlab的仓库支持不好，得使用下面这些步骤

设置git使用 ssh协议 git config --global url."git@github.com:".insteadOf "https://github.com/"

添加ssh key 到gitlab

ssh-keygen 会生成 id_rsa.pub cat ~/.ssh/id_rsa.pub 粘贴到gitlab 右上角头像 Setting -> SSH
keys，或者打开链接https://gitlab.com/profile/keys
修改 go.mod 添加 replace github.com/oaago/common => github.com/oaago/common.git master

设置noproxy域名 go env -w GONOPROXY=\*\*.github.com\*\*

设置private域名 go env -w GOPRIVATE=\*\*.github.com\*\*

自己搭建的gitlab也是如此

git config --global url."git@github.com:".insteadOf "https://github.com/"

go install github.com/oaago/common@main

#### 请把示例写在demo里面

go env -w GOPRIVATE=\*\*github.com\*\*