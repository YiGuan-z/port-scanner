# port-scanner

一个简单的tcp端口扫描器
A simple tcp port scanner

## 使用方式

  ```shell
  git clone https://github.com/YiGuan-z/port-scanner.git
  ``` 
  ```shell
  cd port-scanner
  ```
   ```shell
  make
  ```

- 参数说明
  - -host：指定主机名
  - -start：指定起始端口 默认为21开始
  - -end：指定结束端口 默认为200结束
  - -cache：指定扫描速率 默认为100
- 如果是Linux用户或者是mac用户，可以使用sort命令对输出进行排序，代码如下
  
  ```shell
    go run main.go -host localhost -start 1 -end 65535 | sort -n -k 1
  ```
