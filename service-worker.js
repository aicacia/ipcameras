const e=location.pathname.split("/").slice(0,-1).join("/"),r=[e+"/internal/immutable/entry/app.B9UWvITe.js",e+"/internal/immutable/nodes/0.BcbF22_h.js",e+"/internal/immutable/assets/0.DsL-Oas0.css",e+"/internal/immutable/nodes/1.D2ae2oEC.js",e+"/internal/immutable/nodes/2.Dq4mETqI.js",e+"/internal/immutable/nodes/3.C5KbBJD0.js",e+"/internal/immutable/nodes/4.DDjIQR1J.js",e+"/internal/immutable/nodes/5.BIbr7zkr.js",e+"/internal/immutable/nodes/6.BgU-S-j6.js",e+"/internal/immutable/nodes/7.DmupO-oi.js",e+"/internal/immutable/nodes/8.D2ObsjCu.js",e+"/internal/immutable/chunks/Dropdown.CcNYrred.js",e+"/internal/immutable/assets/Header.Ckn9IwLX.css",e+"/internal/immutable/chunks/Header.bRLdPbOq.js",e+"/internal/immutable/chunks/Icon.CkcrXJuI.js",e+"/internal/immutable/chunks/control.CYgJF_JY.js",e+"/internal/immutable/chunks/debounce.BzpYRo3t.js",e+"/internal/immutable/chunks/each.C8Rgkwia.js",e+"/internal/immutable/chunks/entry.DIDtGCxd.js",e+"/internal/immutable/chunks/i18n-svelte.DgV7mMaz.js",e+"/internal/immutable/chunks/index.BupekJpF.js",e+"/internal/immutable/chunks/index.DHw_VhY_.js",e+"/internal/immutable/chunks/index.De_1-CT6.js",e+"/internal/immutable/chunks/index.QYv70b3f.js",e+"/internal/immutable/chunks/index.av7PPCqZ.js",e+"/internal/immutable/chunks/localAccess.CNYJS0uA.js",e+"/internal/immutable/chunks/localstorageWritable.DviM0pTl.js",e+"/internal/immutable/chunks/notifications.CARCvqH5.js",e+"/internal/immutable/chunks/paths.Dl7yvK5v.js",e+"/internal/immutable/chunks/preload-helper.D6kgxu3v.js",e+"/internal/immutable/chunks/scheduler.Bu9aYH8f.js",e+"/internal/immutable/chunks/util.BKKmp01Z.js",e+"/internal/immutable/entry/start.RoLpwsRk.js",e+"/internal/immutable/chunks/index.qs0H3y-J.js"],m=[e+"/favicon.png",e+"/icon.png",e+"/icon256x256.png",e+"/manifest.json",e+"/robots.txt"],o="1720381911676",l=self,c=`cache-${o}`,u=[...r,...m];l.addEventListener("install",n=>{async function a(){await(await caches.open(c)).addAll(u)}n.waitUntil(a())});l.addEventListener("activate",n=>{async function a(){for(const t of await caches.keys())t!==c&&await caches.delete(t)}n.waitUntil(a())});l.addEventListener("fetch",n=>{if(n.request.method==="GET"){const a=(async()=>{const t=new URL(n.request.url),s=await caches.open(c);if(u.includes(t.pathname))return s.match(t.pathname);try{const i=await fetch(n.request);return i.status===200&&!n.request.url.startsWith("chrome-extension://")&&s.put(n.request,i.clone()),i}catch{return s.match(n.request)}})();n.respondWith(a)}});
