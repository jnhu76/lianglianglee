<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=50&#32;服务编排（下）：基于Helm的服务编排部署实战>
        <link rel="icon" href="/static/favicon.png">
        <title>50 服务编排（下）：基于Helm的服务编排部署实战 </title>
        
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
                            <h1 id="title" data-id="50 服务编排（下）：基于Helm的服务编排部署实战" class="title">50 服务编排（下）：基于Helm的服务编排部署实战</h1>
                            <div><p>你好，我是孔令飞。</p>

<p>上一讲，我介绍了 Helm 的基础知识，并带着你部署了一个简单的应用。掌握Helm的基础知识之后，今天我们就来实战下，一起通过Helm部署一个IAM应用。</p>

<p>通过Helm部署IAM应用，首先需要制作IAM Chart包，然后通过Chart包来一键部署IAM应用。在实际开发中，我们需要将应用部署在不同的环境中，所以我也会给你演示下如何在多环境中部署IAM应用。</p>

<h2 id="制作iam-chart包">制作IAM Chart包</h2>

<p>在部署IAM应用之前，我们首先需要制作一个IAM Chart包。</p>

<p>我们假设IAM项目源码根目录为<code>${IAM_ROOT}</code>，进入 <code>${IAM_ROOT}/deployments</code>目录，在该目录下创建Chart包。具体创建流程分为四个步骤，下面我来详细介绍下。</p>

<p><strong>第一步，</strong>创建一个模板Chart。</p>

<p>Chart是一个组织在文件目录中的集合，目录名称就是Chart名称（没有版本信息）。你可以看看这个 <a href="https://helm.sh/zh/docs/topics/charts" target="_blank">Chart 开发指南</a> ，它介绍了如何开发你自己的Chart。</p>

<p>不过，这里你也可以使用 <code>helm create</code> 命令来快速创建一个模板Chart，并基于该Chart进行修改，得到你自己的Chart。创建命令如下：</p>

<pre><code class="language-bash">$ helm create iam
</code></pre>

<p><code>helm create iam</code>会在当前目录下生成一个<code>iam</code>目录，里面存放的就是Chart文件。Chart目录结构及文件如下：</p>

<pre><code class="language-bash">$ tree -FC iam/
├── charts/                            # [可选]: 该目录中放置当前Chart依赖的其他Chart
├── Chart.yaml                         # YAML文件，用于描述Chart的基本信息，包括名称版本等
├── templates/                         # [可选]: 部署文件模版目录，模版使用的值来自values.yaml和由Tiller提供的值
│   ├── deployment.yaml                # Kubernetes Deployment object
│   ├── _helpers.tpl                   # 用于修改Kubernetes objcet配置的模板
│   ├── hpa.yaml                       # Kubernetes HPA object
│   ├── ingress.yaml                   # Kubernetes Ingress object
│   ├── NOTES.txt                      # [可选]: 放置Chart的使用指南
│   ├── serviceaccount.yaml
│   ├── service.yaml
│   └── tests/                         # 定义了一些测试资源
│       └── test-connection.yaml
└── values.yaml                        # Chart的默认配置文件
</code></pre>

<p>上面的目录中，有两个比较重要的文件：</p>

<ul>
<li>Chart.yaml 文件</li>
<li>templates目录</li>
</ul>

<p>下面我来详细介绍下这两个文件。<strong>我们先来看Chart.yaml 文件。</strong></p>

<p>Chart.yaml用来描述Chart的基本信息，包括名称、版本等，内容如下：</p>

<pre><code class="language-yaml">apiVersion: Chart API 版本 （必需）
name: Chart名称 （必需）
version: 语义化版本（必需）
kubeVersion: 兼容Kubernetes版本的语义化版本（可选）
description: 对这个项目的一句话描述（可选）
type: Chart类型 （可选）
keywords:
  - 关于项目的一组关键字（可选）
home: 项目home页面的URL （可选）
sources:
  - 项目源码的URL列表（可选）
dependencies: # chart 必要条件列表 （可选）
  - name: Chart名称 (nginx)
    version: Chart版本 (&quot;1.2.3&quot;)
    repository: （可选）仓库URL (&quot;https://example.com/charts&quot;) 或别名 (&quot;@repo-name&quot;)
    condition: （可选） 解析为布尔值的YAML路径，用于启用/禁用Chart(e.g. subchart1.enabled )
    tags: # （可选）
      - 用于一次启用/禁用 一组Chart的tag
    import-values: # （可选）
      - ImportValue 保存源值到导入父键的映射。每项可以是字符串或者一对子/父列表项
    alias: （可选） Chart中使用的别名。当你要多次添加相同的Chart时会很有用
maintainers: # （可选）
  - name: 维护者名字 （每个维护者都需要）
    email: 维护者邮箱 （每个维护者可选）
    url: 维护者URL （每个维护者可选）
icon: 用作icon的SVG或PNG图片URL （可选）
appVersion: 包含的应用版本（可选）。不需要是语义化，建议使用引号
deprecated: 不被推荐的Chart（可选，布尔值）
annotations:
  example: 按名称输入的批注列表 （可选）.
</code></pre>

<p><strong>我们再来看下</strong><strong>templates目录</strong><strong>这个文件。</strong></p>

<p>templates目录中包含了应用中各个Kubernetes资源的YAML格式资源定义模板，例如：</p>

<pre><code class="language-yaml">apiVersion: v1
kind: Service
metadata:
  labels:
    app: {{ .Values.pump.name }}
  name: {{ .Values.pump.name }}
spec:
  ports:
  - name: http
    protocol: TCP
    {{- toYaml .Values.pump.service.http| nindent 4 }}
  selector:
    app: {{ .Values.pump.name }}
  sessionAffinity: None
  type: {{ .Values.serviceType }}
</code></pre>

<p><code>{{ .Values.pump.name }}</code>会被<code>deployments/iam/values.yaml</code>文件中<code>pump.name</code>的值替换。上面的模版语法扩展了 Go <code>text/template</code>包的语法：</p>

<pre><code class="language-yaml"># 这种方式定义的模版，会去除test模版尾部所有的空行
{{- define &quot;test&quot;}}
模版内容
{{- end}}

# 去除test模版头部的第一个空行
{{- template &quot;test&quot; }}
</code></pre>

<p>下面是用于YAML文件前置空格的语法：</p>

<pre><code class="language-bash"># 这种方式定义的模版，会去除test模版头部和尾部所有的空行
{{- define &quot;test&quot; -}}
模版内容
{{- end -}}

# 可以在test模版每一行的头部增加4个空格，用于YAML文件的对齐
{{ include &quot;test&quot; | indent 4}}
</code></pre>

<p>最后，这里有三点需要你注意：</p>

<ul>
<li>Chart名称必须是小写字母和数字，单词之间可以使用横杠<code>-</code>分隔，Chart名称中不能用大写字母，也不能用下划线，<code>.</code>号也不行。</li>
<li>尽可能使用<a href="https://semver.org/" target="_blank">SemVer 2</a>来表示版本号。</li>
<li>YAML 文件应该按照双空格的形式缩进(一定不要使用tab键)。</li>
</ul>

<p><strong>第二步，</strong>编辑 <code>iam</code> 目录下的Chart文件。</p>

<p>我们可以基于<code>helm create</code>生成的模板Chart来构建自己的Chart包。这里我们添加了创建iam-apiserver、iam-authz-server、iam-pump、iamctl服务需要的YAML格式的Kubernetes资源文件模板：</p>

<pre><code class="language-bash">$ ls -1 iam/templates/*.yaml
iam/templates/hpa.yaml                                   # Kubernetes HPA模板文件
iam/templates/iam-apiserver-deployment.yaml              # iam-apiserver服务deployment模板文件
iam/templates/iam-apiserver-service.yaml                 # iam-apiserver服务service模板文件
iam/templates/iam-authz-server-deployment.yaml           # iam-authz-server服务deployment模板文件
iam/templates/iam-authz-server-service.yaml              # iam-authz-server服务service模板文件
iam/templates/iamctl-deployment.yaml                     # iamctl服务deployment模板文件
iam/templates/iam-pump-deployment.yaml                   # iam-pump服务deployment模板文件
iam/templates/iam-pump-service.yaml                      # iam-pump服务service模板文件
</code></pre>

<p>模板的具体内容，你可以查看<a href="https://github.com/marmotedu/iam/tree/v1.1.0/deployments/iam/templates" target="_blank">deployments/iam/templates/</a>。</p>

<p>在编辑 Chart 时，我们可以通过 <code>helm lint</code> 验证格式是否正确，例如：</p>

<pre><code class="language-bash">$ helm lint iam
==&gt; Linting iam

1 chart(s) linted, 0 chart(s) failed
</code></pre>

<p><code>0 chart(s) failed</code> 说明当前Iam Chart包是通过校验的。</p>

<p><strong>第三步，</strong>修改Chart的配置文件，添加自定义配置。</p>

<p>我们可以编辑<code>deployments/iam/values.yaml</code>文件，定制自己的配置。具体配置你可以参考<a href="https://github.com/marmotedu/iam/blob/v1.1.0/deployments/iam/values.yaml" target="_blank">deployments/iam/values.yaml</a>。</p>

<p>在修改 <code>values.yaml</code> 文件时，你可以参考下面这些最佳实践。</p>

<ul>
<li>变量名称以小写字母开头，单词按驼峰区分，例如<code>chickenNoodleSoup</code>。</li>
<li>给所有字符串类型的值加上引号。</li>
<li>为了避免整数转换问题，将整型存储为字符串更好，并用 <code>{{ int $value }}</code> 在模板中将字符串转回整型。</li>
<li><code>values.yaml</code>中定义的每个属性都应该文档化。文档字符串应该以它要描述的属性开头，并至少给出一句描述。例如：</li>
</ul>

<pre><code class="language-yaml"># serverHost is the host name for the webserver
serverHost: example
# serverPort is the HTTP listener port for the webserver
serverPort: 9191
</code></pre>

<p>这里需要注意，所有的Helm内置变量都以大写字母开头，以便与用户定义的value进行区分，例如<code>.Release.Name</code>、<code>.Capabilities.KubeVersion</code>。</p>

<p>为了安全，values.yaml中只配置Kubernetes资源相关的配置项，例如Deployment副本数、Service端口等。至于iam-apiserver、iam-authz-server、iam-pump、iamctl组件的配置文件，我们创建单独的ConfigMap，并在Deployment中引用。</p>

<p><strong>第四步，</strong>打包Chart，并上传到Chart仓库中。</p>

<p>这是一个可选步骤，可以根据你的实际需要来选择。如果想了解具体操作，你可以查看 <a href="https://helm.sh/zh/docs/topics/chart_repository" target="_blank">Helm chart 仓库</a>获取更多信息。</p>

<p>最后，IAM应用的Chart包见<a href="https://github.com/marmotedu/iam/tree/v1.1.0/deployments/iam" target="_blank">deployments/iam</a>。</p>

<h2 id="iam-chart部署">IAM Chart部署</h2>

<p>上面，我们制作了IAM应用的Chart包，接下来我们就使用这个Chart包来一键创建IAM应用。IAM Chart部署一共分为10个步骤，你可以跟着我一步步操作下。</p>

<p><strong>第一步，</strong>配置<code>scripts/install/environment.sh</code>。</p>

<p><code>scripts/install/environment.sh</code>文件中包含了各类自定义配置，你主要配置下面这些跟数据库相关的就可以，其他配置使用默认值。</p>

<ul>
<li>MariaDB配置：environment.sh文件中以<code>MARIADB_</code>开头的变量。</li>
<li>Redis配置：environment.sh文件中以<code>REDIS_</code>开头的变量。</li>
<li>MongoDB配置：environment.sh文件中以<code>MONGO_</code>开头的变量。</li>
</ul>

<p><strong>第二步，</strong>创建IAM应用的配置文件。</p>

<pre><code class="language-bash">$ cd ${IAM_ROOT}
$ make gen.defaultconfigs # 生成iam-apiserver、iam-authz-server、iam-pump、iamctl组件的默认配置文件
$ make gen.ca # 生成 CA 证书
</code></pre>

<p>上面的命令会将IAM的配置文件存放在目录<code>${IAM_ROOT}/_output/configs/</code>下。</p>

<p><strong>第三步，</strong>创建 <code>iam</code> 命名空间。</p>

<p>我们将IAM应用涉及到的各类资源都创建在<code>iam</code>命名空间中。将IAM资源创建在独立的命名空间中，不仅方便维护，还可以有效避免影响其他Kubernetes资源。</p>

<pre><code class="language-bash">$ kubectl create namespace iam
</code></pre>

<p><strong>第四步，</strong>将IAM各服务的配置文件，以ConfigMap资源的形式保存在Kubernetes集群中。</p>

<pre><code class="language-bash">$ kubectl -n iam create configmap iam --from-file=${IAM_ROOT}/_output/configs/
$ kubectl -n iam get configmap iam
NAME   DATA   AGE
iam    4      13s
</code></pre>

<p><strong>第五步，</strong>将IAM各服务使用的证书文件，以ConfigMap资源的形式保存在Kubernetes集群中。</p>

<pre><code class="language-bash">$ kubectl -n iam create configmap iam-cert --from-file=${IAM_ROOT}/_output/cert
$ kubectl -n iam get configmap iam-cert
NAME       DATA   AGE
iam-cert   14     12s
</code></pre>

<p><strong>第六步，</strong>创建镜像仓库访问密钥。</p>

<p>在准备阶段，我们开通了<a href="http://ccr.ccs.tencentyun.com" target="_blank">腾讯云镜像仓库服务</a>，并创建了用户<code>10000099</code><code>xxxx</code>，密码为<code>iam59!z$</code>。</p>

<p>接下来，我们就可以创建docker-registry secret了。Kubernetes在下载Docker镜像时，需要docker-registry secret来进行认证。创建命令如下：</p>

<pre><code class="language-bash">$ kubectl -n iam create secret docker-registry ccr-registry --docker-server=ccr.ccs.tencentyun.com --docker-username=10000099xxxx --docker-password='iam59!z$'
</code></pre>

<p><strong>第七步，</strong>创建Docker镜像，并Push到镜像仓库。</p>

<pre><code class="language-bash">$ make push REGISTRY_PREFIX=ccr.ccs.tencentyun.com/marmotedu VERSION=v1.1.0
</code></pre>

<p><strong>第八步，</strong>安装IAM Chart包。</p>

<p>在<a href="https://time.geekbang.org/column/article/420940" target="_blank">49讲</a>里，我介绍了4种安装Chart包的方法。这里，我们通过未打包的IAM Chart路径来安装，安装方法如下：</p>

<pre><code class="language-bash">$ cd ${IAM_ROOT}
$ helm -n iam install iam deployments/iam
NAME: iam
LAST DEPLOYED: Sat Aug 21 17:46:56 2021
NAMESPACE: iam
STATUS: deployed
REVISION: 1
TEST SUITE: None
</code></pre>

<p>执行 <code>helm install</code> 后，Kubernetes会自动部署应用，等到IAM应用的Pod都处在 <code>Running</code> 状态时，说明IAM应用已经成功安装：</p>

<pre><code class="language-bash">$ kubectl -n iam get pods|grep iam
iam-apiserver-cb4ff955-hs827        1/1     Running   0          66s
iam-authz-server-7fccc7db8d-chwnn   1/1     Running   0          66s
iam-pump-78b57b4464-rrlbf           1/1     Running   0          66s
iamctl-59fdc4995-xrzhn              1/1     Running   0          66s
</code></pre>

<p><strong>第九步，</strong>测试IAM应用。</p>

<p>我们通过<code>helm install</code>在<code>iam</code>命令空间下创建了一个测试Deployment <code>iamctl</code>。你可以登陆<code>iamctl</code> Deployment所创建出来的Pod，执行一些运维操作和冒烟测试。登陆命令如下：</p>

<pre><code class="language-bash">$ kubectl -n iam exec -it `kubectl -n iam get pods -l app=iamctl | awk '/iamctl/{print $1}'` -- bash
</code></pre>

<p>登陆到<code>iamctl-xxxxxxxxxx-xxxxx</code> Pod中后，你就可以执行运维操作和冒烟测试了。</p>

<p><strong>先来看运维操作。</strong>iamctl工具以子命令的方式对外提供功能，你可以使用它提供的各类功能，如下图所示：</p>

<p><img src="assets/48f606aebdb2479fb5ca391db04537f8.jpg" alt="图片" /></p>

<p><strong>再来看冒烟测试：</strong></p>

<pre><code class="language-bash"># cd /opt/iam/scripts/install
# ./test.sh iam::test::smoke
</code></pre>

<p>如果<code>./test.sh iam::test::smoke</code>命令打印的输出中，最后一行为<code>congratulations, smoke test passed!</code>字符串，就说明IAM应用安装成功。如下图所示：</p>

<p><img src="assets/9df94e04d1d94b478e6ea191bc9bb2b5.jpg" alt="图片" /></p>

<p><strong>第十步，</strong>销毁EKS集群的资源。</p>

<pre><code class="language-bash">$ kubectl delete namespace iam
</code></pre>

<p>你可以根据需要选择是否删除EKS集群，如果不需要了就可以选择删除。</p>

<h2 id="iam应用多环境部署">IAM应用多环境部署</h2>

<p>在实际的项目开发中，我们需要将IAM应用部署到不同的环境中，不同环境的配置文件是不同的，那么IAM项目是如何进行多环境部署的呢？</p>

<p>IAM项目在configs目录下创建了多个Helm values文件（格式为<code>values-{envName}-env.yaml</code>）：</p>

<ul>
<li>values-test-env.yaml，测试环境Helm values文件。</li>
<li>values-pre-env.yaml，预发环境Helm values文件。</li>
<li>values-prod-env.yaml，生产环境Helm values文件。</li>
</ul>

<p>在部署IAM应用时，我们在命令行指定<code>-f</code>参数，例如：</p>

<pre><code class="language-bash">$ helm -n iam install -f configs/values-test-env.yaml iam deployments/iam # 安装到测试环境。
</code></pre>

<h2 id="总结">总结</h2>

<p>这一讲，我们通过 <code>helm create iam</code> 创建了一个模板Chart，并基于这个模板Chart包进行了二次开发，最终创建了IAM应用的Helm Chart包：<a href="https://github.com/marmotedu/iam/tree/v1.1.0/deployments/iam" target="_blank">deployments/iam</a>。</p>

<p>有了Helm Chart包，我们就可以通过 <code>helm -n iam install iam deployments/iam</code> 命令来一键部署好整个IAM应用。当IAM应用中的所有Pod都处在 <code>Running</code> 状态后，说明IAM应用被成功部署。</p>

<p>最后，我们可以登录iamctl容器，执行 <code>test.sh iam::test::smoke</code> 命令，来对IAM应用进行冒烟测试。</p>

<h2 id="课后练习">课后练习</h2>

<ol>
<li>试着在Helm Chart中加入MariaDB、MongoDB、Redis模板，通过Helm一键部署好整个IAM应用。</li>
<li>试着通过 <code>helm</code> 命令升级、回滚和删除IAM应用。</li>
</ol>

<p>欢迎你在留言区与我交流讨论，我们下一讲见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#acc0c0c095989d9d9c9beccbc1cdc5c082cfc3c1" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0df9aa4ac4653b',t:'MTczNDAwOTI0My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>