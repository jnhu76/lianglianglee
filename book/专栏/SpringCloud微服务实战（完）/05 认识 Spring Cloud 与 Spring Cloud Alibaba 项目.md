<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=05&#32;认识&#32;Spring&#32;Cloud&#32;与&#32;Spring&#32;Cloud&#32;Alibaba&#32;项目>
        <link rel="icon" href="/static/favicon.png">
        <title>05 认识 Spring Cloud 与 Spring Cloud Alibaba 项目 </title>
        
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
                        <a class="menu-item" id="00 开篇导读.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/00%20%e5%bc%80%e7%af%87%e5%af%bc%e8%af%bb.md">00 开篇导读.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 以真实“商场停车”业务切入——需求分析.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/01%20%e4%bb%a5%e7%9c%9f%e5%ae%9e%e2%80%9c%e5%95%86%e5%9c%ba%e5%81%9c%e8%bd%a6%e2%80%9d%e4%b8%9a%e5%8a%a1%e5%88%87%e5%85%a5%e2%80%94%e2%80%94%e9%9c%80%e6%b1%82%e5%88%86%e6%9e%90.md">01 以真实“商场停车”业务切入——需求分析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 具象业务需求再抽象分解——系统设计.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/02%20%e5%85%b7%e8%b1%a1%e4%b8%9a%e5%8a%a1%e9%9c%80%e6%b1%82%e5%86%8d%e6%8a%bd%e8%b1%a1%e5%88%86%e8%a7%a3%e2%80%94%e2%80%94%e7%b3%bb%e7%bb%9f%e8%ae%be%e8%ae%a1.md">02 具象业务需求再抽象分解——系统设计.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 第一个 Spring Boot 子服务——会员服务.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/03%20%e7%ac%ac%e4%b8%80%e4%b8%aa%20Spring%20Boot%20%e5%ad%90%e6%9c%8d%e5%8a%a1%e2%80%94%e2%80%94%e4%bc%9a%e5%91%98%e6%9c%8d%e5%8a%a1.md">03 第一个 Spring Boot 子服务——会员服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 如何维护接口文档供外部调用——在线接口文档管理.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/04%20%e5%a6%82%e4%bd%95%e7%bb%b4%e6%8a%a4%e6%8e%a5%e5%8f%a3%e6%96%87%e6%a1%a3%e4%be%9b%e5%a4%96%e9%83%a8%e8%b0%83%e7%94%a8%e2%80%94%e2%80%94%e5%9c%a8%e7%ba%bf%e6%8e%a5%e5%8f%a3%e6%96%87%e6%a1%a3%e7%ae%a1%e7%90%86.md">04 如何维护接口文档供外部调用——在线接口文档管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 认识 Spring Cloud 与 Spring Cloud Alibaba 项目.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/05%20%e8%ae%a4%e8%af%86%20Spring%20Cloud%20%e4%b8%8e%20Spring%20Cloud%20Alibaba%20%e9%a1%b9%e7%9b%ae.md">05 认识 Spring Cloud 与 Spring Cloud Alibaba 项目.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 服务多不易管理如何破——服务注册与发现.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/06%20%e6%9c%8d%e5%8a%a1%e5%a4%9a%e4%b8%8d%e6%98%93%e7%ae%a1%e7%90%86%e5%a6%82%e4%bd%95%e7%a0%b4%e2%80%94%e2%80%94%e6%9c%8d%e5%8a%a1%e6%b3%a8%e5%86%8c%e4%b8%8e%e5%8f%91%e7%8e%b0.md">06 服务多不易管理如何破——服务注册与发现.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 如何调用本业务模块外的服务——服务调用.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/07%20%e5%a6%82%e4%bd%95%e8%b0%83%e7%94%a8%e6%9c%ac%e4%b8%9a%e5%8a%a1%e6%a8%a1%e5%9d%97%e5%a4%96%e7%9a%84%e6%9c%8d%e5%8a%a1%e2%80%94%e2%80%94%e6%9c%8d%e5%8a%a1%e8%b0%83%e7%94%a8.md">07 如何调用本业务模块外的服务——服务调用.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 服务响应慢或服务不可用怎么办——快速失败与服务降级.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/08%20%e6%9c%8d%e5%8a%a1%e5%93%8d%e5%ba%94%e6%85%a2%e6%88%96%e6%9c%8d%e5%8a%a1%e4%b8%8d%e5%8f%af%e7%94%a8%e6%80%8e%e4%b9%88%e5%8a%9e%e2%80%94%e2%80%94%e5%bf%ab%e9%80%9f%e5%a4%b1%e8%b4%a5%e4%b8%8e%e6%9c%8d%e5%8a%a1%e9%99%8d%e7%ba%a7.md">08 服务响应慢或服务不可用怎么办——快速失败与服务降级.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 热更新一样更新服务的参数配置——分布式配置中心.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/09%20%e7%83%ad%e6%9b%b4%e6%96%b0%e4%b8%80%e6%a0%b7%e6%9b%b4%e6%96%b0%e6%9c%8d%e5%8a%a1%e7%9a%84%e5%8f%82%e6%95%b0%e9%85%8d%e7%bd%ae%e2%80%94%e2%80%94%e5%88%86%e5%b8%83%e5%bc%8f%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83.md">09 热更新一样更新服务的参数配置——分布式配置中心.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 如何高效读取计费规则等热数据——分布式缓存.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/10%20%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e8%af%bb%e5%8f%96%e8%ae%a1%e8%b4%b9%e8%a7%84%e5%88%99%e7%ad%89%e7%83%ad%e6%95%b0%e6%8d%ae%e2%80%94%e2%80%94%e5%88%86%e5%b8%83%e5%bc%8f%e7%bc%93%e5%ad%98.md">10 如何高效读取计费规则等热数据——分布式缓存.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 多实例下的定时任务如何避免重复执行——分布式定时任务.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/11%20%e5%a4%9a%e5%ae%9e%e4%be%8b%e4%b8%8b%e7%9a%84%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e9%87%8d%e5%a4%8d%e6%89%a7%e8%a1%8c%e2%80%94%e2%80%94%e5%88%86%e5%b8%83%e5%bc%8f%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1.md">11 多实例下的定时任务如何避免重复执行——分布式定时任务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 同一套服务如何应对不同终端的需求——服务适配.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/12%20%e5%90%8c%e4%b8%80%e5%a5%97%e6%9c%8d%e5%8a%a1%e5%a6%82%e4%bd%95%e5%ba%94%e5%af%b9%e4%b8%8d%e5%90%8c%e7%bb%88%e7%ab%af%e7%9a%84%e9%9c%80%e6%b1%82%e2%80%94%e2%80%94%e6%9c%8d%e5%8a%a1%e9%80%82%e9%85%8d.md">12 同一套服务如何应对不同终端的需求——服务适配.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 采用消息驱动方式处理扣费通知——集成消息中间件.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/13%20%e9%87%87%e7%94%a8%e6%b6%88%e6%81%af%e9%a9%b1%e5%8a%a8%e6%96%b9%e5%bc%8f%e5%a4%84%e7%90%86%e6%89%a3%e8%b4%b9%e9%80%9a%e7%9f%a5%e2%80%94%e2%80%94%e9%9b%86%e6%88%90%e6%b6%88%e6%81%af%e4%b8%ad%e9%97%b4%e4%bb%b6.md">13 采用消息驱动方式处理扣费通知——集成消息中间件.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 Spring Cloud 与 Dubbo 冲突吗——强强联合.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/14%20Spring%20Cloud%20%e4%b8%8e%20Dubbo%20%e5%86%b2%e7%aa%81%e5%90%97%e2%80%94%e2%80%94%e5%bc%ba%e5%bc%ba%e8%81%94%e5%90%88.md">14 Spring Cloud 与 Dubbo 冲突吗——强强联合.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 破解服务中共性问题的繁琐处理方式——接入 API 网关.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/15%20%e7%a0%b4%e8%a7%a3%e6%9c%8d%e5%8a%a1%e4%b8%ad%e5%85%b1%e6%80%a7%e9%97%ae%e9%a2%98%e7%9a%84%e7%b9%81%e7%90%90%e5%a4%84%e7%90%86%e6%96%b9%e5%bc%8f%e2%80%94%e2%80%94%e6%8e%a5%e5%85%a5%20API%20%e7%bd%91%e5%85%b3.md">15 破解服务中共性问题的繁琐处理方式——接入 API 网关.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 服务压力大系统响应慢如何破——网关流量控制.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/16%20%e6%9c%8d%e5%8a%a1%e5%8e%8b%e5%8a%9b%e5%a4%a7%e7%b3%bb%e7%bb%9f%e5%93%8d%e5%ba%94%e6%85%a2%e5%a6%82%e4%bd%95%e7%a0%b4%e2%80%94%e2%80%94%e7%bd%91%e5%85%b3%e6%b5%81%e9%87%8f%e6%8e%a7%e5%88%b6.md">16 服务压力大系统响应慢如何破——网关流量控制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 集成网关后怎么做安全验证——统一鉴权.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/17%20%e9%9b%86%e6%88%90%e7%bd%91%e5%85%b3%e5%90%8e%e6%80%8e%e4%b9%88%e5%81%9a%e5%ae%89%e5%85%a8%e9%aa%8c%e8%af%81%e2%80%94%e2%80%94%e7%bb%9f%e4%b8%80%e9%89%b4%e6%9d%83.md">17 集成网关后怎么做安全验证——统一鉴权.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 多模块下的接口 API 如何统一管理——聚合 API.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/18%20%e5%a4%9a%e6%a8%a1%e5%9d%97%e4%b8%8b%e7%9a%84%e6%8e%a5%e5%8f%a3%20API%20%e5%a6%82%e4%bd%95%e7%bb%9f%e4%b8%80%e7%ae%a1%e7%90%86%e2%80%94%e2%80%94%e8%81%9a%e5%90%88%20API.md">18 多模块下的接口 API 如何统一管理——聚合 API.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 数据分库后如何确保数据完整性——分布式事务.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/19%20%e6%95%b0%e6%8d%ae%e5%88%86%e5%ba%93%e5%90%8e%e5%a6%82%e4%bd%95%e7%a1%ae%e4%bf%9d%e6%95%b0%e6%8d%ae%e5%ae%8c%e6%95%b4%e6%80%a7%e2%80%94%e2%80%94%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1.md">19 数据分库后如何确保数据完整性——分布式事务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 优惠券如何避免超兑——引入分布式锁.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/20%20%e4%bc%98%e6%83%a0%e5%88%b8%e5%a6%82%e4%bd%95%e9%81%bf%e5%85%8d%e8%b6%85%e5%85%91%e2%80%94%e2%80%94%e5%bc%95%e5%85%a5%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81.md">20 优惠券如何避免超兑——引入分布式锁.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 如何查看各服务的健康状况——系统应用监控.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/21%20%e5%a6%82%e4%bd%95%e6%9f%a5%e7%9c%8b%e5%90%84%e6%9c%8d%e5%8a%a1%e7%9a%84%e5%81%a5%e5%ba%b7%e7%8a%b6%e5%86%b5%e2%80%94%e2%80%94%e7%b3%bb%e7%bb%9f%e5%ba%94%e7%94%a8%e7%9b%91%e6%8e%a7.md">21 如何查看各服务的健康状况——系统应用监控.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 如何确定一次完整的请求过程——服务链路跟踪.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/22%20%e5%a6%82%e4%bd%95%e7%a1%ae%e5%ae%9a%e4%b8%80%e6%ac%a1%e5%ae%8c%e6%95%b4%e7%9a%84%e8%af%b7%e6%b1%82%e8%bf%87%e7%a8%8b%e2%80%94%e2%80%94%e6%9c%8d%e5%8a%a1%e9%93%be%e8%b7%af%e8%b7%9f%e8%b8%aa.md">22 如何确定一次完整的请求过程——服务链路跟踪.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 结束语.md" href="/%e4%b8%93%e6%a0%8f/SpringCloud%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%ae%9e%e6%88%98%ef%bc%88%e5%ae%8c%ef%bc%89/23%20%e7%bb%93%e6%9d%9f%e8%af%ad.md">23 结束语.md</a>
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
                            <h1 id="title" data-id="05 认识 Spring Cloud 与 Spring Cloud Alibaba 项目" class="title">05 认识 Spring Cloud 与 Spring Cloud Alibaba 项目</h1>
                            <div><p>前面我们已经粗略的将项目的骨架搭建完成，并初步引入一些基础支撑功能，为后续的开发打好基础。本篇将介绍 Spring Cloud 及 Spring Cloud Alibaba 两个项目，从理论角度作个整体性掌握，后续进入开发实战作好铺垫工作。 Spring Cloud Alibaba 是 Spring Cloud 的一个子项目，在介绍 Spring Cloud Alibaba 之前，先简单聊一聊 Spring Cloud 的情况。</p>

<h3 id="spring-cloud-介绍">Spring Cloud 介绍</h3>

<p>Spring Cloud 官方文档地址：<a href="https://spring.io/projects/spring-cloud" target="_blank">https://spring.io/projects/spring-cloud</a></p>

<p>它是由很多个组件共同组成的一套微服务技术体系解决方案，目前最新版本是 Hoxton，它的版本并不是我们常见的大版本、小版本的数字形式，Spring Cloud 的版本规划是按伦敦地铁站的名称先后顺序来规划的，目的是为了更好的管理每个 Spring Cloud 子项目的版本，避免自己的版本与子项目的版本号混淆，所以要特别注意两个项目的版本对应情况，以免实际应用中产生不必要的麻烦。</p>

<p>随着新版本的迭代更新，有些低版本的 Spring Cloud 的不建议再应用于生产，比如 Brixton 和 Angel 两个版本在 2017 年已经寿终正寝。 <img src="assets/2020-05-05-021720.jpg" alt="img" /> （ Spring Cloud 与 Spring Boot 的版本对应情况，来源于官网） <img src="assets/2020-05-05-021722.jpg" alt="img" /> ( 某 Spring Cloud Greenwich 版本与子项目的版本对应情况 ) Spring Cloud 基于 Spring Boot 对外提供一整套的微服务架构体系的解决方案，包括配置管理、服务注册与服务发现、路由、端到端的调用、负载均衡、断路器、全局锁、分布式消息等，对于这些功能 Spring Cloud 提供了多种项目选择，可从官网的主要项目列表一窥端倪。 <img src="assets/2020-05-05-021726.jpg" alt="img" /> 看上图有个比较突出的子项目： Spring Cloud netflix，由 Netflix 开发后来又并入 Spring Cloud 大家庭，它主要提供的模块包括：服务发现、断路器和监控、智能路由、客户端负载均衡等。但随着 Spring Cloud 的迭代，不少 Netflix 的组件进行了维护模式，最明显的莫过于 Spring Cloud Gateway 的推出来替代旧有的 Zuul 组件，有项目加入，也会有老旧项目退出舞台，这也是产品迭代的正常节奏。</p>

<p>这些子项目，极大的丰富了 Spring Cloud 在微服务领域中应用范围，几乎无需要借助外部组件，以一已之力打造全生态的微服务架构，并与外部基础运维组件更好的融合在一起。</p>

<p>下面再介绍一个近两年出现的一个新生项目 ： Spring Cloud Alibaba。</p>

<h3 id="spring-cloud-alibaba-介绍">Spring Cloud Alibaba 介绍</h3>

<p>官网地址：<a href="https://github.com/alibaba/spring-cloud-alibaba" target="_blank">https://github.com/alibaba/spring-cloud-alibaba</a> 它是 Spring Cloud 的一个子项目，致力于提供微服务开发的一站式解决方案，项目包含开发分布式应用服务的必需组件，方便开发者通过 Spring Cloud 编程模型轻松使用这些组件来开发分布式应用服务，只需要添加一些注解和少量配置，就可以将 Spring Cloud 应用接入阿里分布式应用解决方案，通过阿里中间件来迅速搭建分布式应用系统。 项目特性见下图： <img src="assets/2020-05-05-021728.jpg" alt="img" /> 包括一些关键组件：</p>

<ul>
<li><strong>Sentinel</strong>：把流量作为切入点，从流量控制、熔断降级、系统负载保护等多个维度保护服务的稳定性。与 Netflix 的 Hystrix 组件类似，但实现方式上更为轻量。</li>
<li><strong>Nacos</strong>：一个更易于构建云原生应用的动态服务发现、配置管理和服务管理平台，同时具备了之前 Netflix Eureka 和 Spring Cloud Config 的功能，而且 UI 操作上更加人性化。</li>
<li><strong>RocketMQ</strong>：一款开源的分布式消息系统，基于高可用分布式集群技术，提供低延时的、高可靠的消息发布与订阅服务，目前已交由 Apache 组织维护。</li>
<li><strong>Dubbo</strong>：Apache Dubbo™ 是一款高性能 Java RPC 框架，自交由 Apache 组织孵化后，目前社区生态很活跃，产生形态越来越丰富。</li>
<li><strong>Seata</strong>：阿里巴巴开源产品，一个易于使用的高性能微服务分布式事务解决方案，由早期内部产品 Fescar 演变而来。</li>
</ul>

<blockquote>
<p>以上组件都是均是开源实现方案。下面提到的几个组件都是结合阿里云的产品形态完成的功能，后续的案例开发实战不引入商业产品，需要的小伙伴可以购买后拿到对应的 API 直接接入即可。</p>
</blockquote>

<ul>
<li><strong>Alibaba Cloud ACM</strong>：一款在分布式架构环境中对应用配置进行集中管理和推送的应用配置中心产品，与开源产品 Nacos 功能类似。</li>
<li><strong>Alibaba Cloud OSS</strong>: 阿里云对象存储服务（Object Storage Service，简称 OSS），是阿里云提供的海量、安全、低成本、高可靠的云存储服务。您可以在任何应用、任何时间、任何地点存储和访问任意类型的数据。</li>
<li><strong>Alibaba Cloud SchedulerX</strong>: 阿里中间件团队开发的一款分布式任务调度产品，提供秒级、精准、高可靠、高可用的定时（基于 Cron 表达式）任务调度服务。</li>
<li><strong>Alibaba Cloud SMS</strong>: 覆盖全球的短信服务，友好、高效、智能的互联化通讯能力，帮助企业迅速搭建客户触达通道。</li>
</ul>

<p>在使用过程中，版本问题同样需要关注。 <img src="assets/2020-05-05-021730.jpg" alt="img" />（截图来源于 Spring 官网）</p>

<h3 id="spring-boot-介绍">Spring Boot 介绍</h3>

<p>不管是 Spring Cloud 还是 Spring Cloud Alibaba 项目，都是基于 Spring Boot ，所以先将 Spring Boot 掌握后，才能更好的应用这两个项目。提供的大量的 starter 组件，更是方便我们快速的应用相应的功能，由于其内置了应用容器( Tomcat ，Jetty ，Undertow )，无须再构建成 war 文件去部署，并遵从约定优于配置的原则，高效开发应用。</p>

<blockquote>
<p>下面的链接是一份 Spring Boot 的全量参数配置，相信对你会有帮助。 <a href="https://docs.spring.io/spring-boot/docs/2.2.4.RELEASE/reference/html/appendix-application-properties.html" target="_blank">https://docs.spring.io/spring-boot/docs/2.2.4.RELEASE/reference/html/appendix-application-properties.html</a></p>
</blockquote>

<p>本篇介绍了课程中使用到的三个关键项目，这里只是做个简单的概念了解，后续将结合实际业务进入开发工作。采用 Spring Boot 构建项目，一般直接是 jar 的形式运行，但有些小伙伴还是有偏爱 war 包情况，哪用 Spring Boot 搭建的项目如何构建出 war 包呢？</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#0f636363363b3e3e3f384f68626e6663216c6062" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f127fb2cbf9ede4',t:'MTczNDA1NjY3Ni4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>