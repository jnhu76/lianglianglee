<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=01&#32;&#32;Docker&#32;安装：入门案例带你了解容器技术原理>
        <link rel="icon" href="/static/favicon.png">
        <title>01  Docker 安装：入门案例带你了解容器技术原理 </title>
        
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
                            <h1 id="title" data-id="01  Docker 安装：入门案例带你了解容器技术原理" class="title">01  Docker 安装：入门案例带你了解容器技术原理</h1>
                            <div><p>咱们第一课时就先聊聊 Docker 的基础内容：Docker 能做什么，怎么安装 Docker，以及容器技术的原理。</p>

<h3 id="docker-能做什么">Docker 能做什么？</h3>

<p>众所周知，Docker 是一个用于开发，发布和运行应用程序的开放平台。通俗地讲，Docker 类似于集装箱。在一艘大船上，各种货物要想被整齐摆放并且相互不受到影响，我们就需要把各种货物进行集装箱标准化。有了集装箱，我们就不需要专门运输水果或者化学用品的船了。我们可以把各种货品通过集装箱打包，然后统一放到一艘船上运输。Docker 要做的就是把各种软件打包成一个集装箱（镜像），然后分发，且在运行的时候可以相互隔离。</p>

<p>到此，相信你已经迫不及待想要体验下了，下面就让我们来安装一个 Docker。</p>

<h3 id="centos-下安装-docker">CentOS 下安装 Docker</h3>

<p>Docker 是跨平台的解决方案，它支持在当前主流的各大平台安装，包括 Ubuntu、RHEL、CentOS、Debian 等 Linux 发行版，同时也可以在 OSX 、Microsoft Windows 等非 Linux 平台下安装使用。</p>

<p>因为 Linux 是 Docker 的原生支持平台，所以推荐你在 Linux 上使用 Docker。由于生产环境中我们使用 CentOS 较多，下面主要针对在 CentOS 平台下安装和使用 Docker 展开介绍。</p>

<h4 id="操作系统要求">操作系统要求</h4>

<p>要安装 Docker，我们需要 CentOS 7 及以上的发行版本。建议使用<code>overlay2</code>存储驱动程序。</p>

<h4 id="卸载已有-docker">卸载已有 Docker</h4>

<p>如果你已经安装过旧版的 Docker，可以先执行以下命令卸载旧版 Docker。</p>

<pre><code>$ sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
</code></pre>

<h4 id="安装-docker">安装 Docker</h4>

<p>首次安装 Docker 之前，需要添加 Docker 安装源。添加之后，我们就可以从已经配置好的源，安装和更新 Docker。添加 Docker 安装源的命令如下：</p>

<pre><code>$ sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
</code></pre>

<p>正常情况下，直接安装最新版本的 Docker 即可，因为最新版本的 Docker 有着更好的稳定性和安全性。你可以使用以下命令安装最新版本的 Docker。</p>

<pre><code>$ sudo yum install docker-ce docker-ce-cli containerd.io
</code></pre>

<p>如果你想要安装指定版本的 Docker，可以使用以下命令：</p>

<pre><code>$ sudo yum list docker-ce --showduplicates | sort -r
docker-ce.x86_64            18.06.1.ce-3.el7                   docker-ce-stable
docker-ce.x86_64            18.06.0.ce-3.el7                   docker-ce-stable
docker-ce.x86_64            18.03.1.ce-1.el7.centos            docker-ce-stable
docker-ce.x86_64            18.03.0.ce-1.el7.centos            docker-ce-stable
docker-ce.x86_64            17.12.1.ce-1.el7.centos            docker-ce-stable
docker-ce.x86_64            17.12.0.ce-1.el7.centos            docker-ce-stable
docker-ce.x86_64            17.09.1.ce-1.el7.centos            docker-ce-stable
</code></pre>

<p>然后选取想要的版本执行以下命令：</p>

<pre><code>$ sudo yum install docker-ce-&lt;VERSION_STRING&gt; docker-ce-cli-&lt;VERSION_STRING&gt; containerd.io
</code></pre>

<p>安装完成后，使用以下命令启动 Docker。</p>

<pre><code>$ sudo systemctl start docker
</code></pre>

<p>这里有一个国际惯例，安装完成后，我们需要使用以下命令启动一个 hello world 的容器。</p>

<pre><code>$ sudo docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
0e03bdcc26d7: Pull complete
Digest: sha256:7f0a9f93b4aa3022c3a4c147a449bf11e0941a1fd0bf4a8e6c9408b2600777c5
Status: Downloaded newer image for hello-world:latest
Hello from Docker!
</code></pre>

<p>运行上述命令，Docker 首先会检查本地是否有<code>hello-world</code>这个镜像，如果发现本地没有这个镜像，Docker 就会去 Docker Hub 官方仓库下载此镜像，然后运行它。最后我们看到该镜像输出 &ldquo;Hello from Docker!&rdquo; 并退出。</p>

<blockquote>
<p>安装完成后默认 docker 命令只能以 root 用户执行，如果想允许普通用户执行 docker 命令，需要执行以下命令 sudo groupadd docker &amp;&amp; sudo gpasswd -a ${USER} docker &amp;&amp; sudo systemctl restart docker ，执行完命令后，退出当前命令行窗口并打开新的窗口即可。</p>
</blockquote>

<p>安装完 Docker，先不着急使用，先来了解下容器的技术原理，这样才能知其所以然。</p>

<h3 id="容器技术原理">容器技术原理</h3>

<p>提起容器就不得不说 chroot，因为 chroot 是最早的容器雏形。chroot 意味着切换根目录，有了 chroot 就意味着我们可以把任何目录更改为当前进程的根目录，这与容器非常相似，下面我们通过一个实例了解下 chroot。</p>

<h4 id="chroot">chroot</h4>

<p>什么是 chroot 呢？下面是 chroot 维基百科定义：</p>

<blockquote>
<p>chroot 是在 Unix 和 Linux 系统的一个操作，针对正在运作的软件行程和它的子进程，改变它外显的根目录。一个运行在这个环境下，经由 chroot 设置根目录的程序，它不能够对这个指定根目录之外的文件进行访问动作，不能读取，也不能更改它的内容。</p>
</blockquote>

<p>通俗地说 ，chroot 就是可以改变某进程的根目录，使这个程序不能访问目录之外的其他目录，这个跟我们在一个容器中是很相似的。下面我们通过一个实例来演示下 chroot。</p>

<p>首先我们在当前目录下创建一个 rootfs 目录：</p>

<pre><code>$ mkdir rootfs
</code></pre>

<p>这里为了方便演示，我使用现成的 busybox 镜像来创建一个系统，镜像的概念和组成后面我会详细讲解，如果你没有 Docker 基础可以把下面的操作命令理解成在 rootfs 下创建了一些目录和放置了一些二进制文件。</p>

<pre><code>$ cd rootfs 
$ docker export $(docker create busybox) -o busybox.tar
$ tar -xf busybox.tar
</code></pre>

<p>执行完上面的命令后，在 rootfs 目录下，我们会得到一些目录和文件。下面我们使用 ls 命令查看一下 rootfs 目录下的内容。</p>

<pre><code>$ ls
bin  busybox.tar  dev  etc  home  proc  root  sys  tmp  usr  var
</code></pre>

<p>可以看到我们在 rootfs 目录下初始化了一些目录，下面让我们通过一条命令来见证 chroot 的神奇之处。使用以下命令，可以启动一个 sh 进程，并且把 /home/centos/rootfs 作为 sh 进程的根目录。</p>

<pre><code>$ chroot /home/centos/rootfs /bin/sh
</code></pre>

<p>此时，我们的命令行窗口已经处于上述命令启动的 sh 进程中。在当前 sh 命令行窗口下，我们使用 ls 命令查看一下当前进程，看是否真的与主机上的其他目录隔离开了。</p>

<pre><code>/ # /bin/ls /
bin  busybox.tar  dev  etc  home  proc  root  sys  tmp  usr  var
</code></pre>

<p>这里可以看到当前进程的根目录已经变成了主机上的 /home/centos/rootfs 目录。这样就实现了当前进程与主机的隔离。到此为止，一个目录隔离的容器就完成了。
但是，此时还不能称之为一个容器，为什么呢？你可以在上一步（使用 chroot 启动命令行窗口）执行以下命令，查看如下路由信息：</p>

<pre><code>/etc # /bin/ip route
default via 172.20.1.1 dev eth0
172.17.0.0/16 dev docker0 scope link  src 172.17.0.1
172.20.1.0/24 dev eth0 scope link  src 172.20.1.3
</code></pre>

<p>执行 ip route 命令后，你可以看到网络信息并没有隔离，实际上进程等信息此时也并未隔离。要想实现一个完整的容器，我们还需要 Linux 的其他三项技术： Namespace、Cgroups 和联合文件系统。</p>

<p>Docker 是利用 Linux 的 Namespace 、Cgroups 和联合文件系统三大机制来保证实现的， 所以它的原理是使用 Namespace 做主机名、网络、PID 等资源的隔离，使用 Cgroups 对进程或者进程组做资源（例如：CPU、内存等）的限制，联合文件系统用于镜像构建和容器运行环境。</p>

<p>后面我会对这些技术进行详细讲解，这里我就简单解释下它们的作用。</p>

<h4 id="namespace">Namespace</h4>

<p>Namespace 是 Linux 内核的一项功能，该功能对内核资源进行隔离，使得容器中的进程都可以在单独的命名空间中运行，并且只可以访问当前容器命名空间的资源。Namespace 可以隔离进程 ID、主机名、用户 ID、文件名、网络访问和进程间通信等相关资源。</p>

<p>Docker 主要用到以下五种命名空间。</p>

<ul>
<li>pid namespace：用于隔离进程 ID。</li>
<li>net namespace：隔离网络接口，在虚拟的 net namespace 内用户可以拥有自己独立的 IP、路由、端口等。</li>
<li>mnt namespace：文件系统挂载点隔离。</li>
<li>ipc namespace：信号量,消息队列和共享内存的隔离。</li>
<li>uts namespace：主机名和域名的隔离。</li>
</ul>

<h4 id="cgroups">Cgroups</h4>

<p>Cgroups 是一种 Linux 内核功能，可以限制和隔离进程的资源使用情况（CPU、内存、磁盘 I/O、网络等）。在容器的实现中，Cgroups 通常用来限制容器的 CPU 和内存等资源的使用。</p>

<h4 id="联合文件系统">联合文件系统</h4>

<p>联合文件系统，又叫 UnionFS，是一种通过创建文件层进程操作的文件系统，因此，联合文件系统非常轻快。Docker 使用联合文件系统为容器提供构建层，使得容器可以实现写时复制以及镜像的分层构建和存储。常用的联合文件系统有 AUFS、Overlay 和 Devicemapper 等。</p>

<h3 id="结语">结语</h3>

<p>容器技术从 1979 年 chroot 的首次问世便已崭露头角，但是到了 2013 年，Dokcer 的横空出世才使得容器技术迅速发展，可见 Docker 对于容器技术的推动力和影响力。</p>

<blockquote>
<p>另外， Docker 还提供了工具和平台来管理容器的生命周期：</p>

<ol>
<li>使用容器开发应用程序及其支持组件。</li>
<li>容器成为分发和测试你的应用程序的单元。</li>
<li>可以将应用程序作为容器或协调服务部署到生产环境中。无论您的生产环境是本地数据中心，云提供商还是两者的混合，其工作原理都相同。</li>
</ol>
</blockquote>

<p>到此，相信你已经了解了实现容器的基本技术原理，并且对 Docker 的作用有了一定认知。那么你知道为什么容器技术在 Docker 出现之前一直没有爆发的根本原因吗？思考后，可以把你的想法写在留言区。</p>

<p>下一课时，我将讲解 Docker 的架构设计以及 Docker 的三大核心概念。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#91fdfdfda8a5a0a0a1a6d1f6fcf0f8fdbff2fefc" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f168b576b8ef667',t:'MTczNDA5OTA5Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>