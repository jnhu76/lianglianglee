<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=06&#32;集群管理：初识&#32;kubectl>
        <link rel="icon" href="/static/favicon.png">
        <title>06 集群管理：初识 kubectl </title>
        
        <link rel="stylesheet" href="/static/index.css">
        <link rel="stylesheet" href="/static/highlight.min.css">
        <script src="/static/highlight.min.js"></script>
        
        <meta name="generator" content="Hexo 4.2.0">
        <script defer src="https://umami.lianglianglee.com/script.js"
         data-website-id="83e5d5db-9d06-40e3-b780-cbae722fdf8c"></script>
    </head>

<body>
    <div class="book-container">
        <div class="book-sidebar">
            <div class="book-brand">
                <a href="/">
                    <img src="/static/favicon.png">
                    <span>技术文章摘抄</span>
                </a>
            </div>
            <div class="book-menu uncollapsible">
                <ul class="uncollapsible">
                    <li><a href="/" class="current-tab">首页</a></li>
                    <li><a href="../">上一级</a></li>
                </ul>
                <ul class="uncollapsible">
                    
                    <li>
                        <a class="menu-item" id="01  开篇： Kubernetes 是什么以及为什么需要它.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/01%20%20%e5%bc%80%e7%af%87%ef%bc%9a%20Kubernetes%20%e6%98%af%e4%bb%80%e4%b9%88%e4%bb%a5%e5%8f%8a%e4%b8%ba%e4%bb%80%e4%b9%88%e9%9c%80%e8%a6%81%e5%ae%83.md">01  开篇： Kubernetes 是什么以及为什么需要它.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 初步认识：Kubernetes 基础概念.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/02%20%e5%88%9d%e6%ad%a5%e8%ae%a4%e8%af%86%ef%bc%9aKubernetes%20%e5%9f%ba%e7%a1%80%e6%a6%82%e5%bf%b5.md">02 初步认识：Kubernetes 基础概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 宏观认识：整体架构.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/03%20%e5%ae%8f%e8%a7%82%e8%ae%a4%e8%af%86%ef%bc%9a%e6%95%b4%e4%bd%93%e6%9e%b6%e6%9e%84.md">03 宏观认识：整体架构.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 搭建 Kubernetes 集群 - 本地快速搭建.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/04%20%e6%90%ad%e5%bb%ba%20Kubernetes%20%e9%9b%86%e7%be%a4%20-%20%e6%9c%ac%e5%9c%b0%e5%bf%ab%e9%80%9f%e6%90%ad%e5%bb%ba.md">04 搭建 Kubernetes 集群 - 本地快速搭建.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 动手实践：搭建一个 Kubernetes 集群 - 生产可用.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/05%20%e5%8a%a8%e6%89%8b%e5%ae%9e%e8%b7%b5%ef%bc%9a%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%20Kubernetes%20%e9%9b%86%e7%be%a4%20-%20%e7%94%9f%e4%ba%a7%e5%8f%af%e7%94%a8.md">05 动手实践：搭建一个 Kubernetes 集群 - 生产可用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 集群管理：初识 kubectl.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/06%20%e9%9b%86%e7%be%a4%e7%ae%a1%e7%90%86%ef%bc%9a%e5%88%9d%e8%af%86%20kubectl.md">06 集群管理：初识 kubectl.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 集群管理：以 Redis 为例-部署及访问.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/07%20%e9%9b%86%e7%be%a4%e7%ae%a1%e7%90%86%ef%bc%9a%e4%bb%a5%20Redis%20%e4%b8%ba%e4%be%8b-%e9%83%a8%e7%bd%b2%e5%8f%8a%e8%ae%bf%e9%97%ae.md">07 集群管理：以 Redis 为例-部署及访问.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 安全重点 认证和授权.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/08%20%e5%ae%89%e5%85%a8%e9%87%8d%e7%82%b9%20%e8%ae%a4%e8%af%81%e5%92%8c%e6%8e%88%e6%9d%83.md">08 安全重点 认证和授权.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 应用发布：部署实际项目.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/09%20%e5%ba%94%e7%94%a8%e5%8f%91%e5%b8%83%ef%bc%9a%e9%83%a8%e7%bd%b2%e5%ae%9e%e9%99%85%e9%a1%b9%e7%9b%ae.md">09 应用发布：部署实际项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 应用管理：初识 Helm.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/10%20%e5%ba%94%e7%94%a8%e7%ae%a1%e7%90%86%ef%bc%9a%e5%88%9d%e8%af%86%20Helm.md">10 应用管理：初识 Helm.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 部署实践：以 Helm 部署项目.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/11%20%e9%83%a8%e7%bd%b2%e5%ae%9e%e8%b7%b5%ef%bc%9a%e4%bb%a5%20Helm%20%e9%83%a8%e7%bd%b2%e9%a1%b9%e7%9b%ae.md">11 部署实践：以 Helm 部署项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 庖丁解牛：kube-apiserver.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/12%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9akube-apiserver.md">12 庖丁解牛：kube-apiserver.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 庖丁解牛：etcd.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/13%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9aetcd.md">13 庖丁解牛：etcd.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 庖丁解牛：controller-manager.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/14%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9acontroller-manager.md">14 庖丁解牛：controller-manager.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 庖丁解牛：kube-scheduler.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/15%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9akube-scheduler.md">15 庖丁解牛：kube-scheduler.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 庖丁解牛：kubelet.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/16%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9akubelet.md">16 庖丁解牛：kubelet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 庖丁解牛：kube-proxy.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/17%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9akube-proxy.md">17 庖丁解牛：kube-proxy.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  庖丁解牛：Container Runtime （Docker）.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/18%20%20%e5%ba%96%e4%b8%81%e8%a7%a3%e7%89%9b%ef%bc%9aContainer%20Runtime%20%ef%bc%88Docker%ef%bc%89.md">18  庖丁解牛：Container Runtime （Docker）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 Troubleshoot.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/19%20Troubleshoot.md">19 Troubleshoot.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 扩展增强：Dashboard.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/20%20%e6%89%a9%e5%b1%95%e5%a2%9e%e5%bc%ba%ef%bc%9aDashboard.md">20 扩展增强：Dashboard.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 扩展增强：CoreDNS.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/21%20%e6%89%a9%e5%b1%95%e5%a2%9e%e5%bc%ba%ef%bc%9aCoreDNS.md">21 扩展增强：CoreDNS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 服务增强：Ingress.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/22%20%e6%9c%8d%e5%8a%a1%e5%a2%9e%e5%bc%ba%ef%bc%9aIngress.md">22 服务增强：Ingress.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 监控实践：对 K8S 集群进行监控.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/23%20%e7%9b%91%e6%8e%a7%e5%ae%9e%e8%b7%b5%ef%bc%9a%e5%af%b9%20K8S%20%e9%9b%86%e7%be%a4%e8%bf%9b%e8%a1%8c%e7%9b%91%e6%8e%a7.md">23 监控实践：对 K8S 集群进行监控.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 总结.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e4%bb%8e%e4%b8%8a%e6%89%8b%e5%88%b0%e5%ae%9e%e8%b7%b5/24%20%e6%80%bb%e7%bb%93.md">24 总结.md</a>
                    </li>
                    
                    <li><a href="https://lianglianglee.com/assets/%E6%8D%90%E8%B5%A0.md">捐赠</a></li>
                </ul>

            </div>
        </div>

        <div class="sidebar-toggle" onclick="sidebar_toggle()" onmouseover="add_inner()" onmouseleave="remove_inner()">
            <div class="sidebar-toggle-inner"></div>
        </div>
        <div class="off-canvas-content">
            <div class="columns">
                <div class="column col-12 col-lg-12">
                    <div class="book-navbar">
                        
                        <header class="navbar">
                            <section class="navbar-section">
                                <a onclick="open_sidebar()">
                                    <i class="icon icon-menu"></i>
                                </a>
                            </section>
                        </header>
                    </div>
                    <div class="book-content" style="max-width: 960px; margin: 0 auto;
    overflow-x: auto;
    overflow-y: hidden;">
                        <div class="book-post">
                            
                            
                            
                            <p id="tip" align="center"></p>
                            <h1 id="title" data-id="06 集群管理：初识 kubectl" class="title">06 集群管理：初识 kubectl</h1>
                            <div><p>从本节开始，我们来学习 K8S 集群管理相关的知识。通过前面的学习，我们知道 K8S 遵循 C/S 架构，官方也提供了 CLI 工具 <code>kubectl</code> 用于完成大多数集群管理相关的功能。当然凡是你可以通过 <code>kubectl</code> 完成的与集群交互的功能，都可以直接通过 API 完成。</p>

<p>对于我们来说 <code>kubectl</code> 并不陌生，在第 3 章讲 K8S 整体架构时，我们首次提到了它。在第 4 章和第 5 章介绍了两种安装 <code>kubectl</code> 的方式故而本章不再赘述安装的部分。</p>

<h2 id="整体概览">整体概览</h2>

<p>首先我们在终端下执行下 <code>kubectl</code>:</p>

<pre><code>➜  ~ kubectl                
kubectl controls the Kubernetes cluster manager.
...
Usage:
  kubectl [flags] [options]
</code></pre>

<p><code>kubectl</code> 已经将命令做了基本的归类，同时显示了其一般的用法 <code>kubectl [flags] [options]</code> 。</p>

<p>使用 <code>kubectl options</code> 可以看到所有全局可用的配置项。</p>

<h2 id="基础配置">基础配置</h2>

<p>在我们的用户家目录，可以看到一个名为 <code>.kube/config</code> 的配置文件，我们来看下其中的内容（此处以本地的 minikube 集群为例）。</p>

<pre><code>➜  ~ ls $HOME/.kube/config 
/home/tao/.kube/config
➜  ~ cat $HOME/.kube/config
apiVersion: v1
clusters:
- cluster:
    certificate-authority: /home/tao/.minikube/ca.crt
    server: https://192.168.99.101:8443
  name: minikube
contexts:
- context:
    cluster: minikube
    user: minikube
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: minikube
  user:
    client-certificate: /home/tao/.minikube/client.crt
    client-key: /home/tao/.minikube/client.key
</code></pre>

<p><code>$HOME/.kube/config</code> 中主要包含着：</p>

<ul>
<li>K8S 集群的 API 地址</li>
<li>用于认证的证书地址</li>
</ul>

<p>当然，我们在第 5 章时，也已经说过，也可以使用 <code>--kubeconfig</code> 或者环境变量 <code>KUBECONFIG</code> 来传递配置文件。</p>

<p>另外如果你并不想使用配置文件的话，你也可以通过使用直接传递相关参数来使用，例如：</p>

<pre><code>➜  ~ kubectl --client-key='/home/tao/.minikube/client.key' --client-certificate='/home/tao/.minikube/client.crt' --server='https://192.168.99.101:8443'  get nodes
NAME       STATUS    ROLES     AGE       VERSION
minikube   Ready     master    2d        v1.11.3
</code></pre>

<h2 id="从-get-说起">从 <code>get</code> 说起</h2>

<p>无论是第 4 章还是第 5 章，当我们创建集群后，我们都做了两个相同的事情，一个是执行 <code>kubectl get nodes</code> 另一个则是 <code>kubectl cluster-info</code>，我们先从查看集群内 <code>Node</code> 开始。</p>

<p>这里我们使用了一个本地已创建好的 <code>minikube</code> 集群。</p>

<pre><code>➜  ~ kubectl get nodes
NAME       STATUS    ROLES     AGE       VERSION
minikube   Ready     master    2d        v1.11.3
➜  ~ kubectl get node
NAME       STATUS    ROLES     AGE       VERSION
minikube   Ready     master    2d        v1.11.3
➜  ~ kubectl get no
NAME       STATUS    ROLES     AGE       VERSION
minikube   Ready     master    2d        v1.11.3
</code></pre>

<p>可以看到以上三种“名称”均可获取当前集群内 <code>Node</code> 信息。这是为了便于使用而增加的别名和缩写。</p>

<p>如果我们想要看到更详细的信息呢？可以通过传递 <code>-o</code> 参数以得到不同格式的输出。</p>

<pre><code>➜  ~ kubectl get nodes -o wide 
NAME       STATUS    ROLES     AGE       VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE            KERNEL-VERSION   CONTAINER-RUNTIME
minikube   Ready     master    2d        v1.11.3   10.0.2.15     &lt;none&gt;        Buildroot 2018.05   4.15.0           docker://17.12.1-ce
</code></pre>

<p>当然也可以传递 <code>-o yaml</code> 或者 <code>-o json</code> 得到更加详尽的信息。</p>

<p>使用 <code>-o json</code> 将内容以 JSON 格式输出时，可以配合 <a href="https://stedolan.github.io/jq/" target="_blank"><code>jq</code></a> 进行内容提取。例如：</p>

<pre><code>➜  ~ kubectl get nodes -o json | jq &quot;.items[] | {name: .metadata.name} + .status.nodeInfo&quot;
{
  &quot;name&quot;: &quot;minikube&quot;,
  &quot;architecture&quot;: &quot;amd64&quot;,
  &quot;bootID&quot;: &quot;d675d75b-e58e-40db-8910-6e5dda9e7cf9&quot;,
  &quot;containerRuntimeVersion&quot;: &quot;docker://17.12.1-ce&quot;,
  &quot;kernelVersion&quot;: &quot;4.15.0&quot;,
  &quot;kubeProxyVersion&quot;: &quot;v1.11.3&quot;,
  &quot;kubeletVersion&quot;: &quot;v1.11.3&quot;,
  &quot;machineID&quot;: &quot;078e2d22629747178397e29cf1c96cc7&quot;,
  &quot;operatingSystem&quot;: &quot;linux&quot;,
  &quot;osImage&quot;: &quot;Buildroot 2018.05&quot;,
  &quot;systemUUID&quot;: &quot;4073906D-69A1-46EE-A08C-0252D9F79893&quot;
}
</code></pre>

<p>以此方法可得到 <code>Node</code> 的基础信息。</p>

<p>那么除了 <code>Node</code> 外我们还可以查看那些资源或别名呢？可以通过 <code>kubectl api-resources</code> 查看服务端支持的 API 资源及别名和描述等信息。</p>

<h2 id="答疑解惑-explain">答疑解惑 <code>explain</code></h2>

<p>当通过上面的命令拿到所有支持的 API 资源列表后，虽然后面基本都有一个简单的说明，是不是仍然感觉一头雾水？</p>

<p>别担心，在我们使用 Linux 的时候，我们有 <code>man</code> ，在使用 <code>kubectl</code> 的时候，我们除了 <code>--help</code> 外还有 <code>explain</code> 可帮我们进行说明。 例如：</p>

<pre><code>➜  ~ kubectl explain node                                                                                                              
KIND:     Node
VERSION:  v1

DESCRIPTION:              
     Node is a worker node in Kubernetes. Each node will have a unique
     identifier in the cache (i.e. in etcd).

# ... 省略输出
</code></pre>

<h2 id="总结">总结</h2>

<p>本节我们大致介绍了 <code>kubectl</code> 的基础使用，尤其是最常见的 <code>get</code> 命令。可通过传递不同参数获取不同格式的输出，配合 <code>jq</code> 工具可方便的进行内容提取。</p>

<p>以及关于 <code>kubectl</code> 的配置文件和无配置文件下通过传递参数直接使用等。</p>

<p>对应于我们前面提到的 K8S 架构，本节相当于 <code>CURD</code> 中的 <code>R</code> 即查询。查询对于我们来说，既是我们了解集群的第一步，同时也是后续验证操作结果或集群状态必不可少的技能。</p>

<p>当然，你在集群管理中可能会遇到各种各样的问题，单纯依靠 <code>get</code> 并不足以定位问题，我们在第 21 节中将介绍 Troubleshoot 的思路及方法。</p>

<p>下节我们来学习关于 <code>C</code> 的部分，即创建。</p>
</div>
                        </div>
                        <div>
                            <div id="prePage" style="float: left">

                            </div>
                            <div id="nextPage" style="float: right">

                            </div>
                        </div>

                    </div>
                </div>
            </div>
            <div class="copyright">
                <hr />
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e68a8a8adfd2d7d7d6d1a6818b878f8ac885898b" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f104d93c8d363c6',t:'MTczNDAzMzY1MS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>