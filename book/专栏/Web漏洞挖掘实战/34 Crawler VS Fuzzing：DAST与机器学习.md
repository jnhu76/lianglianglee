<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=34&#32;Crawler&#32;VS&#32;Fuzzing：DAST与机器学习>
        <link rel="icon" href="/static/favicon.png">
        <title>34 Crawler VS Fuzzing：DAST与机器学习 </title>
        
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
                            <h1 id="title" data-id="34 Crawler VS Fuzzing：DAST与机器学习" class="title">34 Crawler VS Fuzzing：DAST与机器学习</h1>
                            <div><p>你好，我是王昊天。</p>

<p>经过之前的学习，相信你已经对Web漏洞挖掘的基础知识有了一定的理解。其中，我们学习了多种Web漏洞攻击方式，例如SQL注入以及SSRF攻击。事实上，网络攻防与战争是类似的。攻击者就好像侵略者一样想要通过攻击行为去获取一些资源或者满足自己的目的，而Web应用开发者则像抵抗军一样去尽力守护自己的城市。</p>

<p>我对战争也有一定的研究，在八国联军侵略祖国时，我们对敌军飞机的进攻束手无策，只有在看到它之后使用手枪尝试发射子弹击落它们。这无论是对敌人的侦查以及攻击方式都是不够强力的。当我们祖国富强之后，有了雷达探测以及导弹系统，就能及时发现敌人的飞机并将它们击落。其实，在网络攻防中也存在类似于雷达、导弹一样的工具，它们为Crawler与Fuzzing，这一讲就让我们一起学习它吧！</p>

<h2 id="crawler-vs-fuzzing">Crawler VS Fuzzing</h2>

<p>Crawler即网络爬虫，我们可以利用它来爬取一些自己想要的数据，例如豆瓣网的评分以及财经网的股市信息等。在网络攻防中它还具有另一个作用，我们可以用它来爬取待攻击页面的所有链接，这样就可以扩充我们的攻击目标，使我们不仅对当前页面发起攻击，还可以对其相关的页面进行攻击。所以，我们说 <strong>Crawler就像个雷达一样，它可以扫描当前页面的信息</strong>。这样，我们就能获取一些其它可供攻击的目标页面。</p>

<p>而Fuzzing即模糊测试，你可能对它较为陌生，不过请不用担心，事实上它就是一种软件测试技术，我们可以通过提供大量非预期的输入并监视异常结果来发现软件的故障，从而寻找Web应用的漏洞。可见 <strong>Fuzzing 则像一个导弹一样，可以对页面发起强有力的攻击行为</strong>。</p>

<p>接下来，我们将具体学习一下它们。首先，我们来学习Crawler的相关知识。</p>

<h2 id="crawler网络爬虫">Crawler网络爬虫</h2>

<p>Crawler网络爬虫是一种按照一定的规则，从Web应用中获取想要的信息的程序或脚本，我们常将它称为网页蜘蛛、网络机器人。</p>

<p>这么说可能有点抽象，为了让你更好地理解网络爬虫，下面我们一起看一个简单的示例。</p>

<pre><code class="language-python">from urllib import request
print(request.urlopen(request.Request(&quot;https://lev.zone/&quot;)).read().decode(&quot;utf-8&quot;))
</code></pre>

<p>上述是一段 Python 爬虫代码，利用它我们就可以获取到潮汐社区首页的内容。运行代码获取到的输出为：</p>

<pre><code class="language-plain">...
&lt;p class=&quot;tlak_p&quot;&gt;使用潮汐时，你无需进行工具的安装或环境的配置，即可一键开启使用，在线对安全资产进行检测，极速开启主动安全学习之旅。&lt;/p&gt;
...
            &lt;h5 class=&quot;mt-4&quot;&gt;自动化攻击编排&lt;/h5&gt;
            &lt;p class=&quot;tlak_p&quot;&gt;在潮汐，你可以根据接口，调用不同的安全工具，创建自动化的攻击检测编排，也可以直接使用其他小伙伴贡献的自动化编排，体验自动化主动安全检测的魅力！&lt;/p&gt;
...
&lt;a href=&quot;./pages/how-use.html&quot; class=&quot;text-dark icon-move-right&quot;&gt;了解更多
</code></pre>

<p>由于内容太多，这里我截取了一部分输出，可以看到我们已经成功获取到了潮汐页面的内容。之后，你可以利用一些比较好用的 Python 库例如BS4，对数据内容进行解析，提取出我们需要的数据。<strong>这样，我们无需访问页面也可以获取到页面中我们想要的信息啦</strong>。</p>

<p>除此之外，还有很多情况需要Crawler爬虫的支持。一个典型的场景为冷启动问题，例如一个社区在创立之初，它的用户数目肯定是很少的，有一些新加入的用户看到社区没什么人，很容易因此放弃使用该社区。开发者为了避免这种情况，就会使用爬虫去一些人流量较大的Web应用例如微博，爬取一些用户的动态放到自己的社区里，这有利于构造出社区的氛围。</p>

<p>另一种情形为搜索引擎，如今有很多出名的搜索引擎，例如百度以及 Google。它们当中保存有很多的数据，这些数据都是通过爬虫所获取到的。总体来说，现在是一个大数据时代，我们可以通过爬虫这个低成本高收益的方式去获取想要的数据。</p>

<p>以上都是爬虫的通用功能，而对于攻击者而言，爬虫对我们也是有很多帮助的。当我们想要攻击一个Web应用时，例如<a href="http://lev.zone，我们会想要获取这个Web页面的所有链接，这可以扩大我们的攻击面，使得我们攻击成功的概率增加。而需要做到这点，我们**只需要将爬虫获取到的内容做解析，将其中" target="_blank">http://lev.zone，我们会想要获取这个Web页面的所有链接，这可以扩大我们的攻击面，使得我们攻击成功的概率增加。而需要做到这点，我们**只需要将爬虫获取到的内容做解析，将其中</a> <a href=“xxx”> 中的内容进行提取，就可以获取到页面的链接了**。</p>

<p>到这里，你已经对Crawler网络爬虫的实现方式、原理及功能有了一定的了解。下面，让我们一起学习另一部分Fuzzing模糊测试的内容。</p>

<h2 id="fuzzing模糊测试">Fuzzing模糊测试</h2>

<p>Fuzzing技术最早诞生于1950年，那时候的计算机数据主要保存在打孔卡片上，计算机程序想要对数据进行操作就需要读取这些卡片的数据进行计算和输出。</p>

<p>这个设想是可行的，可是有时候会有一些垃圾卡片，计算机在读取它们时会获得一些不正常的输入信息，这些偶然的错误会导致计算机程序产生错误和异常甚至崩溃，这就是Fuzzing最初的来历。</p>

<p>如今，随着计算机技术的发展，Fuzzing技术也在不断发展，现在的Fuzzing是依靠计算机软件自动执行的，包括使用随机函数生成随机的测试用例，然后计算机会将这些输入发送给测试接口，并自动分析系统是否因为这些输入导致异常的发生。<strong>这种依赖于计算机的自动执行方式会使得Fuzzing的效率大大提升。</strong>效率的提高也使得Fuzzing模糊测试成为了目前最主流的漏洞挖掘方案，据了解，近年来80%以上的漏洞都是通过Fuzzing发现的。</p>

<p>不知道你是否还记得我们在学习XSS注入时，曾经学过一款名为XSStrike的XSS检测工具，事实上它就是一款典型的Fuzzing工具。</p>

<p><img src="assets/11cd98c32f1749cc895186ed8529e970.jpg" alt="图片" /></p>

<p>从工具的使用截图中，我们可以看到它按照一定的格式随机生成了很多的payload，然后将它发送给我们指定的链接，并将页面的异常响应概率返回给了我们。这就是Fuzzing模糊测试的功能之一，其实我们还可以利用它来绕WAF、判断Web应用是否有注入漏洞以及 Bug 的存在。</p>

<p>到这里，你已经对Fuzzing模糊测试有了一定的理解，接下来，让我们一起进入到实战部分，将Crawler以及Fuzzing相结合，这可以帮助你更好地理解Crawler与Fuzzing各自的作用与彼此之间的联系。</p>

<h2 id="实战部分">实战部分</h2>

<p>从这一讲的实战部分开始，我们将会用到一个好用的安全工具网站——<strong>潮汐社区版</strong><a href="http://lev.zone" target="_blank">http://lev.zone</a>。它可以帮助我们解决安全工具使用的相关问题，例如安全工具的使用、更新及维护，我们利用它可以便捷地使用各种安全工具。如果你乐于分享，你也可以将自己的工具上传到其中，供其他用户使用。</p>

<p><img src="assets/6099de4e65744150adbf7ce4b553f5e4.jpg" alt="图片" /></p>

<p>接下来，我们需要点击立即体验并注册账号，在注册过程中需要用到邀请码，你使用VefMiMj7N37tHDL7即可，不过它仅可支持有限个用户使用。然后登录进去，就能看到很多的安全工具啦。</p>

<p><img src="assets/3513aa5f17c64000bbfb4cf2984ee914.jpg" alt="图片" /></p>

<p>其中就包括了Crawler工具以及Fuzzing工具。不过我们在使用之前需要进行一定的配置设置。</p>

<p>这里需要点击网页顶栏的文档，然后选择使用说明就可以看到该应用的配置方式。</p>

<p><img src="assets/d301c9275bfc4a9a9f71248ddd46eac0.jpg" alt="图片" /></p>

<p>注意，因为我们仅仅需要使用集成好了的工具，所以我们只用按照文档完成 <strong>Docker安装、VPN介入凭证设置、SSH连接配置以及Docker-compose启动</strong>这四步即可。完成之后，我们就可以任意使用其中的工具了。</p>

<p>首先，我们选择crawler:1.0这款工具，然后用它来创建一个任务，它可以用来爬取一个网站的所有同域名链接。</p>

<p>这里，我们选择谜团<a href="http://mituan.zone中的专项·XSS跨站脚本攻击（实战）作为靶场，将它的链接例如：" target="_blank">http://mituan.zone中的专项·XSS跨站脚本攻击（实战）作为靶场，将它的链接例如：</a></p>

<pre><code class="language-plain">http://be10a2b2f16548f38ed07112904ddaa8.app.mituan.zone
</code></pre>

<p>输入到crawler工具中，点击运行，等待任务完成就会返回爬取到的链接：</p>

<pre><code class="language-plain">http://be10a2b2f16548f38ed07112904ddaa8.app.mituan.zone/level1.php?name=test
</code></pre>

<p>这样，我们就完成了爬虫的功能，接着，我们将爬取到的链接传递给Fuzzing模糊测试工具XSStrike使用。</p>

<p>结合我们之前对XSStrike的学习，从它运行的结果来看，我们可以判断出该页面存在XSS注入攻击。</p>

<p>到这里，我们已经成功运用了Crawler与Fuzzing对一个Web应用发起漏洞探测。</p>

<h2 id="总结">总结</h2>

<p>在这一讲中，我们学习了Crawler网络爬虫与Fuzzing模糊测试。</p>

<p>首先，我们分析了Crawler网络爬虫的原理，即向Web应用发送请求获取响应内容，然后从响应内容中提取需要的信息。并且，我们还知道了爬虫的功能，它不仅可以用来获取信息，还可以用来解决冷启动问题。从攻击者角度来说，爬虫还可以用来扩大攻击面，让我们的思维进行扩展。</p>

<p>之后，我们学习了Fuzzing模糊测试，了解到它其实就是按照一定的规则构造随机输入，然后将输入发送给网页接口进行测试，之后根据响应结果是否有异常来判断Web应用是否存在与此相关的安全问题。</p>

<p>最后，我们利用潮汐社区版，在对该Web应用进行一定的配置之后，对谜团中的靶场XSS跨站脚本攻击（实战）发起攻击行为，先使用Crawler网络爬虫获取攻击页面的链接，然后将获取到的链接交给Fuzzing模糊测试工具XSStrike使用，根据结果，我们发现该靶场存在XSS注入问题。</p>

<h2 id="思考题">思考题</h2>

<p>你认为Fuzzing模糊测试工具的优点有哪些？</p>

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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#e18d8d8dd8d5d0d0d1d6a1868c80888dcf828e8c" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12c60cd8a963f7',t:'MTczNDA1OTU1OC4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>