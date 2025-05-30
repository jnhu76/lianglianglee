<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=06&#32;RPC实战：剖析gRPC源码，动手实现一个完整的RPC>
        <link rel="icon" href="/static/favicon.png">
        <title>06 RPC实战：剖析gRPC源码，动手实现一个完整的RPC </title>
        
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
                        <a class="menu-item" id="00 开篇词 别老想着怎么用好RPC框架，你得多花时间琢磨原理.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%88%ab%e8%80%81%e6%83%b3%e7%9d%80%e6%80%8e%e4%b9%88%e7%94%a8%e5%a5%bdRPC%e6%a1%86%e6%9e%b6%ef%bc%8c%e4%bd%a0%e5%be%97%e5%a4%9a%e8%8a%b1%e6%97%b6%e9%97%b4%e7%90%a2%e7%a3%a8%e5%8e%9f%e7%90%86.md">00 开篇词 别老想着怎么用好RPC框架，你得多花时间琢磨原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 核心原理：能否画张图解释下RPC的通信流程？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/01%20%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86%ef%bc%9a%e8%83%bd%e5%90%a6%e7%94%bb%e5%bc%a0%e5%9b%be%e8%a7%a3%e9%87%8a%e4%b8%8bRPC%e7%9a%84%e9%80%9a%e4%bf%a1%e6%b5%81%e7%a8%8b%ef%bc%9f.md">01 核心原理：能否画张图解释下RPC的通信流程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 协议：怎么设计可扩展且向后兼容的协议？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/02%20%e5%8d%8f%e8%ae%ae%ef%bc%9a%e6%80%8e%e4%b9%88%e8%ae%be%e8%ae%a1%e5%8f%af%e6%89%a9%e5%b1%95%e4%b8%94%e5%90%91%e5%90%8e%e5%85%bc%e5%ae%b9%e7%9a%84%e5%8d%8f%e8%ae%ae%ef%bc%9f.md">02 协议：怎么设计可扩展且向后兼容的协议？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 序列化：对象怎么在网络中传输？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/03%20%e5%ba%8f%e5%88%97%e5%8c%96%ef%bc%9a%e5%af%b9%e8%b1%a1%e6%80%8e%e4%b9%88%e5%9c%a8%e7%bd%91%e7%bb%9c%e4%b8%ad%e4%bc%a0%e8%be%93%ef%bc%9f.md">03 序列化：对象怎么在网络中传输？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 网络通信：RPC框架在网络通信上更倾向于哪种网络IO模型？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/04%20%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1%ef%bc%9aRPC%e6%a1%86%e6%9e%b6%e5%9c%a8%e7%bd%91%e7%bb%9c%e9%80%9a%e4%bf%a1%e4%b8%8a%e6%9b%b4%e5%80%be%e5%90%91%e4%ba%8e%e5%93%aa%e7%a7%8d%e7%bd%91%e7%bb%9cIO%e6%a8%a1%e5%9e%8b%ef%bc%9f.md">04 网络通信：RPC框架在网络通信上更倾向于哪种网络IO模型？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 动态代理：面向接口编程，屏蔽RPC处理流程.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/05%20%e5%8a%a8%e6%80%81%e4%bb%a3%e7%90%86%ef%bc%9a%e9%9d%a2%e5%90%91%e6%8e%a5%e5%8f%a3%e7%bc%96%e7%a8%8b%ef%bc%8c%e5%b1%8f%e8%94%bdRPC%e5%a4%84%e7%90%86%e6%b5%81%e7%a8%8b.md">05 动态代理：面向接口编程，屏蔽RPC处理流程.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 RPC实战：剖析gRPC源码，动手实现一个完整的RPC.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/06%20RPC%e5%ae%9e%e6%88%98%ef%bc%9a%e5%89%96%e6%9e%90gRPC%e6%ba%90%e7%a0%81%ef%bc%8c%e5%8a%a8%e6%89%8b%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e5%ae%8c%e6%95%b4%e7%9a%84RPC.md">06 RPC实战：剖析gRPC源码，动手实现一个完整的RPC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 架构设计：设计一个灵活的RPC框架.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/07%20%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%9a%e8%ae%be%e8%ae%a1%e4%b8%80%e4%b8%aa%e7%81%b5%e6%b4%bb%e7%9a%84RPC%e6%a1%86%e6%9e%b6.md">07 架构设计：设计一个灵活的RPC框架.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 服务发现：到底是要CP还是AP？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/08%20%e6%9c%8d%e5%8a%a1%e5%8f%91%e7%8e%b0%ef%bc%9a%e5%88%b0%e5%ba%95%e6%98%af%e8%a6%81CP%e8%bf%98%e6%98%afAP%ef%bc%9f.md">08 服务发现：到底是要CP还是AP？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 健康检测：这个节点都挂了，为啥还要疯狂发请求？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/09%20%e5%81%a5%e5%ba%b7%e6%a3%80%e6%b5%8b%ef%bc%9a%e8%bf%99%e4%b8%aa%e8%8a%82%e7%82%b9%e9%83%bd%e6%8c%82%e4%ba%86%ef%bc%8c%e4%b8%ba%e5%95%a5%e8%bf%98%e8%a6%81%e7%96%af%e7%8b%82%e5%8f%91%e8%af%b7%e6%b1%82%ef%bc%9f.md">09 健康检测：这个节点都挂了，为啥还要疯狂发请求？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 路由策略：怎么让请求按照设定的规则发到不同的节点上？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/10%20%e8%b7%af%e7%94%b1%e7%ad%96%e7%95%a5%ef%bc%9a%e6%80%8e%e4%b9%88%e8%ae%a9%e8%af%b7%e6%b1%82%e6%8c%89%e7%85%a7%e8%ae%be%e5%ae%9a%e7%9a%84%e8%a7%84%e5%88%99%e5%8f%91%e5%88%b0%e4%b8%8d%e5%90%8c%e7%9a%84%e8%8a%82%e7%82%b9%e4%b8%8a%ef%bc%9f.md">10 路由策略：怎么让请求按照设定的规则发到不同的节点上？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 负载均衡：节点负载差距这么大，为什么收到的流量还一样？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/11%20%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%ef%bc%9a%e8%8a%82%e7%82%b9%e8%b4%9f%e8%bd%bd%e5%b7%ae%e8%b7%9d%e8%bf%99%e4%b9%88%e5%a4%a7%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e6%94%b6%e5%88%b0%e7%9a%84%e6%b5%81%e9%87%8f%e8%bf%98%e4%b8%80%e6%a0%b7%ef%bc%9f.md">11 负载均衡：节点负载差距这么大，为什么收到的流量还一样？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 异常重试：在约定时间内安全可靠地重试.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/12%20%e5%bc%82%e5%b8%b8%e9%87%8d%e8%af%95%ef%bc%9a%e5%9c%a8%e7%ba%a6%e5%ae%9a%e6%97%b6%e9%97%b4%e5%86%85%e5%ae%89%e5%85%a8%e5%8f%af%e9%9d%a0%e5%9c%b0%e9%87%8d%e8%af%95.md">12 异常重试：在约定时间内安全可靠地重试.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 优雅关闭：如何避免服务停机带来的业务损失？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/13%20%e4%bc%98%e9%9b%85%e5%85%b3%e9%97%ad%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e6%9c%8d%e5%8a%a1%e5%81%9c%e6%9c%ba%e5%b8%a6%e6%9d%a5%e7%9a%84%e4%b8%9a%e5%8a%a1%e6%8d%9f%e5%a4%b1%ef%bc%9f.md">13 优雅关闭：如何避免服务停机带来的业务损失？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 优雅启动：如何避免流量打到没有启动完成的节点？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/14%20%e4%bc%98%e9%9b%85%e5%90%af%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e6%b5%81%e9%87%8f%e6%89%93%e5%88%b0%e6%b2%a1%e6%9c%89%e5%90%af%e5%8a%a8%e5%ae%8c%e6%88%90%e7%9a%84%e8%8a%82%e7%82%b9%ef%bc%9f.md">14 优雅启动：如何避免流量打到没有启动完成的节点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 熔断限流：业务如何实现自我保护_.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/15%20%e7%86%94%e6%96%ad%e9%99%90%e6%b5%81%ef%bc%9a%e4%b8%9a%e5%8a%a1%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e8%87%aa%e6%88%91%e4%bf%9d%e6%8a%a4_.md">15 熔断限流：业务如何实现自我保护_.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 业务分组：如何隔离流量？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/16%20%e4%b8%9a%e5%8a%a1%e5%88%86%e7%bb%84%ef%bc%9a%e5%a6%82%e4%bd%95%e9%9a%94%e7%a6%bb%e6%b5%81%e9%87%8f%ef%bc%9f.md">16 业务分组：如何隔离流量？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 异步RPC：压榨单机吞吐量.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/17%20%e5%bc%82%e6%ad%a5RPC%ef%bc%9a%e5%8e%8b%e6%a6%a8%e5%8d%95%e6%9c%ba%e5%90%9e%e5%90%90%e9%87%8f.md">17 异步RPC：压榨单机吞吐量.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 安全体系：如何建立可靠的安全体系？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/18%20%e5%ae%89%e5%85%a8%e4%bd%93%e7%b3%bb%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bb%ba%e7%ab%8b%e5%8f%af%e9%9d%a0%e7%9a%84%e5%ae%89%e5%85%a8%e4%bd%93%e7%b3%bb%ef%bc%9f.md">18 安全体系：如何建立可靠的安全体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 分布式环境下如何快速定位问题？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/19%20%e5%88%86%e5%b8%83%e5%bc%8f%e7%8e%af%e5%a2%83%e4%b8%8b%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e5%ae%9a%e4%bd%8d%e9%97%ae%e9%a2%98%ef%bc%9f.md">19 分布式环境下如何快速定位问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 详解时钟轮在RPC中的应用.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/20%20%e8%af%a6%e8%a7%a3%e6%97%b6%e9%92%9f%e8%bd%ae%e5%9c%a8RPC%e4%b8%ad%e7%9a%84%e5%ba%94%e7%94%a8.md">20 详解时钟轮在RPC中的应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 流量回放：保障业务技术升级的神器.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/21%20%e6%b5%81%e9%87%8f%e5%9b%9e%e6%94%be%ef%bc%9a%e4%bf%9d%e9%9a%9c%e4%b8%9a%e5%8a%a1%e6%8a%80%e6%9c%af%e5%8d%87%e7%ba%a7%e7%9a%84%e7%a5%9e%e5%99%a8.md">21 流量回放：保障业务技术升级的神器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 动态分组：超高效实现秒级扩缩容.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/22%20%e5%8a%a8%e6%80%81%e5%88%86%e7%bb%84%ef%bc%9a%e8%b6%85%e9%ab%98%e6%95%88%e5%ae%9e%e7%8e%b0%e7%a7%92%e7%ba%a7%e6%89%a9%e7%bc%a9%e5%ae%b9.md">22 动态分组：超高效实现秒级扩缩容.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 如何在没有接口的情况下进行RPC调用？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/23%20%e5%a6%82%e4%bd%95%e5%9c%a8%e6%b2%a1%e6%9c%89%e6%8e%a5%e5%8f%a3%e7%9a%84%e6%83%85%e5%86%b5%e4%b8%8b%e8%bf%9b%e8%a1%8cRPC%e8%b0%83%e7%94%a8%ef%bc%9f.md">23 如何在没有接口的情况下进行RPC调用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 如何在线上环境里兼容多种RPC协议？.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/24%20%e5%a6%82%e4%bd%95%e5%9c%a8%e7%ba%bf%e4%b8%8a%e7%8e%af%e5%a2%83%e9%87%8c%e5%85%bc%e5%ae%b9%e5%a4%9a%e7%a7%8dRPC%e5%8d%8f%e8%ae%ae%ef%bc%9f.md">24 如何在线上环境里兼容多种RPC协议？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 RPC框架代码实例详解.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/%e5%8a%a0%e9%a4%90%20RPC%e6%a1%86%e6%9e%b6%e4%bb%a3%e7%a0%81%e5%ae%9e%e4%be%8b%e8%af%a6%e8%a7%a3.md">加餐 RPC框架代码实例详解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐 谈谈我所经历过的RPC.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/%e5%8a%a0%e9%a4%90%20%e8%b0%88%e8%b0%88%e6%88%91%e6%89%80%e7%bb%8f%e5%8e%86%e8%bf%87%e7%9a%84RPC.md">加餐 谈谈我所经历过的RPC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="答疑课堂 基础篇与进阶篇思考题答案合集.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/%e7%ad%94%e7%96%91%e8%af%be%e5%a0%82%20%e5%9f%ba%e7%a1%80%e7%af%87%e4%b8%8e%e8%bf%9b%e9%98%b6%e7%af%87%e6%80%9d%e8%80%83%e9%a2%98%e7%ad%94%e6%a1%88%e5%90%88%e9%9b%86.md">答疑课堂 基础篇与进阶篇思考题答案合集.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 学会从优秀项目的源代码中挖掘知识.md" href="/%e4%b8%93%e6%a0%8f/RPC%e5%ae%9e%e6%88%98%e4%b8%8e%e6%a0%b8%e5%bf%83%e5%8e%9f%e7%90%86/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%ad%a6%e4%bc%9a%e4%bb%8e%e4%bc%98%e7%a7%80%e9%a1%b9%e7%9b%ae%e7%9a%84%e6%ba%90%e4%bb%a3%e7%a0%81%e4%b8%ad%e6%8c%96%e6%8e%98%e7%9f%a5%e8%af%86.md">结束语 学会从优秀项目的源代码中挖掘知识.md</a>
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
                            <h1 id="title" data-id="06 RPC实战：剖析gRPC源码，动手实现一个完整的RPC" class="title">06 RPC实战：剖析gRPC源码，动手实现一个完整的RPC</h1>
                            <div><p>你好，我是何小锋。上一讲我分享了动态代理，其作用总结起来就是一句话：“我们可以通过动态代理技术，屏蔽 RPC 调用的细节，从而让使用者能够面向接口编程。”</p>

<p>到今天为止，我们已经把 RPC 通信过程中要用到的所有基础知识都讲了一遍，但这些内容多属于理论。<strong>这一讲我们就来实战一下，看看具体落实到代码上，我们应该怎么实现一个 RPC 框架？</strong></p>

<p>为了能让咱们快速达成共识，我选择剖析 gRPC 源码（源码地址：<a href="https://github.com/grpc/grpc-java" target="_blank">https://github.com/grpc/grpc-java</a>）。通过分析 gRPC 的通信过程，我们可以清楚地知道在 gRPC 里面这些知识点是怎么落地到具体代码上的。</p>

<p>gRPC 是由 Google 开发并且开源的一款高性能、跨语言的 RPC 框架，当前支持 C、Java 和 Go 等语言，当前 Java 版本最新 Release 版为 1.27.0。gRPC 有很多特点，比如跨语言，通信协议是基于标准的 HTTP/2 设计的，序列化支持 PB（Protocol Buffer）和 JSON，整个调用示例如下图所示：</p>

<p><img src="assets/07c9f3a1b4ff46ee8d06da9f0ec5c2ff.jpg" alt="" /></p>

<p>如果你想快速地了解一个全新框架的工作原理，我个人认为最快的方式就是从使用示例开始，所以现在我们就以最简单的 HelloWord 为例开始了解。</p>

<p>在这个例子里面，我们会定义一个 say 方法，调用方通过 gRPC 调用服务提供方，然后服务提供方会返回一个字符串给调用方。</p>

<p>为了保证调用方和服务提供方能够正常通信，我们需要先约定一个通信过程中的契约，也就是我们在 Java 里面说的定义一个接口，这个接口里面只会包含一个 say 方法。在 gRPC 里面定义接口是通过写 Protocol Buffer 代码，从而把接口的定义信息通过 Protocol Buffer 语义表达出来。HelloWord 的 Protocol Buffer 代码如下所示：</p>

<pre><code>syntax = &quot;proto3&quot;;

option java_multiple_files = true;
option java_package = &quot;io.grpc.hello&quot;;
option java_outer_classname = &quot;HelloProto&quot;;
option objc_class_prefix = &quot;HLW&quot;;

package hello;

service HelloService{
rpc Say(HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
string name = 1;
}

message HelloReply {
string message = 1;
}
</code></pre>

<p>有了这段代码，我们就可以为客户端和服务器端生成消息对象和 RPC 基础代码。我们可以利用 Protocol Buffer 的编译器 protoc，再配合 gRPC Java 插件（protoc-gen-grpc-java），通过命令行 protoc3 加上 plugin 和 proto 目录地址参数，我们就可以生成消息对象和 gRPC 通信所需要的基础代码。如果你的项目是 Maven 工程的话，你还可以直接选择使用 Maven 插件来生成同样的代码。</p>

<h2 id="发送原理">发送原理</h2>

<p>生成完基础代码以后，我们就可以基于生成的代码写下调用端代码，具体如下：</p>

<pre><code>package io.grpc.hello;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;


import java.util.concurrent.TimeUnit;

public class HelloWorldClient {

    private final ManagedChannel channel;
    private final HelloServiceGrpc.HelloServiceBlockingStub blockingStub;
    /**
    * 构建Channel连接
    **/
    public HelloWorldClient(String host, int port) {
        this(ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build());
    }

    /**
    * 构建Stub用于发请求
    **/
    HelloWorldClient(ManagedChannel channel) {
        this.channel = channel;
        blockingStub = HelloServiceGrpc.newBlockingStub(channel);
    }
    
    /**
    * 调用完手动关闭
    **/
    public void shutdown() throws InterruptedException {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
    }

 
    /**
    * 发送rpc请求
    **/
    public void say(String name) {
        // 构建入参对象
        HelloRequest request = HelloRequest.newBuilder().setName(name).build();
        HelloReply response;
        try {
            // 发送请求
            response = blockingStub.say(request);
        } catch (StatusRuntimeException e) {
            return;
        }
        System.out.println(response);
    }

    public static void main(String[] args) throws Exception {
            HelloWorldClient client = new HelloWorldClient(&quot;127.0.0.1&quot;, 50051);
            try {
                client.say(&quot;world&quot;);
            } finally {
                client.shutdown();
            }
    }
}
</code></pre>

<p><strong>调用端代码大致分成三个步骤：</strong></p>

<ul>
<li>首先用 host 和 port 生成 channel 连接；</li>
<li>然后用前面生成的 HelloService gRPC 创建 Stub 类；</li>
<li>最后我们可以用生成的这个 Stub 调用 say 方法发起真正的 RPC 调用，后续其它的 RPC 通信细节就对我们使用者透明了。</li>
</ul>

<p>为了能看清楚里面具体发生了什么，我们需要进入到 ClientCalls.blockingUnaryCall 方法里面看下逻辑细节。但是为了避免太多的细节影响你理解整体流程，我在下面这张图中只画下了最重要的部分。</p>

<p><img src="assets/78aa551db594491d8d9abbb0ee679629.jpg" alt="" /></p>

<p>我们可以看到，在调用端代码里面，我们只需要一行（第48行）代码就可以发起一个 RPC 调用，而具体这个请求是怎么发送到服务提供者那端的呢？这对于我们 gRPC 使用者来说是完全透明的，我们只要关注是怎么创建出 stub 对象的就可以了。</p>

<p>比如入参是一个字符对象，gRPC 是怎么把这个对象传输到服务提供方的呢？因为在<a href="https://time.geekbang.org/column/article/202779" target="_blank">[第 03 讲]</a> 中我们说过，只有二进制才能在网络中传输，但是目前调用端代码的入参是一个字符对象，那在 gRPC 里面我们是怎么把对象转成二进制数据的呢？</p>

<p>回到上面流程图的第3步，在 writePayload 之前，ClientCallImpl 里面有一行代码就是 method.streamRequest(message)，看方法签名我们大概就知道它是用来把对象转成一个 InputStream，有了 InputStream 我们就很容易获得入参对象的二进制数据。这个方法返回值很有意思，就是为啥不直接返回我们想要的二进制数组，而是返回一个 InputStream 对象呢？你可以先停下来想下原因，我们会在最后继续讨论这个问题。</p>

<p>我们接着看 streamRequest 方法的拥有者 method 是个什么对象？我们可以看到 method 是 MethodDescriptor 对象关联的一个实例，而 MethodDescriptor 是用来存放要调用 RPC 服务的接口名、方法名、服务调用的方式以及请求和响应的序列化和反序列化实现类。</p>

<p>大白话说就是，MethodDescriptor 是用来存储一些 RPC 调用过程中的元数据，而在 MethodDescriptor 里面 requestMarshaller 是在绑定请求的时候用来序列化方式对象的，所以当我们调用 method.streamRequest(message) 的时候，实际是调用 requestMarshaller.stream(requestMessage) 方法，而 requestMarshaller 里面会绑定一个 Parser，这个 Parser 才真正地把对象转成了 InputStream 对象。</p>

<p>讲完序列化在 gRPC 里面的应用后，我们再来看下在 gRPC 里面是怎么完成请求数据“断句”的，就是我们在<a href="https://time.geekbang.org/column/article/199651" target="_blank">[第 02 讲]</a> 中说的那个问题——二进制流经过网络传输后，怎么正确地还原请求前语义？</p>

<p>我们在 gRPC 文档中可以看到，gRPC 的通信协议是基于标准的 HTTP/2 设计的，而 HTTP/2 相对于常用的 HTTP/1.X 来说，它最大的特点就是多路复用、双向流，该怎么理解这个特点呢？这就好比我们生活中的单行道和双行道，HTTP/1.X 就是单行道，HTTP/2 就是双行道。</p>

<p>那既然在请求收到后需要进行请求“断句”，那肯定就需要在发送的时候把断句的符号加上，我们看下在 gRPC 里面是怎么加的？</p>

<p>因为 gRPC 是基于 HTTP/2 协议，而 HTTP/2 传输基本单位是 Frame，Frame 格式是以固定 9 字节长度的 header，后面加上不定长的 payload 组成，协议格式如下图所示：</p>

<p><img src="assets/15b37f24e99f4fc9bcf7d07b7e5d7101.jpg" alt="" /></p>

<p>那在 gRPC 里面就变成怎么构造一个 HTTP/2 的 Frame 了。</p>

<p>现在回看我们上面那个流程图的第 4 步，在 write 到 Netty 里面之前，我们看到在 MessageFramer.writePayload 方法里面会间接调用 writeKnownLengthUncompressed 方法，该方法要做的两件事情就是构造 Frame Header 和 Frame Body，然后再把构造的 Frame 发送到 NettyClientHandler，最后将 Frame 写入到 HTTP/2 Stream 中，完成请求消息的发送。</p>

<h2 id="接收原理">接收原理</h2>

<p>讲完 gRPC 的请求发送原理，我们再来看下服务提供方收到请求后会怎么处理？我们还是接着前面的那个例子，先看下服务提供方代码，具体如下：</p>

<pre><code>static class HelloServiceImpl extends HelloServiceGrpc.HelloServiceImplBase {

  @Override
  public void say(HelloRequest req, StreamObserver&lt;HelloReply&gt; responseObserver) {
    HelloReply reply = HelloReply.newBuilder().setMessage(&quot;Hello &quot; + req.getName()).build();
    responseObserver.onNext(reply);
    responseObserver.onCompleted();
  }
}
</code></pre>

<p>上面 HelloServiceImpl 类是按照 gRPC 使用方式实现了 HelloService 接口逻辑，但是对于调用者来说并不能把它调用过来，因为我们没有把这个接口对外暴露，在 gRPC 里面我们是采用 Build 模式对底层服务进行绑定，具体代码如下：</p>

<pre><code>package io.grpc.hello;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.IOException;


public class HelloWorldServer {

  private Server server;

  /**
  * 对外暴露服务
  **/
  private void start() throws IOException {
    int port = 50051;
    server = ServerBuilder.forPort(port)
        .addService(new HelloServiceImpl())
        .build()
        .start();
    Runtime.getRuntime().addShutdownHook(new Thread() {
      @Override
      public void run() {
        HelloWorldServer.this.stop();
      }
    });
  }

  /**
  * 关闭端口
  **/
  private void stop() {
    if (server != null) {
      server.shutdown();
    }
  }

  /**
  * 优雅关闭
  **/
  private void blockUntilShutdown() throws InterruptedException {
    if (server != null) {
      server.awaitTermination();
    }
  }


  public static void main(String[] args) throws IOException, InterruptedException {
    final HelloWorldServer server = new HelloWorldServer();
    server.start();
    server.blockUntilShutdown();
  }
  
}
</code></pre>

<p>服务对外暴露的目的是让过来的请求在被还原成信息后，能找到对应接口的实现。在这之前，我们需要先保证能正常接收请求，通俗地讲就是要先开启一个 TCP 端口，让调用方可以建立连接，并把二进制数据发送到这个连接通道里面，这里依然只展示最重要的部分。</p>

<p><img src="assets/7f955e44cb2442c9ba58b7155a2efb2d.jpg" alt="" /></p>

<p>这四个步骤是用来开启一个 Netty Server，并绑定编解码逻辑的，如果你暂时看不懂，没关系的，我们可以先忽略细节。我们重点看下 NettyServerHandler 就行了，在这个 Handler 里面会绑定一个 FrameListener，gRPC 会在这个 Listener 里面处理收到数据请求的 Header 和 Body，并且也会处理 Ping、RST 命令等，具体流程如下图所示：</p>

<p><img src="assets/0f4e8eaf066747acaa53f91820877928.jpg" alt="" /></p>

<p>在收到 Header 或者 Body 二进制数据后，NettyServerHandler 上绑定的FrameListener 会把这些二进制数据转到 MessageDeframer 里面，从而实现 gRPC 协议消息的解析 。</p>

<p>那你可能会问，这些 Header 和 Body 数据是怎么分离出来的呢？按照我们前面说的，调用方发过来的是一串二进制数据，这就是我们前面开启 Netty Server 的时候绑定 Default HTTP/2FrameReader 的作用，它能帮助我们按照 HTTP/2 协议的格式自动切分出 Header 和 Body 数据来，而对我们上层应用 gRPC 来说，它可以直接拿拆分后的数据来用。</p>

<h2 id="总结">总结</h2>

<p>这是我们基础篇的最后一讲，我们采用剖析 gRPC 源码的方式来学习如何实现一个完整的 RPC。当然整个 gRPC 的代码量可比这多得多，但今天的主要目就是想让你把前面所学的序列化、协议等方面的知识落实到具体代码上，所以我们这儿只分析了 gRPC 收发请求两个过程。</p>

<p>实现了这两个过程，我们就可以完成一个点对点的 RPC 功能，但在实际使用的时候，我们的服务提供方通常都是以一个集群的方式对外提供服务的，所以在 gRPC 里面你还可以看到负载均衡、服务发现等功能。而且 gRPC 采用的是 HTTP/2 协议，我们还可以通过 Stream 方式来调用服务，以提升调用性能。</p>

<p>总的来说，其实我们可以简单地认为<strong>gRPC 就是采用 HTTP/2 协议，并且默认采用 PB 序列化方式的一种 RPC</strong>，它充分利用了 HTTP/2 的多路复用特性，使得我们可以在同一条链路上双向发送不同的 Stream 数据，以解决 HTTP/1.X 存在的性能问题。</p>

<h2 id="课后思考">课后思考</h2>

<p>我们讲到，在 gRPC 调用的时候，我们有一个关键步骤就是把对象转成可传输的二进制，但是在 gRPC 里面，我们并没有直接转成二进制数组，而是返回一个 InputStream，你知道这样做的好处是什么吗？</p>

<p>欢迎留言和我分享你的答案，也欢迎你把文章分享给你的朋友，邀请他加入学习。我们下节课再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e4888888ddd0d5d5d4d3a48389858d88ca878b89" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f120e792c61cdc2',t:'MTczNDA1MjAzOS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>