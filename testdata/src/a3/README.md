## テスト
依存関係のルールが守られている場合、特にエラーは発生しない (3層アーキテクチャ)

## 依存関係のルール

```
controller -> usecase -> domain/repository(抽象),domain/model <- repository(実体)
```

