<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=47&#32;如何编写Kubernetes资源定义文件？>
        <link rel="icon" href="/static/favicon.png">
        <title>47 如何编写Kubernetes资源定义文件？ </title>
        
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
                        <a class="menu-item" id="00 开篇词 从 0 开始搭建一个企业级 Go 应用.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%200%20%e5%bc%80%e5%a7%8b%e6%90%ad%e5%bb%ba%e4%b8%80%e4%b8%aa%e4%bc%81%e4%b8%9a%e7%ba%a7%20Go%20%e5%ba%94%e7%94%a8.md">00 开篇词 从 0 开始搭建一个企业级 Go 应用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 IAM系统概述：我们要实现什么样的 Go 项目？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/01%20IAM%e7%b3%bb%e7%bb%9f%e6%a6%82%e8%bf%b0%ef%bc%9a%e6%88%91%e4%bb%ac%e8%a6%81%e5%ae%9e%e7%8e%b0%e4%bb%80%e4%b9%88%e6%a0%b7%e7%9a%84%20Go%20%e9%a1%b9%e7%9b%ae%ef%bc%9f.md">01 IAM系统概述：我们要实现什么样的 Go 项目？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 环境准备：如何安装和配置一个基本的 Go 开发环境？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/02%20%e7%8e%af%e5%a2%83%e5%87%86%e5%a4%87%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%89%e8%a3%85%e5%92%8c%e9%85%8d%e7%bd%ae%e4%b8%80%e4%b8%aa%e5%9f%ba%e6%9c%ac%e7%9a%84%20Go%20%e5%bc%80%e5%8f%91%e7%8e%af%e5%a2%83%ef%bc%9f.md">02 环境准备：如何安装和配置一个基本的 Go 开发环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 项目部署：如何快速部署 IAM 系统？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/03%20%e9%a1%b9%e7%9b%ae%e9%83%a8%e7%bd%b2%ef%bc%9a%e5%a6%82%e4%bd%95%e5%bf%ab%e9%80%9f%e9%83%a8%e7%bd%b2%20IAM%20%e7%b3%bb%e7%bb%9f%ef%bc%9f.md">03 项目部署：如何快速部署 IAM 系统？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 规范设计（上）：项目开发杂乱无章，如何规范？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/04%20%e8%a7%84%e8%8c%83%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e6%9d%82%e4%b9%b1%e6%97%a0%e7%ab%a0%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%84%e8%8c%83%ef%bc%9f.md">04 规范设计（上）：项目开发杂乱无章，如何规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 规范设计（下）：commit 信息风格迥异、难以阅读，如何规范？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/05%20%e8%a7%84%e8%8c%83%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9acommit%20%e4%bf%a1%e6%81%af%e9%a3%8e%e6%a0%bc%e8%bf%a5%e5%bc%82%e3%80%81%e9%9a%be%e4%bb%a5%e9%98%85%e8%af%bb%ef%bc%8c%e5%a6%82%e4%bd%95%e8%a7%84%e8%8c%83%ef%bc%9f.md">05 规范设计（下）：commit 信息风格迥异、难以阅读，如何规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 目录结构设计：如何组织一个可维护、可扩展的代码目录？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/06%20%e7%9b%ae%e5%bd%95%e7%bb%93%e6%9e%84%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e7%bb%84%e7%bb%87%e4%b8%80%e4%b8%aa%e5%8f%af%e7%bb%b4%e6%8a%a4%e3%80%81%e5%8f%af%e6%89%a9%e5%b1%95%e7%9a%84%e4%bb%a3%e7%a0%81%e7%9b%ae%e5%bd%95%ef%bc%9f.md">06 目录结构设计：如何组织一个可维护、可扩展的代码目录？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 工作流设计：如何设计合理的多人开发模式？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/07%20%e5%b7%a5%e4%bd%9c%e6%b5%81%e8%ae%be%e8%ae%a1%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e5%90%88%e7%90%86%e7%9a%84%e5%a4%9a%e4%ba%ba%e5%bc%80%e5%8f%91%e6%a8%a1%e5%bc%8f%ef%bc%9f.md">07 工作流设计：如何设计合理的多人开发模式？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 研发流程设计（上）：如何设计 Go 项目的开发流程？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/08%20%e7%a0%94%e5%8f%91%e6%b5%81%e7%a8%8b%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%20Go%20%e9%a1%b9%e7%9b%ae%e7%9a%84%e5%bc%80%e5%8f%91%e6%b5%81%e7%a8%8b%ef%bc%9f.md">08 研发流程设计（上）：如何设计 Go 项目的开发流程？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 研发流程设计（下）：如何管理应用的生命周期？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/09%20%e7%a0%94%e5%8f%91%e6%b5%81%e7%a8%8b%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e7%ae%a1%e7%90%86%e5%ba%94%e7%94%a8%e7%9a%84%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%ef%bc%9f.md">09 研发流程设计（下）：如何管理应用的生命周期？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 设计方法：怎么写出优雅的 Go 项目？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/10%20%e8%ae%be%e8%ae%a1%e6%96%b9%e6%b3%95%ef%bc%9a%e6%80%8e%e4%b9%88%e5%86%99%e5%87%ba%e4%bc%98%e9%9b%85%e7%9a%84%20Go%20%e9%a1%b9%e7%9b%ae%ef%bc%9f.md">10 设计方法：怎么写出优雅的 Go 项目？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 设计模式：Go常用设计模式概述.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/11%20%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%ef%bc%9aGo%e5%b8%b8%e7%94%a8%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f%e6%a6%82%e8%bf%b0.md">11 设计模式：Go常用设计模式概述.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 API 风格（上）：如何设计RESTful API？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/12%20API%20%e9%a3%8e%e6%a0%bc%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1RESTful%20API%ef%bc%9f.md">12 API 风格（上）：如何设计RESTful API？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 API 风格（下）：RPC API介绍.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/13%20API%20%e9%a3%8e%e6%a0%bc%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aRPC%20API%e4%bb%8b%e7%bb%8d.md">13 API 风格（下）：RPC API介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 项目管理：如何编写高质量的Makefile？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/14%20%e9%a1%b9%e7%9b%ae%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e7%bc%96%e5%86%99%e9%ab%98%e8%b4%a8%e9%87%8f%e7%9a%84Makefile%ef%bc%9f.md">14 项目管理：如何编写高质量的Makefile？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 研发流程实战：IAM项目是如何进行研发流程管理的？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/15%20%e7%a0%94%e5%8f%91%e6%b5%81%e7%a8%8b%e5%ae%9e%e6%88%98%ef%bc%9aIAM%e9%a1%b9%e7%9b%ae%e6%98%af%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e7%a0%94%e5%8f%91%e6%b5%81%e7%a8%8b%e7%ae%a1%e7%90%86%e7%9a%84%ef%bc%9f.md">15 研发流程实战：IAM项目是如何进行研发流程管理的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 代码检查：如何进行静态代码检查？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/16%20%e4%bb%a3%e7%a0%81%e6%a3%80%e6%9f%a5%ef%bc%9a%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e9%9d%99%e6%80%81%e4%bb%a3%e7%a0%81%e6%a3%80%e6%9f%a5%ef%bc%9f.md">16 代码检查：如何进行静态代码检查？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 API 文档：如何生成 Swagger API 文档 ？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/17%20API%20%e6%96%87%e6%a1%a3%ef%bc%9a%e5%a6%82%e4%bd%95%e7%94%9f%e6%88%90%20Swagger%20API%20%e6%96%87%e6%a1%a3%20%ef%bc%9f.md">17 API 文档：如何生成 Swagger API 文档 ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 错误处理（上）：如何设计一套科学的错误码？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/18%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e4%b8%80%e5%a5%97%e7%a7%91%e5%ad%a6%e7%9a%84%e9%94%99%e8%af%af%e7%a0%81%ef%bc%9f.md">18 错误处理（上）：如何设计一套科学的错误码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 错误处理（下）：如何设计错误包？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/19%20%e9%94%99%e8%af%af%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e9%94%99%e8%af%af%e5%8c%85%ef%bc%9f.md">19 错误处理（下）：如何设计错误包？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 日志处理（上）：如何设计日志包并记录日志？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/20%20%e6%97%a5%e5%bf%97%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e6%97%a5%e5%bf%97%e5%8c%85%e5%b9%b6%e8%ae%b0%e5%bd%95%e6%97%a5%e5%bf%97%ef%bc%9f.md">20 日志处理（上）：如何设计日志包并记录日志？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 日志处理（下）：手把手教你从 0 编写一个日志包.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/21%20%e6%97%a5%e5%bf%97%e5%a4%84%e7%90%86%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%89%8b%e6%8a%8a%e6%89%8b%e6%95%99%e4%bd%a0%e4%bb%8e%200%20%e7%bc%96%e5%86%99%e4%b8%80%e4%b8%aa%e6%97%a5%e5%bf%97%e5%8c%85.md">21 日志处理（下）：手把手教你从 0 编写一个日志包.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 应用构建三剑客：Pflag、Viper、Cobra 核心功能介绍.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/22%20%e5%ba%94%e7%94%a8%e6%9e%84%e5%bb%ba%e4%b8%89%e5%89%91%e5%ae%a2%ef%bc%9aPflag%e3%80%81Viper%e3%80%81Cobra%20%e6%a0%b8%e5%bf%83%e5%8a%9f%e8%83%bd%e4%bb%8b%e7%bb%8d.md">22 应用构建三剑客：Pflag、Viper、Cobra 核心功能介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 应用构建实战：如何构建一个优秀的企业应用框架？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/23%20%e5%ba%94%e7%94%a8%e6%9e%84%e5%bb%ba%e5%ae%9e%e6%88%98%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%e4%bc%98%e7%a7%80%e7%9a%84%e4%bc%81%e4%b8%9a%e5%ba%94%e7%94%a8%e6%a1%86%e6%9e%b6%ef%bc%9f.md">23 应用构建实战：如何构建一个优秀的企业应用框架？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 Web 服务：Web 服务核心功能有哪些，如何实现？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/24%20Web%20%e6%9c%8d%e5%8a%a1%ef%bc%9aWeb%20%e6%9c%8d%e5%8a%a1%e6%a0%b8%e5%bf%83%e5%8a%9f%e8%83%bd%e6%9c%89%e5%93%aa%e4%ba%9b%ef%bc%8c%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%ef%bc%9f.md">24 Web 服务：Web 服务核心功能有哪些，如何实现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 认证机制：应用程序如何进行访问认证？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/25%20%e8%ae%a4%e8%af%81%e6%9c%ba%e5%88%b6%ef%bc%9a%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e8%ae%bf%e9%97%ae%e8%ae%a4%e8%af%81%ef%bc%9f.md">25 认证机制：应用程序如何进行访问认证？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 IAM项目是如何设计和实现访问认证功能的？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/26%20IAM%e9%a1%b9%e7%9b%ae%e6%98%af%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e5%92%8c%e5%ae%9e%e7%8e%b0%e8%ae%bf%e9%97%ae%e8%ae%a4%e8%af%81%e5%8a%9f%e8%83%bd%e7%9a%84%ef%bc%9f.md">26 IAM项目是如何设计和实现访问认证功能的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 权限模型：5大权限模型是如何进行资源授权的？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/27%20%e6%9d%83%e9%99%90%e6%a8%a1%e5%9e%8b%ef%bc%9a5%e5%a4%a7%e6%9d%83%e9%99%90%e6%a8%a1%e5%9e%8b%e6%98%af%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e8%b5%84%e6%ba%90%e6%8e%88%e6%9d%83%e7%9a%84%ef%bc%9f.md">27 权限模型：5大权限模型是如何进行资源授权的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 控制流（上）：通过iam-apiserver设计，看Web服务的构建.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/28%20%e6%8e%a7%e5%88%b6%e6%b5%81%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e9%80%9a%e8%bf%87iam-apiserver%e8%ae%be%e8%ae%a1%ef%bc%8c%e7%9c%8bWeb%e6%9c%8d%e5%8a%a1%e7%9a%84%e6%9e%84%e5%bb%ba.md">28 控制流（上）：通过iam-apiserver设计，看Web服务的构建.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 控制流（下）：iam-apiserver服务核心功能实现讲解.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/29%20%e6%8e%a7%e5%88%b6%e6%b5%81%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aiam-apiserver%e6%9c%8d%e5%8a%a1%e6%a0%b8%e5%bf%83%e5%8a%9f%e8%83%bd%e5%ae%9e%e7%8e%b0%e8%ae%b2%e8%a7%a3.md">29 控制流（下）：iam-apiserver服务核心功能实现讲解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 ORM：CURD 神器 GORM 包介绍及实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/30%20ORM%ef%bc%9aCURD%20%e7%a5%9e%e5%99%a8%20GORM%20%e5%8c%85%e4%bb%8b%e7%bb%8d%e5%8f%8a%e5%ae%9e%e6%88%98.md">30 ORM：CURD 神器 GORM 包介绍及实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 数据流：通过iam-authz-server设计，看数据流服务的设计.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/31%20%e6%95%b0%e6%8d%ae%e6%b5%81%ef%bc%9a%e9%80%9a%e8%bf%87iam-authz-server%e8%ae%be%e8%ae%a1%ef%bc%8c%e7%9c%8b%e6%95%b0%e6%8d%ae%e6%b5%81%e6%9c%8d%e5%8a%a1%e7%9a%84%e8%ae%be%e8%ae%a1.md">31 数据流：通过iam-authz-server设计，看数据流服务的设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 数据处理：如何高效处理应用程序产生的数据？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/32%20%e6%95%b0%e6%8d%ae%e5%a4%84%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e5%a4%84%e7%90%86%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%e4%ba%a7%e7%94%9f%e7%9a%84%e6%95%b0%e6%8d%ae%ef%bc%9f.md">32 数据处理：如何高效处理应用程序产生的数据？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  SDK 设计（上）：如何设计出一个优秀的 Go SDK？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/33%20%20SDK%20%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e5%87%ba%e4%b8%80%e4%b8%aa%e4%bc%98%e7%a7%80%e7%9a%84%20Go%20SDK%ef%bc%9f.md">33  SDK 设计（上）：如何设计出一个优秀的 Go SDK？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 SDK 设计（下）：IAM项目Go SDK设计和实现.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/34%20SDK%20%e8%ae%be%e8%ae%a1%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aIAM%e9%a1%b9%e7%9b%aeGo%20SDK%e8%ae%be%e8%ae%a1%e5%92%8c%e5%ae%9e%e7%8e%b0.md">34 SDK 设计（下）：IAM项目Go SDK设计和实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 效率神器：如何设计和实现一个命令行客户端工具？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/35%20%e6%95%88%e7%8e%87%e7%a5%9e%e5%99%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%be%e8%ae%a1%e5%92%8c%e5%ae%9e%e7%8e%b0%e4%b8%80%e4%b8%aa%e5%91%bd%e4%bb%a4%e8%a1%8c%e5%ae%a2%e6%88%b7%e7%ab%af%e5%b7%a5%e5%85%b7%ef%bc%9f.md">35 效率神器：如何设计和实现一个命令行客户端工具？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 代码测试（上）：如何编写 Go 语言单元测试和性能测试用例？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/36%20%e4%bb%a3%e7%a0%81%e6%b5%8b%e8%af%95%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e7%bc%96%e5%86%99%20Go%20%e8%af%ad%e8%a8%80%e5%8d%95%e5%85%83%e6%b5%8b%e8%af%95%e5%92%8c%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e7%94%a8%e4%be%8b%ef%bc%9f.md">36 代码测试（上）：如何编写 Go 语言单元测试和性能测试用例？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37 代码测试（下）：Go 语言其他测试类型及 IAM 测试介绍.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/37%20%e4%bb%a3%e7%a0%81%e6%b5%8b%e8%af%95%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%85%b6%e4%bb%96%e6%b5%8b%e8%af%95%e7%b1%bb%e5%9e%8b%e5%8f%8a%20IAM%20%e6%b5%8b%e8%af%95%e4%bb%8b%e7%bb%8d.md">37 代码测试（下）：Go 语言其他测试类型及 IAM 测试介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38 性能分析（上）：如何分析 Go 语言代码的性能？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/38%20%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%86%e6%9e%90%20Go%20%e8%af%ad%e8%a8%80%e4%bb%a3%e7%a0%81%e7%9a%84%e6%80%a7%e8%83%bd%ef%bc%9f.md">38 性能分析（上）：如何分析 Go 语言代码的性能？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39 性能分析（下）：API Server性能测试和调优实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/39%20%e6%80%a7%e8%83%bd%e5%88%86%e6%9e%90%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aAPI%20Server%e6%80%a7%e8%83%bd%e6%b5%8b%e8%af%95%e5%92%8c%e8%b0%83%e4%bc%98%e5%ae%9e%e6%88%98.md">39 性能分析（下）：API Server性能测试和调优实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40 软件部署实战（上）：部署方案及负载均衡、高可用组件介绍.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/40%20%e8%bd%af%e4%bb%b6%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e9%83%a8%e7%bd%b2%e6%96%b9%e6%a1%88%e5%8f%8a%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1%e3%80%81%e9%ab%98%e5%8f%af%e7%94%a8%e7%bb%84%e4%bb%b6%e4%bb%8b%e7%bb%8d.md">40 软件部署实战（上）：部署方案及负载均衡、高可用组件介绍.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="41 软件部署实战（中）：IAM 系统生产环境部署实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/41%20%e8%bd%af%e4%bb%b6%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9aIAM%20%e7%b3%bb%e7%bb%9f%e7%94%9f%e4%ba%a7%e7%8e%af%e5%a2%83%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98.md">41 软件部署实战（中）：IAM 系统生产环境部署实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="42 软件部署实战（下）：IAM系统安全加固、水平扩缩容实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/42%20%e8%bd%af%e4%bb%b6%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aIAM%e7%b3%bb%e7%bb%9f%e5%ae%89%e5%85%a8%e5%8a%a0%e5%9b%ba%e3%80%81%e6%b0%b4%e5%b9%b3%e6%89%a9%e7%bc%a9%e5%ae%b9%e5%ae%9e%e6%88%98.md">42 软件部署实战（下）：IAM系统安全加固、水平扩缩容实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="43 技术演进（上）：虚拟化技术演进之路.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/43%20%e6%8a%80%e6%9c%af%e6%bc%94%e8%bf%9b%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e8%99%9a%e6%8b%9f%e5%8c%96%e6%8a%80%e6%9c%af%e6%bc%94%e8%bf%9b%e4%b9%8b%e8%b7%af.md">43 技术演进（上）：虚拟化技术演进之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="44 技术演进（下）：软件架构和应用生命周期技术演进之路.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/44%20%e6%8a%80%e6%9c%af%e6%bc%94%e8%bf%9b%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e8%bd%af%e4%bb%b6%e6%9e%b6%e6%9e%84%e5%92%8c%e5%ba%94%e7%94%a8%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f%e6%8a%80%e6%9c%af%e6%bc%94%e8%bf%9b%e4%b9%8b%e8%b7%af.md">44 技术演进（下）：软件架构和应用生命周期技术演进之路.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="45 基于Kubernetes的云原生架构设计.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/45%20%e5%9f%ba%e4%ba%8eKubernetes%e7%9a%84%e4%ba%91%e5%8e%9f%e7%94%9f%e6%9e%b6%e6%9e%84%e8%ae%be%e8%ae%a1.md">45 基于Kubernetes的云原生架构设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="46 如何制作Docker镜像？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/46%20%e5%a6%82%e4%bd%95%e5%88%b6%e4%bd%9cDocker%e9%95%9c%e5%83%8f%ef%bc%9f.md">46 如何制作Docker镜像？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="47 如何编写Kubernetes资源定义文件？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/47%20%e5%a6%82%e4%bd%95%e7%bc%96%e5%86%99Kubernetes%e8%b5%84%e6%ba%90%e5%ae%9a%e4%b9%89%e6%96%87%e4%bb%b6%ef%bc%9f.md">47 如何编写Kubernetes资源定义文件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="48 IAM 容器化部署实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/48%20IAM%20%e5%ae%b9%e5%99%a8%e5%8c%96%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98.md">48 IAM 容器化部署实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="49 服务编排（上）：Helm服务编排基础知识.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/49%20%e6%9c%8d%e5%8a%a1%e7%bc%96%e6%8e%92%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aHelm%e6%9c%8d%e5%8a%a1%e7%bc%96%e6%8e%92%e5%9f%ba%e7%a1%80%e7%9f%a5%e8%af%86.md">49 服务编排（上）：Helm服务编排基础知识.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="50 服务编排（下）：基于Helm的服务编排部署实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/50%20%e6%9c%8d%e5%8a%a1%e7%bc%96%e6%8e%92%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e5%9f%ba%e4%ba%8eHelm%e7%9a%84%e6%9c%8d%e5%8a%a1%e7%bc%96%e6%8e%92%e9%83%a8%e7%bd%b2%e5%ae%9e%e6%88%98.md">50 服务编排（下）：基于Helm的服务编排部署实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="51 基于 GitHub Actions 的 CI 实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/51%20%e5%9f%ba%e4%ba%8e%20GitHub%20Actions%20%e7%9a%84%20CI%20%e5%ae%9e%e6%88%98.md">51 基于 GitHub Actions 的 CI 实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 Go Modules依赖包管理全讲.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20Go%20Modules%e4%be%9d%e8%b5%96%e5%8c%85%e7%ae%a1%e7%90%86%e5%85%a8%e8%ae%b2.md">特别放送 Go Modules依赖包管理全讲.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 Go Modules实战.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20Go%20Modules%e5%ae%9e%e6%88%98.md">特别放送 Go Modules实战.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 IAM排障指南.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20IAM%e6%8e%92%e9%9a%9c%e6%8c%87%e5%8d%97.md">特别放送 IAM排障指南.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 分布式作业系统设计和实现.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%bd%9c%e4%b8%9a%e7%b3%bb%e7%bb%9f%e8%ae%be%e8%ae%a1%e5%92%8c%e5%ae%9e%e7%8e%b0.md">特别放送 分布式作业系统设计和实现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 给你一份Go项目中最常用的Makefile核心语法.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%bb%99%e4%bd%a0%e4%b8%80%e4%bb%bdGo%e9%a1%b9%e7%9b%ae%e4%b8%ad%e6%9c%80%e5%b8%b8%e7%94%a8%e7%9a%84Makefile%e6%a0%b8%e5%bf%83%e8%af%ad%e6%b3%95.md">特别放送 给你一份Go项目中最常用的Makefile核心语法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="特别放送 给你一份清晰、可直接套用的Go编码规范.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%89%b9%e5%88%ab%e6%94%be%e9%80%81%20%e7%bb%99%e4%bd%a0%e4%b8%80%e4%bb%bd%e6%b8%85%e6%99%b0%e3%80%81%e5%8f%af%e7%9b%b4%e6%8e%a5%e5%a5%97%e7%94%a8%e7%9a%84Go%e7%bc%96%e7%a0%81%e8%a7%84%e8%8c%83.md">特别放送 给你一份清晰、可直接套用的Go编码规范.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="直播加餐 如何从小白进阶成 Go 语言专家？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%9b%b4%e6%92%ad%e5%8a%a0%e9%a4%90%20%e5%a6%82%e4%bd%95%e4%bb%8e%e5%b0%8f%e7%99%bd%e8%bf%9b%e9%98%b6%e6%88%90%20Go%20%e8%af%ad%e8%a8%80%e4%b8%93%e5%ae%b6%ef%bc%9f.md">直播加餐 如何从小白进阶成 Go 语言专家？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 如何让自己的 Go 研发之路走得更远？.md" href="/%e4%b8%93%e6%a0%8f/Go%20%e8%af%ad%e8%a8%80%e9%a1%b9%e7%9b%ae%e5%bc%80%e5%8f%91%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%a6%82%e4%bd%95%e8%ae%a9%e8%87%aa%e5%b7%b1%e7%9a%84%20Go%20%e7%a0%94%e5%8f%91%e4%b9%8b%e8%b7%af%e8%b5%b0%e5%be%97%e6%9b%b4%e8%bf%9c%ef%bc%9f.md">结束语 如何让自己的 Go 研发之路走得更远？.md</a>
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
                            <h1 id="title" data-id="47 如何编写Kubernetes资源定义文件？" class="title">47 如何编写Kubernetes资源定义文件？</h1>
                            <div><p>你好，我是孔令飞。</p>

<p>在接下来的48讲，我会介绍如何基于腾讯云EKS来部署IAM应用。EKS其实是一个标准的Kubernetes集群，在Kubernetes集群中部署应用，需要编写Kubernetes资源的YAML（Yet Another Markup Language）定义文件，例如Service、Deployment、ConfigMap、Secret、StatefulSet等。</p>

<p>这些YAML定义文件里面有很多配置项需要我们去配置，其中一些也比较难理解。为了你在学习下一讲时更轻松，这一讲我们先学习下如何编写Kubernetes YAML文件。</p>

<h2 id="为什么选择yaml格式来定义kubernetes资源">为什么选择YAML格式来定义Kubernetes资源？</h2>

<p>首先解释一下，我们为什么使用YAML格式来定义Kubernetes的各类资源呢？这是因为YAML格式和其他格式（例如XML、JSON等）相比，不仅能够支持丰富的数据，而且结构清晰、层次分明、表达性极强、易于维护，非常适合拿来供开发者配置和管理Kubernetes资源。</p>

<p>其实Kubernetes支持YAML和JSON两种格式，JSON格式通常用来作为接口之间消息传递的数据格式，YAML格式则用于资源的配置和管理。YAML和JSON这两种格式是可以相互转换的，你可以通过在线工具<a href="https://www.json2yaml.com/convert-yaml-to-json" target="_blank">json2yaml</a>，来自动转换YAML和JSON数据格式。</p>

<p>例如，下面是一个YAML文件中的内容：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: Service
metadata:
  name: iam-apiserver
spec:
  clusterIP: 192.168.0.231
  externalTrafficPolicy: Cluster
  ports:
  - name: https
    nodePort: 30443
    port: 8443
    protocol: TCP
    targetPort: 8443
  selector:
    app: iam-apiserver
  sessionAffinity: None
  type: NodePort
</code></pre>

<p>它对应的JSON格式的文件内容为：</p>

<pre><code class="language-json">{
  &quot;apiVersion&quot;: &quot;v1&quot;,
  &quot;kind&quot;: &quot;Service&quot;,
  &quot;metadata&quot;: {
    &quot;name&quot;: &quot;iam-apiserver&quot;
  },
  &quot;spec&quot;: {
    &quot;clusterIP&quot;: &quot;192.168.0.231&quot;,
    &quot;externalTrafficPolicy&quot;: &quot;Cluster&quot;,
    &quot;ports&quot;: [
      {
        &quot;name&quot;: &quot;https&quot;,
        &quot;nodePort&quot;: 30443,
        &quot;port&quot;: 8443,
        &quot;protocol&quot;: &quot;TCP&quot;,
        &quot;targetPort&quot;: 8443
      }
    ],
    &quot;selector&quot;: {
      &quot;app&quot;: &quot;iam-apiserver&quot;
    },
    &quot;sessionAffinity&quot;: &quot;None&quot;,
    &quot;type&quot;: &quot;NodePort&quot;
  }
}
</code></pre>

<p>我就是通过<code>json2yaml</code>在线工具，来转换YAML和JSON的，如下图所示：</p>

<p><img src="assets/e933e6c931544178b1295fc7becf6722.jpg" alt="图片" /></p>

<p>在编写Kubernetes资源定义文件的过程中，如果因为YAML格式文件中的配置项缩进太深，导致不容易判断配置项的层级，那么，你就可以将其转换成JSON格式，通过JSON格式来判断配置型的层级。</p>

<p>如果想学习更多关于YAML的知识，你可以参考<a href="https://yaml.org/spec/1.2/spec.html" target="_blank">YAML 1.2 (3rd Edition)</a>。这里，可以先看看我整理的YAML基本语法：</p>

<ul>
<li>属性和值都是大小写敏感的。</li>
<li>使用缩进表示层级关系。</li>
<li>禁止使用Tab键缩进，只允许使用空格，建议两个空格作为一个层级的缩进。元素左对齐，就说明对齐的两个元素属于同一个级别。</li>
<li>使用 <code>#</code> 进行注释，直到行尾。</li>
<li><code>key: value</code>格式的定义中，冒号后要有一个空格。</li>
<li>短横线表示列表项，使用一个短横线加一个空格；多个项使用同样的缩进级别作为同一列表。</li>
<li>使用 <code>---</code> 表示一个新的YAML文件开始。</li>
</ul>

<p>现在你知道了，Kubernetes支持YAML和JSON两种格式，它们是可以相互转换的。但鉴于YAML格式的各项优点，我建议你使用YAML格式来定义Kubernetes的各类资源。</p>

<h2 id="kubernetes-资源定义概述">Kubernetes 资源定义概述</h2>

<p>Kubernetes中有很多内置的资源，常用的资源有Deployment、StatefulSet、ConfigMap、Service、Secret、Nodes、Pods、Events、Jobs、DaemonSets等。除此之外，Kubernetes还有其他一些资源。如果你觉得Kubernetes内置的资源满足不了需求，还可以自定义资源。</p>

<p>Kubernetes的资源清单可以通过执行以下命令来查看：</p>

<pre><code class="language-bash">$ kubectl api-resources
NAME                              SHORTNAMES   APIVERSION                        NAMESPACED   KIND
bindings                                       v1                                true         Binding
componentstatuses                 cs           v1                                false        ComponentStatus
configmaps                        cm           v1                                true         ConfigMap
endpoints                         ep           v1                                true         Endpoints
events                            ev           v1                                true         Event
</code></pre>

<p>上述输出中，各列的含义如下。</p>

<ul>
<li>NAME：资源名称。</li>
<li>SHORTNAMES：资源名称简写。</li>
<li>APIVERSION：资源的API版本，也称为group。</li>
<li>NAMESPACED：资源是否具有Namespace属性。</li>
<li>KIND：资源类别。</li>
</ul>

<p>这些资源有一些共同的配置，也有一些特有的配置。这里，我们先来看下这些资源共同的配置。</p>

<p>下面这些配置是Kubernetes各类资源都具备的：</p>

<pre><code class="language-yaml">---
apiVersion: &lt;string&gt; # string类型，指定group的名称，默认为core。可以使用 `kubectl api-versions` 命令，来获取当前kubernetes版本支持的所有group。
kind: &lt;string&gt; # string类型，资源类别。
metadata: &lt;Object&gt; # 资源的元数据。
  name: &lt;string&gt; # string类型，资源名称。
  namespace: &lt;string&gt; # string类型，资源所属的命名空间。
  lables: &lt; map[string]string&gt; # map类型，资源的标签。
  annotations: &lt; map[string]string&gt; # map类型，资源的标注。
  selfLink: &lt;string&gt; # 资源的 REST API路径，格式为：/api/&lt;group&gt;/namespaces/&lt;namespace&gt;/&lt;type&gt;/&lt;name&gt;。例如：/api/v1/namespaces/default/services/iam-apiserver
spec: &lt;Object&gt; # 定义用户期望的资源状态（disired state）。
status: &lt;Object&gt; # 资源当前的状态，以只读的方式显示资源的最近状态。这个字段由kubernetes维护，用户无法定义。
</code></pre>

<p>你可以通过<code>kubectl explain &lt;object&gt;</code>命令来查看Object资源对象介绍，并通过<code>kubectl explain &lt;object1&gt;.&lt;object2&gt;</code>来查看<code>&lt;object1&gt;</code>的子对象<code>&lt;object2&gt;</code>的资源介绍，例如：</p>

<pre><code class="language-bash">$ kubectl explain service
$ kubectl explain service.spec
$ kubectl explain service.spec.ports
</code></pre>

<p>Kubernetes资源定义YAML文件，支持以下数据类型：</p>

<ul>
<li>string，表示字符串类型。</li>
<li>object，表示一个对象，需要嵌套多层字段。</li>
<li>map[string]string，表示由key:value组成的映射。</li>
<li>[]string，表示字串列表。</li>
<li>[]object，表示对象列表。</li>
<li>boolean，表示布尔类型。</li>
<li>integer，表示整型。</li>
</ul>

<h2 id="常用的kubernetes资源定义">常用的Kubernetes资源定义</h2>

<p>上面说了，Kubernetes中有很多资源，其中Pod、Deployment、Service、ConfigMap这4类是比较常用的资源，我来一个个介绍下。</p>

<h3 id="pod资源定义">Pod资源定义</h3>

<p>下面是一个Pod的YAML定义：</p>

<pre><code class="language-yaml">apiVersion: v1   # 必须 版本号， 常用v1  apps/v1
kind: Pod	 # 必须
metadata:  # 必须，元数据
  name: string  # 必须，名称
  namespace: string # 必须，命名空间，默认上default,生产环境为了安全性建议新建命名空间分类存放
  labels:   # 非必须，标签，列表值
    - name: string
  annotations:  # 非必须，注解，列表值
    - name: string
spec:  # 必须，容器的详细定义
  containers:  #必须，容器列表，
    - name: string　　　#必须，容器1的名称
      image: string		#必须，容器1所用的镜像
      imagePullPolicy: [Always|Never|IfNotPresent]  #非必须，镜像拉取策略，默认是Always
      command: [string]  # 非必须 列表值，如果不指定，则是一镜像打包时使用的启动命令
      args:　[string] # 非必须，启动参数
      workingDir: string # 非必须，容器内的工作目录
      volumeMounts: # 非必须，挂载到容器内的存储卷配置
        - name: string  # 非必须，存储卷名字，需与【@1】处定义的名字一致
          readOnly: boolean #非必须，定义读写模式，默认是读写
      ports: # 非必须，需要暴露的端口
        - name: string  # 非必须 端口名称
          containerPort: int  # 非必须 端口号
          hostPort: int # 非必须 宿主机需要监听的端口号，设置此值时，同一台宿主机不能存在同一端口号的pod， 建议不要设置此值
          proctocol: [tcp|udp]  # 非必须 端口使用的协议，默认是tcp
      env: # 非必须 环境变量
        - name: string # 非必须 ，环境变量名称
          value: string  # 非必须，环境变量键值对
      resources:  # 非必须，资源限制
        limits:  # 非必须，限制的容器使用资源的最大值，超过此值容器会推出
          cpu: string # 非必须，cpu资源，单位是core，从0.1开始
          memory: string 内存限制，单位为MiB,GiB
        requests:  # 非必须，启动时分配的资源
          cpu: string 
          memory: string
      livenessProbe:   # 非必须，容器健康检查的探针探测方式
        exec: # 探测命令
          command: [string] # 探测命令或者脚本
        httpGet: # httpGet方式
          path: string  # 探测路径，例如 http://ip:port/path
          port: number  
          host: string  
          scheme: string
          httpHeaders:
            - name: string
              value: string
          tcpSocket:  # tcpSocket方式，检查端口是否存在
            port: number
          initialDelaySeconds: 0 #容器启动完成多少秒后的再进行首次探测，单位为s
          timeoutSeconds: 0  #探测响应超时的时间,默认是1s,如果失败，则认为容器不健康，会重启该容器
          periodSeconds: 0  # 探测间隔时间，默认是10s
          successThreshold: 0  # 
          failureThreshold: 0
        securityContext:
          privileged: false
        restartPolicy: [Always|Never|OnFailure]  # 容器重启的策略，
        nodeSelector: object  # 指定运行的宿主机
        imagePullSecrets:  # 容器下载时使用的Secrets名称，需要与valumes.secret中定义的一致
          - name: string
        hostNetwork: false
        volumes: ## 挂载的共享存储卷类型
          - name: string  # 非必须，【@1】
          emptyDir: {}
          hostPath:
            path: string
          secret:  # 类型为secret的存储卷，使用内部的secret内的items值作为环境变量
            secrectName: string
            items:
              - key: string
                path: string
            configMap:  ## 类型为configMap的存储卷
              name: string
              items:
                - key: string
                  path: string
</code></pre>

<p>Pod是Kubernetes中最重要的资源，我们可以通过Pod YAML定义来创建一个Pod，也可以通过DaemonSet、Deployment、ReplicaSet、StatefulSet、Job、CronJob来创建Pod。</p>

<h3 id="deployment资源定义">Deployment资源定义</h3>

<p>Deployment资源定义YAML文件如下：</p>

<pre><code class="language-yaml">apiVersion: apps/v1
kind: Deployment
metadata:
  labels: # 设定资源的标签
    app: iam-apiserver
  name: iam-apiserver
  namespace: default
spec:
  progressDeadlineSeconds: 10 # 指定多少时间内不能完成滚动升级就视为失败，滚动升级自动取消
  replicas: 1 # 声明副本数，建议 &gt;= 2
  revisionHistoryLimit: 5 # 设置保留的历史版本个数，默认是10
  selector: # 选择器
    matchLabels: # 匹配标签
      app: iam-apiserver # 标签格式为key: value对
  strategy: # 指定部署策略
    rollingUpdate:
      maxSurge: 1 # 最大额外可以存在的副本数，可以为百分比，也可以为整数
      maxUnavailable: 1 # 表示在更新过程中能够进入不可用状态的 Pod 的最大值，可以为百分比，也可以为整数
    type: RollingUpdate # 更新策略，包括：重建(Recreate)、RollingUpdate(滚动更新)
  template: # 指定Pod创建模板。注意：以下定义为Pod的资源定义
    metadata: # 指定Pod的元数据
      labels: # 指定Pod的标签
        app: iam-apiserver
    spec:
      affinity:
        podAntiAffinity: # Pod反亲和性，尽量避免同一个应用调度到相同Node
          preferredDuringSchedulingIgnoredDuringExecution: # 软需求
          - podAffinityTerm:
              labelSelector:
                matchExpressions: # 有多个选项，只有同时满足这些条件的节点才能运行 Pod
                - key: app
                  operator: In # 设定标签键与一组值的关系，In、NotIn、Exists、DoesNotExist
                  values:
                  - iam-apiserver
              topologyKey: kubernetes.io/hostname
            weight: 100 # weight 字段值的范围是1-100。
      containers:
      - command: # 指定运行命令
        - /opt/iam/bin/iam-apiserver # 运行参数
        - --config=/etc/iam/iam-apiserver.yaml
        image: ccr.ccs.tencentyun.com/lkccc/iam-apiserver-amd64:v1.0.6 # 镜像名，遵守镜像命名规范
        imagePullPolicy: Always # 镜像拉取策略。IfNotPresent：优先使用本地镜像；Never：使用本地镜像，本地镜像不存在，则报错；Always：默认值，每次都重新拉取镜像
        # lifecycle: # kubernetes支持postStart和preStop事件。当一个容器启动后，Kubernetes将立即发送postStart事件；在容器被终结之前，Kubernetes将发送一个preStop事件
        name: iam-apiserver # 容器名称，与应用名称保持一致
        ports: # 端口设置
        - containerPort: 8443 # 容器暴露的端口
          name: secure # 端口名称
          protocol: TCP # 协议，TCP和UDP
        livenessProbe: # 存活检查，检查容器是否正常，不正常则重启实例
          httpGet: # HTTP请求检查方法
            path: /healthz # 请求路径
            port: 8080 # 检查端口
            scheme: HTTP # 检查协议
          initialDelaySeconds: 5 # 启动延时，容器延时启动健康检查的时间
          periodSeconds: 10 # 间隔时间，进行健康检查的时间间隔
          successThreshold: 1 # 健康阈值，表示后端容器从失败到成功的连续健康检查成功次数
          failureThreshold: 1 # 不健康阈值，表示后端容器从成功到失败的连续健康检查成功次数
          timeoutSeconds: 3 # 响应超时，每次健康检查响应的最大超时时间
        readinessProbe: # 就绪检查，检查容器是否就绪，不就绪则停止转发流量到当前实例
          httpGet: # HTTP请求检查方法
            path: /healthz # 请求路径
            port: 8080 # 检查端口
            scheme: HTTP # 检查协议
          initialDelaySeconds: 5 # 启动延时，容器延时启动健康检查的时间
          periodSeconds: 10 # 间隔时间，进行健康检查的时间间隔
          successThreshold: 1 # 健康阈值，表示后端容器从失败到成功的连续健康检查成功次数
          failureThreshold: 1 # 不健康阈值，表示后端容器从成功到失败的连续健康检查成功次数
          timeoutSeconds: 3 # 响应超时，每次健康检查响应的最大超时时间
        startupProbe: # 启动探针，可以知道应用程序容器什么时候启动了
          failureThreshold: 10
          httpGet:
            path: /healthz
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 5
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 3
        resources: # 资源管理
          limits: # limits用于设置容器使用资源的最大上限,避免异常情况下节点资源消耗过多
            cpu: &quot;1&quot; # 设置cpu limit，1核心 = 1000m
            memory: 1Gi # 设置memory limit，1G = 1024Mi
          requests: # requests用于预分配资源,当集群中的节点没有request所要求的资源数量时,容器会创建失败
            cpu: 250m # 设置cpu request
            memory: 500Mi # 设置memory request
        terminationMessagePath: /dev/termination-log # 容器终止时消息保存路径
        terminationMessagePolicy: File # 仅从终止消息文件中检索终止消息
        volumeMounts: # 挂载日志卷
        - mountPath: /etc/iam/iam-apiserver.yaml # 容器内挂载镜像路径
          name: iam # 引用的卷名称
          subPath: iam-apiserver.yaml # 指定所引用的卷内的子路径，而不是其根路径。
        - mountPath: /etc/iam/cert
          name: iam-cert
      dnsPolicy: ClusterFirst
      restartPolicy: Always # 重启策略，Always、OnFailure、Never
      schedulerName: default-scheduler # 指定调度器的名字
      imagePullSecrets: # 在Pod中设置ImagePullSecrets只有提供自己密钥的Pod才能访问私有仓库
        - name: ccr-registry # 镜像仓库的Secrets需要在集群中手动创建
      securityContext: {} # 指定安全上下文
      terminationGracePeriodSeconds: 5 # 优雅关闭时间，这个时间内优雅关闭未结束，k8s 强制 kill
      volumes: # 配置数据卷，类型详见https://kubernetes.io/zh/docs/concepts/storage/volumes
      - configMap: # configMap 类型的数据卷
          defaultMode: 420 #权限设置0~0777，默认0664
          items:
          - key: iam-apiserver.yaml
            path: iam-apiserver.yaml
          name: iam # configmap名称
        name: iam # 设置卷名称，与volumeMounts名称对应
      - configMap:
          defaultMode: 420
          name: iam-cert
        name: iam-cert
</code></pre>

<p>在部署时，你可以根据需要来配置相应的字段，常见的需要配置的字段为：<code>labels</code>、<code>name</code>、<code>namespace</code>、<code>replicas</code>、<code>command</code>、<code>imagePullPolicy</code>、<code>container.name</code>、<code>livenessProbe</code>、<code>readinessProbe</code>、<code>resources</code>、<code>volumeMounts</code>、<code>volumes</code>、<code>imagePullSecrets</code>等。</p>

<p>另外，在部署应用时，经常需要提供配置文件，供容器内的进程加载使用。最常用的方法是挂载ConfigMap到应用容器中。那么，如何挂载ConfigMap到容器中呢？</p>

<p>引用 ConfigMap 对象时，你可以在 volume 中通过它的名称来引用。你可以自定义 ConfigMap 中特定条目所要使用的路径。下面的配置就显示了如何将名为 <code>log-config</code> 的 ConfigMap 挂载到名为 <code>configmap-pod</code> 的 Pod 中：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: Pod
metadata:
  name: configmap-pod
spec:
  containers:
    - name: test
      image: busybox
      volumeMounts:
        - name: config-vol
          mountPath: /etc/config
  volumes:
    - name: config-vol
      configMap:
        name: log-config
        items:
          - key: log_level
            path: log_level
</code></pre>

<p><code>log-config</code> ConfigMap 以卷的形式挂载，并且存储在 <code>log_level</code> 条目中的所有内容都被挂载到 Pod 的<code>/etc/config/log_level</code> 路径下。 请注意，这个路径来源于卷的 <code>mountPath</code> 和 <code>log_level</code> 键对应的<code>path</code>。</p>

<p>这里需要注意，在使用 ConfigMap 之前，你首先要创建它。接下来，我们来看下ConfigMap定义。</p>

<h3 id="configmap资源定义">ConfigMap资源定义</h3>

<p>下面是一个ConfigMap YAML示例：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: ConfigMap
metadata:
  name: test-config4
data: # 存储配置内容
  db.host: 172.168.10.1 # 存储格式为key: value
  db.port: 3306
</code></pre>

<p>可以看到，ConfigMap的YAML定义相对简单些。假设我们将上述YAML文件保存在了<code>iam-configmap.yaml</code>文件中，我们可以执行以下命令，来创建ConfigMap：</p>

<pre><code class="language-bash">$ kubectl create -f iam-configmap.yaml
</code></pre>

<p>除此之外，kubectl命令行工具还提供了3种创建ConfigMap的方式。我来分别介绍下。</p>

<p>1）通过<code>--from-literal</code>参数创建</p>

<p>创建命令如下：</p>

<pre><code class="language-bash">$ kubectl create configmap iam-configmap --from-literal=db.host=172.168.10.1 --from-literal=db.port='3306'
</code></pre>

<p>2）通过<code>--from-file=&lt;文件&gt;</code>参数创建</p>

<p>创建命令如下：</p>

<pre><code class="language-bash">$ echo -n 172.168.10.1 &gt; ./db.host
$ echo -n 3306 &gt; ./db.port
$ kubectl create cm iam-configmap --from-file=./db.host --from-file=./db.port
</code></pre>

<p><code>--from-file</code>的值也可以是一个目录。当值是目录时，目录中的文件名为key，目录的内容为value。</p>

<p>3）通过<code>--from-env-file</code>参数创建</p>

<p>创建命令如下：</p>

<pre><code class="language-bash">$ cat &lt;&lt; EOF &gt; env.txt
db.host=172.168.10.1
db.port=3306
EOF
$ kubectl create cm iam-configmap --from-env-file=env.txt
</code></pre>

<h3 id="service资源定义">Service资源定义</h3>

<p>Service 是 Kubernetes 另一个核心资源。通过创建 Service，可以为一组具有相同功能的容器应用提供一个统一的入口地址，并且将请求负载到后端的各个容器上。Service资源定义YAML文件如下：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: Service
metadata:
  labels:
    app: iam-apiserver
  name: iam-apiserver
  namespace: default
spec:
  clusterIP: 192.168.0.231 # 虚拟服务地址
  externalTrafficPolicy: Cluster # 表示此服务是否希望将外部流量路由到节点本地或集群范围的端点
  ports: # service需要暴露的端口列表
  - name: https #端口名称
    nodePort: 30443 # 当type = NodePort时，指定映射到物理机的端口号
    port: 8443 # 服务监听的端口号
    protocol: TCP # 端口协议，支持TCP和UDP，默认TCP
    targetPort: 8443 # 需要转发到后端Pod的端口号
  selector: # label selector配置，将选择具有label标签的Pod作为其后端RS
    app: iam-apiserver
  sessionAffinity: None # 是否支持session
  type: NodePort # service的类型，指定service的访问方式，默认为clusterIp
</code></pre>

<p>上面，我介绍了常用的Kubernetes YAML的内容。我们在部署应用的时候，是需要手动编写这些文件的。接下来，我就讲解一些在编写过程中常用的编写技巧。</p>

<h2 id="yaml文件编写技巧">YAML文件编写技巧</h2>

<p>这里我主要介绍三个技巧。</p>

<p>1）使用在线的工具来自动生成模板YAML文件。</p>

<p>YAML文件很复杂，完全从0开始编写一个YAML定义文件，工作量大、容易出错，也没必要。我比较推荐的方式是，使用一些工具来自动生成所需的YAML。</p>

<p>这里我推荐使用<a href="https://k8syaml.com/" target="_blank">k8syaml</a>工具。<code>k8syaml</code>是一个在线的YAML生成工具，当前能够生成Deployment、StatefulSet、DaemonSet类型的YAML文件。<code>k8syaml</code>具有默认值，并且有对各字段详细的说明，可以供我们填参时参考。</p>

<p>2）使用<code>kubectl run</code>命令获取YAML模板：</p>

<pre><code class="language-yaml">$ kubectl run --dry-run=client --image=nginx nginx -o yaml &gt; my-nginx.yaml
$ cat my-nginx.yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Always
status: {}
</code></pre>

<p>然后，我们可以基于这个模板，来修改配置，形成最终的YAML文件。</p>

<p>3）导出集群中已有的资源描述。</p>

<p>有时候，如果我们想创建一个Kubernetes资源，并且发现该资源跟集群中已经创建的资源描述相近或者一致的时候，可以选择导出集群中已经创建资源的YAML描述，并基于导出的YAML文件进行修改，获得所需的YAML。例如：</p>

<pre><code class="language-bash">$ kubectl get deployment iam-apiserver -o yaml &gt; iam-authz-server.yaml
</code></pre>

<p>接着，修改<code>iam-authz-server.yaml</code>。通常，我们需要删除Kubernetes自动添加的字段，例如<code>kubectl.kubernetes.io/last-applied-configuration</code>、<code>deployment.kubernetes.io/revision</code>、<code>creationTimestamp</code>、<code>generation</code>、<code>resourceVersion</code>、<code>selfLink</code>、<code>uid</code>、<code>status</code>。</p>

<p>这些技巧可以帮助我们更好地编写和使用Kubernetes YAML。</p>

<h2 id="使用kubernetes-yaml时的一些推荐工具">使用Kubernetes YAML时的一些推荐工具</h2>

<p>接下来，我再介绍一些比较流行的工具，你可以根据自己的需要进行选择。</p>

<h3 id="kubeval">kubeval</h3>

<p><a href="https://github.com/instrumenta/kubeval" target="_blank">kubeval</a>可以用来验证Kubernetes YAML是否符合Kubernetes API模式。</p>

<p>安装方法如下：</p>

<pre><code class="language-bash">$ wget https://github.com/instrumenta/kubeval/releases/latest/download/kubeval-linux-amd64.tar.gz
$ tar xf kubeval-linux-amd64.tar.gz
$ mv kubeval $HOME/bin
</code></pre>

<p>安装完成后，我们对Kubernetes YAML文件进行验证：</p>

<pre><code class="language-bash">$ kubeval deployments/iam.invalid.yaml
ERR  - iam/templates/iam-configmap.yaml: Duplicate 'ConfigMap' resource 'iam' in namespace ''
</code></pre>

<p>根据提示，查看<code>iam.yaml</code>，发现在<code>iam.yaml</code>文件中，我们定义了两个同名的<code>iam</code> ConfigMap：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: ConfigMap
metadata:
  name: iam
data:
  {}
---
# Source: iam/templates/iam-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: iam
data:
  iam-: &quot;&quot;
  iam-apiserver.yaml: |
    ...
</code></pre>

<p>可以看到，使用<code>kubeval</code>之类的工具，能让我们在部署的早期，不用访问集群就能发现YAML文件的错误。</p>

<h3 id="kube-score">kube-score</h3>

<p><a href="https://github.com/zegl/kube-score" target="_blank">kube-score</a>能够对Kubernetes YAML进行分析，并根据内置的检查对其评分，这些检查是根据安全建议和最佳实践而选择的，例如：</p>

<ul>
<li>以非Root用户启动容器。</li>
<li>为Pods设置健康检查。</li>
<li>定义资源请求和限制。</li>
</ul>

<p>你可以按照这个方法安装：</p>

<pre><code class="language-bash">$ go get github.com/zegl/kube-score/cmd/kube-score
</code></pre>

<p>然后，我们对Kubernetes YAML进行评分：</p>

<pre><code class="language-bash">$ kube-score score -o ci deployments/iam.invalid.yaml
[OK] iam-apiserver apps/v1/Deployment
[OK] iam-apiserver apps/v1/Deployment
[OK] iam-apiserver apps/v1/Deployment
[OK] iam-apiserver apps/v1/Deployment
[CRITICAL] iam-apiserver apps/v1/Deployment: The pod does not have a matching NetworkPolicy
[CRITICAL] iam-apiserver apps/v1/Deployment: Container has the same readiness and liveness probe
[CRITICAL] iam-apiserver apps/v1/Deployment: (iam-apiserver) The pod has a container with a writable root filesystem
[CRITICAL] iam-apiserver apps/v1/Deployment: (iam-apiserver) The container is running with a low user ID
[CRITICAL] iam-apiserver apps/v1/Deployment: (iam-apiserver) The container running with a low group ID
[OK] iam-apiserver apps/v1/Deployment
...
</code></pre>

<p>检查的结果有<code>OK</code>、<code>SKIPPED</code>、<code>WARNING</code>和<code>CRITICAL</code>。<code>CRITICAL</code>是需要你修复的；<code>WARNING</code>是需要你关注的；<code>SKIPPED</code>是因为某些原因略过的检查；<code>OK</code>是验证通过的。</p>

<p>如果你想查看详细的错误原因和解决方案，可以使用<code>-o human</code>选项，例如：</p>

<pre><code class="language-bash">$ kube-score score -o human deployments/iam.invalid.yaml
</code></pre>

<p>上述命令会检查YAML资源定义文件，如果有不合规的地方会报告级别、类别以及错误详情，如下图所示：</p>

<p><img src="assets/948084c6067948fa86fd6a286eb511c6.jpg" alt="图片" /></p>

<p>当然，除了kubeval、kube-score这两个工具，业界还有其他一些Kubernetes检查工具，例如<a href="https://github.com/stelligent/config-lint" target="_blank">config-lint</a>、<a href="https://github.com/cloud66-oss/copper" target="_blank">copper</a>、<a href="https://github.com/open-policy-agent/conftest" target="_blank">conftest</a>、<a href="https://github.com/FairwindsOps/polaris" target="_blank">polaris</a>等。</p>

<p>这些工具，我推荐你这么来选择：首先，使用kubeval工具做最基本的YAML文件验证。验证通过之后，我们就可以进行更多的测试。如果你没有特别复杂的YAML验证要求，只需要用到一些最常见的检查策略，这时候可以使用kube-score。如果你有复杂的验证要求，并且希望能够自定义验证策略，则可以考虑使用copper。当然，<code>polaris</code>、<code>config-lint</code>、<code>copper</code>也值得你去尝试下。</p>

<h2 id="总结">总结</h2>

<p>今天，我主要讲了如何编写Kubernetes YAML文件。</p>

<p>YAML格式具有丰富的数据表达能力、清晰的结构和层次，因此被用于Kubernetes资源的定义文件中。如果你要把应用部署在Kubernetes集群中，就要创建多个关联的K8s资源，如果要创建K8s资源，目前比较多的方式还是编写YAML格式的定义文件。</p>

<p>这一讲我介绍了K8s中最常用的四种资源（Pod、Deployment、Service、ConfigMap）的YAML定义的写法，你可以常来温习。</p>

<p>另外，在编写YAML文件时，也有一些技巧。比如，可以通过在线工具<a href="https://k8syaml.com/" target="_blank">k8syaml</a>来自动生成初版的YAML文件，再基于此YAML文件进行二次修改，从而形成终版。</p>

<p>最后，我还给你分享了编写和使用Kubernetes YAML时，社区提供的多种工具。比如，kubeval可以校验YAML，kube-score可以给YAML文件打分。了解了如何编写Kubernetes YAML文件，下一讲的学习相信你会进行得更顺利。</p>

<h2 id="课后练习">课后练习</h2>

<ol>
<li>思考一下，如何将ConfigMap中的Key挂载到同一个目录中，文件名为Key名？</li>
<li>使用kubeval检查你正在或之前从事过的项目的K8s YAML定义文件，查看报错，并修改和优化。</li>
</ol>

<p>欢迎你在留言区和我交流讨论，我们下一讲见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#b2dedede8b8683838285f2d5dfd3dbde9cd1dddf" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0df8b55937653b',t:'MTczNDAwOTIwNC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>