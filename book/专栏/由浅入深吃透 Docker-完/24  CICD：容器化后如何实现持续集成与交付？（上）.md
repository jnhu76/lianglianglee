<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=24&#32;&#32;CICD：容器化后如何实现持续集成与交付？（上）>
        <link rel="icon" href="/static/favicon.png">
        <title>24  CICD：容器化后如何实现持续集成与交付？（上） </title>
        
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
                            <h1 id="title" data-id="24  CICD：容器化后如何实现持续集成与交付？（上）" class="title">24  CICD：容器化后如何实现持续集成与交付？（上）</h1>
                            <div><p>上一讲，我介绍了 DevOps 的概念与 DevOps 的一些思想。DevOps 的思想可以帮助我们缩短上线周期并且提高软件迭代速度，而 CI/CD 则是 DevOps 思想中最重要的部分。</p>

<p>具体来说，CI/CD 是一种通过在应用开发阶段，引入自动化的手段来频繁地构建应用，并且向客户交付应用的方法。它的核心理念是持续开发、持续部署以及持续交付，它还可以让自动化持续交付贯穿于整个应用生命周期，使得开发和运维统一参与协同支持。</p>

<p>下面我们来详细了解下 CI/CD 。</p>

<h3 id="什么是-ci-cd">什么是 CI/CD</h3>

<h4 id="ci-持续集成-continuous-integration">CI 持续集成（Continuous Integration）</h4>

<p>随着软件功能越来越复杂，一个大型项目要想在规定时间内顺利完成，就需要多位开发人员协同开发。但是，如果我们每个人都负责开发自己的代码，然后集中在某一天将代码合并在一起（称为“合并日”）。你会发现，代码可能会有很多冲突和编译问题，而这个处理过程十分烦琐、耗时，并且需要每一位工程师确认代码是否被覆盖，代码是否完整。这种情况显然不是我们想要看到的，这时持续集成（CI）就可以很好地帮助我们解决这个问题。</p>

<p>CI 持续集成要求开发人员频繁地（甚至是每天）将代码提交到共享分支中。一旦开发人员的代码被合并，将会自动触发构建流程来构建应用，并通过触发自动化测试（单元测试或者集成测试）来验证这些代码的提交，确保这些更改没有对应用造成影响。如果发现提交的代码在测试过程中或者构建过程中有问题，则会马上通知研发人员确认，修改代码并重新提交。通过将以往的定期合并代码的方式，改变为频繁提交代码并且自动构建和测试的方式，可以帮助我们<strong>及早地发现问题和解决冲突，减少代码出错。</strong></p>

<p>传统 CI 流程的实现十分复杂，无法做到标准化交付，而当我们的应用容器化后，应用构建的结果就是 Docker 镜像。代码检查完毕没有缺陷后合并入主分支。此时启动构建流程，构建系统会自动将我们的应用打包成 Docker 镜像，并且推送到镜像仓库。</p>

<h4 id="cd-持续交付-continuous-delivery">CD 持续交付（Continuous Delivery）</h4>

<p>当我们每次完成代码的测试和构建后，我们需要将编译后的镜像快速发布到测试环境，这时我们的持续交付就登场了。持续交付要求我们实现自动化准备测试环境、自动化测试应用、自动化监控代码质量，并且自动化交付生产环境镜像。</p>

<p>在以前，测试环境的构建是非常耗时的，并且很难保证测试环境和研发环境的一致性。但现在，借助于容器技术，我们可以很方便地构建出一个测试环境，并且可以保证开发和测试环境的一致性，这样不仅可以提高测试效率，还可以提高敏捷性。</p>

<p>容器化后的应用交付过程是这样的，我们将测试的环境交由 QA 来维护，当我们确定好本次上线要发布的功能列表时，我们将不同开发人员开发的 feature 分支的代码合并到 release 分支。然后由 QA 来将构建镜像部署到测试环境，结合自动测试和人工测试、自动检测和人工记录，形成完整的测试报告，并且把测试过程中遇到的问题交由开发人员修改，开发修改无误后再次构建测试环境进行测试。测试没有问题后，自动交付生产环境的镜像到镜像仓库。</p>

<h4 id="cd-持续部署-continuous-deployment">CD 持续部署（Continuous Deployment）</h4>

<p>CD 不仅有持续交付的含义，还代表持续部署。经测试无误打包完生产环境的镜像后，我们需要把镜像部署到生产环境，持续部署是最后阶段，它作为持续交付的延伸，可以自动将生产环境的镜像发布到生产环境中。</p>

<p>部署业务首先需要我们有一个资源池，实现资源自动化供给，而且有的应用还希望有自动伸缩的能力，根据外部流量自动调整容器的副本数，而这一切在容器云中都将变得十分简单。</p>

<p>我们可以想象，如果有客户提出了反馈，我们通过持续部署在几分钟内，就能在更改完代码的情况下，将新的应用版本发布到生产环境中（假设通过了自动化测试），这时我们就可以实现快速迭代，持续接收和整合用户的反馈，将用户体验做到极致。</p>

<p>讲了这么多概念，也许你会感觉比较枯燥乏味。下面我们就动动手，利用一些工具搭建一个 DevOps 环境。</p>

<p>搭建 DevOps 环境的工具非常多，这里我选择的工具为 Jenkins、Docker 和 GitLab。Jenkins 和 Docker 都已经介绍过了，这里我再介绍一下 Gitlab。</p>

<p>Gitlab 是由 Gitlab Inc. 开发的一款基于 Git 的代码托管平台，它的功能和 GitHub 类似，可以帮助我们存储代码。除此之外，GitLab 还具有在线编辑 wiki、issue 跟踪等功能，另外最新版本的 GitLab 还集成了 CI/CD 功能，不过这里我们仅仅使用 GitLab 的代码存储功能， CI/CD 还是交给我们的老牌持续集成工具 Jenkins 来做。</p>

<h3 id="docker-jenkins-gitlab-搭建-ci-cd-系统">Docker+Jenkins+GitLab 搭建 CI/CD 系统</h3>

<p>软件安装环境如下。</p>

<ul>
<li>操作系统：CentOS 7</li>
<li>Jenkins：tls 长期维护版</li>
<li>Docker：18.06</li>
<li>GitLab：13.3.8-ce.0</li>
</ul>

<h4 id="第一步-安装-docker">第一步：安装 Docker</h4>

<p>安装 Docker 的步骤可以在[第一讲]的内容中找到，这里就不再赘述。Docker 环境准备好后，我们就可以利用 Docker 来部署 GitLab 和 Jenkins 了。</p>

<h4 id="第二步-安装-gitlab">第二步：安装 GitLab</h4>

<p>GitLab 官方提供了 GitLab 的 Docker 镜像，因此我们只需要执行以下命令就可以快速启动一个 GitLab 服务了。</p>

<pre><code>$ docker run -d \

--hostname localhost \

-p 8080:80 -p 2222:22 \

--name gitlab \

--restart always \

--volume /tmp/gitlab/config:/etc/gitlab \

--volume /tmp/gitlab/logs:/var/log/gitlab \

--volume /tmp/gitlab/data:/var/opt/gitlab \

gitlab/gitlab-ce:13.3.8-ce.0
</code></pre>

<p>这个启动过程可能需要几分钟的时间。当服务启动后我们就可以通过 <a href="http://localhost:8080/" target="_blank">http://localhost:8080</a> 访问到我们的 GitLab 服务了。</p>

<p><img src="assets/CgqCHl-05ceAOEtOAAC7xTjpRgo536.png" alt="Drawing 0.png" /></p>

<p>图1 GitLab 设置密码界面</p>

<p>第一次登陆，GitLab 会要求我们设置管理员密码，我们输入管理员密码后点击确认即可，之后 GitLab 会自动跳转到登录页面。</p>

<p><img src="assets/CgqCHl-05eKAftEiAACO_hts6R8497.png" alt="Drawing 1.png" /></p>

<p>图2 GitLab 登录界面</p>

<p>然后输入默认管理员用户名：<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="a7c6c3cacec9e7c2dfc6cad7cbc289c4c8ca">[email&#160;protected]</a>，密码为我们上一步设置的密码。点击登录即可登录到系统中，至此，GitLab 已经安装成功。</p>

<h4 id="第三步-安装-jenkins">第三步：安装 Jenkins</h4>

<p>Jenkins 官方提供了 Jenkins 的 Docker 镜像，我们使用 Jenkins 镜像就可以一键启动一个 Jenkins 服务。命令如下：</p>

<pre><code># docker run -d --name=jenkins \

-p 8888:8080 \

-u root \

--restart always \

-v /var/run/docker.sock:/var/run/docker.sock \

-v /usr/bin/docker:/usr/bin/docker \

-v /tmp/jenkins_home:/var/jenkins_home \

jenkins/jenkins:lts
</code></pre>

<blockquote>
<p>这里，我将 docker.sock 和 docker 二进制挂载到了 Jenkins 容器中，是为了让 Jenkins 可以直接调用 docker 命令来构建应用镜像。</p>
</blockquote>

<p>Jenkins 的默认密码会在容器启动后打印在容器的日志中，我们可以通过以下命令找到 Jenkins 的默认密码，星号之间的类似于一串 UUID 的随机串就是我们的密码。</p>

<pre><code>$ docker logs -f jenkins

unning from: /usr/share/jenkins/jenkins.war

webroot: EnvVars.masterEnvVars.get(&quot;JENKINS_HOME&quot;)

2020-10-31 16:13:06.472+0000 [id=1]	INFO	org.eclipse.jetty.util.log.Log#initialized: Logging initialized @292ms to org.eclipse.jetty.util.log.JavaUtilLog

2020-10-31 16:13:06.581+0000 [id=1]	INFO	winstone.Logger#logInternal: Beginning extraction from war file

2020-10-31 16:13:08.369+0000 [id=1]	WARNING	o.e.j.s.handler.ContextHandler#setContextPath: Empty contextPath

... 省略部分启动日志

Jenkins initial setup is required. An admin user has been created and a password generated.

Please use the following password to proceed to installation:

*************************************************************

*************************************************************

*************************************************************

Jenkins initial setup is required. An admin user has been created and a password generated.

Please use the following password to proceed to installation:

fb3499944e4845bba9d4b7d9eb4e3932

This may also be found at: /var/jenkins_home/secrets/initialAdminPassword

*************************************************************

*************************************************************

*************************************************************

This may also be found at: /var/jenkins_home/secrets/initialAdminPassword

2020-10-31 16:17:07.577+0000 [id=28]	INFO	jenkins.InitReactorRunner$1#onAttained: Completed initialization

2020-10-31 16:17:07.589+0000 [id=21]	INFO	hudson.WebAppMain$3#run: Jenkins is fully up and running
</code></pre>

<p>之后，我们通过访问 <a href="http://localhost:8888/" target="_blank">http://localhost:8888</a> 就可以访问到 Jenkins 服务了。</p>

<p><img src="assets/Ciqc1F-05giAJYDhAADCpDZzl2M065.png" alt="Drawing 2.png" /></p>

<p>图3 Jenkins 登录界面</p>

<p>然后将日志中的密码粘贴到密码框即可，之后 Jenkins 会自动初始化，我们根据引导，安装推荐的插件即可。</p>

<p><img src="assets/Ciqc1F-05hSAHd7VAAEpFeu4qfY218.png" alt="Drawing 3.png" /></p>

<p>图4 Jenkins 引导页面</p>

<p>选择好安装推荐的插件后，Jenkins 会自动开始初始化一些常用插件。界面如下：</p>

<p><img src="assets/Ciqc1F-05h-AUSiFAAEAxHFCt30058.png" alt="Drawing 4.png" /></p>

<p>图5 Jenkins 插件初始化</p>

<p>插件初始化完后，创建管理员账户和密码，输入用户名、密码和邮箱等信息，然后点击保存并完成即可。</p>

<p><img src="assets/Ciqc1F-05iaANJXNAABXftlfsvQ115.png" alt="Drawing 5.png" /></p>

<p>图6 Jenkins 创建管理员</p>

<p>这里，确认 Jenkins 地址，我们就可以进入到 Jenkins 主页了。</p>

<p><img src="assets/Ciqc1F-05i2AR2lRAADHULl7Ysk145.png" alt="Drawing 6.png" /></p>

<p>图7 Jenkins 主页</p>

<p>然后在系统管理 -&gt; 插件管理 -&gt; 可选插件处，搜索 GitLab 和 Docker ，分别安装相关插件即可，以便我们的 Jenkins 服务和 GitLab 以及 Docker 可以更好地交互。</p>

<p><img src="assets/Ciqc1F-05j2AfmTUAAGVJGSmG1g804.png" alt="Drawing 7.png" />
<img src="assets/Ciqc1F-05kKAMr27AAGh4SQTNsU299.png" alt="Drawing 8.png" /></p>

<p>图8 Jenkins 插件安装</p>

<p>等待插件安装完成， 重启 Jenkins ，我们的 Jenkins 环境就准备完成了。</p>

<p>现在，我们的 Docker+Jenkins+GitLab 环境已经准备完成，后面只需要把我们的代码推送到 GitLab 中，并做相关的配置即可实现推送代码自动构建镜像和发布。</p>

<h3 id="结语">结语</h3>

<p>Docker 的出现解决了 CI/CD 流程中的各种问题，Docker 交付的镜像不仅包含应用程序，也包含了应用程序的运行环境，这很好地解决了开发和线上环境不一致问题。同时 Docker 的出现也极大地提升了 CI/CD 的构建效率，我们仅仅需要编写一个 Dockerfile 并将 Dockerfile 提交到我们的代码仓库即可快速构建出我们的应用，最后，当我们构建好 Docker 镜像后 Docker 可以帮助我们快速发布及更新应用。</p>

<p>那么，你知道 Docker 还可以帮助我们解决 CI/CD 流程中的哪些问题吗？</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#8de1e1e1b4b9bcbcbdbacdeae0ece4e1a3eee2e0" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f16b6648fc5ed08',t:'MTczNDEwMDg1OS4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>