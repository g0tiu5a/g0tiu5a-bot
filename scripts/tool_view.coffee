# Description:
#   Tool view
#
# Commands:
#   tool <genre>       -- show tools filtered by genre
#   tool help          -- help this command
#
tools_json = require './data/tools.json'


module.exports = (robot) ->
  robot.hear /tool (.*)/, (msg) ->
    m = msg.match[1]
    if m in ["rev", "reverse"]
      genre = "reversing"
    else if m in ["stego"]
      genre = "stegano"
    else
      genre = m
    tools = tools_json[genre]
    unless tools?
      msg.send """
      ```
      * tool <genre>  -- ジャンル別にCTFで便利なツールを一覧にして表示する.
          * crypto    --  暗号問題向けのツール. 数学系ライブラリもここに収録
          * forensics --  フォレンジック問題向けのツール
          * misc      --  その他 雑種
          * network   --  ネットワーク問題に対応
          * stego / stegano
          * pwn
          * rev / reverse / reversing
          * web
       * tool help   -- this view
       ```"""
    else
      for tool in tools
        robot.emit 'slack.attachment',
          message: msg.message
          content: [
            {
              fallback: "#{genre}"
              title: "#{tool['name']}"
              title_link: "#{tool['url']}"
              mrkdwn_in: ["text"]
              text: "#{tool['description']}"
            }
          ]
