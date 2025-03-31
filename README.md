# 语音识别
## 接口
- 说明：语音识别
- 路由地址：/asr
- 请求方式：POST
- Content-Type: application/json
- 参数：

|名称|类型|默认值|描述|
|-|-|-|-|
|lang|int32|0|语言 (0-auto, 3-zh, 4-en, 7-yue, 11-ja, 12-ko)|
|norm|int32|14|输出格式化（14-使用ITN，15-不使用ITN）|
|samples|string|""|样本为16000hz，int16数组转byte数组base64字符串（与file二选一）|
|file|string|""|16000hz WAV文件路径（与samples二选一）|
## Example


https://github.com/user-attachments/assets/594c4716-f1c3-41e6-aefe-6ec1c969094d

![image](https://github.com/user-attachments/assets/b20af987-fd99-4ad0-8cfd-365517478d30)
