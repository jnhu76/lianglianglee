<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=17&#32;自动化注入神器（四）：sqlmap的核心功能解析>
        <link rel="icon" href="/static/favicon.png">
        <title>17 自动化注入神器（四）：sqlmap的核心功能解析 </title>
        
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
                            <h1 id="title" data-id="17 自动化注入神器（四）：sqlmap的核心功能解析" class="title">17 自动化注入神器（四）：sqlmap的核心功能解析</h1>
                            <div><p>你好，我是王昊天。</p>

<p>在上节课中，我们重点学习了sqlmap中一个非常重要的算法——页面相似度算法。相信你对页面相似度这个概念会有更加清晰的认知，不但知道它是什么含义，而且知道它是如何计算出来的。解决了这个大难点之后，我在上节课的结尾提出了一个空连接检测功能，有了它，sqlmap就可以大大提高执行效率。完成了检测，sqlmap就进入到实际的SQL注入测试阶段了。</p>

<p>在SQL注入测试阶段，系统首先会检测有哪些注入点，然后对这些注入点逐一发送合适的payload，检测注入是否成功。如果注入成功，那么系统会将注入点存储下来，最后对它们进行输出。</p>

<p>这节课，我们就来正式学习sqlmap的SQL注入测试过程。</p>

<h2 id="注入点检测">注入点检测</h2>

<p>在SQL正式注入测试之前，sqlmap会对每个目标的参数进行过滤。将那些非动态的，不存在注入可能的参数剔除掉，留下可能的注入点。这样sqlmap仅需要对这些可能的注入点进行正式的注入测试即可。</p>

<h3 id="动态参数检测">动态参数检测</h3>

<p>我们首先来看sqlmap是如何检测动态参数的。这部分代码依旧在start函数中，紧接着空连接检测出现。</p>

<pre><code class="language-python"># sqlmap首先对所有可用于注入测试的参数进行简单的优先级排序。
   parameters = list(conf.parameters.keys())
# 定义测试列表的顺序。（从后到前）
   orderList = (PLACE.CUSTOM_POST, PLACE.CUSTOM_HEADER, PLACE.URI, PLACE.POST, PLACE.GET)
# 对测试参数排好序之后，系统开始对参数进行过滤操作。
   proceed = True
   for place in parameters:
       skip = # ...
       if skip:
           continue
       if place not in conf.paramDict:
           continue
       paramDict = conf.paramDict[place]
       paramType = conf.method if conf.method not in (None, HTTPMETHOD.GET, HTTPMETHOD.POST) else place
# ...
       for parameter, value in paramDict.items():
           if not proceed:
               break
# 经过过滤，将该参数加入到测试过的参数中，防止重复测试。
           kb.testedParams.add(paramKey)
</code></pre>

<p>我们可以结合代码中的注释，来理解参数的过滤。首先sqlmap会对待测参数进行一个优先级排序。在排序完成之后，系统会根据用户的配置信息，对这些参数进行过滤操作。这里我举一个例子来让你更加容易理解这一步骤。例如，当用户配置的检测level小于2时，那么系统就会跳过对cookie参数的检测过程。</p>

<p>过滤完成之后，我们就会进入到你最熟悉的一步——SQL注入测试过程。让我们结合代码，分析sqlmap是如何进行SQL注入测试的。</p>

<pre><code class="language-python">if testSqlInj:
# 开始注入测试
    try:
# ...
# 进入启发式注入测试。
        check = heuristicCheckSqlInjection(place, parameter)
# 当启发式注入测试失败，就跳过该参数。
        if check != HEURISTIC_TEST.POSITIVE:
            if conf.smart or (kb.ignoreCasted and check == HEURISTIC_TEST.CASTED):
# ... 
               continue
# ...
# 通过启发式注入测试后，就会进入到SQL注入测试阶段。
        injection = checkSqlInjection(place, parameter, value)
</code></pre>

<h3 id="启发式注入测试">启发式注入测试</h3>

<p>如果一个参数被检测为注入点，那我们就可以对它进行注入测试。为了提高注入测试的效率，系统会过滤一些注入成功率较低的注入点，这需要首先对它进行一个启发式注入测试。下面让我们结合代码，对启发式注入测试有个更具体的理解。</p>

<pre><code class="language-python">def heuristicCheckSqlInjection(place, parameter):

# 如果配置中设置了跳过启发式注入测试，就返回结果None，当使用者没有特殊配置conf.start这个配置项为false，就会跳过该参数的注入检测。
    if conf.skipHeuristics:
        return None

# 初始化参数，并根据用户设置的偏好制作payload。
    origValue = conf.paramDict[place][parameter]
    paramType = conf.method if conf.method not in (None, HTTPMETHOD.GET, HTTPMETHOD.POST) else place

    prefix = &quot;&quot;
    suffix = &quot;&quot;
    randStr = &quot;&quot;

    if conf.prefix or conf.suffix:
        if conf.prefix:
            prefix = conf.prefix

        if conf.suffix:
            suffix = conf.suffix

    while randStr.count('\'') != 1 or randStr.count('\&quot;') != 1:
        randStr = randomStr(length=10, alphabet=HEURISTIC_CHECK_ALPHABET)

    kb.heuristicMode = True

    payload = &quot;%s%s%s&quot; % (prefix, randStr, suffix)
    payload = agent.payload(place, parameter, newValue=payload)

# 利用payload 请求目标页面的响应内容。
    page, _, _ = Request.queryPage(payload, place, content=True, raise404=False)

    kb.heuristicPage = page
    kb.heuristicMode = False
</code></pre>

<p>系统首先会判断，用户是否设置跳过启发式注入测试，如果设置了，则返回<code>None</code>。如果没有设置，那么系统就会获取到用户设置的偏好<code>prefix</code>以及<code>suffix</code>，然后据此构造出合适的payload，并发送给目标，获取到响应内容<code>page</code>。</p>

<pre><code class="language-python"># 检测请求目标的响应中是否有数据库错误。
   parseFilePaths(page)
   result = wasLastResponseDBMSError()
   infoMsg = &quot;heuristic (basic) test shows that %sparameter '%s' might &quot; % (&quot;%s &quot; % paramType if paramType != parameter else &quot;&quot;, parameter)
# 检测page中是否有。
   def _(page):
       return any(_ in (page or &quot;&quot;) for _ in FORMAT_EXCEPTION_STRINGS)
   casting = _(page) and not _(kb.originalPage)
</code></pre>

<p>系统会根据获取到的内容，判断其中的报错信息。其中，如果为数据库报错信息，那么<code>result</code>的值为<code>True</code>。如果是设置在</p>

<p><code>sqlmap/lib/core/settings.py</code>文件中<code>FORMAT_EXCEPTION_SRTINGS</code>配置项中定义的类型转化错误信息，那么就会用<code>casting</code>来储存错误内容。</p>

<pre><code class="language-python">#     ...
# 当存在定义的问题时，发出报错信息。
    if casting:
        errMsg = &quot;possible %s casting detected (e.g. '&quot; % (&quot;integer&quot; if origValue.isdigit() else &quot;type&quot;)

        platform = conf.url.split('.')[-1].lower()
        if platform == WEB_PLATFORM.ASP:
            errMsg += &quot;%s=CInt(request.querystring(\&quot;%s\&quot;))&quot; % (parameter, parameter)
        elif platform == WEB_PLATFORM.ASPX:
            errMsg += &quot;int.TryParse(Request.QueryString[\&quot;%s\&quot;], out %s)&quot; % (parameter, parameter)
        elif platform == WEB_PLATFORM.JSP:
            errMsg += &quot;%s=Integer.parseInt(request.getParameter(\&quot;%s\&quot;))&quot; % (parameter, parameter)
        else:
            errMsg += &quot;$%s=intval($_REQUEST[\&quot;%s\&quot;])&quot; % (parameter, parameter)

        errMsg += &quot;') at the back-end web application&quot;
        logger.error(errMsg)

        if kb.ignoreCasted is None:
            message = &quot;do you want to skip those kind of cases (and save scanning time)? %s &quot; % (&quot;[Y/n]&quot; if conf.multipleTargets else &quot;[y/N]&quot;)
            kb.ignoreCasted = readInput(message, default='Y' if conf.multipleTargets else 'N', boolean=True)

# 当数据库报错时，判断出注入漏洞很可能存在。
    elif result:
        infoMsg += &quot;be injectable&quot;
        if Backend.getErrorParsedDBMSes():
            infoMsg += &quot; (possible DBMS: '%s')&quot; % Format.getErrorParsedDBMSes()
        logger.info(infoMsg)

# 否则判定为不存在注入漏洞。
    else:
        infoMsg += &quot;not be injectable&quot;
        logger.warn(infoMsg)

    kb.heuristicMode = True
    kb.disableHtmlDecoding = True
</code></pre>

<p>最后，函数会根据<code>casting</code>以及<code>result</code>中的内容进行输出。我在这里画了一个它的流程图，帮助你对它的作用进行理解。</p>

<p><img src="assets/b9438b2accb242d1950fa7b679d6f459.jpg" alt="图片" /></p>

<p>图中启发式注入结果分为三种，其中阳性代表该参数大概率可以注入，类型转换和阴性都代表了该参数大概率不可以注入。我们会发现，想要判断是否可以注入，只需要判断有无数据库报错信息就可以了，有的话就认为该参数可注入，否则就认为不可注入。</p>

<p>除了进行启发式SQL注入检测之外，sqlmap还会做一些不属于它的工作，包括进行简单的xss检测和文件包含检测。</p>

<pre><code class="language-python"># 更换payload，检测xss以及文件包含。
randStr1, randStr2 = randomStr(NON_SQLI_CHECK_PREFIX_SUFFIX_LENGTH), randomStr(NON_SQLI_CHECK_PREFIX_SUFFIX_LENGTH)
value = &quot;%s%s%s&quot; % (randStr1, DUMMY_NON_SQLI_CHECK_APPENDIX, randStr2)
payload = &quot;%s%s%s&quot; % (prefix, &quot;'%s&quot; % value, suffix)
payload = agent.payload(place, parameter, newValue=payload)
page, _, _ = Request.queryPage(payload, place, content=True, raise404=False)

paramType = conf.method if conf.method not in (None, HTTPMETHOD.GET, HTTPMETHOD.POST) else place

# 进行xss检测。
if value.upper() in (page or &quot;&quot;).upper():
    infoMsg = &quot;heuristic (XSS) test shows that %sparameter '%s' might be vulnerable to cross-site scripting (XSS) attacks&quot; % (&quot;%s &quot; % paramType if paramType != parameter else &quot;&quot;, parameter)
    logger.info(infoMsg)

    if conf.beep:
        beep()

# 进行文件包含检测。
for match in re.finditer(FI_ERROR_REGEX, page or &quot;&quot;):
    if randStr1.lower() in match.group(0).lower():
        infoMsg = &quot;heuristic (FI) test shows that %sparameter '%s' might be vulnerable to file inclusion (FI) attacks&quot; % (&quot;%s &quot; % paramType if paramType != parameter else &quot;&quot;, parameter)
        logger.info(infoMsg)

        if conf.beep:
            beep()

        break

kb.disableHtmlDecoding = False
kb.heuristicMode = False

return kb.heuristicTest
</code></pre>

<p>最终的检测结果都会在全局变量<code>kb</code>中保存起来，这个全局变量我们在之前的课程中学习过。到此，启发式注入检测的函数已经完成，接下来会进入真正的SQL注入检测，这是sqlmap最核心的功能，没有之一！</p>

<h2 id="checksqlinjection函数">checkSqlInjection函数</h2>

<p>sqlmap对启发式注入的检测结果进行简单地判断后，程序就会进入sqlmap最核心的函数checkSqlInjection中。这个函数用于实现注入检测的核心功能，包括布尔注入、联合注入、报错注入、堆注入等检测。</p>

<p>下面让我们观察它的代码来理解这个注入检测功能。</p>

<pre><code class="language-python">def checkSqlInjection(place, parameter, value):

# 根据参数的类型选择 boundary 。
    injection = InjectionDict()

    threadData = getCurrentThreadData()

    if isDigit(value):
        kb.cache.intBoundaries = kb.cache.intBoundaries or sorted(copy.deepcopy(conf.boundaries), key=lambda boundary: any(_ in (boundary.prefix or &quot;&quot;) or _ in (boundary.suffix or &quot;&quot;) for _ in ('&quot;', '\'')))
        boundaries = kb.cache.intBoundaries
    elif value.isalpha():
        kb.cache.alphaBoundaries = kb.cache.alphaBoundaries or sorted(copy.deepcopy(conf.boundaries), key=lambda boundary: not any(_ in (boundary.prefix or &quot;&quot;) or _ in (boundary.suffix or &quot;&quot;) for _ in ('&quot;', '\'')))
        boundaries = kb.cache.alphaBoundaries
    else:
        boundaries = conf.boundaries
</code></pre>

<p>这个函数首先会判断参数的类型，然后根据参数的不同类型设置合适的闭合方式。解决完寻找注入点以及闭合参数这个问题后，下面让我们进入到payload的选择中。</p>

<p>我们知道，payload的选择和数据库的类型有很大的关系，所以sqlmap在构造payload前，会先尝试探测目标数据库的类型。</p>

<pre><code class="language-python"># 判断是否配置数据库类型。
if conf.dbms is None:

# 探测目标数据库类型。
    if not injection.dbms and PAYLOAD.TECHNIQUE.BOOLEAN in injection.data:
        if not Backend.getIdentifiedDbms() and kb.heuristicDbms is None and not kb.droppingRequests:
            kb.heuristicDbms = heuristicCheckDbms(injection)

# 根据探测结果输出提示信息。
    if kb.reduceTests is None and not conf.testFilter and (intersect(Backend.getErrorParsedDBMSes(), SUPPORTED_DBMS, True) or kb.heuristicDbms or injection.dbms):
        msg = &quot;it looks like the back-end DBMS is '%s'. &quot; % (Format.getErrorParsedDBMSes() or kb.heuristicDbms or joinValue(injection.dbms, '/'))
        msg += &quot;Do you want to skip test payloads specific for other DBMSes? [Y/n]&quot;
        kb.reduceTests = (Backend.getErrorParsedDBMSes() or [kb.heuristicDbms]) if readInput(msg, default='Y', boolean=True) else []
</code></pre>

<p>如果用户在配置中指定了目标数据库的类型，那么就无需探测，用指定类型即可。否则需要用<code>heuristicCheckDbms(injection)</code>函数来判断目标数据库类型。它的判断方法是，发送一些payload给测试目标，然后根据获得的响应判断数据库的类型。</p>

<p>判断出目标数据库的类型之后，系统会根据获得的数据库类型以及用户的配置，挑选适合的测试用例，然后根据这些测试用例以及之前配置的boundary，构造适合的payload。</p>

<pre><code class="language-python"># 配置联合查询的信息。
if stype == PAYLOAD.TECHNIQUE.UNION:
    configUnion(test.request.char)

    if &quot;[CHAR]&quot; in title:
        if conf.uChar is None:
            continue
        else:
            title = title.replace(&quot;[CHAR]&quot;, conf.uChar)
# ...
# 用户指定了测试方法的配置。
if conf.technique and isinstance(conf.technique, list) and stype not in conf.technique:
    debugMsg = &quot;skipping test '%s' because user &quot; % title
    debugMsg += &quot;specified testing of only &quot;
    debugMsg += &quot;%s techniques&quot; % &quot; &amp; &quot;.join(PAYLOAD.SQLINJECTION[_] for _ in conf.technique)
    logger.debug(debugMsg)
    continue

# ...
# 根据指定的数据库以及用户的配置信息，对payload进行筛选。
if conf.technique and isinstance(conf.technique, list) and stype not in conf.technique:
    debugMsg = &quot;skipping test '%s' because user &quot; % title
    debugMsg += &quot;specified testing of only &quot;
    debugMsg += &quot;%s techniques&quot; % &quot; &amp; &quot;.join(PAYLOAD.SQLINJECTION[_] for _ in conf.technique)
    logger.debug(debugMsg)
    continue

# ...
# 对payload去重。
if fstPayload:
    boundPayload = agent.prefixQuery(fstPayload, prefix, where, clause)
    boundPayload = agent.suffixQuery(boundPayload, comment, suffix, where)
    reqPayload = agent.payload(place, parameter, newValue=boundPayload, where=where)
</code></pre>

<p>sqlmap准备完payload之后，就到了你最期待的注入测试环节，这个过程和我们手动测试类似，系统会使用不同的注入测试方法，包括布尔注入、报错注入、时延注入以及联合注入。</p>

<pre><code class="language-python"># 布尔注入
if method == PAYLOAD.METHOD.COMPARISON:
    def genCmpPayload():
        sndPayload = agent.cleanupPayload(test.response.comparison, origValue=value if place not in (PLACE.URI, PLACE.CUSTOM_POST, PLACE.CUSTOM_HEADER) and BOUNDED_INJECTION_MARKER not in (value or &quot;&quot;) else None)

# ...
# 报错注入
elif method == PAYLOAD.METHOD.GREP:
    try:
        page, headers, _ = Request.queryPage(reqPayload, place, content=True, raise404=False)
        output = extractRegexResult(check, page, re.DOTALL | re.IGNORECASE)
        output = output or extractRegexResult(check, threadData.lastHTTPError[2] if wasLastResponseHTTPError() else None, re.DOTALL | re.IGNORECASE)

# ...
# 时延注入
elif method == PAYLOAD.METHOD.TIME:
    trueResult = Request.queryPage(reqPayload, place, timeBasedCompare=True, raise404=False)
    trueCode = threadData.lastCode

# ...
# 联合注入
elif method == PAYLOAD.METHOD.UNION:
    configUnion(test.request.char, test.request.columns)
</code></pre>

<p>做完这些注入测试后，系统会收到响应。我们平时会通过观察响应来判断注入是否成功，但是系统要如何判断呢？聪明的你或许想到了，这就是之前我们学习的页面相似度，我们在学习sqlmap判断waf时就用到了它。其实，根据注入方式的不同，sqlmap对于注入结果的判断方式也是不同的。</p>

<p>在报错注入中，系统会通过对页面的响应结果进行正则匹配，判断响应中是否有报错信息，如果有就判断注入成功，否则判断注入失败。</p>

<pre><code class="language-python"># 报错注入判断注入是否成功。
page, headers, _ = Request.queryPage(reqPayload, place, content=True, raise404=False)
output = extractRegexResult(check, page, re.DOTALL | re.IGNORECASE)
output = output or extractRegexResult(check, threadData.lastHTTPError[2] if wasLastResponseHTTPError() else None, re.DOTALL | re.IGNORECASE)
# ...
injectable = True
</code></pre>

<p>在布尔注入中，系统会判断返回页面的相似度，如果结果为假，那么说明系统会根据错误结果进行不同的响应，这就意味着布尔注入是成功的。</p>

<pre><code class="language-python">falseResult = Request.queryPage(genCmpPayload(), place, raise404=False)

if not falseResult:
    # ...

    injectable = True
</code></pre>

<p>在时延注入中，sqlmap会发送<code>sleep([random])</code>的请求，判断请求时间是否大于“平均时间+7*标准差”，注意这里的标准差是一个时间阈值，如果大于就认为存在时延注入。</p>

<pre><code class="language-python">if trueResult:

    if SLEEP_TIME_MARKER in reqPayload:
        falseResult = Request.queryPage(reqPayload.replace(SLEEP_TIME_MARKER, &quot;0&quot;), place, timeBasedCompare=True, raise404=False)
        if falseResult:
            continue

# ...   
        injectable = True
</code></pre>

<p>在联合注入中，系统会通过unionTest函数来判断联合注入是否存在。它的实现原理比较复杂，我们可以将它简化一下，只需要比较联合注入得到的响应和原本内容是否一致，就可以做出判断，如果不一致，则说明存在联合注入问题。</p>

<pre><code class="language-python">reqPayload, vector = unionTest(comment, place, parameter, value, prefix, suffix)

if isinstance(reqPayload, six.string_types):
    infoMsg = &quot;%sparameter '%s' is '%s' injectable&quot; % (&quot;%s &quot; % paramType if paramType != parameter else &quot;&quot;, parameter, title)
    logger.info(infoMsg)

    injectable = True
</code></pre>

<p>最后系统将结果记录下来，并且输出给使用者，这就是我们在使用sqlmap时看到的结果信息。</p>

<p>至此，经过四讲的学习，我们终于学完了这款自动化注入测试神器，希望你可以了解sqlmap的底层原理，从而更好的使用这款工具。</p>

<h2 id="总结">总结</h2>

<p>在这节课里，我们深入研究了sqlmap的真正SQL注入过程。为了你能更好的理解，我们主要通过观察它的源代码对它进行学习。</p>

<p>在这个过程中，我们首先学习了sqlmap对于注入点的检测，其中包括了动态参数的检测以及启发式注入测试。在实际注入测试的过程中，我们只会对通过检测的参数进行注入的探测。通过这个过程筛选参数，可以提高sqlmap的运行效率。</p>

<p>最后我们进入到最重要的一步中，即真正的注入测试，我们了解了它的测试过程。其中有payload的配置、对目标数据库信息的探测、筛选合适的payload以及实际的注入测试过程。完成测试，系统会根据页面相似度来判断注入结果，而对于不同的注入方式，sqlmap的判断方式也是不同的。我们将联合注入、报错注入、时延注入以及布尔注入的判断方法一一展开，对它们分别进行了介绍。</p>

<p>截止到目前，你已经完成了对SQL注入原理、攻击方式、防御方案以及自动化注入工具sqlmap的学习，结合对sqlmap原理的学习，快去自己尝试一下自动化注入的威力吧！</p>

<h2 id="思考">思考</h2>

<p>sqlmap在实现中有什么值得改进的地方吗？</p>

<p>欢迎在评论区留下你的思考，我们下节课再见。</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#2a464646131e1b1b1a1d6a4d474b434604494547" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12c147dec763f7',t:'MTczNDA1OTM2My4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>