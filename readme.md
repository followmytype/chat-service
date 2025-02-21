# 簡單聊天服務
採用簡易的tcp連線通訊做到即時的收發訊息，功能如下：
1. 用戶註冊登入
2. 用戶上線通知
3. 即時收發消息
4. 廣播消息
5. 收發離線訊息

專案下有兩個目錄，分別為服務端與客戶端，需先啟動服務端後客戶端才能運行，服務端啟動一個即可。
## 使用方法
### Server端
需搭配redis做用戶資料與離線留言儲存
```shell
# ./server
# 編譯執行檔
go build -o main.exe ./main
# 啟動服務器
./main.exe [port]
```
> 可指定運行的port，預設為8889

### Client端
可啟用多個做互相聊天
```shell
# ./client
# 編譯執行檔
go build -o main.exe ./main
# 啟動服務器
./main.exe [addr]
```
> 可指定server位址，預設為`localhost:8889`
