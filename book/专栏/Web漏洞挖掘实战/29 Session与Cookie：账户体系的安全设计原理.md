<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=29&#32;Session与Cookie：账户体系的安全设计原理>
        <link rel="icon" href="/static/favicon.png">
        <title>29 Session与Cookie：账户体系的安全设计原理 </title>
        
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
                        <a class="menu-item" id="00 导读 解读OWASP Top10 2021.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/00%20%e5%af%bc%e8%af%bb%20%e8%a7%a3%e8%af%bbOWASP%20Top10%202021.md">00 导读 解读OWASP Top10 2021.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="00 开篇词 从黑客的视角找漏洞，从安全的角度优雅coding.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/00%20%e5%bc%80%e7%af%87%e8%af%8d%20%e4%bb%8e%e9%bb%91%e5%ae%a2%e7%9a%84%e8%a7%86%e8%a7%92%e6%89%be%e6%bc%8f%e6%b4%9e%ef%bc%8c%e4%bb%8e%e5%ae%89%e5%85%a8%e7%9a%84%e8%a7%92%e5%ba%a6%e4%bc%98%e9%9b%85coding.md">00 开篇词 从黑客的视角找漏洞，从安全的角度优雅coding.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="01 失效的访问控制：攻击者如何获取其他用户信息？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/01%20%e5%a4%b1%e6%95%88%e7%9a%84%e8%ae%bf%e9%97%ae%e6%8e%a7%e5%88%b6%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e5%a6%82%e4%bd%95%e8%8e%b7%e5%8f%96%e5%85%b6%e4%bb%96%e7%94%a8%e6%88%b7%e4%bf%a1%e6%81%af%ef%bc%9f.md">01 失效的访问控制：攻击者如何获取其他用户信息？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="02 路径穿越：你的Web应用系统成了攻击者的资源管理器？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/02%20%e8%b7%af%e5%be%84%e7%a9%bf%e8%b6%8a%ef%bc%9a%e4%bd%a0%e7%9a%84Web%e5%ba%94%e7%94%a8%e7%b3%bb%e7%bb%9f%e6%88%90%e4%ba%86%e6%94%bb%e5%87%bb%e8%80%85%e7%9a%84%e8%b5%84%e6%ba%90%e7%ae%a1%e7%90%86%e5%99%a8%ef%bc%9f.md">02 路径穿越：你的Web应用系统成了攻击者的资源管理器？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="03 敏感数据泄露：攻击者如何获取用户账户？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/03%20%e6%95%8f%e6%84%9f%e6%95%b0%e6%8d%ae%e6%b3%84%e9%9c%b2%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e5%a6%82%e4%bd%95%e8%8e%b7%e5%8f%96%e7%94%a8%e6%88%b7%e8%b4%a6%e6%88%b7%ef%bc%9f.md">03 敏感数据泄露：攻击者如何获取用户账户？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="04 权限不合理：攻击者进来就是root权限？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/04%20%e6%9d%83%e9%99%90%e4%b8%8d%e5%90%88%e7%90%86%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e8%bf%9b%e6%9d%a5%e5%b0%b1%e6%98%afroot%e6%9d%83%e9%99%90%ef%bc%9f.md">04 权限不合理：攻击者进来就是root权限？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="05 CSRF：为什么用户的操作他自己不承认？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/05%20CSRF%ef%bc%9a%e4%b8%ba%e4%bb%80%e4%b9%88%e7%94%a8%e6%88%b7%e7%9a%84%e6%93%8d%e4%bd%9c%e4%bb%96%e8%87%aa%e5%b7%b1%e4%b8%8d%e6%89%bf%e8%ae%a4%ef%bc%9f.md">05 CSRF：为什么用户的操作他自己不承认？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="06 加密失败：使用了加密算法也会被破解吗？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/06%20%e5%8a%a0%e5%af%86%e5%a4%b1%e8%b4%a5%ef%bc%9a%e4%bd%bf%e7%94%a8%e4%ba%86%e5%8a%a0%e5%af%86%e7%ae%97%e6%b3%95%e4%b9%9f%e4%bc%9a%e8%a2%ab%e7%a0%b4%e8%a7%a3%e5%90%97%ef%bc%9f.md">06 加密失败：使用了加密算法也会被破解吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="07 弱编码：程序之间的沟通语言安全吗？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/07%20%e5%bc%b1%e7%bc%96%e7%a0%81%ef%bc%9a%e7%a8%8b%e5%ba%8f%e4%b9%8b%e9%97%b4%e7%9a%84%e6%b2%9f%e9%80%9a%e8%af%ad%e8%a8%80%e5%ae%89%e5%85%a8%e5%90%97%ef%bc%9f.md">07 弱编码：程序之间的沟通语言安全吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="08 数字证书：攻击者可以伪造证书吗？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/08%20%e6%95%b0%e5%ad%97%e8%af%81%e4%b9%a6%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e5%8f%af%e4%bb%a5%e4%bc%aa%e9%80%a0%e8%af%81%e4%b9%a6%e5%90%97%ef%bc%9f.md">08 数字证书：攻击者可以伪造证书吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="09 密码算法问题：数学知识如何提高代码可靠性？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/09%20%e5%af%86%e7%a0%81%e7%ae%97%e6%b3%95%e9%97%ae%e9%a2%98%ef%bc%9a%e6%95%b0%e5%ad%a6%e7%9f%a5%e8%af%86%e5%a6%82%e4%bd%95%e6%8f%90%e9%ab%98%e4%bb%a3%e7%a0%81%e5%8f%af%e9%9d%a0%e6%80%a7%ef%bc%9f.md">09 密码算法问题：数学知识如何提高代码可靠性？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="10 弱随机数生成器：攻击者如何预测随机数？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/10%20%e5%bc%b1%e9%9a%8f%e6%9c%ba%e6%95%b0%e7%94%9f%e6%88%90%e5%99%a8%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e5%a6%82%e4%bd%95%e9%a2%84%e6%b5%8b%e9%9a%8f%e6%9c%ba%e6%95%b0%ef%bc%9f.md">10 弱随机数生成器：攻击者如何预测随机数？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="11 忘记加“盐”：加密结果强度不够吗？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/11%20%e5%bf%98%e8%ae%b0%e5%8a%a0%e2%80%9c%e7%9b%90%e2%80%9d%ef%bc%9a%e5%8a%a0%e5%af%86%e7%bb%93%e6%9e%9c%e5%bc%ba%e5%ba%a6%e4%b8%8d%e5%a4%9f%e5%90%97%ef%bc%9f.md">11 忘记加“盐”：加密结果强度不够吗？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="12 注入（上）：SQL注入起手式.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/12%20%e6%b3%a8%e5%85%a5%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9aSQL%e6%b3%a8%e5%85%a5%e8%b5%b7%e6%89%8b%e5%bc%8f.md">12 注入（上）：SQL注入起手式.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="13 注入（下）：SQL注入技战法及相关安全实践.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/13%20%e6%b3%a8%e5%85%a5%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9aSQL%e6%b3%a8%e5%85%a5%e6%8a%80%e6%88%98%e6%b3%95%e5%8f%8a%e7%9b%b8%e5%85%b3%e5%ae%89%e5%85%a8%e5%ae%9e%e8%b7%b5.md">13 注入（下）：SQL注入技战法及相关安全实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="14 自动化注入神器（一）：sqlmap的设计思路解析.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/14%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b3%a8%e5%85%a5%e7%a5%9e%e5%99%a8%ef%bc%88%e4%b8%80%ef%bc%89%ef%bc%9asqlmap%e7%9a%84%e8%ae%be%e8%ae%a1%e6%80%9d%e8%b7%af%e8%a7%a3%e6%9e%90.md">14 自动化注入神器（一）：sqlmap的设计思路解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="15 自动化注入神器（二）：sqlmap的设计架构解析.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/15%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b3%a8%e5%85%a5%e7%a5%9e%e5%99%a8%ef%bc%88%e4%ba%8c%ef%bc%89%ef%bc%9asqlmap%e7%9a%84%e8%ae%be%e8%ae%a1%e6%9e%b6%e6%9e%84%e8%a7%a3%e6%9e%90.md">15 自动化注入神器（二）：sqlmap的设计架构解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="16 自动化注入神器（三）：sqlmap的核心实现拆解.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/16%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b3%a8%e5%85%a5%e7%a5%9e%e5%99%a8%ef%bc%88%e4%b8%89%ef%bc%89%ef%bc%9asqlmap%e7%9a%84%e6%a0%b8%e5%bf%83%e5%ae%9e%e7%8e%b0%e6%8b%86%e8%a7%a3.md">16 自动化注入神器（三）：sqlmap的核心实现拆解.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="17 自动化注入神器（四）：sqlmap的核心功能解析.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/17%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%b3%a8%e5%85%a5%e7%a5%9e%e5%99%a8%ef%bc%88%e5%9b%9b%ef%bc%89%ef%bc%9asqlmap%e7%9a%84%e6%a0%b8%e5%bf%83%e5%8a%9f%e8%83%bd%e8%a7%a3%e6%9e%90.md">17 自动化注入神器（四）：sqlmap的核心功能解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="19 失效的输入检测（上）：攻击者有哪些绕过方案？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/19%20%e5%a4%b1%e6%95%88%e7%9a%84%e8%be%93%e5%85%a5%e6%a3%80%e6%b5%8b%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bb%95%e8%bf%87%e6%96%b9%e6%a1%88%ef%bc%9f.md">19 失效的输入检测（上）：攻击者有哪些绕过方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="20 失效的输入检测（下）：攻击者有哪些绕过方案？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/20%20%e5%a4%b1%e6%95%88%e7%9a%84%e8%be%93%e5%85%a5%e6%a3%80%e6%b5%8b%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%94%bb%e5%87%bb%e8%80%85%e6%9c%89%e5%93%aa%e4%ba%9b%e7%bb%95%e8%bf%87%e6%96%b9%e6%a1%88%ef%bc%9f.md">20 失效的输入检测（下）：攻击者有哪些绕过方案？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="21 XSS（上）：前端攻防的主战场.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/21%20XSS%ef%bc%88%e4%b8%8a%ef%bc%89%ef%bc%9a%e5%89%8d%e7%ab%af%e6%94%bb%e9%98%b2%e7%9a%84%e4%b8%bb%e6%88%98%e5%9c%ba.md">21 XSS（上）：前端攻防的主战场.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="22 XSS（中）：跨站脚本攻击的危害性.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/22%20XSS%ef%bc%88%e4%b8%ad%ef%bc%89%ef%bc%9a%e8%b7%a8%e7%ab%99%e8%84%9a%e6%9c%ac%e6%94%bb%e5%87%bb%e7%9a%84%e5%8d%b1%e5%ae%b3%e6%80%a7.md">22 XSS（中）：跨站脚本攻击的危害性.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="23 XSS（下）：检测与防御方案解析.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/23%20XSS%ef%bc%88%e4%b8%8b%ef%bc%89%ef%bc%9a%e6%a3%80%e6%b5%8b%e4%b8%8e%e9%98%b2%e5%be%a1%e6%96%b9%e6%a1%88%e8%a7%a3%e6%9e%90.md">23 XSS（下）：检测与防御方案解析.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="24 资源注入：攻击方式为什么会升级？.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/24%20%e8%b5%84%e6%ba%90%e6%b3%a8%e5%85%a5%ef%bc%9a%e6%94%bb%e5%87%bb%e6%96%b9%e5%bc%8f%e4%b8%ba%e4%bb%80%e4%b9%88%e4%bc%9a%e5%8d%87%e7%ba%a7%ef%bc%9f.md">24 资源注入：攻击方式为什么会升级？.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="25 业务逻辑漏洞：好的开始是成功的一半.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/25%20%e4%b8%9a%e5%8a%a1%e9%80%bb%e8%be%91%e6%bc%8f%e6%b4%9e%ef%bc%9a%e5%a5%bd%e7%9a%84%e5%bc%80%e5%a7%8b%e6%98%af%e6%88%90%e5%8a%9f%e7%9a%84%e4%b8%80%e5%8d%8a.md">25 业务逻辑漏洞：好的开始是成功的一半.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="26 包含敏感信息的报错：将安全开发标准应用到项目中.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/26%20%e5%8c%85%e5%90%ab%e6%95%8f%e6%84%9f%e4%bf%a1%e6%81%af%e7%9a%84%e6%8a%a5%e9%94%99%ef%bc%9a%e5%b0%86%e5%ae%89%e5%85%a8%e5%bc%80%e5%8f%91%e6%a0%87%e5%87%86%e5%ba%94%e7%94%a8%e5%88%b0%e9%a1%b9%e7%9b%ae%e4%b8%ad.md">26 包含敏感信息的报错：将安全开发标准应用到项目中.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="27 用户账户安全：账户安全体系设计方案与实践.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/27%20%e7%94%a8%e6%88%b7%e8%b4%a6%e6%88%b7%e5%ae%89%e5%85%a8%ef%bc%9a%e8%b4%a6%e6%88%b7%e5%ae%89%e5%85%a8%e4%bd%93%e7%b3%bb%e8%ae%be%e8%ae%a1%e6%96%b9%e6%a1%88%e4%b8%8e%e5%ae%9e%e8%b7%b5.md">27 用户账户安全：账户安全体系设计方案与实践.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="28 安全配置错误：安全问题不只是代码安全.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/28%20%e5%ae%89%e5%85%a8%e9%85%8d%e7%bd%ae%e9%94%99%e8%af%af%ef%bc%9a%e5%ae%89%e5%85%a8%e9%97%ae%e9%a2%98%e4%b8%8d%e5%8f%aa%e6%98%af%e4%bb%a3%e7%a0%81%e5%ae%89%e5%85%a8.md">28 安全配置错误：安全问题不只是代码安全.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="29 Session与Cookie：账户体系的安全设计原理.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/29%20Session%e4%b8%8eCookie%ef%bc%9a%e8%b4%a6%e6%88%b7%e4%bd%93%e7%b3%bb%e7%9a%84%e5%ae%89%e5%85%a8%e8%ae%be%e8%ae%a1%e5%8e%9f%e7%90%86.md">29 Session与Cookie：账户体系的安全设计原理.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="30 HTTP Header安全标志：协议级别的安全支持.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/30%20HTTP%20Header%e5%ae%89%e5%85%a8%e6%a0%87%e5%bf%97%ef%bc%9a%e5%8d%8f%e8%ae%ae%e7%ba%a7%e5%88%ab%e7%9a%84%e5%ae%89%e5%85%a8%e6%94%af%e6%8c%81.md">30 HTTP Header安全标志：协议级别的安全支持.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="31 易受攻击和过时的组件：DevSecOps与依赖项安全检查.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/31%20%e6%98%93%e5%8f%97%e6%94%bb%e5%87%bb%e5%92%8c%e8%bf%87%e6%97%b6%e7%9a%84%e7%bb%84%e4%bb%b6%ef%bc%9aDevSecOps%e4%b8%8e%e4%be%9d%e8%b5%96%e9%a1%b9%e5%ae%89%e5%85%a8%e6%a3%80%e6%9f%a5.md">31 易受攻击和过时的组件：DevSecOps与依赖项安全检查.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="32 软件和数据完整性故障：SolarWinds事件的幕后⿊⼿.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/32%20%e8%bd%af%e4%bb%b6%e5%92%8c%e6%95%b0%e6%8d%ae%e5%ae%8c%e6%95%b4%e6%80%a7%e6%95%85%e9%9a%9c%ef%bc%9aSolarWinds%e4%ba%8b%e4%bb%b6%e7%9a%84%e5%b9%95%e5%90%8e%e2%bf%8a%e2%bc%bf.md">32 软件和数据完整性故障：SolarWinds事件的幕后⿊⼿.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="33 SSRF：穿越边界防护的利刃.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/33%20SSRF%ef%bc%9a%e7%a9%bf%e8%b6%8a%e8%be%b9%e7%95%8c%e9%98%b2%e6%8a%a4%e7%9a%84%e5%88%a9%e5%88%83.md">33 SSRF：穿越边界防护的利刃.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="34 Crawler VS Fuzzing：DAST与机器学习.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/34%20Crawler%20VS%20Fuzzing%ef%bc%9aDAST%e4%b8%8e%e6%9c%ba%e5%99%a8%e5%ad%a6%e4%b9%a0.md">34 Crawler VS Fuzzing：DAST与机器学习.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="35 自动化攻防：低代码驱动的渗透工具积累.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/35%20%e8%87%aa%e5%8a%a8%e5%8c%96%e6%94%bb%e9%98%b2%ef%bc%9a%e4%bd%8e%e4%bb%a3%e7%a0%81%e9%a9%b1%e5%8a%a8%e7%9a%84%e6%b8%97%e9%80%8f%e5%b7%a5%e5%85%b7%e7%a7%af%e7%b4%af.md">35 自动化攻防：低代码驱动的渗透工具积累.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="36 智能攻防：构建个性化攻防平台.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/36%20%e6%99%ba%e8%83%bd%e6%94%bb%e9%98%b2%ef%bc%9a%e6%9e%84%e5%bb%ba%e4%b8%aa%e6%80%a7%e5%8c%96%e6%94%bb%e9%98%b2%e5%b9%b3%e5%8f%b0.md">36 智能攻防：构建个性化攻防平台.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="大咖助场 数字证书，困境与未来.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/%e5%a4%a7%e5%92%96%e5%8a%a9%e5%9c%ba%20%e6%95%b0%e5%ad%97%e8%af%81%e4%b9%a6%ef%bc%8c%e5%9b%b0%e5%a2%83%e4%b8%8e%e6%9c%aa%e6%9d%a5.md">大咖助场 数字证书，困境与未来.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节策划（一）  视频课内容精选：Web渗透测试工具教学.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/%e6%98%a5%e8%8a%82%e7%ad%96%e5%88%92%ef%bc%88%e4%b8%80%ef%bc%89%20%20%e8%a7%86%e9%a2%91%e8%af%be%e5%86%85%e5%ae%b9%e7%b2%be%e9%80%89%ef%bc%9aWeb%e6%b8%97%e9%80%8f%e6%b5%8b%e8%af%95%e5%b7%a5%e5%85%b7%e6%95%99%e5%ad%a6.md">春节策划（一）  视频课内容精选：Web渗透测试工具教学.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节策划（三） 一套测试题，看看对课程内容的掌握情况.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/%e6%98%a5%e8%8a%82%e7%ad%96%e5%88%92%ef%bc%88%e4%b8%89%ef%bc%89%20%e4%b8%80%e5%a5%97%e6%b5%8b%e8%af%95%e9%a2%98%ef%bc%8c%e7%9c%8b%e7%9c%8b%e5%af%b9%e8%af%be%e7%a8%8b%e5%86%85%e5%ae%b9%e7%9a%84%e6%8e%8c%e6%8f%a1%e6%83%85%e5%86%b5.md">春节策划（三） 一套测试题，看看对课程内容的掌握情况.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="春节策划（二）   给你推荐4本Web安全图书.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/%e6%98%a5%e8%8a%82%e7%ad%96%e5%88%92%ef%bc%88%e4%ba%8c%ef%bc%89%20%20%20%e7%bb%99%e4%bd%a0%e6%8e%a8%e8%8d%904%e6%9c%acWeb%e5%ae%89%e5%85%a8%e5%9b%be%e4%b9%a6.md">春节策划（二）   给你推荐4本Web安全图书.md</a>
                    </li>
                    
                    <li>
                        <a class="menu-item" id="结束语 无畏前行.md" href="/%e4%b8%93%e6%a0%8f/Web%e6%bc%8f%e6%b4%9e%e6%8c%96%e6%8e%98%e5%ae%9e%e6%88%98/%e7%bb%93%e6%9d%9f%e8%af%ad%20%e6%97%a0%e7%95%8f%e5%89%8d%e8%a1%8c.md">结束语 无畏前行.md</a>
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
                            <h1 id="title" data-id="29 Session与Cookie：账户体系的安全设计原理" class="title">29 Session与Cookie：账户体系的安全设计原理</h1>
                            <div><p>你好，我是王昊天。</p>

<p>我有次在访问某个页面时，为了下载一些东西，按照页面要求进行了复杂的登录操作。之后我不小心关闭了当前页面，然后再一次点开这个页面，麻木的准备再来一遍复杂的登录操作时，我神奇地发现，面前的Web应用竟然是登录成功的状态，你知道这是怎么一回事吗？</p>

<p>事实上，这个现象是<strong>由Web账户体系的安全设计所导致的</strong>。在这一讲中，我们将会对它进行学习，这样你就能清楚地知道问题的答案啦。下面我们就正式开始今天的学习。</p>

<p>现在几乎每个大型Web应用都会存在账户体系，当我们需要获取Web应用中的某些服务时，Web应用会首先对我们的身份进行认证。所以接下来，我们会从身份认证的相关基础知识入手。</p>

<h2 id="身份认证">身份认证</h2>

<p>身份认证的方式有多种，我们可以用最典型的账号密码进行认证，除此之外，我们还可以用cookie（session）、Token、数字证书以及手机验证码来验证。这里你可能对于cookie以及Token会比较陌生，不过不用担心，我们会在后面对它们进行详细的讲解。</p>

<p><img src="assets/54b4228f08aa4b55b1169d56ad546dde.jpg" alt="图片" /></p>

<p>在这些认证过程中，可以分为两种类型，即<strong>登录过程的认证以及保持登录的认证</strong>。</p>

<p>为了让你更好地理解它们二者之间的区别，我们一起来看一个示例。</p>

<p><img src="assets/05b33b057e91454b94f5533fd989b9ed.jpg" alt="图片" /></p>

<p>这是谜团（mituan.zone）的登录页面，我们需要输入正确的用户名、密码以及验证码才能通过身份认证，很明显这是登录过程的认证。</p>

<p>当我们登录成功后，我们会发现浏览器中多了一些cookie信息。</p>

<p><img src="assets/e51f35a2dfb6490289a31fb922699e3c.jpg" alt="图片" /></p>

<p>这些cookie信息有一定的有效期。在有效期内，cookie信息会一直存在，它使得我们下次访问这个页面时，无需再次输入账号密码进行登录，而是<strong>可以直接用cookie信息来实现身份认证操作</strong>，这就是保持登录的认证。</p>

<p>你现在知道导入中神奇现象发生的原因了吗？其实在导入部分中的自动登录，就是通过保持登录的认证来实现的。而这种认证方式，主要是通过会话管理来实现的，接下来让我们简单了解下会话管理的作用。</p>

<h2 id="会话管理">会话管理</h2>

<p>在学习会话管理之前，我们首先需要巩固下HTTP协议的知识。</p>

<p>HTTP协议是无状态无连接的协议，服务端对于客户端每次发送的请求都认为它是一个新的请求，上一次会话和下一次会话是没有联系的。因为它无法保存登录状态，所以从协议本身来说，它不适合用来做会话管理。</p>

<p>因此，<strong>我们会使用一个上层应用去实现我们的会话管理功能</strong>。这个应用可以在切换页面时保持登录状态，并且对用户是透明的，这样就使得我们能在短时间内再次访问一个登录过的页面，就会保持登录状态。</p>

<p>经过上述内容的学习，你已经知道了会话管理具有什么作用。接下来，让我们具体学习下会话管理的两种典型方式，即基于session的认证以及基于Token的认证。</p>

<h3 id="基于-session-的认证">基于 session 的认证</h3>

<p>Web应用可以基于session的认证来实现保持登录，它的具体实现方式如下图所示：</p>

<p><img src="assets/b2615aed44804361adeb5ad8cab81c17.jpg" alt="图片" /></p>

<p>用户在首次访问Web应用时，会将自己的账号密码通过POST方式进行上传，然后Web应用服务器会对账号密码进行检查。如果检查通过就会给用户配置一个sessionid，并将它存储在服务器内存中，之后再把这个sessionid发送给用户。</p>

<p>注意这里sessionid的位置可能在URL、隐藏域以及cookie中。由于cookie信息较为隐蔽些，所以<strong>将sessionid放在cookie中相对来说更为安全</strong>，因此这一实现方式也最普遍。</p>

<p>用户在收到Web应用服务器的回应之后，再次对Web应用发起请求的cookie中就会自动包含sessionid信息。Web应用服务器会对其中的sessionid信息进行检查，以获取用户的登录信息，如果信息正确，就让用户处于登录成功的状态，否则需要重新进行登录过程的认证。</p>

<p>值得一提的是，为了安全考虑，Web应用通常会<strong>给sessionid设置一个过期时间</strong>，使得sessionid仅在某个时间段内有效，这样就可以有效地抵御攻击者盗用sessionid绕过身份认证的行为。</p>

<p>到这里，我们已经学习了Web应用是如何利用session进行身份认证的。而这里还有一个很重要的知识点我们有必要深入了解一下，那就是在session进行身份认证中存在的典型攻击方式——<strong>会话固定攻击</strong>。</p>

<p>在之前的学习中，我们知道了sessionid可以存在于URL中。在这种情况下，如果登录前后sessionid不变化，那么攻击者就可以发起会话固定攻击。</p>

<p>这里我已经画出了会话固定攻击的示意图，让我们一起看看吧。</p>

<p><img src="assets/cc4b98324eef431cbafff2194b39b811.jpg" alt="图片" /></p>

<p>攻击者首先访问一个需要登录的网站，获取到Web应用返回的sessionid信息。由于攻击者没有账户密码，所以只能通过发送一个诱骗信息给受害者，使得受害者用这个sessionid实现登录操作。这样攻击者的sessionid就通过了验证，使得攻击者再次用这个sessionid信息访问被攻击网站时，可以直接通过保持登录的认证。</p>

<p>这就是将sessionid信息放在URL中的安全隐患。</p>

<h3 id="基于-token-的认证">基于 Token 的认证</h3>

<p>除了基于session的认证之外，Web应用还可以利用Token来实现会话管理。</p>

<p>基于Token的认证方式，如下图所示，让我们从图中观察它是如何实现的吧。</p>

<p><img src="assets/1b5f5337b4444a9ea80ca0e4d4cf2d09.jpg" alt="图片" /></p>

<p>用户首先需要通过POST方式上传账号密码信息，进行登录过程的认证，Web应用服务器接收到之后，会检查账号密码信息是否正确，如果正确就会生成一个包含密码信息的Token值，这里以JWT（JSON Web Token）为例。</p>

<p>之后服务器会将这个Token信息发送我们的浏览器，接着浏览器会将这个Token信息保存在Header中，使得以后每次请求的Header中都会包含这个Token信息。服务器在接收到Token信息后，会从中提取出用户的账户信息，并对此进行检测，然后将响应发送给我们的浏览器。</p>

<p>这就是基于Token的认证方式，下面让我们以JWT为例进行学习，深入地了解Token的具体形式。</p>

<p><img src="assets/49952823561c4587ac7924783d1384b0.jpg" alt="图片" /></p>

<p>上方方框中的内容是一个完整的JWT信息，它可以根据.分割成三个部分，我们将它不同的部分用不同的颜色进行显示。<strong>接下来，让我们逐一分析JWT各个部分的内容。</strong></p>

<p>第一个部分经过base64解码就变为了蓝色方框中的内容，其中alg的内容设置的是signature中签名使用的算法，而typ的内容则定义了这个Token的类型。</p>

<p>第二部分解码为绿色方框中的内容，它包含了用户相关的信息，Web应用可以根据这些信息来确定用户的身份。</p>

<p>最后一部分解码为橙色方框中的内容，它包含了对Token信息的完整性验证签名。其中需要用到仅有服务器知道的secret信息，这也是导致攻击者无法伪造Token信息的关键。</p>

<p><strong>以上就是JWT的组成结构。</strong>其中Header以及payload用到的都是些通用数据，攻击者很容易就可以伪造出来。唯一有难度的就是对secret签名部分的伪造，事实上，攻击者可以通过密钥爆破的方式，尝试进行Signature信息的伪造。一旦伪造成功，攻击者就可以以任意身份登录这个Web应用，这对Web应用来说是极大的威胁。所以Token信息的设计者，需要有意识地提高secret的复杂度。</p>

<p>到这里，你已经学习了会话管理的两种典型方式。接下来，让我们拓宽视野，简单了解下单点登录的知识。</p>

<h2 id="单点登录">单点登录</h2>

<p>如今的Web应用越来越多，同一个公司可能就会研发出多个Web应用，如果每个应用都需要分开登录注册，那既会使得用户感到不方便，也会增加开发成本。为了解决这个问题，大家通常会采取单点登录方案。</p>

<p>单点登录就是<strong>用户只需要登录一次就可以访问所有相互信任的应用系统</strong>。它把认证的流程统一起来，使得认证的风险集中化。</p>

<p><img src="assets/7b6def82540d419db074bfa736db39da.jpg" alt="图片" /></p>

<p>这样，我们只需要在那统一的登录流程中<strong>做好安全认证措施</strong>，就可以实现对多个应用的身份认证。单点登录既能降低开发成本，也可以提高登录的安全性。</p>

<h2 id="总结">总结</h2>

<p>这节课我们学习了账户体系的安全认证设计。</p>

<p>首先，我们学习了身份认证的方式，了解到除了我们熟悉的登录过程认证之外，还有保持登录认证这一种方式。</p>

<p>接着，我们深入学习了保持登录认证的方式，知道了它是由会话管理方法实现的。然后我们对基于session的会话管理以及基于Token的会话管理进行了全面的学习，我们不仅知道了它们保持登录认证的实现方式，还知道它们存在的安全隐患。</p>

<p>最后，我们了解了一个面对多个应用需要登录验证时的解决方案，即单点登录。使用单点登录既可以统一管理所有的登录认证，还可以降低多个Web应用的开发成本。</p>

<h2 id="思考题">思考题</h2>

<p>你知道在基于session的保持登录认证中，为什么将session信息放置在cookie中会更加安全吗？</p>

<p>欢迎在评论区留下你的思考。如果觉得今天的内容对你有所帮助的话，也欢迎你把课程分享给其他同事或朋友，我们共同学习进步！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#39555555000d0808090e795e54585055175a5654" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12c4bb0e2763f7',t:'MTczNDA1OTUwMy4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>