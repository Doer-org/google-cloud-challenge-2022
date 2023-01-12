# Google Cloud Challenge 2022

## client

- コンポーネント責務の分け方

```
components
|
|
| --- atoms
      // atomsはUIの最小単位、処理は持たず、propsで内部は管理
|
|
| --- molecules
      // moleculesは複数atomsから構成される単位、ドメインに属する、処理は持たない（持つ場合もある）
|
|
| --- organisms
      // organismsは複数moleculesから構成される単位、処理は持つ（もしかしたらorganisms自体がいらないかも）
|
|
| --- templates
      // ドメインごとに何か一部分を変更する時とかに使いそう（今のところ使うところは見えてない）
   |
   |
   | --- shared
         // headerとかどこでも使うもの

```

## 環境

docker-compose環境立ち上げ

TODO: なぜか2回目のmake upで動く

```
make up
```

コマンドのヘルプについて

```
make or make help
```

