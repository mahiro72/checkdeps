## テスト
usecaseから触るrepositoryが抽象ではなく実体を指している場合、エラーが発生する


## 依存関係のルール

```
controller -> usecase -> domain/repository(抽象),domain/model <- repository(実体)
```

