var __wpo = {"assets":{"main":["/scripts/bundle.3735cf55b3cbcd60abfd.js","/scripts/vendor.6379388a82551bfd7aed.js","/scripts/manifest.e1e1f42957bc20351787.js","/scripts/styles.css","/scripts/../","/scripts/vendor.6379388a82551bfd7aed.js.gz","/scripts/bundle.3735cf55b3cbcd60abfd.js.gz","/","https://fonts.googleapis.com/css?family=Roboto:300,400,500"],"additional":[],"optional":[]},"externals":["/","https://fonts.googleapis.com/css?family=Roboto:300,400,500"],"hashesMap":{"fc186d116d9a1ebeae55a56d1fdbda68283c0114":"/scripts/bundle.3735cf55b3cbcd60abfd.js","5fada54e9dae3976c1b1ad000275085ecf12f3e4":"/scripts/vendor.6379388a82551bfd7aed.js","1540135e7cc3d6680987061ad2358577636688f8":"/scripts/manifest.e1e1f42957bc20351787.js","a5b12874b28e999073e01269fdd025b7f7a33ce8":"/scripts/styles.css","3b2791e6166b47a9996af424ed2d0837ff2f835e":"/scripts/../","e82313ddde678fedb206a895d11c0cde40ab44e7":"/scripts/vendor.6379388a82551bfd7aed.js.gz","e16523842b41ba902b68a0a8cf324cf148c5182b":"/scripts/bundle.3735cf55b3cbcd60abfd.js.gz"},"navigateFallbackURL":"/","navigateFallbackForRedirects":true,"strategy":"changed","responseStrategy":"cache-first","version":"2018-1-26 02:06:32","name":"webpack-offline","pluginVersion":"4.8.5","relativePaths":false};

!function(e){function n(r){if(t[r])return t[r].exports;var o=t[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}var t={};n.m=e,n.c=t,n.d=function(e,t,r){n.o(e,t)||Object.defineProperty(e,t,{configurable:!1,enumerable:!0,get:r})},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,n){return Object.prototype.hasOwnProperty.call(e,n)},n.p="/scripts/",n(n.s=0)}([function(e,n,t){"use strict";function r(e,n){return caches.match(e,{cacheName:n}).then(function(t){return c()?t:a(t).then(function(t){return caches.open(n).then(function(n){return n.put(e,t)}).then(function(){return t})})}).catch(function(){})}function o(e,n){return e+(-1!==e.indexOf("?")?"&":"?")+"__uncache="+encodeURIComponent(n)}function i(e){return"navigate"===e.mode||e.headers.get("Upgrade-Insecure-Requests")||-1!==(e.headers.get("Accept")||"").indexOf("text/html")}function c(e){return!e||!e.redirected||!e.ok||"opaqueredirect"===e.type}function a(e){return c(e)?Promise.resolve(e):("body"in e?Promise.resolve(e.body):e.blob()).then(function(n){return new Response(n,{headers:e.headers,status:e.status})})}function s(e){return Object.keys(e).reduce(function(n,t){return n[t]=e[t],n},{})}function u(e,n){console.groupCollapsed("[SW]:",e),n.forEach(function(e){console.log("Asset:",e)}),console.groupEnd()}if(function(){var e=ExtendableEvent.prototype.waitUntil,n=FetchEvent.prototype.respondWith,t=new WeakMap;ExtendableEvent.prototype.waitUntil=function(n){var r=this,o=t.get(r);return o?void o.push(Promise.resolve(n)):(o=[Promise.resolve(n)],t.set(r,o),e.call(r,Promise.resolve().then(function e(){var n=o.length;return Promise.all(o.map(function(e){return e.catch(function(){})})).then(function(){return o.length!=n?e():(t.delete(r),Promise.all(o))})})))},FetchEvent.prototype.respondWith=function(e){return this.waitUntil(e),n.call(this,e)}}(),void 0===f)var f=!1;!function(e,n){function t(){if(!W.additional.length)return Promise.resolve();f&&console.log("[SW]:","Caching additional");var e=void 0;return e="changed"===b?l("additional"):c("additional"),e.catch(function(e){console.error("[SW]:","Cache section `additional` failed to load")})}function c(n){var t=W[n];return caches.open(E).then(function(n){return w(n,t,{bust:e.version,request:e.prefetchRequest})}).then(function(){u("Cached assets: "+n,t)}).catch(function(e){throw console.error(e),e})}function l(n){return d().then(function(t){if(!t)return c(n);var r=t[0],o=t[1],i=t[2],a=i.hashmap,s=i.version;if(!i.hashmap||s===e.version)return c(n);var f=Object.keys(a).map(function(e){return a[e]}),l=o.map(function(e){var n=new URL(e.url);return n.search="",n.hash="",n.toString()}),h=W[n],d=[],p=h.filter(function(e){return-1===l.indexOf(e)||-1===f.indexOf(e)});Object.keys(R).forEach(function(e){var n=R[e];if(-1!==h.indexOf(n)&&-1===p.indexOf(n)&&-1===d.indexOf(n)){var t=a[e];t&&-1!==l.indexOf(t)?d.push([t,n]):p.push(n)}}),u("Changed assets: "+n,p),u("Moved assets: "+n,d);var v=Promise.all(d.map(function(e){return r.match(e[0]).then(function(n){return[e[1],n]})}));return caches.open(E).then(function(n){var t=v.then(function(e){return Promise.all(e.map(function(e){return n.put(e[0],e[1])}))});return Promise.all([t,w(n,p,{bust:e.version,request:e.prefetchRequest})])})})}function h(){return caches.keys().then(function(e){var n=e.map(function(e){if(0===e.indexOf(q)&&0!==e.indexOf(E))return console.log("[SW]:","Delete cache:",e),caches.delete(e)});return Promise.all(n)})}function d(){return caches.keys().then(function(e){for(var n=e.length,t=void 0;n--&&(t=e[n],0!==t.indexOf(q)););if(t){var r=void 0;return caches.open(t).then(function(e){return r=e,e.match(new URL(j,location).toString())}).then(function(e){if(e)return Promise.all([r,r.keys(),e.json()])})}})}function p(){return caches.open(E).then(function(n){var t=new Response(JSON.stringify({version:e.version,hashmap:R}));return n.put(new URL(j,location).toString(),t)})}function v(e,n,t){return r(t,E).then(function(r){return r?(f&&console.log("[SW]:","URL ["+t+"]("+n+") from cache"),r):fetch(e.request).then(function(r){return r.ok?(f&&console.log("[SW]:","URL ["+n+"] from network"),t===n&&function(){var t=r.clone(),o=caches.open(E).then(function(e){return e.put(n,t)}).then(function(){console.log("[SW]:","Cache asset: "+n)});e.waitUntil(o)}(),r):(f&&console.log("[SW]:","URL ["+n+"] wrong response: ["+r.status+"] "+r.type),r)})})}function g(e,n,t){return fetch(e.request).then(function(e){if(e.ok)return f&&console.log("[SW]:","URL ["+n+"] from network"),e;throw new Error("Response is not ok")}).catch(function(){return f&&console.log("[SW]:","URL ["+n+"] from cache if possible"),r(t,E)})}function m(e){return e.catch(function(){}).then(function(e){var n=e&&e.ok,t=e&&"opaqueredirect"===e.type;return n||t&&!F?e:(f&&console.log("[SW]:","Loading navigation fallback ["+C+"] from cache"),r(C,E))})}function w(e,n,t){var r=!1!==t.allowLoaders,i=t&&t.bust,c=t.request||{credentials:"omit",mode:"cors"};return Promise.all(n.map(function(e){return i&&(e=o(e,i)),fetch(e,c).then(a)})).then(function(o){if(o.some(function(e){return!e.ok}))return Promise.reject(new Error("Wrong response status"));var i=[],c=o.map(function(t,o){return r&&i.push(y(n[o],t)),e.put(n[o],t)});return i.length?function(){var r=s(t);r.allowLoaders=!1;var o=c;c=Promise.all(i).then(function(t){var i=[].concat.apply([],t);return n.length&&(o=o.concat(w(e,i,r))),Promise.all(o)})}():c=Promise.all(c),c})}function y(e,n){var t=Object.keys(L).map(function(t){if(-1!==L[t].indexOf(e)&&O[t])return O[t](n.clone())}).filter(function(e){return!!e});return Promise.all(t).then(function(e){return[].concat.apply([],e)})}function x(e){var n=e.url,t=new URL(n),r=void 0;r="navigate"===e.mode?"navigate":t.origin===location.origin?"same-origin":"cross-origin";for(var o=0;o<k.length;o++){var i=k[o];if(i&&(!i.requestTypes||-1!==i.requestTypes.indexOf(r))){var c=void 0;if((c="function"==typeof i.match?i.match(t,e):n.replace(i.match,i.to))&&c!==n)return c}}}var O=n.loaders,k=n.cacheMaps,b=e.strategy,S=e.responseStrategy,W=e.assets,L=e.loaders||{},R=e.hashesMap,U=e.externals,q=e.name,P=e.version,E=q+":"+P,j="__offline_webpack__data";!function(){Object.keys(W).forEach(function(e){W[e]=W[e].map(function(e){var n=new URL(e,location);return n.hash="",-1===U.indexOf(e)&&(n.search=""),n.toString()})}),Object.keys(L).forEach(function(e){L[e]=L[e].map(function(e){var n=new URL(e,location);return n.hash="",-1===U.indexOf(e)&&(n.search=""),n.toString()})}),R=Object.keys(R).reduce(function(e,n){var t=new URL(R[n],location);return t.search="",t.hash="",e[n]=t.toString(),e},{}),U=U.map(function(e){var n=new URL(e,location);return n.hash="",n.toString()})}();var _=[].concat(W.main,W.additional,W.optional),C=e.navigateFallbackURL,F=e.navigateFallbackForRedirects;self.addEventListener("install",function(e){console.log("[SW]:","Install event");var n=void 0;n="changed"===b?l("main"):c("main"),e.waitUntil(n)}),self.addEventListener("activate",function(e){console.log("[SW]:","Activate event");var n=t();n=n.then(p),n=n.then(h),n=n.then(function(){if(self.clients&&self.clients.claim)return self.clients.claim()}),e.waitUntil(n)}),self.addEventListener("fetch",function(e){var n=new URL(e.request.url);n.hash="";var t=n.toString();-1===U.indexOf(t)&&(n.search="",t=n.toString());var r="GET"===e.request.method,o=-1!==_.indexOf(t),c=t;if(!o){var a=x(e.request);a&&(c=a,o=!0)}if(!o&&r&&C&&i(e.request))return void e.respondWith(m(fetch(e.request)));if(!o||!r)return void(n.origin!==location.origin&&-1!==navigator.userAgent.indexOf("Firefox/44.")&&e.respondWith(fetch(e.request)));var s=void 0;s="network-first"===S?g(e,t,c):v(e,t,c),C&&i(e.request)&&(s=m(s)),e.respondWith(s)}),self.addEventListener("message",function(e){var n=e.data;if(n)switch(n.action){case"skipWaiting":self.skipWaiting&&self.skipWaiting()}})}(__wpo,{loaders:{},cacheMaps:[]}),e.exports=t(1)},function(e,n){self.addEventListener("fetch",function(e){e.respondWith(caches.open("reddit-url-cache-v1").then(function(n){return fetch(e.request).then(function(t){return n.put(e.request,t.clone()),t})}))})}]);