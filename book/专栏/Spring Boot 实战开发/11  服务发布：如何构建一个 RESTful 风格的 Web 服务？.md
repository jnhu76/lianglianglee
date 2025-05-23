<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=11&#32;&#32;服务发布：如何构建一个&#32;RESTful&#32;风格的&#32;Web&#32;服务？>
        <link rel="icon" href="/static/favicon.png">
        <title>11  服务发布：如何构建一个 RESTful 风格的 Web 服务？ </title>
        
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
                            <h1 id="title" data-id="11  服务发布：如何构建一个 RESTful 风格的 Web 服务？" class="title">11  服务发布：如何构建一个 RESTful 风格的 Web 服务？</h1>
                            <div><p>通过前面课程的学习，我们已经掌握了构建一个 Spring Boot 应用程序的数据访问层组件实现方法。接下来的几讲，我们将讨论另一层组件，即 Web 服务层的构建方式。</p>

<p>服务与服务之间的交互是系统设计和发展的必然需求，其涉及 Web 服务的发布及消费，今天我们先讨论如何在 Spring Boot 应用程序中发布 Web 服务。</p>

<h3 id="springcss-系统中的服务交互">SpringCSS 系统中的服务交互</h3>

<p>在具体的技术体系介绍之前，我们先来梳理 SpringCSS 案例中服务交互之间的应用场景。</p>

<p>对于客服系统而言，其核心业务流程是生成客服工单，而工单的生成通常需要使用用户账户信息和所关联的订单信息。</p>

<p>在 SpringCSS 案例中，前面几讲我们已经构建了一个用于管理订单的 order-service，接下来我们将分别构建管理用户账户的 account-service 及核心的客服服务 customer-service。</p>

<p>关于三个服务之间的交互方式，我们先通过一张图了解下，如下图所示：</p>

<p><img src="assets/CgpVE1_hv2uAGbTPAABvOdgwGhs113.png" alt="图片6" /></p>

<p>SpringCSS 案例系统中三个服务的交互方式图</p>

<p>实际上，通过上图我们已经可以梳理工单生成 generateCustomerTicket 核心方法的执行流程，这里我们先给出代码的框架，如下代码所示：</p>

<pre><code>public CustomerTicket generateCustomerTicket(Long accountId, String orderNumber) {

        // 创建客服工单对象

        CustomerTicket customerTicket = new CustomerTicket();

 

        // 从远程 account-service 中获取 Account 对象

        Account account = getRemoteAccountById(accountId);

 

        // 从远程 order-service 中获取 Order 读写

        Order order = getRemoteOrderByOrderNumber(orderNumber);

        // 设置 CustomerTicket 对象并保存

        customerTicket.setAccountId(accountId);

        customerTicket.setOrderNumber(order.getOrderNumber());

        customerTicketRepository.save(customerTicket);

 

        return customerTicket;

}
</code></pre>

<p>因 getRemoteAccountById 与 getRemoteOrderByOrderNumber 方法都涉及远程 Web 服务的调用，因此首先我们需要创建 Web 服务。</p>

<p>而 Spring Boot 为我们创建 Web 服务提供了非常强大的组件化支持，简单而方便，我们一起来看一下。</p>

<h3 id="创建-restful-服务">创建 RESTful 服务</h3>

<p>在当下的分布式系统及微服务架构中，RESTful 风格是一种主流的 Web 服务表现方式。</p>

<p>在接下来的内容中，我们将演示如何使用 Spring Boot 创建 RESTful 服务。但在此之前，我们先来理解什么是 RESTful 服务。</p>

<h4 id="理解-restful-架构风格">理解 RESTful 架构风格</h4>

<p>你可能听说过 REST 这个名称，但并不清楚它的含义。</p>

<p>REST（Representational State Transfer，表述性状态转移）本质上只是一种架构风格而不是一种规范，这种架构风格把位于服务器端的访问入口看作一个资源，每个资源都使用 URI（Universal Resource Identifier，统一资源标识符） 得到一个唯一的地址，且在传输协议上使用标准的 HTTP 方法，比如最常见的 GET、PUT、POST 和 DELETE。</p>

<p>下表展示了 RESTful 风格的一些具体示例：</p>

<p><img src="assets/Cip5yF_hv3eAbVtLAAGhaTfDT5s664.png" alt="图片1" /></p>

<p>RESTful 风格示例</p>

<p>另一方面，客户端与服务器端的数据交互涉及序列化问题。关于序列化完成业务对象在网络环境上的传输的实现方式有很多，常见的有文本和二进制两大类。</p>

<p>目前 JSON 是一种被广泛采用的序列化方式，本课程中所有的代码实例我们都将 JSON 作为默认的序列化方式。</p>

<h4 id="使用基础注解">使用基础注解</h4>

<p>在原有 Spring Boot 应用程序的基础上，我们可以通过构建一系列的 Controller 类暴露 RESTful 风格的 HTTP 端点。这里的 Controller 与 Spring MVC 中的 Controller 概念上一致，最简单的 Controller 类如下代码所示：</p>

<pre><code>@RestController

public class HelloController {

 

    @GetMapping(&quot;/&quot;)

    public String index() {

        return &quot;Hello World!&quot;;

    }

}
</code></pre>

<p>从以上代码中可以看到，这里包含了 @RestController 和 @GetMapping 这两个注解。</p>

<p>其中，@RestController 注解继承自 Spring MVC 中的 @Controller 注解，顾名思义就是一个基于 RESTful 风格的 HTTP 端点，并且会自动使用 JSON 实现 HTTP 请求和响应的序列化/反序列化方式。</p>

<p>通过这个特性，在构建 RESTful 服务时，我们可以使用 @RestController 注解代替 @Controller 注解以简化开发。</p>

<p>另外一个 @GetMapping 注解也与 Spring MVC 中的 @RequestMapping 注解类似。我们先来看看 @RequestMapping 注解的定义，该注解所提供的属性都比较容易理解，如下代码所示：</p>

<pre><code>@Target({ElementType.METHOD, ElementType.TYPE})

@Retention(RetentionPolicy.RUNTIME)

@Documented

@Mapping

public @interface RequestMapping {

    String name() default &quot;&quot;;

 

    @AliasFor(&quot;path&quot;)

    String[] value() default {};

 

    @AliasFor(&quot;value&quot;)

    String[] path() default {};

 

    RequestMethod[] method() default {};

 

    String[] params() default {};

 

    String[] headers() default {};

 

    String[] consumes() default {};

    String[] produces() default {};

}
</code></pre>

<p>而 @GetMapping 的注解的定义与 @RequestMapping 非常类似，只是默认使用了 RequestMethod.GET 指定 HTTP 方法，如下代码所示：</p>

<pre><code>@Target(ElementType.METHOD)

@Retention(RetentionPolicy.RUNTIME)

@Documented

@RequestMapping(method = RequestMethod.GET)

public @interface GetMapping {
</code></pre>

<p>Spring Boot 2 中引入的一批新注解中，除了 @GetMapping ，还有 @PutMapping、@PostMapping、@DeleteMapping 等注解，这些注解极大方便了开发人员显式指定 HTTP 的请求方法。当然，你也可以继续使用原先的 @RequestMapping 实现同样的效果。</p>

<p>我们再看一个更加具体的示例，以下代码展示了 account-service 中的 AccountController。</p>

<pre><code>@RestController

@RequestMapping(value = &quot;accounts&quot;)

public class AccountController {

 

    @GetMapping(value = &quot;/{accountId}&quot;)

    public Account getAccountById(@PathVariable(&quot;accountId&quot;) Long accountId) {

        Account account = new Account();

        account.setId(1L);

        account.setAccountCode(&quot;DemoCode&quot;);

        account.setAccountName(&quot;DemoName&quot;);

        return account;

    }

}
</code></pre>

<p>在该 Controller 中，通过静态的业务代码我们完成了根据账号编号（accountId）获取用户账户信息的业务流程。</p>

<p>这里用到了两层 Mapping，第一层的 @RequestMapping 注解在服务层级定义了服务的根路径“/accounts”，第二层的 @GetMapping 注解则在操作级别定义了 HTTP 请求方法的具体路径及参数信息。</p>

<p>到这里，一个典型的 RESTful 服务已经开发完成了，现在我们可以通过 java –jar 命令直接运行 Spring Boot 应用程序了。</p>

<p>在启动日志中，我们发现了以下输出内容（为了显示效果，部分内容做了调整），可以看到自定义的这个 AccountController 已经成功启动并准备接收响应。</p>

<pre><code>RequestMappingHandlerMapping : Mapped &quot;{[/accounts/{accountId}], methods=[GET]}&quot; onto public com.springcss.account.domain.Account com.springcss.account.controller.AccountController.getAccountById (java.lang.Long)
</code></pre>

<p>在本课程中，我们将引入 Postman 来演示如何通过 HTTP 协议暴露的端点进行远程服务访问。</p>

<p>Postman 为我们完成 HTTP 请求和响应过程提供了可视化界面，你可以尝试编写一个 AccountController，并通过 Postman 访问“<a href="http://localhost:8082/accounts/1" target="_blank">http://localhost:8082/accounts/1</a>”端点以得到响应结果。</p>

<p>在前面的 AccountController 中，我们还看到了一个新的注解 @PathVariable，该注解作用于输入的参数，下面我们就来看看如何通过这些注解控制请求的输入。</p>

<h4 id="控制请求输入和输出">控制请求输入和输出</h4>

<p>Spring Boot 提供了一系列简单有用的注解来简化对请求输入的控制过程，常用的包括 @PathVariable、@RequestParam 和 @RequestBody。</p>

<p>其中 @PathVariable 注解用于获取路径参数，即从类似 url/{id} 这种形式的路径中获取 {id} 参数的值。该注解的定义如下代码所示：</p>

<pre><code>@Target(ElementType.PARAMETER)

@Retention(RetentionPolicy.RUNTIME)

@Documented

public @interface PathVariable {

    @AliasFor(&quot;name&quot;)

    String value() default &quot;&quot;;

 

    @AliasFor(&quot;value&quot;)

    String name() default &quot;&quot;;

 

    boolean required() default true;

}
</code></pre>

<p>通常，使用 @PathVariable 注解时，我们只需要指定一个参数的名称即可。我们可以再看一个示例，如下代码所示：</p>

<pre><code>@GetMapping(value = &quot;/{accountName}&quot;)

public Account getAccountByAccountName(@PathVariable(&quot;accountName&quot;) String accountName) {

 

    Account account = accountService.getAccountByAccountName(accountName);

    return account;

}
</code></pre>

<p>@RequestParam 注解的作用与 @PathVariable 注解类似，也是用于获取请求中的参数，但是它面向类似 url?id=XXX 这种路径形式。</p>

<p>该注解的定义如下代码所示，相较 @PathVariable 注解，它只是多了一个设置默认值的 defaultValue 属性。</p>

<pre><code>@Target(ElementType.PARAMETER)

@Retention(RetentionPolicy.RUNTIME)

@Documented

public @interface RequestParam {

    @AliasFor(&quot;name&quot;)

    String value() default &quot;&quot;;

 

    @AliasFor(&quot;value&quot;)

    String name() default &quot;&quot;;

 

    boolean required() default true;

 

    String defaultValue() default ValueConstants.DEFAULT_NONE;

}
</code></pre>

<p>在 HTTP 协议中，content-type 属性用来指定所传输的内容类型，我们可以通过 @RequestMapping 注解中的 produces 属性来设置这个属性。</p>

<p>在设置这个属性时，我们通常会将其设置为“application/json”，如下代码所示：</p>

<pre><code>@RestController

@RequestMapping(value = &quot;accounts&quot;, produces=&quot;application/json&quot;)

public class AccountController {
</code></pre>

<p>@RequestBody 注解用来处理 content-type 为 application/json 类型时的编码内容，通过 @RequestBody 注解可以将请求体中的 JSON 字符串绑定到相应的 JavaBean 上。</p>

<p>如下代码所示就是一个使用 @RequestBody 注解来控制输入的场景。</p>

<pre><code>@PutMapping(value = &quot;/&quot;)

public void updateAccount(@RequestBody Account account) {
</code></pre>

<p>如果使用 @RequestBody 注解，我们可以在 Postman 中输入一个 JSON 字符串来构建输入对象，如下代码所示：</p>

<p><img src="assets/Cip5yF_YgIGACIM7AAAr5-G8i7M764.png" alt="Drawing 1.png" /></p>

<p>使用 Postman 输入 JSON 字符串发起 HTTP 请求示例图</p>

<p>通过以上内容的讲解，我们发现使用注解的操作很简单，接下来我们有必要探讨下控制请求输入的规则。</p>

<p><strong>关于控制请求输入的规则，关键在于按照 RESTful 风格的设计原则设计 HTTP 端点，对于这点业界也存在一些约定。</strong></p>

<ul>
<li>以 Account 这个领域实体为例，如果我们把它视为一种资源，那么 HTTP 端点的根节点命名上通常采用复数形式，即“/accounts”，正如前面的示例代码所示。</li>
<li>在设计 RESTful API 时，我们需要基于 HTTP 语义设计对外暴露的端点的详细路径。针对常见的 CRUD 操作，我们展示了 RESTful API 与非 RESTful API 的一些区别。</li>
</ul>

<p><img src="assets/CgqCHl_hv5yAPnDHAADI8geMxRU064.png" alt="图片3" /></p>

<p>RESTful 风格对比示例</p>

<p>基于以上介绍的控制请求输入的实现方法，我们可以给出 account-service 中 AccountController 类的完整实现过程，如下代码所示：</p>

<pre><code>@RestController

@RequestMapping(value = &quot;accounts&quot;, produces=&quot;application/json&quot;)

public class AccountController {

 

    @Autowired

    private AccountService accountService;

 

    @GetMapping(value = &quot;/{accountId}&quot;)

    public Account getAccountById(@PathVariable(&quot;accountId&quot;) Long accountId) {

        Account account = accountService.getAccountById(accountId);

        return account;

    }

 

    @GetMapping(value = &quot;accountname/{accountName}&quot;)

    public Account getAccountByAccountName(@PathVariable(&quot;accountName&quot;) String accountName) {

 

        Account account = accountService.getAccountByAccountName(accountName);

        return account;

    }

 

    @PostMapping(value = &quot;/&quot;)

    public void addAccount(@RequestBody Account account) {

        accountService.addAccount(account);

    }

 

    @PutMapping(value = &quot;/&quot;)

    public void updateAccount(@RequestBody Account account) {

        accountService.updateAccount(account);

    }

    @DeleteMapping(value = &quot;/&quot;)

    public void deleteAccount(@RequestBody Account account) {

        accountService.deleteAccount(account);

    }

}
</code></pre>

<p>介绍完对请求输入的控制，我们再来讨论如何控制请求的输出。</p>

<p><strong>相较输入控制，输出控制就要简单很多，因为 Spring Boot 所提供的 @RestController 注解已经屏蔽了底层实现的复杂性，我们只需要返回一个普通的业务对象即可。</strong>@RestController 注解相当于是 Spring MVC 中 @Controller 和 @ResponseBody 这两个注解的组合，它们会自动返回 JSON 数据。</p>

<p>这里我们也给出了 order-service 中的 OrderController 实现过程，如下代码所示：</p>

<pre><code>@RestController

@RequestMapping(value=&quot;orders/jpa&quot;)

public class JpaOrderController {

    @Autowired

    JpaOrderService jpaOrderService;

    @GetMapping(value = &quot;/{orderId}&quot;)

    public JpaOrder getOrderById(@PathVariable Long orderId) {

        JpaOrder order = jpaOrderService.getOrderById(orderId);

     return order;

    }

    @GetMapping(value = &quot;orderNumber/{orderNumber}&quot;)

    public JpaOrder getOrderByOrderNumber(@PathVariable String orderNumber) {

        JpaOrder order = jpaOrderService.getOrderByOrderNumber(orderNumber);

//      JpaOrder order = jpaOrderService.getOrderByOrderNumberByExample(orderNumber);

//      JpaOrder order = jpaOrderService.getOrderByOrderNumberBySpecification(orderNumber);

     return order;

    }

    @PostMapping(value = &quot;&quot;)

    public JpaOrder addOrder(@RequestBody JpaOrder order) {

        JpaOrder result = jpaOrderService.addOrder(order);

     return result;

    }

}
</code></pre>

<p>从上面示例可以看到，我们使用了 09 讲中介绍的 Spring Data JPA 完成实体对象及数据访问功能。</p>

<h3 id="小结与预告">小结与预告</h3>

<p>构建 Web 服务是开发 Web 应用程序的基本需求，而设计并实现 RESTful 风格的 Web 服务是开发人员必须具备的开发技能。</p>

<p>基于 Spring Boot 框架，开发人员只需要使用几个注解就能实现复杂的 HTTP 端点，并暴露给其他服务进行使用，工作都变得非常简单。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#6d01010154595c5c5d5a2d0a000c0401430e0200" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1273adae96ede4',t:'MTczNDA1NjE4My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>