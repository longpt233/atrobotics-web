import MessengerChatComponent from './MessengerChat.vue';

export function initFacebookSdk() {
  window.fbAsyncInit = function () {
    window.FB.init({
      xfbml: true,
      version: 'v13.0',
    });
  };

  // load facebook sdk script
  (function (d, s, id) {
    const fjs = d.getElementsByTagName(s)[0];
    if (d.getElementById(id)) return;
    const js = d.createElement(s);
    js.id = id;
    js.setAttribute(
      'src',
      'https://connect.facebook.net/vi_VN/sdk/xfbml.customerchat.js',
    );
    fjs?.parentNode?.insertBefore(js, fjs);
  })(document, 'script', 'facebook-jssdk');
}

initFacebookSdk();

export const MessengerChat = MessengerChatComponent;
