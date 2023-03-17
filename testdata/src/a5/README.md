## テスト
依存関係がぐちゃぐちゃの場合、複数のエラーが発生する


## 依存関係のルール

```
controller -> usecase -> domain/repository(抽象),domain/model <- repository(実体)
```

