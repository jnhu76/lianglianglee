<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=40&#32;Kubernetes的资源模型与资源管理>
        <link rel="icon" href="/static/favicon.png">
        <title>40 Kubernetes的资源模型与资源管理 </title>
        
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
                            <h1 id="title" data-id="40 Kubernetes的资源模型与资源管理" class="title">40 Kubernetes的资源模型与资源管理</h1>
                            <div><p>你好，我是张磊。今天我和你分享的主题是：Kubernetes的资源模型与资源管理。</p>

<p>作为一个容器集群编排与管理项目，Kubernetes为用户提供的基础设施能力，不仅包括了我在前面为你讲述的应用定义和描述的部分，还包括了对应用的资源管理和调度的处理。那么，从今天这篇文章开始，我就来为你详细讲解一下后面这部分内容。</p>

<p>而作为Kubernetes的资源管理与调度部分的基础，我们要从它的资源模型开始说起。</p>

<p>我在前面的文章中已经提到过，在Kubernetes里，Pod是最小的原子调度单位。这也就意味着，所有跟调度和资源管理相关的属性都应该是属于Pod对象的字段。而这其中最重要的部分，就是Pod的CPU和内存配置，如下所示：</p>

<pre><code>apiVersion: v1
kind: Pod
metadata:
  name: frontend
spec:
  containers:
  - name: db
    image: mysql
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: &quot;password&quot;
    resources:
      requests:
        memory: &quot;64Mi&quot;
        cpu: &quot;250m&quot;
      limits:
        memory: &quot;128Mi&quot;
        cpu: &quot;500m&quot;
  - name: wp
    image: wordpress
    resources:
      requests:
        memory: &quot;64Mi&quot;
        cpu: &quot;250m&quot;
      limits:
        memory: &quot;128Mi&quot;
        cpu: &quot;500m&quot;
</code></pre>

<blockquote>
<p>备注：关于哪些属性属于Pod对象，而哪些属性属于Container，你可以在回顾一下第14篇文章《深入解析Pod对象（一）：基本概念》中的相关内容。</p>
</blockquote>

<p>在Kubernetes中，像CPU这样的资源被称作“可压缩资源”（compressible resources）。它的典型特点是，当可压缩资源不足时，Pod只会“饥饿”，但不会退出。</p>

<p>而像内存这样的资源，则被称作“不可压缩资源（incompressible resources）。当不可压缩资源不足时，Pod就会因为OOM（Out-Of-Memory）被内核杀掉。</p>

<p>而由于Pod可以由多个Container组成，所以CPU和内存资源的限额，是要配置在每个Container的定义上的。这样，Pod整体的资源配置，就由这些Container的配置值累加得到。</p>

<p>其中，Kubernetes里为CPU设置的单位是“CPU的个数”。比如，cpu=1指的就是，这个Pod的CPU限额是1个CPU。当然，具体“1个CPU”在宿主机上如何解释，是1个CPU核心，还是1个vCPU，还是1个CPU的超线程（Hyperthread），完全取决于宿主机的CPU实现方式。Kubernetes只负责保证Pod能够使用到“1个CPU”的计算能力。</p>

<p>此外，Kubernetes允许你将CPU限额设置为分数，比如在我们的例子里，CPU limits的值就是500m。所谓500m，指的就是500 millicpu，也就是0.5个CPU的意思。这样，这个Pod就会被分配到1个CPU一半的计算能力。</p>

<p>当然，<strong>你也可以直接把这个配置写成cpu=0.5。但在实际使用时，我还是推荐你使用500m的写法，毕竟这才是Kubernetes内部通用的CPU表示方式。</strong></p>

<p>而对于内存资源来说，它的单位自然就是bytes。Kubernetes支持你使用Ei、Pi、Ti、Gi、Mi、Ki（或者E、P、T、G、M、K）的方式来作为bytes的值。比如，在我们的例子里，Memory requests的值就是64MiB (2的26次方bytes) 。这里要注意区分MiB（mebibyte）和MB（megabyte）的区别。</p>

<blockquote>
<p>备注：1Mi=1024*1024；1M=1000*1000</p>
</blockquote>

<p>此外，不难看到，<strong>Kubernetes里Pod的CPU和内存资源，实际上还要分为limits和requests两种情况</strong>，如下所示：</p>

<pre><code>spec.containers[].resources.limits.cpu
spec.containers[].resources.limits.memory
spec.containers[].resources.requests.cpu
spec.containers[].resources.requests.memory
</code></pre>

<p>这两者的区别其实非常简单：在调度的时候，kube-scheduler只会按照requests的值进行计算。而在真正设置Cgroups限制的时候，kubelet则会按照limits的值来进行设置。</p>

<p>更确切地说，当你指定了requests.cpu=250m之后，相当于将Cgroups的cpu.shares的值设置为(<sup>250</sup>&frasl;<sub>1000</sub>)*1024。而当你没有设置requests.cpu的时候，cpu.shares默认则是1024。这样，Kubernetes就通过cpu.shares完成了对CPU时间的按比例分配。</p>

<p>而如果你指定了limits.cpu=500m之后，则相当于将Cgroups的cpu.cfs_quota_us的值设置为(<sup>500</sup>&frasl;<sub>1000</sub>)*100ms，而cpu.cfs_period_us的值始终是100ms。这样，Kubernetes就为你设置了这个容器只能用到CPU的50%。</p>

<p>而对于内存来说，当你指定了limits.memory=128Mi之后，相当于将Cgroups的memory.limit_in_bytes设置为128 * 1024 * 1024。而需要注意的是，在调度的时候，调度器只会使用requests.memory=64Mi来进行判断。</p>

<p><strong>Kubernetes这种对CPU和内存资源限额的设计，实际上参考了Borg论文中对“动态资源边界”的定义</strong>，既：容器化作业在提交时所设置的资源边界，并不一定是调度系统所必须严格遵守的，这是因为在实际场景中，大多数作业使用到的资源其实远小于它所请求的资源限额。</p>

<p>基于这种假设，Borg在作业被提交后，会主动减小它的资源限额配置，以便容纳更多的作业、提升资源利用率。而当作业资源使用量增加到一定阈值时，Borg会通过“快速恢复”过程，还原作业原始的资源限额，防止出现异常情况。</p>

<p>而Kubernetes的requests+limits的做法，其实就是上述思路的一个简化版：用户在提交Pod时，可以声明一个相对较小的requests值供调度器使用，而Kubernetes真正设置给容器Cgroups的，则是相对较大的limits值。不难看到，这跟Borg的思路相通的。</p>

<p>在理解了Kubernetes资源模型的设计之后，我再来和你谈谈Kubernetes里的QoS模型。在Kubernetes中，不同的requests和limits的设置方式，其实会将这个Pod划分到不同的QoS级别当中。</p>

<p><strong>当Pod里的每一个Container都同时设置了requests和limits，并且requests和limits值相等的时候，这个Pod就属于Guaranteed类别</strong>，如下所示：</p>

<pre><code>apiVersion: v1
kind: Pod
metadata:
  name: qos-demo
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-ctr
    image: nginx
    resources:
      limits:
        memory: &quot;200Mi&quot;
        cpu: &quot;700m&quot;
      requests:
        memory: &quot;200Mi&quot;
        cpu: &quot;700m&quot;
</code></pre>

<p>当这个Pod创建之后，它的qosClass字段就会被Kubernetes自动设置为Guaranteed。需要注意的是，当Pod仅设置了limits没有设置requests的时候，Kubernetes会自动为它设置与limits相同的requests值，所以，这也属于Guaranteed情况。</p>

<p><strong>而当Pod不满足Guaranteed的条件，但至少有一个Container设置了requests。那么这个Pod就会被划分到Burstable类别</strong>。比如下面这个例子：</p>

<pre><code>apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-2
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-2-ctr
    image: nginx
    resources:
      limits
        memory: &quot;200Mi&quot;
      requests:
        memory: &quot;100Mi&quot;
</code></pre>

<p><strong>而如果一个Pod既没有设置requests，也没有设置limits，那么它的QoS类别就是BestEffort</strong>。比如下面这个例子：</p>

<pre><code>apiVersion: v1
kind: Pod
metadata:
  name: qos-demo-3
  namespace: qos-example
spec:
  containers:
  - name: qos-demo-3-ctr
    image: nginx
</code></pre>

<p>那么，Kubernetes为Pod设置这样三种QoS类别，具体有什么作用呢？</p>

<p>实际上，<strong>QoS划分的主要应用场景，是当宿主机资源紧张的时候，kubelet对Pod进行Eviction（即资源回收）时需要用到的。</strong></p>

<p>具体地说，当Kubernetes所管理的宿主机上不可压缩资源短缺时，就有可能触发Eviction。比如，可用内存（memory.available）、可用的宿主机磁盘空间（nodefs.available），以及容器运行时镜像存储空间（imagefs.available）等等。</p>

<p>目前，Kubernetes为你设置的Eviction的默认阈值如下所示：</p>

<pre><code>memory.available&lt;100Mi
nodefs.available&lt;10%
nodefs.inodesFree&lt;5%
imagefs.available&lt;15%
</code></pre>

<p>当然，上述各个触发条件在kubelet里都是可配置的。比如下面这个例子：</p>

<pre><code>kubelet --eviction-hard=imagefs.available&lt;10%,memory.available&lt;500Mi,nodefs.available&lt;5%,nodefs.inodesFree&lt;5% --eviction-soft=imagefs.available&lt;30%,nodefs.available&lt;10% --eviction-soft-grace-period=imagefs.available=2m,nodefs.available=2m --eviction-max-pod-grace-period=600
</code></pre>

<p>在这个配置中，你可以看到<strong>Eviction在Kubernetes里其实分为Soft和Hard两种模式</strong>。</p>

<p>其中，Soft Eviction允许你为Eviction过程设置一段“优雅时间”，比如上面例子里的imagefs.available=2m，就意味着当imagefs不足的阈值达到2分钟之后，kubelet才会开始Eviction的过程。</p>

<p>而Hard Eviction模式下，Eviction过程就会在阈值达到之后立刻开始。</p>

<blockquote>
<p>Kubernetes计算Eviction阈值的数据来源，主要依赖于从Cgroups读取到的值，以及使用cAdvisor监控到的数据。</p>
</blockquote>

<p>当宿主机的Eviction阈值达到后，就会进入MemoryPressure或者DiskPressure状态，从而避免新的Pod被调度到这台宿主机上。</p>

<p>而当Eviction发生的时候，kubelet具体会挑选哪些Pod进行删除操作，就需要参考这些Pod的QoS类别了。</p>

<ul>
<li>首当其冲的，自然是BestEffort类别的Pod。</li>
<li>其次，是属于Burstable类别、并且发生“饥饿”的资源使用量已经超出了requests的Pod。</li>
<li>最后，才是Guaranteed类别。并且，Kubernetes会保证只有当Guaranteed类别的Pod的资源使用量超过了其limits的限制，或者宿主机本身正处于Memory Pressure状态时，Guaranteed的Pod才可能被选中进行Eviction操作。</li>
</ul>

<p>当然，对于同QoS类别的Pod来说，Kubernetes还会根据Pod的优先级来进行进一步地排序和选择。</p>

<p>在理解了Kubernetes里的QoS类别的设计之后，我再来为你讲解一下Kubernetes里一个非常有用的特性：cpuset的设置。</p>

<p>我们知道，在使用容器的时候，你可以通过设置cpuset把容器绑定到某个CPU的核上，而不是像cpushare那样共享CPU的计算能力。</p>

<p>这种情况下，由于操作系统在CPU之间进行上下文切换的次数大大减少，容器里应用的性能会得到大幅提升。事实上，<strong>cpuset方式，是生产环境里部署在线应用类型的Pod时，非常常用的一种方式。</strong></p>

<p>可是，这样的需求在Kubernetes里又该如何实现呢？</p>

<p>其实非常简单。</p>

<ul>
<li>首先，你的Pod必须是Guaranteed的QoS类型；</li>
<li>然后，你只需要将Pod的CPU资源的requests和limits设置为同一个相等的整数值即可。</li>
</ul>

<p>比如下面这个例子：</p>

<pre><code>spec:
  containers:
  - name: nginx
    image: nginx
    resources:
      limits:
        memory: &quot;200Mi&quot;
        cpu: &quot;2&quot;
      requests:
        memory: &quot;200Mi&quot;
        cpu: &quot;2&quot;
</code></pre>

<p>这时候，该Pod就会被绑定在2个独占的CPU核上。当然，具体是哪两个CPU核，是由kubelet为你分配的。</p>

<p>以上，就是Kubernetes的资源模型和QoS类别相关的主要内容。</p>

<h2 id="总结">总结</h2>

<p>在本篇文章中，我先为你详细讲解了Kubernetes里对资源的定义方式和资源模型的设计。然后，我为你讲述了Kubernetes里对Pod进行Eviction的具体策略和实践方式。</p>

<p>正是基于上述讲述，在实际的使用中，我强烈建议你将DaemonSet的Pod都设置为Guaranteed的QoS类型。否则，一旦DaemonSet的Pod被回收，它又会立即在原宿主机上被重建出来，这就使得前面资源回收的动作，完全没有意义了。</p>

<h2 id="思考题">思考题</h2>

<p>为什么宿主机进入MemoryPressure或者DiskPressure状态后，新的Pod就不会被调度到这台宿主机上呢？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#4d21212174797c7c7d7a0d2a202c2421632e2220" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f159b654f437771',t:'MTczNDA4OTI2OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>