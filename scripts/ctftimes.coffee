# Description:
#   CTFtimes
#
# Commands:
# ctftimes event
CTFTIMES_API_URL = "https://ctftime.org/api/v1"

module.exports = (robot) ->
  robot.hear /ctftimes event/, (msg) ->
    limit = 3
    start = new Date().getTime()
    finish = start + 604800000   # 7 days later

    start = Math.floor(start / 1000)
    finish = Math.floor(finish / 1000)
    url = CTFTIMES_API_URL + "/events/?limit=#{limit}&start=#{start}&finish=#{finish}"
    msg.send url
    http = msg.http url
    http.get() (err, res, body) ->
      if res.statusCode is 404
        msg.send "404: nothing event"
      else
        # msg.send body
        jsons = JSON.parse body

        robot.emit 'slack.attachment', {text:"<https://github.com/link/to/a/PR|myrepo #42> fix some broken"}
