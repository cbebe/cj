if(!self.define){let e,i={};const t=(t,r)=>(t=new URL(t+".js",r).href,i[t]||new Promise((i=>{if("document"in self){const e=document.createElement("script");e.src=t,e.onload=i,document.head.appendChild(e)}else e=t,importScripts(t),i()})).then((()=>{let e=i[t];if(!e)throw new Error(`Module ${t} didn’t register its module`);return e})));self.define=(r,s)=>{const n=e||("document"in self?document.currentScript.src:"")||location.href;if(i[n])return;let o={};const c=e=>t(e,n),d={module:{uri:n},exports:o,require:c};i[n]=Promise.all(r.map((e=>d[e]||c(e)))).then((e=>(s(...e),o)))}}define(["./workbox-1f8c92ee"],(function(e){"use strict";self.addEventListener("message",(e=>{e.data&&"SKIP_WAITING"===e.data.type&&self.skipWaiting()})),e.precacheAndRoute([{url:"data.js",revision:"9b017a7cf4c35719eeb2d58ca49ee010"},{url:"index.html",revision:"9e366e27770ec906ea0892eeaa287ce8"},{url:"main.js",revision:"393ddc09dfb6254b125d70223447f8dc"}],{ignoreURLParametersMatching:[/^utm_/,/^fbclid$/]})}));
//# sourceMappingURL=sw.js.map
