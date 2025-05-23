<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=19&#32;Troubleshoot>
        <link rel="icon" href="/static/favicon.png">
        <title>19 Troubleshoot </title>
        
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
                            <h1 id="title" data-id="19 Troubleshoot" class="title">19 Troubleshoot</h1>
                            <div><h2 id="整体概览">整体概览</h2>

<p>通过前面的介绍，我们已经了解到了 K8S 的基础知识，核心组件原理以及如何在 K8S 中部署服务及管理服务等。</p>

<p>但在生产环境中，我们所面临的环境多种多样，可能会遇到各种问题。本节将结合我们已经了解到的知识，介绍一些常见问题定位和解决的思路或方法，以便大家在生产中使用 K8S 能如鱼得水。</p>

<h2 id="应用部署问题">应用部署问题</h2>

<p>首先我们从应用部署相关的问题来入手。这里仍然使用我们的<a href="https://github.com/tao12345666333/saythx" target="_blank">示例项目 SayThx</a>。</p>

<p>clone 该项目，进入到 deploy 目录中，先 <code>kubectl apply -f namespace.yaml</code> 或者 <code>kubectl create ns work</code> 来创建一个用于实验的 <code>Namespace</code> 。</p>

<h3 id="使用-describe-排查问题">使用 <code>describe</code> 排查问题</h3>

<p>对 <code>redis-deployment.yaml</code> 稍作修改，按以下方式操作：</p>

<pre><code>master $ kubectl apply -f redis-deployment.yaml
deployment.apps/saythx-redis created
master $ kubectl -n work get all
NAME                                READY     STATUS             RESTARTS   AGE
pod/saythx-redis-7574c98f5d-v66fx   0/1       ImagePullBackOff   0          9s

NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/saythx-redis   1         1         1            0           9s

NAME                                      DESIRED   CURRENT   READY     AGE
replicaset.apps/saythx-redis-7574c98f5d   1         1         0         9s
</code></pre>

<p>可以看到 <code>Pod</code> 此刻的状态是 <code>ImagePullBackOff</code>，这个状态表示镜像拉取失败，<code>kubelet</code> 退出镜像拉取。</p>

<p>我们在前面的内容中介绍过 <code>kubelet</code> 的作用之一就是负责镜像拉取，而实际上，在镜像方面的错误主要预设了 6 种，分别是 <code>ImagePullBackOff</code>，<code>ImageInspectError</code>，<code>ErrImagePull</code>，<code>ErrImageNeverPull</code>，<code>RegistryUnavailable</code>，<code>InvalidImageName</code>。</p>

<p>当遇到以上所述情况时，便可定位为镜像相关异常。</p>

<p>我们回到上面的问题当中，定位问题所在。</p>

<pre><code>master $ kubectl -n work describe pod/saythx-redis-7574c98f5d-v66fx
Name:               saythx-redis-7574c98f5d-v66fx
Namespace:          work
Priority:           0
PriorityClassName:  &lt;none&gt;
Node:               node01/172.17.0.132
Start Time:         Tue, 18 Dec 2018 17:27:56 +0000
Labels:             app=redis
                    pod-template-hash=3130754918
Annotations:        &lt;none&gt;
Status:             Pending
IP:                 10.40.0.1
Controlled By:      ReplicaSet/saythx-redis-7574c98f5d
Containers:
  redis:
    Container ID:
    Image:          redis:5xx
    Image ID:
    Port:           6379/TCP
    Host Port:      0/TCP
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False
    Restart Count:  0
    Environment:    &lt;none&gt;
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-787w5 (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  default-token-787w5:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-787w5
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  &lt;none&gt;
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type     Reason          Age                 From               Message
  ----     ------          ----                ----               -------
  Normal   Scheduled       11m                 default-scheduler  Successfully assigned work/saythx-redis-7574c98f5d-v66fx to node01
  Normal   SandboxChanged  10m                 kubelet, node01    Pod sandbox changed, it will bekilled and re-created.
  Normal   BackOff         9m (x6 over 10m)    kubelet, node01    Back-off pulling image &quot;redis:5xx&quot;
  Normal   Pulling         9m (x4 over 10m)    kubelet, node01    pulling image &quot;redis:5xx&quot;
  Warning  Failed          9m (x4 over 10m)    kubelet, node01    Failed to pull image &quot;redis:5xx&quot;: rpc error: code = Unknown desc = Error response from daemon: manifest for redis:5xx not found
  Warning  Failed          9m (x4 over 10m)    kubelet, node01    Error: ErrImagePull
  Warning  Failed          49s (x44 over 10m)  kubelet, node01    Error: ImagePullBackOff
</code></pre>

<p>可以看到我们现在 pull 的镜像是 <code>redis:5xx</code> 而实际上并不存在此 tag 的镜像，所以导致拉取失败。</p>

<h3 id="使用-events-排查问题">使用 <code>events</code> 排查问题</h3>

<p>当然，我们还有另一种方式同样可进行问题排查：</p>

<pre><code>master $ kubectl -n work get events
LAST SEEN   FIRST SEEN   COUNT     NAME                                             KIND         SUBOBJECT                TYPE      REASON              SOURCE                  MESSAGE
21m         21m          1         saythx-redis.15717d6361a741a8                    Deployment                        Normal    ScalingReplicaSet   deployment-controller   Scaled up replica set saythx-redis-7574c98f5d to 1
21m         21m          1         saythx-redis-7574c98f5d-qwxgm.15717d6363eb60ff   Pod                        Normal    Scheduled           default-scheduler       Successfully assigned work/saythx-redis-7574c98f5d-qwxgm to node01
21m         21m          1         saythx-redis-7574c98f5d.15717d636309afa8         ReplicaSet                        Normal    SuccessfulCreate    replicaset-controller   Created pod: saythx-redis-7574c98f5d-qwxgm
20m         21m          2         saythx-redis-7574c98f5d-qwxgm.15717d63fa501b3f   Pod          spec.containers{redis}   Normal    BackOff             kubelet, node01         Back-off pulling image &quot;redis:5xx&quot;
20m         21m          2         saythx-redis-7574c98f5d-qwxgm.15717d63fa5049a9   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Error: ImagePullBackOff
20m         21m          3         saythx-redis-7574c98f5d-qwxgm.15717d6393a1993c   Pod          spec.containers{redis}   Normal    Pulling             kubelet, node01         pulling image &quot;redis:5xx&quot;
20m         21m          3         saythx-redis-7574c98f5d-qwxgm.15717d63e11efc7a   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Error: ErrImagePull
20m         21m          3         saythx-redis-7574c98f5d-qwxgm.15717d63e11e9c25   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Failed to pull image &quot;redis:5xx&quot;: rpc error: code = Unknown desc = Error response from daemon: manifest for redis:5xxnot found
20m         20m          1         saythx-redis-54984ff94-2bb6g.15717d6dc03799cd    Pod          spec.containers{redis}   Normal    Killing             kubelet, node01         Killing container with id docker://redis:Need to kill Pod
19m         19m          1         saythx-redis-7574c98f5d-v66fx.15717d72356528ec   Pod                        Normal    Scheduled           default-scheduler       Successfully assigned work/saythx-redis-7574c98f5d-v66fx to node01
19m         19m          1         saythx-redis-7574c98f5d.15717d722f7f1732         ReplicaSet                        Normal    SuccessfulCreate    replicaset-controller   Created pod: saythx-redis-7574c98f5d-v66fx
19m         19m          1         saythx-redis.15717d722b49e758                    Deployment                        Normal    ScalingReplicaSet   deployment-controller   Scaled up replica set saythx-redis-7574c98f5d to 1
19m         19m          1         saythx-redis-7574c98f5d-v66fx.15717d731a09b0ad   Pod                        Normal    SandboxChanged      kubelet, node01         Pod sandbox changed, it will be killed and re-created.
18m         19m          6         saythx-redis-7574c98f5d-v66fx.15717d733ab20b3d   Pod          spec.containers{redis}   Normal    BackOff             kubelet, node01         Back-off pulling image &quot;redis:5xx&quot;
18m         19m          4         saythx-redis-7574c98f5d-v66fx.15717d729de13541   Pod          spec.containers{redis}   Normal    Pulling             kubelet, node01         pulling image &quot;redis:5xx&quot;
18m         19m          4         saythx-redis-7574c98f5d-v66fx.15717d72e6ded95d   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Error: ErrImagePull
18m         19m          4         saythx-redis-7574c98f5d-v66fx.15717d72e6de7b1c   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Failed to pull image &quot;redis:5xx&quot;: rpc error: code = Unknown desc = Error response from daemon: manifest for redis:5xxnot found
4m          19m          66        saythx-redis-7574c98f5d-v66fx.15717d733ab23f2c   Pod          spec.containers{redis}   Warning   Failed              kubelet, node01         Error: ImagePullBackOff
master
</code></pre>

<p>我们在之前介绍时，也提到过 <code>kubelet</code> 或者 <code>kube-scheduler</code> 等组件会接受某些事件等，<code>event</code> 便是用于记录集群内各处发生的事件之类的。</p>

<h3 id="修正错误">修正错误</h3>

<ul>
<li>修正配置文件</li>
</ul>

<p>修正配置文件，然后 <code>kubectl apply -f redis-deployment.yaml</code> 便可应用修正后的配置文件。这种方法比较推荐，并且可以将修改过的位置纳入到版本控制系统中，有利于后续维护。</p>

<ul>
<li>在线修改配置</li>
</ul>

<p>使用 <code>kubectl -n work edit deploy/saythx-redis</code>，会打开默认的编辑器，我们可以将使用的镜像及 tag 修正为 <code>redis:5</code> 保存退出，便会自动应用新的配置。这种做法比较适合比较紧急或者资源是直接通过命令行创建等情况。 <strong>非特殊情况尽量不要在线修改。</strong> 且这样修改并不利于后期维护。</p>

<h3 id="通过详细内容排查错误">通过详细内容排查错误</h3>

<pre><code>master $ kubectl apply -f namespace.yaml
namespace/work created
master $ kubectl apply -f redis-deployment.yaml
deployment.apps/saythx-redis created
master $ vi redis-service.yaml # 稍微做了点修改
master $ kubectl apply -f redis-service.yaml
service/saythx-redis created
master $ kubectl -n work get pods,svc
NAME                               READY     STATUS    RESTARTS   AGE
pod/saythx-redis-8558c7d7d-z8prg   1/1       Running   0          47s

NAME                   TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/saythx-redis   NodePort   10.108.202.170   &lt;none&gt;        6379:32355/TCP   16s
</code></pre>

<p>通过以上的输出，大多数情况下我们的 <code>Service</code> 应该是可以可以正常访问了，现在我们进行下测试：</p>

<pre><code>master $ docker run --rm -it --net host redis redis-cli -p 32355
Unable to find image 'redis:latest' locally
latest: Pulling from library/redis
a5a6f2f73cd8: Pull complete
a6d0f7688756: Pull complete
53e16f6135a5: Pull complete
f52b0cc4e76a: Pull complete
e841feee049e: Pull complete
ccf45e5191d0: Pull complete
Digest: sha256:bf65ecee69c43e52d0e065d094fbdfe4df6e408d47a96e56c7a29caaf31d3c35
Status: Downloaded newer image for redis:latest
Could not connect to Redis at 127.0.0.1:32355: Connection refused
not connected&gt;
</code></pre>

<p>我们先来介绍这里的测试方法。 使用 Docker 的 Redis 官方镜像， <code>--net host</code> 是使用宿主机网络； <code>--rm</code> 表示停止完后即清除； <code>-it</code> 分别表示获取输入及获取 TTY。</p>

<p>通过以上测试发现不能正常连接，故而说明 <code>Service</code> 还是未配置好。使用前面提到的方法也可以进行排查，不过这里提供另一种排查这类问题的思路。</p>

<pre><code>master $ kubectl -n work get endpoints
NAME           ENDPOINTS        AGE
saythx-redis   10.32.0.4:6380   9m
</code></pre>

<p>通过之前的章节，我们已经知道 <code>Service</code> 工作的时候是按 <code>Endpoints</code> 来的，这里我们发现此处的 <code>Endpoints</code> 是 <code>6380</code> 与我们预期的 <code>6379</code> 并不相同。所以问题定位于端口配置有误。</p>

<p>前面已经说过修正方法了，不再赘述。当修正完成后，再次验证：</p>

<pre><code>master $ kubectl -n work get endpoints
NAME           ENDPOINTS        AGE
saythx-redis   10.32.0.4:6379   15m
</code></pre>

<p><code>Endpoints</code> 已经正常，验证下服务是否可用：</p>

<pre><code>master $ docker run --rm -it --net host redis redis-cli -p 32355
127.0.0.1:32355&gt; ping
PONG
</code></pre>

<p>验证无误。</p>

<h2 id="集群问题">集群问题</h2>

<p>由于我们有多个节点，况且在集群搭建和维护过程中，也会比较常见到集群相关的问题。这里我们先举个实际例子进行分析：</p>

<pre><code>master $ kubectl get nodes
NAME      STATUS     ROLES     AGE       VERSION
master    Ready      master    58m       v1.11.3
node01    NotReady   &lt;none&gt;    58m       v1.11.3
</code></pre>

<p>通过 kubectl 查看，发现有一个节点 NotReady ，这在搭建集群的过程中也有可能遇到。</p>

<pre><code>master $ kubectl get  node/node01 -o yaml
apiVersion: v1
kind: Node
metadata:
  annotations:
    kubeadm.alpha.kubernetes.io/cri-socket: /var/run/dockershim.sock
    node.alpha.kubernetes.io/ttl: &quot;0&quot;
    volumes.kubernetes.io/controller-managed-attach-detach: &quot;true&quot;
  creationTimestamp: 2018-12-19T16:46:59Z
  labels:
    beta.kubernetes.io/arch: amd64
    beta.kubernetes.io/os: linux
    kubernetes.io/hostname: node01
  name: node01
  resourceVersion: &quot;4850&quot;
  selfLink: /api/v1/nodes/node01
  uid: b440d3d5-03ad-11e9-917e-0242ac110035
spec: {}
status:
  addresses:
  - address: 172.17.0.66
    type: InternalIP
  - address: node01
    type: Hostname
  allocatable:
    cpu: &quot;4&quot;
    ephemeral-storage: &quot;89032026784&quot;
    hugepages-1Gi: &quot;0&quot;
    hugepages-2Mi: &quot;0&quot;
    memory: 3894652Ki
    pods: &quot;110&quot;
  capacity:
    cpu: &quot;4&quot;
    ephemeral-storage: 96605932Ki
    hugepages-1Gi: &quot;0&quot;
    hugepages-2Mi: &quot;0&quot;
    memory: 3997052Ki
    pods: &quot;110&quot;
  conditions:
  - lastHeartbeatTime: 2018-12-19T17:42:16Z
    lastTransitionTime: 2018-12-19T17:43:00Z
    message: Kubelet stopped posting node status.
    reason: NodeStatusUnknown
    status: Unknown
    type: OutOfDisk
  - lastHeartbeatTime: 2018-12-19T17:42:16Z
    lastTransitionTime: 2018-12-19T17:43:00Z
    message: Kubelet stopped posting node status.
    reason: NodeStatusUnknown
    status: Unknown
    type: MemoryPressure
  - lastHeartbeatTime: 2018-12-19T17:42:16Z
    lastTransitionTime: 2018-12-19T17:43:00Z
    message: Kubelet stopped posting node status.
    reason: NodeStatusUnknown
    status: Unknown
    type: DiskPressure
  - lastHeartbeatTime: 2018-12-19T17:42:16Z
    lastTransitionTime: 2018-12-19T16:46:59Z
    message: kubelet has sufficient PID available
    reason: KubeletHasSufficientPID
    status: &quot;False&quot;
    type: PIDPressure
  - lastHeartbeatTime: 2018-12-19T17:42:16Z
    lastTransitionTime: 2018-12-19T17:43:00Z
    message: Kubelet stopped posting node status.
    reason: NodeStatusUnknown
    status: Unknown
    type: Ready
  daemonEndpoints:
    kubeletEndpoint:
      Port: 10250
  ...
</code></pre>

<p>我们之前介绍 <code>kubelet</code> 时说过， <code>kubelet</code> 的作用之一便是将自身注册至 <code>kube-apiserver</code>。</p>

<p>这里的 message 信息说明 <code>kubelet</code> 不再向 <code>kube-apiserver</code> 发送心跳包之类的了，所以被判定为 NotReady 的状态。</p>

<p>接下来，我们登录 node01 机器查看 <code>kubelet</code> 的状态。</p>

<pre><code>node01 $ systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: enabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─kubeadm.conf
   Active: inactive (dead) since Wed 2018-12-19 17:42:17 UTC; 18min ago
     Docs: https://kubernetes.io/docs/home/
  Process: 1693 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_CONFIG_ARGS $KUBELET_
 Main PID: 1693 (code=exited, status=0/SUCCESS)
</code></pre>

<p>可以看到该机器上 <code>kubelet</code> 没有启动。现在将其启动，稍等片刻看看节群中 <code>Node</code> 的状态。</p>

<pre><code>master $ kubectl get nodes
NAME      STATUS    ROLES     AGE       VERSION
master    Ready     master    1h        v1.11.3
node01    Ready     &lt;none&gt;    1h        v1.11.3
</code></pre>

<h2 id="总结">总结</h2>

<p>本节我们介绍了 K8S 中常用的问题排查和解决思路，但实际生产环境中情况会有和更多不确定因素，掌握本节中介绍的基础，有利于之后生产环境中进行常规问题的排查。</p>

<p>当然，本节只是介绍通过 kubectl 来定位和解决问题，个别情况下我们需要登录相关的节点，实际去使用 <code>Docker</code> 工具等进行问题的详细排查。</p>

<p>至此，K8S 的基础原理和常规问题排查思路等都已经通过包括本节在内的 19 小节介绍完毕，相信你现在已经迫不及待的想要使用 K8S 了。</p>

<p>不过 kubectl 作为命令行工具也许有些人会不习惯使用，下节，我们将介绍 K8S 的扩展组件 <code>kube-dashboard</code> 了解它的主要功能及带给我们的便利。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#f79b9b9bcec3c6c6c7c0b7909a969e9bd994989a" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f104faaafd963c6',t:'MTczNDAzMzczNy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>