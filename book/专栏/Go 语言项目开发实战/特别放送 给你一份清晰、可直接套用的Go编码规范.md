<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=特别放送&#32;给你一份清晰、可直接套用的Go编码规范>
        <link rel="icon" href="/static/favicon.png">
        <title>特别放送 给你一份清晰、可直接套用的Go编码规范 </title>
        
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
                            <h1 id="title" data-id="特别放送 给你一份清晰、可直接套用的Go编码规范" class="title">特别放送 给你一份清晰、可直接套用的Go编码规范</h1>
                            <div><p>你好，我是孔令飞。</p>

<p>我们在上一讲学习了“写出优雅Go项目的方法论”，那一讲内容很丰富，是我多年Go项目开发的经验沉淀，需要你多花一些时间好好消化吸收。吃完大餐之后，咱们今天来一期特别放送，就是上一讲我提到过的编码规范。这一讲里，为了帮你节省时间和精力，我会给你一份清晰、可直接套用的 Go 编码规范，帮助你编写一个高质量的 Go 应用。</p>

<p>这份规范，是我参考了Go官方提供的编码规范，以及Go社区沉淀的一些比较合理的规范之后，加入自己的理解总结出的，它比很多公司内部的规范更全面，你掌握了，以后在面试大厂的时候，或者在大厂里写代码的时候，都会让人高看你一眼，觉得你code很专业。</p>

<p>这份编码规范中包含代码风格、命名规范、注释规范、类型、控制结构、函数、GOPATH 设置规范、依赖管理和最佳实践九类规范。如果你觉得这些规范内容太多了，看完一遍也记不住，这完全没关系。你可以多看几遍，也可以在用到时把它翻出来，在实际应用中掌握。这篇特别放送的内容，更多是作为写代码时候的一个参考手册。</p>

<h2 id="1-代码风格">1. 代码风格</h2>

<h3 id="1-1-代码格式">1.1 代码格式</h3>

<ul>
<li><p>代码都必须用 <code>gofmt</code> 进行格式化。</p></li>

<li><p>运算符和操作数之间要留空格。</p></li>

<li><p>建议一行代码不超过120个字符，超过部分，请采用合适的换行方式换行。但也有些例外场景，例如import行、工具自动生成的代码、带tag的struct字段。</p></li>

<li><p>文件长度不能超过800行。</p></li>

<li><p>函数长度不能超过80行。</p></li>

<li><p>import规范</p>

<ul>
<li>代码都必须用 goimports进行格式化（建议将代码Go代码编辑器设置为：保存时运行 goimports）。-
- 不要使用相对路径引入包，例如 import …/util/net 。-
- 包名称与导入路径的最后一个目录名不匹配时，或者多个相同包名冲突时，则必须使用导入别名。</li>
</ul></li>
</ul>

<pre><code>// bad
	&quot;github.com/dgrijalva/jwt-go/v4&quot;

	//good
	jwt &quot;github.com/dgrijalva/jwt-go/v4&quot;
</code></pre>

<ul>
<li>*   导入的包建议进行分组，匿名包的引用使用一个新的分组，并对匿名包引用进行说明。</li>
</ul>

<pre><code>	import (
		// go 标准包
		&quot;fmt&quot;

		// 第三方包
	    &quot;github.com/jinzhu/gorm&quot;
	    &quot;github.com/spf13/cobra&quot;
	    &quot;github.com/spf13/viper&quot;

		// 匿名包单独分组，并对匿名包引用进行说明
	    // import mysql driver
	    _ &quot;github.com/jinzhu/gorm/dialects/mysql&quot;

		// 内部包
	    v1 &quot;github.com/marmotedu/api/apiserver/v1&quot;
	    metav1 &quot;github.com/marmotedu/apimachinery/pkg/meta/v1&quot;
	    &quot;github.com/marmotedu/iam/pkg/cli/genericclioptions&quot;
	)
</code></pre>

<h3 id="1-2-声明-初始化和定义">1.2 声明、初始化和定义</h3>

<ul>
<li>当函数中需要使用到多个变量时，可以在函数开始处使用var声明。在函数外部声明必须使用 <code>var</code> ，不要采用 <code>:=</code> ，容易踩到变量的作用域的问题。</li>
</ul>

<pre><code>var (
	Width  int
	Height int
)
</code></pre>

<ul>
<li>在初始化结构引用时，请使用&amp;T{}代替new(T)，以使其与结构体初始化一致。</li>
</ul>

<pre><code>// bad
sptr := new(T)
sptr.Name = &quot;bar&quot;

// good
sptr := &amp;T{Name: &quot;bar&quot;}
</code></pre>

<ul>
<li>struct 声明和初始化格式采用多行，定义如下。</li>
</ul>

<pre><code>type User struct{
    Username  string
    Email     string
}

user := User{
	Username: &quot;colin&quot;,
	Email: &quot;<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="086b676461663c383c486e677065696164266b6765">[email&#160;protected]</a>&quot;,
}
</code></pre>

<ul>
<li>相似的声明放在一组，同样适用于常量、变量和类型声明。</li>
</ul>

<pre><code>// bad
import &quot;a&quot;
import &quot;b&quot;

// good
import (
  &quot;a&quot;
  &quot;b&quot;
)
</code></pre>

<ul>
<li>尽可能指定容器容量，以便为容器预先分配内存，例如：</li>
</ul>

<pre><code>v := make(map[int]string, 4)
v := make([]string, 0, 4)
</code></pre>

<ul>
<li>在顶层，使用标准var关键字。请勿指定类型，除非它与表达式的类型不同。</li>
</ul>

<pre><code>// bad
var _s string = F()

func F() string { return &quot;A&quot; }

// good
var _s = F()
// 由于 F 已经明确了返回一个字符串类型，因此我们没有必要显式指定_s 的类型
// 还是那种类型

func F() string { return &quot;A&quot; }
</code></pre>

<ul>
<li>对于未导出的顶层常量和变量，使用_作为前缀。</li>
</ul>

<pre><code>// bad
const (
  defaultHost = &quot;127.0.0.1&quot;
  defaultPort = 8080
)

// good
const (
  _defaultHost = &quot;127.0.0.1&quot;
  _defaultPort = 8080
)
</code></pre>

<ul>
<li>嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。</li>
</ul>

<pre><code>// bad
type Client struct {
  version int
  http.Client
}

// good
type Client struct {
  http.Client

  version int
}
</code></pre>

<h3 id="1-3-错误处理">1.3 错误处理</h3>

<ul>
<li><code>error</code>作为函数的值返回，必须对<code>error</code>进行处理，或将返回值赋值给明确忽略。对于<code>defer xx.Close()</code>可以不用显式处理。</li>
</ul>

<pre><code>func load() error {
	// normal code
}

// bad
load()

// good
 _ = load()
</code></pre>

<ul>
<li><code>error</code>作为函数的值返回且有多个返回值的时候，<code>error</code>必须是最后一个参数。</li>
</ul>

<pre><code>// bad
func load() (error, int) {
	// normal code
}

// good
func load() (int, error) {
	// normal code
}
</code></pre>

<ul>
<li>尽早进行错误处理，并尽早返回，减少嵌套。</li>
</ul>

<pre><code>// bad
if err != nil {
	// error code
} else {
	// normal code
}

// good
if err != nil {
	// error handling
	return err
}
// normal code
</code></pre>

<ul>
<li>如果需要在 if 之外使用函数调用的结果，则应采用下面的方式。</li>
</ul>

<pre><code>// bad
if v, err := foo(); err != nil {
	// error handling
}

// good
v, err := foo()
if err != nil {
	// error handling
}
</code></pre>

<ul>
<li>错误要单独判断，不与其他逻辑组合判断。</li>
</ul>

<pre><code>// bad
v, err := foo()
if err != nil || v  == nil {
	// error handling
	return err
}

// good
v, err := foo()
if err != nil {
	// error handling
	return err
}

if v == nil {
	// error handling
	return errors.New(&quot;invalid value v&quot;)
}
</code></pre>

<ul>
<li>如果返回值需要初始化，则采用下面的方式。</li>
</ul>

<pre><code>v, err := f()
if err != nil {
    // error handling
    return // or continue.
}
// use v
</code></pre>

<ul>
<li><p>错误描述建议</p>

<ul>
<li>告诉用户他们可以做什么，而不是告诉他们不能做什么。</li>
<li>当声明一个需求时，用must 而不是should。例如，must be greater than 0、must match regex ‘[a-z]+’。</li>
<li>当声明一个格式不对时，用must not。例如，must not contain。</li>
<li>当声明一个动作时用may not。例如，may not be specified when otherField is empty、only name may be specified。</li>
<li>引用文字字符串值时，请在单引号中指示文字。例如，ust not contain ‘…’。</li>
<li>当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than <code>request</code>。</li>
<li>指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more than、higher than)。</li>
<li>指定数字范围时，请尽可能使用包含范围。</li>
<li>建议 Go 1.13 以上，error 生成方式为 <code>fmt.Errorf(&quot;module xxx: %w&quot;, err)</code>。</li>
<li>错误描述用小写字母开头，结尾不要加标点符号，例如：</li>
</ul></li>
</ul>

<pre><code>	// bad
	errors.New(&quot;Redis connection failed&quot;)
	errors.New(&quot;redis connection failed.&quot;)

	// good
	errors.New(&quot;redis connection failed&quot;)
</code></pre>

<h3 id="1-4-panic处理">1.4 panic处理</h3>

<ul>
<li>在业务逻辑处理中禁止使用panic。</li>
<li>在main包中，只有当程序完全不可运行时使用panic，例如无法打开文件、无法连接数据库导致程序无法正常运行。</li>
<li>在main包中，使用 <code>log.Fatal</code> 来记录错误，这样就可以由log来结束程序，或者将panic抛出的异常记录到日志文件中，方便排查问题。</li>
<li>可导出的接口一定不能有panic。</li>
<li>包内建议采用error而不是panic来传递错误。</li>
</ul>

<h3 id="1-5-单元测试">1.5 单元测试</h3>

<ul>
<li>单元测试文件名命名规范为 <code>example_test.go</code>。</li>
<li>每个重要的可导出函数都要编写测试用例。</li>
<li>因为单元测试文件内的函数都是不对外的，所以可导出的结构体、函数等可以不带注释。</li>
<li>如果存在 <code>func (b *Bar) Foo</code> ，单测函数可以为 <code>func TestBar_Foo</code>。</li>
</ul>

<h3 id="1-6-类型断言失败处理">1.6 类型断言失败处理</h3>

<p>type assertion 的单个返回值针对不正确的类型将产生 panic。请始终使用 “comma ok”的惯用法。</p>

<pre><code>// bad
t := n.(int)

// good
t, ok := n.(int)
if !ok {
	// error handling
}
// normal code
</code></pre>

<h2 id="2-命名规范">2. 命名规范</h2>

<p>命名规范是代码规范中非常重要的一部分，一个统一的、短小的、精确的命名规范可以大大提高代码的可读性，也可以借此规避一些不必要的Bug。</p>

<h3 id="2-1-包命名">2.1 包命名</h3>

<ul>
<li>包名必须和目录名一致，尽量采取有意义、简短的包名，不要和标准库冲突。</li>
<li>包名全部小写，没有大写或下划线，使用多级目录来划分层级。</li>
<li>项目名可以通过中划线来连接多个单词。</li>
<li>包名以及包所在的目录名，不要使用复数，例如，是<code>net/url</code>，而不是<code>net/urls</code>。</li>
<li>不要用 common、util、shared 或者 lib 这类宽泛的、无意义的包名。</li>
<li>包名要简单明了，例如 net、time、log。</li>
</ul>

<h3 id="2-2-函数命名">2.2 函数命名</h3>

<ul>
<li>函数名采用驼峰式，首字母根据访问控制决定使用大写或小写，例如：MixedCaps或者mixedCaps。</li>
<li>代码生成工具自动生成的代码(如xxxx.pb.go)和为了对相关测试用例进行分组，而采用的下划线(如TestMyFunction_WhatIsBeingTested)排除此规则。</li>
</ul>

<h3 id="2-3-文件命名">2.3 文件命名</h3>

<ul>
<li>文件名要简短有意义。</li>
<li>文件名应小写，并使用下划线分割单词。</li>
</ul>

<h3 id="2-4-结构体命名">2.4 结构体命名</h3>

<ul>
<li>采用驼峰命名方式，首字母根据访问控制决定使用大写或小写，例如MixedCaps或者mixedCaps。</li>
<li>结构体名不应该是动词，应该是名词，比如 Node、NodeSpec。</li>
<li>避免使用Data、Info这类无意义的结构体名。</li>
<li>结构体的声明和初始化应采用多行，例如：</li>
</ul>

<pre><code>// User 多行声明
type User struct {
    Name  string
    Email string
}

// 多行初始化
u := User{
    UserName: &quot;colin&quot;,
    Email:    &quot;<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="abc8c4c7c2c59f9b9febcdc4d3c6cac2c785c8c4c6">[email&#160;protected]</a>&quot;,
}
</code></pre>

<h3 id="2-5-接口命名">2.5 接口命名</h3>

<ul>
<li><p>接口命名的规则，基本和结构体命名规则保持一致：</p>

<ul>
<li>单个函数的接口名以 “er&rdquo;”作为后缀（例如Reader，Writer），有时候可能导致蹩脚的英文，但是没关系。</li>
<li>两个函数的接口名以两个函数名命名，例如ReadWriter。</li>
<li>三个以上函数的接口名，类似于结构体名。</li>
</ul></li>
</ul>

<p>例如：</p>

<pre><code>	// Seeking to an offset before the start of the file is an error.
	// Seeking to any positive offset is legal, but the behavior of subsequent
	// I/O operations on the underlying object is implementation-dependent.
	type Seeker interface {
	    Seek(offset int64, whence int) (int64, error)
	}

	// ReadWriter is the interface that groups the basic Read and Write methods.
	type ReadWriter interface {
	    Reader
	    Writer
	}
</code></pre>

<h3 id="2-6-变量命名">2.6 变量命名</h3>

<ul>
<li><p>变量名必须遵循<strong>驼峰式</strong>，首字母根据访问控制决定使用大写或小写。</p></li>

<li><p>在相对简单（对象数量少、针对性强）的环境中，可以将一些名称由完整单词简写为单个字母，例如：</p>

<ul>
<li>user 可以简写为 u；</li>
<li>userID 可以简写 uid。</li>
</ul></li>

<li><p>特有名词时，需要遵循以下规则：</p>

<ul>
<li>如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient。</li>
<li>其他情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID。</li>
</ul></li>
</ul>

<p>下面列举了一些常见的特有名词。</p>

<pre><code>// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
    &quot;API&quot;:   true,
    &quot;ASCII&quot;: true,
    &quot;CPU&quot;:   true,
    &quot;CSS&quot;:   true,
    &quot;DNS&quot;:   true,
    &quot;EOF&quot;:   true,
    &quot;GUID&quot;:  true,
    &quot;HTML&quot;:  true,
    &quot;HTTP&quot;:  true,
    &quot;HTTPS&quot;: true,
    &quot;ID&quot;:    true,
    &quot;IP&quot;:    true,
    &quot;JSON&quot;:  true,
    &quot;LHS&quot;:   true,
    &quot;QPS&quot;:   true,
    &quot;RAM&quot;:   true,
    &quot;RHS&quot;:   true,
    &quot;RPC&quot;:   true,
    &quot;SLA&quot;:   true,
    &quot;SMTP&quot;:  true,
    &quot;SSH&quot;:   true,
    &quot;TLS&quot;:   true,
    &quot;TTL&quot;:   true,
    &quot;UI&quot;:    true,
    &quot;UID&quot;:   true,
    &quot;UUID&quot;:  true,
    &quot;URI&quot;:   true,
    &quot;URL&quot;:   true,
    &quot;UTF8&quot;:  true,
    &quot;VM&quot;:    true,
    &quot;XML&quot;:   true,
    &quot;XSRF&quot;:  true,
    &quot;XSS&quot;:   true,
}
</code></pre>

<ul>
<li>若变量类型为bool类型，则名称应以Has，Is，Can或Allow开头，例如：</li>
</ul>

<pre><code>var hasConflict bool
var isExist bool
var canManage bool
var allowGitHook bool
</code></pre>

<ul>
<li>局部变量应当尽可能短小，比如使用buf指代buffer，使用i指代index。</li>
<li>代码生成工具自动生成的代码可排除此规则(如xxx.pb.go里面的Id)</li>
</ul>

<h3 id="2-7-常量命名">2.7 常量命名</h3>

<ul>
<li>常量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。</li>
<li>如果是枚举类型的常量，需要先创建相应类型：</li>
</ul>

<pre><code>// Code defines an error code type.
type Code int

// Internal errors.
const (
    // ErrUnknown - 0: An unknown error occurred.
    ErrUnknown Code = iota
    // ErrFatal - 1: An fatal error occurred.
    ErrFatal
)
</code></pre>

<h3 id="2-8-error的命名">2.8 Error的命名</h3>

<ul>
<li>Error类型应该写成FooError的形式。</li>
</ul>

<pre><code>type ExitError struct {
	// ....
}
</code></pre>

<ul>
<li>Error变量写成ErrFoo的形式。</li>
</ul>

<pre><code>var ErrFormat = errors.New(&quot;unknown format&quot;)
</code></pre>

<h2 id="3-注释规范">3. 注释规范</h2>

<ul>
<li>每个可导出的名字都要有注释，该注释对导出的变量、函数、结构体、接口等进行简要介绍。</li>
<li>全部使用单行注释，禁止使用多行注释。</li>
<li>和代码的规范一样，单行注释不要过长，禁止超过 120 字符，超过的请使用换行展示，尽量保持格式优雅。</li>
<li>注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾，格式为 <code>// 名称 描述.</code> 。例如：</li>
</ul>

<pre><code>// bad
// logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	// normal code
}

// good
// PrintFlags logs the flags in the flagset.
func PrintFlags(flags *pflag.FlagSet) {
	// normal code
}
</code></pre>

<ul>
<li><p>所有注释掉的代码在提交code review前都应该被删除，否则应该说明为什么不删除，并给出后续处理建议。</p></li>

<li><p>在多段注释之间可以使用空行分隔加以区分，如下所示：</p></li>
</ul>

<pre><code>// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
</code></pre>

<h3 id="3-1-包注释">3.1 包注释</h3>

<ul>
<li>每个包都有且仅有一个包级别的注释。</li>
<li>包注释统一用 <code>//</code> 进行注释，格式为 <code>// Package 包名 包描述</code> ，例如：</li>
</ul>

<pre><code>// Package genericclioptions contains flags which can be added to you command, bound, completed, and produce
// useful helper functions.
package genericclioptions
</code></pre>

<h3 id="3-2-变量-常量注释">3.2 变量/常量注释</h3>

<ul>
<li>每个可导出的变量/常量都必须有注释说明，格式为<code>// 变量名 变量描述</code>，例如：</li>
</ul>

<pre><code>// ErrSigningMethod defines invalid signing method error.
var ErrSigningMethod = errors.New(&quot;Invalid signing method&quot;)
</code></pre>

<ul>
<li>出现大块常量或变量定义时，可在前面注释一个总的说明，然后在每一行常量的前一行或末尾详细注释该常量的定义，例如：</li>
</ul>

<pre><code>// Code must start with 1xxxxx.    
const (                         
    // ErrSuccess - 200: OK.          
    ErrSuccess int = iota + 100001    
                                                   
    // ErrUnknown - 500: Internal server error.    
    ErrUnknown    

    // ErrBind - 400: Error occurred while binding the request body to the struct.    
    ErrBind    
                                                  
    // ErrValidation - 400: Validation failed.    
    ErrValidation 
)
</code></pre>

<h3 id="3-3-结构体注释">3.3 结构体注释</h3>

<ul>
<li>每个需要导出的结构体或者接口都必须有注释说明，格式为 <code>// 结构体名 结构体描述.</code>。</li>
<li>结构体内的可导出成员变量名，如果意义不明确，必须要给出注释，放在成员变量的前一行或同一行的末尾。例如：</li>
</ul>

<pre><code>// User represents a user restful resource. It is also used as gorm model.
type User struct {
    // Standard object's metadata.
    metav1.ObjectMeta `json:&quot;metadata,omitempty&quot;`

    Nickname string `json:&quot;nickname&quot; gorm:&quot;column:nickname&quot;`
    Password string `json:&quot;password&quot; gorm:&quot;column:password&quot;`
    Email    string `json:&quot;email&quot; gorm:&quot;column:email&quot;`
    Phone    string `json:&quot;phone&quot; gorm:&quot;column:phone&quot;`
    IsAdmin  int    `json:&quot;isAdmin,omitempty&quot; gorm:&quot;column:isAdmin&quot;`
}
</code></pre>

<h3 id="3-4-方法注释">3.4 方法注释</h3>

<ul>
<li>每个需要导出的函数或者方法都必须有注释，格式为<code>// 函数名 函数描述.</code>，例如：</li>
</ul>

<pre><code>// BeforeUpdate run before update database record.
func (p *Policy) BeforeUpdate() (err error) {
	// normal code
	return nil
}
</code></pre>

<h3 id="3-5-类型注释">3.5 类型注释</h3>

<ul>
<li>每个需要导出的类型定义和类型别名都必须有注释说明，格式为 <code>// 类型名 类型描述.</code> ，例如：</li>
</ul>

<pre><code>// Code defines an error code type.
type Code int
</code></pre>

<h2 id="4-类型">4. 类型</h2>

<h3 id="4-1-字符串">4.1 字符串</h3>

<ul>
<li>空字符串判断。</li>
</ul>

<pre><code>// bad
if s == &quot;&quot; {
    // normal code
}

// good
if len(s) == 0 {
    // normal code
}
</code></pre>

<ul>
<li>[]byte/string相等比较。</li>
</ul>

<pre><code>// bad
var s1 []byte
var s2 []byte
...
bytes.Equal(s1, s2) == 0
bytes.Equal(s1, s2) != 0

// good
var s1 []byte
var s2 []byte
...
bytes.Compare(s1, s2) == 0
bytes.Compare(s1, s2) != 0
</code></pre>

<ul>
<li>复杂字符串使用raw字符串避免字符转义。</li>
</ul>

<pre><code>// bad
regexp.MustCompile(&quot;\\.&quot;)

// good
regexp.MustCompile(`\.`)
</code></pre>

<h3 id="4-2-切片">4.2 切片</h3>

<ul>
<li>空slice判断。</li>
</ul>

<pre><code>// bad
if len(slice) = 0 {
    // normal code
}

// good
if slice != nil &amp;&amp; len(slice) == 0 {
    // normal code
}
</code></pre>

<p>上面判断同样适用于map、channel。</p>

<ul>
<li>声明slice。</li>
</ul>

<pre><code>// bad
s := []string{}
s := make([]string, 0)

// good
var s []string
</code></pre>

<ul>
<li>slice复制。</li>
</ul>

<pre><code>// bad
var b1, b2 []byte
for i, v := range b1 {
   b2[i] = v
}
for i := range b1 {
   b2[i] = b1[i]
}

// good
copy(b2, b1)
</code></pre>

<ul>
<li>slice新增。</li>
</ul>

<pre><code>// bad
var a, b []int
for _, v := range a {
    b = append(b, v)
}

// good
var a, b []int
b = append(b, a...)
</code></pre>

<h3 id="4-3-结构体">4.3 结构体</h3>

<ul>
<li>struct初始化。</li>
</ul>

<p>struct以多行格式初始化。</p>

<pre><code>type user struct {
	Id   int64
	Name string
}

u1 := user{100, &quot;Colin&quot;}

u2 := user{
    Id:   200,
    Name: &quot;Lex&quot;,
}
</code></pre>

<h2 id="5-控制结构">5. 控制结构</h2>

<h3 id="5-1-if">5.1 if</h3>

<ul>
<li>if 接受初始化语句，约定如下方式建立局部变量。</li>
</ul>

<pre><code>if err := loadConfig(); err != nil {
	// error handling
	return err
}
</code></pre>

<ul>
<li>if 对于bool类型的变量，应直接进行真假判断。</li>
</ul>

<pre><code>var isAllow bool
if isAllow {
	// normal code
}
</code></pre>

<h3 id="5-2-for">5.2 for</h3>

<ul>
<li>采用短声明建立局部变量。</li>
</ul>

<pre><code>sum := 0
for i := 0; i &lt; 10; i++ {
    sum += 1
}
</code></pre>

<ul>
<li>不要在 for 循环里面使用 defer，defer只有在函数退出时才会执行。</li>
</ul>

<pre><code>// bad
for file := range files {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()
	// normal code
}

// good
for file := range files {
	func() {
		fd, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fd.Close()
		// normal code
	}()
}
</code></pre>

<h3 id="5-3-range">5.3 range</h3>

<ul>
<li>如果只需要第一项（key），就丢弃第二个。</li>
</ul>

<pre><code>for key := range keys {
// normal code
}
</code></pre>

<ul>
<li>如果只需要第二项，则把第一项置为下划线。</li>
</ul>

<pre><code>sum := 0
for _, value := range array {
    sum += value
}
</code></pre>

<h3 id="5-4-switch">5.4 switch</h3>

<ul>
<li>必须要有default。</li>
</ul>

<pre><code>switch os := runtime.GOOS; os {
    case &quot;linux&quot;:
        fmt.Println(&quot;Linux.&quot;)
    case &quot;darwin&quot;:
        fmt.Println(&quot;OS X.&quot;)
    default:
        fmt.Printf(&quot;%s.\n&quot;, os)
}
</code></pre>

<h3 id="5-5-goto">5.5 goto</h3>

<ul>
<li>业务代码禁止使用 <code>goto</code> 。</li>
<li>框架或其他底层源码尽量不用。</li>
</ul>

<h2 id="6-函数">6. 函数</h2>

<ul>
<li><p>传入变量和返回变量以小写字母开头。</p></li>

<li><p>函数参数个数不能超过5个。</p></li>

<li><p>函数分组与顺序</p>

<ul>
<li>函数应按粗略的调用顺序排序。</li>
<li>同一文件中的函数应按接收者分组。</li>
</ul></li>

<li><p>尽量采用值传递，而非指针传递。</p></li>

<li><p>传入参数是 map、slice、chan、interface ，不要传递指针。</p></li>
</ul>

<h3 id="6-1-函数参数">6.1 函数参数</h3>

<ul>
<li>如果函数返回相同类型的两个或三个参数，或者如果从上下文中不清楚结果的含义，使用命名返回，其他情况不建议使用命名返回，例如：</li>
</ul>

<pre><code>func coordinate() (x, y float64, err error) {
	// normal code
}
</code></pre>

<ul>
<li>传入变量和返回变量都以小写字母开头。</li>
<li>尽量用值传递，非指针传递。</li>
<li>参数数量均不能超过5个。</li>
<li>多返回值最多返回三个，超过三个请使用 struct。</li>
</ul>

<h3 id="6-2-defer">6.2 defer</h3>

<ul>
<li>当存在资源创建时，应紧跟defer释放资源(可以大胆使用defer，defer在Go1.14版本中，性能大幅提升，defer的性能损耗即使在性能敏感型的业务中，也可以忽略)。</li>
<li>先判断是否错误，再defer释放资源，例如：</li>
</ul>

<pre><code>rep, err := http.Get(url)
if err != nil {
    return err
}

defer resp.Body.Close()
</code></pre>

<h3 id="6-3-方法的接收器">6.3 方法的接收器</h3>

<ul>
<li>推荐以类名第一个英文首字母的小写作为接收器的命名。</li>
<li>接收器的命名在函数超过20行的时候不要用单字符。</li>
<li>接收器的命名不能采用me、this、self这类易混淆名称。</li>
</ul>

<h3 id="6-4-嵌套">6.4 嵌套</h3>

<ul>
<li>嵌套深度不能超过4层。</li>
</ul>

<h3 id="6-5-变量命名">6.5 变量命名</h3>

<ul>
<li>变量声明尽量放在变量第一次使用的前面，遵循就近原则。</li>
<li>如果魔法数字出现超过两次，则禁止使用，改用一个常量代替，例如：</li>
</ul>

<pre><code>// PI ...
const Prise = 3.14

func getAppleCost(n float64) float64 {
	return Prise * n
}

func getOrangeCost(n float64) float64 {
	return Prise * n
}
</code></pre>

<h2 id="7-gopath-设置规范">7. GOPATH 设置规范</h2>

<ul>
<li>Go 1.11 之后，弱化了 GOPATH 规则，已有代码（很多库肯定是在1.11之前建立的）肯定符合这个规则，建议保留 GOPATH 规则，便于维护代码。</li>
<li>建议只使用一个 GOPATH，不建议使用多个 GOPATH。如果使用多个GOPATH，编译生效的 bin 目录是在第一个 GOPATH 下。</li>
</ul>

<h2 id="8-依赖管理">8. 依赖管理</h2>

<ul>
<li>Go 1.11 以上必须使用 Go Modules。</li>
<li>使用Go Modules作为依赖管理的项目时，不建议提交vendor目录。</li>
<li>使用Go Modules作为依赖管理的项目时，必须提交go.sum文件。</li>
</ul>

<h2 id="9-最佳实践">9. 最佳实践</h2>

<ul>
<li>尽量少用全局变量，而是通过参数传递，使每个函数都是“无状态”的。这样可以减少耦合，也方便分工和单元测试。</li>
<li>在编译时验证接口的符合性，例如：</li>
</ul>

<pre><code>type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}
var _ http.Handler = LogHandler{}
</code></pre>

<ul>
<li>服务器处理请求时，应该创建一个context，保存该请求的相关信息（如requestID），并在函数调用链中传递。</li>
</ul>

<h3 id="9-1-性能">9.1 性能</h3>

<ul>
<li>string 表示的是不可变的字符串变量，对 string 的修改是比较重的操作，基本上都需要重新申请内存。所以，如果没有特殊需要，需要修改时多使用 []byte。</li>
<li>优先使用 strconv 而不是 fmt。</li>
</ul>

<h3 id="9-2-注意事项">9.2 注意事项</h3>

<ul>
<li>append 要小心自动分配内存，append 返回的可能是新分配的地址。</li>
<li>如果要直接修改 map 的 value 值，则 value 只能是指针，否则要覆盖原来的值。</li>
<li>map 在并发中需要加锁。</li>
<li>编译过程无法检查 interface{} 的转换，只能在运行时检查，小心引起 panic。</li>
</ul>

<h2 id="总结">总结</h2>

<p>这一讲，我向你介绍了九类常用的编码规范。但今天的最后，我要在这里提醒你一句：规范是人定的，你也可以根据需要，制定符合你项目的规范。这也是我在之前的课程里一直强调的思路。但同时我也建议你采纳这些业界沉淀下来的规范，并通过工具来确保规范的执行。</p>

<p>今天的内容就到这里啦，欢迎你在下面的留言区谈谈自己的看法，我们下一讲见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#a1cdcdcd989590909196e1c6ccc0c8cd8fc2cecc" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f0dfb7ddcd2653b',t:'MTczNDAwOTMxOC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>