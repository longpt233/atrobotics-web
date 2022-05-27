import ZaloChatComponent from './ZaloChat.vue';

export function initZalo() {
  // load zalo sdk script
  (function (d, s, id) {
    const fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) {
      return;
    }
    const js = d.createElement(s);
    js.id = id;
    js.setAttribute('src', 'https://sp.zalo.me/plugins/sdk.js');
    fjs.parentNode?.insertBefore(js, fjs);
  })(document, 'script', 'zalo-jssdk');
}

initZalo();

export const ZaloChat = ZaloChatComponent;
