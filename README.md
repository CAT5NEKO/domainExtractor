# リストの右側だけ抽出するやつ
非常にニッチですが、例えばMisskeyのブロックリストをhtml形式でエクスポートした際に

```dotenv
1   hoge.ltd
2   fuga.ltd
.
.
.
123 piyo.ltd
```
みたいな感じで抽出されたときに、数字めっちゃ邪魔だな。って思ったときに使うツールです。
実行するとドメインの部分だけを抽出してテキストファイルにエクスポートします。

## 使い方
同一のディレクトリ内に`block.txt`というファイルを作成し、その中にhtmlからコピーした行数入りのブロックリストを貼り付けます。

次にビルド、または実行を押します。

## カスタム
正規表現の部分を弄れば、様々な形式のリストに対応できます。
