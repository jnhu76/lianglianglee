<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=38&#32;&#32;WebSocket：沙盒里的TCP>
        <link rel="icon" href="/static/favicon.png">
        <title>38  WebSocket：沙盒里的TCP </title>
        
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
                        <a class="menu-item" id="00 开篇词｜To Be a HTTP Hero.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/00%20%e5%bc%80%e7%af%87%e8%af%8d%ef%bd%9cTo%20Be%20a%20HTTP%20Hero.md">00 开篇词｜To Be a HTTP Hero.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01  时势与英雄：HTTP的前世今生.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/01%20%20%e6%97%b6%e5%8a%bf%e4%b8%8e%e8%8b%b1%e9%9b%84%ef%bc%9aHTTP%e7%9a%84%e5%89%8d%e4%b8%96%e4%bb%8a%e7%94%9f.md">01  时势与英雄：HTTP的前世今生.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02  HTTP是什么？HTTP又不是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/02%20%20HTTP%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9fHTTP%e5%8f%88%e4%b8%8d%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">02  HTTP是什么？HTTP又不是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03  HTTP世界全览（上）：与HTTP相关的各种概念.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/03%20%20HTTP%e4%b8%96%e7%95%8c%e5%85%a8%e8%a7%88%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e4%b8%8eHTTP%e7%9b%b8%e5%85%b3%e7%9a%84%e5%90%84%e7%a7%8d%e6%a6%82%e5%bf%b5.md">03  HTTP世界全览（上）：与HTTP相关的各种概念.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04  HTTP世界全览（下）：与HTTP相关的各种协议.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/04%20%20HTTP%e4%b8%96%e7%95%8c%e5%85%a8%e8%a7%88%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e4%b8%8eHTTP%e7%9b%b8%e5%85%b3%e7%9a%84%e5%90%84%e7%a7%8d%e5%8d%8f%e8%ae%ae.md">04  HTTP世界全览（下）：与HTTP相关的各种协议.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05  常说的“四层”和“七层”到底是什么？“五层”“六层”哪去了？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/05%20%20%e5%b8%b8%e8%af%b4%e7%9a%84%e2%80%9c%e5%9b%9b%e5%b1%82%e2%80%9d%e5%92%8c%e2%80%9c%e4%b8%83%e5%b1%82%e2%80%9d%e5%88%b0%e5%ba%95%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f%e2%80%9c%e4%ba%94%e5%b1%82%e2%80%9d%e2%80%9c%e5%85%ad%e5%b1%82%e2%80%9d%e5%93%aa%e5%8e%bb%e4%ba%86%ef%bc%9f.md">05  常说的“四层”和“七层”到底是什么？“五层”“六层”哪去了？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06  域名里有哪些门道？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/06%20%20%e5%9f%9f%e5%90%8d%e9%87%8c%e6%9c%89%e5%93%aa%e4%ba%9b%e9%97%a8%e9%81%93%ef%bc%9f.md">06  域名里有哪些门道？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07  自己动手，搭建HTTP实验环境.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/07%20%20%e8%87%aa%e5%b7%b1%e5%8a%a8%e6%89%8b%ef%bc%8c%e6%90%ad%e5%bb%baHTTP%e5%ae%9e%e9%aa%8c%e7%8e%af%e5%a2%83.md">07  自己动手，搭建HTTP实验环境.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08  键入网址再按下回车，后面究竟发生了什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/08%20%20%e9%94%ae%e5%85%a5%e7%bd%91%e5%9d%80%e5%86%8d%e6%8c%89%e4%b8%8b%e5%9b%9e%e8%bd%a6%ef%bc%8c%e5%90%8e%e9%9d%a2%e7%a9%b6%e7%ab%9f%e5%8f%91%e7%94%9f%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">08  键入网址再按下回车，后面究竟发生了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09  HTTP报文是什么样子的？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/09%20%20HTTP%e6%8a%a5%e6%96%87%e6%98%af%e4%bb%80%e4%b9%88%e6%a0%b7%e5%ad%90%e7%9a%84%ef%bc%9f.md">09  HTTP报文是什么样子的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10  应该如何理解请求方法？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/10%20%20%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e7%90%86%e8%a7%a3%e8%af%b7%e6%b1%82%e6%96%b9%e6%b3%95%ef%bc%9f.md">10  应该如何理解请求方法？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11  你能写出正确的网址吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/11%20%20%e4%bd%a0%e8%83%bd%e5%86%99%e5%87%ba%e6%ad%a3%e7%a1%ae%e7%9a%84%e7%bd%91%e5%9d%80%e5%90%97%ef%bc%9f.md">11  你能写出正确的网址吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12  响应状态码该怎么用？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/12%20%20%e5%93%8d%e5%ba%94%e7%8a%b6%e6%80%81%e7%a0%81%e8%af%a5%e6%80%8e%e4%b9%88%e7%94%a8%ef%bc%9f.md">12  响应状态码该怎么用？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13  HTTP有哪些特点？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/13%20%20HTTP%e6%9c%89%e5%93%aa%e4%ba%9b%e7%89%b9%e7%82%b9%ef%bc%9f.md">13  HTTP有哪些特点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14  HTTP有哪些优点？又有哪些缺点？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/14%20%20HTTP%e6%9c%89%e5%93%aa%e4%ba%9b%e4%bc%98%e7%82%b9%ef%bc%9f%e5%8f%88%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bc%ba%e7%82%b9%ef%bc%9f.md">14  HTTP有哪些优点？又有哪些缺点？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15  海纳百川：HTTP的实体数据.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/15%20%20%e6%b5%b7%e7%ba%b3%e7%99%be%e5%b7%9d%ef%bc%9aHTTP%e7%9a%84%e5%ae%9e%e4%bd%93%e6%95%b0%e6%8d%ae.md">15  海纳百川：HTTP的实体数据.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16  把大象装进冰箱：HTTP传输大文件的方法.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/16%20%20%e6%8a%8a%e5%a4%a7%e8%b1%a1%e8%a3%85%e8%bf%9b%e5%86%b0%e7%ae%b1%ef%bc%9aHTTP%e4%bc%a0%e8%be%93%e5%a4%a7%e6%96%87%e4%bb%b6%e7%9a%84%e6%96%b9%e6%b3%95.md">16  把大象装进冰箱：HTTP传输大文件的方法.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17  排队也要讲效率：HTTP的连接管理.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/17%20%20%e6%8e%92%e9%98%9f%e4%b9%9f%e8%a6%81%e8%ae%b2%e6%95%88%e7%8e%87%ef%bc%9aHTTP%e7%9a%84%e8%bf%9e%e6%8e%a5%e7%ae%a1%e7%90%86.md">17  排队也要讲效率：HTTP的连接管理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18  四通八达：HTTP的重定向和跳转.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/18%20%20%e5%9b%9b%e9%80%9a%e5%85%ab%e8%be%be%ef%bc%9aHTTP%e7%9a%84%e9%87%8d%e5%ae%9a%e5%90%91%e5%92%8c%e8%b7%b3%e8%bd%ac.md">18  四通八达：HTTP的重定向和跳转.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19  让我知道你是谁：HTTP的Cookie机制.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/19%20%20%e8%ae%a9%e6%88%91%e7%9f%a5%e9%81%93%e4%bd%a0%e6%98%af%e8%b0%81%ef%bc%9aHTTP%e7%9a%84Cookie%e6%9c%ba%e5%88%b6.md">19  让我知道你是谁：HTTP的Cookie机制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20  生鲜速递：HTTP的缓存控制.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/20%20%20%e7%94%9f%e9%b2%9c%e9%80%9f%e9%80%92%ef%bc%9aHTTP%e7%9a%84%e7%bc%93%e5%ad%98%e6%8e%a7%e5%88%b6.md">20  生鲜速递：HTTP的缓存控制.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21  良心中间商：HTTP的代理服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/21%20%20%e8%89%af%e5%bf%83%e4%b8%ad%e9%97%b4%e5%95%86%ef%bc%9aHTTP%e7%9a%84%e4%bb%a3%e7%90%86%e6%9c%8d%e5%8a%a1.md">21  良心中间商：HTTP的代理服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22  冷链周转：HTTP的缓存代理.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/22%20%20%e5%86%b7%e9%93%be%e5%91%a8%e8%bd%ac%ef%bc%9aHTTP%e7%9a%84%e7%bc%93%e5%ad%98%e4%bb%a3%e7%90%86.md">22  冷链周转：HTTP的缓存代理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23  HTTPS是什么？SSLTLS又是什么？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/23%20%20HTTPS%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9fSSLTLS%e5%8f%88%e6%98%af%e4%bb%80%e4%b9%88%ef%bc%9f.md">23  HTTPS是什么？SSLTLS又是什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24  固若金汤的根本（上）：对称加密与非对称加密.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/24%20%20%e5%9b%ba%e8%8b%a5%e9%87%91%e6%b1%a4%e7%9a%84%e6%a0%b9%e6%9c%ac%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%af%b9%e7%a7%b0%e5%8a%a0%e5%af%86%e4%b8%8e%e9%9d%9e%e5%af%b9%e7%a7%b0%e5%8a%a0%e5%af%86.md">24  固若金汤的根本（上）：对称加密与非对称加密.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25  固若金汤的根本（下）：数字签名与证书.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/25%20%20%e5%9b%ba%e8%8b%a5%e9%87%91%e6%b1%a4%e7%9a%84%e6%a0%b9%e6%9c%ac%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%95%b0%e5%ad%97%e7%ad%be%e5%90%8d%e4%b8%8e%e8%af%81%e4%b9%a6.md">25  固若金汤的根本（下）：数字签名与证书.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26  信任始于握手：TLS1.2连接过程解析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/26%20%20%e4%bf%a1%e4%bb%bb%e5%a7%8b%e4%ba%8e%e6%8f%a1%e6%89%8b%ef%bc%9aTLS1.2%e8%bf%9e%e6%8e%a5%e8%bf%87%e7%a8%8b%e8%a7%a3%e6%9e%90.md">26  信任始于握手：TLS1.2连接过程解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27  更好更快的握手：TLS1.3特性解析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/27%20%20%e6%9b%b4%e5%a5%bd%e6%9b%b4%e5%bf%ab%e7%9a%84%e6%8f%a1%e6%89%8b%ef%bc%9aTLS1.3%e7%89%b9%e6%80%a7%e8%a7%a3%e6%9e%90.md">27  更好更快的握手：TLS1.3特性解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28  连接太慢该怎么办：HTTPS的优化.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/28%20%20%e8%bf%9e%e6%8e%a5%e5%a4%aa%e6%85%a2%e8%af%a5%e6%80%8e%e4%b9%88%e5%8a%9e%ef%bc%9aHTTPS%e7%9a%84%e4%bc%98%e5%8c%96.md">28  连接太慢该怎么办：HTTPS的优化.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29  我应该迁移到HTTPS吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/29%20%20%e6%88%91%e5%ba%94%e8%af%a5%e8%bf%81%e7%a7%bb%e5%88%b0HTTPS%e5%90%97%ef%bc%9f.md">29  我应该迁移到HTTPS吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30  时代之风（上）：HTTP2特性概览.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/30%20%20%e6%97%b6%e4%bb%a3%e4%b9%8b%e9%a3%8e%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aHTTP2%e7%89%b9%e6%80%a7%e6%a6%82%e8%a7%88.md">30  时代之风（上）：HTTP2特性概览.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31  时代之风（下）：HTTP2内核剖析.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/31%20%20%e6%97%b6%e4%bb%a3%e4%b9%8b%e9%a3%8e%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aHTTP2%e5%86%85%e6%a0%b8%e5%89%96%e6%9e%90.md">31  时代之风（下）：HTTP2内核剖析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32  未来之路：HTTP3展望.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/32%20%20%e6%9c%aa%e6%9d%a5%e4%b9%8b%e8%b7%af%ef%bc%9aHTTP3%e5%b1%95%e6%9c%9b.md">32  未来之路：HTTP3展望.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33  我应该迁移到HTTP2吗？.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/33%20%20%e6%88%91%e5%ba%94%e8%af%a5%e8%bf%81%e7%a7%bb%e5%88%b0HTTP2%e5%90%97%ef%bc%9f.md">33  我应该迁移到HTTP2吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34  Nginx：高性能的Web服务器.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/34%20%20Nginx%ef%bc%9a%e9%ab%98%e6%80%a7%e8%83%bd%e7%9a%84Web%e6%9c%8d%e5%8a%a1%e5%99%a8.md">34  Nginx：高性能的Web服务器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35  OpenResty：更灵活的Web服务器.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/35%20%20OpenResty%ef%bc%9a%e6%9b%b4%e7%81%b5%e6%b4%bb%e7%9a%84Web%e6%9c%8d%e5%8a%a1%e5%99%a8.md">35  OpenResty：更灵活的Web服务器.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36  WAF：保护我们的网络服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/36%20%20WAF%ef%bc%9a%e4%bf%9d%e6%8a%a4%e6%88%91%e4%bb%ac%e7%9a%84%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">36  WAF：保护我们的网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="37  CDN：加速我们的网络服务.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/37%20%20CDN%ef%bc%9a%e5%8a%a0%e9%80%9f%e6%88%91%e4%bb%ac%e7%9a%84%e7%bd%91%e7%bb%9c%e6%9c%8d%e5%8a%a1.md">37  CDN：加速我们的网络服务.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="38  WebSocket：沙盒里的TCP.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/38%20%20WebSocket%ef%bc%9a%e6%b2%99%e7%9b%92%e9%87%8c%e7%9a%84TCP.md">38  WebSocket：沙盒里的TCP.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="39  HTTP性能优化面面观（上）.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/39%20%20HTTP%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e9%9d%a2%e9%9d%a2%e8%a7%82%ef%bc%88%e4%b8%8a%ef%bc%89.md">39  HTTP性能优化面面观（上）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="40  HTTP性能优化面面观（下）.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/40%20%20HTTP%e6%80%a7%e8%83%bd%e4%bc%98%e5%8c%96%e9%9d%a2%e9%9d%a2%e8%a7%82%ef%bc%88%e4%b8%8b%ef%bc%89.md">40  HTTP性能优化面面观（下）.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语  做兴趣使然的Hero.md" href="/%e4%b8%93%e6%a0%8f/%e9%80%8f%e8%a7%86HTTP%e5%8d%8f%e8%ae%ae/%e7%bb%93%e6%9d%9f%e8%af%ad%20%20%e5%81%9a%e5%85%b4%e8%b6%a3%e4%bd%bf%e7%84%b6%e7%9a%84Hero.md">结束语  做兴趣使然的Hero.md</a>
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
                            <h1 id="title" data-id="38  WebSocket：沙盒里的TCP" class="title">38  WebSocket：沙盒里的TCP</h1>
                            <div><p>在之前讲 TCP/IP 协议栈的时候，我说过有“TCP Socket”，它实际上是一种功能接口，通过这些接口就可以使用 TCP/IP 协议栈在传输层收发数据。</p>

<p>那么，你知道还有一种东西叫“WebSocket”吗？</p>

<p>单从名字上看，“Web”指的是 HTTP，“Socket”是套接字调用，那么这两个连起来又是什么意思呢？</p>

<p>所谓“望文生义”，大概你也能猜出来，“WebSocket”就是运行在“Web”，也就是 HTTP 上的 Socket 通信规范，提供与“TCP Socket”类似的功能，使用它就可以像“TCP Socket”一样调用下层协议栈，任意地收发数据。</p>

<p><img src="assets/ee6685c7d3c673b95e46d582828eee28.png" alt="img" /></p>

<p>更准确地说，“WebSocket”是一种基于 TCP 的轻量级网络通信协议，在地位上是与 HTTP“平级”的。</p>

<h2 id="为什么要有-websocket">为什么要有 WebSocket</h2>

<p>不过，已经有了被广泛应用的 HTTP 协议，为什么要再出一个 WebSocket 呢？它有哪些好处呢？</p>

<p>其实 WebSocket 与 HTTP/2 一样，都是为了解决 HTTP 某方面的缺陷而诞生的。HTTP/2 针对的是“队头阻塞”，而 WebSocket 针对的是“请求 - 应答”通信模式。</p>

<p>那么，“请求 - 应答”有什么不好的地方呢？</p>

<p>“请求 - 应答”是一种“<strong>半双工</strong>”的通信模式，虽然可以双向收发数据，但同一时刻只能一个方向上有动作，传输效率低。更关键的一点，它是一种“被动”通信模式，服务器只能“被动”响应客户端的请求，无法主动向客户端发送数据。</p>

<p>虽然后来的 HTTP/2、HTTP/3 新增了 Stream、Server Push 等特性，但“请求 - 应答”依然是主要的工作方式。这就导致 HTTP 难以应用在动态页面、即时消息、网络游戏等要求“<strong>实时通信</strong>”的领域。</p>

<p>在 WebSocket 出现之前，在浏览器环境里用 JavaScript 开发实时 Web 应用很麻烦。因为浏览器是一个“受限的沙盒”，不能用 TCP，只有 HTTP 协议可用，所以就出现了很多“变通”的技术，“<strong>轮询</strong>”（polling）就是比较常用的的一种。</p>

<p>简单地说，轮询就是不停地向服务器发送 HTTP 请求，问有没有数据，有数据的话服务器就用响应报文回应。如果轮询的频率比较高，那么就可以近似地实现“实时通信”的效果。</p>

<p>但轮询的缺点也很明显，反复发送无效查询请求耗费了大量的带宽和 CPU 资源，非常不经济。</p>

<p>所以，为了克服 HTTP“请求 - 应答”模式的缺点，WebSocket 就“应运而生”了。它原来是 HTML5 的一部分，后来“自立门户”，形成了一个单独的标准，RFC 文档编号是 6455。</p>

<h2 id="websocket-的特点">WebSocket 的特点</h2>

<p>WebSocket 是一个真正“<strong>全双工</strong>”的通信协议，与 TCP 一样，客户端和服务器都可以随时向对方发送数据，而不用像 HTTP“你拍一，我拍一”那么“客套”。于是，服务器就可以变得更加“主动”了。一旦后台有新的数据，就可以立即“推送”给客户端，不需要客户端轮询，“实时通信”的效率也就提高了。</p>

<p>WebSocket 采用了二进制帧结构，语法、语义与 HTTP 完全不兼容，但因为它的主要运行环境是浏览器，为了便于推广和应用，就不得不“搭便车”，在使用习惯上尽量向 HTTP 靠拢，这就是它名字里“Web”的含义。</p>

<p>服务发现方面，WebSocket 没有使用 TCP 的“IP 地址 + 端口号”，而是延用了 HTTP 的 URI 格式，但开头的协议名不是“http”，引入的是两个新的名字：“<strong>ws</strong>”和“<strong>wss</strong>”，分别表示明文和加密的 WebSocket 协议。</p>

<p>WebSocket 的默认端口也选择了 80 和 443，因为现在互联网上的防火墙屏蔽了绝大多数的端口，只对 HTTP 的 80、443 端口“放行”，所以 WebSocket 就可以“伪装”成 HTTP 协议，比较容易地“穿透”防火墙，与服务器建立连接。具体是怎么“伪装”的，我稍后再讲。</p>

<p>下面我举几个 WebSocket 服务的例子，你看看，是不是和 HTTP 几乎一模一样：</p>

<pre><code>ws://www.chrono.com
ws://www.chrono.com:8080/srv
wss://www.chrono.com:445/im?user_id=xxx
</code></pre>

<p>要注意的一点是，WebSocket 的名字容易让人产生误解，虽然大多数情况下我们会在浏览器里调用 API 来使用 WebSocket，但它不是一个“调用接口的集合”，而是一个通信协议，所以我觉得把它理解成“<strong>TCP over Web</strong>”会更恰当一些。</p>

<h2 id="websocket-的帧结构">WebSocket 的帧结构</h2>

<p>刚才说了，WebSocket 用的也是二进制帧，有之前 HTTP/2、HTTP/3 的经验，相信你这次也能很快掌握 WebSocket 的报文结构。</p>

<p>不过 WebSocket 和 HTTP/2 的关注点不同，WebSocket 更<strong>侧重于“实时通信”</strong>，而 HTTP/2 更侧重于提高传输效率，所以两者的帧结构也有很大的区别。</p>

<p>WebSocket 虽然有“帧”，但却没有像 HTTP/2 那样定义“流”，也就不存在“多路复用”“优先级”等复杂的特性，而它自身就是“全双工”的，也就不需要“服务器推送”。所以综合起来，WebSocket 的帧学习起来会简单一些。</p>

<p>下图就是 WebSocket 的帧结构定义，长度不固定，最少 2 个字节，最多 14 字节，看着好像很复杂，实际非常简单。</p>

<p><img src="assets/29d33e972dda5a27aa4773eea896a8c4.png" alt="img" /></p>

<p>开头的两个字节是必须的，也是最关键的。</p>

<p>第一个字节的第一位“<strong>FIN</strong>”是消息结束的标志位，相当于 HTTP/2 里的“END_STREAM”，表示数据发送完毕。一个消息可以拆成多个帧，接收方看到“FIN”后，就可以把前面的帧拼起来，组成完整的消息。</p>

<p>“FIN”后面的三个位是保留位，目前没有任何意义，但必须是 0。</p>

<p>第一个字节的后 4 位很重要，叫<strong>“Opcode</strong>”，操作码，其实就是帧类型，比如 1 表示帧内容是纯文本，2 表示帧内容是二进制数据，8 是关闭连接，9 和 10 分别是连接保活的 PING 和 PONG。</p>

<p>第二个字节第一位是掩码标志位“<strong>MASK</strong>”，表示帧内容是否使用异或操作（xor）做简单的加密。目前的 WebSocket 标准规定，客户端发送数据必须使用掩码，而服务器发送则必须不使用掩码。</p>

<p>第二个字节后 7 位是“<strong>Payload len</strong>”，表示帧内容的长度。它是另一种变长编码，最少 7 位，最多是 7+64 位，也就是额外增加 8 个字节，所以一个 WebSocket 帧最大是 2^64。</p>

<p>长度字段后面是“<strong>Masking-key</strong>”，掩码密钥，它是由上面的标志位“MASK”决定的，如果使用掩码就是 4 个字节的随机数，否则就不存在。</p>

<p>这么分析下来，其实 WebSocket 的帧头就四个部分：“<strong>结束标志位 + 操作码 + 帧长度 + 掩码</strong>”，只是使用了变长编码的“小花招”，不像 HTTP/2 定长报文头那么简单明了。</p>

<p>我们的实验环境利用 OpenResty 的“lua-resty-websocket”库，实现了一个简单的 WebSocket 通信，你可以访问 URI“/38-1”，它会连接后端的 WebSocket 服务“ws://127.0.0.<sup>1</sup>&frasl;<sub>38</sub>-0”，用 Wireshark 抓包就可以看到 WebSocket 的整个通信过程。</p>

<p>下面的截图是其中的一个文本帧，因为它是客户端发出的，所以需要掩码，报文头就在两个字节之外多了四个字节的“Masking-key”，总共是 6 个字节。</p>

<p><img src="assets/c91ee4815097f5f9059ab798bb841594.png" alt="img" /></p>

<p>而报文内容经过掩码，不是直接可见的明文，但掩码的安全强度几乎是零，用“Masking-key”简单地异或一下就可以转换出明文。</p>

<h2 id="websocket-的握手">WebSocket 的握手</h2>

<p>和 TCP、TLS 一样，WebSocket 也要有一个握手过程，然后才能正式收发数据。</p>

<p>这里它还是搭上了 HTTP 的“便车”，利用了 HTTP 本身的“协议升级”特性，“伪装”成 HTTP，这样就能绕过浏览器沙盒、网络防火墙等等限制，这也是 WebSocket 与 HTTP 的另一个重要关联点。</p>

<p>WebSocket 的握手是一个标准的 HTTP GET 请求，但要带上两个协议升级的专用头字段：</p>

<ul>
<li>“Connection: Upgrade”，表示要求协议“升级”；</li>
<li>“Upgrade: websocket”，表示要“升级”成 WebSocket 协议。</li>
</ul>

<p>另外，为了防止普通的 HTTP 消息被“意外”识别成 WebSocket，握手消息还增加了两个额外的认证用头字段（所谓的“挑战”，Challenge）：</p>

<ul>
<li>Sec-WebSocket-Key：一个 Base64 编码的 16 字节随机数，作为简单的认证密钥；</li>
<li>Sec-WebSocket-Version：协议的版本号，当前必须是 13。</li>
</ul>

<p><img src="assets/8f007bb0e403b6cc28493565f709c997.png" alt="img" /></p>

<p>服务器收到 HTTP 请求报文，看到上面的四个字段，就知道这不是一个普通的 GET 请求，而是 WebSocket 的升级请求，于是就不走普通的 HTTP 处理流程，而是构造一个特殊的“101 Switching Protocols”响应报文，通知客户端，接下来就不用 HTTP 了，全改用 WebSocket 协议通信。（有点像 TLS 的“Change Cipher Spec”）</p>

<p>WebSocket 的握手响应报文也是有特殊格式的，要用字段“Sec-WebSocket-Accept”验证客户端请求报文，同样也是为了防止误连接。</p>

<p>具体的做法是把请求头里“Sec-WebSocket-Key”的值，加上一个专用的 UUID “258EAFA5-E914-47DA-95CA-C5AB0DC85B11”，再计算 SHA-1 摘要。</p>

<pre><code>encode_base64(
  sha1( 
    Sec-WebSocket-Key + '258EAFA5-E914-47DA-95CA-C5AB0DC85B11' ))
</code></pre>

<p>客户端收到响应报文，就可以用同样的算法，比对值是否相等，如果相等，就说明返回的报文确实是刚才握手时连接的服务器，认证成功。</p>

<p>握手完成，后续传输的数据就不再是 HTTP 报文，而是 WebSocket 格式的二进制帧了。</p>

<p><img src="assets/84e9fa337f2b4c2c9f14760feb41c903.png" alt="img" /></p>

<h2 id="小结">小结</h2>

<p>浏览器是一个“沙盒”环境，有很多的限制，不允许建立 TCP 连接收发数据，而有了 WebSocket，我们就可以在浏览器里与服务器直接建立“TCP 连接”，获得更多的自由。</p>

<p>不过自由也是有代价的，WebSocket 虽然是在应用层，但使用方式却与“TCP Socket”差不多，过于“原始”，用户必须自己管理连接、缓存、状态，开发上比 HTTP 复杂的多，所以是否要在项目中引入 WebSocket 必须慎重考虑。</p>

<ol>
<li>HTTP 的“请求 - 应答”模式不适合开发“实时通信”应用，效率低，难以实现动态页面，所以出现了 WebSocket；</li>
<li>WebSocket 是一个“全双工”的通信协议，相当于对 TCP 做了一层“薄薄的包装”，让它运行在浏览器环境里；</li>
<li>WebSocket 使用兼容 HTTP 的 URI 来发现服务，但定义了新的协议名“ws”和“wss”，端口号也沿用了 80 和 443；</li>
<li>WebSocket 使用二进制帧，结构比较简单，特殊的地方是有个“掩码”操作，客户端发数据必须掩码，服务器则不用；</li>
<li>WebSocket 利用 HTTP 协议实现连接握手，发送 GET 请求要求“协议升级”，握手过程中有个非常简单的认证机制，目的是防止误连接。</li>
</ol>

<h2 id="课下作业">课下作业</h2>

<ol>
<li>WebSocket 与 HTTP/2 有很多相似点，比如都可以从 HTTP/1 升级，都采用二进制帧结构，你能比较一下这两个协议吗？</li>
<li>试着自己解释一下 WebSocket 里的”Web“和”Socket“的含义。</li>
<li>结合自己的实际工作，你觉得 WebSocket 适合用在哪些场景里？</li>
</ol>

<p>欢迎你把自己的学习体会写在留言区，与我和其他同学一起讨论。如果你觉得有所收获，也欢迎把文章分享给你的朋友。</p>

<p><img src="assets/4b81de6b5c57db92ed7808344482ef5b.png" alt="unpreview" /></p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#650909095c5154545552250208040c094b060a08" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f1a003f99b7888b',t:'MTczNDEzNTM0Mi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>