<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=21&#32;扩展增强：CoreDNS>
        <link rel="icon" href="/static/favicon.png">
        <title>21 扩展增强：CoreDNS </title>
        
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
                            <h1 id="title" data-id="21 扩展增强：CoreDNS" class="title">21 扩展增强：CoreDNS</h1>
                            <div><h2 id="整体概览">整体概览</h2>

<p>通过前面的学习，我们知道在 K8S 中有一套默认的<a href="https://github.com/kubernetes/dns" target="_blank">集群内 DNS 服务</a>，我们通常把它叫做 <code>kube-dns</code>，它基于 SkyDNS，为我们在服务注册发现方面提供了很大的便利。</p>

<p>比如，在我们的示例项目 <a href="https://github.com/tao12345666333/saythx" target="_blank">SayThx</a> 中，各组件便是依赖 DNS 进行彼此间的调用。</p>

<p>本节，我们将介绍的 <a href="https://coredns.io/" target="_blank">CoreDNS</a> 是 CNCF 旗下又一孵化项目，在 K8S 1.9 版本中加入并进入 Alpha 阶段。我们当前是以 K8S 1.11 的版本进行介绍，它并不是默认的 DNS 服务，但是<a href="https://github.com/kubernetes/enhancements/issues/427" target="_blank">它作为 K8S 的 DNS 插件的功能已经 GA</a> 。</p>

<p>CoreDNS 在 K8S 1.13 版本中才正式成为<a href="https://kubernetes.io/blog/2018/12/03/kubernetes-1-13-release-announcement/" target="_blank">默认的 DNS 服务</a>。</p>

<h2 id="coredns-是什么">CoreDNS 是什么</h2>

<p>首先，我们需要明确 CoreDNS 是一个独立项目，它不仅可支持在 K8S 中使用，你也可以在你任何需要 DNS 服务的时候使用它。</p>

<p>CoreDNS 使用 Go 语言实现，部署非常方便。</p>

<p>它的扩展性很强，很多功能特性都是通过插件完成的，它不仅有大量的<a href="https://coredns.io/plugins/" target="_blank">内置插件</a>，同时也有很丰富的<a href="https://coredns.io/explugins/" target="_blank">第三方插件</a>。甚至你自己<a href="https://coredns.io/2016/12/19/writing-plugins-for-coredns/" target="_blank">写一个插件</a>也非常的容易。</p>

<h2 id="如何安装使用-coredns">如何安装使用 CoreDNS</h2>

<p>我们这里主要是为了说明如何在 K8S 环境中使用它，所以对于独立安装部署它不做说明。</p>

<p>本小册中我们使用的是 K8S 1.11 版本，在第 5 小节 《搭建 Kubernetes 集群》中，我们介绍了使用 <code>kubeadm</code> 搭建集群。</p>

<p>使用 <code>kubeadm</code> 创建集群时候 <code>kubeadm init</code> 可以传递 <code>--feature-gates</code> 参数，用于启用一些额外的特性。</p>

<p>比如在之前版本中，我们可以通过 <code>kubeadm init --feature-gates CoreDNS=true</code> 在创建集群时候启用 CoreDNS。</p>

<p>而在 1.11 版本中，使用 <code>kubeadm</code> 创建集群时 <code>CoreDNS</code> 已经被默认启用，这也从侧面证明了 CoreDNS 在 K8S 中达到了生产可用的状态。</p>

<p>我们来看一下创建集群时的日志输出：</p>

<pre><code>[root@master ~]# kubeadm init
[init] using Kubernetes version: v1.11.3               
[preflight] running pre-flight checks
...
[bootstraptoken] creating the &quot;cluster-info&quot; ConfigMap in the &quot;kube-public&quot; namespace
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes master has initialized successfully!
</code></pre>

<p>可以看到创建时已经启用了 CoreDNS 的扩展，待集群创建完成后，可用过以下方式进行查看：</p>

<pre><code>master $ kubectl -n kube-system get all  -l k8s-app=kube-dns -o wide
NAME                           READY     STATUS    RESTARTS   AGE       IP          NODE      NOMINATED NODE
pod/coredns-78fcdf6894-5zbx4   1/1       Running   0          1h        10.32.0.3   node01    &lt;none&gt;
pod/coredns-78fcdf6894-cxdw8   1/1       Running   0          1h        10.32.0.2   node01    &lt;none&gt;

NAME               TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)         AGE       SELECTOR
service/kube-dns   ClusterIP   10.96.0.10   &lt;none&gt;        53/UDP,53/TCP   1h        k8s-app=kube-dns

NAME                      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE       CONTAINERS   IMAGES                     SELECTOR
deployment.apps/coredns   2         2         2            2           1h        coredns      k8s.gcr.io/coredns:1.1.3   k8s-app=kube-dns

NAME                                 DESIRED   CURRENT   READY     AGE       CONTAINERS   IMAGES                   SELECTOR
replicaset.apps/coredns-78fcdf6894   2         2         2         1h        coredns      k8s.gcr.io/coredns:1.1.3   k8s-app=kube-dns,pod-template-hash=3497892450
</code></pre>

<p>这里主要是为了兼容 K8S 原有的 <code>kube-dns</code> 所以标签和 <code>Service</code> 的名字都还使用了 <code>kube-dns</code>，但实际在运行的则是 CoreDNS。</p>

<h2 id="验证-coredns-功能">验证 CoreDNS 功能</h2>

<p>从上面的输出我们看到 CoreDNS 的 <code>Pod</code> 运行正常，现在测试下它是否能正确解析。仍然以我们的示例项目 <a href="https://github.com/tao12345666333/saythx" target="_blank">SayThx</a> 为例，先 clone 项目，进入到项目的 deploy 目录中。</p>

<pre><code>master $ cd saythx/deploy/
master $ ls
backend-deployment.yaml  frontend-deployment.yaml  namespace.yaml         redis-service.yaml
backend-service.yaml     frontend-service.yaml     redis-deployment.yaml  work-deployment.yaml
master $ kubectl apply -f namespace.yaml
namespace/work created
master $ kubectl apply -f redis-deployment.yaml
deployment.apps/saythx-redis created
master $ kubectl apply -f redis-service.yaml
service/saythx-redis created
</code></pre>

<ul>
<li>查看其部署情况：</li>
</ul>

<pre><code>master $ kubectl -n work get all
NAME                               READY     STATUS    RESTARTS   AGE
pod/saythx-redis-8558c7d7d-8v4lp   1/1       Running   0          2m

NAME                   TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/saythx-redis   NodePort   10.109.215.147   &lt;none&gt;        6379:31438/TCP   2m

NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/saythx-redis   1         1         1            1           2m

NAME                                     DESIRED   CURRENT   READY     AGE
replicaset.apps/saythx-redis-8558c7d7d   1         1         1         2m
</code></pre>

<ul>
<li>验证 DNS 是否正确解析:</li>
</ul>

<pre><code># 使用 AlpineLinux 的镜像创建一个 Pod 并进入其中
master $ kubectl run alpine -it --rm --restart='Never' --image='alpine' sh
If you don't see a command prompt, try pressing enter.
/ # apk add --no-cache bind-tools
fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/community/x86_64/APKINDEX.tar.gz
(1/5) Installing libgcc (6.4.0-r9)
(2/5) Installing json-c (0.13.1-r0)
(3/5) Installing libxml2 (2.9.8-r1)
(4/5) Installing bind-libs (9.12.3-r0)
(5/5) Installing bind-tools (9.12.3-r0)
Executing busybox-1.28.4-r2.trigger
OK: 9 MiB in 18 packages

# 安装完 dig 命令所在包之后，使用 dig 命令进行验证
/ # dig @10.32.0.2 saythx-redis.work.svc.cluster.local +noall +answer

; &lt;&lt;&gt;&gt; DiG 9.12.3 &lt;&lt;&gt;&gt; @10.32.0.2 saythx-redis.work.svc.cluster.local +noall +answer
; (1 server found)
;; global options: +cmd
saythx-redis.work.svc.cluster.local. 5 IN A     10.109.215.147
</code></pre>

<p>通过以上操作，可以看到相应的 <code>Service</code> 记录可被正确解析。这里有几个点需要注意：</p>

<ul>
<li>域名解析是可跨 <code>Namespace</code> 的</li>
</ul>

<p>刚才的示例中，我们没有指定 <code>Namespace</code> 所以刚才我们所在的 <code>Namespace</code> 是 <code>default</code>。而我们的解析实验成功了。说明 CoreDNS 的解析是全局的。<strong>虽然解析是全局的，但不代表网络互通</strong></p>

<ul>
<li>域名有特定格式</li>
</ul>

<p>可以看到刚才我们使用的完整域名是 <code>saythx-redis.work.svc.cluster.local</code> , 注意开头的便是 <strong>Service 名.Namespace 名</strong> 当然，我们也可以直接通过 <code>host</code> 命令查询:</p>

<pre><code>  / # host -t srv  saythx-redis.work
  saythx-redis.work.svc.cluster.local has SRV record 0 100 6379 saythx-redis.work.svc.cluster.local.
</code></pre>

<h2 id="配置和监控">配置和监控</h2>

<p>CoreDNS 使用 <code>ConfigMap</code> 的方式进行配置，但是如果更改了配置，<code>Pod</code> 重启后才会生效。</p>

<p>我们通过以下命令可查看其配置：</p>

<pre><code>master $ kubectl -n kube-system get configmap coredns -o yaml
apiVersion: v1
data:
  Corefile: |
    .:53 {
        errors
        health
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           upstream
           fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        proxy . /etc/resolv.conf
        cache 30
        reload
    }
kind: ConfigMap
metadata:
  creationTimestamp: 2018-12-22T16:45:47Z
  name: coredns
  namespace: kube-system
  resourceVersion: &quot;217&quot;
  selfLink: /api/v1/namespaces/kube-system/configmaps/coredns
  uid: 0882e51b-0609-11e9-b25e-0242ac110057
</code></pre>

<p><code>Corefile</code> 便是它的配置文件，可以看到它启动了类似 <code>kubernetes</code>, <code>prometheus</code> 等插件。</p>

<p>注意 <code>kubernetes</code> 插件的配置，使用的域是 <code>cluster.local</code> ，这也是上面我们提到域名格式时候后半部分未解释的部分。</p>

<p>至于 <code>prometheus</code> 插件，则是监听在 9153 端口上提供了符合 Prometheus 标准的 metrics 接口，可用于监控等。关于监控的部分，可参考第 23 节。</p>

<h2 id="总结">总结</h2>

<p>在本节中，我们介绍了 CoreDNS 的基本情况，它是以 Go 编写的灵活可扩展的 DNS 服务器。</p>

<p>使用 CoreDNS 代替 kube-dns 主要是为了解决一些 kube-dns 时期的问题，比如说原先 kube-dns 的时候，一个 Pod 中还需要包含 <code>kube-dns</code>, <code>sidecar</code> 和 <code>dnsmasq</code> 的容器，而每当 <code>dnsmasq</code> 出现漏洞时，就不得不让 K8S 发个安全补丁才能进行更新。</p>

<p>CoreDNS 有丰富的插件，可以满足更多样的应用需求，同时 <code>kubernetes</code> 插件还包含了一些独特的功能，比如 Pod 验证之类的，可增加安全性。</p>

<p>同时 CoreDNS 在 1.13 版本中会作为默认的 DNS 服务器使用，所以应该给它更多的关注。</p>

<p>在下节，我们将介绍 <code>Ingress</code>，看看如果使用不同于之前使用的 <code>NodePort</code> 的方式将服务暴露于外部。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b6dadada8f8287878681f6d1dbd7dfda98d5d9db" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1050186d4063c6',t:'MTczNDAzMzc1NS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>