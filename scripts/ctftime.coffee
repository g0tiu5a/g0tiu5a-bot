# Description:
#   CTFtime
#
# Commands:
#   * ctftime event
#      show events between now and 7 days later limits 3
CTFTIME_API_URL = "https://ctftime.org/api/v1"
moment = require "moment"

# sleep function from https://jsfiddle.net/axjw668v/8/
sleep = (delay) ->
  start = new Date().getTime()
  while new Date().getTime() < start + delay
    true

module.exports = (robot) ->
  robot.hear /ctftime event/, (msg) ->
    limit = 3
    start = new Date().getTime()
    start = Math.floor(start / 1000)
    finish = Math.floor(finish / 1000)
    # 最大直近3つのイベントを表示する
    url = CTFTIME_API_URL + "/events/?limit=#{limit}&start=#{start}"
    http = msg.http url
    http.get() (err, res, body) ->
      if res.statusCode is 404
        msg.send "404: nothing event"
      else if JSON.parse body == []
        msg.send "Nothing event"
      else
        events = JSON.parse body
        for event in events
          # Hide offline event
          if event.location == ""
            sleep(500) # 0.5
            robot.emit 'slack.attachment',
            message: msg.message
            content: [
              {
                fallback: "#{event.title} - #{event.url}"
                title: event.title
                title_link: event.url

                fields: [
                  {
                    title: "format"
                    value: event.format
                    short: true
                  },
                  {
                    title: "weight"
                    value: event.weight
                    short: true
                  },
                  {
                    title: "start"
                    value: moment(event.start).format("YYYY/MM/DD HH:mm:ssZ")
                    short: true
                  },
                  {
                    title: "finish"
                    value: moment(event.finish).format("YYYY/MM/DD HH:mm:ssZ")
                    short: true
                  }
                ]
                color: "#F35A00"
              }
            ]

  robot.hear /ctftime ours/, (msg) ->
    OUR_TEAM_ID = 16931  # Team g0tiu5a ID

    url = CTFTIME_API_URL + "/teams/#{OUR_TEAM_ID}/"
    http = msg.http url
    http.get() (err, res, body) ->
      if res.statusCode is 404
        msg.send "404: nothing team"
      else
        team = JSON.parse body
        latest_ratings = team.rating[0][new Date().getFullYear().toString()]
        robot.emit 'slack.attachment',
        message: msg.message
        content: [
          {
            fallback: "#{team.name}"
            title: team.name

            fields: [
              {
                title: "Latest rating points"
                value: latest_ratings.rating_points
                short: true
              },
              {
                title: "Latest raring place"
                value: "#{latest_ratings.rating_place} th"
                short: true
              },
              {
                title: "country"
                value: ":flag-#{team.country}:"
                short: true
              }
            ]
            color: "#764FA5"
          }
        ]
