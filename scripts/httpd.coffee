# Description:
#   HTTP SERVER
module.exports = (robot) ->
  robot.router.get "/", (req, res) ->
    res.end robot.version
