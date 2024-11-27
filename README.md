# comfy4go

A Golang implementation that utilizes the ComfyUI Web API to achieve server-side functionality.

一个调用ComfyUI Web API来实现服务器端功能的Golang实现。

<p/>

<b>CurrentVersion：[v0.3.5](https://github.com/comfyanonymous/ComfyUI/releases/tag/v0.3.5) </b>

<b>当前支持的版本：[v0.3.5](https://github.com/comfyanonymous/ComfyUI/releases/tag/v0.3.5) </b>

PS：有时候，我可能会在继续某一个没有打Tag的开发版本基础上进行开发，这时我还是会使用这个版本之前的有Tag的稳定版本作为支持的当前版本号。

## 开发初衷

这是一个用于调用ComfyUI中Web接口的项目。

在编写这个项目之前，我在pkg.go.dev上搜了一些其他的实现，并且下载使用了。但是不知道是因为长时间不维护了，还是comfyUI改了API，我看到的情况是已经不能直接连接当前的版本了，会报错。

写这个实现的初衷，只是希望简单的有一个库，能够帮助我去调用搭建好的ComfyUI

PS:因为ComfyUI的更新有时候比较频繁，我并不会一直追着更新这个库。我的频率应该是在自己需要的前提下，每隔一段时间来更新需要的功能调用，并附上example

## 开发计划

因为事实上我们并不需要实现所有ComfyUI提供的Web接口，而是只需要实现一部分，就可以完成对接ComfyUI，调用其功能即可。

所以，以下是我计划进行开发的功能，如果你在使用我的这个库，并认为需要增加一下功能支持的话，可以发起issue或者联系我来说明情况~

<table>
    <thead>
        <td>序号</td>
        <td>计划功能</td>
        <td>开发状态</td>
    </thead>
    <tr>
        <td>01</td>
        <td>用户列表（方便分用户以及用户所属工作流调用）</td>
        <td><color style="color:green;">已完成</color></td>
    </tr>
    <tr>
        <td>02</td>
        <td>工作流列表以及详情</td>
        <td><color style="color:green;">已完成</color></td>
    </tr>
    <tr>
        <td>03</td>
        <td>图片上传以及管理</td>
        <td><color style="color:green;">已完成</color></td>
    </tr>
    <tr>
        <td>04</td>
        <td>图片的定时清理</td>
        <td><color style="color:red;">计划中</color></td>
    </tr>
    <tr>
        <td>05</td>
        <td>工作流执行以及消息返回</td>
        <td><color style="color:green;">已完成</color></td>
    </tr>
</table>

## 安装说明

获得/安装命令
```shell
  go get github.com/iazkaban/comfy4go
```

## 使用说明

参考pkg.go.dev