<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;PV、PVC体系是不是多此一举？从本地持久化卷谈起>
        <link rel="icon" href="/static/favicon.png">
        <title>29 PV、PVC体系是不是多此一举？从本地持久化卷谈起 </title>
        
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
                        <a class="menu-item" id="00 开篇词 打通“容器技术”的任督二脉.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e6%89%93%e9%80%9a%e2%80%9c%e5%ae%b9%e5%99%a8%e6%8a%80%e6%9c%af%e2%80%9d%e7%9a%84%e4%bb%bb%e7%9d%a3%e4%ba%8c%e8%84%89.md">00 开篇词 打通“容器技术”的任督二脉.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 预习篇 · 小鲸鱼大事记（一）：初出茅庐.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/01%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20%e5%b0%8f%e9%b2%b8%e9%b1%bc%e5%a4%a7%e4%ba%8b%e8%ae%b0%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%88%9d%e5%87%ba%e8%8c%85%e5%ba%90.md">01 预习篇 · 小鲸鱼大事记（一）：初出茅庐.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 预习篇 · 小鲸鱼大事记（二）：崭露头角.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/02%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20%e5%b0%8f%e9%b2%b8%e9%b1%bc%e5%a4%a7%e4%ba%8b%e8%ae%b0%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%b4%ad%e9%9c%b2%e5%a4%b4%e8%a7%92.md">02 预习篇 · 小鲸鱼大事记（二）：崭露头角.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 预习篇 · 小鲸鱼大事记（三）：群雄并起.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/03%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20%e5%b0%8f%e9%b2%b8%e9%b1%bc%e5%a4%a7%e4%ba%8b%e8%ae%b0%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e7%be%a4%e9%9b%84%e5%b9%b6%e8%b5%b7.md">03 预习篇 · 小鲸鱼大事记（三）：群雄并起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 预习篇 · 小鲸鱼大事记（四）：尘埃落定.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/04%20%e9%a2%84%e4%b9%a0%e7%af%87%20%c2%b7%20%e5%b0%8f%e9%b2%b8%e9%b1%bc%e5%a4%a7%e4%ba%8b%e8%ae%b0%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e5%b0%98%e5%9f%83%e8%90%bd%e5%ae%9a.md">04 预习篇 · 小鲸鱼大事记（四）：尘埃落定.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 白话容器基础（一）：从进程说开去.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/05%20%e7%99%bd%e8%af%9d%e5%ae%b9%e5%99%a8%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e4%bb%8e%e8%bf%9b%e7%a8%8b%e8%af%b4%e5%bc%80%e5%8e%bb.md">05 白话容器基础（一）：从进程说开去.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 白话容器基础（二）：隔离与限制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/06%20%e7%99%bd%e8%af%9d%e5%ae%b9%e5%99%a8%e5%9f%ba%e7%a1%80%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e9%9a%94%e7%a6%bb%e4%b8%8e%e9%99%90%e5%88%b6.md">06 白话容器基础（二）：隔离与限制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 白话容器基础（三）：深入理解容器镜像.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/07%20%e7%99%bd%e8%af%9d%e5%ae%b9%e5%99%a8%e5%9f%ba%e7%a1%80%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3%e5%ae%b9%e5%99%a8%e9%95%9c%e5%83%8f.md">07 白话容器基础（三）：深入理解容器镜像.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 白话容器基础（四）：重新认识Docker容器.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/08%20%e7%99%bd%e8%af%9d%e5%ae%b9%e5%99%a8%e5%9f%ba%e7%a1%80%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9a%e9%87%8d%e6%96%b0%e8%ae%a4%e8%af%86Docker%e5%ae%b9%e5%99%a8.md">08 白话容器基础（四）：重新认识Docker容器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 从容器到容器云：谈谈Kubernetes的本质.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/09%20%e4%bb%8e%e5%ae%b9%e5%99%a8%e5%88%b0%e5%ae%b9%e5%99%a8%e4%ba%91%ef%bc%9a%e8%b0%88%e8%b0%88Kubernetes%e7%9a%84%e6%9c%ac%e8%b4%a8.md">09 从容器到容器云：谈谈Kubernetes的本质.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 Kubernetes一键部署利器：kubeadm.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/10%20Kubernetes%e4%b8%80%e9%94%ae%e9%83%a8%e7%bd%b2%e5%88%a9%e5%99%a8%ef%bc%9akubeadm.md">10 Kubernetes一键部署利器：kubeadm.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 从0到1：搭建一个完整的Kubernetes集群.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/11%20%e4%bb%8e0%e5%88%b01%ef%bc%9a%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%e5%ae%8c%e6%95%b4%e7%9a%84Kubernetes%e9%9b%86%e7%be%a4.md">11 从0到1：搭建一个完整的Kubernetes集群.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 牛刀小试：我的第一个容器化应用.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/12%20%e7%89%9b%e5%88%80%e5%b0%8f%e8%af%95%ef%bc%9a%e6%88%91%e7%9a%84%e7%ac%ac%e4%b8%80%e4%b8%aa%e5%ae%b9%e5%99%a8%e5%8c%96%e5%ba%94%e7%94%a8.md">12 牛刀小试：我的第一个容器化应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 为什么我们需要Pod？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/13%20%e4%b8%ba%e4%bb%80%e4%b9%88%e6%88%91%e4%bb%ac%e9%9c%80%e8%a6%81Pod%ef%bc%9f.md">13 为什么我们需要Pod？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 深入解析Pod对象（一）：基本概念.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/14%20%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90Pod%e5%af%b9%e8%b1%a1%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e5%9f%ba%e6%9c%ac%e6%a6%82%e5%bf%b5.md">14 深入解析Pod对象（一）：基本概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 深入解析Pod对象（二）：使用进阶.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/15%20%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90Pod%e5%af%b9%e8%b1%a1%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e4%bd%bf%e7%94%a8%e8%bf%9b%e9%98%b6.md">15 深入解析Pod对象（二）：使用进阶.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 编排其实很简单：谈谈“控制器”模型.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/16%20%e7%bc%96%e6%8e%92%e5%85%b6%e5%ae%9e%e5%be%88%e7%ae%80%e5%8d%95%ef%bc%9a%e8%b0%88%e8%b0%88%e2%80%9c%e6%8e%a7%e5%88%b6%e5%99%a8%e2%80%9d%e6%a8%a1%e5%9e%8b.md">16 编排其实很简单：谈谈“控制器”模型.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 经典PaaS的记忆：作业副本与水平扩展.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/17%20%e7%bb%8f%e5%85%b8PaaS%e7%9a%84%e8%ae%b0%e5%bf%86%ef%bc%9a%e4%bd%9c%e4%b8%9a%e5%89%af%e6%9c%ac%e4%b8%8e%e6%b0%b4%e5%b9%b3%e6%89%a9%e5%b1%95.md">17 经典PaaS的记忆：作业副本与水平扩展.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 深入理解StatefulSet（一）：拓扑状态.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/18%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3StatefulSet%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9a%e6%8b%93%e6%89%91%e7%8a%b6%e6%80%81.md">18 深入理解StatefulSet（一）：拓扑状态.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 深入理解StatefulSet（二）：存储状态.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/19%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3StatefulSet%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e5%ad%98%e5%82%a8%e7%8a%b6%e6%80%81.md">19 深入理解StatefulSet（二）：存储状态.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 深入理解StatefulSet（三）：有状态应用实践.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/20%20%e6%b7%b1%e5%85%a5%e7%90%86%e8%a7%a3StatefulSet%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9a%e6%9c%89%e7%8a%b6%e6%80%81%e5%ba%94%e7%94%a8%e5%ae%9e%e8%b7%b5.md">20 深入理解StatefulSet（三）：有状态应用实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 容器化守护进程的意义：DaemonSet.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/21%20%e5%ae%b9%e5%99%a8%e5%8c%96%e5%ae%88%e6%8a%a4%e8%bf%9b%e7%a8%8b%e7%9a%84%e6%84%8f%e4%b9%89%ef%bc%9aDaemonSet.md">21 容器化守护进程的意义：DaemonSet.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 撬动离线业务：Job与CronJob.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/22%20%e6%92%ac%e5%8a%a8%e7%a6%bb%e7%ba%bf%e4%b8%9a%e5%8a%a1%ef%bc%9aJob%e4%b8%8eCronJob.md">22 撬动离线业务：Job与CronJob.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 声明式API与Kubernetes编程范式.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/23%20%e5%a3%b0%e6%98%8e%e5%bc%8fAPI%e4%b8%8eKubernetes%e7%bc%96%e7%a8%8b%e8%8c%83%e5%bc%8f.md">23 声明式API与Kubernetes编程范式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 深入解析声明式API（一）：API对象的奥秘.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/24%20%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90%e5%a3%b0%e6%98%8e%e5%bc%8fAPI%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9aAPI%e5%af%b9%e8%b1%a1%e7%9a%84%e5%a5%a5%e7%a7%98.md">24 深入解析声明式API（一）：API对象的奥秘.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 深入解析声明式API（二）：编写自定义控制器.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/25%20%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90%e5%a3%b0%e6%98%8e%e5%bc%8fAPI%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9a%e7%bc%96%e5%86%99%e8%87%aa%e5%ae%9a%e4%b9%89%e6%8e%a7%e5%88%b6%e5%99%a8.md">25 深入解析声明式API（二）：编写自定义控制器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 基于角色的权限控制：RBAC.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/26%20%e5%9f%ba%e4%ba%8e%e8%a7%92%e8%89%b2%e7%9a%84%e6%9d%83%e9%99%90%e6%8e%a7%e5%88%b6%ef%bc%9aRBAC.md">26 基于角色的权限控制：RBAC.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 聪明的微创新：Operator工作原理解读.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/27%20%e8%81%aa%e6%98%8e%e7%9a%84%e5%be%ae%e5%88%9b%e6%96%b0%ef%bc%9aOperator%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86%e8%a7%a3%e8%af%bb.md">27 聪明的微创新：Operator工作原理解读.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 PV、PVC、StorageClass，这些到底在说啥？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/28%20PV%e3%80%81PVC%e3%80%81StorageClass%ef%bc%8c%e8%bf%99%e4%ba%9b%e5%88%b0%e5%ba%95%e5%9c%a8%e8%af%b4%e5%95%a5%ef%bc%9f.md">28 PV、PVC、StorageClass，这些到底在说啥？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 PV、PVC体系是不是多此一举？从本地持久化卷谈起.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/29%20PV%e3%80%81PVC%e4%bd%93%e7%b3%bb%e6%98%af%e4%b8%8d%e6%98%af%e5%a4%9a%e6%ad%a4%e4%b8%80%e4%b8%be%ef%bc%9f%e4%bb%8e%e6%9c%ac%e5%9c%b0%e6%8c%81%e4%b9%85%e5%8c%96%e5%8d%b7%e8%b0%88%e8%b5%b7.md">29 PV、PVC体系是不是多此一举？从本地持久化卷谈起.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 编写自己的存储插件：FlexVolume与CSI.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/30%20%e7%bc%96%e5%86%99%e8%87%aa%e5%b7%b1%e7%9a%84%e5%ad%98%e5%82%a8%e6%8f%92%e4%bb%b6%ef%bc%9aFlexVolume%e4%b8%8eCSI.md">30 编写自己的存储插件：FlexVolume与CSI.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 容器存储实践：CSI插件编写指南.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/31%20%e5%ae%b9%e5%99%a8%e5%ad%98%e5%82%a8%e5%ae%9e%e8%b7%b5%ef%bc%9aCSI%e6%8f%92%e4%bb%b6%e7%bc%96%e5%86%99%e6%8c%87%e5%8d%97.md">31 容器存储实践：CSI插件编写指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 浅谈容器网络.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/32%20%e6%b5%85%e8%b0%88%e5%ae%b9%e5%99%a8%e7%bd%91%e7%bb%9c.md">32 浅谈容器网络.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 深入解析容器跨主机网络.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/33%20%e6%b7%b1%e5%85%a5%e8%a7%a3%e6%9e%90%e5%ae%b9%e5%99%a8%e8%b7%a8%e4%b8%bb%e6%9c%ba%e7%bd%91%e7%bb%9c.md">33 深入解析容器跨主机网络.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 Kubernetes网络模型与CNI网络插件.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/34%20Kubernetes%e7%bd%91%e7%bb%9c%e6%a8%a1%e5%9e%8b%e4%b8%8eCNI%e7%bd%91%e7%bb%9c%e6%8f%92%e4%bb%b6.md">34 Kubernetes网络模型与CNI网络插件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 解读Kubernetes三层网络方案.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/35%20%e8%a7%a3%e8%af%bbKubernetes%e4%b8%89%e5%b1%82%e7%bd%91%e7%bb%9c%e6%96%b9%e6%a1%88.md">35 解读Kubernetes三层网络方案.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 为什么说Kubernetes只有soft multi-tenancy？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/36%20%e4%b8%ba%e4%bb%80%e4%b9%88%e8%af%b4Kubernetes%e5%8f%aa%e6%9c%89soft%20multi-tenancy%ef%bc%9f.md">36 为什么说Kubernetes只有soft multi-tenancy？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 找到容器不容易：Service、DNS与服务发现.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/37%20%e6%89%be%e5%88%b0%e5%ae%b9%e5%99%a8%e4%b8%8d%e5%ae%b9%e6%98%93%ef%bc%9aService%e3%80%81DNS%e4%b8%8e%e6%9c%8d%e5%8a%a1%e5%8f%91%e7%8e%b0.md">37 找到容器不容易：Service、DNS与服务发现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 从外界连通Service与Service调试“三板斧”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/38%20%e4%bb%8e%e5%a4%96%e7%95%8c%e8%bf%9e%e9%80%9aService%e4%b8%8eService%e8%b0%83%e8%af%95%e2%80%9c%e4%b8%89%e6%9d%bf%e6%96%a7%e2%80%9d.md">38 从外界连通Service与Service调试“三板斧”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 谈谈Service与Ingress.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/39%20%e8%b0%88%e8%b0%88Service%e4%b8%8eIngress.md">39 谈谈Service与Ingress.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 Kubernetes的资源模型与资源管理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/40%20Kubernetes%e7%9a%84%e8%b5%84%e6%ba%90%e6%a8%a1%e5%9e%8b%e4%b8%8e%e8%b5%84%e6%ba%90%e7%ae%a1%e7%90%86.md">40 Kubernetes的资源模型与资源管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 十字路口上的Kubernetes默认调度器.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/41%20%e5%8d%81%e5%ad%97%e8%b7%af%e5%8f%a3%e4%b8%8a%e7%9a%84Kubernetes%e9%bb%98%e8%ae%a4%e8%b0%83%e5%ba%a6%e5%99%a8.md">41 十字路口上的Kubernetes默认调度器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 Kubernetes默认调度器调度策略解析.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/42%20Kubernetes%e9%bb%98%e8%ae%a4%e8%b0%83%e5%ba%a6%e5%99%a8%e8%b0%83%e5%ba%a6%e7%ad%96%e7%95%a5%e8%a7%a3%e6%9e%90.md">42 Kubernetes默认调度器调度策略解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 Kubernetes默认调度器的优先级与抢占机制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/43%20Kubernetes%e9%bb%98%e8%ae%a4%e8%b0%83%e5%ba%a6%e5%99%a8%e7%9a%84%e4%bc%98%e5%85%88%e7%ba%a7%e4%b8%8e%e6%8a%a2%e5%8d%a0%e6%9c%ba%e5%88%b6.md">43 Kubernetes默认调度器的优先级与抢占机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 Kubernetes GPU管理与Device Plugin机制.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/44%20Kubernetes%20GPU%e7%ae%a1%e7%90%86%e4%b8%8eDevice%20Plugin%e6%9c%ba%e5%88%b6.md">44 Kubernetes GPU管理与Device Plugin机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 幕后英雄：SIG-Node与CRI.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/45%20%e5%b9%95%e5%90%8e%e8%8b%b1%e9%9b%84%ef%bc%9aSIG-Node%e4%b8%8eCRI.md">45 幕后英雄：SIG-Node与CRI.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 解读 CRI 与 容器运行时.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/46%20%e8%a7%a3%e8%af%bb%20CRI%20%e4%b8%8e%20%e5%ae%b9%e5%99%a8%e8%bf%90%e8%a1%8c%e6%97%b6.md">46 解读 CRI 与 容器运行时.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 绝不仅仅是安全：Kata Containers 与 gVisor.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/47%20%e7%bb%9d%e4%b8%8d%e4%bb%85%e4%bb%85%e6%98%af%e5%ae%89%e5%85%a8%ef%bc%9aKata%20Containers%20%e4%b8%8e%20gVisor.md">47 绝不仅仅是安全：Kata Containers 与 gVisor.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 Prometheus、Metrics Server与Kubernetes监控体系.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/48%20Prometheus%e3%80%81Metrics%20Server%e4%b8%8eKubernetes%e7%9b%91%e6%8e%a7%e4%bd%93%e7%b3%bb.md">48 Prometheus、Metrics Server与Kubernetes监控体系.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 Custom Metrics_ 让Auto Scaling不再“食之无味”.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/49%20Custom%20Metrics_%20%e8%ae%a9Auto%20Scaling%e4%b8%8d%e5%86%8d%e2%80%9c%e9%a3%9f%e4%b9%8b%e6%97%a0%e5%91%b3%e2%80%9d.md">49 Custom Metrics_ 让Auto Scaling不再“食之无味”.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 让日志无处可逃：容器日志收集与管理.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/50%20%e8%ae%a9%e6%97%a5%e5%bf%97%e6%97%a0%e5%a4%84%e5%8f%af%e9%80%83%ef%bc%9a%e5%ae%b9%e5%99%a8%e6%97%a5%e5%bf%97%e6%94%b6%e9%9b%86%e4%b8%8e%e7%ae%a1%e7%90%86.md">50 让日志无处可逃：容器日志收集与管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 谈谈Kubernetes开源社区和未来走向.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/51%20%e8%b0%88%e8%b0%88Kubernetes%e5%bc%80%e6%ba%90%e7%a4%be%e5%8c%ba%e5%92%8c%e6%9c%aa%e6%9d%a5%e8%b5%b0%e5%90%91.md">51 谈谈Kubernetes开源社区和未来走向.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="52 答疑：在问题中解决问题，在思考中产生思考.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/52%20%e7%ad%94%e7%96%91%ef%bc%9a%e5%9c%a8%e9%97%ae%e9%a2%98%e4%b8%ad%e8%a7%a3%e5%86%b3%e9%97%ae%e9%a2%98%ef%bc%8c%e5%9c%a8%e6%80%9d%e8%80%83%e4%b8%ad%e4%ba%a7%e7%94%9f%e6%80%9d%e8%80%83.md">52 答疑：在问题中解决问题，在思考中产生思考.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 2019 年，容器技术生态会发生些什么？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%202019%20%e5%b9%b4%ef%bc%8c%e5%ae%b9%e5%99%a8%e6%8a%80%e6%9c%af%e7%94%9f%e6%80%81%e4%bc%9a%e5%8f%91%e7%94%9f%e4%ba%9b%e4%bb%80%e4%b9%88%ef%bc%9f.md">特别放送 2019 年，容器技术生态会发生些什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 基于 Kubernetes 的云原生应用管理，到底应该怎么做？.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e5%9f%ba%e4%ba%8e%20Kubernetes%20%e7%9a%84%e4%ba%91%e5%8e%9f%e7%94%9f%e5%ba%94%e7%94%a8%e7%ae%a1%e7%90%86%ef%bc%8c%e5%88%b0%e5%ba%95%e5%ba%94%e8%af%a5%e6%80%8e%e4%b9%88%e5%81%9a%ef%bc%9f.md">特别放送 基于 Kubernetes 的云原生应用管理，到底应该怎么做？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 Kubernetes：赢开发者赢天下.md" href="/%e4%b8%93%e6%a0%8f/%e6%b7%b1%e5%85%a5%e5%89%96%e6%9e%90Kubernetes/%e7%bb%93%e6%9d%9f%e8%af%ad%20Kubernetes%ef%bc%9a%e8%b5%a2%e5%bc%80%e5%8f%91%e8%80%85%e8%b5%a2%e5%a4%a9%e4%b8%8b.md">结束语 Kubernetes：赢开发者赢天下.md</a>
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
                            <h1 id="title" data-id="29 PV、PVC体系是不是多此一举？从本地持久化卷谈起" class="title">29 PV、PVC体系是不是多此一举？从本地持久化卷谈起</h1>
                            <div><p>你好，我是张磊。今天我和你分享的主题是：PV、PVC体系是不是多此一举？从本地持久化卷谈起。</p>

<p>在上一篇文章中，我为你详细讲解了PV、PVC持久化存储体系在Kubernetes项目中的设计和实现原理。而在文章最后的思考题中，我为你留下了这样一个讨论话题：像PV、PVC这样的用法，是不是有“过度设计”的嫌疑？</p>

<p>比如，我们公司的运维人员可以像往常一样维护一套NFS或者Ceph服务器，根本不必学习Kubernetes。而开发人员，则完全可以靠“复制粘贴”的方式，在Pod的YAML文件里填上Volumes字段，而不需要去使用PV和PVC。</p>

<p>实际上，如果只是为了职责划分，PV、PVC体系确实不见得比直接在Pod里声明Volumes字段有什么优势。</p>

<p>不过，你有没有想过这样一个问题，如果<a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes/#types-of-persistent-volumes" target="_blank">Kubernetes内置的20种持久化数据卷实现</a>，都没办法满足你的容器存储需求时，该怎么办？</p>

<p>这个情况乍一听起来有点不可思议。但实际上，凡是鼓捣过开源项目的读者应该都有所体会，“不能用”“不好用”“需要定制开发”，这才是落地开源基础设施项目的三大常态。</p>

<p>而在持久化存储领域，用户呼声最高的定制化需求，莫过于支持“本地”持久化存储了。</p>

<p>也就是说，用户希望Kubernetes能够直接使用宿主机上的本地磁盘目录，而不依赖于远程存储服务，来提供“持久化”的容器Volume。</p>

<p>这样做的好处很明显，由于这个Volume直接使用的是本地磁盘，尤其是SSD盘，它的读写性能相比于大多数远程存储来说，要好得多。这个需求对本地物理服务器部署的私有Kubernetes集群来说，非常常见。</p>

<p>所以，Kubernetes在v1.10之后，就逐渐依靠PV、PVC体系实现了这个特性。这个特性的名字叫作：Local Persistent Volume。</p>

<p>不过，首先需要明确的是，<strong>Local Persistent Volume并不适用于所有应用</strong>。事实上，它的适用范围非常固定，比如：高优先级的系统应用，需要在多个不同节点上存储数据，并且对I/O较为敏感。典型的应用包括：分布式数据存储比如MongoDB、Cassandra等，分布式文件系统比如GlusterFS、Ceph等，以及需要在本地磁盘上进行大量数据缓存的分布式应用。</p>

<p>其次，相比于正常的PV，一旦这些节点宕机且不能恢复时，Local Persistent Volume的数据就可能丢失。这就要求<strong>使用Local Persistent Volume的应用必须具备数据备份和恢复的能力</strong>，允许你把这些数据定时备份在其他位置。</p>

<p>接下来，我就为你深入讲解一下这个特性。</p>

<p>不难想象，Local Persistent Volume的设计，主要面临两个难点。</p>

<p><strong>第一个难点在于</strong>：如何把本地磁盘抽象成PV。</p>

<p>可能你会说，Local Persistent Volume，不就等同于hostPath加NodeAffinity吗？</p>

<p>比如，一个Pod可以声明使用类型为Local的PV，而这个PV其实就是一个hostPath类型的Volume。如果这个hostPath对应的目录，已经在节点A上被事先创建好了。那么，我只需要再给这个Pod加上一个nodeAffinity=nodeA，不就可以使用这个Volume了吗？</p>

<p>事实上，<strong>你绝不应该把一个宿主机上的目录当作PV使用</strong>。这是因为，这种本地目录的存储行为完全不可控，它所在的磁盘随时都可能被应用写满，甚至造成整个宿主机宕机。而且，不同的本地目录之间也缺乏哪怕最基础的I/O隔离机制。</p>

<p>所以，一个Local Persistent Volume对应的存储介质，一定是一块额外挂载在宿主机的磁盘或者块设备（“额外”的意思是，它不应该是宿主机根目录所使用的主硬盘）。这个原则，我们可以称为“<strong>一个PV一块盘</strong>”。</p>

<p><strong>第二个难点在于</strong>：调度器如何保证Pod始终能被正确地调度到它所请求的Local Persistent Volume所在的节点上呢？</p>

<p>造成这个问题的原因在于，对于常规的PV来说，Kubernetes都是先调度Pod到某个节点上，然后，再通过“两阶段处理”来“持久化”这台机器上的Volume目录，进而完成Volume目录与容器的绑定挂载。</p>

<p>可是，对于Local PV来说，节点上可供使用的磁盘（或者块设备），必须是运维人员提前准备好的。它们在不同节点上的挂载情况可以完全不同，甚至有的节点可以没这种磁盘。</p>

<p>所以，这时候，调度器就必须能够知道所有节点与Local Persistent Volume对应的磁盘的关联关系，然后根据这个信息来调度Pod。</p>

<p>这个原则，我们可以称为“<strong>在调度的时候考虑Volume分布</strong>”。在Kubernetes的调度器里，有一个叫作VolumeBindingChecker的过滤条件专门负责这个事情。在Kubernetes v1.11中，这个过滤条件已经默认开启了。</p>

<p>基于上述讲述，在开始使用Local Persistent Volume之前，你首先需要在集群里配置好磁盘或者块设备。在公有云上，这个操作等同于给虚拟机额外挂载一个磁盘，比如GCE的Local SSD类型的磁盘就是一个典型例子。</p>

<p>而在我们部署的私有环境中，你有两种办法来完成这个步骤。</p>

<ul>
<li>第一种，当然就是给你的宿主机挂载并格式化一个可用的本地磁盘，这也是最常规的操作；</li>
<li>第二种，对于实验环境，你其实可以在宿主机上挂载几个RAM Disk（内存盘）来模拟本地磁盘。</li>
</ul>

<p>接下来，我会使用第二种方法，在我们之前部署的Kubernetes集群上进行实践。</p>

<p><strong>首先</strong>，在名叫node-1的宿主机上创建一个挂载点，比如/mnt/disks；<strong>然后</strong>，用几个RAM Disk来模拟本地磁盘，如下所示：</p>

<pre><code># 在node-1上执行
$ mkdir /mnt/disks
$ for vol in vol1 vol2 vol3; do
    mkdir /mnt/disks/$vol
    mount -t tmpfs $vol /mnt/disks/$vol
done
</code></pre>

<p>需要注意的是，如果你希望其他节点也能支持Local Persistent Volume的话，那就需要为它们也执行上述操作，并且确保这些磁盘的名字（vol1、vol2等）都不重复。</p>

<p>接下来，我们就可以为这些本地磁盘定义对应的PV了，如下所示：</p>

<pre><code>apiVersion: v1
kind: PersistentVolume
metadata:
  name: example-pv
spec:
  capacity:
    storage: 5Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Delete
  storageClassName: local-storage
  local:
    path: /mnt/disks/vol1
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - node-1
</code></pre>

<p>可以看到，这个PV的定义里：local字段，指定了它是一个Local Persistent Volume；而path字段，指定的正是这个PV对应的本地磁盘的路径，即：/mnt/disks/vol1。</p>

<p>当然了，这也就意味着如果Pod要想使用这个PV，那它就必须运行在node-1上。所以，在这个PV的定义里，需要有一个nodeAffinity字段指定node-1这个节点的名字。这样，调度器在调度Pod的时候，就能够知道一个PV与节点的对应关系，从而做出正确的选择。<strong>这正是Kubernetes实现“在调度的时候就考虑Volume分布”的主要方法。</strong></p>

<p><strong>接下来</strong>，我们就可以使用kubect create来创建这个PV，如下所示：</p>

<pre><code>$ kubectl create -f local-pv.yaml 
persistentvolume/example-pv created

$ kubectl get pv
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY  STATUS      CLAIM             STORAGECLASS    REASON    AGE
example-pv   5Gi        RWO            Delete           Available                     local-storage             16s
</code></pre>

<p>可以看到，这个PV创建后，进入了Available（可用）状态。</p>

<p>而正如我在上一篇文章里所建议的那样，使用PV和PVC的最佳实践，是你要创建一个StorageClass来描述这个PV，如下所示：</p>

<pre><code>kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: local-storage
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
</code></pre>

<p>这个StorageClass的名字，叫作local-storage。需要注意的是，在它的provisioner字段，我们指定的是no-provisioner。这是因为Local Persistent Volume目前尚不支持Dynamic Provisioning，所以它没办法在用户创建PVC的时候，就自动创建出对应的PV。也就是说，我们前面创建PV的操作，是不可以省略的。</p>

<p>与此同时，这个StorageClass还定义了一个volumeBindingMode=WaitForFirstConsumer的属性。它是Local Persistent Volume里一个非常重要的特性，即：<strong>延迟绑定</strong>。</p>

<p>我们知道，当你提交了PV和PVC的YAML文件之后，Kubernetes就会根据它们俩的属性，以及它们指定的StorageClass来进行绑定。只有绑定成功后，Pod才能通过声明这个PVC来使用对应的PV。</p>

<p>可是，如果你使用的是Local Persistent Volume的话，就会发现，这个流程根本行不通。</p>

<p>比如，现在你有一个Pod，它声明使用的PVC叫作pvc-1。并且，我们规定，这个Pod只能运行在node-2上。</p>

<p>而在Kubernetes集群中，有两个属性（比如：大小、读写权限）相同的Local类型的PV。</p>

<p>其中，第一个PV的名字叫作pv-1，它对应的磁盘所在的节点是node-1。而第二个PV的名字叫作pv-2，它对应的磁盘所在的节点是node-2。</p>

<p>假设现在，Kubernetes的Volume控制循环里，首先检查到了pvc-1和pv-1的属性是匹配的，于是就将它们俩绑定在一起。</p>

<p>然后，你用kubectl create创建了这个Pod。</p>

<p>这时候，问题就出现了。</p>

<p>调度器看到，这个Pod所声明的pvc-1已经绑定了pv-1，而pv-1所在的节点是node-1，根据“调度器必须在调度的时候考虑Volume分布”的原则，这个Pod自然会被调度到node-1上。</p>

<p>可是，我们前面已经规定过，这个Pod根本不允许运行在node-1上。所以。最后的结果就是，这个Pod的调度必然会失败。</p>

<p><strong>这就是为什么，在使用Local Persistent Volume的时候，我们必须想办法推迟这个“绑定”操作。</strong></p>

<p>那么，具体推迟到什么时候呢？</p>

<p><strong>答案是：推迟到调度的时候。</strong></p>

<p>所以说，StorageClass里的volumeBindingMode=WaitForFirstConsumer的含义，就是告诉Kubernetes里的Volume控制循环（“红娘”）：虽然你已经发现这个StorageClass关联的PVC与PV可以绑定在一起，但请不要现在就执行绑定操作（即：设置PVC的VolumeName字段）。</p>

<p>而要等到第一个声明使用该PVC的Pod出现在调度器之后，调度器再综合考虑所有的调度规则，当然也包括每个PV所在的节点位置，来统一决定，这个Pod声明的PVC，到底应该跟哪个PV进行绑定。</p>

<p>这样，在上面的例子里，由于这个Pod不允许运行在pv-1所在的节点node-1，所以它的PVC最后会跟pv-2绑定，并且Pod也会被调度到node-2上。</p>

<p>所以，通过这个延迟绑定机制，原本实时发生的PVC和PV的绑定过程，就被延迟到了Pod第一次调度的时候在调度器中进行，从而保证了这个<strong>绑定结果不会影响Pod的正常调度</strong>。</p>

<p>当然，在具体实现中，调度器实际上维护了一个与Volume Controller类似的控制循环，专门负责为那些声明了“延迟绑定”的PV和PVC进行绑定工作。</p>

<p>通过这样的设计，这个额外的绑定操作，并不会拖慢调度器的性能。而当一个Pod的PVC尚未完成绑定时，调度器也不会等待，而是会直接把这个Pod重新放回到待调度队列，等到下一个调度周期再做处理。</p>

<p>在明白了这个机制之后，我们就可以创建StorageClass了，如下所示：</p>

<pre><code>$ kubectl create -f local-sc.yaml 
storageclass.storage.k8s.io/local-storage created
</code></pre>

<p>接下来，我们只需要定义一个非常普通的PVC，就可以让Pod使用到上面定义好的Local Persistent Volume了，如下所示：</p>

<pre><code>kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: example-local-claim
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
  storageClassName: local-storage
</code></pre>

<p>可以看到，这个PVC没有任何特别的地方。唯一需要注意的是，它声明的storageClassName是local-storage。所以，将来Kubernetes的Volume Controller看到这个PVC的时候，不会为它进行绑定操作。</p>

<p>现在，我们来创建这个PVC：</p>

<pre><code>$ kubectl create -f local-pvc.yaml 
persistentvolumeclaim/example-local-claim created

$ kubectl get pvc
NAME                  STATUS    VOLUME    CAPACITY   ACCESS MODES   STORAGECLASS    AGE
example-local-claim   Pending                                       local-storage   7s
</code></pre>

<p>可以看到，尽管这个时候，Kubernetes里已经存在了一个可以与PVC匹配的PV，但这个PVC依然处于Pending状态，也就是等待绑定的状态。</p>

<p>然后，我们编写一个Pod来声明使用这个PVC，如下所示：</p>

<pre><code>kind: Pod
apiVersion: v1
metadata:
  name: example-pv-pod
spec:
  volumes:
    - name: example-pv-storage
      persistentVolumeClaim:
       claimName: example-local-claim
  containers:
    - name: example-pv-container
      image: nginx
      ports:
        - containerPort: 80
          name: &quot;http-server&quot;
      volumeMounts:
        - mountPath: &quot;/usr/share/nginx/html&quot;
          name: example-pv-storage
</code></pre>

<p>这个Pod没有任何特别的地方，你只需要注意，它的volumes字段声明要使用前面定义的、名叫example-local-claim的PVC即可。</p>

<p>而我们一旦使用kubectl create创建这个Pod，就会发现，我们前面定义的PVC，会立刻变成Bound状态，与前面定义的PV绑定在了一起，如下所示：</p>

<pre><code>$ kubectl create -f local-pod.yaml 
pod/example-pv-pod created

$ kubectl get pvc
NAME                  STATUS    VOLUME       CAPACITY   ACCESS MODES   STORAGECLASS    AGE
example-local-claim   Bound     example-pv   5Gi        RWO            local-storage   6h
</code></pre>

<p>也就是说，在我们创建的Pod进入调度器之后，“绑定”操作才开始进行。</p>

<p>这时候，我们可以尝试在这个Pod的Volume目录里，创建一个测试文件，比如：</p>

<pre><code>$ kubectl exec -it example-pv-pod -- /bin/sh
# cd /usr/share/nginx/html
# touch test.txt
</code></pre>

<p>然后，登录到node-1这台机器上，查看一下它的 /mnt/disks/vol1目录下的内容，你就可以看到刚刚创建的这个文件：</p>

<pre><code># 在node-1上
$ ls /mnt/disks/vol1
test.txt
</code></pre>

<p>而如果你重新创建这个Pod的话，就会发现，我们之前创建的测试文件，依然被保存在这个持久化Volume当中：</p>

<pre><code>$ kubectl delete -f local-pod.yaml 

$ kubectl create -f local-pod.yaml 

$ kubectl exec -it example-pv-pod -- /bin/sh
# ls /usr/share/nginx/html
# touch test.txt
</code></pre>

<p>这就说明，像Kubernetes这样构建出来的、基于本地存储的Volume，完全可以提供容器持久化存储的功能。所以，像StatefulSet这样的有状态编排工具，也完全可以通过声明Local类型的PV和PVC，来管理应用的存储状态。</p>

<p><strong>需要注意的是，我们上面手动创建PV的方式，即Static的PV管理方式，在删除PV时需要按如下流程执行操作：</strong></p>

<ol>
<li><p>删除使用这个PV的Pod；</p></li>

<li><p>从宿主机移除本地磁盘（比如，umount它）；</p></li>

<li><p>删除PVC；</p></li>

<li><p>删除PV。</p></li>
</ol>

<p>如果不按照这个流程的话，这个PV的删除就会失败。</p>

<p>当然，由于上面这些创建PV和删除PV的操作比较繁琐，Kubernetes其实提供了一个Static Provisioner来帮助你管理这些PV。</p>

<p>比如，我们现在的所有磁盘，都挂载在宿主机的/mnt/disks目录下。</p>

<p>那么，当Static Provisioner启动后，它就会通过DaemonSet，自动检查每个宿主机的/mnt/disks目录。然后，调用Kubernetes API，为这些目录下面的每一个挂载，创建一个对应的PV对象出来。这些自动创建的PV，如下所示：</p>

<pre><code>$ kubectl get pv
NAME                CAPACITY    ACCESSMODES   RECLAIMPOLICY   STATUS      CLAIM     STORAGECLASS    REASON    AGE
local-pv-ce05be60   1024220Ki   RWO           Delete          Available             local-storage             26s

$ kubectl describe pv local-pv-ce05be60 
Name:  local-pv-ce05be60
...
StorageClass: local-storage
Status:  Available
Claim:  
Reclaim Policy: Delete
Access Modes: RWO
Capacity: 1024220Ki
NodeAffinity:
  Required Terms:
      Term 0:  kubernetes.io/hostname in [node-1]
Message: 
Source:
    Type: LocalVolume (a persistent volume backed by local storage on a node)
    Path: /mnt/disks/vol1
</code></pre>

<p>这个PV里的各种定义，比如StorageClass的名字、本地磁盘挂载点的位置，都可以通过provisioner的<a href="https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume/helm" target="_blank">配置文件指定</a>。当然，provisioner也会负责前面提到的PV的删除工作。</p>

<p>而这个provisioner本身，其实也是一个我们前面提到过的<a href="https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume" target="_blank">External Provisioner</a>，它的部署方法，在<a href="https://github.com/kubernetes-incubator/external-storage/tree/master/local-volume#option-1-using-the-local-volume-static-provisioner" target="_blank">对应的文档里</a>有详细描述。这部分内容，就留给你课后自行探索了。</p>

<h2 id="总结">总结</h2>

<p>在今天这篇文章中，我为你详细介绍了Kubernetes里Local Persistent Volume的实现方式。</p>

<p>可以看到，正是通过PV和PVC，以及StorageClass这套存储体系，这个后来新添加的持久化存储方案，对Kubernetes已有用户的影响，几乎可以忽略不计。作为用户，你的Pod的YAML和PVC的YAML，并没有任何特殊的改变，这个特性所有的实现只会影响到PV的处理，也就是由运维人员负责的那部分工作。</p>

<p>而这，正是这套存储体系带来的“解耦”的好处。</p>

<p>其实，Kubernetes很多看起来比较“繁琐”的设计（比如“声明式API”，以及我今天讲解的“PV、PVC体系”）的主要目的，都是希望为开发者提供更多的“可扩展性”，给使用者带来更多的“稳定性”和“安全感”。这两个能力的高低，是衡量开源基础设施项目水平的重要标准。</p>

<h2 id="思考题">思考题</h2>

<p>正是由于需要使用“延迟绑定”这个特性，Local Persistent Volume目前还不能支持Dynamic Provisioning。你是否能说出，为什么“延迟绑定”会跟Dynamic Provisioning有冲突呢？</p>

<p>感谢你的收听，欢迎你给我留言，也欢迎分享给更多的朋友一起阅读。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#c0acacacf9f4f1f1f0f780a7ada1a9aceea3afad" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f159903acac7771',t:'MTczNDA4OTE3MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>