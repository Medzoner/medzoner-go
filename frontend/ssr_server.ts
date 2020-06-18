import * as Vue from 'vue';
import * as VueRender from 'vue-server-renderer';
import * as express from 'express';

async function bootstrap() {

    const server = express();
    const context = {
        title: 'vue ssr',
        metas: `
        <meta name="keyword" content="vue,ssr">
        <meta name="description" content="vue srr demo">
    `,
    };

    let render = VueRender.createRenderer({
        template: require('fs').readFileSync('./frontend/index.template.html', 'utf-8')
    });

    server.get('/', async (req, res) => {
        let tpl = `<div id="#app">The visited URL is: {{ url }}</div>`;
        let app = new Vue({
            data: {
                url: req.url
            },
            template: tpl,
        });

        await render.renderToString(app, context,(err: Error, html: string) => {
            if (err) {
                res.status(500).end('Internal Server Error')
                return;
            }

            res.end(html);
        });
    });

    server.listen(3666);

}
bootstrap();