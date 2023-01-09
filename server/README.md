# server

## API document

基本的にはopenapiを参考にしてください。

ここにはよく使われるであろうAPIと一部補足が書いてあります。

#### event作成 POST /events
name,adminが必須です。

下記が参考です。

```json
{
    "name":"hoge",
    "admin":"3f94bd72-1ec6-43e0-a1bf-cf5601016cf4",
}
```

作成したイベントがresponseとして返ってきます。

注 : adminはこのAPIをたたいたタイミングでイベントに参加したとみなされます。
つまり、他のAPIを用いてadminをeventに参加させる処理を行う必要はありません。

#### event取得 GET /events/{id}
idでイベントを取得します

このidが共有するURLに使われます。


#### ユーザーのイベント参加 POST /events/{id}/participants
nameのみが必須です。

commentがある場合は、以下のようにrequest bodyに入れてください。

```json
{
    "name":"hoge",
    "comment":"らーめん好き",
}
```

#### eventの状態を更新する PATCH /events/{id}/state

いまdocumentを書いててとんでもないことに気づきました。

event参加可能数を保持するプロパティを用意してませんでした。

明日やるのでお待ちください。

話を戻して、このAPIはstateのみが必須です。
cancelかcloseを送って状態を更新できます。


<br><br>

ここが厄介です。

#### eventの参加者を取得する GET /events/{id}/users

このAPIで参加者を取得できますが、同時にコメントは取得できません。

#### eventの参加者のコメントを取得する GET /events/{id}/comments

このAPIを同時にたたいて、コメント一覧も取得してください。

