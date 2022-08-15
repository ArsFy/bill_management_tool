process.env.NTBA_FIX_319 = 1;
const config = require('./config.json');

const TelegramBot = require('node-telegram-bot-api');
const bot = new TelegramBot(config.bot_token, { polling: true });

bot.on("polling_error", console.log);

bot.on('message', async (msg) => {
    const chatId = msg.chat.id;
    let msgInfo = msg.text != undefined ? msg.text.split(" ") : "";

    switch (msgInfo[0]) {
        case "/start": case `/start${config.bot_name}`:
            bot.sendMessage(chatId, "BMTool Bot")
    }
})