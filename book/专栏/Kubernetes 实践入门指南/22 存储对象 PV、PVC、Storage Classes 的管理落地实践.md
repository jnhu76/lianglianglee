<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=22&#32;存储对象&#32;PV、PVC、Storage&#32;Classes&#32;的管理落地实践>
        <link rel="icon" href="/static/favicon.png">
        <title>22 存储对象 PV、PVC、Storage Classes 的管理落地实践 </title>
        
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
                        <a class="menu-item" id="00 为什么我们要学习 Kubernetes 技术.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/00%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e4%bb%ac%e8%a6%81%e5%ad%a6%e4%b9%a0%20Kubernetes%20%e6%8a%80%e6%9c%af.md">00 为什么我们要学习 Kubernetes 技术.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 重新认识 Kubernetes 的核心组件.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/01%20%e9%87%8d%e6%96%b0%e8%ae%a4%e8%af%86%20Kubernetes%20%e7%9a%84%e6%a0%b8%e5%bf%83%e7%bb%84%e4%bb%b6.md">01 重新认识 Kubernetes 的核心组件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 深入理解 Kubernets 的编排对象.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/02%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%20Kubernets%20%e7%9a%84%e7%bc%96%e6%8e%92%e5%af%b9%e8%b1%a1.md">02 深入理解 Kubernets 的编排对象.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 DevOps 场景下落地 K8s 的困难分析.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/03%20DevOps%20%e5%9c%ba%e6%99%af%e4%b8%8b%e8%90%bd%e5%9c%b0%20K8s%20%e7%9a%84%e5%9b%b0%e9%9a%be%e5%88%86%e6%9e%90.md">03 DevOps 场景下落地 K8s 的困难分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 微服务应用场景下落地 K8s 的困难分析.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/04%20%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af%e4%b8%8b%e8%90%bd%e5%9c%b0%20K8s%20%e7%9a%84%e5%9b%b0%e9%9a%be%e5%88%86%e6%9e%90.md">04 微服务应用场景下落地 K8s 的困难分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 解决 K8s 落地难题的方法论提炼.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/05%20%e8%a7%a3%e5%86%b3%20K8s%20%e8%90%bd%e5%9c%b0%e9%9a%be%e9%a2%98%e7%9a%84%e6%96%b9%e6%b3%95%e8%ae%ba%e6%8f%90%e7%82%bc.md">05 解决 K8s 落地难题的方法论提炼.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 练习篇：K8s 核心实践知识掌握.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/06%20%e7%bb%83%e4%b9%a0%e7%af%87%ef%bc%9aK8s%20%e6%a0%b8%e5%bf%83%e5%ae%9e%e8%b7%b5%e7%9f%a5%e8%af%86%e6%8e%8c%e6%8f%a1.md">06 练习篇：K8s 核心实践知识掌握.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 容器引擎 containerd 落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/07%20%e5%ae%b9%e5%99%a8%e5%bc%95%e6%93%8e%20containerd%20%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">07 容器引擎 containerd 落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 K8s 集群安装工具 kubeadm 的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/08%20K8s%20%e9%9b%86%e7%be%a4%e5%ae%89%e8%a3%85%e5%b7%a5%e5%85%b7%20kubeadm%20%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">08 K8s 集群安装工具 kubeadm 的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 南北向流量组件 IPVS 的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/09%20%e5%8d%97%e5%8c%97%e5%90%91%e6%b5%81%e9%87%8f%e7%bb%84%e4%bb%b6%20IPVS%20%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">09 南北向流量组件 IPVS 的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 东西向流量组件 Calico 的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/10%20%e4%b8%9c%e8%a5%bf%e5%90%91%e6%b5%81%e9%87%8f%e7%bb%84%e4%bb%b6%20Calico%20%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">10 东西向流量组件 Calico 的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 服务发现 DNS 的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/11%20%e6%9c%8d%e5%8a%a1%e5%8f%91%e7%8e%b0%20DNS%20%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">11 服务发现 DNS 的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 练习篇：K8s 集群配置测验.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/12%20%e7%bb%83%e4%b9%a0%e7%af%87%ef%bc%9aK8s%20%e9%9b%86%e7%be%a4%e9%85%8d%e7%bd%ae%e6%b5%8b%e9%aa%8c.md">12 练习篇：K8s 集群配置测验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 理解对方暴露服务的对象 Ingress 和 Service.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/13%20%e7%90%86%e8%a7%a3%e5%af%b9%e6%96%b9%e6%9a%b4%e9%9c%b2%e6%9c%8d%e5%8a%a1%e7%9a%84%e5%af%b9%e8%b1%a1%20Ingress%20%e5%92%8c%20Service.md">13 理解对方暴露服务的对象 Ingress 和 Service.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 应用网关 OpenResty 对接 K8s 实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/14%20%e5%ba%94%e7%94%a8%e7%bd%91%e5%85%b3%20OpenResty%20%e5%af%b9%e6%8e%a5%20K8s%20%e5%ae%9e%e8%b7%b5.md">14 应用网关 OpenResty 对接 K8s 实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Service 层引流技术实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/15%20Service%20%e5%b1%82%e5%bc%95%e6%b5%81%e6%8a%80%e6%9c%af%e5%ae%9e%e8%b7%b5.md">15 Service 层引流技术实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 Cilium 容器网络的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/16%20Cilium%20%e5%ae%b9%e5%99%a8%e7%bd%91%e7%bb%9c%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">16 Cilium 容器网络的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 应用流量的优雅无损切换实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/17%20%e5%ba%94%e7%94%a8%e6%b5%81%e9%87%8f%e7%9a%84%e4%bc%98%e9%9b%85%e6%97%a0%e6%8d%9f%e5%88%87%e6%8d%a2%e5%ae%9e%e8%b7%b5.md">17 应用流量的优雅无损切换实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 练习篇：应用流量无损切换技术测验.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/18%20%e7%bb%83%e4%b9%a0%e7%af%87%ef%bc%9a%e5%ba%94%e7%94%a8%e6%b5%81%e9%87%8f%e6%97%a0%e6%8d%9f%e5%88%87%e6%8d%a2%e6%8a%80%e6%9c%af%e6%b5%8b%e9%aa%8c.md">18 练习篇：应用流量无损切换技术测验.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 使用 Rook 构建生产可用存储环境实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/19%20%e4%bd%bf%e7%94%a8%20Rook%20%e6%9e%84%e5%bb%ba%e7%94%9f%e4%ba%a7%e5%8f%af%e7%94%a8%e5%ad%98%e5%82%a8%e7%8e%af%e5%a2%83%e5%ae%9e%e8%b7%b5.md">19 使用 Rook 构建生产可用存储环境实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 有状态应用的默认特性落地分析.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/20%20%e6%9c%89%e7%8a%b6%e6%80%81%e5%ba%94%e7%94%a8%e7%9a%84%e9%bb%98%e8%ae%a4%e7%89%b9%e6%80%a7%e8%90%bd%e5%9c%b0%e5%88%86%e6%9e%90.md">20 有状态应用的默认特性落地分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 案例：分布式 MySQL 集群工具 Vitess 实践分析.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/21%20%e6%a1%88%e4%be%8b%ef%bc%9a%e5%88%86%e5%b8%83%e5%bc%8f%20MySQL%20%e9%9b%86%e7%be%a4%e5%b7%a5%e5%85%b7%20Vitess%20%e5%ae%9e%e8%b7%b5%e5%88%86%e6%9e%90.md">21 案例：分布式 MySQL 集群工具 Vitess 实践分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 存储对象 PV、PVC、Storage Classes 的管理落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/22%20%e5%ad%98%e5%82%a8%e5%af%b9%e8%b1%a1%20PV%e3%80%81PVC%e3%80%81Storage%20Classes%20%e7%9a%84%e7%ae%a1%e7%90%86%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">22 存储对象 PV、PVC、Storage Classes 的管理落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 K8s 集群中存储对象灾备的落地实践.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/23%20K8s%20%e9%9b%86%e7%be%a4%e4%b8%ad%e5%ad%98%e5%82%a8%e5%af%b9%e8%b1%a1%e7%81%be%e5%a4%87%e7%9a%84%e8%90%bd%e5%9c%b0%e5%ae%9e%e8%b7%b5.md">23 K8s 集群中存储对象灾备的落地实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 练习篇：K8s 集群配置测验.md" href="/%e4%b8%93%e6%a0%8f/Kubernetes%20%e5%ae%9e%e8%b7%b5%e5%85%a5%e9%97%a8%e6%8c%87%e5%8d%97/24%20%e7%bb%83%e4%b9%a0%e7%af%87%ef%bc%9aK8s%20%e9%9b%86%e7%be%a4%e9%85%8d%e7%bd%ae%e6%b5%8b%e9%aa%8c.md">24 练习篇：K8s 集群配置测验.md</a>
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
                            <h1 id="title" data-id="22 存储对象 PV、PVC、Storage Classes 的管理落地实践" class="title">22 存储对象 PV、PVC、Storage Classes 的管理落地实践</h1>
                            <div><p>谈到 Kubernetes 存储对象的管理，大多数读者使用最多的就是 Local、NFS 存储类型。因为基于本地卷的挂载使用很少出现问题，并不会出现有什么困难的场景需要用心学习整理。但是从我这里出发想带领读者一起，往更深层的对象实现细节和云原生的存储运维角度出发，看看我们能怎么管理这些资源才是落地的实践。</p>

<h3 id="了解-pv-pvc-storageclass">了解 PV、PVC、StorageClass</h3>

<p>StorageClass 是描述存储类的方法。 不同的类型可能会映射到不同的服务质量等级或备份策略，或是由集群管理员制定的任意策略。 Kubernetes 本身并不清楚各种类代表的是什么。这个类的概念在其他存储系统中有时被称为“配置文件”。</p>

<p>每个 StorageClass 都包含 provisioner、parameters 和 reclaimPolicy 字段，这些字段会在 StorageClass 需要动态分配 PersistentVolume 时会使用到。</p>

<p>StorageClass 对象的命名很重要，用户使用这个命名来请求生成一个特定的类。当创建 StorageClass 对象时，管理员设置 StorageClass 对象的命名和其他参数，一旦创建了对象就不能再对其更新。参考范例如下：</p>

<pre><code class="language-yaml">apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
allowVolumeExpansion: true
mountOptions:
  - debug
volumeBindingMode: Immediate

</code></pre>

<p>持久卷（PersistentVolume，PV）是集群中的一块存储，可以由管理员事先供应，或者使用存储类（StorageClass）来动态供应。 持久卷是全局集群资源，就像 Node 也是全局集群资源一样，没有 Namespace 隔离的概念。PV 持久卷和普通的 Volumes 一样，也是使用卷插件来实现的，只是它们拥有自己独立的生命周期。 此 API 对象中记述了存储的实现细节，无论其背后是 NFS、iSCSI 还是特定于云平台的存储系统。</p>

<p>持久卷申领（PersistentVolumeClaim，PVC）表达的是用户对存储的请求。概念上与 Pod 类似。 Pod 会耗用节点资源，而 PVC 申领会消耗 PV 资源。Pod 可以请求特定数量的资源（CPU 和内存）；同样 PVC 申领也可以请求特定 PV 的大小和访问模式。</p>

<p>尽管 PersistentVolumeClaim 允许用户消耗抽象的存储资源，常见的情况是针对不同的用户需要提供具有不同属性（如性能）的 PersistentVolume 卷。 集群管理员需要能够提供不同性质的 PersistentVolume，并且这些 PV 卷之间的差别不仅限于卷大小和访问模式，同时又不能将卷是如何实现的这些细节暴露给用户。 为了满足这类需求，就有了存储类（StorageClass）资源。</p>

<p>所以总结下来，对于存储资源，我们默认指代的就是铁三角 API 对象：StorageClass、PersistentVolume、PersistentVolumeClaim。</p>

<h3 id="了解-csi">了解 CSI</h3>

<p>从 Kubernetes v1.13 开始 CSI 进入稳定可用阶段，所以用户有必要了解这个容器存储接口。CSI 卷类型是一种外部引用驱动的 CSI 卷插件，用于 Pod 与在同一节点上运行的外部 CSI 卷驱动程序交互。部署 CSI 兼容卷驱动后，用户可以使用 CSI 作为卷类型来挂载驱动提供的存储。</p>

<p><img src="assets/c1e6bc40-2ce1-11eb-90f6-fbd19bda6e6e" alt="19-1-csi-arch" /></p>

<p>一直以来，存储插件的测试、维护等事宜都由 Kubernetes 社区来完成，即使有贡献者提供协作也不容易合并到主分支发布。另外，存储插件需要随 Kubernetes 一同发布，如果存储插件存在问题有可能会影响 Kubernetes 其他组件的正常运行。</p>

<p>鉴于此，Kubernetes 和 CNCF 决定把容器存储进行抽象，通过标准接口的形式把存储部分移到容器编排系统外部去。CSI 的设计目的是定义一个行业标准，该标准将使存储供应商能够自己实现，维护和部署他们的存储插件。这些存储插件会以 Sidecar Container 形式运行在 Kubernetes 上并为容器平台提供稳定的存储服务。</p>

<p>如上 CSI 设计图：<strong>浅绿色</strong>表示从 Kubernetes 社区中抽离出来且可复用的组件，负责连接 CSI 插件（右侧）以及和 Kubernetes 集群交互：</p>

<ul>
<li>Driver-registrar：使用 Kubelet 注册 CSI 驱动程序的 sidecar 容器，并将 NodeId （通过 GetNodeID 调用检索到 CSI endpoint）添加到 Kubernetes Node API 对象的 annotation 里面。</li>
<li>External-provisioner：监听 Kubernetes PersistentVolumeClaim 对象的 sidecar 容器，并触发对 CSI 端点的 CreateVolume 和 DeleteVolume 操作；</li>
<li>External-attacher：可监听 Kubernetes VolumeAttachment 对象并触发 ControllerPublish 和 ControllerUnPublish 操作的 sidecar 容器，负责 attache/detache 卷到 Node 节点上。</li>
</ul>

<p>右侧<strong>浅灰色</strong>表示第三方实现的存储插件驱动，分别有三个服务：</p>

<ul>
<li>CSI identify：标志插件服务，并维持插件健康状态</li>
<li>CSI Controller：创建/删除、attaching/detaching、快照等</li>
<li>CSI Node：attach/mount、umount/detach</li>
</ul>

<p>通过对比 Kubernetes 的内置 Volume Plugin，以及外置 Provisioner 和 CSI 三种方式，在对接比较常见的存储时，可以使用不需要改动的内置方案，因为开箱即用，但是缺点也非常明显，只支持有限的存储类型，可拓展性较差甚至有版本限制，另外官方宣布以后新特性将不再添加到其中。相比之下，使用 CSI 则可以实现和 Kubernetes 的核心组件解耦，并能支持更多的存储类型和高级特性，因而也是推荐使用的一种供应方式。由于对编排系统而言是非侵入式插件部署，因而更受存储服务商的青睐。</p>

<h3 id="采用-volume-snapshots-备份">采用 Volume Snapshots 备份</h3>

<p>与 API 资源 PersistentVolume 和 PersistentVolumeClaim 用于为用户和管理员提供卷的方式类似，VolumeSnapshotContent 和 VolumeSnapshot API 资源被提供用于为用户和管理员创建卷快照。</p>

<p>VolumeSnapshotContent 是指从群集中由管理员配置的卷中获取的快照。它是集群中的资源，就像 PersistentVolume 是集群资源一样。</p>

<p>VolumeSnapshot 是用户对卷的快照请求。它类似于 PersistentVolumeClaim。</p>

<p>VolumeSnapshotClass 允许您指定属于 VolumeSnapshot 的不同属性。这些属性可能在存储系统上同一个卷的快照中会有所不同，因此不能使用 PersistentVolumeClaim 的同一个 StorageClass 来表达。</p>

<p>卷快照为 Kubernetes 用户提供了一种标准化的备份恢复方法，可以在特定时间点复制卷的内容，而无需创建一个全新的卷。例如，数据库管理员可以通过该功能在执行编辑或删除修改之前备份数据库。</p>

<p>用户在使用该功能时需要注意以下几点。</p>

<ul>
<li>API 对象 VolumeSnapshot、VolumeSnapshotContent 和 VolumeSnapshotClass 是 CRD，不是核心 API 的一部分。 VolumeSnapshot 支持仅适用于 CSI 驱动程序。</li>
<li>作为 VolumeSnapshot 测试版部署过程的一部分，Kubernetes 团队提供了一个部署到控制平面的快照控制器，以及一个名为 csi-snapshotter 的帮助容器，与 CSI 驱动程序一起部署。快照控制器监视 VolumeSnapshot 和 VolumeSnapshotContent 对象，并负责动态供应中 VolumeSnapshotContent 对象的创建和删除。帮助容器 csi-snapshotter 监视 VolumeSnapshotContent 对象，并触发针对 CSI 端点的 CreateSnapshot 和 DeleteSnapshot 操作。</li>
<li>CSI 驱动程序可能已经实现或没有实现卷快照功能。已提供卷快照支持的 CSI 驱动程序可能会使用 csi-snapshotter。详情请参见 CSI 驱动程序文档。</li>
<li>CRD 和 快照控制器的安装是 Kubernetes 发行版的责任。</li>
</ul>

<h4 id="volumesnapshot-和-volumesnapshotcontent-的生命周期"><strong>VolumeSnapshot 和 VolumeSnapshotContent 的生命周期</strong></h4>

<p>VolumeSnapshotContents 是集群中的资源。VolumeSnapshot 是对这些资源的请求。VolumeSnapshotContents 和 VolumeSnapshot 之间的交互遵循这个生命周期。</p>

<p><strong>1. 供应卷快照</strong></p>

<p>有两种方式可以配置快照：预配置或动态配置。</p>

<p><strong>2. 预备</strong></p>

<p>群集管理员会创建一些 VolumeSnapshotContents。它们携带了存储系 统上真实卷照的详细信息，可供集群用户使用。它们存在于 Kubernetes API 中，可供消费。</p>

<p><strong>3. 动态</strong></p>

<p>您可以请求从 PersistentVolumeClaim 动态获取快照，而不是使用预先存在的快照。VolumeSnapshotClass 指定了存储提供商的特定参数，以便在获取快照时使用。</p>

<p><strong>4. 绑定</strong></p>

<p>快照控制器处理 VolumeSnapshot 对象与适当的 VolumeSnapshotContent 对象的绑定，在预供应和动态供应的情况下都是如此。绑定是一个一对一的映射。</p>

<p>在预供应绑定的情况下，VolumeSnapshot 将保持未绑定状态，直到请求的 VolumeSnapshotContent 对象被创建。</p>

<p><strong>5. 作为快照源保护的持久性卷索赔</strong></p>

<p>这个保护的目的是为了确保在使用中的 PersistentVolumeClaim API 对象在快照时不会被从系统中移除（因为这可能导致数据丢失）。</p>

<p>当一个 PersistentVolumeClaim 的快照被取走时，该 PersistentVolumeClaim 是在使用中的。如果您删除了一个正在使用的 PersistentVolumeClaim API 对象作为快照源，PersistentVolumeClaim 对象不会被立即删除。相反，PersistentVolumeClaim 对象的删除会被推迟到快照准备好或中止之后。</p>

<p><strong>6. 删除</strong></p>

<p>删除是通过删除 VolumeSnapshot 对象来触发的，将遵循 DeletionPolicy。如果 DeletionPolicy 是 Delete，那么底层存储快照将和 VolumeSnapshotContent 对象一起被删除。如果 DeletionPolicy 是 Retain，那么底层快照和 VolumeSnapshotContent 都会保留。</p>

<h4 id="volumesnapshots"><strong>VolumeSnapshots</strong></h4>

<p>每个 VolumeSnapshot 包含一个规格和一个状态：</p>

<pre><code class="language-yaml">apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshot
metadata:
  name: new-snapshot-test
spec:
  volumeSnapshotClassName: csi-hostpath-snapclass
  source:
    persistentVolumeClaimName: pvc-test

</code></pre>

<p>persistentVolumeClaimName 是快照的 PersistentVolumeClaim 数据源的名称。动态供应快照时需要该字段。</p>

<p>卷快照可以通过使用属性 volumeSnapshotClassName 指定 VolumeSnapshotClass 的名称来请求特定的类。如果没有设置任何内容，则使用默认的类（如果可用）。</p>

<p>对于预设的快照，您需要指定一个 volumeSnapshotContentName 作为快照的源，如下例所示。对于预置快照，volumeSnapshotContentName 源字段是必需的。</p>

<pre><code class="language-yaml">apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshot
metadata:
  name: test-snapshot
spec:
  source:
    volumeSnapshotContentName: test-content

</code></pre>

<h4 id="volume-snapshot-contents"><strong>Volume Snapshot Contents</strong></h4>

<p>每个 VolumeSnapshotContent 包含一个规格和状态。在动态供应中快照通用控制器会创建 VolumeSnapshotContent 对象。下面是一个例子:</p>

<pre><code class="language-yaml">apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshotContent
metadata:
  name: snapcontent-72d9a349-aacd-42d2-a240-d775650d2455
spec:
  deletionPolicy: Delete
  driver: hostpath.csi.k8s.io
  source:
    volumeHandle: ee0cfb94-f8d4-11e9-b2d8-0242ac110002
  volumeSnapshotClassName: csi-hostpath-snapclass
  volumeSnapshotRef:
    name: new-snapshot-test
    namespace: default
    uid: 72d9a349-aacd-42d2-a240-d775650d2455

</code></pre>

<p>volumeHandle 是在存储后端创建的卷的唯一标识符，由 CSI 驱动程序在卷创建期间返回。动态供应快照时需要该字段。它指定了快照的卷源。</p>

<p>对于预配置的快照，群集管理员负责创建 VolumeSnapshotContent 对象，具体如下:</p>

<pre><code class="language-yaml">apiVersion: snapshot.storage.k8s.io/v1beta1
kind: VolumeSnapshotContent
metadata:
  name: new-snapshot-content-test
spec:
  deletionPolicy: Delete
  driver: hostpath.csi.k8s.io
  source:
    snapshotHandle: 7bdd0de3-aaeb-11e8-9aae-0242ac110002
  volumeSnapshotRef:
    name: new-snapshot-test
    namespace: default

</code></pre>

<p>snapshotHandle 是在存储后端创建的卷快照的唯一标识符。该字段对预置快照是必需的。它指定了该 VolumeSnapshotContent 所代表的存储系统上的 CSI 快照 ID。</p>

<h3 id="从快照中恢复卷">从快照中恢复卷</h3>

<p>您可以通过使用 PersistentVolumeClaim 对象中的 dataSource 字段来提供一个新的卷，并从快照中恢复数据。下面是一个例子：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: restore-pvc
spec:
  storageClassName: csi-hostpath-sc
  dataSource:
    name: new-snapshot-test
    kind: VolumeSnapshot
    apiGroup: snapshot.storage.k8s.io
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi

</code></pre>

<h3 id="总结">总结</h3>

<p>随着有状态应用的广泛采用，有状态的存储对象资源开始被广泛使用，这样 Kubernetes 生态中开始包含了更多对存储资源的管理对象。其中最重要的就是备份和恢复对象。当然目前最新的 API 对象中包含的还是测试阶段的 VolumeSnapshot 的概念，无法直接生产可用。另外注意的是 CSI 卷驱动的外置插件架构设计，目前是生产可用，请不要在使用内置驱动来挂载卷了。根据存储驱动落地的情况，大量的 NFS 存储的管理仍然是当前最重要的部分，大家只需要掌握创建、备份、恢复操作就已经算掌握了 90% 的存储对象使用技能。对于分布式存储 Ceph，我们可以在开发测试环节大量使用，等待 Ceph 驱动的成熟时机就可以大量采用。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#fd919191c4c9cccccdcabd9a909c9491d39e9290" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1055936d6a63c6',t:'MTczNDAzMzk3OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>