const e=location.pathname.split("/").slice(0,-1).join("/"),u=[e+"/internal/immutable/entry/app.BaMxkMDU.js",e+"/internal/immutable/nodes/0.a4j1ZdYF.js",e+"/internal/immutable/assets/0.DsL-Oas0.css",e+"/internal/immutable/nodes/1.BEfDo_fE.js",e+"/internal/immutable/nodes/2.DVFRQt-r.js",e+"/internal/immutable/nodes/3.BCRGhAq9.js",e+"/internal/immutable/nodes/4.vGWYHGeF.js",e+"/internal/immutable/nodes/5.Cxqp33ZF.js",e+"/internal/immutable/nodes/6.BgU-S-j6.js",e+"/internal/immutable/nodes/7.Dm-9Ejf1.js",e+"/internal/immutable/nodes/8.AqLcBBs7.js",e+"/internal/immutable/chunks/Dropdown.CcNYrred.js",e+"/internal/immutable/assets/Header.Ckn9IwLX.css",e+"/internal/immutable/chunks/Header.v5g8fz1y.js",e+"/internal/immutable/chunks/Icon.CkcrXJuI.js",e+"/internal/immutable/chunks/control.CYgJF_JY.js",e+"/internal/immutable/chunks/debounce.BzpYRo3t.js",e+"/internal/immutable/chunks/each.C8Rgkwia.js",e+"/internal/immutable/chunks/entry.qQdFD1p2.js",e+"/internal/immutable/chunks/httpAccess.TKbybTVh.js",e+"/internal/immutable/chunks/i18n-svelte.DgV7mMaz.js",e+"/internal/immutable/chunks/index.BupekJpF.js",e+"/internal/immutable/chunks/index.C2SR94LA.js",e+"/internal/immutable/chunks/index.DHw_VhY_.js",e+"/internal/immutable/chunks/index.De_1-CT6.js",e+"/internal/immutable/chunks/index.QYv70b3f.js",e+"/internal/immutable/chunks/localstorageWritable.DviM0pTl.js",e+"/internal/immutable/chunks/notifications.CARCvqH5.js",e+"/internal/immutable/chunks/paths.Cw_2S-rN.js",e+"/internal/immutable/chunks/preload-helper.D6kgxu3v.js",e+"/internal/immutable/chunks/scheduler.Bu9aYH8f.js",e+"/internal/immutable/chunks/util.BKKmp01Z.js",e+"/internal/immutable/entry/start.BfGCAbUe.js",e+"/internal/immutable/chunks/index.2su4BVk8.js"],m=[e+"/favicon.png",e+"/icon.png",e+"/icon256x256.png",e+"/manifest.json",e+"/robots.txt"],o="1720697987238",l=self,c=`cache-${o}`,r=[...u,...m];l.addEventListener("install",n=>{async function a(){await(await caches.open(c)).addAll(r)}n.waitUntil(a())});l.addEventListener("activate",n=>{async function a(){for(const t of await caches.keys())t!==c&&await caches.delete(t)}n.waitUntil(a())});l.addEventListener("fetch",n=>{if(n.request.method==="GET"){const a=(async()=>{const t=new URL(n.request.url),s=await caches.open(c);if(r.includes(t.pathname))return s.match(t.pathname);try{const i=await fetch(n.request);return i.status===200&&!n.request.url.startsWith("chrome-extension://")&&s.put(n.request,i.clone()),i}catch{return s.match(n.request)}})();n.respondWith(a)}});
