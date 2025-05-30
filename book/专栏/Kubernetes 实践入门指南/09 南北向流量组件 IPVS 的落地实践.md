<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;南北向流量组件&#32;IPVS&#32;的落地实践>
        <link rel="icon" href="/static/favicon.png">
        <title>09 南北向流量组件 IPVS 的落地实践 </title>
        
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
                            <h1 id="title" data-id="09 南北向流量组件 IPVS 的落地实践" class="title">09 南北向流量组件 IPVS 的落地实践</h1>
                            <div><p>我们知道 Kubernetes 工作节点的流量管理都是由 kube-proxy 来管理的。kube-proxy 利用了 iptables 的网络流量转换能力，在整个集群的数据层面创建了一层集群虚拟网络，也就是大家在 Service 对象中会看到的术语 ClusterIP，即集群网络 IP。既然 iptables 已经很完美的支持流量的负载均衡，并能实现南北向流量的反向代理功能，为什么我们还要让用户使用另外一个系统组件 IPVS 来代替它呢？</p>

<p>主要原因还是 iptables 能承载的 Service 对象规模有限，超过 1000 个以上就开始出现性能瓶颈了。目前 Kubernetes 默认推荐代理就是 IPVS 模式，这个推荐方案迫使我们需要开始了解 IPVS 的机制，熟悉它的应用范围和对比 iptables 的优缺点，让我们能有更多的精力放在应用开发上。</p>

<h3 id="一次大规模的-service-性能评测引入的-ipvs">一次大规模的 Service 性能评测引入的 IPVS</h3>

<p>iptables 一直是 Kubernetes 集群依赖的系统组件，它同时也是 Liinux 的内核模块，一般实践过程中我们都不会感知到它的性能问题。社区中有华为的开发者在 KuberCon 2018 中引入了一个问题：</p>

<blockquote>
<p>在超大规模如 10000 个 Service 的场景下，kube-proxy 的南北向流量转发性能还能保持高效吗？</p>
</blockquote>

<p>通过测试数据发现，答案是否定的。在 Pod 实例规模达到上万个实例的时候，iptables 就开始对系统性能产生影响了。我们需要知道哪些原因导致 iptables 不能稳定工作。</p>

<p>首先，IPVS 模式 和 iptables 模式同样基于 Netfilter，在生成负载均衡规则的时候，IPVS 是基于哈希表转发流量，iptables 则采用遍历一条一条规则来转发，因为 iptables 匹配规则需要从上到下一条一条规则的匹配，肯定对 CPU 消耗增大并且转发效率随着规则规模的扩大而降低。反观 IPVS 的哈希查表方案，在生成 Service 负载规则后，查表范围有限，所以转发性能上直接秒杀了 iptables 模式。</p>

<p>其次，这里我们要清楚的是，iptables 毕竟是为防火墙模型配置的工具，和专业的负载均衡组件 IPVS 的实现目标不同，这里并不能说 IPVS 就比 iptables 优秀，因为 IPVS 模式启用之后，仅仅只是取代南北向流量的转发，东西向流量的 NAT 转换仍然需要 iptables 来支撑。为了让大家对它们性能对比的影响有一个比较充分的理解，可以看下图：</p>

<p><img src="assets/d1b69cc0-dad0-11ea-a4e8-99c9f3b124a4.jpg" alt="ipvs-iptables-compare" /></p>

<p>从图中可以看到，当 Service 对象实例超过 1000 的时候，iptables 和 IPVS 对 CPU 的影响才会产生差异，规模越大影响也越明显。很明显这是因为它们的转发规则的查询方式不同导致了性能的差异。</p>

<p>除了规则匹配的检索优势，IPVS 对比 iptables 还提供了一些更灵活的负载均衡算法特性如：</p>

<ul>
<li>rr：轮询调度</li>
<li>lc：最小连接数</li>
<li>dh：目标地址散列</li>
<li>sh：源地址散列</li>
<li>sed：最短期望延迟</li>
<li>nq：最少队列</li>
</ul>

<p>而在 iptables 中我们默认只有轮询调度特性。</p>

<p>此刻让我们来重温一下 IPVS（IP Virtual Server）的定义，它是构建在 Linux Netfilter 模块之上的实现南北向流量负载均衡的 Linux 内核模块。IPVS 支持 4 层地址请求的转发，把 Service 层的唯一虚拟 IP 绑定到容器副本 IP 组上，实现流量负载。所以它天然符合 Kubernetes 对 Service 的定义实现。</p>

<h3 id="ipvs-使用简介">IPVS 使用简介</h3>

<p>安装 ipvsadm，这是 LVS 的管理工具：</p>

<pre><code class="language-bash">sudo apt-get install -y ipvsadm

</code></pre>

<p>创建虚拟服务：</p>

<pre><code class="language-bash">sudo ipvsadm -A -t 100.100.100.100:80 -s rr

</code></pre>

<p>使用容器创建 2 个实例：</p>

<pre><code class="language-bash">$ docker run -d -p 8000:8000 --name first -t jwilder/whoami
cd977829ae0c76236a1506c497d5ce1628f1f701f8ed074916b21fc286f3d0d1

$ docker run -d -p 8001:8000 --name second -t jwilder/whoami
5886b1ed7bd4095cb02b32d1642866095e6f4ce1750276bd9fc07e91e2fbc668

</code></pre>

<p>查出容器地址：</p>

<pre><code class="language-bash">$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' first
172.17.0.2

$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' second
172.17.0.3

$ curl 172.17.0.2:8000
I'm cd977829ae0c

</code></pre>

<p>配置 IP 组绑定到虚拟服务 IP 上：</p>

<pre><code>$ sudo ipvsadm -a -t 100.100.100.100:80 -r 172.17.0.2:8000 -m
$ sudo ipvsadm -a -t 100.100.100.100:80 -r 172.17.0.3:8000 -m

$ ipvsadm -l
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -&gt; RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  100.100.100.100:http rr
  -&gt; 172.17.0.2:8000              Masq    1      0          0
  -&gt; 172.17.0.3:8000              Masq    1      0          0

</code></pre>

<p>到此，IPVS 的负载配置步骤就给大家演示完了，kube-proxy 也是如此操作 IPVS 的。</p>

<h3 id="kube-proxy-的-ipvs-模式配置参数详解">kube-proxy 的 IPVS 模式配置参数详解</h3>

<ul>
<li><strong>&ndash;proxy-mode 参数</strong>：当你配置为 <code>--proxy-mode=ipvs</code>，立即激活 IPVS NAT 转发模式，实现 POD 端口的流量负载。</li>
<li><strong>&ndash;ipvs-scheduler 参数</strong>：修改负载均衡策略，默认为 rr——轮询调度。</li>
<li><strong>&ndash;cleanup-ipvs 参数</strong>：启动 IPVS 前清除历史遗留的规则。</li>
<li><strong>&ndash;ipvs-sync-period 和 &ndash;ipvs-min-sync-period 参数</strong>：配置定期同步规则的周期，例如 5s，必须大于 0。</li>
<li><strong>&ndash;ipvs-exclude-cidrs 参数</strong>：过滤自建的 IPVS 规则，让你可以保留之前的负载均衡流量。</li>
</ul>

<h3 id="ipvs-模式下的例外情况">IPVS 模式下的例外情况</h3>

<p>IPVS 毕竟是为了负载均衡流量使用的，还有一些场景下它是爱莫能助的。比如包过滤、端口回流、SNAT 等需求下还是要依靠 iptables 来完成。另外，还有 4 种情况下 IPVS 模式会回退到 iptables 模式。</p>

<ul>
<li>kube-proxy 启动并开启 <code>--masquerade-all=true</code> 参数</li>
<li>kube-proxy 启动时定义了集群网络</li>
<li>支持 Loadbalancer 类型的 Service</li>
<li>支持 NodePort 类型的 Service</li>
</ul>

<p>当然，为了优化 iptables 的过多规则生成，kube-proxy 还引入 ipset 工具来减少 iptables 规则。以下表格就是 IPVS 模式下维护的 ipset 规则：</p>

<table>
<thead>
<tr>
<th align="left"><strong>set name</strong></th>
<th align="left"><strong>members</strong></th>
<th align="left"><strong>usage</strong></th>
</tr>
</thead>

<tbody>
<tr>
<td align="left">KUBE-CLUSTER-IP</td>
<td align="left">All Service IP + port</td>
<td align="left">masquerade for cases that <code>masquerade-all=true</code> or <code>clusterCIDR</code> specified</td>
</tr>

<tr>
<td align="left">KUBE-LOOP-BACK</td>
<td align="left">All Service IP + port + IP</td>
<td align="left">masquerade for resolving hairpin issue</td>
</tr>

<tr>
<td align="left">KUBE-EXTERNAL-IP</td>
<td align="left">Service External IP + port</td>
<td align="left">masquerade for packets to external IPs</td>
</tr>

<tr>
<td align="left">KUBE-LOAD-BALANCER</td>
<td align="left">Load Balancer ingress IP + port</td>
<td align="left">masquerade for packets to Load Balancer type service</td>
</tr>

<tr>
<td align="left">KUBE-LOAD-BALANCER-LOCAL</td>
<td align="left">Load Balancer ingress IP + port with <code>externalTrafficPolicy=local</code></td>
<td align="left">accept packets to Load Balancer with <code>externalTrafficPolicy=local</code></td>
</tr>

<tr>
<td align="left">KUBE-LOAD-BALANCER-FW</td>
<td align="left">Load Balancer ingress IP + port with <code>loadBalancerSourceRanges</code></td>
<td align="left">Drop packets for Load Balancer type Service with <code>loadBalancerSourceRanges</code> specified</td>
</tr>

<tr>
<td align="left">KUBE-LOAD-BALANCER-SOURCE-CIDR</td>
<td align="left">Load Balancer ingress IP + port + source CIDR</td>
<td align="left">accept packets for Load Balancer type Service with <code>loadBalancerSourceRanges</code> specified</td>
</tr>

<tr>
<td align="left">KUBE-NODE-PORT-TCP</td>
<td align="left">NodePort type Service TCP port</td>
<td align="left">masquerade for packets to NodePort(TCP)</td>
</tr>

<tr>
<td align="left">KUBE-NODE-PORT-LOCAL-TCP</td>
<td align="left">NodePort type Service TCP port with <code>externalTrafficPolicy=local</code></td>
<td align="left">accept packets to NodePort Service with <code>externalTrafficPolicy=local</code></td>
</tr>

<tr>
<td align="left">KUBE-NODE-PORT-UDP</td>
<td align="left">NodePort type Service UDP port</td>
<td align="left">masquerade for packets to NodePort(UDP)</td>
</tr>

<tr>
<td align="left">KUBE-NODE-PORT-LOCAL-UDP</td>
<td align="left">NodePort type service UDP port with <code>externalTrafficPolicy=local</code></td>
<td align="left">accept packets to NodePort Service with <code>externalTrafficPolicy=local</code></td>
</tr>
</tbody>
</table>

<h3 id="使用-ipvs-的注意事项">使用 IPVS 的注意事项</h3>

<p>Kubernetes 使用 IPVS 模式可以覆盖大部分场景下的流量负载，但是对于长连接 TCP 请求的水平扩展分流是无能为力的。这是因为 IPVS 并没有能力对 keepalive_requests 做一些限制。一旦你遇到这样的场景，临时解决办法是把连接方式从长连接变为短连接。如设定请求值（比如 1000）之后服务端会在 HTTP 的 Header 头标记 <code>Connection:close</code>，通知客户端处理完当前的请求后关闭连接，新的请求需要重新建立 TCP 连接，所以这个过程中不会出现请求失败，同时又达到了将长连接按需转换为短连接的目的。当然长期的解决之道，你需要在集群前置部署一组 Nginx 或 HAProxy 集群，有效帮助你限制长连接请求的阈值，从而轻松实现流量的弹性扩容。</p>

<h3 id="实践总结">实践总结</h3>

<p>IPVS 模式的引入是社区进行高性能集群测试而引入的优化方案，通过内核已经存在的 IPVS 模块替换掉 iptables 的负载均衡实现，可以说是一次非常成功的最佳实践。因为 IPVS 内置在 Kernel 中，其实 Kernel 的版本对 IPVS 还是有很大影响的，在使用中一定需要注意。当笔者在揣摩着 IPVS 和 iptables 配合使用的纠结中，其实 Linux 社区已经在推进一个新技术 eBPF（柏克莱封包过滤器）技术，准备一举取代 iptables 和 IPVS。如果你没有听说过这个技术，你一定看过这个 Cilium 这个容器网络解决方案，它就是基于 eBPF 实现的：</p>

<p><img src="assets/8bb729b0-d661-11ea-8e3d-27a3708e9ea4.jpg" alt="8-2-cilium-diagram" /></p>

<p>通过 eBPF 技术，目前已经在高版本的 Kernel 之上实现了流量转发和容器网络互连，期待有一天可以完美替换 iptables 和 IPVS，为我们提供功能更强、性能更好的流量管理组件。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#026e6e6e3b363333323542656f636b6e2c616d6f" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1052e9195063c6',t:'MTczNDAzMzg3MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>