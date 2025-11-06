# puppeteer-stream start error

Error: net::ERR_BLOCKED_BY_CLIENT at chrome-extension://jjndjgheafjngoipoacpjgeicjeomjli/options.html#55200
at navigate (D:\v5livepuppeteer\node_modules\puppeteer-core\lib\cjs\puppeteer\cdp\Frame.js:186:27)
at async Deferred.race (D:\v5livepuppeteer\node_modules\puppeteer-core\lib\cjs\puppeteer\util\Deferred.js:36:20)
at async CdpFrame.goto (D:\v5livepuppeteer\node_modules\puppeteer-core\lib\cjs\puppeteer\cdp\Frame.js:152:25)
at async CdpPage.goto (D:\v5livepuppeteer\node_modules\puppeteer-core\lib\cjs\puppeteer\api\Page.js:588:20)

update puppeteer-stream 3.0.22

and

# puppeteer-stream captured error
3.0.22 - Error during recording: Extension has not been invoked for the current page (see activeTab permission). Chrome pages cannot be captured.

args: [
'--allowlisted-extension-id=jjndjgheafjngoipoacpjgeicjeomjli'
],


# linux web puppeteer-stream record 字体问题
需要安装字体
sudo apt install -y fonts-noto-cjk fonts-noto-color-emoji fonts-liberation fonts-dejavu-core