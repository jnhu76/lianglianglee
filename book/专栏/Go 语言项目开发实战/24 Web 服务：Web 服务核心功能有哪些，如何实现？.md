<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=24&#32;Web&#32;服务：Web&#32;服务核心功能有哪些，如何实现？>
        <link rel="icon" href="/static/favicon.png">
        <title>24 Web 服务：Web 服务核心功能有哪些，如何实现？ </title>
        
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
                            <h1 id="title" data-id="24 Web 服务：Web 服务核心功能有哪些，如何实现？" class="title">24 Web 服务：Web 服务核心功能有哪些，如何实现？</h1>
                            <div><p>你好，我是孔令飞。从今天开始，我们进入实战第三站：服务开发。在这个部分，我会讲解 IAM项目各个服务的构建方式，帮助你掌握Go 开发阶段的各个技能点。</p>

<p>在Go项目开发中，绝大部分情况下，我们是在写能提供某种功能的后端服务，这些功能以RPC API 接口或者RESTful API接口的形式对外提供，能提供这两种API接口的服务也统称为Web服务。今天这一讲，我就通过介绍RESTful API风格的Web服务，来给你介绍下如何实现Web服务的核心功能。</p>

<p>那今天我们就来看下，Web服务的核心功能有哪些，以及如何开发这些功能。</p>

<h2 id="web服务的核心功能">Web服务的核心功能</h2>

<p>Web服务有很多功能，为了便于你理解，我将这些功能分成了基础功能和高级功能两大类，并总结在了下面这张图中：</p>

<p><img src="assets/a9d0d1b4fd8b4b5c85edd7e8f2a5edee.jpg" alt="" /></p>

<p>下面，我就按图中的顺序，来串讲下这些功能。</p>

<p>要实现一个Web服务，首先我们要选择通信协议和通信格式。在Go项目开发中，有HTTP+JSON 和 gRPC+Protobuf两种组合可选。因为iam-apiserver主要提供的是REST风格的API接口，所以选择的是HTTP+JSON组合。</p>

<p><strong>Web服务最核心的功能是路由匹配。</strong>路由匹配其实就是根据<code>(HTTP方法, 请求路径)</code>匹配到处理这个请求的函数，最终由该函数处理这次请求，并返回结果，过程如下图所示：</p>

<p><img src="assets/860b990aec28491ea03e3854d3b79fad.jpg" alt="" /></p>

<p>一次HTTP请求经过路由匹配，最终将请求交由<code>Delete(c *gin.Context)</code>函数来处理。变量<code>c</code>中存放了这次请求的参数，在Delete函数中，我们可以进行参数解析、参数校验、逻辑处理，最终返回结果。</p>

<p>对于大型系统，可能会有很多个API接口，API接口随着需求的更新迭代，可能会有多个版本，为了便于管理，我们需要<strong>对路由进行分组</strong>。</p>

<p>有时候，我们需要在一个服务进程中，同时开启HTTP服务的80端口和HTTPS的443端口，这样我们就可以做到：对内的服务，访问80端口，简化服务访问复杂度；对外的服务，访问更为安全的HTTPS服务。显然，我们没必要为相同功能启动多个服务进程，所以这时候就需要Web服务能够支持<strong>一进程多服务</strong>的功能。</p>

<p>我们开发Web服务最核心的诉求是：输入一些参数，校验通过后，进行业务逻辑处理，然后返回结果。所以Web服务还应该能够进行<strong>参数解析</strong>、<strong>参数校验</strong>、<strong>逻辑处理</strong>、<strong>返回结果</strong>。这些都是Web服务的业务处理功能。</p>

<p>上面这些是Web服务的基本功能，此外，我们还需要支持一些高级功能。</p>

<p>在进行HTTP请求时，经常需要针对每一次请求都设置一些通用的操作，比如添加Header、添加RequestID、统计请求次数等，这就要求我们的Web服务能够支持<strong>中间件</strong>特性。</p>

<p>为了保证系统安全，对于每一个请求，我们都需要进行<strong>认证</strong>。Web服务中，通常有两种认证方式，一种是基于用户名和密码，一种是基于Token。认证通过之后，就可以继续处理请求了。</p>

<p>为了方便定位和跟踪某一次请求，需要支持<strong>RequestID</strong>，定位和跟踪RequestID主要是为了排障。</p>

<p>最后，当前的软件架构中，很多采用了前后端分离的架构。在前后端分离的架构中，前端访问地址和后端访问地址往往是不同的，浏览器为了安全，会针对这种情况设置跨域请求，所以Web服务需要能够处理浏览器的<strong>跨域</strong>请求。</p>

<p>到这里，我就把Web服务的基础功能和高级功能串讲了一遍。当然，上面只介绍了Web服务的核心功能，还有很多其他的功能，你可以通过学习<a href="https://github.com/gin-gonic/gin" target="_blank">Gin的官方文档</a>来了解。</p>

<p>你可以看到，Web服务有很多核心功能，这些功能我们可以基于net/http包自己封装。但在实际的项目开发中， 我们更多会选择使用基于net/http包进行封装的优秀开源Web框架。本实战项目选择了Gin框架。</p>

<p>接下来，我们主要看下Gin框架是如何实现以上核心功能的，这些功能我们在实际的开发中可以直接拿来使用。</p>

<h2 id="为什么选择gin框架">为什么选择Gin框架？</h2>

<p>优秀的Web框架有很多，我们为什么要选择Gin呢？在回答这个问题之前，我们先来看下选择Web框架时的关注点。</p>

<p>在选择Web框架时，我们可以关注如下几点：</p>

<ul>
<li>路由功能；</li>
<li>是否具备middleware/filter能力；</li>
<li>HTTP 参数（path、query、form、header、body）解析和返回；</li>
<li>性能和稳定性；</li>
<li>使用复杂度；</li>
<li>社区活跃度。</li>
</ul>

<p>按 GitHub Star 数来排名，当前比较火的 Go Web 框架有 Gin、Beego、Echo、Revel 、Martini。经过调研，我从中选择了Gin框架，原因是Gin具有如下特性：</p>

<ul>
<li>轻量级，代码质量高，性能比较高；</li>
<li>项目目前很活跃，并有很多可用的 Middleware；</li>
<li>作为一个 Web 框架，功能齐全，使用起来简单。</li>
</ul>

<p>那接下来，我就先详细介绍下Gin框架。</p>

<p><a href="https://github.com/gin-gonic/gin" target="_blank">Gin</a>是用Go语言编写的Web框架，功能完善，使用简单，性能很高。Gin核心的路由功能是通过一个定制版的<a href="https://github.com/julienschmidt/httprouter" target="_blank">HttpRouter</a>来实现的，具有很高的路由性能。</p>

<p>Gin有很多功能，这里我给你列出了它的一些核心功能：</p>

<ul>
<li>支持HTTP方法：GET、POST、PUT、PATCH、DELETE、OPTIONS。</li>
<li>支持不同位置的HTTP参数：路径参数（path）、查询字符串参数（query）、表单参数（form）、HTTP头参数（header）、消息体参数（body）。</li>
<li>支持HTTP路由和路由分组。</li>
<li>支持middleware和自定义middleware。</li>
<li>支持自定义Log。</li>
<li>支持binding和validation，支持自定义validator。可以bind如下参数：query、path、body、header、form。</li>
<li>支持重定向。</li>
<li>支持basic auth middleware。</li>
<li>支持自定义HTTP配置。</li>
<li>支持优雅关闭。</li>
<li>支持HTTP2。</li>
<li>支持设置和获取cookie。</li>
</ul>

<h2 id="gin是如何支持web服务基础功能的">Gin是如何支持Web服务基础功能的？</h2>

<p>接下来，我们先通过一个具体的例子，看下Gin是如何支持Web服务基础功能的，后面再详细介绍这些功能的用法。</p>

<p>我们创建一个webfeature目录，用来存放示例代码。因为要演示HTTPS的用法，所以需要创建证书文件。具体可以分为两步。</p>

<p>第一步，执行以下命令创建证书：</p>

<pre><code>cat &lt;&lt; 'EOF' &gt; ca.pem
-----BEGIN CERTIFICATE-----
MIICSjCCAbOgAwIBAgIJAJHGGR4dGioHMA0GCSqGSIb3DQEBCwUAMFYxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxDzANBgNVBAMTBnRlc3RjYTAeFw0xNDExMTEyMjMxMjla
Fw0yNDExMDgyMjMxMjlaMFYxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0
YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxDzANBgNVBAMT
BnRlc3RjYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAwEDfBV5MYdlHVHJ7
+L4nxrZy7mBfAVXpOc5vMYztssUI7mL2/iYujiIXM+weZYNTEpLdjyJdu7R5gGUu
g1jSVK/EPHfc74O7AyZU34PNIP4Sh33N+/A5YexrNgJlPY+E3GdVYi4ldWJjgkAd
Qah2PH5ACLrIIC6tRka9hcaBlIECAwEAAaMgMB4wDAYDVR0TBAUwAwEB/zAOBgNV
HQ8BAf8EBAMCAgQwDQYJKoZIhvcNAQELBQADgYEAHzC7jdYlzAVmddi/gdAeKPau
sPBG/C2HCWqHzpCUHcKuvMzDVkY/MP2o6JIW2DBbY64bO/FceExhjcykgaYtCH/m
oIU63+CFOTtR7otyQAWHqXa7q4SbCDlG7DyRFxqG0txPtGvy12lgldA2+RgcigQG
Dfcog5wrJytaQ6UA0wE=
-----END CERTIFICATE-----
EOF

cat &lt;&lt; 'EOF' &gt; server.key
-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAOHDFScoLCVJpYDD
M4HYtIdV6Ake/sMNaaKdODjDMsux/4tDydlumN+fm+AjPEK5GHhGn1BgzkWF+slf
3BxhrA/8dNsnunstVA7ZBgA/5qQxMfGAq4wHNVX77fBZOgp9VlSMVfyd9N8YwbBY
AckOeUQadTi2X1S6OgJXgQ0m3MWhAgMBAAECgYAn7qGnM2vbjJNBm0VZCkOkTIWm
V10okw7EPJrdL2mkre9NasghNXbE1y5zDshx5Nt3KsazKOxTT8d0Jwh/3KbaN+YY
tTCbKGW0pXDRBhwUHRcuRzScjli8Rih5UOCiZkhefUTcRb6xIhZJuQy71tjaSy0p
dHZRmYyBYO2YEQ8xoQJBAPrJPhMBkzmEYFtyIEqAxQ/o/A6E+E4w8i+KM7nQCK7q
K4JXzyXVAjLfyBZWHGM2uro/fjqPggGD6QH1qXCkI4MCQQDmdKeb2TrKRh5BY1LR
81aJGKcJ2XbcDu6wMZK4oqWbTX2KiYn9GB0woM6nSr/Y6iy1u145YzYxEV/iMwff
DJULAkB8B2MnyzOg0pNFJqBJuH29bKCcHa8gHJzqXhNO5lAlEbMK95p/P2Wi+4Hd
aiEIAF1BF326QJcvYKmwSmrORp85AkAlSNxRJ50OWrfMZnBgzVjDx3xG6KsFQVk2
ol6VhqL6dFgKUORFUWBvnKSyhjJxurlPEahV6oo6+A+mPhFY8eUvAkAZQyTdupP3
XEFQKctGz+9+gKkemDp7LBBMEMBXrGTLPhpEfcjv/7KPdnFHYmhYeBTBnuVmTVWe
F98XJ7tIFfJq
-----END PRIVATE KEY-----
EOF

cat &lt;&lt; 'EOF' &gt; server.pem
-----BEGIN CERTIFICATE-----
MIICnDCCAgWgAwIBAgIBBzANBgkqhkiG9w0BAQsFADBWMQswCQYDVQQGEwJBVTET
MBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50ZXJuZXQgV2lkZ2l0cyBQ
dHkgTHRkMQ8wDQYDVQQDEwZ0ZXN0Y2EwHhcNMTUxMTA0MDIyMDI0WhcNMjUxMTAx
MDIyMDI0WjBlMQswCQYDVQQGEwJVUzERMA8GA1UECBMISWxsaW5vaXMxEDAOBgNV
BAcTB0NoaWNhZ28xFTATBgNVBAoTDEV4YW1wbGUsIENvLjEaMBgGA1UEAxQRKi50
ZXN0Lmdvb2dsZS5jb20wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAOHDFSco
LCVJpYDDM4HYtIdV6Ake/sMNaaKdODjDMsux/4tDydlumN+fm+AjPEK5GHhGn1Bg
zkWF+slf3BxhrA/8dNsnunstVA7ZBgA/5qQxMfGAq4wHNVX77fBZOgp9VlSMVfyd
9N8YwbBYAckOeUQadTi2X1S6OgJXgQ0m3MWhAgMBAAGjazBpMAkGA1UdEwQCMAAw
CwYDVR0PBAQDAgXgME8GA1UdEQRIMEaCECoudGVzdC5nb29nbGUuZnKCGHdhdGVy
em9vaS50ZXN0Lmdvb2dsZS5iZYISKi50ZXN0LnlvdXR1YmUuY29thwTAqAEDMA0G
CSqGSIb3DQEBCwUAA4GBAJFXVifQNub1LUP4JlnX5lXNlo8FxZ2a12AFQs+bzoJ6
hM044EDjqyxUqSbVePK0ni3w1fHQB5rY9yYC5f8G7aqqTY1QOhoUk8ZTSTRpnkTh
y4jjdvTZeLDVBlueZUTDRmy2feY5aZIU18vFDK08dTG0A87pppuv1LNIR3loveU8
-----END CERTIFICATE-----
EOF
</code></pre>

<p>第二步，创建main.go文件：</p>

<pre><code>package main

import (
	&quot;fmt&quot;
	&quot;log&quot;
	&quot;net/http&quot;
	&quot;sync&quot;
	&quot;time&quot;

	&quot;github.com/gin-gonic/gin&quot;
	&quot;golang.org/x/sync/errgroup&quot;
)

type Product struct {
	Username    string    `json:&quot;username&quot; binding:&quot;required&quot;`
	Name        string    `json:&quot;name&quot; binding:&quot;required&quot;`
	Category    string    `json:&quot;category&quot; binding:&quot;required&quot;`
	Price       int       `json:&quot;price&quot; binding:&quot;gte=0&quot;`
	Description string    `json:&quot;description&quot;`
	CreatedAt   time.Time `json:&quot;createdAt&quot;`
}

type productHandler struct {
	sync.RWMutex
	products map[string]Product
}

func newProductHandler() *productHandler {
	return &amp;productHandler{
		products: make(map[string]Product),
	}
}

func (u *productHandler) Create(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	// 1. 参数解析
	var product Product
	if err := c.ShouldBindJSON(&amp;product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{&quot;error&quot;: err.Error()})
		return
	}

	// 2. 参数校验
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{&quot;error&quot;: fmt.Sprintf(&quot;product %s already exist&quot;, product.Name)})
		return
	}
	product.CreatedAt = time.Now()

	// 3. 逻辑处理
	u.products[product.Name] = product
	log.Printf(&quot;Register product %s success&quot;, product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}

func (u *productHandler) Get(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	product, ok := u.products[c.Param(&quot;name&quot;)]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{&quot;error&quot;: fmt.Errorf(&quot;can not found product %s&quot;, c.Param(&quot;name&quot;))})
		return
	}

	c.JSON(http.StatusOK, product)
}

func router() http.Handler {
	router := gin.Default()
	productHandler := newProductHandler()
	// 路由分组、中间件、认证
	v1 := router.Group(&quot;/v1&quot;)
	{
		productv1 := v1.Group(&quot;/products&quot;)
		{
			// 路由匹配
			productv1.POST(&quot;&quot;, productHandler.Create)
			productv1.GET(&quot;:name&quot;, productHandler.Get)
		}
	}

	return router
}

func main() {
	var eg errgroup.Group

	// 一进程多端口
	insecureServer := &amp;http.Server{
		Addr:         &quot;:8080&quot;,
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	secureServer := &amp;http.Server{
		Addr:         &quot;:8443&quot;,
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil &amp;&amp; err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS(&quot;server.pem&quot;, &quot;server.key&quot;)
		if err != nil &amp;&amp; err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
</code></pre>

<p>运行以上代码：</p>

<pre><code>$ go run main.go
</code></pre>

<p>打开另外一个终端，请求HTTP接口：</p>

<pre><code># 创建产品
$ curl -XPOST -H&quot;Content-Type: application/json&quot; -d'{&quot;username&quot;:&quot;colin&quot;,&quot;name&quot;:&quot;iphone12&quot;,&quot;category&quot;:&quot;phone&quot;,&quot;price&quot;:8000,&quot;description&quot;:&quot;cannot afford&quot;}' http://127.0.0.1:8080/v1/products
{&quot;username&quot;:&quot;colin&quot;,&quot;name&quot;:&quot;iphone12&quot;,&quot;category&quot;:&quot;phone&quot;,&quot;price&quot;:8000,&quot;description&quot;:&quot;cannot afford&quot;,&quot;createdAt&quot;:&quot;2021-06-20T11:17:03.818065988+08:00&quot;}

# 获取产品信息
$ curl -XGET http://127.0.0.1:8080/v1/products/iphone12
{&quot;username&quot;:&quot;colin&quot;,&quot;name&quot;:&quot;iphone12&quot;,&quot;category&quot;:&quot;phone&quot;,&quot;price&quot;:8000,&quot;description&quot;:&quot;cannot afford&quot;,&quot;createdAt&quot;:&quot;2021-06-20T11:17:03.818065988+08:00&quot;}
</code></pre>

<p>示例代码存放地址为<a href="https://github.com/marmotedu/gopractise-demo/tree/master/gin/webfeature" target="_blank">webfeature</a>。</p>

<p>另外，Gin项目仓库中也包含了很多使用示例，如果你想详细了解，可以参考 <a href="https://github.com/gin-gonic/examples" target="_blank">gin examples</a>。</p>

<p>下面，我来详细介绍下Gin是如何支持Web服务基础功能的。</p>

<h3 id="http-https支持">HTTP/HTTPS支持</h3>

<p>因为Gin是基于net/http包封装的一个Web框架，所以它天然就支持HTTP/HTTPS。在上述代码中，通过以下方式开启一个HTTP服务：</p>

<pre><code>insecureServer := &amp;http.Server{
	Addr:         &quot;:8080&quot;,
	Handler:      router(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}
...
err := insecureServer.ListenAndServe()
</code></pre>

<p>通过以下方式开启一个HTTPS服务：</p>

<pre><code>secureServer := &amp;http.Server{
	Addr:         &quot;:8443&quot;,
	Handler:      router(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}
...
err := secureServer.ListenAndServeTLS(&quot;server.pem&quot;, &quot;server.key&quot;)
</code></pre>

<h3 id="json数据格式支持">JSON数据格式支持</h3>

<p>Gin支持多种数据通信格式，例如application/json、application/xml。可以通过<code>c.ShouldBindJSON</code>函数，将Body中的JSON格式数据解析到指定的Struct中，通过<code>c.JSON</code>函数返回JSON格式的数据。</p>

<h3 id="路由匹配">路由匹配</h3>

<p>Gin支持两种路由匹配规则。</p>

<p><strong>第一种匹配规则是精确匹配。</strong>例如，路由为/products/:name，匹配情况如下表所示：</p>

<p><img src="assets/f1fe6c4bff9c413d87b2a7561f46f7bc.jpg" alt="" /></p>

<p><strong>第二种匹配规则是模糊匹配。</strong>例如，路由为/products/*name，匹配情况如下表所示：</p>

<p><img src="assets/b0e30669260b4de3a94c99c25d643789.jpg" alt="" /></p>

<h3 id="路由分组">路由分组</h3>

<p>Gin通过Group函数实现了路由分组的功能。路由分组是一个非常常用的功能，可以将相同版本的路由分为一组，也可以将相同RESTful资源的路由分为一组。例如：</p>

<pre><code>v1 := router.Group(&quot;/v1&quot;, gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;}))
{
    productv1 := v1.Group(&quot;/products&quot;)
    {
        // 路由匹配
        productv1.POST(&quot;&quot;, productHandler.Create)
        productv1.GET(&quot;:name&quot;, productHandler.Get)
    }

    orderv1 := v1.Group(&quot;/orders&quot;)
    {
        // 路由匹配
        orderv1.POST(&quot;&quot;, orderHandler.Create)
        orderv1.GET(&quot;:name&quot;, orderHandler.Get)
    }
}

v2 := router.Group(&quot;/v2&quot;, gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;}))
{
    productv2 := v2.Group(&quot;/products&quot;)
    {
        // 路由匹配
        productv2.POST(&quot;&quot;, productHandler.Create)
        productv2.GET(&quot;:name&quot;, productHandler.Get)
    }
}
</code></pre>

<p>通过将路由分组，可以对相同分组的路由做统一处理。比如上面那个例子，我们可以通过代码</p>

<pre><code>v1 := router.Group(&quot;/v1&quot;, gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;}))
</code></pre>

<p>给所有属于v1分组的路由都添加gin.BasicAuth中间件，以实现认证功能。中间件和认证，这里你先不用深究，下面讲高级功能的时候会介绍到。</p>

<h3 id="一进程多服务">一进程多服务</h3>

<p>我们可以通过以下方式实现一进程多服务：</p>

<pre><code>var eg errgroup.Group
insecureServer := &amp;http.Server{...}
secureServer := &amp;http.Server{...}

eg.Go(func() error {
	err := insecureServer.ListenAndServe()
	if err != nil &amp;&amp; err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return err
})
eg.Go(func() error {
	err := secureServer.ListenAndServeTLS(&quot;server.pem&quot;, &quot;server.key&quot;)
	if err != nil &amp;&amp; err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return err
}

if err := eg.Wait(); err != nil {
	log.Fatal(err)
})
</code></pre>

<p>上述代码实现了两个相同的服务，分别监听在不同的端口。这里需要注意的是，为了不阻塞启动第二个服务，我们需要把ListenAndServe函数放在goroutine中执行，并且调用eg.Wait()来阻塞程序进程，从而让两个HTTP服务在goroutine中持续监听端口，并提供服务。</p>

<h3 id="参数解析-参数校验-逻辑处理-返回结果">参数解析、参数校验、逻辑处理、返回结果</h3>

<p>此外，Web服务还应该具有参数解析、参数校验、逻辑处理、返回结果4类功能，因为这些功能联系紧密，我们放在一起来说。</p>

<p>在productHandler的Create方法中，我们通过<code>c.ShouldBindJSON</code>来解析参数，接下来自己编写校验代码，然后将product信息保存在内存中（也就是业务逻辑处理），最后通过<code>c.JSON</code>返回创建的product信息。代码如下：</p>

<pre><code>func (u *productHandler) Create(c *gin.Context) {
	u.Lock()
	defer u.Unlock()

	// 1. 参数解析
	var product Product
	if err := c.ShouldBindJSON(&amp;product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{&quot;error&quot;: err.Error()})
		return
	}

	// 2. 参数校验
	if _, ok := u.products[product.Name]; ok {
		c.JSON(http.StatusBadRequest, gin.H{&quot;error&quot;: fmt.Sprintf(&quot;product %s already exist&quot;, product.Name)})
		return
	}
	product.CreatedAt = time.Now()

	// 3. 逻辑处理
	u.products[product.Name] = product
	log.Printf(&quot;Register product %s success&quot;, product.Name)

	// 4. 返回结果
	c.JSON(http.StatusOK, product)
}
</code></pre>

<p>那这个时候，你可能会问：HTTP的请求参数可以存在不同的位置，Gin是如何解析的呢？这里，我们先来看下HTTP有哪些参数类型。HTTP具有以下5种参数类型：</p>

<ul>
<li>路径参数（path）。例如<code>gin.Default().GET(&quot;/user/:name&quot;, nil)</code>， name就是路径参数。</li>
<li>查询字符串参数（query）。例如<code>/welcome?firstname=Lingfei&amp;lastname=Kong</code>，firstname和lastname就是查询字符串参数。</li>
<li>表单参数（form）。例如<code>curl -X POST -F 'username=colin' -F 'password=colin1234' http://mydomain.com/login</code>，username和password就是表单参数。</li>
<li>HTTP头参数（header）。例如<code>curl -X POST -H 'Content-Type: application/json' -d '{&quot;username&quot;:&quot;colin&quot;,&quot;password&quot;:&quot;colin1234&quot;}' http://mydomain.com/login</code>，Content-Type就是HTTP头参数。</li>
<li>消息体参数（body）。例如<code>curl -X POST -H 'Content-Type: application/json' -d '{&quot;username&quot;:&quot;colin&quot;,&quot;password&quot;:&quot;colin1234&quot;}' http://mydomain.com/login</code>，username和password就是消息体参数。</li>
</ul>

<p>Gin提供了一些函数，来分别读取这些HTTP参数，每种类别会提供两种函数，一种函数可以直接读取某个参数的值，另外一种函数会把同类HTTP参数绑定到一个Go结构体中。比如，有如下路径参数：</p>

<pre><code>gin.Default().GET(&quot;/:name/:id&quot;, nil)
</code></pre>

<p>我们可以直接读取每个参数：</p>

<pre><code>name := c.Param(&quot;name&quot;)
action := c.Param(&quot;action&quot;)
</code></pre>

<p>也可以将所有的路径参数，绑定到结构体中：</p>

<pre><code>type Person struct {
    ID string `uri:&quot;id&quot; binding:&quot;required,uuid&quot;`
    Name string `uri:&quot;name&quot; binding:&quot;required&quot;`
}

if err := c.ShouldBindUri(&amp;person); err != nil {
    // normal code
    return
}
</code></pre>

<p>Gin在绑定参数时，是通过结构体的tag来判断要绑定哪类参数到结构体中的。这里要注意，不同的HTTP参数有不同的结构体tag。</p>

<ul>
<li>路径参数：uri。</li>
<li>查询字符串参数：form。</li>
<li>表单参数：form。</li>
<li>HTTP头参数：header。</li>
<li>消息体参数：会根据Content-Type，自动选择使用json或者xml，也可以调用ShouldBindJSON或者ShouldBindXML直接指定使用哪个tag。</li>
</ul>

<p>针对每种参数类型，Gin都有对应的函数来获取和绑定这些参数。这些函数都是基于如下两个函数进行封装的：</p>

<ol>
<li>ShouldBindWith(obj interface{}, b binding.Binding) error</li>
</ol>

<p>非常重要的一个函数，很多ShouldBindXXX函数底层都是调用ShouldBindWith函数来完成参数绑定的。该函数会根据传入的绑定引擎，将参数绑定到传入的结构体指针中，<strong>如果绑定失败，只返回错误内容，但不终止HTTP请求。</strong>ShouldBindWith支持多种绑定引擎，例如 binding.JSON、binding.Query、binding.Uri、binding.Header等，更详细的信息你可以参考 <a href="https://github.com/gin-gonic/gin/blob/v1.7.2/binding/binding.go#L72" target="_blank">binding.go</a>。</p>

<ol>
<li>MustBindWith(obj interface{}, b binding.Binding) error</li>
</ol>

<p>这是另一个非常重要的函数，很多BindXXX函数底层都是调用MustBindWith函数来完成参数绑定的。该函数会根据传入的绑定引擎，将参数绑定到传入的结构体指针中，<strong>如果绑定失败，返回错误并终止请求，返回HTTP 400错误。</strong>MustBindWith所支持的绑定引擎跟ShouldBindWith函数一样。</p>

<p>Gin基于ShouldBindWith和MustBindWith这两个函数，又衍生出很多新的Bind函数。这些函数可以满足不同场景下获取HTTP参数的需求。Gin提供的函数可以获取5个类别的HTTP参数。</p>

<ul>
<li>路径参数：ShouldBindUri、BindUri；</li>
<li>查询字符串参数：ShouldBindQuery、BindQuery；</li>
<li>表单参数：ShouldBind；</li>
<li>HTTP头参数：ShouldBindHeader、BindHeader；</li>
<li>消息体参数：ShouldBindJSON、BindJSON等。</li>
</ul>

<p>每个类别的Bind函数，详细信息你可以参考<a href="https://github.com/marmotedu/geekbang-go/blob/master/Gin%E6%8F%90%E4%BE%9B%E7%9A%84Bind%E5%87%BD%E6%95%B0.md" target="_blank">Gin提供的Bind函数</a>。</p>

<p>这里要注意，Gin并没有提供类似ShouldBindForm、BindForm这类函数来绑定表单参数，但我们可以通过ShouldBind来绑定表单参数。当HTTP方法为GET时，ShouldBind只绑定Query类型的参数；当HTTP方法为POST时，会先检查content-type是否是json或者xml，如果不是，则绑定Form类型的参数。</p>

<p>所以，ShouldBind可以绑定Form类型的参数，但前提是HTTP方法是POST，并且content-type不是application/json、application/xml。</p>

<p>在Go项目开发中，我建议使用ShouldBindXXX，这样可以确保我们设置的HTTP Chain（Chain可以理解为一个HTTP请求的一系列处理插件）能够继续被执行。</p>

<h2 id="gin是如何支持web服务高级功能的">Gin是如何支持Web服务高级功能的？</h2>

<p>上面介绍了Web服务的基础功能，这里我再来介绍下高级功能。Web服务可以具备多个高级功能，但比较核心的高级功能是中间件、认证、RequestID、跨域和优雅关停。</p>

<h3 id="中间件">中间件</h3>

<p>Gin支持中间件，HTTP请求在转发到实际的处理函数之前，会被一系列加载的中间件进行处理。在中间件中，可以解析HTTP请求做一些逻辑处理，例如：跨域处理或者生成X-Request-ID并保存在context中，以便追踪某个请求。处理完之后，可以选择中断并返回这次请求，也可以选择将请求继续转交给下一个中间件处理。当所有的中间件都处理完之后，请求才会转给路由函数进行处理。具体流程如下图：</p>

<p><img src="assets/00cdcb5bb15e455f96c49b74e0ac6fb9.jpg" alt="" /></p>

<p>通过中间件，可以实现对所有请求都做统一的处理，提高开发效率，并使我们的代码更简洁。但是，因为所有的请求都需要经过中间件的处理，可能会增加请求延时。对于中间件特性，我有如下建议：</p>

<ul>
<li>中间件做成可加载的，通过配置文件指定程序启动时加载哪些中间件。</li>
<li>只将一些通用的、必要的功能做成中间件。</li>
<li>在编写中间件时，一定要保证中间件的代码质量和性能。</li>
</ul>

<p>在Gin中，可以通过gin.Engine的Use方法来加载中间件。中间件可以加载到不同的位置上，而且不同的位置作用范围也不同，例如：</p>

<pre><code>router := gin.New()
router.Use(gin.Logger(), gin.Recovery()) // 中间件作用于所有的HTTP请求
v1 := router.Group(&quot;/v1&quot;).Use(gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;})) // 中间件作用于v1 group
v1.POST(&quot;/login&quot;, Login).Use(gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;})) //中间件只作用于/v1/login API接口
</code></pre>

<p>Gin框架本身支持了一些中间件。</p>

<ul>
<li><strong>gin.Logger()：</strong>Logger中间件会将日志写到gin.DefaultWriter，gin.DefaultWriter默认为 os.Stdout。</li>
<li><strong>gin.Recovery()：</strong>Recovery中间件可以从任何panic恢复，并且写入一个500状态码。</li>
<li><strong>gin.CustomRecovery(handle gin.RecoveryFunc)：</strong>类似Recovery中间件，但是在恢复时还会调用传入的handle方法进行处理。</li>
<li><strong>gin.BasicAuth()：</strong>HTTP请求基本认证（使用用户名和密码进行认证）。</li>
</ul>

<p>另外，Gin还支持自定义中间件。中间件其实是一个函数，函数类型为gin.HandlerFunc，HandlerFunc底层类型为func(*Context)。如下是一个Logger中间件的实现：</p>

<pre><code>package main

import (
	&quot;log&quot;
	&quot;time&quot;

	&quot;github.com/gin-gonic/gin&quot;
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置变量example
		c.Set(&quot;example&quot;, &quot;12345&quot;)

		// 请求之前

		c.Next()

		// 请求之后
		latency := time.Since(t)
		log.Print(latency)

		// 访问我们发送的状态
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET(&quot;/test&quot;, func(c *gin.Context) {
		example := c.MustGet(&quot;example&quot;).(string)

		// it would print: &quot;12345&quot;
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(&quot;:8080&quot;)
}
</code></pre>

<p>另外，还有很多开源的中间件可供我们选择，我把一些常用的总结在了表格里：</p>

<p><img src="assets/39e1403ebaee455a91ef34f39f887583.jpg" alt="" /></p>

<h3 id="认证-requestid-跨域">认证、RequestID、跨域</h3>

<p>认证、RequestID、跨域这三个高级功能，都可以通过Gin的中间件来实现，例如：</p>

<pre><code>router := gin.New()

// 认证
router.Use(gin.BasicAuth(gin.Accounts{&quot;foo&quot;: &quot;bar&quot;, &quot;colin&quot;: &quot;colin404&quot;}))

// RequestID
router.Use(requestid.New(requestid.Config{
    Generator: func() string {
        return &quot;test&quot;
    },
}))

// 跨域
// CORS for https://foo.com and https://github.com origins, allowing:
// - PUT and PATCH methods
// - Origin header
// - Credentials share
// - Preflight requests cached for 12 hours
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{&quot;https://foo.com&quot;},
    AllowMethods:     []string{&quot;PUT&quot;, &quot;PATCH&quot;},
    AllowHeaders:     []string{&quot;Origin&quot;},
    ExposeHeaders:    []string{&quot;Content-Length&quot;},
    AllowCredentials: true,
    AllowOriginFunc: func(origin string) bool {
        return origin == &quot;https://github.com&quot;
    },
    MaxAge: 12 * time.Hour,
}))
</code></pre>

<h3 id="优雅关停">优雅关停</h3>

<p>Go项目上线后，我们还需要不断迭代来丰富项目功能、修复Bug等，这也就意味着，我们要不断地重启Go服务。对于HTTP服务来说，如果访问量大，重启服务的时候可能还有很多连接没有断开，请求没有完成。如果这时候直接关闭服务，这些连接会直接断掉，请求异常终止，这就会对用户体验和产品口碑造成很大影响。因此，这种关闭方式不是一种优雅的关闭方式。</p>

<p>这时候，我们期望HTTP服务可以在处理完所有请求后，正常地关闭这些连接，也就是优雅地关闭服务。我们有两种方法来优雅关闭HTTP服务，分别是借助第三方的Go包和自己编码实现。</p>

<p>方法一：借助第三方的Go包</p>

<p>如果使用第三方的Go包来实现优雅关闭，目前用得比较多的包是<a href="https://github.com/fvbock/endless" target="_blank">fvbock/endless</a>。我们可以使用fvbock/endless来替换掉net/http的ListenAndServe方法，例如：</p>

<pre><code>router := gin.Default()
router.GET(&quot;/&quot;, handler)
// [...]
endless.ListenAndServe(&quot;:4242&quot;, router)
</code></pre>

<p>方法二：编码实现</p>

<p>借助第三方包的好处是可以稍微减少一些编码工作量，但缺点是引入了一个新的依赖包，因此我更倾向于自己编码实现。Go 1.8版本或者更新的版本，http.Server内置的Shutdown方法，已经实现了优雅关闭。下面是一个示例：</p>

<pre><code>// +build go1.8

package main

import (
	&quot;context&quot;
	&quot;log&quot;
	&quot;net/http&quot;
	&quot;os&quot;
	&quot;os/signal&quot;
	&quot;syscall&quot;
	&quot;time&quot;

	&quot;github.com/gin-gonic/gin&quot;
)

func main() {
	router := gin.Default()
	router.GET(&quot;/&quot;, func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, &quot;Welcome Gin Server&quot;)
	})

	srv := &amp;http.Server{
		Addr:    &quot;:8080&quot;,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil &amp;&amp; err != http.ErrServerClosed {
			log.Fatalf(&quot;listen: %s\n&quot;, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	&lt;-quit
	log.Println(&quot;Shutting down server...&quot;)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(&quot;Server forced to shutdown:&quot;, err)
	}

	log.Println(&quot;Server exiting&quot;)
}
</code></pre>

<p>上面的示例中，需要把srv.ListenAndServe放在goroutine中执行，这样才不会阻塞到srv.Shutdown函数。因为我们把srv.ListenAndServe放在了goroutine中，所以需要一种可以让整个进程常驻的机制。</p>

<p>这里，我们借助了有缓冲channel，并且调用signal.Notify函数将该channel绑定到SIGINT、SIGTERM信号上。这样，收到SIGINT、SIGTERM信号后，quilt通道会被写入值，从而结束阻塞状态，程序继续运行，执行srv.Shutdown(ctx)，优雅关停HTTP服务。</p>

<h2 id="总结">总结</h2>

<p>今天我们主要学习了Web服务的核心功能，以及如何开发这些功能。在实际的项目开发中， 我们一般会使用基于net/http包进行封装的优秀开源Web框架。</p>

<p>当前比较火的Go Web框架有 Gin、Beego、Echo、Revel、Martini。你可以根据需要进行选择。我比较推荐Gin，Gin也是目前比较受欢迎的Web框架。Gin Web框架支持Web服务的很多基础功能，例如 HTTP/HTTPS、JSON格式的数据、路由分组和匹配、一进程多服务等。</p>

<p>另外，Gin还支持Web服务的一些高级功能，例如 中间件、认证、RequestID、跨域和优雅关停等。</p>

<h2 id="课后练习">课后练习</h2>

<ol>
<li>使用 Gin 框架编写一个简单的Web服务，要求该Web服务可以解析参数、校验参数，并进行一些简单的业务逻辑处理，最终返回处理结果。欢迎在留言区分享你的成果，或者遇到的问题。</li>
<li>思考下，如何给iam-apiserver的/healthz接口添加一个限流中间件，用来限制请求/healthz的频率。</li>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#93ffffffaaa7a2a2a3a4d3f4fef2faffbdf0fcfe" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0df11a4aae653b',t:'MTczNDAwODg5Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>