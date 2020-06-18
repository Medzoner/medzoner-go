"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const vue_1 = require("vue");
const vue_server_renderer_1 = require("vue-server-renderer");
const express_1 = require("express");
async function bootstrap() {
    const server = express_1.default();
    const context = {
        title: 'vue ssr',
        metas: `
        <meta name="keyword" content="vue,ssr">
        <meta name="description" content="vue srr demo">
    `,
    };
    let tpl = `<div>The visited URL is: {{ url }}</div>`;
    console.log(tpl);
    let render = vue_server_renderer_1.default.createRenderer();
    server.get('/', async (req, res) => {
        let app = new vue_1.default({
            data: {
                url: req.url
            },
            template: tpl,
        });
        await render.renderToString(app, context, (err, html) => {
            console.log(html);
            if (err) {
                return;
            }
            res.end(html);
        });
    });
    server.listen(3000);
}
bootstrap();
//# sourceMappingURL=main.js.map