module.exports = (robot) ->
  robot.hear /test/, (msg) ->
    robot.emit 'slack.attachment',
    # msg.messageは必須
    message: msg.message
    content: [{
        # see https://api.slack.com/docs/attachments
        text: "Attachment text"
        fallback: "Attachment fallback"
        fields: [{
            title: "Field title"
            value: "Field value"
        }]
    }]
