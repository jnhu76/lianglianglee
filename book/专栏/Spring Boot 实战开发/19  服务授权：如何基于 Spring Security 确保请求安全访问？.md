<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=19&#32;&#32;服务授权：如何基于&#32;Spring&#32;Security&#32;确保请求安全访问？>
        <link rel="icon" href="/static/favicon.png">
        <title>19  服务授权：如何基于 Spring Security 确保请求安全访问？ </title>
        
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
                        <a class="menu-item" id="00 开篇词  从零开始：为什么要学习 Spring Boot？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%20%e4%bb%8e%e9%9b%b6%e5%bc%80%e5%a7%8b%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ad%a6%e4%b9%a0%20Spring%20Boot%ef%bc%9f.md">00 开篇词  从零开始：为什么要学习 Spring Boot？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  家族生态：如何正确理解 Spring 家族的技术体系？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/01%20%20%e5%ae%b6%e6%97%8f%e7%94%9f%e6%80%81%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e7%90%86%e8%a7%a3%20Spring%20%e5%ae%b6%e6%97%8f%e7%9a%84%e6%8a%80%e6%9c%af%e4%bd%93%e7%b3%bb%ef%bc%9f.md">01  家族生态：如何正确理解 Spring 家族的技术体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  案例驱动：如何剖析一个 Spring Web 应用程序？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/02%20%20%e6%a1%88%e4%be%8b%e9%a9%b1%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%89%96%e6%9e%90%e4%b8%80%e4%b8%aa%20Spring%20Web%20%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">02  案例驱动：如何剖析一个 Spring Web 应用程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  多维配置：如何使用 Spring Boot 中的配置体系？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/03%20%20%e5%a4%9a%e7%bb%b4%e9%85%8d%e7%bd%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Spring%20Boot%20%e4%b8%ad%e7%9a%84%e9%85%8d%e7%bd%ae%e4%bd%93%e7%b3%bb%ef%bc%9f.md">03  多维配置：如何使用 Spring Boot 中的配置体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  定制配置：如何创建和管理自定义的配置信息？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/04%20%20%e5%ae%9a%e5%88%b6%e9%85%8d%e7%bd%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e5%88%9b%e5%bb%ba%e5%92%8c%e7%ae%a1%e7%90%86%e8%87%aa%e5%ae%9a%e4%b9%89%e7%9a%84%e9%85%8d%e7%bd%ae%e4%bf%a1%e6%81%af%ef%bc%9f.md">04  定制配置：如何创建和管理自定义的配置信息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  自动配置：如何正确理解 Spring Boot 自动配置实现原理？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/05%20%20%e8%87%aa%e5%8a%a8%e9%85%8d%e7%bd%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e7%90%86%e8%a7%a3%20Spring%20Boot%20%e8%87%aa%e5%8a%a8%e9%85%8d%e7%bd%ae%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%ef%bc%9f.md">05  自动配置：如何正确理解 Spring Boot 自动配置实现原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  基础规范：如何理解 JDBC 关系型数据库访问规范？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/06%20%20%e5%9f%ba%e7%a1%80%e8%a7%84%e8%8c%83%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%20JDBC%20%e5%85%b3%e7%b3%bb%e5%9e%8b%e6%95%b0%e6%8d%ae%e5%ba%93%e8%ae%bf%e9%97%ae%e8%a7%84%e8%8c%83%ef%bc%9f.md">06  基础规范：如何理解 JDBC 关系型数据库访问规范？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  数据访问：如何使用 JdbcTemplate 访问关系型数据库？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/07%20%20%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20JdbcTemplate%20%e8%ae%bf%e9%97%ae%e5%85%b3%e7%b3%bb%e5%9e%8b%e6%95%b0%e6%8d%ae%e5%ba%93%ef%bc%9f.md">07  数据访问：如何使用 JdbcTemplate 访问关系型数据库？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  数据访问：如何剖析 JdbcTemplate 数据访问实现原理？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/08%20%20%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%ef%bc%9a%e5%a6%82%e4%bd%95%e5%89%96%e6%9e%90%20JdbcTemplate%20%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%ef%bc%9f.md">08  数据访问：如何剖析 JdbcTemplate 数据访问实现原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  数据抽象：Spring Data 如何对数据访问过程进行统一抽象？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/09%20%20%e6%95%b0%e6%8d%ae%e6%8a%bd%e8%b1%a1%ef%bc%9aSpring%20Data%20%e5%a6%82%e4%bd%95%e5%af%b9%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%e8%bf%87%e7%a8%8b%e8%bf%9b%e8%a1%8c%e7%bb%9f%e4%b8%80%e6%8a%bd%e8%b1%a1%ef%bc%9f.md">09  数据抽象：Spring Data 如何对数据访问过程进行统一抽象？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  ORM 集成：如何使用 Spring Data JPA 访问关系型数据库？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/10%20%20ORM%20%e9%9b%86%e6%88%90%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Spring%20Data%20JPA%20%e8%ae%bf%e9%97%ae%e5%85%b3%e7%b3%bb%e5%9e%8b%e6%95%b0%e6%8d%ae%e5%ba%93%ef%bc%9f.md">10  ORM 集成：如何使用 Spring Data JPA 访问关系型数据库？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  服务发布：如何构建一个 RESTful 风格的 Web 服务？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/11%20%20%e6%9c%8d%e5%8a%a1%e5%8f%91%e5%b8%83%ef%bc%9a%e5%a6%82%e4%bd%95%e6%9e%84%e5%bb%ba%e4%b8%80%e4%b8%aa%20RESTful%20%e9%a3%8e%e6%a0%bc%e7%9a%84%20Web%20%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">11  服务发布：如何构建一个 RESTful 风格的 Web 服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  服务调用：如何使用 RestTemplate 消费 RESTful 服务？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/12%20%20%e6%9c%8d%e5%8a%a1%e8%b0%83%e7%94%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20RestTemplate%20%e6%b6%88%e8%b4%b9%20RESTful%20%e6%9c%8d%e5%8a%a1%ef%bc%9f.md">12  服务调用：如何使用 RestTemplate 消费 RESTful 服务？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  服务调用：如何正确理解 RestTemplate 远程调用实现原理？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/13%20%20%e6%9c%8d%e5%8a%a1%e8%b0%83%e7%94%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e7%90%86%e8%a7%a3%20RestTemplate%20%e8%bf%9c%e7%a8%8b%e8%b0%83%e7%94%a8%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86%ef%bc%9f.md">13  服务调用：如何正确理解 RestTemplate 远程调用实现原理？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  消息驱动：如何使用 KafkaTemplate 集成 Kafka？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/14%20%20%e6%b6%88%e6%81%af%e9%a9%b1%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20KafkaTemplate%20%e9%9b%86%e6%88%90%20Kafka%ef%bc%9f.md">14  消息驱动：如何使用 KafkaTemplate 集成 Kafka？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  消息驱动：如何使用 JmsTemplate 集成 ActiveMQ？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/15%20%20%e6%b6%88%e6%81%af%e9%a9%b1%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20JmsTemplate%20%e9%9b%86%e6%88%90%20ActiveMQ%ef%bc%9f.md">15  消息驱动：如何使用 JmsTemplate 集成 ActiveMQ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  消息驱动：如何使用 RabbitTemplate 集成 RabbitMQ？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/16%20%20%e6%b6%88%e6%81%af%e9%a9%b1%e5%8a%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20RabbitTemplate%20%e9%9b%86%e6%88%90%20RabbitMQ%ef%bc%9f.md">16  消息驱动：如何使用 RabbitTemplate 集成 RabbitMQ？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  安全架构：如何理解 Spring 安全体系的整体架构？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/17%20%20%e5%ae%89%e5%85%a8%e6%9e%b6%e6%9e%84%ef%bc%9a%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%20Spring%20%e5%ae%89%e5%85%a8%e4%bd%93%e7%b3%bb%e7%9a%84%e6%95%b4%e4%bd%93%e6%9e%b6%e6%9e%84%ef%bc%9f.md">17  安全架构：如何理解 Spring 安全体系的整体架构？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  用户认证：如何基于 Spring Security 构建用户认证体系？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/18%20%20%e7%94%a8%e6%88%b7%e8%ae%a4%e8%af%81%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e6%9e%84%e5%bb%ba%e7%94%a8%e6%88%b7%e8%ae%a4%e8%af%81%e4%bd%93%e7%b3%bb%ef%bc%9f.md">18  用户认证：如何基于 Spring Security 构建用户认证体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  服务授权：如何基于 Spring Security 确保请求安全访问？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/19%20%20%e6%9c%8d%e5%8a%a1%e6%8e%88%e6%9d%83%ef%bc%9a%e5%a6%82%e4%bd%95%e5%9f%ba%e4%ba%8e%20Spring%20Security%20%e7%a1%ae%e4%bf%9d%e8%af%b7%e6%b1%82%e5%ae%89%e5%85%a8%e8%ae%bf%e9%97%ae%ef%bc%9f.md">19  服务授权：如何基于 Spring Security 确保请求安全访问？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  服务监控：如何使用 Actuator 组件实现系统监控？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/20%20%20%e6%9c%8d%e5%8a%a1%e7%9b%91%e6%8e%a7%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Actuator%20%e7%bb%84%e4%bb%b6%e5%ae%9e%e7%8e%b0%e7%b3%bb%e7%bb%9f%e7%9b%91%e6%8e%a7%ef%bc%9f.md">20  服务监控：如何使用 Actuator 组件实现系统监控？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  指标定制：如何实现自定义度量指标和 Actuator 端点？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/21%20%20%e6%8c%87%e6%a0%87%e5%ae%9a%e5%88%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0%e8%87%aa%e5%ae%9a%e4%b9%89%e5%ba%a6%e9%87%8f%e6%8c%87%e6%a0%87%e5%92%8c%20Actuator%20%e7%ab%af%e7%82%b9%ef%bc%9f.md">21  指标定制：如何实现自定义度量指标和 Actuator 端点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  运行管理：如何使用 Admin Server 管理 Spring 应用程序？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/22%20%20%e8%bf%90%e8%a1%8c%e7%ae%a1%e7%90%86%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Admin%20Server%20%e7%ae%a1%e7%90%86%20Spring%20%e5%ba%94%e7%94%a8%e7%a8%8b%e5%ba%8f%ef%bc%9f.md">22  运行管理：如何使用 Admin Server 管理 Spring 应用程序？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  数据测试：如何使用 Spring 测试数据访问层组件？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/23%20%20%e6%95%b0%e6%8d%ae%e6%b5%8b%e8%af%95%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Spring%20%e6%b5%8b%e8%af%95%e6%95%b0%e6%8d%ae%e8%ae%bf%e9%97%ae%e5%b1%82%e7%bb%84%e4%bb%b6%ef%bc%9f.md">23  数据测试：如何使用 Spring 测试数据访问层组件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  服务测试：如何使用 Spring 测试 Web 服务层组件？.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/24%20%20%e6%9c%8d%e5%8a%a1%e6%b5%8b%e8%af%95%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bd%bf%e7%94%a8%20Spring%20%e6%b5%8b%e8%af%95%20Web%20%e6%9c%8d%e5%8a%a1%e5%b1%82%e7%bb%84%e4%bb%b6%ef%bc%9f.md">24  服务测试：如何使用 Spring 测试 Web 服务层组件？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  以终为始：Spring Boot 总结和展望.md" href="/%e4%b8%93%e6%a0%8f/Spring%20Boot%20%e5%ae%9e%e6%88%98%e5%bc%80%e5%8f%91/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e4%bb%a5%e7%bb%88%e4%b8%ba%e5%a7%8b%ef%bc%9aSpring%20Boot%20%e6%80%bb%e7%bb%93%e5%92%8c%e5%b1%95%e6%9c%9b.md">结束语  以终为始：Spring Boot 总结和展望.md</a>
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
                            <h1 id="title" data-id="19  服务授权：如何基于 Spring Security 确保请求安全访问？" class="title">19  服务授权：如何基于 Spring Security 确保请求安全访问？</h1>
                            <div><p>18 讲中，我们集中讨论了如何通过 WebSecurityConfigurerAdapter 完成对用户认证体系的构建。这一讲我们将继续使用这个配置类完成对服务访问的授权控制。</p>

<p>在日常开发过程中，我们需要对 Web 应用中的不同 HTTP 端点进行不同粒度的权限控制，并且希望这种控制方法足够灵活。而借助 Spring Security 框架，我们就可以对其进行简单实现，下面我们一起来看下。</p>

<h3 id="对-http-端点进行访问授权管理">对 HTTP 端点进行访问授权管理</h3>

<p>在一个 Web 应用中，权限管理的对象是通过 Controller 层暴露的一个个 HTTP 端点，而这些 HTTP 端点就是需要授权访问的资源。</p>

<p>开发人员使用 Spring Security 中提供的一系列丰富技术组件，即可通过简单的设置对权限进行灵活管理。</p>

<h4 id="使用配置方法">使用配置方法</h4>

<p>实现访问授权的第一种方法是使用配置方法，关于配置方法的处理过程也是位于 WebSecurityConfigurerAdapter 类中，但使用的是 configure(HttpSecurity http) 方法，如下代码所示：</p>

<pre><code>protected void configure(HttpSecurity http) throws Exception {

 

        http

            .authorizeRequests()

            .anyRequest()

            .authenticated()

            .and()

            .formLogin()

            .and()

            .httpBasic();

}
</code></pre>

<p>上述代码就是 Spring Security 中作用于访问授权的默认实现方法，这里用到了多个常见的配置方法。</p>

<p>回想 18 课时中的内容，访问任何端点时，一旦在代码类路径中引入了 Spring Security 框架，就会弹出一个登录界面从而完成用户认证。因为认证是授权的前置流程，认证结束后就可以进入授权环节。</p>

<p>结合这些配置方法的名称，我们简单分析一下实现这种默认的授权效果的具体步骤。</p>

<p>首先，通过 HttpSecurity 类的 authorizeRequests() 方法，我们可以对所有访问 HTTP 端点的 HttpServletRequest 进行限制。</p>

<p>其次，anyRequest().authenticated() 语句指定了所有请求都需要执行认证，也就是说没有通过认证的用户无法访问任何端点。</p>

<p>然后，formLogin() 语句指定了用户需要使用表单进行登录，即会弹出一个登录界面。</p>

<p>最后， httpBasic() 语句使用 HTTP 协议中的 Basic Authentication 方法完成认证。</p>

<p>18 讲中我们也演示了如何使用 Postman 完成认证的方式，这里就不过多赘述了。</p>

<p>当然，Spring Security 中还提供了很多其他有用的配置方法供开发人员灵活使用，下表中我们进行了列举，一起来看下。</p>

<p><img src="assets/CgpVE2AGpjqARlFfAAE9RHgwruA457.png" alt="Lark20210119-172757.png" /></p>

<p><strong>基于上表中的配置方法，我们就可以通过 HttpSecurity 实现自定义的授权策略。</strong></p>

<p>比方说，我们希望针对“/orders”根路径下的所有端点进行访问控制，且只允许认证通过的用户访问，那么可以创建一个继承了 WebSecurityConfigurerAdapter 类的 SpringCssSecurityConfig，并覆写其中的 configure(HttpSecurity http) 方法来实现，如下代码所示：</p>

<pre><code>@Configuration

public class SecurityConfig extends WebSecurityConfigurerAdapter {

 

    @Override

    public void configure(HttpSecurity http) throws Exception {

 

        http.authorizeRequests()

            .antMatchers(&quot;/orders/**&quot;)

            .authenticated();

    }

}
</code></pre>

<p><strong>请注意：虽然上表中的这些配置方法非常有用，但是由于我们无法基于一些来自环境和业务的参数灵活控制访问规则，也就存在一定的局限性。</strong></p>

<p>为此，Spring Security 还提供了一个 access() 方法，该方法允许开发人员传入一个表达式进行更细粒度的权限控制，这里，我们将引入Spring 框架提供的一种动态表达式语言—— SpEL（Spring Expression Language 的简称）。</p>

<p>只要 SpEL 表达式的返回值为 true，access() 方法就允许用户访问，如下代码所示：</p>

<pre><code>@Override

public void configure(HttpSecurity http) throws Exception {

 

        http.authorizeRequests()

            .antMatchers(&quot;/orders&quot;)

            .access(&quot;hasRole('ROLE_USER')&quot;);

}
</code></pre>

<p>上述代码中，假设访问“/orders”端点的请求必须具备“ROLE_USER”角色，通过 access 方法中的 hasRole 方法我们即可灵活地实现这个需求。当然，除了使用 hasRole 外，我们还可以使用 authentication、isAnonymous、isAuthenticated、permitAll 等表达式进行实现。因这些表达式的作用与前面介绍的配置方法一致，我们就不过多赘述。</p>

<h4 id="使用注解">使用注解</h4>

<p>除了使用配置方法，Spring Security 还为我们提供了 @PreAuthorize 注解实现类似的效果，该注解定义如下代码所示：</p>

<pre><code>@Target({ ElementType.METHOD, ElementType.TYPE })

@Retention(RetentionPolicy.RUNTIME)

@Inherited

@Documented

public @interface PreAuthorize {

 

    //通过 SpEL 表达式设置访问控制

    String value();

}
</code></pre>

<p>可以看到 @PreAuthorize 的原理与前面介绍的 access() 方法一样，即通过传入一个 SpEL 表达式设置访问控制，如下所示代码就是一个典型的使用示例：</p>

<pre><code>@RestController

@RequestMapping(value=&quot;orders&quot;)

public class OrderController {

 

    @PostMapping(value = &quot;/&quot;)

    @PreAuthorize(&quot;hasRole(ROLE_ADMIN)&quot;)

    public void addOrder(@RequestBody Order order) {

        …

    }

}
</code></pre>

<p>从这个示例中可以看到，在“/orders/”这个 HTTP 端点上，我们添加了一个 @PreAuthorize 注解用来限制只有角色为“ROLE_ADMIN”的用户才能访问该端点。</p>

<p>其实，Spring Security 中用于授权的注解还有 @PostAuthorize，它与 @PreAuthorize 注解是一组，主要用于请求结束之后检查权限。因这种情况比较少见，这里我们不再继续展开，你可以翻阅相关资料学习。</p>

<h3 id="实现多维度访问授权方案">实现多维度访问授权方案</h3>

<p>我们知道 HTTP 端点是 Web 应用程序的一种资源，而每个 Web 应用程序对于自身资源的保护粒度因服务而异。对于一般的 HTTP 端点，用户可能通过认证就可以访问；对于一些重要的 HTTP 端点，用户在已认证的基础上还会有一些附加要求。</p>

<p>接下来，我们将讨论对资源进行保护的三种粒度级别。</p>

<ul>
<li><strong>用户级别：</strong> 该级别是最基本的资源保护级别，只要是认证用户就可能访问服务内的各种资源。</li>
<li><strong>用户+角色级别：</strong> 该级别在认证用户级别的基础上，还要求用户属于某一个或多个特定角色。</li>
<li><strong>用户+角色+操作级别：</strong> 该级别在认证用户+角色级别的基础上，对某些 HTTP 操作方法做了访问限制。</li>
</ul>

<p>基于配置方法和注解，我们可以轻松实现上述三种访问授权方案。</p>

<h4 id="使用用户级别保护服务访问">使用用户级别保护服务访问</h4>

<p>这次，我们来到 SpringCSS 案例系统中的 customer-service，先来回顾一下 CustomerController 的内容，如下所示：</p>

<pre><code>@RestController

@RequestMapping(value=&quot;customers&quot;)

public class CustomerController {

    @Autowired

    private CustomerTicketService customerTicketService; 

    @PostMapping(value = &quot;/{accountId}/{orderNumber}&quot;)

    public CustomerTicket generateCustomerTicket( @PathVariable(&quot;accountId&quot;) Long accountId,

            @PathVariable(&quot;orderNumber&quot;) String orderNumber) {

        CustomerTicket customerTicket = customerTicketService.generateCustomerTicket(accountId, orderNumber);

        return customerTicket;

    }

    @GetMapping(value = &quot;/{id}&quot;)

    public CustomerTicket getCustomerTicketById(@PathVariable Long id) {

        CustomerTicket customerTicket = customerTicketService.getCustomerTicketById(id);

     return customerTicket;

    }

    @GetMapping(value = &quot;/{pageIndex}/{pageSize}&quot;)

    public List&lt;CustomerTicket&gt; getCustomerTicketList( @PathVariable(&quot;pageIndex&quot;) int pageIndex, @PathVariable(&quot;pageSize&quot;) int pageSize) {

        List&lt;CustomerTicket&gt; customerTickets = customerTicketService.getCustomerTickets(pageIndex, pageSize);

        return customerTickets;

    }

 

	@DeleteMapping(value = &quot;/{id}&quot;)

    public void deleteCustomerTicket( @PathVariable(&quot;id&quot;) Long id) {

        customerTicketService.deleteCustomerTicket(id);

    }

}
</code></pre>

<p>因为 CustomerController 是 SpringCSS 案例中的核心入口，所以我们认为它的所有端点都应该受到保护。于是，在 customer-service 中，我们创建了一个 SpringCssSecurityConfig 类继承 WebSecurityConfigurerAdapter，如下代码所示：</p>

<pre><code>@Configuration

public class SpringCssSecurityConfig extends WebSecurityConfigurerAdapter {

 

    @Override

    public void configure(HttpSecurity http) throws Exception {

 

        http.authorizeRequests()

            .anyRequest()

            .authenticated();

    }

}
</code></pre>

<p>位于 configure() 方法中的 .anyRequest().authenticated() 语句指定了访问 customer-service 下的所有端点的任何请求都需要进行验证。因此，当我们使用普通的 HTTP 请求访问 CustomerController 中的任何 URL（例如<a href="http://localhost:8083/customers/1" target="_blank">http://localhost:8083/customers/1</a>），将会得到如下图代码所示的错误信息，该错误信息明确指出资源的访问需要进行认证。</p>

<pre><code>{

    &quot;error&quot;: &quot;access_denied&quot;,

    &quot;error_description&quot;: &quot;Full authentication is required to access to this resource&quot;

}
</code></pre>

<p>记得 18 讲中覆写 WebSecurityConfigurerAdapter 的 config(AuthenticationManagerBuilder auth) 方法时提供了一个用户名“springcss_user”，现在我们就用这个用户名来添加用户认证信息并再次访问该端点。显然，因为此时我们传入的是有效的用户信息，所以可以满足认证要求。</p>

<h4 id="使用用户-角色级别保护服务访问">使用用户+角色级别保护服务访问</h4>

<p>对于某些安全性要求比较高的 HTTP 端点，我们通常需要限定访问的角色。</p>

<p>例如，customer-service 服务中涉及客户工单管理等核心业务，我们认为不应该给所有的认证用户开放资源访问入口，而应该限定只有角色为“ADMIN”的管理员才开放。这时，我们就可以使用认证用户+角色保护服务的访问控制机制，具体的示例代码如下所示：</p>

<pre><code>@Configuration

public class SpringCssSecurityConfig extends WebSecurityConfigurerAdapter {

 

    @Override

    public void configure(HttpSecurity http) throws Exception {

 

        http.authorizeRequests()

            .antMatchers(&quot;/customers/**&quot;)

            .hasRole(&quot;ADMIN&quot;)

            .anyRequest()

            .authenticated();

    }

}
</code></pre>

<p>在上述代码中可以看到，我们使用了 HttpSecurity 类中的 antMatchers(&ldquo;/customer/<strong>&rdquo;) 和 hasRole(&ldquo;ADMIN&rdquo;) 方法为访问&rdquo;/customers/</strong>&ldquo;的请求限定了角色，只有&rdquo;ADMIN&rdquo;角色的认证用户才能访问以&rdquo;/customers/&ldquo;为根地址的所有 URL。</p>

<p>如果我们使用了认证用户+角色的方式保护服务访问，使用角色为“USER”的认证用户“springcss_user”访问 customer-service 时就会出现如下所示的“access_denied”错误信息：</p>

<pre><code>{

    &quot;error&quot;: &quot;access_denied&quot;,

    &quot;error_description&quot;: &quot;Access is denied&quot;

}
</code></pre>

<p>而我们使用具有“ADMIN”角色的“springcss_admin”用户访问 customer-service 时，将会得到正常的返回信息，关于这点你可以自己做一些尝试。</p>

<h4 id="使用用户-角色-操作级别保护服务访问">使用用户+角色+操作级别保护服务访问</h4>

<p>最后一种保护服务访问的策略粒度划分最细，在认证用户+角色的基础上，我们需要再对具体的 HTTP 操作进行限制。</p>

<p>在 customer-service 中，我们认为所有对客服工单的删除操作都很危险，因此可以使用 http.antMatchers(HttpMethod.DELETE, &ldquo;/customers/**&rdquo;) 方法对删除操作进行保护，示例代码如下：</p>

<pre><code>@Configuration

public class SpringCssSecurityConfig extends WebSecurityConfigurerAdapter {

 

    @Override

    public void configure(HttpSecurity http) throws Exception{

        http.authorizeRequests()

                .antMatchers(HttpMethod.DELETE, &quot;/customers/**&quot;)

                .hasRole(&quot;ADMIN&quot;)

                .anyRequest()

                .authenticated();

    }

}
</code></pre>

<p>上述代码的效果在于对“/customers”端点执行删除操作时，我们需要使用具有“ADMIN”角色的“springcss_admin”用户，执行其他操作时不需要。因为如果我们使用“springcss_user”账户执行删除操作，还是会出现“access_denied”错误信息。</p>

<h3 id="小结与预告">小结与预告</h3>

<p>通过 19 讲的学习，我们明确了 Web 应用程序中访问授权控制的三种粒度，并基于 SpringCSS 案例给出了三种粒度下的控制实现方式。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#d9b5b5b5e0ede8e8e9ee99beb4b8b0b5f7bab6b4" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1275db69d8ede4',t:'MTczNDA1NjI3My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>