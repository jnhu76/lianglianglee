<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=加餐5&#32;&#32;安全新技术：IoT、IPv6、区块链中的安全新问题>
        <link rel="icon" href="/static/favicon.png">
        <title>加餐5  安全新技术：IoT、IPv6、区块链中的安全新问题 </title>
        
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
                        <a class="menu-item" id="00 开篇词 别说你没被安全困扰过.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e5%88%ab%e8%af%b4%e4%bd%a0%e6%b2%a1%e8%a2%ab%e5%ae%89%e5%85%a8%e5%9b%b0%e6%89%b0%e8%bf%87.md">00 开篇词 别说你没被安全困扰过.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 安全的本质：数据被窃取后，你能意识到问题来源吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/01%20%e5%ae%89%e5%85%a8%e7%9a%84%e6%9c%ac%e8%b4%a8%ef%bc%9a%e6%95%b0%e6%8d%ae%e8%a2%ab%e7%aa%83%e5%8f%96%e5%90%8e%ef%bc%8c%e4%bd%a0%e8%83%bd%e6%84%8f%e8%af%86%e5%88%b0%e9%97%ae%e9%a2%98%e6%9d%a5%e6%ba%90%e5%90%97%ef%bc%9f.md">01 安全的本质：数据被窃取后，你能意识到问题来源吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 安全原则：我们应该如何上手解决安全问题？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/02%20%e5%ae%89%e5%85%a8%e5%8e%9f%e5%88%99%ef%bc%9a%e6%88%91%e4%bb%ac%e5%ba%94%e8%af%a5%e5%a6%82%e4%bd%95%e4%b8%8a%e6%89%8b%e8%a7%a3%e5%86%b3%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98%ef%bc%9f.md">02 安全原则：我们应该如何上手解决安全问题？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 密码学基础：如何让你的密码变得“不可见”？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/03%20%e5%af%86%e7%a0%81%e5%ad%a6%e5%9f%ba%e7%a1%80%ef%bc%9a%e5%a6%82%e4%bd%95%e8%ae%a9%e4%bd%a0%e7%9a%84%e5%af%86%e7%a0%81%e5%8f%98%e5%be%97%e2%80%9c%e4%b8%8d%e5%8f%af%e8%a7%81%e2%80%9d%ef%bc%9f.md">03 密码学基础：如何让你的密码变得“不可见”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 身份认证：除了账号密码，我们还能怎么做身份认证？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/04%20%e8%ba%ab%e4%bb%bd%e8%ae%a4%e8%af%81%ef%bc%9a%e9%99%a4%e4%ba%86%e8%b4%a6%e5%8f%b7%e5%af%86%e7%a0%81%ef%bc%8c%e6%88%91%e4%bb%ac%e8%bf%98%e8%83%bd%e6%80%8e%e4%b9%88%e5%81%9a%e8%ba%ab%e4%bb%bd%e8%ae%a4%e8%af%81%ef%bc%9f.md">04 身份认证：除了账号密码，我们还能怎么做身份认证？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 访问控制：如何选取一个合适的数据保护方案？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/05%20%e8%ae%bf%e9%97%ae%e6%8e%a7%e5%88%b6%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%89%e5%8f%96%e4%b8%80%e4%b8%aa%e5%90%88%e9%80%82%e7%9a%84%e6%95%b0%e6%8d%ae%e4%bf%9d%e6%8a%a4%e6%96%b9%e6%a1%88%ef%bc%9f.md">05 访问控制：如何选取一个合适的数据保护方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 XSS：当你“被发送”了一条微博时，到底发生了什么？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/06%20XSS%ef%bc%9a%e5%bd%93%e4%bd%a0%e2%80%9c%e8%a2%ab%e5%8f%91%e9%80%81%e2%80%9d%e4%ba%86%e4%b8%80%e6%9d%a1%e5%be%ae%e5%8d%9a%e6%97%b6%ef%bc%8c%e5%88%b0%e5%ba%95%e5%8f%91%e7%94%9f%e4%ba%86%e4%bb%80%e4%b9%88%ef%bc%9f.md">06 XSS：当你“被发送”了一条微博时，到底发生了什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 SQL注入：明明设置了强密码，为什么还会被别人登录？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/07%20SQL%e6%b3%a8%e5%85%a5%ef%bc%9a%e6%98%8e%e6%98%8e%e8%ae%be%e7%bd%ae%e4%ba%86%e5%bc%ba%e5%af%86%e7%a0%81%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e8%bf%98%e4%bc%9a%e8%a2%ab%e5%88%ab%e4%ba%ba%e7%99%bb%e5%bd%95%ef%bc%9f.md">07 SQL注入：明明设置了强密码，为什么还会被别人登录？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 CSRF_SSRF：为什么避免了XSS，还是“被发送”了一条微博？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/08%20CSRF_SSRF%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e9%81%bf%e5%85%8d%e4%ba%86XSS%ef%bc%8c%e8%bf%98%e6%98%af%e2%80%9c%e8%a2%ab%e5%8f%91%e9%80%81%e2%80%9d%e4%ba%86%e4%b8%80%e6%9d%a1%e5%be%ae%e5%8d%9a%ef%bc%9f.md">08 CSRF_SSRF：为什么避免了XSS，还是“被发送”了一条微博？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 反序列化漏洞：使用了编译型语言，为什么还是会被注入？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/09%20%e5%8f%8d%e5%ba%8f%e5%88%97%e5%8c%96%e6%bc%8f%e6%b4%9e%ef%bc%9a%e4%bd%bf%e7%94%a8%e4%ba%86%e7%bc%96%e8%af%91%e5%9e%8b%e8%af%ad%e8%a8%80%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e8%bf%98%e6%98%af%e4%bc%9a%e8%a2%ab%e6%b3%a8%e5%85%a5%ef%bc%9f.md">09 反序列化漏洞：使用了编译型语言，为什么还是会被注入？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 信息泄露：为什么黑客会知道你的代码逻辑？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/10%20%e4%bf%a1%e6%81%af%e6%b3%84%e9%9c%b2%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e9%bb%91%e5%ae%a2%e4%bc%9a%e7%9f%a5%e9%81%93%e4%bd%a0%e7%9a%84%e4%bb%a3%e7%a0%81%e9%80%bb%e8%be%91%ef%bc%9f.md">10 信息泄露：为什么黑客会知道你的代码逻辑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 插件漏洞：我的代码看起来很安全，为什么还会出现漏洞？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/11%20%e6%8f%92%e4%bb%b6%e6%bc%8f%e6%b4%9e%ef%bc%9a%e6%88%91%e7%9a%84%e4%bb%a3%e7%a0%81%e7%9c%8b%e8%b5%b7%e6%9d%a5%e5%be%88%e5%ae%89%e5%85%a8%ef%bc%8c%e4%b8%ba%e4%bb%80%e4%b9%88%e8%bf%98%e4%bc%9a%e5%87%ba%e7%8e%b0%e6%bc%8f%e6%b4%9e%ef%bc%9f.md">11 插件漏洞：我的代码看起来很安全，为什么还会出现漏洞？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 Linux系统安全：多人共用服务器，如何防止别人干“坏事”？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/13%20Linux%e7%b3%bb%e7%bb%9f%e5%ae%89%e5%85%a8%ef%bc%9a%e5%a4%9a%e4%ba%ba%e5%85%b1%e7%94%a8%e6%9c%8d%e5%8a%a1%e5%99%a8%ef%bc%8c%e5%a6%82%e4%bd%95%e9%98%b2%e6%ad%a2%e5%88%ab%e4%ba%ba%e5%b9%b2%e2%80%9c%e5%9d%8f%e4%ba%8b%e2%80%9d%ef%bc%9f.md">13 Linux系统安全：多人共用服务器，如何防止别人干“坏事”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 网络安全：和别人共用Wi-Fi时，你的信息会被窃取吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/14%20%e7%bd%91%e7%bb%9c%e5%ae%89%e5%85%a8%ef%bc%9a%e5%92%8c%e5%88%ab%e4%ba%ba%e5%85%b1%e7%94%a8Wi-Fi%e6%97%b6%ef%bc%8c%e4%bd%a0%e7%9a%84%e4%bf%a1%e6%81%af%e4%bc%9a%e8%a2%ab%e7%aa%83%e5%8f%96%e5%90%97%ef%bc%9f.md">14 网络安全：和别人共用Wi-Fi时，你的信息会被窃取吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 Docker安全：在虚拟的环境中，就不用考虑安全了吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/15%20Docker%e5%ae%89%e5%85%a8%ef%bc%9a%e5%9c%a8%e8%99%9a%e6%8b%9f%e7%9a%84%e7%8e%af%e5%a2%83%e4%b8%ad%ef%bc%8c%e5%b0%b1%e4%b8%8d%e7%94%a8%e8%80%83%e8%99%91%e5%ae%89%e5%85%a8%e4%ba%86%e5%90%97%ef%bc%9f.md">15 Docker安全：在虚拟的环境中，就不用考虑安全了吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 数据库安全：数据库中的数据是如何被黑客拖取的？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/16%20%e6%95%b0%e6%8d%ae%e5%ba%93%e5%ae%89%e5%85%a8%ef%bc%9a%e6%95%b0%e6%8d%ae%e5%ba%93%e4%b8%ad%e7%9a%84%e6%95%b0%e6%8d%ae%e6%98%af%e5%a6%82%e4%bd%95%e8%a2%ab%e9%bb%91%e5%ae%a2%e6%8b%96%e5%8f%96%e7%9a%84%ef%bc%9f.md">16 数据库安全：数据库中的数据是如何被黑客拖取的？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 分布式安全：上百个分布式节点，不会出现“内奸”吗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/17%20%e5%88%86%e5%b8%83%e5%bc%8f%e5%ae%89%e5%85%a8%ef%bc%9a%e4%b8%8a%e7%99%be%e4%b8%aa%e5%88%86%e5%b8%83%e5%bc%8f%e8%8a%82%e7%82%b9%ef%bc%8c%e4%b8%8d%e4%bc%9a%e5%87%ba%e7%8e%b0%e2%80%9c%e5%86%85%e5%a5%b8%e2%80%9d%e5%90%97%ef%bc%9f.md">17 分布式安全：上百个分布式节点，不会出现“内奸”吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="18 安全标准和框架：怎样依“葫芦”画出好“瓢”？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/18%20%e5%ae%89%e5%85%a8%e6%a0%87%e5%87%86%e5%92%8c%e6%a1%86%e6%9e%b6%ef%bc%9a%e6%80%8e%e6%a0%b7%e4%be%9d%e2%80%9c%e8%91%ab%e8%8a%a6%e2%80%9d%e7%94%bb%e5%87%ba%e5%a5%bd%e2%80%9c%e7%93%a2%e2%80%9d%ef%bc%9f.md">18 安全标准和框架：怎样依“葫芦”画出好“瓢”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 防火墙：如何和黑客“划清界限”？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/19%20%e9%98%b2%e7%81%ab%e5%a2%99%ef%bc%9a%e5%a6%82%e4%bd%95%e5%92%8c%e9%bb%91%e5%ae%a2%e2%80%9c%e5%88%92%e6%b8%85%e7%95%8c%e9%99%90%e2%80%9d%ef%bc%9f.md">19 防火墙：如何和黑客“划清界限”？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 WAF：如何为漏洞百出的Web应用保驾护航？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/20%20WAF%ef%bc%9a%e5%a6%82%e4%bd%95%e4%b8%ba%e6%bc%8f%e6%b4%9e%e7%99%be%e5%87%ba%e7%9a%84Web%e5%ba%94%e7%94%a8%e4%bf%9d%e9%a9%be%e6%8a%a4%e8%88%aa%ef%bc%9f.md">20 WAF：如何为漏洞百出的Web应用保驾护航？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 IDS：当黑客绕过了防火墙，你该如何发现？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/21%20IDS%ef%bc%9a%e5%bd%93%e9%bb%91%e5%ae%a2%e7%bb%95%e8%bf%87%e4%ba%86%e9%98%b2%e7%81%ab%e5%a2%99%ef%bc%8c%e4%bd%a0%e8%af%a5%e5%a6%82%e4%bd%95%e5%8f%91%e7%8e%b0%ef%bc%9f.md">21 IDS：当黑客绕过了防火墙，你该如何发现？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 RASP：写规则写得烦了？尝试一下更底层的IDS.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/22%20RASP%ef%bc%9a%e5%86%99%e8%a7%84%e5%88%99%e5%86%99%e5%be%97%e7%83%a6%e4%ba%86%ef%bc%9f%e5%b0%9d%e8%af%95%e4%b8%80%e4%b8%8b%e6%9b%b4%e5%ba%95%e5%b1%82%e7%9a%84IDS.md">22 RASP：写规则写得烦了？尝试一下更底层的IDS.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 SIEM：一个人管理好几个安全工具，如何高效运营？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/23%20SIEM%ef%bc%9a%e4%b8%80%e4%b8%aa%e4%ba%ba%e7%ae%a1%e7%90%86%e5%a5%bd%e5%87%a0%e4%b8%aa%e5%ae%89%e5%85%a8%e5%b7%a5%e5%85%b7%ef%bc%8c%e5%a6%82%e4%bd%95%e9%ab%98%e6%95%88%e8%bf%90%e8%90%a5%ef%bc%9f.md">23 SIEM：一个人管理好几个安全工具，如何高效运营？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 SDL：怎样才能写出更“安全”的代码？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/24%20SDL%ef%bc%9a%e6%80%8e%e6%a0%b7%e6%89%8d%e8%83%bd%e5%86%99%e5%87%ba%e6%9b%b4%e2%80%9c%e5%ae%89%e5%85%a8%e2%80%9d%e7%9a%84%e4%bb%a3%e7%a0%81%ef%bc%9f.md">24 SDL：怎样才能写出更“安全”的代码？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 业务安全体系：对比基础安全，业务安全有哪些不同？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/25%20%e4%b8%9a%e5%8a%a1%e5%ae%89%e5%85%a8%e4%bd%93%e7%b3%bb%ef%bc%9a%e5%af%b9%e6%af%94%e5%9f%ba%e7%a1%80%e5%ae%89%e5%85%a8%ef%bc%8c%e4%b8%9a%e5%8a%a1%e5%ae%89%e5%85%a8%e6%9c%89%e5%93%aa%e4%ba%9b%e4%b8%8d%e5%90%8c%ef%bc%9f.md">25 业务安全体系：对比基础安全，业务安全有哪些不同？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 产品安全方案：如何降低业务对黑灰产的诱惑？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/26%20%e4%ba%a7%e5%93%81%e5%ae%89%e5%85%a8%e6%96%b9%e6%a1%88%ef%bc%9a%e5%a6%82%e4%bd%95%e9%99%8d%e4%bd%8e%e4%b8%9a%e5%8a%a1%e5%af%b9%e9%bb%91%e7%81%b0%e4%ba%a7%e7%9a%84%e8%af%b1%e6%83%91%ef%bc%9f.md">26 产品安全方案：如何降低业务对黑灰产的诱惑？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 风控系统：如何从海量业务数据中，挖掘黑灰产？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/27%20%e9%a3%8e%e6%8e%a7%e7%b3%bb%e7%bb%9f%ef%bc%9a%e5%a6%82%e4%bd%95%e4%bb%8e%e6%b5%b7%e9%87%8f%e4%b8%9a%e5%8a%a1%e6%95%b0%e6%8d%ae%e4%b8%ad%ef%bc%8c%e6%8c%96%e6%8e%98%e9%bb%91%e7%81%b0%e4%ba%a7%ef%bc%9f.md">27 风控系统：如何从海量业务数据中，挖掘黑灰产？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 机器学习：如何教会机器识别黑灰产？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/28%20%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0%ef%bc%9a%e5%a6%82%e4%bd%95%e6%95%99%e4%bc%9a%e6%9c%ba%e5%99%a8%e8%af%86%e5%88%ab%e9%bb%91%e7%81%b0%e4%ba%a7%ef%bc%9f.md">28 机器学习：如何教会机器识别黑灰产？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 设备指纹：面对各种虚拟设备，如何进行对抗？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/29%20%e8%ae%be%e5%a4%87%e6%8c%87%e7%ba%b9%ef%bc%9a%e9%9d%a2%e5%af%b9%e5%90%84%e7%a7%8d%e8%99%9a%e6%8b%9f%e8%ae%be%e5%a4%87%ef%bc%8c%e5%a6%82%e4%bd%95%e8%bf%9b%e8%a1%8c%e5%af%b9%e6%8a%97%ef%bc%9f.md">29 设备指纹：面对各种虚拟设备，如何进行对抗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 安全运营：“黑灰产”打了又来，如何正确处置？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/30%20%e5%ae%89%e5%85%a8%e8%bf%90%e8%90%a5%ef%bc%9a%e2%80%9c%e9%bb%91%e7%81%b0%e4%ba%a7%e2%80%9d%e6%89%93%e4%ba%86%e5%8f%88%e6%9d%a5%ef%bc%8c%e5%a6%82%e4%bd%95%e6%ad%a3%e7%a1%ae%e5%a4%84%e7%bd%ae%ef%bc%9f.md">30 安全运营：“黑灰产”打了又来，如何正确处置？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐1 数据安全：如何防止内部员工泄露商业机密？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e5%8a%a0%e9%a4%901%20%e6%95%b0%e6%8d%ae%e5%ae%89%e5%85%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e9%98%b2%e6%ad%a2%e5%86%85%e9%83%a8%e5%91%98%e5%b7%a5%e6%b3%84%e9%9c%b2%e5%95%86%e4%b8%9a%e6%9c%ba%e5%af%86%ef%bc%9f.md">加餐1 数据安全：如何防止内部员工泄露商业机密？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐2 前端安全：如何打造一个可信的前端环境？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e5%8a%a0%e9%a4%902%20%e5%89%8d%e7%ab%af%e5%ae%89%e5%85%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e6%89%93%e9%80%a0%e4%b8%80%e4%b8%aa%e5%8f%af%e4%bf%a1%e7%9a%84%e5%89%8d%e7%ab%af%e7%8e%af%e5%a2%83%ef%bc%9f.md">加餐2 前端安全：如何打造一个可信的前端环境？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐3 职业发展：应聘安全工程师，我需要注意什么？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e5%8a%a0%e9%a4%903%20%e8%81%8c%e4%b8%9a%e5%8f%91%e5%b1%95%ef%bc%9a%e5%ba%94%e8%81%98%e5%ae%89%e5%85%a8%e5%b7%a5%e7%a8%8b%e5%b8%88%ef%bc%8c%e6%88%91%e9%9c%80%e8%a6%81%e6%b3%a8%e6%84%8f%e4%bb%80%e4%b9%88%ef%bc%9f.md">加餐3 职业发展：应聘安全工程师，我需要注意什么？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐4 个人成长：学习安全，哪些资源我必须要知道？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e5%8a%a0%e9%a4%904%20%e4%b8%aa%e4%ba%ba%e6%88%90%e9%95%bf%ef%bc%9a%e5%ad%a6%e4%b9%a0%e5%ae%89%e5%85%a8%ef%bc%8c%e5%93%aa%e4%ba%9b%e8%b5%84%e6%ba%90%e6%88%91%e5%bf%85%e9%a1%bb%e8%a6%81%e7%9f%a5%e9%81%93%ef%bc%9f.md">加餐4 个人成长：学习安全，哪些资源我必须要知道？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="加餐5  安全新技术：IoT、IPv6、区块链中的安全新问题.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e5%8a%a0%e9%a4%905%20%20%e5%ae%89%e5%85%a8%e6%96%b0%e6%8a%80%e6%9c%af%ef%bc%9aIoT%e3%80%81IPv6%e3%80%81%e5%8c%ba%e5%9d%97%e9%93%be%e4%b8%ad%e7%9a%84%e5%ae%89%e5%85%a8%e6%96%b0%e9%97%ae%e9%a2%98.md">加餐5  安全新技术：IoT、IPv6、区块链中的安全新问题.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="模块串讲（一）Web安全：如何评估用户数据和资产数据面临的威胁？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e6%a8%a1%e5%9d%97%e4%b8%b2%e8%ae%b2%ef%bc%88%e4%b8%80%ef%bc%89Web%e5%ae%89%e5%85%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e8%af%84%e4%bc%b0%e7%94%a8%e6%88%b7%e6%95%b0%e6%8d%ae%e5%92%8c%e8%b5%84%e4%ba%a7%e6%95%b0%e6%8d%ae%e9%9d%a2%e4%b8%b4%e7%9a%84%e5%a8%81%e8%83%81%ef%bc%9f.md">模块串讲（一）Web安全：如何评估用户数据和资产数据面临的威胁？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="模块串讲（三）安全防御工具：如何选择和规划公司的安全防御体系？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e6%a8%a1%e5%9d%97%e4%b8%b2%e8%ae%b2%ef%bc%88%e4%b8%89%ef%bc%89%e5%ae%89%e5%85%a8%e9%98%b2%e5%be%a1%e5%b7%a5%e5%85%b7%ef%bc%9a%e5%a6%82%e4%bd%95%e9%80%89%e6%8b%a9%e5%92%8c%e8%a7%84%e5%88%92%e5%85%ac%e5%8f%b8%e7%9a%84%e5%ae%89%e5%85%a8%e9%98%b2%e5%be%a1%e4%bd%93%e7%b3%bb%ef%bc%9f.md">模块串讲（三）安全防御工具：如何选择和规划公司的安全防御体系？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="模块串讲（二）Linux系统和应用安全：如何大范围提高平台安全性？.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e6%a8%a1%e5%9d%97%e4%b8%b2%e8%ae%b2%ef%bc%88%e4%ba%8c%ef%bc%89Linux%e7%b3%bb%e7%bb%9f%e5%92%8c%e5%ba%94%e7%94%a8%e5%ae%89%e5%85%a8%ef%bc%9a%e5%a6%82%e4%bd%95%e5%a4%a7%e8%8c%83%e5%9b%b4%e6%8f%90%e9%ab%98%e5%b9%b3%e5%8f%b0%e5%ae%89%e5%85%a8%e6%80%a7%ef%bc%9f.md">模块串讲（二）Linux系统和应用安全：如何大范围提高平台安全性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 在与黑客的战役中，我们都是盟友！.md" href="/%e4%b8%93%e6%a0%8f/%e5%ae%89%e5%85%a8%e6%94%bb%e9%98%b2%e6%8a%80%e8%83%bd30%e8%ae%b2/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e5%9c%a8%e4%b8%8e%e9%bb%91%e5%ae%a2%e7%9a%84%e6%88%98%e5%bd%b9%e4%b8%ad%ef%bc%8c%e6%88%91%e4%bb%ac%e9%83%bd%e6%98%af%e7%9b%9f%e5%8f%8b%ef%bc%81.md">结束语 在与黑客的战役中，我们都是盟友！.md</a>
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
                            <h1 id="title" data-id="加餐5  安全新技术：IoT、IPv6、区块链中的安全新问题" class="title">加餐5  安全新技术：IoT、IPv6、区块链中的安全新问题</h1>
                            <div><p>加餐5 安全新技术：IoT、IPv6、区块链中的安全新问题</p>

<p>你好，我是何为舟。欢迎来到安全专栏的第5次加餐时间。</p>

<p>随着科技的快速发展，各种新的技术和概念不断出现，持续出现的新技术会不断推动安全的发展。虽然，每一个新技术都会衍生出新的安全威胁和隐患，但是，这些新的安全问题也正是安全行业保持活力的源泉。所以，对于安全人员来说，这些新技术的出现既是一种挑战，也是一种机遇。</p>

<p>近几年，IoT、IPv6和区块链是三个热度很高的新技术，我也最常听到三个热词。今天，我们就一起来探讨一下，这几个新技术都面临哪些独特的安全问题。</p>

<h2 id="独特的iot安全">独特的IoT安全</h2>

<p>毫无疑问，IoT（Internet of Things，物联网）是最近十年来比较火热的一个技术。对比于当前的网络环境，IoT的网络主要有以下几个特点：</p>

<ul>
<li>设备更多：每一件小的物品都有可能成为联入互联网的设备</li>
<li>设备性能更低：受限于体积和供电量，单台设备能够搭载的硬件配置都不高</li>
<li>更加开放：由于设备的数量和类型众多，无法统一标准，因此IoT的网络环境也更加开放</li>
</ul>

<p>那么，这些特点会给安全性带来哪些新的挑战呢？关于这个问题，我推荐你玩一玩《看门狗》这款游戏，它很好地描绘了一个未来IoT城市中会面临的各类安全问题。那在此之前，我先和你分享一下我对这些新挑战的思考。</p>

<p><strong>我认为最明显的问题就是认证更加复杂了。</strong></p>

<p>在使用电脑或者手机连入网络的时候，我们可以手动输入密码来完成认证。但是，当我们想要将各类小硬件连入网络的时候，没有键盘和屏幕可以供我们输入密码。为了解决这个认证问题，目前小米等IoT厂商的解决方案是，先让手机直接控制设备，配置好WiFi密码后，再让设备连入网络。</p>

<p>但是，这其实又引发了一个新的问题，如何确认是你本人在控制设备，而不是黑客呢？针对这个问题，现在也有对应的解决方案，那就是在短时间内开放设备的控制权限，限制手机在这个时间内完成对设备的控制。</p>

<p>仔细观察的话，你会发现这个解决方案有一个假设前提：黑客没办法在短时间内发现并控制设备。在当前的环境下，这个前提是成立的。但是随着技术的发展，IoT设备可能充斥在我们身边的每一个角落里，当有一个设备被黑客控制了之后，它很可能会时刻监控这周围的环境，一旦发现其他的设备开放控制权限，就会立即黑入。可以说，通过这样的攻击方式，任何一个设备都有可能被黑客所控制。</p>

<p>因此，<strong>如何确保IoT中设备与网络、设备与设备之间的通信是可信的</strong>，是未来认证技术需要面临的主要挑战之一。</p>

<p><strong>其次，我认为物理攻击会越来越流行。</strong></p>

<p>物理攻击实际上是安全领域内的<strong>降维打击。</strong> 换句话说，当底层的硬件被黑客控制之后，我们就无法保障运行在硬件之上的系统和软件的安全性了。</p>

<p>IoT的发展，事实上正让物理攻击变得越来越容易。我总结了一张物理攻击的发展过程图，你可以看到，随着IoT越来越小、越来越智能，和我们的联系越来越紧密，物理攻击的难度也变得越来越低。在未来，公共区域内的所有设备甚至都有可以成为黑客的囊中之物。</p>

<p><img src="assets/ef3afbdebca141cebf4e9e7147bded0f.jpg" alt="" /></p>

<p>因此，<strong>如何对物理攻击进行有效的防控</strong>，也是未来安全中需要解决的主要挑战之一。</p>

<p><strong>除了带来新的安全挑战，IoT能够造成的安全威胁也变得更加复杂了</strong>。</p>

<p>目前来说，黑客利用IoT设备发起的最主要的攻击还是DDoS攻击，即黑客利用海量的IoT设备向目标服务器发送巨大的网络流量，导致服务器无法响应正常请求。</p>

<p>随着IoT的发展，黑客能够控制的设备越来越多，能够导致的影响也会越来越大。你一定在很多电影中看到过类似的情景，比如，黑客通过操纵汽车控制医疗设备等方式，导致人员伤亡。</p>

<p>因此，<strong>如何保护IoT设备免受黑客的攻击</strong>，同样会成为未来安全的主要挑战之一。</p>

<h2 id="ipv6对安全的影响">IPv6对安全的影响</h2>

<p>因为IPv4的地址空间短缺问题，IPv6是国家重点推进的一个技术方向。目前三大运营商已经完成了改造，各大互联网公司也已经接到了兼容IPv6的强制要求，我相信国内应该会很快推广和普及IPv6。</p>

<p>IPv6和IPv4相比最大区别就是IP地址变得非常庞大了。那么，庞大的IP地址对于安全来说，又意味着什么呢？</p>

<p>我认为对于黑客来说，最大的影响就是<strong>网络扫描不再可能</strong>。</p>

<p>我们知道，找到攻击目标是黑客发起攻击的第一步。因此，很多黑客会通过扫描网络来发现目标。目前，性能最优的扫描工具是<a href="https://github.com/robertdavidgraham/masscan" target="_blank">M</a><a href="https://github.com/robertdavidgraham/masscan" target="_blank">asscan</a>，它能够在5分钟内扫遍全部IPv4的地址空间。</p>

<p>而IPv6的地址空间是IPv4的2^96倍，黑客想要利用现有的扫描工具快速遍历IPv6的地址空间，显然是不可能的。因此，黑客就只能通过其他方式去精准定位目标了。</p>

<p>除了对黑客有影响以外，<strong>庞大的IP地址对公司安全来说，也同样是一种负担</strong>。</p>

<p>IP地址变多就意味着黑客手中的IP资源变多了，同时，IPv6的高变化频率还会让同一个设备的IP经常性地发生变化。因此，使用了IPv6之后，我们就很难利用黑名单对IP进行标记和处罚了。</p>

<p>另外，仍然有待观察的一点是，<strong>IPv6的复用性是否会比IPv4更低</strong>。</p>

<p>IPv4由于地址匮乏，有很高的复用性（一个学校可能都在共用一个IP地址），这让我们很难根据IP去定位到一个具体的位置或者人。</p>

<p>而IPv6的地址空间是足够的（每一粒沙子都能分配到一个IP地址），因此，IP复用就不再是一个刚需了。所以，如果IPv6的复用性远低于IPv4的话，就能让IP的定位变得更准确。那么对于安全工作来说，想要找到黑客也会更加容易。</p>

<h2 id="区块链中的安全问题">区块链中的安全问题</h2>

<p>最后，我们再来聊一聊近两年兴起的区块链。目前，区块链最成功的应用形式，就是以比特币为代表的各类虚拟货币。那么，比特币和区块链的安全性如何呢？它们又面临什么样的安全威胁呢？下面，我们一起来看。</p>

<p>我们都知道，区块链的思想是去中心化，即将数据和算力分散到每一个小的计算节点中，最终，以少数服从多数的形式来完成数据的计算和存储。这实际上是一种对完整性的保障。这么说你可能还不理解，我举个例子。</p>

<p>以货币为例，我们现在通过支付宝、微信等电子货币来完成日常交易，事实上是将钱交由支付宝和微信这样的中心机构进行集中保管。而对于支付宝、微信来说，理论上是可以对用户的余额进行篡改的，不过，因为受到了多方面限制，这一操作是无法实现的。</p>

<p>但是在比特币中，因为不存在中心机构，每个用户的余额由所有人共同保管，因此没有任何一个节点可以实现篡改。</p>

<p>但如果你仔细想想的话，就会发现这种近乎完美的完整性保障，是通过牺牲机密性来完成的。也就是说，在支付宝中，你无法知道其他用户的余额，但是在比特币中，每一笔交易和每一个用户的余额都是公开的信息，因此比特币不提供任何针对机密性的保护措施（比如，你可以在<a href="https://www.blockchain.com/explorer?view=btc_blocks" target="_blank">blockchain</a>看到所有的比特币信息）。</p>

<p>尽管比特币本身的完整性无可挑剔，但仍然无法阻止由于用户个人密钥丢失而导致的资产损失。这就好比你安装了一个特别结实的门，但只要钥匙丢了，门的存在就毫无意义了。事实上，目前大部分的比特币安全事件，都是黑客成功盗取了用户或者公司系统的比特币密钥之后，再去盗取对应账号的余额。</p>

<p>另外，比特币是目前黑客们主要使用的货币之一。其原因在于，它是匿名的（注意：匿名不是机密性，匿名是指你无法通过比特币的账号，关联到某个具体的人）。这也就保证了，即使警方知道了黑客的账户，也没办法抓到黑客。而且，由于比特币的去中心化，警方也没办法封停黑客的账户，追回被盗的比特币。</p>

<p>所以，比特币这样一种去中心化且匿名的货币体系，既不保险，也不利于政府的管控，因此国内对于以区块链为基础的电子货币落地，始终不认可。</p>

<h2 id="总结">总结</h2>

<p>今天，我们主要对 IoT、IPv6和区块链这三个热门技术及其安全性进行了盘点。这些新的技术都具备其独特的应用场景，也都带有独特的安全问题。这些问题既可能是这些技术本身所存在的一些缺陷，也可能是对已有的安全防御工作产生的威胁。</p>

<p>我们不仅要对这些新技术进行持续的关注，还要思考它们会产生的新安全需求，然后去学习对应的新知识。这也是安全人员提升自我价值，保持思维活力的有效手段。</p>

<h2 id="思考题">思考题</h2>

<p>最后，咱们来看一道思考题。</p>

<p>除了我们今天讲的这三种技术，你还接触过哪些新的技术呢？不妨和我的一样，把你对这些新技术的安全思考都写下来。</p>

<p>欢迎留言和我分享你的思考和疑惑，也欢迎你把文章分享给你的朋友。我们下一讲再见！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#fc909090c5c8cdcdcccbbc9b919d9590d29f9391" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f144a0e9e8c9508',t:'MTczNDA3NTQ1MC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>