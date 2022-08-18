process.env.NTBA_FIX_319 = 1;
const config = require('./config.json');

const qs = require("qs");
const axios = require("axios");
const TelegramBot = require('node-telegram-bot-api');
const bot = new TelegramBot(config.bot_token, { polling: true });

bot.on("polling_error", console.log);

let use = "";
let help = `list              - List all obj
create <obj_name> - Create a new obj
use <obj_name>    - Set default use obj
add [Operator Name] [+123 or -123] [what is it used for] [Timestamp (optional)] - Add Obj`

bot.on('message', async (msg) => {
    const chatId = msg.chat.id;
    let msgInfo = msg.text != undefined ? msg.text.split(" ") : "";

    switch (msgInfo[0]) {
        case "/start": case `/start${config.bot_name}`:
            bot.sendMessage(chatId, "BMTool Bot\n" + help);
            break;
        case "/list": case `/list${config.bot_name}`:
            axios({
                method: "POST",
                url: config.server + "/api/obj_list"
            }).then(res => {
                if (res.data.count != 0) {
                    bot.sendMessage(chatId, `Use '/use <obj_name>' to select a obj\n${res.data.list.join(" ")}`)
                } else {
                    bot.sendMessage(chatId, `obj list is empty, use command '/create obj_name'`)
                }
            }).catch(e => {
                bot.sendMessage(chatId, `Server Err`)
            })
            break;
        case "/create": case `/create${config.bot_name}`:
            if (msgInfo.length == 2 && msgInfo[1] != "") {
                axios({
                    method: "POST",
                    url: config.server + "/api/create_obj",
                    data: qs.stringify({ "name": msgInfo[1] })
                }).then(res => {
                    if (res.data.status == 200) {
                        switch (res.data.info) {
                            case "ok":
                                bot.sendMessage(chatId, `Created successfully: ${msgInfo[1]}`)
                                break;
                            case "isExist":
                                bot.sendMessage(chatId, `Created Error: ${msgInfo[1]} is Exist`)
                                break
                        }
                    }
                }).catch(e => {
                    bot.sendMessage(chatId, `Server Err`)
                })
            } else {
                bot.sendMessage(chatId, `missing name: '/create obj_name'`)
            }
            break
        case "/use": case `/use${config.bot_name}`:
            if (msgInfo.length == 2 && msgInfo[1] != "") {
                axios({
                    method: "POST",
                    url: config.server + "/api/is_obj_exist",
                    data: qs.stringify({ "name": msgInfo[1] })
                }).then(res => {
                    if (res.data.isExist) {
                        use = msgInfo[1]
                        bot.sendMessage(chatId, `Select: ${msgInfo[1]}`)
                    } else {
                        bot.sendMessage(chatId, `Select Err: ${msgInfo[1]} is not exist, use '/create obj_name'`)
                    }
                }).catch(e => {
                    bot.sendMessage(chatId, `Server Err`)
                })
            } else {
                bot.sendMessage(chatId, `input value missing: '/use obj_name'`)
            }
            break
        case "/add": case `/add${config.bot_name}`:
            if (msgInfo.length > 3 && msgInfo[1] != "" && msgInfo[2] != "" && msgInfo[3] != "") {
                if (use != "") {
                    axios({
                        method: "POST",
                        url: config.server + "/api/add_record",
                        data: qs.stringify({ "name": use, "operator": msgInfo[1], "change": msgInfo[2], "comment": msgInfo[3], "time": msgInfo.length == 4 ? Math.ceil(new Date() / 1000) : (isNaN(Number(msgInfo[4])) ? Math.ceil(new Date() / 1000) : Number(msgInfo[4])) })
                    }).then(res => {
                        if (res.data.status == 200) {
                            bot.sendMessage(chatId, `Add successfully`)
                        } else {
                            bot.sendMessage(chatId, `Add Err`)
                        }
                    }).catch(e => {
                        bot.sendMessage(chatId, `Server Err`)
                    })
                } else {
                    bot.sendMessage(chatId, `must be selected an obj to use: '/use obj_name'`)
                }
            } else {
                bot.sendMessage(chatId, `input value missing: '/add [Operator Name] [+123 or -123] [what is it used for] [Timestamp (optional)]'`)
            }
            break
    }
})

console.log("Start bot...")