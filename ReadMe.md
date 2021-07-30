# Firestore Emulator Testing

## 必要なもの
### Firestore Emulator
[こちら](https://qiita.com/castaneai/items/c7d68cbee1a6e3655247)の記事をみるとよいかも？


## 実行
```
# Emulatorの起動
gcloud beta emulators firestore start --host-port=localhost:8812

# 実行例
go run cmd/structs/main.go -e envfile
```