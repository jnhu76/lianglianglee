<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=09&#32;&#32;资源隔离：为什么构建容器需要&#32;Namespace&#32;？>
        <link rel="icon" href="/static/favicon.png">
        <title>09  资源隔离：为什么构建容器需要 Namespace ？ </title>
        
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
                        <a class="menu-item" id="00 溯本求源，吃透 Docker！.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/00%20%e6%ba%af%e6%9c%ac%e6%b1%82%e6%ba%90%ef%bc%8c%e5%90%83%e9%80%8f%20Docker%ef%bc%81.md">00 溯本求源，吃透 Docker！.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  Docker 安装：入门案例带你了解容器技术原理.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/01%20%20Docker%20%e5%ae%89%e8%a3%85%ef%bc%9a%e5%85%a5%e9%97%a8%e6%a1%88%e4%be%8b%e5%b8%a6%e4%bd%a0%e4%ba%86%e8%a7%a3%e5%ae%b9%e5%99%a8%e6%8a%80%e6%9c%af%e5%8e%9f%e7%90%86.md">01  Docker 安装：入门案例带你了解容器技术原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  核心概念：镜像、容器、仓库，彻底掌握 Docker 架构核心设计理念.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/02%20%20%e6%a0%b8%e5%bf%83%e6%a6%82%e5%bf%b5%ef%bc%9a%e9%95%9c%e5%83%8f%e3%80%81%e5%ae%b9%e5%99%a8%e3%80%81%e4%bb%93%e5%ba%93%ef%bc%8c%e5%bd%bb%e5%ba%95%e6%8e%8c%e6%8f%a1%20Docker%20%e6%9e%b6%e6%9e%84%e6%a0%b8%e5%bf%83%e8%ae%be%e8%ae%a1%e7%90%86%e5%bf%b5.md">02  核心概念：镜像、容器、仓库，彻底掌握 Docker 架构核心设计理念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  镜像使用：Docker 环境下如何配置你的镜像？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/03%20%20%e9%95%9c%e5%83%8f%e4%bd%bf%e7%94%a8%ef%bc%9aDocker%20%e7%8e%af%e5%a2%83%e4%b8%8b%e5%a6%82%e4%bd%95%e9%85%8d%e7%bd%ae%e4%bd%a0%e7%9a%84%e9%95%9c%e5%83%8f%ef%bc%9f.md">03  镜像使用：Docker 环境下如何配置你的镜像？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  容器操作：得心应手掌握 Docker 容器基本操作.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/04%20%20%e5%ae%b9%e5%99%a8%e6%93%8d%e4%bd%9c%ef%bc%9a%e5%be%97%e5%bf%83%e5%ba%94%e6%89%8b%e6%8e%8c%e6%8f%a1%20Docker%20%e5%ae%b9%e5%99%a8%e5%9f%ba%e6%9c%ac%e6%93%8d%e4%bd%9c.md">04  容器操作：得心应手掌握 Docker 容器基本操作.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  仓库访问：怎样搭建属于你的私有仓库？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/05%20%20%e4%bb%93%e5%ba%93%e8%ae%bf%e9%97%ae%ef%bc%9a%e6%80%8e%e6%a0%b7%e6%90%ad%e5%bb%ba%e5%b1%9e%e4%ba%8e%e4%bd%a0%e7%9a%84%e7%a7%81%e6%9c%89%e4%bb%93%e5%ba%93%ef%bc%9f.md">05  仓库访问：怎样搭建属于你的私有仓库？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  最佳实践：如何在生产中编写最优 Dockerfile？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/06%20%20%e6%9c%80%e4%bd%b3%e5%ae%9e%e8%b7%b5%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9c%a8%e7%94%9f%e4%ba%a7%e4%b8%ad%e7%bc%96%e5%86%99%e6%9c%80%e4%bc%98%20Dockerfile%ef%bc%9f.md">06  最佳实践：如何在生产中编写最优 Dockerfile？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  Docker 安全：基于内核的弱隔离系统如何保障安全性？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/07%20%20Docker%20%e5%ae%89%e5%85%a8%ef%bc%9a%e5%9f%ba%e4%ba%8e%e5%86%85%e6%a0%b8%e7%9a%84%e5%bc%b1%e9%9a%94%e7%a6%bb%e7%b3%bb%e7%bb%9f%e5%a6%82%e4%bd%95%e4%bf%9d%e9%9a%9c%e5%ae%89%e5%85%a8%e6%80%a7%ef%bc%9f.md">07  Docker 安全：基于内核的弱隔离系统如何保障安全性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  容器监控：容器监控原理及 cAdvisor 的安装与使用.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/08%20%20%e5%ae%b9%e5%99%a8%e7%9b%91%e6%8e%a7%ef%bc%9a%e5%ae%b9%e5%99%a8%e7%9b%91%e6%8e%a7%e5%8e%9f%e7%90%86%e5%8f%8a%20cAdvisor%20%e7%9a%84%e5%ae%89%e8%a3%85%e4%b8%8e%e4%bd%bf%e7%94%a8.md">08  容器监控：容器监控原理及 cAdvisor 的安装与使用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  资源隔离：为什么构建容器需要 Namespace ？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/09%20%20%e8%b5%84%e6%ba%90%e9%9a%94%e7%a6%bb%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e6%9e%84%e5%bb%ba%e5%ae%b9%e5%99%a8%e9%9c%80%e8%a6%81%20Namespace%20%ef%bc%9f.md">09  资源隔离：为什么构建容器需要 Namespace ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  资源限制：如何通过 Cgroups 机制实现资源限制？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/10%20%20%e8%b5%84%e6%ba%90%e9%99%90%e5%88%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20Cgroups%20%e6%9c%ba%e5%88%b6%e5%ae%9e%e7%8e%b0%e8%b5%84%e6%ba%90%e9%99%90%e5%88%b6%ef%bc%9f.md">10  资源限制：如何通过 Cgroups 机制实现资源限制？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  组件组成：剖析 Docker 组件作用及其底层工作原理.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/11%20%20%e7%bb%84%e4%bb%b6%e7%bb%84%e6%88%90%ef%bc%9a%e5%89%96%e6%9e%90%20Docker%20%e7%bb%84%e4%bb%b6%e4%bd%9c%e7%94%a8%e5%8f%8a%e5%85%b6%e5%ba%95%e5%b1%82%e5%b7%a5%e4%bd%9c%e5%8e%9f%e7%90%86.md">11  组件组成：剖析 Docker 组件作用及其底层工作原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  网络模型：剖析 Docker 网络实现及 Libnetwork 底层原理.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/12%20%20%e7%bd%91%e7%bb%9c%e6%a8%a1%e5%9e%8b%ef%bc%9a%e5%89%96%e6%9e%90%20Docker%20%e7%bd%91%e7%bb%9c%e5%ae%9e%e7%8e%b0%e5%8f%8a%20Libnetwork%20%e5%ba%95%e5%b1%82%e5%8e%9f%e7%90%86.md">12  网络模型：剖析 Docker 网络实现及 Libnetwork 底层原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  数据存储：剖析 Docker 卷与持久化数据存储的底层原理.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/13%20%20%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%ef%bc%9a%e5%89%96%e6%9e%90%20Docker%20%e5%8d%b7%e4%b8%8e%e6%8c%81%e4%b9%85%e5%8c%96%e6%95%b0%e6%8d%ae%e5%ad%98%e5%82%a8%e7%9a%84%e5%ba%95%e5%b1%82%e5%8e%9f%e7%90%86.md">13  数据存储：剖析 Docker 卷与持久化数据存储的底层原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  文件存储驱动：AUFS 文件系统原理及生产环境的最佳配置.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/14%20%20%e6%96%87%e4%bb%b6%e5%ad%98%e5%82%a8%e9%a9%b1%e5%8a%a8%ef%bc%9aAUFS%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%8e%9f%e7%90%86%e5%8f%8a%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%e7%9a%84%e6%9c%80%e4%bd%b3%e9%85%8d%e7%bd%ae.md">14  文件存储驱动：AUFS 文件系统原理及生产环境的最佳配置.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  文件存储驱动：Devicemapper 文件系统原理及生产环境的最佳配置.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/15%20%20%e6%96%87%e4%bb%b6%e5%ad%98%e5%82%a8%e9%a9%b1%e5%8a%a8%ef%bc%9aDevicemapper%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%8e%9f%e7%90%86%e5%8f%8a%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%e7%9a%84%e6%9c%80%e4%bd%b3%e9%85%8d%e7%bd%ae.md">15  文件存储驱动：Devicemapper 文件系统原理及生产环境的最佳配置.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  文件存储驱动：OverlayFS 文件系统原理及生产环境的最佳配置.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/16%20%20%e6%96%87%e4%bb%b6%e5%ad%98%e5%82%a8%e9%a9%b1%e5%8a%a8%ef%bc%9aOverlayFS%20%e6%96%87%e4%bb%b6%e7%b3%bb%e7%bb%9f%e5%8e%9f%e7%90%86%e5%8f%8a%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%e7%9a%84%e6%9c%80%e4%bd%b3%e9%85%8d%e7%bd%ae.md">16  文件存储驱动：OverlayFS 文件系统原理及生产环境的最佳配置.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  原理实践：自己动手使用 Golang 开发 Docker（上）.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/17%20%20%e5%8e%9f%e7%90%86%e5%ae%9e%e8%b7%b5%ef%bc%9a%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%e4%bd%bf%e7%94%a8%20Golang%20%e5%bc%80%e5%8f%91%20Docker%ef%bc%88%e4%b8%8a%ef%bc%89.md">17  原理实践：自己动手使用 Golang 开发 Docker（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  原理实践：自己动手使用 Golang 开发 Docker（下）.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/18%20%20%e5%8e%9f%e7%90%86%e5%ae%9e%e8%b7%b5%ef%bc%9a%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%e4%bd%bf%e7%94%a8%20Golang%20%e5%bc%80%e5%8f%91%20Docker%ef%bc%88%e4%b8%8b%ef%bc%89.md">18  原理实践：自己动手使用 Golang 开发 Docker（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  如何使用 Docker Compose 解决开发环境的依赖？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/19%20%20%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Docker%20Compose%20%e8%a7%a3%e5%86%b3%e5%bc%80%e5%8f%91%e7%8e%af%e5%a2%83%e7%9a%84%e4%be%9d%e8%b5%96%ef%bc%9f.md">19  如何使用 Docker Compose 解决开发环境的依赖？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  如何在生产环境中使用 Docker Swarm 调度容器？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/20%20%20%e5%a6%82%e4%bd%95%e5%9c%a8%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%e4%b8%ad%e4%bd%bf%e7%94%a8%20Docker%20Swarm%20%e8%b0%83%e5%ba%a6%e5%ae%b9%e5%99%a8%ef%bc%9f.md">20  如何在生产环境中使用 Docker Swarm 调度容器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  如何使 Docker 和 Kubernetes 结合发挥容器的最大价值？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/21%20%20%e5%a6%82%e4%bd%95%e4%bd%bf%20Docker%20%e5%92%8c%20Kubernetes%20%e7%bb%93%e5%90%88%e5%8f%91%e6%8c%a5%e5%ae%b9%e5%99%a8%e7%9a%84%e6%9c%80%e5%a4%a7%e4%bb%b7%e5%80%bc%ef%bc%9f.md">21  如何使 Docker 和 Kubernetes 结合发挥容器的最大价值？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  多阶级构建：Docker 下如何实现镜像多阶级构建？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/22%20%20%e5%a4%9a%e9%98%b6%e7%ba%a7%e6%9e%84%e5%bb%ba%ef%bc%9aDocker%20%e4%b8%8b%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e9%95%9c%e5%83%8f%e5%a4%9a%e9%98%b6%e7%ba%a7%e6%9e%84%e5%bb%ba%ef%bc%9f.md">22  多阶级构建：Docker 下如何实现镜像多阶级构建？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  DevOps：容器化后如何通过 DevOps 提高协作效能？.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/23%20%20DevOps%ef%bc%9a%e5%ae%b9%e5%99%a8%e5%8c%96%e5%90%8e%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%20DevOps%20%e6%8f%90%e9%ab%98%e5%8d%8f%e4%bd%9c%e6%95%88%e8%83%bd%ef%bc%9f.md">23  DevOps：容器化后如何通过 DevOps 提高协作效能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  CICD：容器化后如何实现持续集成与交付？（上）.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/24%20%20CICD%ef%bc%9a%e5%ae%b9%e5%99%a8%e5%8c%96%e5%90%8e%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%e4%b8%8e%e4%ba%a4%e4%bb%98%ef%bc%9f%ef%bc%88%e4%b8%8a%ef%bc%89.md">24  CICD：容器化后如何实现持续集成与交付？（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  CICD：容器化后如何实现持续集成与交付？（下）.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/25%20%20CICD%ef%bc%9a%e5%ae%b9%e5%99%a8%e5%8c%96%e5%90%8e%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e6%8c%81%e7%bb%ad%e9%9b%86%e6%88%90%e4%b8%8e%e4%ba%a4%e4%bb%98%ef%bc%9f%ef%bc%88%e4%b8%8b%ef%bc%89.md">25  CICD：容器化后如何实现持续集成与交付？（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 结束语  展望未来：Docker 的称霸之路.md" href="/%e4%b8%93%e6%a0%8f/%e7%94%b1%e6%b5%85%e5%85%a5%e6%b7%b1%e5%90%83%e9%80%8f%20Docker-%e5%ae%8c/26%20%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e5%b1%95%e6%9c%9b%e6%9c%aa%e6%9d%a5%ef%bc%9aDocker%20%e7%9a%84%e7%a7%b0%e9%9c%b8%e4%b9%8b%e8%b7%af.md">26 结束语  展望未来：Docker 的称霸之路.md</a>
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
                            <h1 id="title" data-id="09  资源隔离：为什么构建容器需要 Namespace ？" class="title">09  资源隔离：为什么构建容器需要 Namespace ？</h1>
                            <div><p>我们知道， Docker 是使用 Linux 的 Namespace 技术实现各种资源隔离的。那么究竟什么是 Namespace，各种 Namespace 都有什么作用，为什么 Docker 需要 Namespace呢？下面我带你一一揭秘。</p>

<p>首先我们来了解一下什么是 Namespace。</p>

<h3 id="什么是-namespace">什么是 Namespace？</h3>

<p>下面是 Namespace 的维基百科定义：</p>

<blockquote>
<p>Namespace 是 Linux 内核的一项功能，该功能对内核资源进行分区，以使一组进程看到一组资源，而另一组进程看到另一组资源。Namespace 的工作方式通过为一组资源和进程设置相同的 Namespace 而起作用，但是这些 Namespace 引用了不同的资源。资源可能存在于多个 Namespace 中。这些资源可以是进程 ID、主机名、用户 ID、文件名、与网络访问相关的名称和进程间通信。</p>
</blockquote>

<p>简单来说，Namespace 是 Linux 内核的一个特性，该特性可以实现在同一主机系统中，对进程 ID、主机名、用户 ID、文件名、网络和进程间通信等资源的隔离。Docker 利用 Linux 内核的 Namespace 特性，实现了每个容器的资源相互隔离，从而保证容器内部只能访问到自己 Namespace 的资源。</p>

<p>最新的 Linux 5.6 内核中提供了 8 种类型的 Namespace：</p>

<table>
<thead>
<tr>
<th>Namespace 名称</th>
<th>作用</th>
<th>内核版本</th>
</tr>
</thead>

<tbody>
<tr>
<td>Mount（mnt）</td>
<td>隔离挂载点</td>
<td>2.4.19</td>
</tr>

<tr>
<td>Process ID (pid)</td>
<td>隔离进程 ID</td>
<td>2.6.24</td>
</tr>

<tr>
<td>Network (net)</td>
<td>隔离网络设备，端口号等</td>
<td>2.6.29</td>
</tr>

<tr>
<td>Interprocess Communication (ipc)</td>
<td>隔离 System V IPC 和 POSIX message queues</td>
<td>2.6.19</td>
</tr>

<tr>
<td>UTS Namespace(uts)</td>
<td>隔离主机名和域名</td>
<td>2.6.19</td>
</tr>

<tr>
<td>User Namespace (user)</td>
<td>隔离用户和用户组</td>
<td>3.8</td>
</tr>

<tr>
<td>Control group (cgroup) Namespace</td>
<td>隔离 Cgroups 根目录</td>
<td>4.6</td>
</tr>

<tr>
<td>Time Namespace</td>
<td>隔离系统时间</td>
<td>5.6</td>
</tr>
</tbody>
</table>
<p>虽然 Linux 内核提供了8种 Namespace，但是最新版本的 Docker 只使用了其中的前6 种，分别为Mount Namespace、PID Namespace、Net Namespace、IPC Namespace、UTS Namespace、User Namespace。</p>

<p>下面，我们详细了解下 Docker 使用的 6 种 Namespace的作用分别是什么。</p>

<h3 id="各种-namespace-的作用">各种 Namespace 的作用？</h3>

<h4 id="1-mount-namespace">（1）Mount Namespace</h4>

<p>Mount Namespace 是 Linux 内核实现的第一个 Namespace，从内核的 2.4.19 版本开始加入。它可以用来隔离不同的进程或进程组看到的挂载点。通俗地说，就是可以实现在不同的进程中看到不同的挂载目录。使用 Mount Namespace 可以实现容器内只能看到自己的挂载信息，在容器内的挂载操作不会影响主机的挂载目录。</p>

<p>下面我们通过一个实例来演示下 Mount Namespace。在演示之前，我们先来认识一个命令行工具 unshare。unshare 是 util-linux 工具包中的一个工具，CentOS 7 系统默认已经集成了该工具，<strong>使用 unshare 命令可以实现创建并访问不同类型的 Namespace</strong>。</p>

<p>首先我们使用以下命令创建一个 bash 进程并且新建一个 Mount Namespace：</p>

<pre><code>$ sudo unshare --mount --fork /bin/bash

[root@centos7 centos]#
</code></pre>

<p>执行完上述命令后，这时我们已经在主机上创建了一个新的 Mount Namespace，并且当前命令行窗口加入了新创建的 Mount Namespace。下面我通过一个例子来验证下，在独立的 Mount Namespace 内创建挂载目录是不影响主机的挂载目录的。</p>

<p>首先在 /tmp 目录下创建一个目录。</p>

<pre><code>[root@centos7 centos]# mkdir /tmp/tmpfs
</code></pre>

<p>创建好目录后使用 mount 命令挂载一个 tmpfs 类型的目录。命令如下：</p>

<pre><code>[root@centos7 centos]# mount -t tmpfs -o size=20m tmpfs /tmp/tmpfs
</code></pre>

<p>然后使用 df 命令查看一下已经挂载的目录信息：</p>

<pre><code>[root@centos7 centos]# df -h

Filesystem      Size  Used Avail Use% Mounted on

/dev/vda1       500G  1.4G  499G   1% /

devtmpfs         16G     0   16G   0% /dev

tmpfs            16G     0   16G   0% /dev/shm

tmpfs            16G     0   16G   0% /sys/fs/cgroup

tmpfs            16G   57M   16G   1% /run

tmpfs           3.2G     0  3.2G   0% /run/user/1000

tmpfs            20M     0   20M   0% /tmp/tmpfs
</code></pre>

<p>可以看到 /tmp/tmpfs 目录已经被正确挂载。为了验证主机上并没有挂载此目录，我们新打开一个命令行窗口，同样执行 df 命令查看主机的挂载信息：</p>

<pre><code>[centos@centos7 ~]$ df -h

Filesystem      Size  Used Avail Use% Mounted on

devtmpfs         16G     0   16G   0% /dev

tmpfs            16G     0   16G   0% /dev/shm

tmpfs            16G   57M   16G   1% /run

tmpfs            16G     0   16G   0% /sys/fs/cgroup

/dev/vda1       500G  1.4G  499G   1% /

tmpfs           3.2G     0  3.2G   0% /run/user/1000
</code></pre>

<p>通过上面输出可以看到主机上并没有挂载 /tmp/tmpfs，可见我们独立的 Mount Namespace 中执行 mount 操作并不会影响主机。</p>

<p>为了进一步验证我们的想法，我们继续在当前命令行窗口查看一下当前进程的 Namespace 信息，命令如下：</p>

<pre><code>[root@centos7 centos]# ls -l /proc/self/ns/

total 0

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 ipc -&gt; ipc:[4026531839]

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 mnt -&gt; mnt:[4026532239]

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 net -&gt; net:[4026531956]

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 pid -&gt; pid:[4026531836]

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 user -&gt; user:[4026531837]

lrwxrwxrwx. 1 root root 0 Sep  4 08:20 uts -&gt; uts:[4026531838]
</code></pre>

<p>然后新打开一个命令行窗口，使用相同的命令查看一下主机上的 Namespace 信息：</p>

<pre><code>[centos@centos7 ~]$ ls -l /proc/self/ns/

total 0

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 ipc -&gt; ipc:[4026531839]

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 mnt -&gt; mnt:[4026531840]

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 net -&gt; net:[4026531956]

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 pid -&gt; pid:[4026531836]

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 user -&gt; user:[4026531837]

lrwxrwxrwx. 1 centos centos 0 Sep  4 08:20 uts -&gt; uts:[4026531838]
</code></pre>

<p>通过对比两次命令的输出结果，我们可以看到，除了 Mount Namespace 的 ID 值不一样外，其他Namespace 的 ID 值均一致。</p>

<p>通过以上结果我们可以得出结论，<strong>使用 unshare 命令可以新建 Mount Namespace，并且在新建的 Mount Namespace 内 mount 是和外部完全隔离的。</strong></p>

<h4 id="2-pid-namespace">（2）PID Namespace</h4>

<p>PID Namespace 的作用是用来隔离进程。在不同的 PID Namespace 中，进程可以拥有相同的 PID 号，利用 PID Namespace 可以实现每个容器的主进程为 1 号进程，而容器内的进程在主机上却拥有不同的PID。例如一个进程在主机上 PID 为 122，使用 PID Namespace 可以实现该进程在容器内看到的 PID 为 1。</p>

<p>下面我们通过一个实例来演示下 PID Namespace的作用。首先我们使用以下命令创建一个 bash 进程，并且新建一个 PID Namespace：</p>

<pre><code>$ sudo unshare --pid --fork --mount-proc /bin/bash

[root@centos7 centos]#
</code></pre>

<p>执行完上述命令后，我们在主机上创建了一个新的 PID Namespace，并且当前命令行窗口加入了新创建的 PID Namespace。在当前的命令行窗口使用 ps aux 命令查看一下进程信息：</p>

<pre><code>[root@centos7 centos]# ps aux

USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND

root         1  0.0  0.0 115544  2004 pts/0    S    10:57   0:00 bash

root        10  0.0  0.0 155444  1764 pts/0    R+   10:59   0:00 ps aux
</code></pre>

<p>通过上述命令输出结果可以看到当前 Namespace 下 bash 为 1 号进程，而且我们也看不到主机上的其他进程信息。</p>

<h4 id="3-uts-namespace">（3）UTS Namespace</h4>

<p>UTS Namespace 主要是用来隔离主机名的，它允许每个 UTS Namespace 拥有一个独立的主机名。例如我们的主机名称为 docker，使用 UTS Namespace 可以实现在容器内的主机名称为 lagoudocker 或者其他任意自定义主机名。</p>

<p>同样我们通过一个实例来验证下 UTS Namespace 的作用，首先我们使用 unshare 命令来创建一个 UTS Namespace：</p>

<pre><code>$ sudo unshare --uts --fork /bin/bash

[root@centos7 centos]#
</code></pre>

<p>创建好 UTS Namespace 后，当前命令行窗口已经处于一个独立的 UTS Namespace 中，下面我们使用 hostname 命令（hostname 可以用来查看主机名称）设置一下主机名：</p>

<pre><code>root@centos7 centos]# hostname -b lagoudocker
</code></pre>

<p>然后再查看一下主机名：</p>

<pre><code>[root@centos7 centos]# hostname

lagoudocker
</code></pre>

<p>通过上面命令的输出，我们可以看到当前UTS Namespace 内的主机名已经被修改为 lagoudocker。然后我们新打开一个命令行窗口，使用相同的命令查看一下主机的 hostname：</p>

<pre><code>[centos@centos7 ~]$ hostname

centos7
</code></pre>

<p>可以看到主机的名称仍然为 centos7，并没有被修改。由此，可以验证 UTS Namespace 可以用来隔离主机名。</p>

<h4 id="4-ipc-namespace">（4）IPC Namespace</h4>

<p>IPC Namespace 主要是用来隔离进程间通信的。例如 PID Namespace 和 IPC Namespace 一起使用可以实现同一 IPC Namespace 内的进程彼此可以通信，不同 IPC Namespace 的进程却不能通信。</p>

<p>同样我们通过一个实例来验证下IPC Namespace的作用，首先我们使用 unshare 命令来创建一个 IPC Namespace：</p>

<pre><code>$ sudo unshare --ipc --fork /bin/bash

[root@centos7 centos]#
</code></pre>

<p>下面我们需要借助两个命令来实现对 IPC Namespace 的验证。</p>

<ul>
<li>ipcs -q 命令：用来查看系统间通信队列列表。</li>
<li>ipcmk -Q 命令：用来创建系统间通信队列。</li>
</ul>

<p>我们首先使用 ipcs -q 命令查看一下当前 IPC Namespace 下的系统通信队列列表：</p>

<pre><code>[centos@centos7 ~]$ ipcs -q

------ Message Queues --------

key        msqid      owner      perms      used-bytes   messages
</code></pre>

<p>由上可以看到当前无任何系统通信队列，然后我们使用 ipcmk -Q 命令创建一个系统通信队列：</p>

<pre><code>[root@centos7 centos]# ipcmk -Q

Message queue id: 0
</code></pre>

<p>再次使用 ipcs -q 命令查看当前 IPC Namespace 下的系统通信队列列表：</p>

<pre><code>[root@centos7 centos]# ipcs -q

------ Message Queues --------

key        msqid      owner      perms      used-bytes   messages

0x73682a32 0          root       644        0            0
</code></pre>

<p>可以看到我们已经成功创建了一个系统通信队列。然后我们新打开一个命令行窗口，使用ipcs -q 命令查看一下主机的系统通信队列：</p>

<pre><code>[centos@centos7 ~]$ ipcs -q

------ Message Queues --------

key        msqid      owner      perms      used-bytes   messages
</code></pre>

<p>通过上面的实验，可以发现，在单独的 IPC Namespace 内创建的系统通信队列在主机上无法看到。即 IPC Namespace 实现了系统通信队列的隔离。</p>

<h4 id="5-user-namespace">（5）User Namespace</h4>

<p>User Namespace 主要是用来隔离用户和用户组的。一个比较典型的应用场景就是在主机上以非 root 用户运行的进程可以在一个单独的 User Namespace 中映射成 root 用户。使用 User Namespace 可以实现进程在容器内拥有 root 权限，而在主机上却只是普通用户。</p>

<p>User Namesapce 的创建是可以不使用 root 权限的。下面我们以普通用户的身份创建一个 User Namespace，命令如下：</p>

<pre><code>[centos@centos7 ~]$ unshare --user -r /bin/bash

[root@centos7 ~]#
</code></pre>

<blockquote>
<p>CentOS7 默认允许创建的 User Namespace 为 0，如果执行上述命令失败（ unshare 命令返回的错误为 unshare: unshare failed: Invalid argument ），需要使用以下命令修改系统允许创建的 User Namespace 数量，命令为：echo 65535 &gt; /proc/sys/user/max_user_namespaces，然后再次尝试创建 User Namespace。</p>
</blockquote>

<p>然后执行 id 命令查看一下当前的用户信息：</p>

<pre><code>[root@centos7 ~]# id

uid=0(root) gid=0(root) groups=0(root),65534(nfsnobody) context=unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023
</code></pre>

<p>通过上面的输出可以看到我们在新的 User Namespace 内已经是 root 用户了。下面我们使用只有主机 root 用户才可以执行的 reboot 命令来验证一下，在当前命令行窗口执行 reboot 命令：</p>

<pre><code>[root@centos7 ~]# reboot

Failed to open /dev/initctl: Permission denied

Failed to talk to init daemon.
</code></pre>

<p>可以看到，我们在新创建的 User Namespace 内虽然是 root 用户，但是并没有权限执行 reboot 命令。这说明在隔离的 User Namespace 中，并不能获取到主机的 root 权限，也就是说 User Namespace 实现了用户和用户组的隔离。</p>

<h4 id="6-net-namespace">（6）Net Namespace</h4>

<p>Net Namespace 是用来隔离网络设备、IP 地址和端口等信息的。Net Namespace 可以让每个进程拥有自己独立的 IP 地址，端口和网卡信息。例如主机 IP 地址为 172.16.4.1 ，容器内可以设置独立的 IP 地址为 192.168.1.1。</p>

<p>同样用实例验证，我们首先使用 ip a 命令查看一下主机上的网络信息：</p>

<pre><code>$ ip a

1: lo: &lt;LOOPBACK,UP,LOWER_UP&gt; mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000

    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

    inet 127.0.0.1/8 scope host lo

       valid_lft forever preferred_lft forever

    inet6 ::1/128 scope host

       valid_lft forever preferred_lft forever

2: eth0: &lt;BROADCAST,MULTICAST,UP,LOWER_UP&gt; mtu 1500 qdisc pfifo_fast state UP group default qlen 1000

    link/ether 02:11:b0:14:01:0c brd ff:ff:ff:ff:ff:ff

    inet 172.20.1.11/24 brd 172.20.1.255 scope global dynamic eth0

       valid_lft 86063337sec preferred_lft 86063337sec

    inet6 fe80::11:b0ff:fe14:10c/64 scope link

       valid_lft forever preferred_lft forever

3: docker0: &lt;NO-CARRIER,BROADCAST,MULTICAST,UP&gt; mtu 1500 qdisc noqueue state DOWN group default

    link/ether 02:42:82:8d:a0:df brd ff:ff:ff:ff:ff:ff

    inet 172.17.0.1/16 scope global docker0

       valid_lft forever preferred_lft forever

    inet6 fe80::42:82ff:fe8d:a0df/64 scope link

       valid_lft forever preferred_lft forever
</code></pre>

<p>然后我们使用以下命令创建一个 Net Namespace：</p>

<pre><code>$ sudo unshare --net --fork /bin/bash

[root@centos7 centos]#
</code></pre>

<p>同样的我们使用 ip a 命令查看一下网络信息：</p>

<pre><code>[root@centos7 centos]# ip a

1: lo: &lt;LOOPBACK&gt; mtu 65536 qdisc noop state DOWN group default qlen 1000

    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
</code></pre>

<p>可以看到，宿主机上有 lo、eth0、docker0 等网络设备，而我们新建的 Net Namespace 内则与主机上的网络设备不同。</p>

<h3 id="为什么-docker-需要-namespace">为什么 Docker 需要 Namespace？</h3>

<p>Linux 内核从 2002 年 2.4.19 版本开始加入了 Mount Namespace，而直到内核 3.8 版本加入了 User Namespace 才为容器提供了足够的支持功能。</p>

<p>当 Docker 新建一个容器时， 它会创建这六种 Namespace，然后将容器中的进程加入这些 Namespace 之中，使得 Docker 容器中的进程只能看到当前 Namespace 中的系统资源。</p>

<p>正是由于 Docker 使用了 Linux 的这些 Namespace 技术，才实现了 Docker 容器的隔离，可以说没有 Namespace，就没有 Docker 容器。</p>

<h3 id="小结">小结</h3>

<p>到此，相信你已经了解了什么是 Namespace。Namespace 是 Linux 内核的一个特性，该特性可以实现在同一主机系统中对进程 ID、主机名、用户 ID、文件名、网络和进程间通信等资源的隔离。Docker 正是结合了这六种 Namespace 的功能，才诞生了 Docker 容器。</p>

<p>最后，试想下，当我们使用 docker run &ndash;net=host 命令启动容器时，容器是否和主机共享同一个 Net Namespace？思考后，可以把你的想法写在留言区。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#ee828282d7dadfdfded9ae89838f8782c08d8183" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f168d1afdd8f667',t:'MTczNDA5OTE2OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>