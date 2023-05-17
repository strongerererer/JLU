# JLU
吉林大学毕设项目

配置运行方式：
1.下载并安装node.js(Vue使用)：https://nodejs.org/en

2.进入JLU\server\src\roomserver\web，运行build.bat

<img width="245" alt="image" src="https://github.com/strongerererer/JLU/assets/96679694/a803ca80-d000-459e-b370-a1e4c2c6a374">

3.如果配置vite报错，需要执行npm install命令

<img width="652" alt="image" src="https://github.com/strongerererer/JLU/assets/96679694/fd22b3ab-4192-4219-ba43-73493ca56cfe">

4.JLU\server\bin中找到config.json.example，删除后缀，在proxy中放上自己的代理端口(openai的GPT—api需要开启代理)，如：“http://127.0.0.1:xxxx”

<img width="294" alt="image" src="https://github.com/strongerererer/JLU/assets/96679694/c10de791-808b-4dc9-9184-f4d43b0edc4c">



5.注意，需要在LLM.go中，用上你自己的OpenAI Key，github会自动检测出GPT key有泄露情况让其失效。https://platform.openai.com/account/api-keys
