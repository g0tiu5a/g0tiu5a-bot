# Description:
#   CTFtimes
#
# Commands:
# ctftimes event
CTFTIMES_API_URL = "https://ctftime.org/api/v1"
moment = require "moment"

module.exports = (robot) ->
  robot.hear /ctftimes event/, (msg) ->
    limit = 3
    start = new Date().getTime()
    finish = start + 604800000   # 7 days later

    start = Math.floor(start / 1000)
    finish = Math.floor(finish / 1000)
    url = CTFTIMES_API_URL + "/events/?limit=#{limit}&start=#{start}&finish=#{finish}"
    http = msg.http url
    http.get() (err, res, body) ->
      if res.statusCode is 404
        msg.send "404: nothing event"
      else
        # msg.send body
        events = JSON.parse body
        for event in events
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
